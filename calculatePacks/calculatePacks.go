package calculatepacks

import (
	"fmt"
	"sort"
)

func recursive(innerSize int, packSizes []int, sizes []int, count int, packSizesNeeded [][]int) {

	var sumOfSizes int
	for _, size := range sizes {
		sumOfSizes += size
	}
	fmt.Println("sumOfSizes", sumOfSizes)
	if innerSize+sumOfSizes >= count {
		var newSizes = []int{}
		sizes = append(sizes, innerSize)
		newSizes = append(newSizes, sizes...)
		fmt.Println("newSizes", newSizes)
		packSizesNeeded = append(packSizesNeeded, newSizes)
		fmt.Println("packSizesNeeded", packSizesNeeded)

	} else {
		// packSizesNeeded = append(packSizesNeeded, make([]int, sizes, innerSize))
		var newSizes = []int{}
		sizes = append(sizes, innerSize)
		newSizes = append(newSizes, sizes...)
		recursive(innerSize, packSizes, newSizes, count, packSizesNeeded)
	}

}

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

	packSizesNeeded := make([][]int, 0, len(order))
	// fmt.Println("Pack sizes needed", packSizesNeeded)

	for _, size := range packSizes {
		// Start at first el (250)
		if size < count {
			sizes := []int{size}
			// fmt.Println("sizes", sizes)
			fmt.Println("packSizes", packSizes)
			fmt.Println("sizes", sizes)
			fmt.Println("count", count)
			fmt.Println("packSizesNeeded", packSizesNeeded)
			for _, innerSize := range packSizes {
				recursive(innerSize, packSizes, sizes, count, packSizesNeeded)
			}
		} else if size > count {
			appendItem := []int{size}
			packSizesNeeded = append(packSizesNeeded, appendItem)
			break
		}

		// If greater, return

		// If not greater, start nested loop

		// Start at first el

		// If outer loop el + inner loop el > return
	}

	fmt.Println(packSizesNeeded)
	return order
}
