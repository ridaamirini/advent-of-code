package main

import (
	"advent-of-code/helper"
	"fmt"
	"math"
	"path"
	"regexp"
	"runtime"
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
			helper.TestingValue{Result: result01, Expect: 10011},
			helper.TestingValue{Result: result02, Expect: 2994},
		},
	)
}

// Task code
type TargetArea struct {
	xFrom, xTo, yFrom, yTo int
}

func (t *TargetArea) highestPosition() int {
	yf := math.Abs(float64(t.yFrom))
	y := int(yf)

	return (y - 1) * y / 2
}

func (t *TargetArea) checkLaunch(x, y int) bool {
	nX, nY := 0, 0

	for nX <= t.xTo && nY >= t.yFrom {
		nX, nY = nX+x, nY+y

		if t.hits(nX, nY) {
			return true
		}

		if x > 0 {
			x--
		}

		y--
	}

	return false
}

func (t *TargetArea) hits(x, y int) bool {
	return x >= t.xFrom && x <= t.xTo && y >= t.yFrom && y <= t.yTo
}

func part01(input []string) int {
	ta := createTargetArea(input[0])

	return ta.highestPosition()
}

func part02(input []string) int {
	ta := createTargetArea(input[0])
	var hits int

	for y := ta.yFrom; y <= int(math.Abs(float64(ta.yFrom))); y++ {
		for x := 0; x <= ta.xTo; x++ {
			if ta.checkLaunch(x, y) {
				hits++
			}
		}
	}

	return hits
}

func createTargetArea(input string) TargetArea {
	pattern := regexp.MustCompile(`(-?\d+)..(-?\d+)`)
	res := pattern.FindAllStringSubmatch(input, 2)

	t := TargetArea{}
	t.xFrom, _ = strconv.Atoi(res[0][1])
	t.xTo, _ = strconv.Atoi(res[0][2])
	t.yFrom, _ = strconv.Atoi(res[1][1])
	t.yTo, _ = strconv.Atoi(res[1][2])

	return t
}
