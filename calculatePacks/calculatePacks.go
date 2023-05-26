package calculatepacks

func calculatepacks(count int32) map[int32]int {
	order := map[int32]int{
		250:  1,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	return order
}
