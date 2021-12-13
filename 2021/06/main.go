package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Sata51/adventOfCode/pkg/utils"
	"github.com/bradfitz/iter"
)

type lanternfish struct {
	internalTimer int
}

func (lf *lanternfish) String() string {
	return fmt.Sprintf("%d", lf.internalTimer)
}

func main() {
	dt := utils.Load(2021, "06-real")
	fishInitialClock := strings.Split(dt.String(), ",")

	part1(fishInitialClock, 18)
	part1(fishInitialClock, 80)

	part2(fishInitialClock, 18)
	part2(fishInitialClock, 80)
	part2(fishInitialClock, 256)

}

func part1(fishInitialClock []string, days int) {

	fishes := make([]*lanternfish, 0)

	for _, fish := range fishInitialClock {
		fishes = append(fishes, &lanternfish{
			internalTimer: utils.MustParseInt(fish),
		})
	}

	// log.Printf("Initial state: %v", fishes)
	for range iter.N(days) {
		toSpawn := 0
		for _, fish := range fishes {
			if fish.internalTimer == 0 {
				fish.internalTimer = 7 // reset to 6
				toSpawn++
			}
			fish.internalTimer--
		}
		for range iter.N(toSpawn) {
			fishes = append(fishes, &lanternfish{
				internalTimer: 8,
			})
		}
		// log.Printf("After %03d day: %d", i+1, len(fishes))
	}

	log.Printf("After %d days: %d", days, len(fishes))
}

// Avoid counting a fish as it is, only count generation and for index 0 increment index 6 (old fish) and index 8 (new created fish)
// After all, sum all generations
func part2(fishInitialClock []string, days int) {
	fishes := make([]int, 9)

	for _, fish := range fishInitialClock {
		fishes[utils.MustParseInt(fish)]++
	}

	for range iter.N(days) {
		sum := 0
		for _, count := range fishes {
			sum += count
		}

		new_fishes := make([]int, 9)

		for index, count := range fishes {
			if index == 0 {
				new_fishes[6] += count
				new_fishes[8] += count
			} else {
				new_fishes[index-1] += count
			}
		}
		fishes = new_fishes
	}

	sum := 0
	for _, count := range fishes {
		sum += count
	}

	log.Printf("After %d days: %d", days, sum)
}
