package dp

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < len(dp); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(dp[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = grid[i][j] + max(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}
