/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */
//f[n]=max(f[n-1]*1,(n-1)*1,...,f[1]*n-1,n-1)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	ans := make([]int, n+1, n+1)
	ans[1], ans[2] = 1, 1
	for i := 3; i <= n; i++ {
		ans[i] = 1
		for j := 1; j < i; j++ {
			ans[i] = max(max(ans[i-j]*j, (i-j)*j), ans[i])
		}
	}
	return ans[n]
}

