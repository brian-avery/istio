package examples

import (
	"fmt"
	"io/ioutil"
	"path"
)

type outputTypeImpl interface {
	Write(name string, output []byte)
}

type textOutput struct{}

func (textOutput) Write(name string, output []byte) {
	_, filename := path.Split(name)
	fmt.Printf("Writing %s\n", filename+"_output.txt")
	ioutil.WriteFile(path.Join("output/", filename+"_output.txt"), output, 0644)

}

type yamlOutput struct{}

func (yamlOutput) Write(name string, output []byte) {
	_, filename := path.Split(name)
	fmt.Printf("Writing %s\n", filename+"_output.yaml")
	ioutil.WriteFile(path.Join("output/", filename+"_output.yaml"), output, 0644)
}

type jsonOutput struct{}

func (jsonOutput) Write(name string, output []byte) {
	_, filename := path.Split(name)
	fmt.Printf("Writing %s\n", filename+"_output.json")
	ioutil.WriteFile(path.Join("output/", filename+"_output.json"), output, 0644)
}
