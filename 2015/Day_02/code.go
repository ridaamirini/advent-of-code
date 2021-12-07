package main

import (
	"advent-of-code/helper"
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
			helper.TestingValue{Result: result01, Expect: 1586300},
			helper.TestingValue{Result: result02, Expect: 3737498},
		},
	)
}

// Task code
func part01(input []string) int {
	var sum int

	for _, s := range input {
		numbers := convertToInt(strings.Split(s, "x"))
		l := numbers[0]
		w := numbers[1]
		h := numbers[2]
		a := l * w
		b := w * h
		c := h * l
		slack := min([]int{a, b, c})

		sum += (2*a + 2*b + 2*c) + slack
	}

	return sum
}

func part02(input []string) int {
	var sum int

	for _, s := range input {
		numbers := convertToInt(strings.Split(s, "x"))
		l := numbers[0]
		w := numbers[1]
		h := numbers[2]
		sides := []int{l, w, h}
		sort.Ints(sides)

		sum += (2*sides[0] + 2*sides[1]) + l*w*h
	}

	return sum
}

func min(n []int) int {
	m := n[0]

	for _, number := range n {
		if m > number {
			m = number
		}
	}

	return m
}

func convertToInt(s []string) []int {
	var r []int

	for _, v := range s {
		n, _ := strconv.Atoi(v)
		r = append(r, n)
	}

	return r
}
