package arraylist

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[i]+nums[right] == 0 {
				res := []int{nums[left], nums[i], nums[right]}
				result = append(result, res)
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				right--
				left++
			} else if nums[left]+nums[i]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}

	}

	return result
}
