package stocks

// 2 1 4
func maxProfit5(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[0][0] = -prices[0]
			dp[0][1] = 0
		} else if i == 1 {
			dp[1][0] = max(dp[0][0], -prices[1])
			dp[1][1] = max(dp[0][1], dp[0][0]+prices[1])
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		}
	}

	return dp[len(prices)-1][1]
}
