package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func parseLines(input []string) ([][]int, []string) {
	split := input[:len(input)-1]
	lines := make([][]int, len(split))

	for i, line := range split {
		fields := strings.Fields(line)
		lines[i] = make([]int, len(fields))
		for j, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}
			lines[i][j] = num
		}
	}

	ops := strings.Fields(input[len(input)-1])
	return lines, ops
}

func transposeInput[T int | string](input [][]T) [][]T {
	if len(input) == 0 {
		return [][]T{}
	}
	transposed := make([][]T, len(input[0]))
	for i := range transposed {
		transposed[i] = make([]T, len(input))
	}
	for i := range input {
		for j := range input[i] {
			transposed[j][i] = input[i][j]
		}
	}
	return transposed
}

func calculateLine(line []int, op string) int {
	acc := line[0]
	for _, line := range line[1:] {
		switch op {
		case "+":
			acc += line
		case "*":
			acc *= line
		}
	}
	return acc
}

func part1(input []string) int {

	lines, ops := parseLines(input)
	transposed := transposeInput(lines)

	result := 0
	for i, line := range transposed {
		result += calculateLine(line, ops[i])
	}
	return result
}

func part2(input []string) int {
	lines := input[:len(input)-1]
	re := regexp.MustCompile(`[*+][ ]*`)
	ops := re.FindAllString(input[len(input)-1], -1)

	result := 0

	curr_Pos := 0
	for i, op := range ops {
		transposedLine := make([]int, len(op)-1)
		for j := 0; j < len(op)-1; j++ {
			num_acc := ""
			for _, line := range lines {
				num_acc += strings.TrimSpace(string(line[curr_Pos]))
			}
			num, err := strconv.Atoi(num_acc)
			if err != nil {
				log.Fatal(err)
			}
			transposedLine[j] = num
			curr_Pos += 1
		}
		result += calculateLine(transposedLine, strings.TrimSpace(ops[i]))
		curr_Pos += 1
	}
	return result
}

func main() {

	sample, err := utils.ReadLines("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	sample2, err := utils.ReadLinesWithSpaces("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	fmt.Println("Test Part 1:", part1(sample))
	fmt.Printf("Part 1 took %v\n", time.Since(start))

	start = time.Now()
	fmt.Println("Test Part 2:", part2(sample2))
	fmt.Printf("Part 2 took %v\n", time.Since(start))

	input, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input2, err := utils.ReadLinesWithSpaces("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	fmt.Println("Part 1:", part1(input))
	fmt.Printf("Part 1 took %v\n", time.Since(start))

	start = time.Now()
	fmt.Println("Part 2:", part2(input2))
	fmt.Printf("Part 2 took %v\n", time.Since(start))
}
