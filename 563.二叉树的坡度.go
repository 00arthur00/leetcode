/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
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
func findTilt(root *TreeNode) int {
	titl:=0
	postorder(root,&titl)
	return titl
}
//返回当前节点和
func postorder(root *TreeNode, titl *int) int{
	if root==nil{
		return 0
	}
	left:=postorder(root.Left,titl)
	right:=postorder(root.Right,titl)
	diff:=left-right
	if diff < 0{
		diff = -diff
	}
	*titl += diff
	return left+right+root.Val
}
// @lc code=end

