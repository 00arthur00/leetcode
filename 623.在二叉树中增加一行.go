/*
 * @lc app=leetcode.cn id=623 lang=golang
 *
 * [623] 在二叉树中增加一行
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
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	return helper(root,v,1,d)
}
func helper(root *TreeNode, v int,level, d int) *TreeNode {
	//root的情况
	if level > d-1{
		return &TreeNode{
			Val:v,
			Left:root,
		}
	}
	//root为空的叶子节点,直接返回
	if root==nil{
		return nil
	}
	//非root,直接添加层
	if level==d-1{
		head:=root
		root.Left = &TreeNode{
			Val:v,
			Left:head.Left,
		}
		root.Right = &TreeNode{
			Val:v,
			Right:head.Right,
		}
		return root
	}
	root.Left = helper(root.Left,v,level+1,d)
	root.Right = helper(root.Right,v,level+1,d)
	return root
}
// @lc code=end

