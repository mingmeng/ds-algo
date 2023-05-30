package binarytree

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	midIndex := len(nums) / 2
	root := &TreeNode{Val: nums[midIndex]}

	leftSlice := nums[:midIndex]
	rightSlice := nums[midIndex+1:]

	if len(leftSlice) > 0 {
		root.Left = sortedArrayToBST(leftSlice)
	}
	if len(rightSlice) > 0 {
		root.Right = sortedArrayToBST(rightSlice)
	}

	return root
}
