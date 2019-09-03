/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	result := 0
	for tmp > 0 {
		result = result*10 + tmp%10
		tmp = tmp / 10
	}
	return result == x
}

