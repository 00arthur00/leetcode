/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, sum int) int {
	if root==nil{
		return 0
	}
	return countSum(root,sum)+pathSum(root.Left,sum)+pathSum(root.Right,sum)
}
func countSum(root *TreeNode, sum int) int{
	if root==nil{
		return 0
	}
	count:=0
	if root.Val==sum{
		count++
	}
	return count + 
		countSum(root.Left, sum - root.Val) +
		countSum(root.Right, sum - root.Val)
}
// @lc code=end

