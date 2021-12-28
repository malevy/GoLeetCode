package goleetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	scenarios := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1}, expected: []int{2}},
		{input: []int{2, 1}, expected: []int{2, 2}},
		{input: []int{1, 9}, expected: []int{2, 0}},
		{input: []int{9}, expected: []int{1, 0}},
	}

	passed, failed := 0, 0
	for _, s := range scenarios {
		actual := plusOne(s.input)
		if !CollectionEquals(actual, s.expected) {
			t.Errorf("expected value of %v but was %v", s.expected, actual)
			failed++
		} else {
			passed++
		}
	}
	t.Logf("%d passed; %d failed", passed, failed)
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] = digits[i] % 10
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
