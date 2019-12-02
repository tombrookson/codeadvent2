package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	defaultFilename = "program.json"
)

func parseFile(file string) ([]int, error) {
	var opcodes []int
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &opcodes)
	return opcodes, err
}

func splitOpcodes(opcodes []int) [][]int {
	var opcodeLines [][]int

	lineSize := 4

	for i := 0; i < len(opcodes); i += lineSize {
		end := i + lineSize

		if end > len(opcodes) {
			end = len(opcodes)
		}

		opcodeLines = append(opcodeLines, opcodes[i:end])
	}

	return opcodeLines
}

func opcodeOne(pos1 int, pos2 int, pos3 int, program []int) {
	program[pos3] = program[pos1] + program[pos2]
}

func opcodeTwo(pos1 int, pos2 int, pos3 int, program []int) {
	program[pos3] = program[pos1] * program[pos2]
}

func main() {
	opcodes, err := parseFile(defaultFilename)
	if err != nil {
		fmt.Println(err)
	}

	opcodeLines := splitOpcodes(opcodes)

	for _, opcodeLine := range opcodeLines {
		optype := opcodeLine[0]

		if optype == 1 {
			opcodeOne(opcodeLine[1], opcodeLine[2], opcodeLine[3], opcodes)
		}
		if optype == 2 {
			opcodeTwo(opcodeLine[1], opcodeLine[2], opcodeLine[3], opcodes)
		}
		if optype == 99 {
			break
		}
	}

	fmt.Println(opcodes)
}
