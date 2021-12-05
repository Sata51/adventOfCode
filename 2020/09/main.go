package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	data := parse(utils.Load("09-real").ToStringSlice())
	// fmt.Println(data)

	incorrect := foundError(5, data)
	fmt.Println(incorrect)

	rangeFound := foundContiguous(data, incorrect)
	if rangeFound != nil {
		smallest, largest := getSmallestAndLargest(rangeFound)
		fmt.Println(rangeFound, smallest, largest, smallest+largest)
	}
}

func getSmallestAndLargest(data []int) (min, max int) {
	min = math.MaxInt64
	max = math.MinInt64
	for _, v := range data {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}

func getContiguous(data []int, start, size int) []int {
	return data[start : start+size]
}
func sum(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	return sum
}

func foundContiguous(data []int, incorrect int) []int {
	for i := 2; i < len(data); i++ {
		for j := 0; j < len(data)-i; j++ {
			contiguous := getContiguous(data, j, i)
			// fmt.Println(contiguous)
			// if len(contiguous) != i {
			// 	fmt.Println("Incorrect length", len(contiguous), i)
			// }
			if sum(contiguous) == incorrect {
				return contiguous
			}
		}
	}
	return nil
}

func foundError(preamble int, data []int) int {
	for i, startData := range data[preamble:] {
		iterator := data[i : i+1+preamble]
		found := false
		// fmt.Println(iterator, i, startData)
	iteratorLoop:
		for _, iData := range iterator {
			for _, jData := range iterator {
				if iData != jData {
					// fmt.Printf("is %d sum of %d and %d -> %t\n", startData, iData, jData, iData+jData == startData)
					if iData+jData == startData {
						found = true
						break iteratorLoop
					}
				}
			}
		}
		if !found {
			return startData
		}
	}
	return -1
}

func parse(lines []string) []int {
	retVal := make([]int, 0)
	for _, line := range lines {
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		retVal = append(retVal, val)
	}
	return retVal
}
