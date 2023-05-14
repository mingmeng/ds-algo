package binarytree

var arr = make([][]int, 0)

func order(r *TreeNode, depth int) {
	if r == nil {
		return
	}

	if depth == len(arr) {
		arr = append(arr, []int{})
	}

	arr[depth] = append(arr[depth], r.Val)

	order(r.Left, depth+1)
	order(r.Right, depth+1)
}

func levelOrder(root *TreeNode) [][]int {
	depth := 0

	order(root, depth)

	return arr
}
