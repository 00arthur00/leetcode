/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	result:=make([]int,0)
	stack:=make([]*TreeNode,0)

	stack = append(stack,root)
	
	for len(stack)>0{
		node:=stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result,node.Val)
		if node.Right!=nil{
			stack = append(stack,node.Right)
		}
		if node.Left!=nil{
			stack = append(stack,node.Left)
		}
	}

	return result
}
func preorderTraversal1(root *TreeNode) []int {
	result:=make([]int,0)
	helper(root,&result)
	return result
}
func helper(root *TreeNode,result *[]int){
	if root==nil{
		return 
	}
	*result = append(*result,root.Val)
	helper(root.Left,result)
	helper(root.Right,result)
}
// @lc code=end

