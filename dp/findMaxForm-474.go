package dp

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		var one, zero int
		for _, ch := range strs[i] {
			switch ch {
			case '1':
				one++
			case '0':
				zero++
			}
		}

		for j := m; j >= 0; i-- {
			for k := n; k >= 0; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zero][k-one]+1)
			}
		}
	}

	return dp[m][n]
}
