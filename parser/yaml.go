package parser

import (
	"github.com/emperorsixpacks/trace/model"
	"gopkg.in/yaml.v3"
)

func ParseYaml(data []byte) (*model.Flow, error) {
	var flow model.Flow
	err := yaml.Unmarshal(data, &flow)
	if err != nil {
		return nil, err
	}
	return &flow, nil
}
