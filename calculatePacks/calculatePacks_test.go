package calculatepacks

import "testing"

func TestCalculatePacks1(t *testing.T) {
	expected := OrderBreakdown{
		_250:  1,
		_500:  0,
		_1000: 0,
		_2000: 0,
		_5000: 0,
	}
	result := calculatepacks(1)
	if result != expected {
		t.Errorf("\033[31mFAILED - Expected %d GOT %d\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %d GOT %d\033[0m", expected, result)
	}
}

func TestCalculatePacks250(t *testing.T) {
	expected := OrderBreakdown{
		_250:  1,
		_500:  0,
		_1000: 0,
		_2000: 0,
		_5000: 0,
	}
	result := calculatepacks(250)
	if result != expected {
		t.Errorf("\033[31mFAILED - Expected %d GOT %d\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %d GOT %d\033[0m", expected, result)
	}
}

func TestCalculatePacks251(t *testing.T) {
	expected := OrderBreakdown{
		_250:  0,
		_500:  1,
		_1000: 0,
		_2000: 0,
		_5000: 0,
	}
	result := calculatepacks(1)
	if result != expected {
		t.Errorf("\033[31mFAILED - Expected %d GOT %d\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %d GOT %d\033[0m", expected, result)
	}
}

func TestCalculatePacks501(t *testing.T) {
	expected := OrderBreakdown{
		_250:  1,
		_500:  1,
		_1000: 0,
		_2000: 0,
		_5000: 0,
	}
	result := calculatepacks(1)
	if result != expected {
		t.Errorf("\033[31mFAILED - Expected %d GOT %d\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %d GOT %d\033[0m", expected, result)
	}
}

func TestCalculatePacks12001(t *testing.T) {
	expected := OrderBreakdown{
		_250:  1,
		_500:  0,
		_1000: 0,
		_2000: 1,
		_5000: 2,
	}
	result := calculatepacks(1)
	if result != expected {
		t.Errorf("\033[31mFAILED - Expected %d GOT %d\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %d GOT %d\033[0m", expected, result)
	}
}
