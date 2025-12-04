package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func check_invalid_id_part1(id int) bool {
	id_str := strconv.Itoa(id)
	id_len := len(id_str)
	if id_len%2 != 0 {
		return false
	}
	first_half := id_str[:id_len/2]
	second_half := id_str[id_len/2:]
	return first_half == second_half
}

func check_invalid_id_part2(id int) bool {
	id_str := strconv.Itoa(id)
	id_len := len(id_str)

	if id_len == 1 {
		return false
	}

	for pattern_length := 1; pattern_length <= id_len/2; pattern_length++ {
		if id_len%pattern_length != 0 {
			continue
		} else {
			is_invalid := true
			for i := 0; i < id_len; i += 1 {
				if id_str[i] != id_str[i%pattern_length] {
					is_invalid = false
				}
			}
			if is_invalid {
				return true
			}
		}
	}
	return false
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
			if check_invalid_id_part1(id) {
				result += id
			}
		}

	}
	return result
}

func part2(input string) int {
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
			if check_invalid_id_part2(id) {
				result += id
			}
		}

	}
	return result
}

func main() {

	sample, err := utils.ReadInput("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Test Part 1:", part1(sample))
	fmt.Println("Test Part 2:", part2(sample))

	input, err := utils.ReadInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
