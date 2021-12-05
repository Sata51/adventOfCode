package main

import (
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type Opcode string

const OpcodeForward = "forward" // Move forward one step.
const OpcodeUP = "up"           // Decrease the depth
const OpcodeDOWN = "down"       // Increase the depth

type (
	Instruction struct {
		Opcode Opcode
		times  int
	}
	Submarine struct {
		horizontalPosition int
		verticalPosition   int
		aim                int
	}
)

func (s *Submarine) Move1(instruction Instruction) {
	switch instruction.Opcode {
	case OpcodeForward:
		s.horizontalPosition += instruction.times
	case OpcodeUP:
		s.verticalPosition -= instruction.times
	case OpcodeDOWN:
		s.verticalPosition += instruction.times
	}
}

func (s *Submarine) Move2(instruction Instruction) {
	switch instruction.Opcode {
	case OpcodeForward:
		s.horizontalPosition += instruction.times
		s.verticalPosition += s.aim * instruction.times
	case OpcodeUP:
		s.aim -= instruction.times
	case OpcodeDOWN:
		s.aim += instruction.times

	}
}

func parseInstruction(s string) Instruction {
	spl := strings.Split(s, " ")
	return Instruction{
		Opcode: Opcode(spl[0]),
		times:  utils.MustParseInt(spl[1]),
	}
}
