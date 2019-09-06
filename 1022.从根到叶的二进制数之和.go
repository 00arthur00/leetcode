/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	sum, currValue := 0, 0
	helper(root, &sum, currValue)
	return sum
}
func helper(root *TreeNode, sum *int, preValue int) {
	//空,则直接返回
	if root == nil {
		return
	}
	//更新当前节点值
	currValue := preValue<<1 | root.Val

	if root.Left == nil && root.Right == nil {
		//left right为空,为子节点,计算和
		*sum += currValue
		return
	} else {
		//有一个节点不为空,则接着遍历
		helper(root.Left, sum, currValue)
		helper(root.Right, sum, currValue)
	}
}

