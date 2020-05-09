/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
 func findMode(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return nil
	}
	m := &Mode{
		count:    0,
		maxCount: -1,
		pre:      nil,
	}
	m.helper(root, &res)
	return res
}

type Mode struct {
	pre      *TreeNode
	maxCount int
	count    int
}

func (m *Mode) helper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	m.helper(root.Left, res)
	if m.pre != nil && root.Val == m.pre.Val {
		m.count++
	} else {
		m.count = 1
	}
	if m.count > m.maxCount {
		m.maxCount = m.count
		*res = []int{root.Val}
	} else if m.count == m.maxCount {
		*res = append(*res, root.Val)
	}
	m.pre = root
	m.helper(root.Right, res)
}
// @lc code=end

