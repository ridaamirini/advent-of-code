package main

import (
	"advent-of-code/helper"
	"container/list"
	"fmt"
	"path"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	_, f, _, _ := runtime.Caller(0)
	cwd := path.Join(path.Dir(f))

	input := helper.ReadInput(cwd, helper.NewLine)
	iValues := input.Strings()

	// Execute
	start := time.Now()
	result01 := part01(iValues)
	result02 := part02(iValues)
	executionTime := helper.ExecutionTime(time.Since(start))

	fmt.Printf("Solution Part 1: %v\n", result01)
	fmt.Printf("Solution Part 2: %v\n", result02)
	fmt.Printf("Execution time: %s\n", executionTime)

	helper.SaveBenchmarkTime(executionTime, cwd)

	// Testing
	helper.TestResults(
		[]helper.TestingValue{
			helper.TestingValue{Result: result01, Expect: 566},
			helper.TestingValue{Result: result02, Expect: 891684},
		},
	)
}

// Task code
var (
	Up         = Point{0, -1}
	Down       = Point{0, 1}
	Left       = Point{-1, 0}
	Right      = Point{1, 0}
	Directions = [4]Point{Up, Down, Left, Right}
)

type Heightmap map[Point]int

type Point struct {
	x, y int
}

func (p *Point) add(position Point) Point {
	return Point{p.x + position.x, p.y + position.y}
}

func (p Point) neighbors() (n [4]Point) {
	for i, d := range Directions {
		n[i] = p.add(d)
	}

	return n
}

func (hm Heightmap) lowPoints() (points []Point) {
	for point, _ := range hm {
		if hm.isLowPoint(point) {
			points = append(points, point)
		}
	}

	return points
}

func (hm Heightmap) isLowPoint(p Point) bool {
	height, ok := hm[p]
	if !ok {
		return false
	}

	for _, np := range p.neighbors() {
		if n, ok := hm[np]; ok && n <= height {
			return false
		}
	}

	return true
}

func (hm Heightmap) basinSize(p Point) (size int) {
	q := list.New()
	q.PushBack(p)

	visited := map[Point]bool{p: true}

	for e := q.Front(); e != nil; e = e.Next() {
		pos := e.Value.(Point)

		height, ok := hm[pos]
		if !ok || height == 9 {
			continue
		}

		for _, nPos := range pos.neighbors() {
			if _, ok := visited[nPos]; ok {
				continue
			}
			visited[nPos] = true
			q.PushBack(nPos)
		}

		size++
	}

	return size
}

func part01(input []string) (riskLevel int) {
	heightMap := createHeightmap(input)

	for _, point := range heightMap.lowPoints() {
		riskLevel += heightMap[point] + 1
	}

	return riskLevel
}

func part02(input []string) int {
	heightMap := createHeightmap(input)
	lowPoints := heightMap.lowPoints()
	basinSizes := make([]int, len(lowPoints))

	for i, point := range lowPoints {
		basinSizes[i] = heightMap.basinSize(point)
	}

	sort.Ints(basinSizes)

	result := 1
	for _, basin := range basinSizes[len(basinSizes)-3:] {
		result *= basin
	}

	return result
}

func createHeightmap(input []string) Heightmap {
	heightMap := make(map[Point]int, 0)

	for y, l := range input {
		for x, s := range strings.Split(l, "") {
			n, _ := strconv.Atoi(s)
			heightMap[Point{x, y}] = n
		}
	}

	return heightMap
}
