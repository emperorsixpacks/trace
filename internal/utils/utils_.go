package utils

import (
	"fmt"
	"os"
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
