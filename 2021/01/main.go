package main

import (
	"log"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	dt := utils.Load(2021, "01-real")
	numbers := dt.ToIntSlice()

	part1(numbers)
	part2(numbers)

}

func part1(numbers []int) {

	// First number is not part of the loop
	prev := numbers[0]
	increased := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > prev {
			increased++
		}
		prev = numbers[i]
	}

	log.Printf("Increased %d times", increased)
}

func part2(numbers []int) {
	//Transform as a windowed array
	windowed := make([]int, 0)

	for i := 0; i < len(numbers)-2; i++ {
		current := numbers[i]
		next := numbers[i+1]
		nextNext := numbers[i+2]

		windowed = append(windowed, current+next+nextNext)
	}

	prev := windowed[0]
	increased := 0
	for i := 1; i < len(windowed); i++ {
		if windowed[i] > prev {
			increased++
		}
		prev = windowed[i]
	}

	log.Printf("Increased %d times", increased)

}
