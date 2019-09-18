// Copyright 2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v2

import (
	"fmt"
	"time"

	"istio.io/istio/pkg/util/protomarshal"

	xdsapi "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/gogo/protobuf/types"

	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/networking/util"
)

func (s *DiscoveryServer) pushRoute(con *XdsConnection, push *model.PushContext, version string) error {
	pushStart := time.Now()
	rawRoutes := s.generateRawRoutes(con, push)
	if s.DebugConfigs {
		for _, r := range rawRoutes {
			con.RouteConfigs[r.Name] = r
			if adsLog.DebugEnabled() {
				resp, _ := protomarshal.ToJSONWithIndent(r, " ")
				adsLog.Debugf("RDS: Adding route:%s for node:%v", resp, con.node.ID)
			}
		}
	}

	response := routeDiscoveryResponse(rawRoutes, version, push.Version)
	err := con.send(response)
	rdsPushTime.Record(time.Since(pushStart).Seconds())
	if err != nil {
		adsLog.Warnf("RDS: Send failure for node:%v: %v", con.node.ID, err)
		recordSendError(rdsSendErrPushes, err)
		return err
	}
	rdsPushes.Increment()

	adsLog.Infof("RDS: PUSH for node:%s routes:%d", con.node.ID, len(rawRoutes))
	return nil
}

func (s *DiscoveryServer) pushDeltaRoute(con *XdsConnection, push *model.PushContext, removedResources []string) error {
	rawRoutes := s.generateRawRoutes(con, push)
	if s.DebugConfigs {
		for _, r := range rawRoutes {
			con.RouteConfigs[r.Name] = r
			if adsLog.DebugEnabled() {
				resp, _ := protomarshal.ToJSONWithIndent(r, " ")
				adsLog.Debugf("RDS: Adding route %s for node %v", resp, con.modelNode)
			}
		}
	}

	response := deltaRouteDiscoveryResponse(rawRoutes, removedResources)
	if err := con.sendDelta(response); err != nil {
		adsLog.Warnf("ADS: RDS: Send failure for node %v, closing grpc %v", con.modelNode, err)
		deltaRdsSendErrPushes.Increment()
		return err
	}
	deltaRdsPushes.Increment()

	adsLog.Infof("ADS: RDS: PUSH for node: %s addr:%s routes:%d", con.modelNode.ID, con.PeerAddr, len(rawRoutes))
	return nil
}

func (s *DiscoveryServer) generateRawRoutes(con *XdsConnection, push *model.PushContext) []*xdsapi.RouteConfiguration {
	rawRoutes := s.ConfigGenerator.BuildHTTPRoutes(s.Env, con.node, push, con.Routes)
	// Now validate each route
	for _, r := range rawRoutes {
		if err := r.Validate(); err != nil {
			adsLog.Errorf("RDS: Generated invalid routes for route:%s for node:%v: %v, %v", r.Name, con.node.ID, err, r)
			rdsBuildErrPushes.Increment()
			// Generating invalid routes is a bug.
			// Instead of panic, which will break down the whole cluster. Just ignore it here, let envoy process it.
		}
	}
	return rawRoutes
}

func (s *DiscoveryServer) generateRouteConfig(con *XdsConnection, push *model.PushContext, routeName string) (*xdsapi.RouteConfiguration, error) {
	//BAVERY_TODO: Fix this
	routeConfig := s.ConfigGenerator.BuildHTTPRoutes(s.Env, con.modelNode, push, []string{routeName})
	if routeConfig == nil {
		adsLog.Warnf("RDS: got nil value for route %s for node %v", routeName, con.modelNode)
		return nil, nil
	}

	//BAVERY_TODO: Fix this
	if err := routeConfig[0].Validate(); err != nil {
		retErr := fmt.Errorf("RDS: Generated invalid route %s for node %v: %v", routeName, con.modelNode, err)
		adsLog.Errorf("RDS: Generated invalid routes for route: %s for node: %v: %v, %v", routeName, con.modelNode, err, routeConfig)
		deltaRdsBuildErrPushes.Increment()
		// Generating invalid routes is a bug.
		// Panic instead of trying to recover from that, since we can't
		// assume anything about the state.
		panic(retErr.Error())
	}

	return routeConfig[0], nil
}

func deltaRouteDiscoveryResponse(routeConfigs []*xdsapi.RouteConfiguration, removedResources []string) *xdsapi.DeltaDiscoveryResponse {
	resp := &xdsapi.DeltaDiscoveryResponse{
		Nonce:             nonce(),
		SystemVersionInfo: versionInfo(),
	}

	for _, routeConfig := range routeConfigs {
		marshaledRoute, _ := types.MarshalAny(routeConfig)
		resp.Resources = append(resp.Resources, &xdsapi.Resource{
			Name: routeConfig.Name,
			//BAVERY_TODO: Add version.... md5? Something else? xDS increments an integer, but I wonder if we can do one better (i.e. repeatable)
			Resource: marshaledRoute,
		})
	}
	resp.RemovedResources = removedResources
	return resp
}

func routeDiscoveryResponse(rs []*xdsapi.RouteConfiguration, version string, noncePrefix string) *xdsapi.DiscoveryResponse {
	resp := &xdsapi.DiscoveryResponse{
		TypeUrl:     RouteType,
		VersionInfo: version,
		Nonce:       nonce(noncePrefix),
	}
	for _, rc := range rs {
		rr := util.MessageToAny(rc)
		resp.Resources = append(resp.Resources, rr)
	}

	return resp
}
