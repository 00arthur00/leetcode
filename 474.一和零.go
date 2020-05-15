/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, str := range strs {
		zero, one := countZeroOne(str)
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}
func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}
func countZeroOne(str string)(zero,one int){
	for _,c:=range str{
		if c=='0'{
			zero++
		}
		if c=='1'{
			one++
		}
	}
	return
}
// @lc code=end

