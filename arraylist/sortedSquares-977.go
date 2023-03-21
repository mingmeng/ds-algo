package arraylist

func sortedSquares(nums []int) []int {
	var starts, ends = 0, len(nums) - 1
	result := make([]int, len(nums))
	pos := len(nums) - 1
	for starts <= ends {
		if nums[starts]*nums[starts] < nums[ends]*nums[ends] {
			result[pos] = nums[ends] * nums[ends]
			ends--
		} else {
			result[pos] = nums[starts] * nums[starts]
			starts++
		}
		pos--
	}
	return result
}
