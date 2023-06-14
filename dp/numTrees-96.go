package dp

func numTrees(n int) int {
	if n == 1 {
		return n
	} else if n == 2 {
		return n
	}

	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		for j := 0; j <= i; j++ {
			dp[i] += dp[j] * dp[i-j]
		}
	}
	return dp[n]
}
