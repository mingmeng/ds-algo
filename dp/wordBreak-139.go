package dp

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, s2 := range wordDict {
		wordSet[s2] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			subStr := s[j:i]
			if wordSet[subStr] && dp[j] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}
