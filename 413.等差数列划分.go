/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
/*
d[i]=d[i-1]+1 if A[i]-A[i-1]==A[i-1]-A[i-2]
*/

func numberOfArithmeticSlices(A []int) int {
	dp:=make([]int,len(A))
	sum:=0
	for i:=2;i<len(A);i++{
		if A[i]-A[i-1]==A[i-1]-A[i-2]{
			dp[i] = dp[i-1]+1
			sum += dp[i]
		}
	}
	
	return sum
}
// @lc code=end

