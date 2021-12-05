package main

import (
	"fmt"
	"log"
	"math"
)

type move struct {
	from     position
	to       position
	affected map[string]struct{}
}

func getGridSize(moves []move) (int, int) {
	maxX := 0
	maxY := 0

	for _, m := range moves {
		if m.from.x > maxX {
			maxX = m.from.x
		}
		if m.from.y > maxY {
			maxY = m.from.y
		}

		if m.to.x > maxX {
			maxX = m.to.x
		}
		if m.to.y > maxY {
			maxY = m.to.y
		}
	}

	return maxX, maxY
}

func (m move) getAffected() {
	m.affected[m.from.String()] = struct{}{}
	m.affected[m.to.String()] = struct{}{}
	if m.isHorizontal() { // all x are equal
		start := int(math.Min(float64(m.from.y), float64(m.to.y)))
		end := int(math.Max(float64(m.from.y), float64(m.to.y)))
		for y := start; y <= end; y++ {
			m.affected[position{m.from.x, y}.String()] = struct{}{}
		}
	}
	if m.isVertical() { // all y are equal
		start := int(math.Min(float64(m.from.x), float64(m.to.x)))
		end := int(math.Max(float64(m.from.x), float64(m.to.x)))
		for x := start; x <= end; x++ {
			m.affected[position{x, m.from.y}.String()] = struct{}{}
		}
	}

	if m.isDiagonal() {

		if m.goingTop() && m.goingLeft() {
			// top left
			// at each step, decrease x and y
			// until x,y == to.x, to.y
			nextDestination := position{m.from.x, m.from.y}
			lastDestination := position{m.to.x, m.to.y}
			for nextDestination != lastDestination {
				nextDestination.x--
				nextDestination.y--
				m.affected[nextDestination.String()] = struct{}{}
			}
		}
		if m.goingTop() && !m.goingLeft() {
			// top right
			// at each step, decrease x and increase y
			// until x,y == to.x, to.y
			nextDestination := position{m.from.x, m.from.y}
			lastDestination := position{m.to.x, m.to.y}
			for nextDestination != lastDestination {
				nextDestination.x--
				nextDestination.y++
				m.affected[nextDestination.String()] = struct{}{}
			}
		}
		if !m.goingTop() && m.goingLeft() {
			// bottom left
			// at each step, increase x and decrease y
			// until x,y == to.x, to.y
			nextDestination := position{m.from.x, m.from.y}
			lastDestination := position{m.to.x, m.to.y}
			for nextDestination != lastDestination {
				nextDestination.x++
				nextDestination.y--
				m.affected[nextDestination.String()] = struct{}{}
			}
		}

		if !m.goingTop() && !m.goingLeft() {
			// bottom right
			// at each step, increase x and increase y
			// until x,y == to.x, to.y
			nextDestination := position{m.from.x, m.from.y}
			lastDestination := position{m.to.x, m.to.y}
			for nextDestination != lastDestination {
				nextDestination.x++
				nextDestination.y++
				m.affected[nextDestination.String()] = struct{}{}
			}
		}

		log.Printf("Start %s, end %s", m.from.String(), m.to.String())
		log.Printf("pass by %v", m.affected)

	}
}

func (m move) String() string {
	affected := make([]string, 0)
	for k := range m.affected {
		affected = append(affected, fmt.Sprintf("[%s]", k))
	}
	return fmt.Sprintf("%s -> %s || %s", m.from.String(), m.to.String(), affected)
}

func (m move) isHorizontal() bool {
	return m.from.x == m.to.x
}

func (m move) isVertical() bool {
	return m.from.y == m.to.y
}

func (m move) isDiagonal() bool {
	return !m.isHorizontal() && !m.isVertical()
}

func (m move) goingTop() bool {
	return m.from.x > m.to.x
}

func (m move) goingLeft() bool {
	return m.from.y > m.to.y
}

func keepOnlyHorizontalAndVertical(moves []move) []move {
	retVal := make([]move, 0)

	for _, m := range moves {
		if m.isVertical() || m.isHorizontal() {
			retVal = append(retVal, m)
		}
	}

	return retVal
}
