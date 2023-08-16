package backtrack

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	path := make([]int, 0)
	results := make([][]int, 0)

	var backtrack func(nums []int, index int)
	backtrack = func(nums []int, index int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		results = append(results, tmp)

		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}

			path = append(path, nums[i])
			backtrack(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(nums, 0)

	return results
}
