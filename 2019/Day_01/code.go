package main

import (
	"advent-of-code/helper"
	"fmt"
	"math"
	"path"
	"runtime"
	"time"
)

func main() {
	_, f, _, _ := runtime.Caller(0)
	cwd := path.Join(path.Dir(f))

	input := helper.ReadInput(cwd, helper.NewLine)
	iValues := input.Ints()

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
			helper.TestingValue{Result: result01, Expect: 3226822},
			helper.TestingValue{Result: result02, Expect: 4837367},
		},
	)
}

// Task code
func part01(input []int) int {
	var r int

	for _, m := range input {
		r += int(math.Floor(float64(m)/3)) - 2
	}

	return r
}

func part02(input []int) int {
	var r int

	for _, m := range input {
		fuel := m

		for fuel > 0 {
			fuel = int(math.Floor(float64(fuel)/3)) - 2

			if fuel > 0 {
				r += fuel
			}
		}
	}

	return r
}
