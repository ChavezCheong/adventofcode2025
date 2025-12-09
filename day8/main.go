package main

import (
	"advent-of-code-2025/utils"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
	z int
}

type PointandDistance struct {
	pointA   int
	pointB   int
	distance int
}

type Graph struct {
	adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
	g.adj[v] = append(g.adj[v], u)
}

func (g *Graph) DFS(v int, visited map[int]bool) {
	visited[v] = true
	for _, neighbor := range g.adj[v] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

func findDistance(a, b Point) int {
	return int(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return
	}

	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

func (uf *UnionFind) CountComponents() int {
	rootSet := make(map[int]bool)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		rootSet[root] = true
	}
	return len(rootSet)
}

func connectKPoints(distances []PointandDistance, numPoints, k int) [][]int {
	uf := NewUnionFind(numPoints)

	for itr := 0; itr < k; itr++ {
		uf.Union(distances[itr].pointA, distances[itr].pointB)
	}

	// Group points by their root parent
	componentMap := make(map[int][]int)
	for i := 0; i < numPoints; i++ {
		root := uf.Find(i)
		componentMap[root] = append(componentMap[root], i)
	}

	// Convert map to slice of components
	components := [][]int{}
	for _, component := range componentMap {
		components = append(components, component)
	}

	return components
}

func part1(input []string, k int) int {
	numPoints := len(input)
	distances := make([]PointandDistance, 0)
	for i := 0; i < numPoints; i++ {
		for j := i + 1; j < numPoints; j++ {
			var pointA, pointB Point
			fmt.Sscanf(input[i], "%d,%d,%d", &pointA.x, &pointA.y, &pointA.z)
			fmt.Sscanf(input[j], "%d,%d,%d", &pointB.x, &pointB.y, &pointB.z)
			distance := findDistance(pointA, pointB)
			distances = append(distances, PointandDistance{pointA: i, pointB: j, distance: distance})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
	components := connectKPoints(distances, numPoints, k)

	componentSizes := make([]int, len(components))
	for i, component := range components {
		componentSizes[i] = len(component)
	}
	sort.Ints(componentSizes)
	return componentSizes[len(componentSizes)-1] * componentSizes[len(componentSizes)-2] * componentSizes[len(componentSizes)-3]
}

func part2(input []string) int {
	numPoints := len(input)
	distances := make([]PointandDistance, 0)
	for i := 0; i < numPoints; i++ {
		for j := i + 1; j < numPoints; j++ {
			var pointA, pointB Point
			fmt.Sscanf(input[i], "%d,%d,%d", &pointA.x, &pointA.y, &pointA.z)
			fmt.Sscanf(input[j], "%d,%d,%d", &pointB.x, &pointB.y, &pointB.z)
			distance := findDistance(pointA, pointB)
			distances = append(distances, PointandDistance{pointA: i, pointB: j, distance: distance})
		}
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	uf := NewUnionFind(numPoints)
	finalPointA := -1
	finalPointB := -1

	for _, edge := range distances {
		rootA := uf.Find(edge.pointA)
		rootB := uf.Find(edge.pointB)

		if rootA != rootB {
			uf.Union(edge.pointA, edge.pointB)

			// Early exit when only 1 component remains
			if uf.CountComponents() == 1 {
				finalPointA = edge.pointA
				finalPointB = edge.pointB
				break
			}
		}
	}
	finalPointAX, err := strconv.Atoi(strings.Split(input[finalPointA], ",")[0])
	if err != nil {
		log.Fatal(err)
	}
	finalPointBX, err := strconv.Atoi(strings.Split(input[finalPointB], ",")[0])
	if err != nil {
		log.Fatal(err)
	}

	return finalPointAX * finalPointBX
}

func main() {

	sample, err := utils.ReadLines("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	fmt.Println("Test Part 1:", part1(sample, 10))
	fmt.Printf("Part 1 took %v\n", time.Since(start))

	start = time.Now()
	fmt.Println("Test Part 2:", part2(sample))
	fmt.Printf("Part 2 took %v\n", time.Since(start))

	input, err := utils.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	fmt.Println("Part 1:", part1(input, 1000))
	fmt.Printf("Part 1 took %v\n", time.Since(start))

	start = time.Now()
	fmt.Println("Part 2:", part2(input))
	fmt.Printf("Part 2 took %v\n", time.Since(start))
}
