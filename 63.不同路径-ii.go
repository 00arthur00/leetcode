/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m:=len(obstacleGrid)
	n:=len(obstacleGrid[0])
	dp:=make([]int,len(obstacleGrid[0])+1)
	dp[1]=1
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if obstacleGrid[i][j]==1{
				dp[j+1]=0
				continue
			}
			dp[j+1]+=dp[j]
		}
	}
	return dp[n]
}
// @lc code=end

