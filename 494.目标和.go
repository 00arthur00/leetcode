/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */
/*
sum(P)-sum(N) = S
sum(P)+sum(N) = sum(nums)

2*sum(P) = S+sum(nums)

sum(P)为偶数
求和为(S+sum(nums))/2的子数组是否不能在
*/
// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	sum:=sumArray(nums)
	if sum<S||(sum+S)%2!=0{
		return 0
	}
	target:=(sum+S)/2
	dp:=make([]int,target+1)
	dp[0]=1
	for _,n:=range nums{
		for i:=target;i>=n;i--{
			dp[i]+=dp[i-n]
		}
	}
	return dp[target]
}

func sumArray(nums []int) int{
	ans:=0
	for _,n:=range nums{
		ans+=n
	}
	return ans
}
// @lc code=end

