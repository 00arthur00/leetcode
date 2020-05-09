/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	if N==0{
		return 0
	}
	if N==1{
		return 1
	}
	f_0:=0
	f_1:=1
	ans:= 0
	for i:=2;i<=N;i++{
		ans = f_0+f_1
		f_0=f_1
		f_1=ans
	}
	return ans
}
// @lc code=end

