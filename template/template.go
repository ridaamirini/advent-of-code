package main

import (
	"advent-of-code/helper"
	"fmt"
	"path"
	"runtime"
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
	//result02 := part02(iValues)
	executionTime := helper.ExecutionTime(time.Since(start))

	fmt.Printf("Solution Part 1: %v\n", result01)
	//fmt.Printf("Solution Part 2: %v\n", result02)
	fmt.Printf("Execution time: %s\n", executionTime)

	helper.SaveBenchmarkTime(executionTime, cwd)

	// Testing
	//helper.TestResults(
	//	[]helper.TestingValue{
	//		helper.TestingValue{Result: result01, Expect: nil},
	//		helper.TestingValue{Result: result02, Expect: nil},
	//	},
	//)
}

// Task code
func part01(input []string) {

}

//func part02(input []string) {
//
//}
