/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层次遍历
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//遍历解法,用队列来处理
func levelOrderIterator(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		nq := []*TreeNode{}
		output := []int{}
		for _, tn := range q {
			if tn != nil {
				nq = append(nq, tn.Left, tn.Right)
				output = append(output, tn.Val)
			}
		}
		if len(nq) > 0 {
			ans = append(ans, output)
		}
		q = nq
	}
	return ans
}

//递归解法,类似于线序遍历,但添加了一个level变量用于控制结果的输出
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	helper(root, 0, &ans)
	return ans
}
func helper(node *TreeNode, level int, output *[][]int) {
	if node == nil {
		return
	}
	if level == len(*output) {
		*output = append(*output, []int{})
	}
	(*output)[level] = append((*output)[level], node.Val)
	if node.Left != nil {
		helper(node.Left, level+1, output)
	}
	if node.Right != nil {
		helper(node.Right, level+1, output)
	}
}
