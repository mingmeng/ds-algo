package arraylist

import "math"

func minSubArrayLen(target int, nums []int) int {
	var result = math.MaxInt32

	sum := 0
	for fast, slow := 0, 0; fast < len(nums); fast++ {
		sum += nums[fast]
		for sum >= target {
			result = min(result, fast-slow+1)
			sum -= nums[slow]
			slow++
		}
	}
	if result == math.MaxInt32 {
		return 0
	} else {
		return result
	}
}
