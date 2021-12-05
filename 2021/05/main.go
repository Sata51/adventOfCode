package main

import (
	"log"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	dt := utils.Load(2021, "05-real")
	input := dt.ToStringSlice()

	moves := parseInput(input)
	xMax, yMax := getGridSize(moves)

	part1(moves, xMax, yMax)
	part2(moves, xMax, yMax)

}

func part1(moves []move, xMax, yMax int) {
	grid := makeGrid(xMax, yMax)
	hvMoves := keepOnlyHorizontalAndVertical(moves)

	// log.Printf("source: %v || %d", moves, len(moves))
	// log.Printf("filtered: %v || %d", hvMoves, len(hvMoves))

	// log.Printf("\nInitial grid:\n%v\n", grid)

	for _, m := range hvMoves {
		m.getAffected()
		// log.Printf("\nProcessing move: %v\n", m)
		grid.markAffected(m)
		// log.Printf("\nGrid after move: \n%v\n", grid)
	}

	// log.Printf("\nFinal grid:\n%v\n", grid)

	dangerous := 0
	for _, p := range grid.positions {
		if p >= 2 {
			dangerous++
		}
	}
	log.Printf("Dangerous positions: %d\n", dangerous)
}

func part2(moves []move, xMax, yMax int) {
	grid := makeGrid(xMax, yMax)

	// log.Printf("source: %v || %d", moves, len(moves))

	// log.Printf("\nInitial grid:\n%v\n", grid)

	for _, m := range moves {
		m.getAffected()
		// log.Printf("\nProcessing move: %v\n", m)
		grid.markAffected(m)
		// log.Printf("\nGrid after move: \n%v\n", grid)
	}

	log.Printf("\nFinal grid:\n%v\n", grid)

	dangerous := 0
	for _, p := range grid.positions {
		if p >= 2 {
			dangerous++
		}
	}
	log.Printf("Dangerous positions: %d\n", dangerous)
}
