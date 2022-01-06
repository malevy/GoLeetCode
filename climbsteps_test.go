package goleetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	scenarios := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 3},
	}

	passed, failed := 0, 0
	for _, s := range scenarios {
		actual := climbStairs(s.input)
		if actual != s.expected {
			t.Errorf("expected value of %v but was %v", s.expected, actual)
			failed++
		} else {
			passed++
		}
	}
	t.Logf("%d passed; %d failed", passed, failed)
}

func climbStairs(n int) int {

	cache := map[int]int{
		1: 1,
		2: 2,
	}

	return countPaths(n, &cache)
}

func countPaths(n int, cache *map[int]int) int {

	v, found := (*cache)[n]

	if found {
		return v
	}

	v1, found := (*cache)[n-1]
	if !found {
		v1 = countPaths(n-1, cache)
		(*cache)[n-1] = v1
	}

	v2, found := (*cache)[n-2]
	if !found {
		v2 = countPaths(n-2, cache)
		(*cache)[n-2] = v2
	}

	return v1 + v2

}
