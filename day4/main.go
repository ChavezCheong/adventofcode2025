package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
)

func surroundingAccessibleCheck(row int, col int, input [][]rune) bool {
	count := 0
	directions := [][]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}
	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]
		if newRow >= 0 && newRow < len(input) && newCol >= 0 && newCol < len(input[0]) {
			if input[newRow][newCol] == '@' {
				count++
			}
		}
	}
	return count < 4
}

func countingAccessible(input [][]rune) int {
	count := 0

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == '@' {
				if surroundingAccessibleCheck(row, col, input) {
					count++
				}
			}

		}
	}
	return count
}

func graphAfterIteration(input [][]rune) [][]rune {
	newGraph := make([][]rune, len(input))
	for row := 0; row < len(input); row++ {
		newGraph[row] = make([]rune, len(input[row]))
		for col := 0; col < len(input[row]); col++ {
			if input[row][col] == '@' {
				if surroundingAccessibleCheck(row, col, input) {
					newGraph[row][col] = '.'
				} else {
					newGraph[row][col] = '@'
				}
			} else {
				newGraph[row][col] = input[row][col]
			}
		}
	}
	return newGraph
}

func part1(input [][]rune) int {
	result := countingAccessible(input)
	return result
}

func part2(input [][]rune) int {
	result := 0
	for countingAccessible(input) > 0 {
		result += countingAccessible(input)
		input = graphAfterIteration(input)
	}
	return result
}

func main() {

	sample, err := utils.ReadToGraph("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Test Part 1:", part1(sample))
	fmt.Println("Test Part 2:", part2(sample))

	input, err := utils.ReadToGraph("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
