package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type orientation string
type side string
type order string

type instruction struct {
	order interface{}
	value int
}
type boat struct {
	orientation   orientation
	east          int
	north         int
	wayPointEast  int
	wayPointNorth int
}

//Orientation constant
const (
	OrientationNorth orientation = "N"
	OrientationWest  orientation = "W"
	OrientationSouth orientation = "S"
	OrientationEast  orientation = "E"
	SideLeft         side        = "L"
	SideRight        side        = "R"
	OrderForward     order       = "F"
)

var (
	axis = []orientation{OrientationNorth, OrientationEast, OrientationSouth, OrientationWest}
)

func newBoat() *boat {
	return &boat{
		orientation:   OrientationEast,
		east:          0,
		north:         0,
		wayPointEast:  10,
		wayPointNorth: 1,
	}
}

func (b *boat) String() string {
	return fmt.Sprintf("o:%s, e:%d, n:%d we:%d wn:%d", b.orientation, b.east, b.north, b.wayPointEast, b.wayPointNorth)
}

func main() {
	instructions := parse(utils.Load("12-real").ToStringSlice())
	fmt.Println(instructions)

	b := newBoat()

	for _, i := range instructions {
		// i.process(b)
		i.processWayPoint(b)
		fmt.Println(i, b)
	}

	fmt.Println(b.getManhattanDistance())
}

func (b *boat) moveNorth(value int) {
	b.north += value
}
func (b *boat) moveWayPointNorth(value int) {
	b.wayPointNorth += value
}
func (b *boat) moveSouth(value int) {
	b.north -= value
}
func (b *boat) moveWayPointSouth(value int) {
	b.wayPointNorth -= value
}
func (b *boat) moveWest(value int) {
	b.east -= value
}
func (b *boat) moveWayPointWest(value int) {
	b.wayPointEast -= value
}
func (b *boat) moveEast(value int) {
	b.east += value
}
func (b *boat) moveWayPointEast(value int) {
	b.wayPointEast += value
}
func (b *boat) moveForward(value int) {
	switch b.orientation {
	case OrientationNorth:
		b.moveNorth(value)
	case OrientationSouth:
		b.moveSouth(value)
	case OrientationWest:
		b.moveWest(value)
	case OrientationEast:
		b.moveEast(value)
	}
}
func (b *boat) moveWayPointForward(value int) {
	b.north += b.wayPointNorth * value
	b.east += b.wayPointEast * value
}

func getOrientationIndex(or orientation) int {
	for i, o := range axis {
		if o == or {
			return i
		}
	}
	return -1
}

func (b *boat) turnLeft(value int) {
	currentOrientation := b.orientation
	currentOrientationIndex := getOrientationIndex(currentOrientation)
	currentOrientationIndex += len(axis)
	currentOrientationIndex -= value / 90
	currentOrientationIndex %= len(axis)
	b.orientation = axis[currentOrientationIndex]
}

//turnWayPointAround: for right pass negative value
func (b *boat) turnWayPointAround(value int) {
	radians := (math.Pi / 180) * float64(value)

	distance := math.Sqrt(
		math.Pow(float64(b.wayPointEast), 2) +
			math.Pow(float64(b.wayPointNorth), 2),
	)

	angle := math.Atan2(
		float64(b.wayPointNorth), float64(b.wayPointEast),
	) + radians

	b.wayPointEast = int(math.Round(distance * math.Cos(angle)))
	b.wayPointNorth = int(math.Round(distance * math.Sin(angle)))
}

func (b *boat) turnRight(value int) {
	currentOrientation := b.orientation
	currentOrientationIndex := getOrientationIndex(currentOrientation)
	currentOrientationIndex += len(axis)
	currentOrientationIndex += value / 90
	currentOrientationIndex %= len(axis)
	b.orientation = axis[currentOrientationIndex]
}

func (b *boat) getManhattanDistance() int {
	return int(math.Abs(float64(b.north)) + math.Abs(float64(b.east)))
}

func (i instruction) process(b *boat) {
	switch i.order {
	case OrientationNorth:
		b.moveNorth(i.value)
	case OrientationSouth:
		b.moveSouth(i.value)
	case OrientationWest:
		b.moveWest(i.value)
	case OrientationEast:
		b.moveEast(i.value)
	case SideLeft:
		b.turnLeft(i.value)
	case SideRight:
		b.turnRight(i.value)
	case OrderForward:
		b.moveForward(i.value)
	}
}

func (i instruction) processWayPoint(b *boat) {
	switch i.order {
	case OrientationNorth:
		b.moveWayPointNorth(i.value)
	case OrientationSouth:
		b.moveWayPointSouth(i.value)
	case OrientationWest:
		b.moveWayPointWest(i.value)
	case OrientationEast:
		b.moveWayPointEast(i.value)
	case SideLeft:
		b.turnWayPointAround(i.value)
	case SideRight:
		b.turnWayPointAround(-i.value)
	case OrderForward:
		b.moveWayPointForward(i.value)
	}
}

func (i instruction) String() string {
	return fmt.Sprintf("[%v %d]", i.order, i.value)
}

func parse(lines []string) []instruction {
	retVal := make([]instruction, 0)

	for _, line := range lines {
		thisInstruction := instruction{}
		switch line[0] {
		case 'N':
			thisInstruction.order = OrientationNorth
		case 'W':
			thisInstruction.order = OrientationWest
		case 'S':
			thisInstruction.order = OrientationSouth
		case 'E':
			thisInstruction.order = OrientationEast
		case 'L':
			thisInstruction.order = SideLeft
		case 'R':
			thisInstruction.order = SideRight
		case 'F':
			thisInstruction.order = OrderForward
		}
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		thisInstruction.value = val
		retVal = append(retVal, thisInstruction)
	}

	return retVal
}
