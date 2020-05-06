/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode)  {
	if root==nil{
		return
	}

	leftMax:=root
	if leftMax.Left!=nil{
		leftMax = leftMax.Left
		for leftMax.Right!=nil{
			leftMax = leftMax.Right
		}
		leftMax.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	
	flatten(root.Right)	
}
// @lc code=end

