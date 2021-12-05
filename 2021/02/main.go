package main

import (
	"log"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	dt := utils.Load(2021, "02-real")
	instructionsStr := dt.ToStringSlice()
	instructions := make([]Instruction, 0)
	for _, s := range instructionsStr {
		instructions = append(instructions, parseInstruction(s))
	}

	part1(instructions)
	part2(instructions)

}

func part1(instructions []Instruction) {

	submarine := &Submarine{}

	for _, instruction := range instructions {
		submarine.Move1(instruction)
	}

	log.Printf("Final position is %d, %d, so result is %d", submarine.horizontalPosition, submarine.verticalPosition, submarine.horizontalPosition*submarine.verticalPosition)
}

func part2(instructions []Instruction) {
	submarine := &Submarine{}

	for _, instruction := range instructions {
		submarine.Move2(instruction)
	}

	log.Printf("Final position is %d, %d, so result is %d", submarine.horizontalPosition, submarine.verticalPosition, submarine.horizontalPosition*submarine.verticalPosition)
}
