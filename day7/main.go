package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"slices"
	"time"
)

func countBeams(manifold [][]rune) (int, []int) {
	count := 0
	beams := make([]int, len(manifold[0]))
	beams[slices.IndexFunc(manifold[0], func(r rune) bool { return r == 'S' })] = 1

	for _, line := range manifold[1:] {
		for i, _ := range beams {
			if line[i] == '^' && beams[i] > 0 {
				count += 1
				beams[i-1] += beams[i]
				beams[i+1] += beams[i]
				beams[i] = 0
			}
		}
	}
	return count, beams
}

func part1(input [][]rune) int {
	count, _ := countBeams(input)
	return count
}

func part2(input [][]rune) int {
	_, beams := countBeams(input)
	total := 0
	for _, b := range beams {
		total += b
	}
	return total
}

func main() {

	sample, err := utils.ReadToGraph("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	fmt.Println("Test Part 1:", part1(sample))
	fmt.Printf("Part 1 took %v\n", time.Since(start))

	start = time.Now()
	fmt.Println("Test Part 2:", part2(sample))
	fmt.Printf("Part 2 took %v\n", time.Since(start))

	input, err := utils.ReadToGraph("input.txt")
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
