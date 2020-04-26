/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder)!=len(postorder) || len(inorder)==0{
		return nil
	}
	rootVal:=postorder[len(postorder)-1]
	root:=&TreeNode{Val:rootVal}
	pivot:=0
	for i,val:=range inorder{
		if val==rootVal{
			pivot=i
			break
		}
	}
	root.Left = buildTree(inorder[:pivot],postorder[:pivot])
	root.Right = buildTree(inorder[pivot+1:],postorder[pivot:len(postorder)-1])
	return root
}
// @lc code=end

