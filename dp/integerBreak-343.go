package dp

func integerBreak(n int) int {
	if n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[2] = 1

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}
