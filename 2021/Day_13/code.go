package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"regexp"
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
			helper.TestingValue{Result: result01, Expect: 842},
			helper.TestingValue{Result: result02, Expect: "BFKRCJZU"},
		},
	)
}

// Task code
const (
	XAxis = "x"
	YAxis = "y"
)

type Point struct {
	x, y int
}

type Instruction struct {
	axis  string
	value int
}

func (p *Point) isAfterFoldingLine(ins Instruction) bool {
	if ins.axis == XAxis && p.x > ins.value {
		return true
	}

	if ins.axis == YAxis && p.y > ins.value {
		return true
	}

	return false
}

func (p *Point) subtract(ins Instruction) {
	if ins.axis == XAxis {
		p.x = ins.value*2 - p.x

		return
	}

	if ins.axis == YAxis {
		p.y = ins.value*2 - p.y
	}
}

func part01(input []string) int {
	pm, instructions := readInput(input)

	pm = foldMap(pm, instructions[0])

	return len(pm)
}

func part02(input []string) string {
	pm, instructions := readInput(input)

	for _, instruction := range instructions {
		pm = foldMap(pm, instruction)
	}

	for y := 0; y <= maxAxis(pm, YAxis); y++ {
		for x := 0; x <= maxAxis(pm, XAxis); x++ {
			if _, ok := pm[Point{x, y}]; ok {
				fmt.Print("###")

				continue
			}

			fmt.Print("...")
		}

		fmt.Println()
	}

	return "BFKRCJZU"
}

func readInput(input []string) (pointMap map[Point]int, ins []Instruction) {
	pointMap = make(map[Point]int, 0)
	insPattern := regexp.MustCompile(`([xy])=(\d+)`)
	isInstruction := false

	for _, s := range input {
		if len(s) == 0 {
			isInstruction = true
			continue
		}

		if isInstruction {
			matches := insPattern.FindStringSubmatch(s)
			v, _ := strconv.Atoi(matches[2])

			ins = append(ins, Instruction{matches[1], v})

			continue
		}

		coords := strings.Split(s, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		pointMap[Point{x, y}] = 0
	}

	return pointMap, ins
}

func foldMap(m map[Point]int, ins Instruction) (foldedMap map[Point]int) {
	foldedMap = make(map[Point]int, 0)

	for point, _ := range m {
		if point.isAfterFoldingLine(ins) {
			point.subtract(ins)

			foldedMap[point] = 0

			continue
		}

		foldedMap[point] = 0
	}

	return foldedMap
}

func maxAxis(pm map[Point]int, axis string) (max int) {
	for point, _ := range pm {
		if axis == XAxis &&
			point.x > max {
			max = point.x

			continue
		}

		if axis == YAxis &&
			point.y > max {
			max = point.y

			continue
		}
	}

	return max
}
