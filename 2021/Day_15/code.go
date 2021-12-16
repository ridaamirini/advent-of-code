package main

import (
	"advent-of-code/helper"
	"container/heap"
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
			helper.TestingValue{Result: result01, Expect: 393},
			helper.TestingValue{Result: result02, Expect: 2823},
		},
	)
}

// Task code
var (
	Up         = Point{0, -1}
	Down       = Point{0, 1}
	Left       = Point{-1, 0}
	Right      = Point{1, 0}
	Directions = [4]Point{Up, Down, Left, Right}
)

type Point struct {
	x, y int
}

type PointCosts struct {
	Point
	costs int
}

type PriorityQueue []*PointCosts

func (p *Point) add(position Point) Point {
	return Point{p.x + position.x, p.y + position.y}
}

func (p Point) neighbors() (n [4]Point) {
	for i, d := range Directions {
		n[i] = p.add(d)
	}

	return n
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i].costs < pq[j].costs }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*PointCosts)) }

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	pw := old[n-1]
	*pq = old[0 : n-1]
	return pw
}

func part01(input []string) int {
	riskLevelMap := createMap(input, 1)
	startPoint := Point{0, 0}
	endPoint := Point{len(input[0]) - 1, len(input) - 1}

	return dijkstra(riskLevelMap, startPoint, endPoint)
}

func part02(input []string) int {
	scale := 5
	riskLevelMap := createMap(input, scale)
	startPoint := Point{0, 0}
	endPoint := Point{len(input[0])*scale - 1, len(input)*scale - 1}

	return dijkstra(riskLevelMap, startPoint, endPoint)
}

func createMap(input []string, scale int) (m map[Point]int) {
	m = make(map[Point]int, 0)
	tileXMax, tileYMax := len(input[0]), len(input)
	mXMax, mYMax := tileXMax*scale, tileYMax*scale

	for y := 0; y < mYMax; y++ {
		for x := 0; x < mXMax; x++ {
			nx := x % tileXMax
			ny := y % tileYMax
			v, _ := strconv.Atoi(string(input[ny][nx]))

			if y >= tileYMax || x >= tileXMax {
				v++

				if prevLeftV, ok := m[Point{x - tileXMax, y}]; ok {
					v = prevLeftV + 1
				} else if prevUpV, ok := m[Point{x, y - tileYMax}]; ok {
					v = prevUpV + 1
				}

				if v > 9 {
					v = 1
				}
			}

			m[Point{x, y}] = v
		}
	}

	return m
}

func dijkstra(m map[Point]int, p Point, ep Point) int {
	costs := make(map[Point]int, 0)
	pq := PriorityQueue([]*PointCosts{{p, 0}})

	heap.Init(&pq)

	visited := map[Point]bool{}

	for pq.Len() > 0 {
		pc := heap.Pop(&pq).(*PointCosts)
		cp := Point{pc.x, pc.y}

		if visited[cp] {
			continue
		}

		visited[cp] = true
		costs[cp] = pc.costs

		if cp == ep {
			break
		}

		for _, point := range cp.neighbors() {
			c, ok := m[point]

			if !ok || visited[point] {
				continue
			}

			heap.Push(&pq, &PointCosts{point, pc.costs + c})
		}
	}

	return costs[ep]
}
