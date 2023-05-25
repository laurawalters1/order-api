package calculatepacks

type OrderBreakdown struct {
	_250  int32
	_500  int32
	_1000 int32
	_2000 int32
	_5000 int32
}

func calculatepacks(count int32) OrderBreakdown {
	return OrderBreakdown{
		_250:  1,
		_500:  0,
		_1000: 0,
		_2000: 0,
		_5000: 0,
	}
}
