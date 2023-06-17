package binarytree

import "container/list"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	levelFloor := make([][]int, 0)

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		length := queue.Len()

		floor := make([]int, length)
		for i := 0; i < length; i++ {
			// 把队头pop出来
			node := queue.Remove(queue.Front()).(*TreeNode)
			floor[i] = node.Val
			// 对每个pop出来的元素，把不为空的左右节点push进去
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		levelFloor = append(levelFloor, floor)
	}

	result := make([]int, len(levelFloor))
	for i, e := range levelFloor {
		arrLen := len(e)
		result[i] = e[arrLen-1]
	}

	return result
}
