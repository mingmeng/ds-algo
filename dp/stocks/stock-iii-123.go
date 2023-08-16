package stocks

func maxProfit3(prices []int) int {
	maxK := 2

	dp := make([][][]int, len(prices))
	for i := range dp {
		dp[i] = make([][]int, maxK)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, maxK)
		}
	}

	for i := 0; i < len(prices); i++ {
		for k := maxK; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0] = -prices[0]
				dp[i][k][1] = 0
			} else {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k-1][1]-prices[i])
				dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k][0]+prices[i])
			}
		}
	}

	return dp[len(dp)-1][maxK][1]
}
