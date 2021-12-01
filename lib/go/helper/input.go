package helper

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

const (
	NewLine = "\n"
)

type Input struct {
	data []string
}

func ReadInput(currentDir string, delimiter string) *Input {
	buffer, err := ioutil.ReadFile(fmt.Sprintf("%s/input.txt", currentDir))
	if err != nil {
		panic(err)
	}

	return &Input{strings.Split(string(buffer), delimiter)}
}

func (i *Input) Strings() []string {
	return i.data
}

func (i *Input) Ints() []int {
	result := make([]int, 0)

	for _, l := range i.data {
		value, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		result = append(result, value)
	}

	return result
}

func (i *Input) Int64s(base int) []int64 {
	result := make([]int64, 0)

	for _, l := range i.data {
		value, err := strconv.ParseInt(l, base, 64)
		if err != nil {
			panic(err)
		}

		result = append(result, value)
	}

	return result
}

func (i *Input) Floats() []float64 {
	result := make([]float64, 0)

	for _, l := range i.data {
		value, err := strconv.ParseFloat(l, 64)
		if err != nil {
			panic(err)
		}

		result = append(result, value)
	}

	return result
}

func (i *Input) BigInts(base int) []*big.Int {
	result := make([]*big.Int, 0)

	for _, l := range i.data {
		value := new(big.Int)
		_, ok := value.SetString(l, base)
		if !ok {
			panic("Could not parse value to big.Int")
		}

		result = append(result, value)
	}

	return result
}

func (i *Input) BigFloat() []*big.Float {
	result := make([]*big.Float, 0)

	for _, l := range i.data {
		value := new(big.Float)
		_, ok := value.SetString(l)
		if !ok {
			panic("Could not parse value to big.Float")
		}

		result = append(result, value)
	}

	return result
}
