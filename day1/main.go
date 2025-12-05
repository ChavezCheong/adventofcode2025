package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"strconv"
	"time"
)

func circular_add(current int, addition int, modulo int) int {
	result := (current + addition) % modulo
	if result < 0 {
		result += modulo
	}
	return result
}

func circular_subtract(current int, subtraction int, modulo int) int {
	result := (current - subtraction) % modulo
	if result < 0 {
		result += modulo
	}
	return result
}

func part1(input []string) int {
	count := 0
	current_dial := 50

	for _, turn := range input {
		direction := turn[0]
		turn_number, err := strconv.Atoi(turn[1:])
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		if direction == 'L' {
			current_dial = circular_subtract(current_dial, turn_number, 100)
		} else {
			current_dial = circular_add(current_dial, turn_number, 100)
		}
		if current_dial == 0 {
			count += 1
		}
	}
	return count
}

func part2(input []string) int {
	count := 0
	current_dial := 50

	for _, turn := range input {
		direction := turn[0]
		turn_number, err := strconv.Atoi(turn[1:])
		if err != nil {
			fmt.Println("Error converting string to int", err)
		}
		if direction == 'L' {
			dial_raw := current_dial - turn_number
			if dial_raw <= 0 {
				if current_dial == 0 {
					count += (-dial_raw / 100)
				} else {
					count += (-dial_raw / 100) + 1
				}
			}
			current_dial = circular_subtract(current_dial, turn_number, 100)
		} else {
			dial_raw := current_dial + turn_number
			if dial_raw >= 100 {
				count += (dial_raw / 100)
			}
			current_dial = circular_add(current_dial, turn_number, 100)
		}
	}
	return count
}

func main() {

	sample, err := utils.ReadLines("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	fmt.Println("Test Part 1:", part1(sample))
	fmt.Printf("Part 1 took %v\n", time.Since(start))

	start = time.Now()
	fmt.Println("Test Part 2:", part2(sample))
	fmt.Printf("Part 2 took %v\n", time.Since(start))

	input, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	fmt.Println("Part 1:", part1(input))
	fmt.Printf("Part 1 took %v\n", time.Since(start))

	start = time.Now()
	fmt.Println("Part 2:", part2(input))
	fmt.Printf("Part 2 took %v\n", time.Since(start))
}
