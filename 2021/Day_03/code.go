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

	helper.TestResults(
		[]helper.TestingValue{
			helper.TestingValue{Result: result01, Expect: 845186},
			helper.TestingValue{Result: result02, Expect: 4636702},
		},
	)
}

// Task code
func part01(input []string) int {
	var gammaBin, epsilonBin string

	byteLength := len(input[0])
	for i := 0; i < byteLength; i++ {
		mb, _ := commonBit(byteByPos(input, i))
		var lb uint8
		if mb == '0' {
			lb = '1'
		} else {
			lb = '0'
		}

		gammaBin += string(mb)
		epsilonBin += string(lb)
	}

	gamma, _ := strconv.ParseUint(gammaBin, 2, byteLength)
	epsilon, _ := strconv.ParseUint(epsilonBin, 2, byteLength)

	return int(gamma) * int(epsilon)
}

func part02(input []string) int {
	byteLength := len(input[0])
	oxygen, _ := strconv.ParseUint(filterBytes(input, true), 2, byteLength)
	carbon, _ := strconv.ParseUint(filterBytes(input, false), 2, byteLength)

	return int(oxygen) * int(carbon)
}

func commonBit(byte string) (bit uint8, isEqual bool) {
	zero := strings.Count(byte, "0")
	one := strings.Count(byte, "1")

	if zero == one {
		return '0', true
	}

	if zero > one {
		return '0', false
	}

	return '1', false
}

func byteByPos(bytes []string, pos int) string {
	r := ""

	for _, b := range bytes {
		r += string(b[pos])
	}

	return r
}

func filterBytes(input []string, mostCommon bool) string {

	bytes := input
	for pos := 0; pos < len(bytes[0]); pos++ {
		if len(bytes) == 1 {
			break
		}

		mb, eq := commonBit(byteByPos(bytes, pos))
		var lb uint8
		if mb == '0' {
			lb = '1'
		} else {
			lb = '0'
		}

		buffer := make([]string, 0)
		for _, b := range bytes {
			if eq {
				if mostCommon {
					if b[pos] == '1' {
						buffer = append(buffer, b)
					}

					continue
				}

				if b[pos] == '0' {
					buffer = append(buffer, b)
					continue
				}

				continue
			}

			if !mostCommon {
				if b[pos] == lb {
					buffer = append(buffer, b)
				}

				continue
			}

			if b[pos] == mb {
				buffer = append(buffer, b)
			}
		}
		bytes = buffer
	}

	return bytes[0]
}
