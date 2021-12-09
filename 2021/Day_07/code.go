package main

import (
	"advent-of-code/helper"
	"fmt"
	"math"
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
			helper.TestingValue{Result: result01, Expect: 352707},
			helper.TestingValue{Result: result02, Expect: 95519693},
		},
	)
}

// Task code
func part01(input []string) int {
	coords := parseCoords(input[0])
	max := maxCoord(coords)
	fuelList := make(map[int]int, 0)

	for i := 1; i <= max; i++ {
		var fuel int

		for _, c := range coords {
			if i == c {
				continue
			}

			fuel += int(math.Abs(float64(i - c)))
		}

		fuelList[i] = fuel
	}

	return minFuel(fuelList)
}

func part02(input []string) int {
	coords := parseCoords(input[0])
	max := maxCoord(coords)
	fuelList := make(map[int]int, 0)

	for i := 1; i <= max; i++ {
		var fuel int

		for _, c := range coords {
			if i == c {
				continue
			}

			base := int(math.Abs(float64(i - c)))
			fuel += (base * (base + 1)) / 2
		}

		fuelList[i] = fuel
	}

	return minFuel(fuelList)
}

func minFuel(list map[int]int) int {
	var min int
	for _, min = range list {
		break
	}

	for _, f := range list {
		if f < min {
			min = f
		}
	}

	return min
}

func maxCoord(coords []int) int {
	var max int

	for _, c := range coords {
		if c > max {
			max = c
		}
	}

	return max
}

func parseCoords(line string) []int {
	var coords []int

	for _, s := range strings.Split(line, ",") {
		c, _ := strconv.Atoi(s)
		coords = append(coords, c)
	}

	return coords
}
