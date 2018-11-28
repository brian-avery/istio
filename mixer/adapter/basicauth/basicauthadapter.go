// nolint: lll
//go:generate $GOPATH/src/istio.io/istio/bin/mixer_codegen.sh -a mixer/adapter/basicauth/config/config.proto -x "-n basicauthadapter -t authorization"
package adapter

import (
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"strings"
	"time"

	"istio.io/istio/mixer/adapter/basicauth/authenticators"

	"github.com/gogo/googleapis/google/rpc"
	"google.golang.org/grpc"

	"istio.io/api/mixer/adapter/model/v1beta1"
	"istio.io/istio/mixer/adapter/basicauth/config"
	"istio.io/istio/mixer/pkg/status"
	"istio.io/istio/mixer/template/authorization"
	"istio.io/istio/pkg/log"
)

type (
	// Server is basic server interface
	Server interface {
		Addr() string
		Close() error
		Run(shutdown chan error)
	}

	// BasicAuthAdapter supports authorization template.
	BasicAuthAdapter struct {
		listener               net.Listener
		server                 *grpc.Server
		basicAuthAuthenticator *authenticators.BasicAuthAuthenticator
	}
)

type authenticator interface {
	Authenticate(token string) (bool, error)
}

const (
	SchemaBasic = "basic"
)

var (
	ValidDuration       = 5 * time.Second
	ValidUseCount int32 = 1000
)

var _ authorization.HandleAuthorizationServiceServer = &BasicAuthAdapter{}

func (s *BasicAuthAdapter) getCheckResult(rpcStatus rpc.Status) *v1beta1.CheckResult {
	return &v1beta1.CheckResult{
		Status:        rpcStatus,
		ValidDuration: ValidDuration,
		ValidUseCount: ValidUseCount,
	}
}

//HandleAuthorization records authorization requests
func (s *BasicAuthAdapter) HandleAuthorization(ctx context.Context, r *authorization.HandleAuthorizationRequest) (*v1beta1.CheckResult, error) {
	log.Infof("HandleAuthorization received request: %+v\n\n", *r)

	if r == nil || r.Instance == nil || r.Instance.Subject == nil {
		return s.getCheckResult(status.New(rpc.UNAUTHENTICATED)), fmt.Errorf("no authorization info present")
	}

	cfg := &config.Params{}
	if r.AdapterConfig != nil {
		if err := cfg.Unmarshal(r.AdapterConfig.Value); err != nil {
			log.Errorf("error unmarshaling adapter config: %v", err)
			return s.getCheckResult(status.New(rpc.UNAUTHENTICATED)), err
		}
	}

	log.Infof("Received adapter config: %+v\n", cfg)

	user := r.Instance.Subject.User
	token, err := getBasicToken(user)
	if err != nil {
		return s.getCheckResult(status.New(rpc.UNAUTHENTICATED)), err
	}

	user, password, err := getTokenSegments(string(token))
	if err != nil {
		log.Errorf("unable to get token segments: %s", err.Error())
		return s.getCheckResult(status.New(rpc.UNAUTHENTICATED)), err
	}

	if s.basicAuthAuthenticator == nil {
		cfg.Htpasswd = "/volume/htpasswd"
		s.basicAuthAuthenticator, err = authenticators.NewBasicAuthAdapter(cfg.Htpasswd)
		if err != nil {
			log.Errorf("unable to create basic auth adapter: %s", err.Error())
			return s.getCheckResult(status.New(rpc.UNAUTHENTICATED)), err
		}
	}

	if ok := s.basicAuthAuthenticator.Validate(user, password); ok {
		return s.getCheckResult(status.OK), nil
	}
	return s.getCheckResult(status.New(rpc.UNAUTHENTICATED)), nil
}

func getBasicToken(user string) (string, error) {
	firstSpace := strings.Index(user, " ")

	if firstSpace == -1 {
		return "", fmt.Errorf("missing authorization schema")
	}

	schema := strings.ToLower(user[:firstSpace])
	token := strings.TrimSpace(user[firstSpace:])
	if schema != SchemaBasic {
		return "", fmt.Errorf("invalid basic auth token")
	}
	return token, nil
}

func getTokenSegments(token string) (string, string, error) {
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Errorf("unable to decode basic token: %s", err.Error())
		return "", "", fmt.Errorf("unable to decode basic token: %s", err.Error())
	}

	tokenSegments := strings.Split(string(decodedToken), ":")
	if len(tokenSegments) != 2 {
		return "", "", fmt.Errorf("token was not a base64 encoded username and password")
	}
	return tokenSegments[0], tokenSegments[1], nil
}

// Addr returns the listening address of the server
func (s *BasicAuthAdapter) Addr() string {
	return s.listener.Addr().String()
}

// Run starts the server run
func (s *BasicAuthAdapter) Run(shutdown chan error) {
	shutdown <- s.server.Serve(s.listener)
}

// Close gracefully shuts down the server; used for testing
func (s *BasicAuthAdapter) Close() error {
	if s.server != nil {
		s.server.GracefulStop()
	}

	if s.listener != nil {
		_ = s.listener.Close()
	}

	return nil
}

// NewBasicAuthAdapter creates a new adapter that listens at provided port.
func NewBasicAuthAdapter(addr string) (Server, error) {
	if addr == "" {
		addr = "0"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", addr))
	if err != nil {
		return nil, fmt.Errorf("unable to listen on socket: %v", err)
	}
	s := &BasicAuthAdapter{
		listener: listener,
	}
	fmt.Printf("listening on \"%v\"\n", s.Addr())
	s.server = grpc.NewServer()
	authorization.RegisterHandleAuthorizationServiceServer(s.server, s)
	return s, nil
}
