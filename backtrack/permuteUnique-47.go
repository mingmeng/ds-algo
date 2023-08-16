package backtrack

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0)
	result := make([][]int, 0)
	used := make([]bool, len(nums))

	var dfs func(nums []int, used []bool)
	dfs = func(nums []int, used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i] == false {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(nums, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(nums, used)

	return result
}
