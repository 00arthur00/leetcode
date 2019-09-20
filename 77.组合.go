/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	bt(path, n, 1, k, &ans)
	return ans
}

func bt(path []int, n, start, limit int, output *[][]int) {
	if limit == 0 {
		copyOfResult := make([]int, len(path))
		copy(copyOfResult, path)
		*output = append(*output, copyOfResult)
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		bt(path, n, i+1, limit-1, output)
		path = path[0 : len(path)-1]
	}
}

