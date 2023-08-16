package stack

import "container/list"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 先打出nums2 对应的下一个最大元素表
	map2 := make(map[int]int)
	stack := list.New()
	stack.PushBack(0)
	for i := 0; i < len(nums2); i++ {
		if nums2[i] > stack.Back().Value.(int) {
			for stack.Len() > 0 && nums2[i] > stack.Back().Value.(int) {
				map2[stack.Back().Value.(int)] = nums2[i]
				stack.Remove(stack.Back())
			}
		}

		stack.PushBack(nums2[i])
	}

	result := make([]int, len(nums1))

	for i, n := range nums1 {
		if next, ok := map2[n]; ok {
			result[i] = next
		} else {
			result[i] = -1
		}
	}

	return result
}
