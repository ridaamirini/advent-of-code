package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
	"unicode"
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
			helper.TestingValue{Result: result01, Expect: 4659},
			helper.TestingValue{Result: result02, Expect: 148962},
		},
	)
}

// Task code
const (
	Start = "start"
	End   = "end"
)

type Cave struct {
	value string
}

func (c *Cave) isStart() bool {
	return c.value == Start
}

func (c *Cave) isEnd() bool {
	return c.value == End
}

func (c *Cave) isSmallCave() bool {
	v := []rune(c.value)

	for _, r := range v {
		if unicode.IsLower(r) {
			return true
		}
	}

	return false
}

func part01(input []string) int {
	caveMap := createCaveMap(input)
	s := Cave{Start}
	visited := make(map[Cave]int, 0)
	paths := 0

	traverse(s, caveMap, 1, &paths, &visited)

	return paths
}

func part02(input []string) int {
	caveMap := createCaveMap(input)
	s := Cave{Start}
	visited := make(map[Cave]int, 0)
	paths := 0

	traverse(s, caveMap, 2, &paths, &visited)

	return paths
}

func createCaveMap(input []string) map[Cave][]Cave {
	res := make(map[Cave][]Cave, 0)

	for _, s := range input {
		o := strings.Split(s, "-")
		caves := []Cave{
			{o[0]},
			{o[1]},
		}
		caveOne, caveTwo := caves[0], caves[1]

		res[caveOne] = append(res[caveOne], caveTwo)
		res[caveTwo] = append(res[caveTwo], caveOne)
	}

	return res
}

func traverse(s Cave, cm map[Cave][]Cave, smallCaveMaxVisit int, paths *int, visited *map[Cave]int) {
	if s.isEnd() {
		*paths++
		return
	}

	if s.isSmallCave() {
		(*visited)[s]++

		visitedSmallCaves := 0
		for cave, _ := range *visited {
			if (*visited)[cave] > 1 {
				visitedSmallCaves++
			}

			if (*visited)[cave] > smallCaveMaxVisit {
				(*visited)[s]--

				return
			}
		}

		if visitedSmallCaves > 1 {
			(*visited)[s]--
			return
		}
	}

	for _, d := range cm[s] {
		if d.isStart() {
			continue
		}

		traverse(d, cm, smallCaveMaxVisit, paths, visited)
	}

	if s.isSmallCave() {
		(*visited)[s]--
	}
}
