/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */
/*
if n<0{
	min,max = max,min
}
maxF(n) = max(n,n*maxF(n-1))
minF(n) = min(n,n*minF(n-1))
*/
func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxProduct(nums []int) int {
	ans, min, max := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			max, min = min, max
		}
		max, min = maxInt(nums[i], max*nums[i]), minInt(nums[i], min*nums[i])
		ans = maxInt(max, ans)
	}
	return ans
}

