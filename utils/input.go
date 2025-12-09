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

func ReadInputWithSpaces(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadLines(filename string) ([]string, error) {
	content, err := ReadInput(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(content, "\n"), nil
}

func ReadLinesWithSpaces(filename string) ([]string, error) {
	content, err := ReadInputWithSpaces(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(content, "\n"), nil
}

func ReadToGraph(filename string) ([][]rune, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	graph := make([][]rune, len(lines))
	for i, line := range lines {
		graph[i] = []rune(line)
	}

	return graph, nil
}
