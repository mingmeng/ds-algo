package binarytree

func buildTreeInAndPost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	// find the root node
	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rootValue,
	}
	if len(postorder) == 1 {
		return root
	}

	// split left and right side in in order sequence
	var rootIndex int
	for index, node := range inorder {
		if node == rootValue {
			rootIndex = index
			break // in fact no need, as the description said no repeat value
		}
	}

	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	// cut down the whole root
	postorder = postorder[:len(postorder)-1]
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder):]

	root.Left = buildTreeInAndPost(leftInorder, leftPostorder)
	root.Right = buildTreeInAndPost(rightInorder, rightPostorder)

	return root
}
