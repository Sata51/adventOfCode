package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	dt := utils.Load(2021, "08-real")
	lines := dt.ToStringSlice()

	// Part 1
	part1(lines)
	part2(lines)

}

func part1(lines []string) {
	count := 0
	for _, line := range lines {
		inputs := strings.Split(line, " | ")
		toDecode := strings.Split(inputs[1], " ")

		for _, item := range toDecode {
			length := len(item)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}

	log.Println(count)
}

func part2(lines []string) {
	answer := 0
	for _, line := range lines {
		inputs := strings.Split(line, " | ")
		hints := strToChars(strings.Split(inputs[0], " "))
		outputValues := strToChars(strings.Split(inputs[1], " "))
		// fmt.Println("Hints", hints, "Output values to decode", outputValues)

		// Positions
		// ┌-0-┐
		// 1   2
		// |-3-|
		// 4   5
		// └-6-┘

		// Fill the 7 positions with the 7 possible characters
		positions := make([][]string, 7)
		for i := range positions {
			positions[i] = []string{"a", "b", "c", "d", "e", "f", "g"}
		}

		// 1 is the only number with 2 positions
		// The characters must be in position 2 and 5
		hintWithLength2 := filterLength(hints, 2)[0]
		// fmt.Println("Length 2", hintWithLength2)
		positions[2] = []string{}
		positions[5] = []string{}
		for _, char := range hintWithLength2 {
			positions = removeItemAtEachPositions(positions, char)
			positions[2] = append(positions[2], char)
			positions[5] = append(positions[5], char)
		}
		// fmt.Println("Positions", positions)

		// 7 is the only number with 3 positions
		// The character that is not in the hintWithLength2 must be in position 0
		hintWithLength3 := filterLength(hints, 3)[0]
		// fmt.Println("Length 3", hintWithLength3)
		positions[0] = []string{}
		for _, char := range hintWithLength3 {
			if hasItem(positions[2], char) {
				continue
			}
			positions = removeItemAtEachPositions(positions, char)
			positions[0] = append(positions[0], char)
		}
		// fmt.Println("Positions", positions)

		// 4 is the only number with 4 positions
		// The 2 characters that are not in the hintWithLength2 must be in position 1 and 3
		hintWithLength4 := filterLength(hints, 4)[0]
		// fmt.Println("Length 4", hintWithLength4)
		positions[1] = []string{}
		positions[3] = []string{}
		for _, char := range hintWithLength4 {
			if hasItem(positions[2], char) {
				continue
			}
			positions = removeItemAtEachPositions(positions, char)
			positions[1] = append(positions[1], char)
			positions[3] = append(positions[3], char)
		}
		// fmt.Println("Positions", positions)

		// 0, 6 and 9 are the numbers with 6 positions
		// The 6 is the only number that doesn't have 2 positions that are in the hintWithLength2
		// Use this to find the number 6
		hint6 := []string{}
		hints0or9 := [][]string{}
		hintsWithLength6 := filterLength(hints, 6)
		// fmt.Println("Length 6", hintsWithLength6)
		for _, hint := range hintsWithLength6 {
			if !(hasItem(hint, hintWithLength2[0]) && hasItem(hint, hintWithLength2[1])) {
				hint6 = hint
			} else {
				hints0or9 = append(hints0or9, hint)
			}
		}
		// fmt.Println("Hint 6", hint6)

		// Number 6 has position 5 but not position 2
		// Use this to find out which of the 2 characters are in position 2 and 5
		for _, char := range positions[5] {
			if hasItem(hint6, char) {
				positions[5] = []string{char}
			} else {
				positions[2] = []string{char}
			}
		}
		// fmt.Println("Positions", positions)

		// The 0 is the only 6 positions number that has the 2 characters from position 4
		// Use this to find number 0 and 9
		hint0 := []string{}
		hint9 := []string{}
		// fmt.Println("Hints 0 or 9", hints0or9)
		for _, hint := range hints0or9 {
			if hasItem(hint, positions[4][0]) && hasItem(hint, positions[4][1]) {
				hint0 = hint
			} else {
				hint9 = hint
			}
		}
		// fmt.Println("Hint 0", hint0, "Hint 9", hint9)

		// Number 9 has position 6 but not position 4
		// Use this to find out which of the 2 characters are in position 4 and 6
		for _, char := range positions[4] {
			if hasItem(hint9, char) {
				positions[6] = []string{char}
			} else {
				positions[4] = []string{char}
			}
		}
		// fmt.Println("Positions", positions)

		// Number 0 has position 1 but not position 3
		// Use this to find out which of the 2 characters are in position 1 and 3
		for _, char := range positions[1] {
			if hasItem(hint0, char) {
				positions[1] = []string{char}
			} else {
				positions[3] = []string{char}
			}
		}
		// fmt.Println("Positions", positions)
		// Done finding the right character for all positions!

		// List the positions per number
		posNumbers := [][]int{
			{0, 1, 2, 4, 5, 6},
			{2, 5},
			{0, 2, 3, 4, 6},
			{0, 2, 3, 5, 6},
			{1, 2, 3, 5},
			{0, 1, 3, 5, 6},
			{0, 1, 3, 4, 5, 6},
			{0, 2, 5},
			{0, 1, 2, 3, 4, 5, 6},
			{0, 1, 2, 3, 5, 6},
		}

		// Match the characters to the numbers
		posChars := [][]string{}
		for _, numbers := range posNumbers {
			chars := []string{}
			for _, number := range numbers {
				chars = append(chars, positions[number][0])
			}
			posChars = append(posChars, chars)
		}

		// Convert the output values to single digits by matching the characters
		var outputInts []int

		for _, outputValue := range outputValues {
			for key, posChar := range posChars {
				if len(posChar) != len(outputValue) {
					continue
				}

				match := true

				for _, char := range posChar {
					if !hasItem(outputValue, char) {
						match = false
					}
				}

				if match {
					outputInts = append(outputInts, key)
				}
			}
		}

		// Merge the digits to a single number
		var str string
		for i := range outputInts {
			str += strconv.Itoa(outputInts[i])
		}
		num, _ := strconv.Atoi(str)

		answer += num
		fmt.Println(outputValues, num)
	}

	log.Printf("Answer: %d", answer)

}

func strToChars(strings []string) [][]string {
	var charsSlice [][]string
	for _, str := range strings {
		var chars []string
		for _, rune := range str {
			chars = append(chars, string(rune))
		}
		charsSlice = append(charsSlice, chars)
	}
	return charsSlice
}

func filterLength(strings [][]string, length int) [][]string {
	var filtered [][]string
	for _, str := range strings {
		if len(str) == length {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func removeItemAtEachPositions(positions [][]string, item string) [][]string {
	for i := range positions {
		positions[i] = removeItemByValue(positions[i], item)
	}
	return positions
}

func removeItemByValue(strings []string, item string) []string {
	var filtered []string
	for _, str := range strings {
		if str != item {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func hasItem(strings []string, item string) bool {
	for _, str := range strings {
		if str == item {
			return true
		}
	}
	return false
}
