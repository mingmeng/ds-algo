package backtrack

func subsets(nums []int) [][]int {
	subSetsPath := make([]int, 0)
	subSetsResults := make([][]int, 0)

	var backtrack func(nums []int, index int)
	backtrack = func(nums []int, index int) {
		tmp := make([]int, len(subSetsPath))
		copy(tmp, subSetsPath)
		subSetsResults = append(subSetsResults, tmp)
		if index >= len(nums) {
			return
		}

		for i := index; i < len(nums); i++ {
			subSetsPath = append(subSetsPath, nums[i])
			backtrack(nums, i+1)
			subSetsPath = subSetsPath[:len(nums)-1]
		}
	}

	backtrack(nums, 0)

	return subSetsResults
}
