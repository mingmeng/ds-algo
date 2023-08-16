package arraylist

func searchRange(nums []int, target int) []int {

	l := searchLeft(nums, target)
	r := searchRight(nums, target)

	if l == -1 && r == -1 {
		return []int{l, r}
	} else if l == -1 {
		return []int{r, r}
	} else if r == -1 {
		return []int{l, l}
	}
	return []int{l, r}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if left <= 0 || left > len(nums)-1 {
		return -1
	}

	if nums[left] == target {
		return left
	}
	return -1
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if right <= 0 || right > len(nums)-1 || nums[right-1] != target {
		return -1
	}

	return right - 1
}
