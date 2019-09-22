/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	q := []*TreeNode{root}
	helper(q, &ans)
	return ans
}
func helper(q []*TreeNode, output *[][]int) {
	nq := make([]*TreeNode, 0)
	ld := make([]int, 0)
	for _, node := range q {
		if node != nil {
			ld = append(ld, node.Val)
			nq = append(nq, node.Left, node.Right)
		}
	}

	if len(ld) > 0 {
		helper(nq, output)
		*output = append(*output, ld)
	}
}

