package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func readFile(path string) ([]byte, error) {
	out, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("File %s not found", path)
		}
		return nil, err
	}
	return out, nil
}

func loadYamlFile(path string) ([]byte, error) {
	out, err := readFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(out, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func loadJsonFile(path string) ([]byte, error) {
	out, err := readFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(out, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
