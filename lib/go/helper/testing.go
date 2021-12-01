package helper

import "fmt"

type TestingValue struct {
	Result interface{}
	Expect interface{}
}

func (v *TestingValue) isEqual() bool {
	return v.Result == v.Expect
}

func TestResults(values []TestingValue) {
	if len(values) == 0 {
		fmt.Println("Skipped tests!")
	}

	for i, tv := range values {
		pos := i + 1

		if tv.isEqual() {
			fmt.Printf("\033[32mPart %d: passed\033[37m\n", pos)
			continue
		}

		fmt.Printf(
			"\033[31mPart %d: failed with Expected: '%v' but get '%v'. \033[37m\n",
			pos,
			tv.Expect,
			tv.Result,
		)
	}
}
