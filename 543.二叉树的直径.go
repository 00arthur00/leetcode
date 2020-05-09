/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
//后序遍历
func diameterOfBinaryTree(root *TreeNode) int {
	if root==nil{
		return 0
	}
	max:=-1
	maxLeafDepth(root,&max)
	return max
}

func maxLeafDepth(root *TreeNode, max *int) int{
	if root==nil{
		return 0
	}
	cur := 0
	curSum:=0
	if root.Left!=nil{
		left:= 1+ maxLeafDepth(root.Left, max)
		cur = left
		curSum =  cur
	}
	if root.Right!=nil{
		right:=maxLeafDepth(root.Right,max)
		if cur < 1+right{
			cur = 1+right
		}
		curSum += 1+right
	}
	if curSum > *max{
		 *max = curSum
	}
	return cur
}
// @lc code=end

