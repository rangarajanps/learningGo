package dog

import (
	"testing"
)

func TestYears(t *testing.T) {

	testData := map[uint]uint{
		3:   21,
		7:   49,
		25:  175,
		123: 861,
	}

	for input, expected := range testData {
		actualResult := Years(input)
		if actualResult != expected {
			t.Errorf("TestYears of (%d) was incorrect, got: %d, want: %d.", input, actualResult, expected)
		}
	}
}
