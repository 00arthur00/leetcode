/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	
	//计算节点的数值
	memo:=make(map[*TreeNode]int)
	helper(root,memo)
	
	//计算数值对应的出现次数
	valCount:=make(map[int]int)
	for _,val:=range memo{
		valCount[val]++
	}

	//计算最大出现次数
	maxCount:=-1
	for _,count:=range valCount{
		if count > maxCount{
			maxCount = count
		}
	}
	
	//返回结果
	ans:=make([]int,0)
	for val,count:=range valCount{
		if count==maxCount{
			ans = append(ans,val)
		}
	}
	return ans
}
func helper(root *TreeNode, memo map[*TreeNode]int) int{
	if root==nil{
		return 0
	}
	if val,ok:=memo[root];ok{
		return val
	}
	left := helper(root.Left, memo)
	right := helper(root.Right,memo)
	cur:=left+right+root.Val
	memo[root]=cur
	return cur
}
// @lc code=end

