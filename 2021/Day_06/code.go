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
			helper.TestingValue{Result: result01, Expect: 385391},
			helper.TestingValue{Result: result02, Expect: 1728611055389},
		},
	)
}

// Task code
type Fish struct {
	days int
}

func (f *Fish) reproduce() *Fish {
	f.days--

	if f.days < 0 {
		fmt.Println(f.days)
		f.days = 6

		return &Fish{8}
	}

	return nil
}

func part01(input []string) int {
	return calcFishes(80, parsFishSchool(input[0]))
}

func part02(input []string) int {
	return calcFishes(256, parsFishSchool(input[0]))
}

func parsFishSchool(line string) map[int]int {
	fishes := make(map[int]int, 0)

	for _, s := range strings.Split(line, ",") {
		fish, _ := strconv.Atoi(s)
		fishes[fish]++
	}

	return fishes
}

func calcFishes(days int, startSchool map[int]int) int {
	fishSchool := startSchool
	for d := 1; d <= days; d++ {
		iteration := make(map[int]int, 0)

		for fd, f := range fishSchool {
			if fd == 0 {
				iteration[6] += f
				iteration[8] = f
			} else {
				iteration[fd-1] += f
			}
		}

		fishSchool = iteration
	}

	var count int
	for _, f := range fishSchool {
		count += f
	}

	return count
}
