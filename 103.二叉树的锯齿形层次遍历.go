/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	left2right := true
	for len(q) > 0 {
		nq := make([]*TreeNode, 0)
		levelData := make([]int, 0, len(q))
		for _, node := range q {
			if node != nil {
				levelData = append(levelData, node.Val)
				nq = append(nq, node.Left, node.Right)
			}
		}
		if left2right {
			left2right = false
		} else {
			for i, j := 0, len(levelData)-1; i < j; i, j = i+1, j-1 {
				levelData[i], levelData[j] = levelData[j], levelData[i]
			}
			left2right = true
		}
		if len(levelData) > 0 {
			ans = append(ans, levelData)
		}
		q = nq
	}
	return ans
}

