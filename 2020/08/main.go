package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type op string

const (
	opNop op = "nop"
	opAcc op = "acc"
	opJmp op = "jmp"
)

type operation struct {
	kind               op
	value              int
	numberOfExecutions int
}

func main() {
	instructions := parse(utils.Load("08-real").ToStringSlice())

	fmt.Println(instructions)

	number := execute(instructions)

	fmt.Println(number)

	processedAcc := testSequence(utils.Load("08-real").ToStringSlice())
	fmt.Println(processedAcc)

}

func permute(lastChangedIndex int, instructions map[int]operation) (int, map[int]operation) {
	for i := lastChangedIndex; i < len(instructions); i++ {
		hasPermut := false
		thisOperation := instructions[i]
		switch thisOperation.kind {
		case opNop:
			thisOperation.kind = opJmp
			hasPermut = true
		case opJmp:
			thisOperation.kind = opNop
			hasPermut = true
		}

		if hasPermut {
			instructions[i] = thisOperation
			return i + 1, instructions
		}
	}
	return lastChangedIndex, instructions
}

func testSequence(lines []string) int {
	for {
		lastChangeIndex := 0
	resetIteration:
		thisLoopInstruction := parse(lines) // fresh read
		lastChangeIndex, thisLoopInstruction = permute(lastChangeIndex, thisLoopInstruction)
		iterationIndex := 0
		iterationAcc := 0
		for {
			oldIndex := iterationIndex
			if oldIndex >= len(thisLoopInstruction) {
				return iterationAcc
			}
			if thisOperation, ok := thisLoopInstruction[oldIndex]; ok {
				fmt.Printf("Process operation %d %v\n", oldIndex, thisOperation)
				if thisOperation.numberOfExecutions == 1 {
					goto resetIteration
				}
				thisOperation.numberOfExecutions++
				switch thisOperation.kind {
				case opNop:
					iterationIndex++
				case opAcc:
					iterationIndex++
					iterationAcc += thisOperation.value
				case opJmp:
					iterationIndex += thisOperation.value
				}
				thisLoopInstruction[oldIndex] = thisOperation
			} else {
				fmt.Printf("No operation found at index %d\n", iterationIndex)
				return iterationAcc
			}
		}
	}
}

func execute(instructions map[int]operation) int {
	acc := 0
	index := 0

	for {
		oldIndex := index
		if thisOperation, ok := instructions[oldIndex]; ok {
			fmt.Printf("Process operation %d %v\n", oldIndex, thisOperation)
			// block second time execution
			if thisOperation.numberOfExecutions == 1 {
				return acc
			}
			thisOperation.numberOfExecutions++
			switch thisOperation.kind {
			case opNop:
				index++
			case opAcc:
				index++
				acc += thisOperation.value
			case opJmp:
				index += thisOperation.value
			}
			instructions[oldIndex] = thisOperation
		} else {
			fmt.Printf("No operation found at index %d\n", index)
			return acc
		}
	}

}

func parse(lines []string) map[int]operation {
	retVal := make(map[int]operation)
	for i, v := range lines {
		thisOperation := operation{numberOfExecutions: 0}
		opLine := strings.Split(v, " ")
		switch op := opLine[0]; op {
		case "acc":
			thisOperation.kind = opAcc
		case "jmp":
			thisOperation.kind = opJmp
		case "nop":
			thisOperation.kind = opNop
		}
		negative := false
		if opLine[1][0] == '-' {
			negative = true
		}
		rest := opLine[1][1:]
		val, err := strconv.Atoi(rest)
		if err != nil {
			panic(err)
		}
		if negative {
			thisOperation.value = val * -1
		} else {
			thisOperation.value = val
		}

		retVal[i] = thisOperation
	}
	return retVal
}
