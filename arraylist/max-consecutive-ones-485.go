package arraylist

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	length := len(nums)

	ones := 0
	for i := length - 1; i >= 0; i-- {
		if nums[i] == 0 {
			ones = 0
		} else {
			ones++
		}

		result = max(result, ones)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
