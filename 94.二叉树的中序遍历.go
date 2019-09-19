/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	helper(root, &ans)
	return ans
}
func helper(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, ans)
	*ans = append(*ans, root.Val)
	helper(root.Right, ans)
}

