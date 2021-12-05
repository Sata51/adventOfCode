package main

import (
	"fmt"

	"github.com/Sata51/adventOfCode/pkg/utils"
	"github.com/bradfitz/iter"
)

type pocketDimension struct {
	grid map[utils.Vector]bool
}

func newPocketDimension(input []string, dimensions int) pocketDimension {
	pd := pocketDimension{
		grid: make(map[utils.Vector]bool),
	}

	for y, line := range input {
		for x, char := range line {
			if char != '#' {
				continue
			}
			if dimensions == 3 {
				pd.grid[utils.Vector3{X: x, Y: y, Z: 0}] = true
			} else if dimensions == 4 {
				pd.grid[utils.Vector4{X: x, Y: y, Z: 0, W: 0}] = true
			} else {
				panic("Dimensions not implemented")
			}
		}
	}

	return pd
}

func main() {
	input := utils.Load("real").ToStringSlice()

	fmt.Println("----------------Part 1----------------")
	pd3 := newPocketDimension(input, 3)
	fmt.Println("Result after 6 cycles:", pd3.run(6))

	fmt.Println("----------------Part 2----------------")
	pd4 := newPocketDimension(input, 4)
	fmt.Println("Result after 6 cycles:", pd4.run(6))
}

func (pd *pocketDimension) runCycle() {
	searchGrid := make(map[utils.Vector]bool)
	for k, v := range pd.grid {
		searchGrid[k] = v

		// Make sure we also check its neighbors
		for _, neighbor := range k.Nearby() {
			if _, ok := searchGrid[neighbor]; !ok {
				searchGrid[neighbor] = false
			}
		}
	}

	newGrid := make(map[utils.Vector]bool)

	for vector, active := range searchGrid {
		activeNeighbors := 0
		for _, neighbor := range vector.Nearby() {
			if searchGrid[neighbor] {
				activeNeighbors++
			}
		}

		if active && (activeNeighbors == 2 || activeNeighbors == 3) {
			newGrid[vector] = true
		} else if !active && activeNeighbors == 3 {
			newGrid[vector] = true
		}
	}

	pd.grid = newGrid
}

func (pd *pocketDimension) countActive() int {
	retVal := 0
	for _, active := range pd.grid {
		if active {
			retVal++
		}
	}
	return retVal
}

func (pd *pocketDimension) run(cycles int) int {
	for range iter.N(cycles) {
		pd.runCycle()
	}

	return pd.countActive()
}
