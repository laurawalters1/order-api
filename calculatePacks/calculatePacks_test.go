package calculatepacks

import "testing"

func TestCalculatePacks(t *testing.T) {
	var expected int32 = 1
	result := calculatepacks(1)
	if result != expected {
		t.Errorf("FAILED - Expected %d GOT %d", expected, result)
	} else {
		t.Logf("PASSED - Expected %d GOT %d", expected, result)
	}
}
