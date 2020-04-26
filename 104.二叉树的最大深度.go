/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
    if root==nil{
		return 0
	}
	
	left := 1+maxDepth(root.Left)
	right := 1+ maxDepth(root.Right)

	if left > right{
		return left
	}
	return right
}
// @lc code=end

