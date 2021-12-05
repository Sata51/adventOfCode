package main

import (
	"log"

	"github.com/Sata51/adventOfCode/pkg/utils"
	"github.com/bradfitz/iter"
)

func main() {
	dt := utils.Load("03-real")
	strInputs := dt.ToStringSlice()

	part1(strInputs)
	part2(strInputs)

}

func part1(strInput []string) {
	gammaRate := utils.BinaryStringToInt(getGammaRate(strInput))
	epsilonRate := utils.BinaryStringToInt(utils.InvertBinaryString(getGammaRate((strInput))))

	log.Printf("Gamma rate: %v", gammaRate)
	log.Printf("Epsilon rate: %v", epsilonRate)

	log.Printf("Power consumption: %v", gammaRate*epsilonRate)
}

func part2(strInput []string) {
	oxygenGeneratorRating := utils.BinaryStringToInt(getOxygenGeneratorRating(strInput))
	co2ScrubberRating := utils.BinaryStringToInt(getCo2ScrubberRating(strInput))

	log.Printf("Oxygen generator rating: %v", oxygenGeneratorRating)
	log.Printf("CO2 scrubber rating: %v", co2ScrubberRating)

	log.Printf("Result is %d", oxygenGeneratorRating*co2ScrubberRating)
}

func getOxygenGeneratorRating(strInput []string) string {
	filtered := strInput

	// Pour chaque position
	for i := 0; i < len(strInput[0]); i++ {
		// On prend le most common
		mostCommon := mostCommonAtPosition(i, filtered)
		filtered = filter(filtered, i, mostCommon)
		if len(filtered) == 1 {
			break
		}
	}

	return filtered[0]
}

func getCo2ScrubberRating(strInput []string) string {
	filtered := strInput

	// Pour chaque position
	for i := 0; i < len(strInput[0]); i++ {
		// On prend le most common
		leastCommon := leastCommonAtPosition(i, filtered)
		filtered = filter(filtered, i, leastCommon)
		if len(filtered) == 1 {
			break
		}
	}

	return filtered[0]
}

func filter(inputs []string, position int, value string) []string {
	filtered := make([]string, 0)
	for _, str := range inputs {
		if str[position] == value[0] {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func getGammaRate(strInput []string) string {
	// Split all inputs into slices
	binaryOutput := ""

	for i := range iter.N(len(strInput[0])) {
		binaryOutput += mostCommonAtPosition(i, strInput)
	}

	return binaryOutput
}

func mostCommonAtPosition(position int, inputs []string) string {
	cnt := make(map[rune]int)
	// Only work for binary
	cnt['0'] = 0
	cnt['1'] = 0
	for _, str := range inputs {
		c := str[position]
		cnt[rune(c)]++
	}

	if cnt['0'] > cnt['1'] {
		return "0"
	}
	if cnt['0'] < cnt['1'] {
		return "1"
	}
	return "1"
}

func leastCommonAtPosition(position int, inputs []string) string {
	cnt := make(map[rune]int)
	// Only work for binary
	cnt['0'] = 0
	cnt['1'] = 0
	for _, str := range inputs {
		c := str[position]
		cnt[rune(c)]++
	}

	if cnt['0'] < cnt['1'] {
		return "0"
	}
	if cnt['0'] > cnt['1'] {
		return "1"
	}
	return "0"
}
