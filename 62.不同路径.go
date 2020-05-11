/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp:=make([]int,n+1)
	dp[1] = 1
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			dp[j] += dp[j-1]
		}
	}
	return dp[n]
}
// @lc code=end

