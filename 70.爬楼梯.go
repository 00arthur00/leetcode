/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	n_1 := 2
	n_2 := 1
	result := n_1
	for i := 3; i <= n; i++ {
		result = n_2 + n_1
		n_2 = n_1
		n_1 = result
	}
	return result
}

