/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */
func combinationSum3(k int, n int) [][]int {
	var sum int
	path := make([]int, 0)
	ans := make([][]int, 0)
	bt(k, n, 1, 9, &sum, path, &ans)
	return ans
}
func bt(k, n, start, end int, sum *int, path []int, output *[][]int) {
	if len(path) == k && *sum == n {
		copyOfResult := make([]int, k)
		copy(copyOfResult, path)
		*output = append(*output, copyOfResult)
		return
	}
	for i := start; i <= end; i++ {
		*sum += i
		path = append(path, i)
		bt(k, n, i+1, end, sum, path, output)
		*sum -= i
		path = path[0 : len(path)-1]
	}
}

