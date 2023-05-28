package binarytree

import "fmt"

func binaryTreePaths(root *TreeNode) []string {
	var path string
	var pathResult []string

	var backtrack func(root *TreeNode, path string, pathResult []string) []string
	backtrack = func(root *TreeNode, path string, pathResult []string) []string {
		path += fmt.Sprint(root.Val)
		if root.Left == nil && root.Right == nil {
			pathResult = append(pathResult, path)
			return pathResult
		}

		if root.Left != nil {
			pathResult = backtrack(root.Left, path+"->", pathResult)
		}
		if root.Right != nil {
			pathResult = backtrack(root.Right, path+"->", pathResult)
		}

		return pathResult
	}

	return backtrack(root, path, pathResult)
}
