package invdutil

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type YamlConfigFile struct {
	Apikey string
}

func ReadAPIKeyFromYaml(path string) (string, error) {
	filename, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var yamlConfig YamlConfigFile

	err = yaml.Unmarshal(yamlFile, &yamlConfig)

	if err != nil {
		return "", err
	}

	return yamlConfig.Apikey, nil
}
