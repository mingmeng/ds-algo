package arraylist

import "sort"

func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	diff := 1
	for i := 0; i < len(nums)-1; i++ {

		if nums[i] != nums[i+1] {
			diff++
			if diff == 3 {
				return nums[i+1]
			}
		}
	}

	return nums[0]
}
