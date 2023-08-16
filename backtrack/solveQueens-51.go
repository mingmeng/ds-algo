package backtrack

import "strings"

func solveNQueens(n int) [][]string {
	var solveNQueensResults [][]string

	square := make([]string, n)
	for i := range square {
		square[i] = strings.Repeat(".", n)
	}

	var backtrack func(square []string, row, n int)
	backtrack = func(square []string, row, n int) {
		if len(square) == row {
			tmp := make([]string, len(square))
			copy(tmp, square)
			solveNQueensResults = append(solveNQueensResults, tmp)
			return
		}

		for col := 0; col < n; col++ {
			if isValid(square, n, row, col) {
				r := []byte(square[row])
				r[col] = 'Q'
				square[row] = string(r)
				backtrack(square, row+1, n)
				r[col] = '.'
				square[row] = string(r)
			}
		}
	}
	backtrack(square, 0, n)
	return solveNQueensResults
}

func isValid(square []string, n, row, col int) bool {
	for i := 0; i < n; i++ {
		if i != row && square[i][col] == 'Q' {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if square[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}

	for i, j := row-1, col+1; i >= 0 && j < n; {
		if square[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	return true
}
