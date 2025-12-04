package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"strconv"
)

func maxKDigits(digits []int, k int) []int {
	n := len(digits)
	drop := n - k
	var stack []int

	for _, digit := range digits {
		for drop > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			drop--
		}
		stack = append(stack, digit)
	}
	return stack[:k]
}

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
		digit_list := maxKDigits(bank_digits, 2)
		battery_charge := 0
		for _, digit := range digit_list {
			battery_charge = battery_charge*10 + digit
		}
		result += battery_charge
	}

	return result
}

func part2(input []string) int {
	result := 0

	for _, bank := range input {
		bank_digits := convertToDigitList(bank)
		digit_list := maxKDigits(bank_digits, 12)
		battery_charge := 0
		for _, digit := range digit_list {
			battery_charge = battery_charge*10 + digit
		}
		result += battery_charge
	}

	return result
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
