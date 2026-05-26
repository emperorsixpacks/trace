package utils

import "github.com/emperorsixpacks/trace/model"

var reader = map[model.FileFormat]func(string) ([]byte, error){
	model.YAML: loadYamlFile,
	model.JSON: loadJsonFile,
}

func LoadFile(path string, filetype model.FileFormat) error {
	return reader[filetype](path)
}
