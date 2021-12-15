package main

import (
	"advent-of-code/helper"
	"fmt"
	"math"
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
			helper.TestingValue{Result: result01, Expect: 3831},
			helper.TestingValue{Result: result02, Expect: 5725739914282},
		},
	)
}

// Task code
func part01(input []string) int {
	return calculate(input, 10)
}

func part02(input []string) int {
	return calculate(input, 40)
}

func calculate(input []string, maxSteps int) int {
	template := input[0]
	var pairList []string
	polymerFreq := make(map[string]int, 0)
	rules := readRules(input[2:])

	for i := 0; i < len(template)-1; i++ {
		pair := template[i : i+2]

		if _, ok := polymerFreq[pair]; !ok {
			pairList = append(pairList, pair)
		}

		polymerFreq[pair]++
	}

	for steps := 0; steps < maxSteps; steps++ {
		pairList, polymerFreq = countFrequency(pairList, polymerFreq, rules)
	}

	count := make(map[string]int, 0)
	for p, o := range polymerFreq {
		count[string(p[0])] += o
		count[string(p[1])] += o
	}

	count[string(template[0])]++
	count[string(template[len(template)-1])]++

	for c, o := range count {
		count[c] = int(math.RoundToEven(float64(o / 2)))
	}

	least, most := leastAndMostCharCount(count)

	return most - least
}

func readRules(input []string) map[string]string {
	pm := make(map[string]string, 0)

	for _, l := range input {
		v := strings.Split(l, " -> ")

		pm[v[0]] = v[1]
	}

	return pm
}

func countFrequency(pairList []string, freq map[string]int, rules map[string]string) ([]string, map[string]int) {
	nFreq := copyMap(freq)
	nPairList := pairList

	for _, p := range pairList {
		for from, to := range rules {
			if p == from {
				nPStart := string(p[0]) + to
				nPEnd := to + string(p[1])

				_, startExists := nFreq[nPStart]
				if !startExists {
					nPairList = append(nPairList, nPStart)
				}

				_, endExists := nFreq[nPEnd]
				if !endExists {
					nPairList = append(nPairList, nPEnd)
				}

				o := freq[p]
				nFreq[p] -= o
				nFreq[nPStart] += o
				nFreq[nPEnd] += o

				break
			}

		}
	}

	return nPairList, nFreq
}

func leastAndMostCharCount(charMap map[string]int) (least int, most int) {
	for _, v := range charMap {
		least = v
		break
	}

	for _, amount := range charMap {
		if amount < least {
			least = amount
		}

		if amount > most {
			most = amount
		}
	}

	return least, most
}

func copyMap(h map[string]int) map[string]int {
	c := make(map[string]int, 0)

	for k, v := range h {
		c[k] = v
	}

	return c
}
