package dp

// babad
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	dp := make([][]bool, 0)
	maxLen := 0
	startPos := 0

	for i := 0; i < len(s); i++ {
		left, right := i, i
		for left >= 0 && right <= len(s)-1 {
			if left == right {
				dp[left][right] = true
			} else {
				if s[left] == s[right] && dp[left-1][right-1] {
					dp[left][right] = dp[left-1][right-1]
					maxLen = max(maxLen, right-left)
					startPos = left
				} else {
					dp[left][right] = false
					break
				}
			}
		}
	}

	return s[startPos : startPos+maxLen]
}
