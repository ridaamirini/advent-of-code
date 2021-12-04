package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
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
			helper.TestingValue{Result: result01, Expect: 44088},
			helper.TestingValue{Result: result02, Expect: 23670},
		},
	)
}

// Task code
type Board struct {
	rows []*Row
}

type Row struct {
	cols []*Column
}

type Column struct {
	value  int
	marked bool
}

type Winner struct {
	board Board
	place int
	bingo int
}

func (b *Board) addLine(line string) {
	row := &Row{}

	for _, n := range strings.Fields(line) {
		value, _ := strconv.Atoi(n)

		row.cols = append(row.cols, &Column{value: value})
	}

	b.rows = append(b.rows, row)
}

func (b *Board) markNumber(number int) {
	for _, row := range b.rows {
		for _, col := range row.cols {
			if col.value == number {
				col.marked = true
			}
		}
	}
}

func (b *Board) winner() bool {
	rowLength := len(b.rows[0].cols)
	colsLength := len(b.rows)

	for i := 0; i < colsLength; i++ {
		inColumn := 0

		for _, row := range b.rows {
			if row.cols[i].marked {
				inColumn++
			}
		}

		if inColumn == rowLength {
			return true
		}
	}

	for _, row := range b.rows {
		inRow := 0

		for _, col := range row.cols {
			if col.marked {
				inRow++
			}
		}

		if inRow == colsLength {
			return true
		}
	}

	return false
}

func part01(input []string) int {
	numbers := strings.Split(input[0], ",")
	boards := createBoards(input[2:])

	var bingo int
	var winningBoard *Board

outer:
	for i := 0; i < len(numbers); i++ {
		n, _ := strconv.Atoi(numbers[i])

		for _, board := range boards {
			board.markNumber(n)

			if board.winner() {
				bingo = n
				winningBoard = board
				break outer
			}
		}
	}

	var unmarkedNumbers int
	for _, row := range winningBoard.rows {
		for _, col := range row.cols {
			if !col.marked {
				unmarkedNumbers += col.value
			}
		}
	}

	return unmarkedNumbers * bingo
}

func part02(input []string) int {
	numbers := strings.Split(input[0], ",")
	boards := createBoards(input[2:])
	winningBoards := make(map[string]Winner, 0)

	for i := 0; i < len(numbers); i++ {
		n, _ := strconv.Atoi(numbers[i])

		for bi, board := range boards {
			k := strconv.Itoa(bi)
			_, exits := winningBoards[k]

			if !exits {
				board.markNumber(n)

				if board.winner() {
					winningBoards[k] = Winner{
						board: *board,
						place: len(winningBoards) + 1,
						bingo: n,
					}
				}
			}
		}
	}

	var lastWinner Winner
	for _, w := range winningBoards {
		if w.place == len(winningBoards) {
			lastWinner = w
			break
		}
	}

	var unmarkedNumbers int
	for _, row := range lastWinner.board.rows {
		for _, col := range row.cols {
			if !col.marked {
				unmarkedNumbers += col.value
			}
		}
	}

	return unmarkedNumbers * lastWinner.bingo
}

func createBoards(data []string) map[int]*Board {
	boards := make(map[int]*Board, 0)
	bCount := 0

	for _, l := range data {
		if len(l) == 0 {
			bCount++
			continue
		}

		board, exists := boards[bCount]
		if !exists {
			board = &Board{}
			boards[bCount] = board
		}

		board.addLine(l)
	}

	return boards
}
