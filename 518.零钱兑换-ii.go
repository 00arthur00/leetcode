/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
func change(amount int, coins []int) int {
	if amount<0{
		return 0
	}

	dp:=make([]int, amount+1)
	dp[0]=1
	for _,c:=range coins{
		for i:=c;i<=amount;i++{
			dp[i] += dp[i-c]
		}
	}
	return dp[amount]
}

// @lc code=end

