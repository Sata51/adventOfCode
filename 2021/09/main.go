package main

import (
	"log"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	dt := utils.Load(2021, "09-fake")
	heightmap := make([][]int, 0)

	for _, line := range dt.ToStringSlice() {
		atLine := make([]int, 0)
		for _, number := range strings.Split(line, "") {
			atLine = append(atLine, utils.MustParseInt(number))
		}
		heightmap = append(heightmap, atLine)
	}

	part1(heightmap)
	part2()
}

func part1(heightmap [][]int) {
	// log.Println(heightmap)
	lowPoints := make([]int, 0)
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			nearest := getNearest(heightmap, i, j)
			isLowPoint := true
			for _, n := range nearest {
				if n <= heightmap[i][j] {
					isLowPoint = false
				}
			}
			// log.Println(heightmap[i][j], nearest, isLowPoint)
			if isLowPoint {
				lowPoints = append(lowPoints, heightmap[i][j])
			}
		}
	}

	// log.Println(lowPoints)
	riskLevel := make([]int, 0)
	for _, lp := range lowPoints {
		riskLevel = append(riskLevel, lp+1)
	}
	// log.Println(riskLevel)
	log.Printf("Part 1 : %d", utils.SumSlice(riskLevel))
}

type position struct {
	x int
	y int
}

func getNearest(heightmap [][]int, x, y int) []int {
	nearest := make([]int, 0)
	upPosition := position{x: x - 1, y: y}
	downPosition := position{x: x + 1, y: y}
	leftPosition := position{x: x, y: y - 1}
	rightPosition := position{x: x, y: y + 1}

	if isValidPosition(heightmap, upPosition) {
		nearest = append(nearest, heightmap[upPosition.x][upPosition.y])
	}
	if isValidPosition(heightmap, downPosition) {
		nearest = append(nearest, heightmap[downPosition.x][downPosition.y])
	}
	if isValidPosition(heightmap, leftPosition) {
		nearest = append(nearest, heightmap[leftPosition.x][leftPosition.y])
	}
	if isValidPosition(heightmap, rightPosition) {
		nearest = append(nearest, heightmap[rightPosition.x][rightPosition.y])
	}

	return nearest
}

func isValidPosition(heightmap [][]int, p position) bool {
	return p.x >= 0 && p.x < len(heightmap) && p.y >= 0 && p.y < len(heightmap[p.x])
}

func part2() {

}
