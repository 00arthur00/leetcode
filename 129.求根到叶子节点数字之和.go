/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根到叶子节点数字之和
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
func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
	
}
func helper(root *TreeNode,cur int)int{
	
	if root==nil{
		return 0
	}
	
	cur = cur*10+root.Val
	if root.Left==nil&&root.Right==nil{
		return cur
	}

	return helper(root.Left,cur)+helper(root.Right,cur)
}
// @lc code=end

