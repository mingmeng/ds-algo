package arraylist

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if target == nums[middle] {
			return middle
		}

		if nums[left] <= nums[middle] {
			if target >= nums[left] && target < nums[middle] {
				right = middle + 1
			} else {
				left = middle - 1
			}
		} else {
			if target <= nums[right] && target > nums[middle] {
				left = middle - 1
			} else {
				right = middle + 1
			}
		}
	}
	return -1
}
