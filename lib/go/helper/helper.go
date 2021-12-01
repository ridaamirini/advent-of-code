package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func ExecutionTime(duration time.Duration) string {
	if duration.Milliseconds() == 0 {
		return fmt.Sprintf("%dμs", duration.Microseconds())
	}

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
