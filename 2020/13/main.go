package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

func main() {
	data := utils.Load("real").ToStringSlice()

	step1(data)
	step2(data)
}

func step1(data []string) {
	start, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	bus := parseBus(data[1])
	fmt.Println(bus)
	needToWait, busID := getCloser(start, bus)
	fmt.Println(needToWait, busID, needToWait*busID)
}

func step2(data []string) {
	bus := parseBus2(data[1])
	fmt.Println(bus)
	closerTimestampContiguous := getCloserContiguous(bus)
	fmt.Println(closerTimestampContiguous)
}

func getCloserContiguous(buses map[int]int) int {
	t := buses[0]
	delta := buses[0]

	for dt, bus := range buses {
		for {
			if (t+dt)%bus == 0 {
				break
			}
			t += delta
		}
		delta = lcm(delta, bus)
	}
	return t
}

func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	return a * b / gcd(a, b)
}

func getCloser(start int, bus []int) (needToWait int, busID int) {
	closer := make([]float64, len(bus))
	for i := range closer {
		closer[i] = math.MaxFloat64
	}
busloop:
	for i, b := range bus {
		fmt.Printf("for %d bus start iteration\n", b)
		data := 0
		for data <= start+b {
			newDiff := float64(data) - float64(start)
			if newDiff > 0 {
				if newDiff < closer[i] {
					closer[i] = newDiff
				} else {
					fmt.Println("continue busLoop")
					continue busloop
				}
			}
			data += b
		}
	}
	smallIndex := -1
	smallValue := math.MaxFloat64
	for i, c := range closer {
		if c < smallValue {
			smallValue = c
			smallIndex = i
		}
	}

	needToWait = int(smallValue)
	busID = bus[smallIndex]
	return
}

func parseBus(line string) []int {
	bus := make([]int, 0)
	for _, s := range strings.Split(line, ",") {
		if s == "x" {
			continue
		}
		busLine, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		bus = append(bus, busLine)
	}

	return bus
}

func parseBus2(line string) map[int]int {
	retVal := make(map[int]int)
	for i, s := range strings.Split(line, ",") {
		if s == "x" {
			continue
		}
		busID, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		retVal[i] = busID
	}
	return retVal
}
