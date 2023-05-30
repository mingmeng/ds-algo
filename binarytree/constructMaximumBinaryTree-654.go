package binarytree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	maxInSlice := nums[0]
	maxIndex := 0
	for index, num := range nums {
		if num > maxInSlice {
			maxInSlice = num
			maxIndex = index
		}
	}
	root := &TreeNode{Val: maxInSlice}

	leftSlice := nums[:maxIndex]
	rightSlice := nums[maxIndex+1:]

	if len(leftSlice) > 0 {
		root.Left = constructMaximumBinaryTree(leftSlice)
	}
	if len(rightSlice) > 0 {
		root.Right = constructMaximumBinaryTree(rightSlice)
	}

	return root
}
