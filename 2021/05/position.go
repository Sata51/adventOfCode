package main

import "fmt"

type position struct {
	x int
	y int
}

func (p position) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
