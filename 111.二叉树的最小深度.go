/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		nextQ := make([]*TreeNode, 0)
		level++
		for _, node := range q {
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				nextQ = append(nextQ, node.Left)
			}
			if node.Right != nil {
				nextQ = append(nextQ, node.Right)
			}
		}
		q = nextQ
	}
	return level
}

//DFS
func minDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		l, r := 1+minDepth(root.Left), 1+minDepth(root.Right)
		if l > r {
			return r
		}
		return l
	} else if root.Left != nil {
		return 1 + minDepth(root.Left)
	}

	return 1 + minDepth(root.Right)
}

