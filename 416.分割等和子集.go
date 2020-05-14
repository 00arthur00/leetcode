/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {

	sum:=0
	for _,n:=range nums{
		sum += n
	}

	if sum%2!=0{
		return false
	}

	target := sum/2
	return hasSubSet(nums,target)
}

func hasSubSet(nums[]int, target int) bool{

	ans:=make([]bool,target+1)
	ans[0]=true
	for _,n:=range nums{
		for i:=target;i>=n;i--{//从后往前计算防止,被覆盖
			ans[i] =  ans[i]||ans[i-n]
		}
	}
	return ans[target]
}
// @lc code=end

