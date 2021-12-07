package main

import (
	"advent-of-code/helper"
	"crypto/md5"
	"encoding/hex"
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
			helper.TestingValue{Result: result01, Expect: 254575},
			helper.TestingValue{Result: result02, Expect: 1038736},
		},
	)
}

// Task code
func part01(input []string) int {
	return findLowestPositiveNumber(input[0], "00000")
}

func part02(input []string) int {
	return findLowestPositiveNumber(input[0], "000000")
}

func findLowestPositiveNumber(key string, prefix string) int {
	counter := 0

	for {
		b := []byte(fmt.Sprintf("%s%d", key, counter))
		hash := md5.Sum(b)

		if strings.HasPrefix(hex.EncodeToString(hash[:]), prefix) {
			return counter
		}

		counter++
	}
}
