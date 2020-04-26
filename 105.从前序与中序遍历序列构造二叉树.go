/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder)!= len(inorder){
		return nil
	}
	if len(preorder)==0{
		return nil
	}
	root:=&TreeNode{Val:preorder[0]}
	pivot:=0
	for i:=0;i<len(inorder);i++{
		if inorder[i]==preorder[0]{
			pivot = i
			break
		}
	}
	root.Left = buildTree(preorder[1:pivot+1],inorder[:pivot])
	root.Right = buildTree(preorder[pivot+1:],inorder[pivot+1:])
	return root
}
// @lc code=end

