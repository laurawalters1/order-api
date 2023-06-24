package calculatepacks

import (
	"fmt"
	"math"

	"github.com/laurawalters1/order-api/utils"
)

func recursive(count int, packSizes []int, accumulatedSizes []int) []int {
	var sum int = utils.GetSum(accumulatedSizes)
	if sum >= count {
		return accumulatedSizes
	} else {

		for i, size := range packSizes {
			var prevNumsSum = utils.GetSum(packSizes[0 : i+1]) // returns sum of previous and current number in the packsizes array

			// BREAKING CONDITION
			if sum+size >= count {
				accumulatedSizes = append(accumulatedSizes, size)
				return accumulatedSizes
			} else if prevNumsSum+sum >= count { // if previous & curr number + sum >= count current number can be added as we know we won't need a larger number
				accumulatedSizes = append(accumulatedSizes, size)
				return recursive(count, packSizes, accumulatedSizes)
			} else if i == len(packSizes)-1 { // if current value is the largest packsize and count has not been met, we must add it to the array and loop back round
				accumulatedSizes = append(accumulatedSizes, size)
				return recursive(count, packSizes, accumulatedSizes)
			} else if (utils.GetSum(packSizes[i+1:i+2]))-size >= count {
				var numNeeded = int(math.Ceil(float64(count-prevNumsSum) / float64(size)))
				for y := 0; y < numNeeded; y++ {
					accumulatedSizes = append(accumulatedSizes, size)
				}
				return accumulatedSizes
			}
		}

		// Could check if the target value is at least *currentVal* less than accumulatedSizes + *nextVal*, if it is, add x *currentVal* to meet the target
		// EXAMPLE:
		// targetVal = 249
		// currVal = 50
		// nextVal = 500
		// if ((nextVal - currVal) + sumVal ) >= targetVal {
		// x = (targetVal - sumVal) / currVal -- Then add x number of currVals to the accumulatedSizes array and return
		// }
		return accumulatedSizes
	}

}

func getOptimalSlice(arr [][]int) []int {
	var returnVal []int = arr[0]
	var optimalArrSum int

	for i, array := range arr {
		sum := utils.GetSum(array)
		if i == 0 {
			optimalArrSum = sum
			returnVal = array
		} else if sum < optimalArrSum {
			optimalArrSum = sum
			returnVal = array
		}
	}
	return returnVal
}

func CalculatePacks(count int, order []int) map[int]int {

	fmt.Println(order)

	var orderMap = make(map[int]int)
	for _, size := range order {
		orderMap[size] = 0
	}
	fmt.Println(orderMap)

	// order := map[int]int{
	// 	250:  0,
	// 	500:  0,
	// 	1000: 0,
	// 	2000: 0,
	// 	5000: 0,
	// }

	// PUT PACK SIZES IN ORDER
	packSizes := utils.OrderPackSizes(orderMap)

	reversePackSizes := utils.Reverse(packSizes)

	// This variable will store arrays which contain all the number combinations that fulfill the order
	var allCalculations = [][]int{}

	for _, size := range reversePackSizes {
		acc := []int{size}
		var calculation []int = recursive(count, packSizes, acc)
		allCalculations = append(allCalculations, calculation)

	}
	optimalSlice := getOptimalSlice(allCalculations)

	for _, size := range optimalSlice {
		orderMap[size]++
	}
	return orderMap

}
