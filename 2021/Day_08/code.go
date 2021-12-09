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
			helper.TestingValue{Result: result01, Expect: 532},
			helper.TestingValue{Result: result02, Expect: 1011284},
		},
	)
}

// Task code
func part01(input []string) int {
	numberList := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	var count int

	for _, s := range input {
		_, digits := parsePattern(s)

		for _, d := range digits {
			if _, ok := numberList[len(d)]; ok {
				count++
			}
		}
	}

	return count
}

func part02(input []string) int {
	var count int

	for _, s := range input {
		signal, digits := parsePattern(s)
		n := deduce(signal)
		decoded, _ := strconv.Atoi(
			fmt.Sprintf("%d%d%d%d", n[digits[0]], n[digits[1]], n[digits[2]], n[digits[3]]),
		)

		count += decoded
	}

	return count
}

func parsePattern(input string) (signals []string, digits []string) {
	data := strings.Split(input, " | ")
	signals = strings.Split(data[0], " ")
	digits = strings.Split(data[1], " ")

	for i, s := range signals {
		signals[i] = normalize(s)
	}

	for i, d := range digits {
		digits[i] = normalize(d)
	}

	return signals, digits
}

func deduce(signal []string) map[string]int {
	var len2, len3, len4, len7 string
	var len5, len6 []string
	var six string

	for _, si := range signal {
		switch len(si) {
		case 2:
			len2 = si
		case 3:
			len3 = si
		case 4:
			len4 = si
		case 7:
			len7 = si
		case 5:
			len5 = append(len5, si)
		default:
			len6 = append(len6, si)
		}
	}

	numberList := map[string]int{
		len2: 1,
		len4: 4,
		len3: 7,
		len7: 8,
	}

	for i, d := range len6 {
		if strings.Index(d, string(len2[0])) == -1 ||
			strings.Index(d, string(len2[1])) == -1 {
			numberList[d] = 6
			six = d
			len6 = append(len6[:i], len6[i+1:]...)

			break
		}
	}

	for i, d := range len6 {
		if strings.Index(d, string(len4[0])) == -1 ||
			strings.Index(d, string(len4[1])) == -1 ||
			strings.Index(d, string(len4[2])) == -1 ||
			strings.Index(d, string(len4[3])) == -1 {
			numberList[d] = 0
			len6 = append(len6[:i], len6[i+1:]...)

			break
		}
	}

	numberList[len6[0]] = 9

	for i, d := range len5 {
		if strings.Index(d, string(len2[0])) != -1 &&
			strings.Index(d, string(len2[1])) != -1 {
			numberList[d] = 3
			len5 = append(len5[:i], len5[i+1:]...)

			break
		}
	}

	for i, d := range len5 {
		if strings.Index(six, string(d[0])) != -1 &&
			strings.Index(six, string(d[1])) != -1 &&
			strings.Index(six, string(d[2])) != -1 &&
			strings.Index(six, string(d[3])) != -1 &&
			strings.Index(six, string(d[4])) != -1 {
			numberList[d] = 5
			len5 = append(len5[:i], len5[i+1:]...)

			break
		}
	}

	numberList[len5[0]] = 2

	return numberList
}

func normalize(v string) string {
	s := strings.Split(v, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
