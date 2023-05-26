package calculatepacks

import (
	"reflect"
	"testing"
)

func TestCalculatePacks1(t *testing.T) {
	expected := map[int32]int{
		250:  1,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := calculatepacks(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %%v: %v GOT %%v: %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks250(t *testing.T) {
	expected := map[int32]int{
		250:  1,
		500:  0,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := calculatepacks(250)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks251(t *testing.T) {
	expected := map[int32]int{
		250:  0,
		500:  1,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := calculatepacks(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks501(t *testing.T) {
	expected := map[int32]int{
		250:  1,
		500:  1,
		1000: 0,
		2000: 0,
		5000: 0,
	}

	result := calculatepacks(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}

func TestCalculatePacks12001(t *testing.T) {
	expected := map[int32]int{
		250:  1,
		500:  0,
		1000: 0,
		2000: 1,
		5000: 2,
	}

	result := calculatepacks(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("\033[31mFAILED - Expected %v GOT %v\033[0m", expected, result)
	} else {
		t.Logf("\033[32mPASSED - Expected %v GOT %v\033[0m", expected, result)
	}
}
