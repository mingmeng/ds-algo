package dp

func minCostClimbingStairs(cost []int) int {
	// 题目里面说了 cost >= 2 暂时不需要考虑

	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
