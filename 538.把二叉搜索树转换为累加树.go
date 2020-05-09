/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
	var post *TreeNode
	revertinorder(root, &post)
	return root
}
func revertinorder(root *TreeNode, post **TreeNode){
	if root==nil{
		return
	}
	revertinorder(root.Right,post)
	if *post != nil{
		root.Val += (*post).Val
	}
	*post = root
	revertinorder(root.Left, post)
}
// @lc code=end

