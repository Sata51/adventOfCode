package main

import (
	"fmt"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
)

type field struct {
	name      string
	lowRange  rrange
	highRange rrange
}

type rrange struct {
	start int
	stop  int
}

type ticket struct {
	value []int
}

func main() {
	data := utils.Load("real")
	fields, myTicket, nearbyTickets := parseInput(data.String())

	fmt.Println("step1", step1(nearbyTickets, fields))
	fmt.Println("step2", step2(myTicket, nearbyTickets, fields))
}

func step2(myTicket ticket, nearbyTickets []ticket, fields []field) int {
	validTickets := make([]ticket, 0)
	validTickets = append(validTickets, myTicket)
	for _, t := range nearbyTickets {
		if _, ok := t.isValid(fields); ok {
			validTickets = append(validTickets, t)
		}
	}

	columnValidForField := make(map[string]map[int]struct{})
	// For every fields
	for _, f := range fields {
		// For every column
		columnValidForField[f.name] = make(map[int]struct{})
	columnloop:
		for i := 0; i < len(myTicket.value); i++ {
			allEntry := takeAllEntry(i, validTickets)
			// For every entry in column
			for _, e := range allEntry {
				if f.lowRange.isInRange(e) || f.highRange.isInRange(e) {
					// do nothing
				} else {
					continue columnloop
				}
			}
			columnValidForField[f.name][i] = struct{}{}
		}
	}
	columnForField := make(map[string]int)
validLoop:
	for len(columnValidForField) > 0 {
		for k, v := range columnValidForField {
			if len(v) > 1 {
				continue validLoop
			}
			// There is only one but since its an map we must iterate
			for key := range v {
				columnForField[k] = key
			}
			delete(columnValidForField, k)
			// remove columnForField[k] from all other fields
			for _, v1 := range columnValidForField {
				delete(v1, columnForField[k])
			}
		}
	}
	retVal := 1
	for k, v := range columnForField {
		if strings.HasPrefix(k, "departure") {
			retVal *= myTicket.value[v]
		}
	}

	return retVal
}

func takeAllEntry(index int, validTickets []ticket) []int {
	retVal := make([]int, len(validTickets))
	for i, entry := range validTickets {
		retVal[i] = entry.value[index]
	}
	return retVal
}

func step1(nearbyTickets []ticket, fields []field) int {
	errorRate := 0
	for _, t := range nearbyTickets {
		rate, _ := t.isValid(fields)
		errorRate += rate
	}
	return errorRate
}

func (t ticket) isValid(f []field) (int, bool) {
entryloop:
	for _, entry := range t.value {
		for _, verifier := range f {
			if verifier.lowRange.isInRange(entry) || verifier.highRange.isInRange(entry) {
				continue entryloop
			}
		}
		return entry, false
	}
	return 0, true
}

func (r rrange) isInRange(value int) bool {
	if r.start <= value && value <= r.stop {
		return true
	}
	return false
}

func parseInput(s string) (fields []field, myTicket ticket, nearbyTickets []ticket) {
	fields = make([]field, 0)
	nearbyTickets = make([]ticket, 0)

	data := strings.Split(s, "\n\n")

	for _, line := range strings.Split(data[0], "\n") {
		fields = append(fields, parseField(line))
	}
	myTicket = parseTicket(strings.Split(data[1], "\n")[1])

	for _, line := range strings.Split(data[2], "\n")[1:] {
		nearbyTickets = append(nearbyTickets, parseTicket(line))
	}

	return
}

func parseTicket(s string) ticket {
	thisTicket := ticket{value: make([]int, 0)}
	for _, v := range strings.Split(s, ",") {
		thisTicket.value = append(thisTicket.value, utils.MustParseInt(v))
	}
	return thisTicket
}

func parseField(s string) field {
	thisField := field{}
	splitted := strings.Split(s, ":")
	thisField.name = strings.TrimSpace(splitted[0])
	ranges := strings.ReplaceAll(strings.TrimSpace(splitted[1]), " or ", " ")
	rangeSplitted := strings.Split(ranges, " ")
	thisField.lowRange = parseRange(rangeSplitted[0])
	thisField.highRange = parseRange(rangeSplitted[1])
	return thisField
}

func parseRange(s string) rrange {
	splitted := strings.Split(s, "-")
	return rrange{
		start: utils.MustParseInt(splitted[0]),
		stop:  utils.MustParseInt(splitted[1]),
	}
}
