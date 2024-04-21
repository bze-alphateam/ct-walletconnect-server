package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func LoadConfig(path string) (*App, error) {
	// Open the file located at the path
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file's contents
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the YAML bytes into the App struct
	var app App
	err = yaml.Unmarshal(bytes, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}
