/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//1. 递归解法
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)

}
func isMirror(l, r *TreeNode) bool {
	//都为空,则返回true
	if l == nil && r == nil {
		return true
	}
	//非全空,一个未空,一个不未空,返回false
	if l == nil || r == nil {
		return false
	}
	//val相等,且左子树的左节点==右字数的右节点&&左子树的右节点==右字数的左节点
	return l.Val == r.Val && isMirror(l.Right, r.Left) && isMirror(l.Left, r.Right)
}

//2. 非递归,类广度优先遍历解法
func isSymmetricIterator(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := []*TreeNode{root.Left, root.Right}
	for len(q) > 0 {
		//两个两个出队列
		l, r := q[0], q[1]
		q = q[2:len(q)]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		//左子树的左节点==右子树的右节点,先这两个进入Q
		//左子树的右节点==右子树的左节点,再这两个进入Q
		q = append(q, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}

