package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"runtime"
	"sort"
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
			helper.TestingValue{Result: result01, Expect: 464991},
			helper.TestingValue{Result: result02, Expect: 3662008566},
		},
	)
}

// Task code
const (
	BracketOpened       = "("
	BracketClosed       = ")"
	SquareBracketOpened = "["
	SquareBracketClosed = "]"
	BraceOpened         = "{"
	BraceClosed         = "}"
	CrocodileOpened     = "<"
	CrocodileClosed     = ">"
)

var OpeningBrackets = map[string]interface{}{
	BracketOpened:       nil,
	SquareBracketOpened: nil,
	BraceOpened:         nil,
	CrocodileOpened:     nil,
}

var ClosingBrackets = map[string]interface{}{
	BracketClosed:       nil,
	SquareBracketClosed: nil,
	BraceClosed:         nil,
	CrocodileClosed:     nil,
}

var SyntaxErrorScorePoints = map[string]int{
	BracketClosed:       3,
	SquareBracketClosed: 57,
	BraceClosed:         1197,
	CrocodileClosed:     25137,
}

var AutocompleteScorePoints = map[string]int{
	BracketClosed:       1,
	SquareBracketClosed: 2,
	BraceClosed:         3,
	CrocodileClosed:     4,
}

func part01(input []string) int {
	var wrongBrackets []string

	for _, line := range input {
		var lifo []rune

		for _, s := range strings.Split(line, "") {
			r := rune(s[0])

			if isOpeningBracket(s) {
				lifo = append(lifo, r)

				continue
			}

			lel := lifo[len(lifo)-1]
			if isCorrectClosingBracket(lel, r) {
				lifo = lifo[:len(lifo)-1]

				continue
			}

			wrongBrackets = append(wrongBrackets, s)

			break
		}
	}

	return calcSyntaxScore(wrongBrackets)
}

func part02(input []string) int {
	var missingBracketsScore []int

outer:
	for _, line := range input {
		var score int
		var lifo []rune

		for _, s := range strings.Split(line, "") {
			r := rune(s[0])

			if isOpeningBracket(s) {
				lifo = append(lifo, r)

				continue
			}

			lel := lifo[len(lifo)-1]
			if isCorrectClosingBracket(lel, r) {
				lifo = lifo[:len(lifo)-1]

				continue
			}

			continue outer
		}

		for len(lifo) > 0 {
			c := determineClosingBracket(lifo[len(lifo)-1])
			points, _ := AutocompleteScorePoints[c]
			score = (score * 5) + points
			lifo = lifo[:len(lifo)-1]
		}

		missingBracketsScore = append(missingBracketsScore, score)
	}

	sort.Ints(missingBracketsScore)

	return missingBracketsScore[len(missingBracketsScore)/2]
}

func calcSyntaxScore(list []string) int {
	var res int

	for _, s := range list {
		points, _ := SyntaxErrorScorePoints[s]
		res += points
	}

	return res
}

func isOpeningBracket(c string) bool {
	_, ok := OpeningBrackets[c]
	return ok
}

func isCorrectClosingBracket(o rune, c rune) bool {
	return o+1 == c || o+2 == c
}

func determineClosingBracket(r rune) string {
	b := r + 1

	if _, isClosingBracket := ClosingBrackets[string(b)]; isClosingBracket {
		return string(b)
	}

	return string(r + 2)
}
