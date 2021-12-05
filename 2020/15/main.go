package main

import (
	"fmt"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type test struct {
	input  string
	result int
}

func main() {
	tests := []test{
		{input: "0,3,6", result: 436},
		{input: "1,3,2", result: 1},
		{input: "2,1,3", result: 10},
		{input: "1,2,3", result: 27},
		{input: "2,3,1", result: 78},
		{input: "3,2,1", result: 438},
		{input: "3,1,2", result: 1836},
	}

	for _, t := range tests {
		p1 := process(t.input, 2020)
		if t.result != p1 {
			panic("Wrong algo")
		}
	}
	fmt.Println(process("2,0,6,12,1,3", 2020))
	fmt.Println(process("2,0,6,12,1,3", 30000000))
}

func process(input string, maxTurn int) int {
	strNums := strings.Split(input, ",")
	nums := []int{}
	for _, s := range strNums {
		nums = append(nums, utils.MustParseInt(s))
	}
	turn := 0
	num := 0
	lomemSize := 10000000
	mem := map[int]int{}
	lomem := make([]int, lomemSize)

	for ; turn < len(nums); turn++ {
		num = nums[turn]
		lomem[num] = turn + 1
	}
	for ; turn != maxTurn; turn++ {
		var (
			m  int
			ok bool
		)
		if num < lomemSize {
			m = lomem[num]
			ok = m != 0
			lomem[num] = turn
		} else {
			m, ok = mem[num]
			mem[num] = turn
		}
		if !ok {
			num = 0
		} else {
			num = turn - m
		}
	}

	return num
}
