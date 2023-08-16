package arraylist

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left, right := 0, 0

	result := 0

	for right = 0; right < len(s); right++ {
		pos, ok := m[s[right]]
		if ok {
			left = max(pos, left)
		}

		if result < right-left+1 {
			result = right - left + 1
		}
		m[s[right]] = pos
	}
	return result
}
