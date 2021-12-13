package main

import (
	"log"
	"math"
	"sort"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	dt := utils.Load(2021, "07-real")

	positionsInStr := strings.Split(dt.String(), ",")
	positionsInInt := make([]int, 0)
	for _, v := range positionsInStr {
		positionsInInt = append(positionsInInt, utils.MustParseInt(v))
	}

	// Part 1
	part1(positionsInInt)
	part2(positionsInInt)
}

func part1(crabPosition []int) {
	// Sort the source array
	sort.Ints(crabPosition)
	// For all positions between the lowest and the highest int

	lowest := crabPosition[0]
	highest := crabPosition[len(crabPosition)-1]
	lowestCost := math.MaxInt64
	for i := lowest; i <= highest; i++ {
		// count required moves to reach the target position (i)
		count := 0
		for _, v := range crabPosition {
			count += utils.Abs(v - i)
		}
		if count < lowestCost {
			lowestCost = count
		}
	}

	log.Printf("Lowest cost: %d", lowestCost)
}

func part2(crabPosition []int) {
	// Sort the source array
	sort.Ints(crabPosition)
	// For all positions between the lowest and the highest int

	lowest := crabPosition[0]
	highest := crabPosition[len(crabPosition)-1]
	lowestCost := math.MaxInt64
	for i := lowest; i <= highest; i++ {
		// count required moves to reach the target position (i)
		count := 0
		for _, v := range crabPosition {
			theSum := utils.Abs(v - i)
			count += (theSum * (theSum + 1)) / 2
		}
		// log.Printf("Alignment to position %d: cost %d", i, count)
		if count < lowestCost {
			lowestCost = count
		}
	}

	log.Printf("Lowest cost: %d", lowestCost)
}
