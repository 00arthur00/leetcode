/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
 func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) > 0 {
		k--
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node.Val)
		if k == 0 {
			return node.Val
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			node = node.Right
			for node.Left != nil {
				stack = append(stack, node.Left)
				node = node.Left
			}
		}
	}
	return -1
}
// @lc code=end

