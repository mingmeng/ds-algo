package binarytree

import (
	"container/list"
)

func preorderTraversalIterate(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)

	stack := list.New()
	stack.PushBack(root.Val)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)

		result = append(result, node.Val)

		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}

	return result
}

func postOrderTraversalIterate(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)

	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)

		result = append(result, node.Val)

		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	// 懒得去写reverse函数了，直接写过程里面了，实际上最好抽出来方便理解
	for i := 0; i < len(result)/2; i++ {
		temp := result[i]
		result[i] = result[len(result)-i-1]
		result[len(result)-i-1] = temp
	}

	return result
}

func inOrderTraversalIterate(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)

	stack := list.New()

	cur := root
	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			cur = stack.Remove(stack.Back()).(*TreeNode)
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}

	return result
}
