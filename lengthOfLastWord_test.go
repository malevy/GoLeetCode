package goleetcode

import (
	"strings"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	scenarios := []struct {
		input    string
		expected int
	}{
		{input: "a", expected: 1},
		{input: "hello world", expected: 5},
		{input: "hello world   ", expected: 5},
		{input: "  hello world   ", expected: 5},
	}

	for _, s := range scenarios {
		actual := lengthOfLastWord(s.input)
		if actual != s.expected {
			t.Errorf("expected value of %v but was %v", s.expected, actual)
		}
	}
}

func lengthOfLastWord(s string) int {
	trimmed := strings.TrimRight(s, " ")
	index := strings.LastIndex(trimmed, " ")
	if index == -1 {
		return len(trimmed)
	} else {
		return len(trimmed) - index - 1
	}
}
