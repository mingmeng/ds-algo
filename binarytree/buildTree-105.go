package binarytree

func buildTree105Helper(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}

	if len(preOrder) == 1 && len(inOrder) == 1 {
		return &TreeNode{
			Val: inOrder[0],
		}
	}

	rootValue := preOrder[0]
	root := &TreeNode{
		Val:   rootValue,
		Left:  nil,
		Right: nil,
	}
	rootPos := -1
	for i, v := range inOrder {
		if rootValue == v {
			rootPos = i
		}
	}

	leftInorder := inOrder[:rootPos]
	rightInorder := inOrder[rootPos+1:]

	preOrder = preOrder[1:]
	leftPreOrder := preOrder[:len(leftInorder)]
	rightPreOrder := preOrder[len(leftInorder):]

	root.Left = buildTree105Helper(leftPreOrder, leftInorder)
	root.Right = buildTree105Helper(rightPreOrder, rightInorder)

	return root
}
