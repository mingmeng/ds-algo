package dp

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bp01([]int) {
	// the backpack contains 4
	weight := []int{1, 3, 4}
	value := []int{15, 25, 35}
	bagweight := 4

	dp := make([][]int, 3)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, bagweight)
	}

	for j := weight[0]; j < bagweight; j++ {
		dp[0][j] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 1; j < bagweight; j++ {
			if j < weight[j] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return
}

func bp01floor1() {
	// the backpack contains 4
	weight := []int{1, 3, 4}
	value := []int{15, 25, 35}
	bagWeight := 4

	dp := make([]int, bagWeight+1)
	for i := weight[0]; i < len(dp); i++ {
		dp[i] = value[0]
	}

	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= 0; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
}
