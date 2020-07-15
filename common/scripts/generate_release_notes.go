package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"

	"github.com/ghodss/yaml"
	//"gopkg.in/russross/blackfriday.v2" // use blackfriday to generate an HTML file https://stackoverflow.com/questions/23124008/how-can-i-render-markdown-to-a-golang-templatehtml-or-tmpl-with-blackfriday
)

var(
	RN_TEMPLATE_DIR="./releasenotes/templates"
	RN_DIR="./releasenotes/notes"
)

type ReleaseNote struct {
	Kind string `json:"kind"`
	Area string `json:"area"`
	Issue string `json:"issue"`
	ReleaseNotes string `json:"releaseNotes"`
	UpgradeNotes string `json:"upgradeNotes"`
	SecurityNotes string `json:"securityNotes"`
}

func main() {
	notes, err := getReleaseNotesFiles(RN_DIR)
	if err != nil {
		fmt.Printf("Failed to get release notes:%s", err.Error())
		return
	}

	for _, note := range notes {
		fmt.Printf("Releasenotes: %s \n upgradeNotes: %s \n issue: %s\n Kind: %s\n\n ", note.ReleaseNotes, note.UpgradeNotes, note.Issue, note.Kind)
	}

	getTemplates(RN_TEMPLATE_DIR)


	return
}

func getReleaseNotesFiles(dir string) ([]ReleaseNote, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to list release notes: %s", err.Error())	
	}


	releaseNotes := []ReleaseNote{}
	for _, file := range files {
		//todo: make sure file is a yaml file
		yamlContents, err := ioutil.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read YAML file %s:%s", file.Name(), err.Error())
		}

		var note ReleaseNote
		if err = yaml.Unmarshal(yamlContents, &note); err != nil {
			return nil, fmt.Errorf("failed to parse YAML: %s", err.Error())
		}

		releaseNotes = append(releaseNotes, note)


	}
	return releaseNotes, nil
}


func getTemplates(dir string) (*int, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to find templates: %s", err.Error())
	}

	for _, file := range files {
		htmlCommentRegex := regexp.MustCompile("<!--(\w\s)*-->")

		markdownContents, err := ioutil.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("could not read template:%s", err.Error())
		}
		comments := htmlCommentRegex.FindAllStringSubmatch(fmt.Sprintf("%q",markdownContents), -1)
		
		fmt.Printf("Comments\n")
		for _, comment := range comments {
			fmt.Printf("\t\t%+v\n", comment)
			for _, comment2 := range comment {
				fmt.Printf("\t%+v\n", comment2)
			}
		}

		fmt.Printf("Comments done\n\n\n\n")


	}
	//populate templates
	//* find templates
	//* find comments
	//* if they match a specific pattern, populate values
	return nil, nil
}
