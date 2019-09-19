/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */
func missingNumber(nums []int) int {
	res := len(nums)
	for i, n := range nums {
		res = res ^ i ^ n
	}
	return res
}

