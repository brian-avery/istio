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
// limitations under the License.
package examples

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"testing"

	testEnv "istio.io/istio/pkg/test/env"
	"istio.io/istio/pkg/test/framework/components/environment/kube"
	"istio.io/istio/pkg/test/scopes"
)

type scriptTestType struct {
	script    string
	output    outputType
	validator ValidationFunction
	namespace string
}

func newStepScript(namespace string, script string, output outputType, validator ValidationFunction) testStep {
	return scriptTestType{
		script:    script,
		output:    output,
		validator: validator,
		namespace: namespace,
	}
}

func (test scriptTestType) Run(env *kube.Environment, t *testing.T) (string, error) {
	scopes.CI.Infof(fmt.Sprintf("Executing %s\n", test.script))

	script, err := ioutil.ReadFile(test.script)
	if err != nil {
		return "", fmt.Errorf("test framework failed to read script: %s", err)
	}

	//TODO: Does #!/bin/sh need to be removed?

	//replace @.*@ with the correct paths
	atMatch := regexp.MustCompile("@.*@")
	script = atMatch.ReplaceAllFunc(script, func(input []byte) []byte {
		trimmed := input[1 : len(input)-1]
		return []byte(path.Join(testEnv.IstioSrc, string(trimmed)))
	})

	cmd := exec.Command("bash")
	cmd.Stdin = strings.NewReader(string(script))
	cmd.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", env.Settings().KubeConfig))
	output, err := cmd.CombinedOutput()

	//if a validation function is provided, execute that and use errors from that instead.
	//if a validation function is not provided, return errors from execution instead.
	if test.validator != nil {
		return string(output), test.validator(string(output))
	}

	test.output.Write(test.script, output)
	return string(output), err
}
func (test scriptTestType) Copy(path string) error {
	return copyFile(test.script, path)
}
