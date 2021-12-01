package main

import (
	"advent-of-code/helper"
	"fmt"
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
			helper.TestingValue{Result: result01, Expect: 1713},
			helper.TestingValue{Result: result02, Expect: 1734},
		},
	)
}

// Task code
func part01(input []int) int {
	count := 0

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}

	return count
}

func part02(input []int) int {
	iLen := len(input)
	count := 0
	depth := input[0] + input[1] + input[2]

	for i := 1; i < iLen-2; i++ {
		cDepth := input[i] + input[i+1] + input[i+2]

		if cDepth > depth {
			count++
		}

		depth = cDepth
	}

	return count
}
