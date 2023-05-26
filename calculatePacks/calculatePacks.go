package calculatepacks

import (
	"fmt"
	"math"
	"sort"
)

func recursive(innerSize int, packSizes []int, sizes []int, count int, packSizesNeeded [][]int) {

}

func orderPackSizes(order map[int]int) []int {
	// CREATE ORDERED SLICE OF PACKSIZES
	packSizes := make([]int, 0, len(order))
	for size := range order {
		packSizes = append(packSizes, size)
	}
	sort.Ints(packSizes)
	return packSizes
}

func getNumberRequiredForEachSize(packSizes []int, count int) map[int]float64 {
	var requiredNums = map[int]float64{}
	for _, size := range packSizes {

		requiredNums[size] = math.Ceil(float64(count) / float64(size))
		// fmt.Printf("%T", size)
	}
	return requiredNums
}

func allValuesEqualOne(numsMap map[int]float64) bool {
	returnVal := true
	for _, num := range numsMap {
		if num != 1 {
			returnVal = false
		}
		// fmt.Printf("%T", size)
	}
	return returnVal
}

func CalculatePacks(count int) map[int]int {
	order := map[int]int{
		250:  0,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	// PUT PACK SIZES IN ORDER
	packSizes := orderPackSizes(order)

	// GET HOW MANY OF EACH SIZE WOULD BE REQUIRED IF YOU COULD ONLY TAKE ONE SIZE
	var requiredNumsForEachSize = getNumberRequiredForEachSize(packSizes, count)
	fmt.Println(requiredNumsForEachSize)
	if allValuesEqualOne(requiredNumsForEachSize) {
		order[packSizes[0]]++
		return order
	} else {
		return order
	}

	// fmt.Println(requiredNumsForEachSize)

}
