package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"runtime"
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
			helper.TestingValue{Result: result01, Expect: 2081},
			helper.TestingValue{Result: result02, Expect: 2341},
		},
	)
}

// Task code
const (
	North = "^"
	East  = ">"
	South = "v"
	West  = "<"
)

func part01(input []string) int {
	commands := strings.Split(input[0], "")
	houses := map[string]bool{"0-0": true}
	coords := []int{0, 0}

	for _, c := range commands {
		y, x := coords[0], coords[1]

		switch c {
		case North:
			y--
		case East:
			x++
		case South:
			y++
		case West:
			x--
		}

		coords = []int{y, x}
		houses[fmt.Sprintf("%d-%d", y, x)] = true
	}

	return len(houses)
}

func part02(input []string) int {
	commands := strings.Split(input[0], "")
	houses := map[string]bool{"0-0": true}
	coords := [][2]int{{0, 0}, {0, 0}}

	for i, c := range commands {
		sr := 0
		if i%2 == 0 {
			sr = 1
		}
		y, x := coords[sr][0], coords[sr][1]

		switch c {
		case North:
			y--
		case East:
			x++
		case South:
			y++
		case West:
			x--
		}

		coords[sr] = [2]int{y, x}
		houses[fmt.Sprintf("%d-%d", y, x)] = true
	}

	return len(houses)
}
