package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func convertToIntervalsAndPoints(input []string) ([][]int, []int) {
	intervals := [][]int{}
	points := []int{}
	for _, line := range input {
		if strings.Contains(line, "-") {
			var start, end int
			startEnd := strings.Split(line, "-")
			start, err := strconv.Atoi(startEnd[0])
			if err != nil {
				log.Fatal(err)
			}
			end, err = strconv.Atoi(startEnd[1])
			if err != nil {
				log.Fatal(err)
			}
			intervals = append(intervals, []int{start, end})
		} else if strings.TrimSpace(line) != "" {
			point, err := strconv.Atoi(strings.TrimSpace(line))
			if err != nil {
				log.Fatal(err)
			}
			points = append(points, point)
		} else {
			continue
		}
	}
	return intervals, points
}

func mergeAndSortIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastMerged := merged[len(merged)-1]
		current := intervals[i]
		if current[0] <= lastMerged[1] {
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

func isPointInIntervals(point int, intervals [][]int) bool {
	if len(intervals) == 0 {
		return false
	}
	index := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] > point
	}) - 1
	if index < 0 {
		return false
	}
	interval := intervals[index]
	return point >= interval[0] && point <= interval[1]
}

func part1(input []string) int {
	intervals, points := convertToIntervalsAndPoints(input)
	mergedIntervals := mergeAndSortIntervals(intervals)
	count := 0
	for _, point := range points {
		if isPointInIntervals(point, mergedIntervals) {
			count++
		}
	}
	return count
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
