package stocks

func maxProfitFee(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0] - fee
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(prices)-1][1]
}
