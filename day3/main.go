package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"strconv"
)

func convertToDigitList(n string) []int {
	var digits []int
	for _, r := range n {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			log.Fatal(err)
		}
		digits = append(digits, digit)
	}
	return digits
}

func part1(input []string) int {
	result := 0

	for _, bank := range input {
		bank_digits := convertToDigitList(bank)
		max_digit := -1
		max_digit_index := -1
		second_max_digit := -1

		for i, digit := range bank_digits {
			if digit > max_digit {
				max_digit = digit
				max_digit_index = i
			}
			if max_digit == 9 {
				break
			}
		}

		if max_digit_index != len(bank_digits)-1 {
			for _, digit := range bank_digits[max_digit_index+1:] {
				if digit > second_max_digit {
					second_max_digit = digit
				}
			}
		} else {
			for _, digit := range bank_digits[:max_digit_index] {
				if digit > second_max_digit {
					second_max_digit = digit
				}
			}
		}

		if max_digit_index == len(bank_digits)-1 {
			result += second_max_digit*10 + max_digit
		} else {
			result += max_digit*10 + second_max_digit
		}
	}

	return result
}

func part2(input []string) int {
	return 0
}

func main() {

	sample, err := utils.ReadLines("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Test Part 1:", part1(sample))
	fmt.Println("Test Part 2:", part2(sample))

	input, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
