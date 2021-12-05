package main

import (
	"fmt"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
	"github.com/bradfitz/iter"
)

type (
	BingoBoard struct {
		size   int
		grid   []*gridEntry
		hasWin bool
	}

	gridEntry struct {
		x      int
		y      int
		value  int
		marked bool
	}
)

func (b *BingoBoard) String() string {
	var sb strings.Builder
	// Pour chaque ligne
	for i := range iter.N(b.size) {
		// Pour chaque colonne
		for j := range iter.N(b.size) {
			for _, c := range b.grid {
				if c.x == i && c.y == j {
					sb.WriteString(c.String())
					if c.marked {
						sb.WriteString("X")
					} else {
						sb.WriteString("O")
					}
					sb.WriteString(" ")
				}
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (c *gridEntry) String() string {
	return fmt.Sprintf("%02d", c.value)
}

func NewBingoBoard(input []string) *BingoBoard {

	b := &BingoBoard{
		size: len(input),
		grid: make([]*gridEntry, 0),
	}

	for i, line := range input {
		spl := r.FindAllString(line, -1)
		for j, v := range spl {
			if v == "" {
				continue
			}
			b.grid = append(b.grid, &gridEntry{
				x:      i,
				y:      j,
				value:  utils.MustParseInt(v),
				marked: false,
			})
		}
	}

	return b
}

func (b *BingoBoard) Mark(value int) {
	for _, c := range b.grid {
		if c.value == value {
			c.marked = true
		}
	}
}

func (b *BingoBoard) wins() bool {
	// Check vertically
	toCheck := make([]*gridEntry, 0)
	// For every column
	for i := range iter.N(b.size) {
		for _, c := range b.grid {
			if c.x == i {
				toCheck = append(toCheck, c)
			}
		}

		isMarked := true
		for _, v := range toCheck {
			if v.marked == false {
				isMarked = false
			}
		}
		if isMarked {
			return true
		}
		toCheck = make([]*gridEntry, 0) //reset toCheck
	}
	// Check horizontally
	for i := range iter.N(b.size) {
		for _, c := range b.grid {
			if c.y == i {
				toCheck = append(toCheck, c)
			}
		}

		isMarked := true
		for _, v := range toCheck {
			if v.marked == false {
				isMarked = false
			}
		}

		if isMarked {
			return true
		}
		toCheck = make([]*gridEntry, 0) //reset toCheck
	}

	return false
}
