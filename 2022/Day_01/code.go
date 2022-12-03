package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"runtime"
	"sort"
	"strconv"
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
			helper.TestingValue{Result: result01, Expect: 67027},
			helper.TestingValue{Result: result02, Expect: 197291},
		},
	)
}

type CalorieList []int

func NewCalorieList(input []string) CalorieList {
	list := CalorieList{}

	tmp := 0
	for _, s := range input {
		if "" == s {
			list = append(list, tmp)
			tmp = 0

			continue
		}

		i, _ := strconv.Atoi(s)
		tmp += i
	}

	return list
}

// Task code
func part01(input []string) int {
	l := NewCalorieList(input)
	sort.Ints(l)

	return l[len(l)-1]
}

func part02(input []string) int {
	l := NewCalorieList(input)
	sort.Ints(l)
	length := len(l)

	return l[length-1] + l[length-2] + l[length-3]
}
