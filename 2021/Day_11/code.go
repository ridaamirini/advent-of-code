package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
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
			helper.TestingValue{Result: result01, Expect: 1617},
			helper.TestingValue{Result: result02, Expect: 258},
		},
	)
}

var (
	Up         = Point{0, -1}
	UpLeft     = Point{-1, -1}
	UpRight    = Point{1, -1}
	Down       = Point{0, 1}
	DownLeft   = Point{-1, 1}
	DownRight  = Point{1, 1}
	Left       = Point{-1, 0}
	Right      = Point{1, 0}
	Directions = [8]Point{Up, UpLeft, UpRight, Down, DownLeft, DownRight, Left, Right}
)

type Point struct {
	x, y int
}

func (p *Point) add(position Point) Point {
	return Point{p.x + position.x, p.y + position.y}
}

func (p Point) neighbors() (n [8]Point) {
	for i, d := range Directions {
		n[i] = p.add(d)
	}

	return n
}

// Task code
func part01(input []string) int {
	var flashes int
	dm, dmOrder := createDumboMap(input)

	for step := 1; step <= 100; step++ {
		for point, _ := range dm {
			dm[point]++
		}

		for {
			flashing := make(map[Point]int, 0)

			for _, pp := range dmOrder {
				point := *pp

				if dm[point] > 9 {
					flashing[point] = 0

					for _, p := range point.neighbors() {
						_, ok := dm[p]

						if !ok {
							continue
						}

						if dm[p] != 0 {
							dm[p]++
						}
					}
				}
			}

			if len(flashing) == 0 {
				break
			}

			for point, _ := range flashing {
				dm[point] = 0
			}

			flashes += len(flashing)
		}
	}

	//outputMap(dm, dmOrder, len(input[0]))

	return flashes
}

func part02(input []string) int {
	dm, dmOrder := createDumboMap(input)
	steps := 1

	for {
		var flashes int

		for point, _ := range dm {
			dm[point]++
		}

		for {
			flashing := make(map[Point]int, 0)

			for _, pp := range dmOrder {
				point := *pp

				if dm[point] > 9 {
					flashing[point] = 0

					for _, p := range point.neighbors() {
						_, ok := dm[p]

						if !ok {
							continue
						}

						if dm[p] != 0 {
							dm[p]++
						}
					}
				}
			}

			if len(flashing) == 0 {
				break
			}

			for point, _ := range flashing {
				dm[point] = 0
			}

			flashes += len(flashing)
		}

		if flashes == len(dm) {
			break
		}

		steps++
	}

	//outputMap(dm, dmOrder, len(input[0]))

	return steps
}

func createDumboMap(input []string) (dumboMap map[Point]int, order []*Point) {
	dumboMap = make(map[Point]int, 0)

	for y, line := range input {
		for x, dumbo := range line {
			n, _ := strconv.Atoi(string(dumbo))
			p := Point{x, y}
			dumboMap[p] = n
			order = append(order, &p)
		}
	}

	return dumboMap, order
}

func outputMap(dm map[Point]int, dmOrder []*Point, lineLength int) {
	var out string
	for i, point := range dmOrder {
		if i != 0 && i%lineLength == 0 {
			out += "\n"
		}

		v := dm[*point]

		out += fmt.Sprintf("%d", v)
	}

	fmt.Println(out)
}
