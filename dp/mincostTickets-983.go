package dp

// 在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为 days 的数组给出。每一项是一个从 1 到 365 的整数。
//
// 火车票有 三种不同的销售方式 ：
//
// 一张 为期一天 的通行证售价为 costs[0] 美元；
// 一张 为期七天 的通行证售价为 costs[1] 美元；
// 一张 为期三十天 的通行证售价为 costs[2] 美元。
// 通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张 为期 7 天 的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。
//
// 返回 你想要完成在给定的列表 days 中列出的每一天的旅行所需要的最低消费 。
func mincostTickets(days []int, costs []int) int {
	maxDay := days[len(days)-1]
	minDay := days[0]

	dp := make([]int, maxDay+31)

	for d, i := maxDay, len(days)-1; d >= minDay; d-- {
		if d == days[i] {
			dp[d] = min(dp[d+30]+costs[2], dp[d+7]+costs[1])
			dp[d] = min(dp[d], dp[d+1]+costs[0])
			i--
		} else {
			dp[d] = dp[d+1]
		}
	}

	return dp[minDay]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
