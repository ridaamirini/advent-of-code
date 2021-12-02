package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"runtime"
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
			helper.TestingValue{Result: result01, Expect: 1936494},
			helper.TestingValue{Result: result02, Expect: 1997106066},
		},
	)
}

// Task code
const (
	StepForward = "forward"
	StepDown    = "down"
	StepUp      = "up"
)

func part01(input []string) int {
	var h, d int

	for _, s := range input {
		v := strings.Fields(s)
		step := v[0]
		number, _ := strconv.Atoi(v[1])

		if step == StepForward {
			h += number

			continue
		}

		if step == StepUp {
			d -= number

			continue
		}

		if step == StepDown {
			d += number

			continue
		}
	}

	return h * d
}

func part02(input []string) int {
	var h, d, a int

	for _, s := range input {
		v := strings.Fields(s)
		step := v[0]
		number, _ := strconv.Atoi(v[1])

		if step == StepForward {
			h += number
			d += number * a

			continue
		}

		if step == StepUp {
			a -= number

			continue
		}

		if step == StepDown {
			a += number

			continue
		}
	}

	return h * d
}
