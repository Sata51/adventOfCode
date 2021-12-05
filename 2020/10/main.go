package main

import (
	"fmt"
	"sort"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	data := utils.Load(2020, "10-real")
	input := data.ToIntSlice()

	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)
	first(input)
	second(input)
}

func first(data []int) {
	fmt.Println(data)
	one, three := getDiffs(data)
	fmt.Println(one, three, "result", one*three)

}

func second(data []int) {
	accumulator := map[int]int{0: 1}
	for _, i := range data {
		accumulator[i] = accumulator[i-1] + accumulator[i-2] + accumulator[i-3]
	}
	fmt.Println(accumulator[data[len(data)-1]])
}

func getDiffs(data []int) (numberOf1, numberOf3 int) {
	numberOf1 = 0
	numberOf3 = 0

	start := 0

	for _, number := range data {
		diff := number - start
		fmt.Println(start, number, diff)
		switch diff {
		case 1:
			numberOf1++
		case 3:
			numberOf3++
		}

		start = number
	}

	return
}
