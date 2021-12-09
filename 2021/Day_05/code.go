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
			helper.TestingValue{Result: result01, Expect: 7473},
			helper.TestingValue{Result: result02, Expect: 24164},
		},
	)
}

// Task code
type Diagram struct {
	coords map[Point]int
}

type Point struct {
	x, y int
}

func (d *Diagram) addSegment(l LineSegment) {
	if l.isHorizontal() {
		min, max := l.horizontalRange()

		for x := min; x <= max; x++ {
			d.coords[Point{x, l.y1}]++
		}

		return
	}

	if l.isVertical() {
		min, max := l.verticalRange()

		for y := min; y <= max; y++ {
			d.coords[Point{l.x1, y}]++
		}

		return
	}

	if l.x1 < l.x2 {
		y := l.y1

		for x := l.x1; x <= l.x2; x++ {
			d.coords[Point{x, y}]++

			if l.y1 < l.y2 {
				y++
			} else {
				y--
			}
		}
	}

	if l.x1 > l.x2 {
		y := l.y2

		for x := l.x2; x <= l.x1; x++ {
			d.coords[Point{x, y}]++

			if l.y1 > l.y2 {
				y++
			} else {
				y--
			}
		}
	}
}

func (d *Diagram) countOverlaps() int {
	var count int

	for _, c := range d.coords {
		if c > 1 {
			count++
		}
	}

	return count
}

type LineSegment struct {
	x1, y1, x2, y2 int
}

func (l *LineSegment) horizontalRange() (int, int) {
	n1, n2 := float64(l.x1), float64(l.x2)

	return int(math.Min(n1, n2)), int(math.Max(n1, n2))
}

func (l *LineSegment) verticalRange() (int, int) {
	n1, n2 := float64(l.y1), float64(l.y2)

	return int(math.Min(n1, n2)), int(math.Max(n1, n2))
}

func (l *LineSegment) isStraight() bool {
	return l.isHorizontal() || l.isVertical()
}

func (l *LineSegment) isVertical() bool {
	return l.x1 == l.x2
}

func (l *LineSegment) isHorizontal() bool {
	return l.y1 == l.y2
}

func part01(input []string) int {
	d := &Diagram{coords: make(map[Point]int, 0)}
	pointsPattern := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	for _, l := range input {
		values := pointsPattern.FindStringSubmatch(l)
		x1, _ := strconv.Atoi(values[1])
		y1, _ := strconv.Atoi(values[2])
		x2, _ := strconv.Atoi(values[3])
		y2, _ := strconv.Atoi(values[4])
		seg := LineSegment{x1, y1, x2, y2}

		if !(seg.isVertical() || seg.isHorizontal()) {
			continue
		}

		d.addSegment(seg)
	}

	return d.countOverlaps()
}

func part02(input []string) int {
	d := &Diagram{coords: make(map[Point]int, 0)}
	pointsPattern := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

	for _, l := range input {
		values := pointsPattern.FindStringSubmatch(l)
		x1, _ := strconv.Atoi(values[1])
		y1, _ := strconv.Atoi(values[2])
		x2, _ := strconv.Atoi(values[3])
		y2, _ := strconv.Atoi(values[4])
		seg := LineSegment{x1, y1, x2, y2}

		d.addSegment(seg)
	}

	return d.countOverlaps()
}
