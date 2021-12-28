package goleetcode

func CollectionEquals(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	counts := make(map[int]int)

	for _, v := range x {
		counts[v] += 1
	}

	missCount := 0
	for _, v := range y {
		cnt, ok := counts[v]
		if ok && cnt > 0 {
			counts[v] -= 1
		} else {
			missCount++
		}
	}

	return missCount == 0
}
