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
	"os"
	"path"
	"testing"

	"istio.io/istio/pkg/test/env"
	"istio.io/istio/pkg/test/framework/components/environment/kube"
	"istio.io/istio/pkg/test/scopes"

	"istio.io/istio/pkg/test/framework"
)

type outputType outputTypeImpl

var (
	// TextOutput describes a test which returns text output
	TextOutput outputType = textOutput{}
	// YamlOutput describes a test which returns yaml output
	YamlOutput outputType = yamlOutput{}
	// JSONOutput describes a test which returns json output
	JSONOutput outputType = jsonOutput{}
)

const (
	istioPath = "istio.io/istio/"
)

//ValidationFunction is used to define a function that will be used to verify output
type ValidationFunction func(stdOut string) error

// Example manages the steps in a test, executes them, and records the output
type Example struct {
	name  string
	steps []testStep
	t     *testing.T
}

// New returns an instance of an example test
func New(t *testing.T, name string) *Example {
	return &Example{
		name:  name,
		steps: make([]testStep, 0),
		t:     t,
	}
}

// RunScript adds a directive to run a script
func (example *Example) RunScript(namespace string, script string, output outputType, validator ValidationFunction) *Example {
	example.steps = append(example.steps, newStepScript(namespace, script, output, validator))
	return example
}

// AddFile adds an existing file
func (example *Example) Apply(namespace string, path string) *Example {
	fullPath := getFullPath(istioPath + path)
	example.steps = append(example.steps, newStepFile(namespace, fullPath, false))
	return example
}

// Delete deletes an existing file
func (example *Example) Delete(namespace string, path string) {
	fullPath := getFullPath(istioPath + path)
	example.steps = append(example.steps, newStepFile(namespace, fullPath, true))
}

type testFunc func(t *testing.T) error

// Exec registers a callback to be invoked synchronously. This is typically used for
// validation logic to ensure command-lines worked as intended
func (example *Example) Exec(testFunction testFunc) *Example {
	example.steps = append(example.steps, newStepFunction(testFunction))
	return example
}

// getFullPath reads the current gopath from environment variables and appends
// the specified path to it
func getFullPath(scriptPath string) string {
	return path.Join(env.IstioTop, "/src/"+scriptPath)

}

// Run runs the scripts and capture output
// TODO: this overrides os.Stdout/os.Stderr and is not thread-safe
func (example *Example) Run() {
	scopes.CI.Infof(fmt.Sprintf("Executing test %s (%d steps)", example.name, len(example.steps)))
	//create directory if it doesn't exist
	if _, err := os.Stat(example.name); os.IsNotExist(err) {
		err := os.Mkdir(example.name, os.ModePerm)
		if err != nil {
			example.t.Fatalf("test framework failed to create directory: %s", err)
		}
	}

	framework.
		NewTest(example.t).
		Run(func(ctx framework.TestContext) {
			kubeEnv, ok := ctx.Environment().(*kube.Environment)
			if !ok {
				example.t.Fatalf("test framework unable to get Kubernetes environment")
			}
			for _, step := range example.steps {
				output, err := step.Run(kubeEnv, example.t)
				example.t.Log(output)
				if err != nil {
					example.t.Fatal(err)
				}
			}
		})
}
