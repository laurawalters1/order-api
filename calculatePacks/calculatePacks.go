package calculatepacks

import (
	"fmt"
)

func CalculatePacks(count int32) map[int32]int {
	order := map[int32]int{
		250:  1,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	for key, element := range order {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}
	return order
}
