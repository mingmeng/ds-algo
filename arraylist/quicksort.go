package arraylist

func quicksort(nums []int, left, right int) []int {
	if right < left {
		return nums
	}
	// 选择最左边的元素
	pivot := nums[left]

	l := left
	r := right
	for l < r {
		for nums[r] >= pivot && r > l {
			r--
		}

		for nums[l] <= pivot && l < r {
			l++
		}

		nums[l], nums[r] = nums[r], nums[l]
	}

	nums[l], nums[left] = nums[left], nums[l]
	quicksort(nums, left, l-1)
	quicksort(nums, r+1, right)

	return nums
}
