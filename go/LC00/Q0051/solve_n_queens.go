package Q0051

import "bytes"

var res = make([][]string, 0)
var rowBuf bytes.Buffer

func solveNQueens(n int) [][]string {
	res = res[:0]
	board := make([][]bool, n)
	rowBuf.Grow(n)

	for i := range board {
		board[i] = make([]bool, n)
	}

	backTrack(board, 0, n)
	return res
}

func backTrack(board [][]bool, row, n int) {
	if row == n {
		b := make([]string, n)
		for i, r := range board {
			for _, c := range r {
				if c {
					rowBuf.WriteRune('Q')
				} else {
					rowBuf.WriteRune('.')
				}
			}
			b[i] = rowBuf.String()
			rowBuf.Reset()
		}
		res = append(res, b)
		return
	}
	for i := range board[row] {
		if !isValid(board, row, i) {
			continue
		}
		board[row][i] = true
		backTrack(board, row+1, n)
		board[row][i] = false
	}

}

func isValid(board [][]bool, row int, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < len(board); {
		if board[i][j] {
			return false
		}
		i--
		j++
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] {
			return false
		}
		i--
		j--
	}

	return true
}
