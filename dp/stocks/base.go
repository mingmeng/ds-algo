package stocks

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestSumAfterKNegations(nums []int, k int) int {

	for i := 0; i < k; i++ {
		min := nums[0]
		for index := range nums {
			if nums[index] < min {
				nums[index] = -nums[index]
			}
		}
	}

	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
