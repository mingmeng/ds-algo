package arraylist

func findKthLargest(nums []int, k int) int {
	split(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func split(nums []int, k, left, right int) {
	if left > right {
		return
	}
	if index := partition(nums, left, right); index == k {
		return
	} else if index < k {
		partition(nums, index+1, right)
	} else {
		partition(nums, left, index-1)
	}

}

func partition(nums []int, left, right int) int {
	if right < left {
		return -1
	}

	pivot := nums[left]

	l := left + 1
	r := right

	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		for l < r && nums[l] < pivot {
			l++
		}

		nums[l], nums[r] = nums[r], nums[l]
	}

	nums[l] = pivot

	return l
}
