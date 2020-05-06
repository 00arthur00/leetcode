/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
		if root==nil{
			return nil
		}
		result:=make([]int,0)
		stack:=make([]*TreeNode,0)
	
		stack = append(stack,root)
		
		for len(stack)>0{
			node:=stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append([]int{node.Val},result...)
			if node.Left!=nil{
				stack = append(stack,node.Left)
			}
			if node.Right!=nil{
				stack = append(stack,node.Right)
			}
		}
	
		return result
}
// @lc code=end

