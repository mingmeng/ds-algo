package string

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	var slow = 0

	maxLength := 0
	for fast := 1; fast < len(s); fast++ {
		for i := fast - 1; i >= slow; i-- {
			if s[i] == s[fast] {
				slow = i + 1
				break
			}
		}

		maxLength = max(maxLength, fast-slow+1)
	}

	return maxLength
}
