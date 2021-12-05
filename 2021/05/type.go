package main

import (
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func parseInput(input []string) []move {
	moves := make([]move, 0)

	for _, line := range input {
		spl := strings.Split(line, " -> ")
		from := strings.Split(spl[0], ",")
		to := strings.Split(spl[1], ",")

		moves = append(moves, move{
			from: position{
				x: utils.MustParseInt(from[0]),
				y: utils.MustParseInt(from[1]),
			},
			to: position{
				x: utils.MustParseInt(to[0]),
				y: utils.MustParseInt(to[1]),
			},
			affected: make(map[string]struct{}),
		})

	}

	return moves
}
