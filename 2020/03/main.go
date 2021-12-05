package main

import (
	"fmt"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type pattern struct {
	right int
	down  int
}

func main() {
	p := []pattern{
		{right: 1, down: 1},
		{right: 3, down: 1},
		{right: 5, down: 1},
		{right: 7, down: 1},
		{right: 1, down: 2},
	}

	data := utils.Load(2020, "03-real")
	row := data.ToStringSlice()

	multipliedResults := 1
	for _, pat := range p {
		result := evalTrees(row, pat.right, pat.down)
		fmt.Printf("%d\n", result)
		multipliedResults *= result
	}

	fmt.Printf("%d\n", multipliedResults)

}

func evalTrees(rows []string, right, down int) int {
	trees := 0
	x := 0
	y := 0
rowLoop:
	for {
		rowAtY := rows[y]
		elementAtX := rowAtY[x%len(rowAtY)]
		if elementAtX == '#' {
			trees++
		}
		x += right
		y += down
		if y >= len(rows) {
			break rowLoop
		}
	}
	return trees
}
