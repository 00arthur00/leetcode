/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(num int) []int {
	dp:=make([]int,num+1)
	for i:=1;i<=num;i++{
		if i&0x01==0x01{
			dp[i] += dp[i/2]+1
			continue
		}
		dp[i] = dp[i/2]
	}
	return dp
}
// @lc code=end

