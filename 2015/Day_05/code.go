package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"regexp"
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
			helper.TestingValue{Result: result01, Expect: 255},
			helper.TestingValue{Result: result02, Expect: 55},
		},
	)
}

// Task code
func part01(input []string) int {
	count := 0
	vowel := regexp.MustCompile(`[aeiou]{1}.*[aeiou]{1}.*[aeiou]{1}.*`)
	restricted := regexp.MustCompile(`(ab|cd|pq|xy)`)
	hasRepeatedLetters := func(s string, offset int) bool {
		for i := 0; i < len(s)-1; i++ {
			if s[i] == s[i+1] {
				return true
			}
		}

		return false
	}

	for _, s := range input {
		if vowel.MatchString(s) &&
			hasRepeatedLetters(s, 1) &&
			!restricted.MatchString(s) {
			count++
		}
	}

	return count
}

func part02(input []string) int {
	count := 0
	hasRepeatedLetters := func(s string) bool {
		for i := 0; i < len(s)-2; i++ {
			if s[i] == s[i+2] {
				return true
			}
		}

		return false
	}

	for _, s := range input {
		if hasRepeatedLetters(s) &&
			hasRepeatedDoubleLetters(s) {
			count++
		}
	}

	return count
}

func hasRepeatedDoubleLetters(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		pOne := s[i : i+2]

		if strings.Count(s, pOne) > 1 {
			return true
		}

	}

	return false
}
