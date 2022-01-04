package goleetcode

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	scenarios := []struct {
		input    int
		expected int
	}{
		{input: 4, expected: 2},
		{input: 8, expected: 2},
	}

	passed, failed := 0, 0
	for _, s := range scenarios {
		actual := mySqrt(s.input)
		if actual != s.expected {
			t.Errorf("expected value of %v but was %v", s.expected, actual)
			failed++
		} else {
			passed++
		}
	}
	t.Logf("%d passed; %d failed", passed, failed)
}

func mySqrt(x int) int {
	target := float64(x)
	approxRoot := target / 2

	for math.Abs(target-(approxRoot*approxRoot)) > 0.005 {
		approxRoot = (approxRoot + (target / approxRoot)) / 2.0
	}

	return int(approxRoot)
}
