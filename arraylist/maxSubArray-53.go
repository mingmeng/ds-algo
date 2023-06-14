package arraylist

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	result := math.MinInt64

	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]

		if count > result {
			result = count
		}
		if count < 0 {
			count = 0
		}
	}

	return result
}
