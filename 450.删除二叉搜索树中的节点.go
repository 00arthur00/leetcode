/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func successor(root *TreeNode) int{
	root = root.Right
	for root!=nil&&root.Left!=nil{
		root = root.Left
	}
	return root.Val
}

func presuccessor(root *TreeNode) int{
	root = root.Left
	for root!=nil&&root.Right!=nil{
		root = root.Right
	}
	return root.Val
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root==nil{
		return nil
	}

	if key>root.Val{
		root.Right = deleteNode(root.Right,key)
	}else if key<root.Val{
		root.Left =deleteNode(root.Left,key)
	}else{
		if root.Left==nil&&root.Right==nil{
			return nil
		}else if root.Right!=nil{
			root.Val = successor(root)
			root.Right=deleteNode(root.Right,root.Val)
		}else if root.Left!=nil{
			root.Val = presuccessor(root)
			root.Left = deleteNode(root.Left,root.Val)
		}
	}
	return root
}
// @lc code=end

