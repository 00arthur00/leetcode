/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath, qPath := make([]*TreeNode, 0), make([]*TreeNode, 0)
	getpath(root, p, &pPath)
	getpath(root, q, &qPath)
	minLen := len(pPath)
	if len(qPath) < len(pPath) {
		minLen = len(qPath)
	}

	var node *TreeNode
	for i := 0; i < minLen; i++ {
		if pPath[i] != qPath[i] {
			break
		}
		node = pPath[i]
	}
	return node
}
func getpath(root, node *TreeNode, path *[]*TreeNode) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root)
	if root == node {
		return true
	}
	if getpath(root.Left, node, path) {
		return true
	}
	if getpath(root.Right, node, path) {
		return true
	}
	*path = (*path)[:len(*path)-1]
	return false
}
// @lc code=end

