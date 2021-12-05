package main

import (
	"fmt"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	yes := parseInput(utils.Load(2020, "06-real").String())
	fmt.Printf("yes in input: %d\n", yes)
}

func intersection(a, b []rune) (c []rune) {
	m := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func parseInput(s string) int {
	grouped := strings.Split(s, "\n\n")
	sum := 0
	for _, g := range grouped {
		personned := strings.Split(g, "\n")
		yesPerPerson := make([][]rune, 0)
		for _, p := range personned {
			yesPerPerson = append(yesPerPerson, []rune(p))
		}
		if len(yesPerPerson) == 1 {
			sum += len(yesPerPerson[0])
		} else if len(yesPerPerson) > 1 {
			base := yesPerPerson[0]
			for _, p := range yesPerPerson {
				base = intersection(base, p)
			}
			sum += len(base)
		}

		// inGroup := strings.ReplaceAll(g, "\n", "")
		// lettersInGroup := make(map[rune]struct{})
		// for _, l := range inGroup {
		// 	lettersInGroup[l] = struct{}{}
		// }
		// fmt.Printf("%s %d\n", inGroup, len(lettersInGroup))
		// sum += len(lettersInGroup)
	}
	return sum
}
