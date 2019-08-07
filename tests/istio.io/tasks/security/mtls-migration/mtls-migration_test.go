// Copyright 2019 Istio Authors
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
package tests

import (
	"testing"

	"istio.io/istio/pkg/test/istio.io/examples"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/environment"

	"istio.io/istio/pkg/test/framework/components/istio"
)

var (
	ist istio.Instance
)

func TestMain(m *testing.M) {
	framework.NewSuite("mtls-migration", m).
		SetupOnEnv(environment.Kube, istio.Setup(&ist, nil)).
		RequireEnvironment(environment.Kube).
		Run()
}

func validateInitialPolicies(output string) error {
	//verify that only the following exist:
	// NAMESPACE      NAME                          AGE
	// istio-system   grafana-ports-mtls-disabled   3m

	return nil
}

func validateInitialDestinationRules(output string) error {
	//verify that only the following exists:
	//NAMESPACE      NAME              AGE
	//istio-system   istio-policy      25m
	//istio-system   istio-telemetry   25m

	return nil
}

func curlVerify200FromAllRequests(output string) error {
	//verify 200ok from all requests

	return nil
}

func Verify200FromFirstTwoRequestsAnd503FromThird(output string) error {
	//verify 200 from first 2 requests and 503 from 3rd request

	return nil
}

//https://istio.io/docs/tasks/security/mtls-migration/
//https://github.com/istio/istio.io/blob/release-1.2/content/docs/tasks/security/mtls-migration/index.md
func TestMTLS(t *testing.T) {
	examples.New(t, "mtls-migration").
		AddScript("", "curl-foo-bar-legacy.sh", examples.TextOutput, nil).
		AddScript("", "verify-initial-policies.sh", examples.TextOutput, validateInitialPolicies).
		AddScript("", "verify-initial-destinationrules.sh", examples.TextOutput, validateInitialDestinationRules).
		AddScript("", "configure-mtls-destinationrule.sh", examples.TextOutput, nil).
		AddScript("", "curl-foo-bar-legacy.sh", examples.TextOutput, nil).
		AddScript("", "httpbin-foo-mtls-only.sh", examples.TextOutput, nil).
		AddScript("", "curl-foo-bar-legacy.sh", examples.TextOutput, nil).
		AddScript("", "cleanup.sh", examples.TextOutput, nil).
		Run()
}
