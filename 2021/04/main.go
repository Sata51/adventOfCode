package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

var r *regexp.Regexp = regexp.MustCompile("[^\\s]+")

func main() {
	dt := utils.Load("04-real")

	input := dt.ToStringSlice()

	part1(input)
	part2(input)

}

func part1(input []string) {
	log.Printf("Part1")

	selectedNumbers := getSelectedNumbers(input[0])

	// Remove the two first line
	input = input[2:]

	boards := make([]*BingoBoard, 0)

	boardLine := make([]string, 0)
	for _, line := range input {
		if len(line) == 0 {
			boards = append(boards, NewBingoBoard(boardLine))
			boardLine = make([]string, 0) // Reset the current board line
			continue
		}
		boardLine = append(boardLine, line)
	}
	boards = append(boards, NewBingoBoard(boardLine))

	numberWhoMakeAWinner, winner := bingoNumbersFirst(selectedNumbers, boards)
	if winner == nil {
		log.Panic("No winner")
	}

	sumOfUnmarked := 0
	for _, c := range winner.grid {
		if c.marked == false {
			sumOfUnmarked += c.value
		}
	}

	log.Printf("\n%s\nNumberWhoMakeAWinner: %d\nSumOfUnmarked: %d", winner, numberWhoMakeAWinner, sumOfUnmarked)

	log.Printf("Result : %d", numberWhoMakeAWinner*sumOfUnmarked)

}

func part2(input []string) {
	log.Printf("Part2")

	selectedNumbers := getSelectedNumbers(input[0])

	// Remove the two first line
	input = input[2:]

	boards := make([]*BingoBoard, 0)

	boardLine := make([]string, 0)
	for _, line := range input {
		if len(line) == 0 {
			boards = append(boards, NewBingoBoard(boardLine))
			boardLine = make([]string, 0) // Reset the current board line
			continue
		}
		boardLine = append(boardLine, line)
	}
	boards = append(boards, NewBingoBoard(boardLine))

	numberWhoMakeAWinner, winner := bingoNumbersLast(selectedNumbers, boards)
	if winner == nil {
		log.Panic("No winner")
	}

	sumOfUnmarked := 0
	for _, c := range winner.grid {
		if c.marked == false {
			sumOfUnmarked += c.value
		}
	}

	log.Printf("\n%s\nNumberWhoMakeAWinner: %d\nSumOfUnmarked: %d", winner, numberWhoMakeAWinner, sumOfUnmarked)

	log.Printf("Result : %d", numberWhoMakeAWinner*sumOfUnmarked)

}

func bingoNumbersFirst(input []int, boards []*BingoBoard) (int, *BingoBoard) {
	for _, n := range input {
		for _, b := range boards {
			b.Mark(n)
			if b.wins() {
				return n, b
			}
		}
	}

	return -1, nil
}

func bingoNumbersLast(input []int, boards []*BingoBoard) (int, *BingoBoard) {

	var lastWins *BingoBoard = nil
	var lastNumberWhoMakeAWinner int = -1

inputLoop:
	for _, n := range input {
		for _, b := range boards {
			b.Mark(n)
			if !b.hasWin && b.wins() {
				b.hasWin = true
				lastWins = b
				lastNumberWhoMakeAWinner = n
				if hasAllWinned(boards) {
					break inputLoop
				}
			}
		}
	}

	return lastNumberWhoMakeAWinner, lastWins
}

func hasAllWinned(boards []*BingoBoard) bool {
	for _, b := range boards {
		if !b.hasWin {
			return false
		}
	}
	return true
}

func getSelectedNumbers(input string) []int {
	nb := strings.Split(input, ",")
	selectedNumbers := make([]int, 0)

	for _, n := range nb {
		selectedNumbers = append(selectedNumbers, utils.MustParseInt((n)))
	}

	return selectedNumbers
}
