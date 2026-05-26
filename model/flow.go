package model

type Flow struct {
	Name  string `yaml:"name"`
	Steps []Step `yaml:"steps"`
}

type Step struct {
	Name   string
	Action string
	Args   []string
}
