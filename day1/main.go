package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"strconv"
)

func part1(input []string) int {
	// Your solution here
	count := 0
	current_dial := 50

	for _, turn := range input {
		direction := turn[0]
		turn_number, err := strconv.Atoi(turn[1:])
		if err != nil {
			fmt.Println("Error converting string to int:")
		}
		if direction == 'L' {
			current_dial = (current_dial - turn_number) % 100
		} else {
			current_dial = (current_dial + turn_number) % 100
		}
		if current_dial == 0 {
			count += 1
		}
	}
	return count
}

func part2(input string) int {
	// Your solution here
	return 0
}

func main() {

	test_day_1_input, err := utils.ReadLines("part1_sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Test Part 1:", part1(test_day_1_input))
	// fmt.Println("Test Part 2:", part2(test_input))

	day_1_input, err := utils.ReadLines("part1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(day_1_input))
	// fmt.Println("Part 2:", part2(input))
}
