package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

const (
	NewLine = "\n"
)

func ReadInput(currentDir string, delimiter string) []string {
	buffer, err := ioutil.ReadFile(fmt.Sprintf("%s/input.txt", currentDir))
	if err != nil {
		panic(err)
	}

	return strings.Split(string(buffer), delimiter)
}

func ExecutionTime(duration time.Duration) string {
	return fmt.Sprintf("%dms", duration.Milliseconds())
}

func SaveBenchmarkTime(execTime string, currentDir string) {
	fileName := fmt.Sprintf("%s/benchmark.json", currentDir)
	result := make(map[string]interface{})

	data, _ := ioutil.ReadFile(fileName)

	json.Unmarshal(data, &result)

	result["GO"] = execTime

	bufferJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(fileName, bufferJson, 0644)
}

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
			"\033[31mPart %d: failed with Expected: '%s' but get '%s'. \033[37m\n",
			pos,
			tv.Expect,
			tv.Result,
		)
	}
}
