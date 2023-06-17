package backtrack

var combinationSum3Result [][]int
var combinationSum3Path []int

func dfs(n, k, start, sum int) {
	if len(combinationSum3Path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, combinationSum3Path)
			combinationSum3Result = append(combinationSum3Result, tmp)
		}
		return
	}

	for i := start; i <= 9; i++ {
		combinationSum3Path = append(combinationSum3Path, i)
		sum += i
		dfs(n, k, i, sum)
		sum -= i
		combinationSum3Path = combinationSum3Path[:len(combinationSum3Path)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	combinationSum3Result = make([][]int, 0)
	combinationSum3Path = make([]int, 0)
	dfs(n, k, 1, 0)
	return combinationSum3Result
}
