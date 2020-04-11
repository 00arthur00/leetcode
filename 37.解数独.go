/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	bt(board, 0, 0)
}
func bt(board [][]byte, x, y int) bool {
	if x == 9 {
		return isValid(board, x, y)
	}
	if board[x][y] == '.' {
		for i := 1; i <= 9; i++ {
			board[x][y] = '0' + i
			if isValid(board, x, y) {
				if bt(board, x+1, y+1) {
					return true
				}
			}
		}
	} else {
		return bt(board, x+1, y+1)
	}
	return false
}
func isValid(board [][]byte, x, y index) bool {
	return true
}

// @lc code=end

