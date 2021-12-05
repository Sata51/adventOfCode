package main

import (
	"fmt"

	"github.com/Sata51/adventOfCode/pkg/utils"
	"github.com/bradfitz/iter"
)

func main() {
	hight := -1
	seats := make(map[int]position)
	for _, sr := range utils.Load("05-real").ToStringSlice() {
		p := splitName(sr)
		p.getRow()
		p.getSeat()
		p.getSeatID()
		seats[p.seatID] = p
		if p.seatID > hight {
			hight = p.seatID
		}
	}

	for i := range iter.N(hight) {
		if _, ok := seats[i]; !ok {
			fmt.Printf("Missing %d\n", i)
		}
	}

	fmt.Printf("%d\n", hight)
}

type position struct {
	rowInput    string
	columnInput string
	row         int
	column      int
	seatID      int
}

func (p position) String() string {
	return fmt.Sprintf("%s %s %d %d : %d", p.rowInput, p.columnInput, p.row, p.column, p.seatID)
}

func splitName(s string) position {
	return position{
		rowInput:    s[0:7],
		columnInput: s[7:],
		row:         -1,
		column:      -1,
		seatID:      -1,
	}
}

func (p *position) getRow() {
	rows := make([]int, 128)
	for i := range iter.N(128) {
		rows[i] = i
	}
	for _, r := range p.rowInput {
		switch r {
		case 'F':
			rows = rows[:len(rows)/2]
		case 'B':
			rows = rows[len(rows)/2:]
		}
	}
	p.row = rows[0]
}

func (p *position) getSeat() {
	rows := make([]int, 8)
	for i := range iter.N(8) {
		rows[i] = i
	}
	for _, r := range p.columnInput {
		switch r {
		case 'L':
			rows = rows[:len(rows)/2]
		case 'R':
			rows = rows[len(rows)/2:]
		}
	}
	p.column = rows[0]
}

func (p *position) getSeatID() {
	p.seatID = p.row*8 + p.column
}
