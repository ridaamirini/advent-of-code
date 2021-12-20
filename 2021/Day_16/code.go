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
			helper.TestingValue{Result: result01, Expect: 984},
			helper.TestingValue{Result: result02, Expect: 1015320896946},
		},
	)
}

// Task code
const (
	TypeLiteralValue = 4
)

type Packet struct {
	version      int
	typeID       int
	lengthTypeID int
	length       int
	rawValue     int
	children     []Packet
}

func (p *Packet) versionSum() int {
	sum := p.version

	for _, cp := range p.children {
		sum += cp.versionSum()
	}

	return sum
}

func (p *Packet) determineValue() int {
	switch p.typeID {
	case 0:
		{
			var result int
			for _, cp := range p.children {
				result += cp.determineValue()
			}

			return result
		}
	case 1:
		{
			result := 1
			for _, cp := range p.children {
				result *= cp.determineValue()
			}

			return result
		}
	case 2:
		{
			min := p.children[0].determineValue()
			for _, cp := range p.children[1:] {
				val := cp.determineValue()
				if val < min {
					min = val
				}
			}

			return min
		}
	case 3:
		{
			max := p.children[0].determineValue()
			for _, cp := range p.children[1:] {
				val := cp.determineValue()
				if val > max {
					max = val
				}
			}

			return max
		}
	case 4:
		return p.rawValue
	case 5:
		{
			firstVal := p.children[0].determineValue()
			twoVal := p.children[1].determineValue()

			if firstVal > twoVal {
				return 1
			}

			return 0
		}
	case 6:
		{
			firstVal := p.children[0].determineValue()
			twoVal := p.children[1].determineValue()

			if firstVal < twoVal {
				return 1
			}

			return 0
		}
	case 7:
		{
			firstVal := p.children[0].determineValue()
			twoVal := p.children[1].determineValue()

			if firstVal == twoVal {
				return 1
			}

			return 0
		}
	default:
		return 0
	}
}

func part01(input []string) int {
	pos := 0
	p := parsePacket(hex2bin(input[0]), &pos)

	return p.versionSum()
}

func part02(input []string) int {
	pos := 0
	p := parsePacket(hex2bin(input[0]), &pos)

	return p.determineValue()
}

func parsePacket(ps string, idx *int) Packet {
	version, typeId := ps[*idx:*idx+3], ps[*idx+3:*idx+6]
	p := Packet{version: bin2int(version), typeID: bin2int(typeId)}
	*idx += 6

	if p.typeID == TypeLiteralValue {
		var vb string
		var check byte

		for check != '0' {
			vb += ps[*idx+1 : *idx+5]
			check = ps[*idx]
			*idx += 5
		}

		p.rawValue = bin2int(vb)
	} else {
		p.lengthTypeID, _ = strconv.Atoi(string(ps[*idx]))
		*idx++

		if p.lengthTypeID == 0 {
			p.length = bin2int(ps[*idx : *idx+15])
			*idx += 15
			endIdx := *idx + p.length

			for *idx < endIdx {
				cp := parsePacket(ps, idx)
				p.children = append(p.children, cp)
			}
		} else {
			p.length = bin2int(ps[*idx : *idx+11])
			*idx += 11

			for i := 0; i < p.length; i++ {
				cp := parsePacket(ps, idx)
				p.children = append(p.children, cp)
			}
		}
	}

	return p
}

func hex2bin(hex string) (bin string) {
	l := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	for _, i := range hex {
		bin += l[string(i)]
	}

	return bin
}

func bin2int(bin string) int {
	res, _ := strconv.ParseInt(bin, 2, 64)

	return int(res)
}
