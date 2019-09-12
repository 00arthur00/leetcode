/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans int

func longestUnivaluePath1(root *TreeNode) int {
	ans = 0
	arrowLength(root)
	return ans
}
func arrowLength(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := arrowLength(node.Left)
	right := arrowLength(node.Right)
	arrowLeft, arrowRight := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		arrowLeft += left + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		arrowRight += right + 1
	}
	if arrowLeft+arrowRight > ans {
		ans = arrowLeft + arrowRight
	}
	if arrowLeft > arrowRight {
		return arrowLeft
	}
	return arrowRight
}

func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	nodeCount(root, &ans)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func nodeCount(root *TreeNode, sum *int) int {
	if root == nil {
		return 0
	}
	left, right := nodeCount(root.Left, sum), nodeCount(root.Right, sum)
	ll, rl := 0, 0
	if root.Left != nil && root.Left.Val == root.Val {
		ll += 1 + left
	}
	if root.Right != nil && root.Right.Val == root.Val {
		rl += 1 + right
	}
	if *sum < ll+rl {
		*sum = ll + rl
	}
	return max(ll, rl)
}

