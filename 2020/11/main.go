package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

//AreaKind kind of area
type AreaKind string

// Seat constants
const (
	kindFloor     AreaKind = "."
	kindSeatEmpty AreaKind = "L"
	kindSeatTaken AreaKind = "#"
	kindNone      AreaKind = " "
)

// Field struct
type Field struct {
	s    [][]AreaKind
	w, h int
}

// Life struct
type Life struct {
	a, b *Field
	w, h int
}

//NewField return a new field
func NewField(w, h int) *Field {
	s := make([][]AreaKind, h)
	for i := range s {
		s[i] = make([]AreaKind, w)
	}
	return &Field{s, w, h}
}

// Set (x, y) to field
func (f *Field) Set(x, y int, kind AreaKind) {
	f.s[y][x] = kind
}

func (f *Field) isOccupied(x, y int) bool {
	if x >= 0 && x < f.w && y >= 0 && y < f.h {
		return f.s[y][x] == kindSeatTaken
	}
	return false
}
func (f *Field) isFloor(x, y int) bool {
	if x >= 0 && x < f.w && y >= 0 && y < f.h {
		return f.s[y][x] == kindFloor
	}
	return true
}

func (f *Field) isEmpty(x, y int) bool {
	if x >= 0 && x < f.w && y >= 0 && y < f.h {
		return f.s[y][x] == kindSeatEmpty
	}
	return false
}

func (f *Field) get(x, y int) AreaKind {
	if x >= 0 && x < f.w && y >= 0 && y < f.h {
		return f.s[y][x]
	}
	return kindFloor
}

func (f *Field) isSeat(x, y int) bool {
	if x >= 0 && x < f.w && y >= 0 && y < f.h {
		place := f.s[y][x]
		switch place {
		case kindSeatEmpty, kindSeatTaken:
			return true
		}
	}
	return false
}

// Next return data for next round
func (f *Field) Next(x, y int) (bool, AreaKind) {
	if f.isFloor(x, y) {
		return false, kindFloor
	}
	occupied := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.isOccupied(x+i, y+j) {
				occupied++
			}
		}
	}
	if f.isEmpty(x, y) && occupied == 0 {
		return true, kindSeatTaken
	}
	if f.isOccupied(x, y) && occupied >= 4 {
		return true, kindSeatEmpty
	}
	return false, f.get(x, y)
}

func (f *Field) hasTop(x, y int) bool {
	for y >= 0 {
		y--
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}

func (f *Field) hasBottom(x, y int) bool {
	for y <= f.h {
		y++
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}

func (f *Field) hasRight(x, y int) bool {
	for x <= f.w {
		x++
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}
func (f *Field) hasLeft(x, y int) bool {
	for x >= 0 {
		x--
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}

func (f *Field) hasTopRight(x, y int) bool {
	for y >= 0 && x <= f.w {
		y--
		x++
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}
func (f *Field) hasBottomRight(x, y int) bool {
	for y <= f.h && x <= f.w {
		y++
		x++
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}

func (f *Field) hasBottomLeft(x, y int) bool {
	for y <= f.h && x >= 0 {
		y++
		x--
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}
func (f *Field) hasTopLeft(x, y int) bool {
	for y >= 0 && x >= 0 {
		y--
		x--
		if !f.isSeat(x, y) {
			continue
		}
		return f.isOccupied(x, y)
	}
	return false
}

// NextDirection returns data for next round using raycast to direction
func (f *Field) NextDirection(x, y int) (bool, AreaKind) {
	if f.isFloor(x, y) {
		return false, kindFloor
	}
	occupied := 0
	if f.hasTop(x, y) {
		occupied++
	}
	if f.hasBottom(x, y) {
		occupied++
	}
	if f.hasRight(x, y) {
		occupied++
	}
	if f.hasLeft(x, y) {
		occupied++
	}
	if f.hasTopRight(x, y) {
		occupied++
	}
	if f.hasBottomRight(x, y) {
		occupied++
	}
	if f.hasBottomLeft(x, y) {
		occupied++
	}
	if f.hasTopLeft(x, y) {
		occupied++
	}

	if f.isEmpty(x, y) && occupied == 0 {
		return true, kindSeatTaken
	}
	if f.isOccupied(x, y) && occupied >= 5 {
		return true, kindSeatEmpty
	}
	return false, f.get(x, y)
}

// NewLife return a new life
func NewLife(lines []string) *Life {
	a := NewField(len(lines[0]), len(lines))
	for y, line := range lines {
		seats := strings.Split(line, "")
		for x, seat := range seats {
			thisKind := kindFloor
			switch seat {
			case string(kindSeatEmpty):
				thisKind = kindSeatEmpty
			case string(kindSeatTaken):
				thisKind = kindSeatTaken
			}
			a.Set(x, y, thisKind)
		}
	}
	return &Life{
		a: a,
		b: NewField(len(lines[0]), len(lines)),
		w: len(lines[0]),
		h: len(lines),
	}
}

// Step advance life to next step
func (l *Life) Step() bool {
	stepHasChanged := false
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			hasChange, next := l.a.NextDirection(x, y)
			if hasChange {
				stepHasChanged = true
			}
			l.b.Set(x, y, next)
		}
	}
	l.a, l.b = l.b, l.a
	return stepHasChanged
}

// String format Field
func (f *Field) String() string {
	var buf bytes.Buffer
	for y := 0; y < f.h; y++ {
		for x := 0; x < f.w; x++ {
			b := byte(' ')
			switch f.get(x, y) {
			case kindSeatEmpty:
				b = 'L'
			case kindFloor:
				b = '.'
			case kindSeatTaken:
				b = '#'
			default:
				b = 'a'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

//String format Life
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			switch l.a.get(x, y) {
			case kindSeatEmpty:
				b = 'L'
			case kindFloor:
				b = '.'
			case kindSeatTaken:
				b = '#'
			case kindNone:
				b = ' '
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

func (l *Life) countOccupied() int {
	sum := 0
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			if l.a.isOccupied(x, y) {
				sum++
			}
		}
	}
	return sum
}

func main() {
	lines := utils.Load("11-real").ToStringSlice()
	l := NewLife(lines)
	hasChange := true
	for hasChange {
		hasChange = l.Step()
		// fmt.Print("\x0c", l) // Clear screen
		// bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	fmt.Println(l.countOccupied())
}
