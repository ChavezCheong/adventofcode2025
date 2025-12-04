package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func check_invalid_id(id int) bool {
	id_str := strconv.Itoa(id)
	id_len := len(id_str)
	if id_len%2 != 0 {
		return false
	}
	first_half := id_str[:id_len/2]
	second_half := id_str[id_len/2:]
	return first_half == second_half
}

func part1(input string) int {
	result := 0

	products := strings.Split(input, ",")
	for _, product := range products {
		id_range := strings.Split(product, "-")
		id_start, err := strconv.Atoi(id_range[0])
		if err != nil {
			log.Fatal(err)
		}
		id_end, err := strconv.Atoi(id_range[1])
		if err != nil {
			log.Fatal(err)
		}
		for id := id_start; id <= id_end; id++ {
			if check_invalid_id(id) {
				result += id
			}
		}

	}
	return result
}

func part2(input string) int {
	return 0
}

func main() {

	part1_sample, err := utils.ReadInput("sample1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Test Part 1:", part1(part1_sample))
	fmt.Println("Test Part 2:", part2(part1_sample))

	part1_input, err := utils.ReadInput("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(part1_input))
	fmt.Println("Part 2:", part2(part1_input))
}
