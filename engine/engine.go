package engine

import (
	"fmt"

	"github.com/emperorsixpacks/trace/internal/utils"
	"github.com/emperorsixpacks/trace/parser"
)

func Runner(file string) {
	data, err := utils.LoadFile(file)
	if err != nil {
		panic(err)
	}
	flow, err := parser.ParseYaml(data)
	if err != nil {
		panic(err)
	}
	for _, step := range flow.Steps {
		fmt.Println(step.Name)
	}
}
