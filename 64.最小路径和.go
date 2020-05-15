/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}
func minPathSum(grid [][]int) int {
	if len(grid)==0||len(grid[0])==0{
		return 0
	}
	
	dp:=make([]int,len(grid[0]))
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[i]);j++{
			if j==0{
				dp[j] = dp[j]
			}else if i==0{
				dp[j] = dp[j-1]
			}else{
				dp[j] = min(dp[j],dp[j-1])
			}
			dp[j]+=grid[i][j]
		}
	}
	return dp[len(grid[0])-1]
}
// @lc code=end

