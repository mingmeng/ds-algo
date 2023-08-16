package backtrack

func permute(nums []int) [][]int {

	path := make([]int, 0)
	used := make([]bool, len(nums))
	results := make([][]int, 0)
	var permuteDFS func(num []int, used []bool)
	permuteDFS = func(num []int, used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			results = append(results, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			permuteDFS(nums, used)
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	permuteDFS(nums, used)

	return results
}
