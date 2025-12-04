package utils

import (
	"os"
	"strings"
)

func ReadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func ReadLines(filename string) ([]string, error) {
	content, err := ReadInput(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(content, "\n"), nil
}
