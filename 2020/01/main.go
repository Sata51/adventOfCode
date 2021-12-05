package main

import (
	"fmt"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {

	raw := utils.Load("01-real")
	input := raw.ToIntSlice()
	fmt.Printf("%d\n", getData(input))

}

func getData(input []int) int {
	for _, i := range input {
		for _, j := range input {
			for _, k := range input {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	return -1
}
