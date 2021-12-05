package main

import (
	"fmt"
	"log"
	"strings"
)

type grid struct {
	positions map[string]int
	xMax      int
	yMax      int
}

func makeGrid(maxX, maxY int) grid {
	g := grid{
		positions: make(map[string]int),
		xMax:      maxX,
		yMax:      maxY,
	}

	for x := 0; x <= g.xMax; x++ {
		for y := 0; y <= g.yMax; y++ {
			g.positions[position{x, y}.String()] = 0
		}
	}

	log.Printf("Create grid of (%d,%d)", g.xMax, g.yMax)

	return g
}

func (g grid) String() string {
	sb := strings.Builder{}

	for y := 0; y <= g.yMax; y++ {
		for x := 0; x <= g.xMax; x++ {
			gridPosition, ok := g.positions[position{x, y}.String()]
			if !ok {
				sb.WriteString("&")
			} else {
				if gridPosition == 0 {
					sb.WriteString(".")
				} else {
					sb.WriteString(fmt.Sprintf("%d", gridPosition))
				}
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (g grid) markAffected(moves move) {
	for k := range moves.affected {
		g.positions[k] += 1
	}
}
