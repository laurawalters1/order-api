package calculatepacks

import (
	"fmt"
	"sort"
)

func CalculatePacks(count int) map[int]int {
	order := map[int]int{
		250:  1,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	// CREATE ORDERED SLICE OF PACKSIZES
	packSizes := make([]int, 0, len(order))
	for size := range order {
		packSizes = append(packSizes, size)
	}
	sort.Ints(packSizes)
	fmt.Println(packSizes)

	packSizesNeeded := make([][]int, 0, len(order))

	for _, size := range packSizes {
		// Start at first el (250)
		if !(size > count) {
			for _, innerSize := range packSizes {
				if innerSize+size > count {
					packSizesNeeded = append(packSizesNeeded, make([]int, size, innerSize))
				}
			}
		} else if size > count {
			packSizesNeeded = append(packSizesNeeded, make([]int, size))
			break
		}

		// If greater, return

		// If not greater, start nested loop

		// Start at first el

		// If outer loop el + inner loop el > return
	}

	return order
}
