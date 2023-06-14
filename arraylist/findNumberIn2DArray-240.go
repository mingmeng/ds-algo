package arraylist

// findNumberIn2DArray
func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	for i, j := n-1, 0; i > 0 && j < m; {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
