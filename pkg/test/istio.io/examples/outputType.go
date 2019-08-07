package examples

import (
	"io/ioutil"
)

type outputTypeImpl interface {
	Write(name string, output []byte)
}

//todo: do we need a path to write to?
//todo: what if script is a path? Should we split?
type textOutput struct{}

func (textOutput) Write(name string, output []byte) {
	ioutil.WriteFile(name+".txt", output, 0644)

}

type yamlOutput struct{}

func (yamlOutput) Write(name string, output []byte) {
	ioutil.WriteFile(name+".yaml", output, 0644)
}

type jsonOutput struct{}

func (jsonOutput) Write(name string, output []byte) {
	ioutil.WriteFile(name+".json", output, 0644)
}
