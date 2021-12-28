package goleetcode

import "testing"

func TestMaxSubarrySum(t *testing.T) {
	scenarios := []struct {
		input    []int
		expected int
	}{
		{input: []int{}, expected: 0},
		{input: []int{5}, expected: 5},
		{input: []int{-1, 2}, expected: 2},
		{input: []int{1, 2}, expected: 3},
		{input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
		{input: []int{5, 4, -1, 7, 8}, expected: 23},
	}

	for _, s := range scenarios {
		actual := maxSubarraySum(s.input)
		if actual != s.expected {
			t.Errorf("expected value of %v but was %v", s.expected, actual)
		}
	}
}

func maxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	localMax := nums[0]
	for _, v := range nums[1:] {
		if v > localMax+v {
			localMax = v
		} else {
			localMax += v
		}

		if localMax > max {
			max = localMax
		}
	}

	return max
}
