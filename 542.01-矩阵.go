import "math"

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */
func updateMatrix(matrix [][]int) [][]int {
	locations := make([]*location, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				locations = append(locations, &location{i, j})
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				min := math.MaxInt32
				for _, l := range locations {
					min = minInt(min, abs(l.x-i)+abs(l.y-j))
				}
				matrix[i][j] = min
			}
		}
	}
	return matrix
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type location struct {
	x, y int
}
