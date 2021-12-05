package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

var store = make(map[string]map[string]int)
var lookFor = "shiny gold"

func main() {
	parse(utils.Load(2020, "07-real").ToStringSlice())

	result := make(map[string]struct{})

	for bag := range store {
		if bag == lookFor {
			continue
		}
		if canCarry(bag, 0) {
			result[bag] = struct{}{}
		}
	}

	// fmt.Printf("subsequent %d\n", len(result))

	rslt := countBag(lookFor, 1)

	fmt.Printf("%v\n", rslt)
}

func countBag(bag string, level int) int {
	// fmt.Println(level, bag)
	sum := 0
	if val, ok := store[bag]; ok {
		if len(val) == 0 {
			return 0
		}
		for subBags, quantity := range val {
			subCount := countBag(subBags, level+1)
			fmt.Println(strings.Repeat("\t", level), quantity, subBags, subCount)
			fmt.Printf("%d + %d*%d \n", quantity, quantity, subCount)
			if subCount == 0 {
				sum += quantity
			} else {
				sum += quantity + quantity*subCount
			}
		}
	}

	return sum
}

func canCarry(bag string, level int) bool {
	if val, ok := store[bag]; ok {
		if len(val) == 0 {
			return false
		}
		for childBag := range val {
			// fmt.Printf("%s%s %d: evaluate %s\n", strings.Repeat("\t", level), bag, level, childBag)
			if childBag == lookFor {
				// fmt.Printf("%sFound %s - %s\n", strings.Repeat("\t", level), bag, childBag)
				return true
			}
			if canCarry(childBag, level+1) {
				return true
			}
		}
	}
	return false
}

func parse(lines []string) {
	for _, line := range lines {
		line = strings.ReplaceAll(line, "bags", "")
		line = strings.ReplaceAll(line, "contain", "")
		line = strings.ReplaceAll(line, "bag", "")
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, "  ", " ")
		splitted := strings.Split(line, " ")
		bagName := splitted[:2]
		splitted = splitted[2:]
		clean := make([]string, 0)
		for _, s := range splitted {
			s = strings.TrimSpace(s)
			if len(s) != 0 {
				clean = append(clean, s)
			}
		}
		bag := strings.TrimSpace(strings.Join(bagName, " "))

		subBags := make(map[string]int)

		// fmt.Printf("bag: %s %v %d\n", bag, clean, len(clean))
		for len(clean) >= 3 {
			quantity := clean[0]
			bagName := clean[1:3]
			quantityInt, err := strconv.Atoi(quantity)
			if err != nil {
				panic(err)
			}
			subBags[strings.TrimSpace(strings.Join(bagName, " "))] = quantityInt
			clean = clean[3:]
		}
		store[bag] = subBags
	}
}
