package authenticators

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fsnotify/fsnotify"
	"istio.io/istio/pkg/log"
)

type BasicAuthAuthenticator struct {
	HTPasswdContent string
}

func NewBasicAuthAdapter(path string) (*BasicAuthAuthenticator, error) {
	data, err := getFileContents(path)
	if err != nil {
		return nil, fmt.Errorf("could not read htpasswd file: %s", err.Error())
	}
	adapter := &BasicAuthAuthenticator{
		HTPasswdContent: data,
	}
	//adapter.monitor(path)
	return adapter, nil
}

//Authenticate accepts
func (auth BasicAuthAuthenticator) Authenticate(token string) (success bool, err error) {
	// basic authentication stores tokens in the format b64(user:password).
	// HTPasswd uses user:md5(password). Transform to match.
	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Errorf("unable to decode basic token: %s", err.Error())
		return false, err
	}
	log.Infof("Have decoded token: %s", decodedToken)

	//break on colon
	tokenSegments := strings.Split(string(decodedToken), ":")
	if len(tokenSegments) != 2 {
		log.Errorf("token was not of the format base64(user:password)")
		return false, fmt.Errorf("token was not of the format base64(user:password)")
	}

	log.Infof("Segments retrieved: %+v", tokenSegments)

	passwdSum := md5.Sum([]byte(tokenSegments[1]))
	hash := tokenSegments[0] + ":" + fmt.Sprintf("%x", passwdSum)
	log.Infof("Got token: %s Hash: %s\n", token, hash)

	//check to see if the htpasswd contains the credentials.
	if strings.Contains(auth.HTPasswdContent, hash) {
		log.Infof("Authenticated successfully.\n")
		return true, nil
	}
	return false, nil
}

func getFileContents(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("unable to read changes to HTPassword file")
	}

	log.Infof("Read htpasswd: %+v", string(data))

	return string(data), nil
}

func (auth BasicAuthAuthenticator) monitor(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Warnf("faild to create watcher for HTPasswd file: %s", err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Infof("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Infof("modified file:", event.Name)
					if data, err := getFileContents(path); err != nil {
						log.Errorf("could not read modified file contents: %s", err.Error())
					} else {
						auth.HTPasswdContent = data
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Errorf("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Errorf(err.Error())
	}
	<-done
}
