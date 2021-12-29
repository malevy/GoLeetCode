package goleetcode

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func TestAddBinary(t *testing.T) {
	scenarios := []struct {
		a        string
		b        string
		expected string
	}{
		{a: "0", b: "0", expected: "0"},
		{a: "1", b: "0", expected: "1"},
		{a: "0", b: "1", expected: "1"},
		{a: "1", b: "1", expected: "10"},
		{a: "11", b: "1", expected: "100"},
		{a: "1010", b: "1011", expected: "10101"},
	}

	passed, failed := 0, 0
	for _, s := range scenarios {
		actual := addBinary(s.a, s.b)
		if actual != s.expected {
			t.Errorf("expected value of %q but was %q", s.expected, actual)
			failed++
		} else {
			passed++
		}
	}
	t.Logf("%d passed; %d failed", passed, failed)
}

func addBinary(a string, b string) string {

	var maxLength int = int(math.Max(float64(len(a)), float64(len(b))))
	indexA, indexB := len(a)-1, len(b)-1

	sum := make([]string, maxLength)

	valueAt := func(s string, i int) int {
		v := 0
		if i >= 0 {
			v = int(s[i] - '0')
		}
		return v
	}

	carry := 0
	for indexA >= 0 || indexB >= 0 {
		result := valueAt(a, indexA) + valueAt(b, indexB) + carry
		resultIndex := int(math.Max(float64(indexA), float64(indexB)))
		sum[resultIndex] = strconv.Itoa(result % 2)
		carry = result / 2

		indexA--
		indexB--
	}
	if carry == 1 {
		sum = append([]string{"1"}, sum...)
	}

	return strings.Join(sum, "")
}
