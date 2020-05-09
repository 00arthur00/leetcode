/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一个树的子树
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
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s==nil&&t==nil{
		return true
	}
	if s==nil||t==nil{
		return false
	}
	if isSubtree(s.Left,t){
		return true
	}
	if isSameTree(s,t){
		return true
	}
	if isSubtree(s.Right,t){
		return true
	}
	return false
}
func isSameTree(l,r *TreeNode) bool{
	if l==nil&&r==nil{
		return true
	}
	if l!=nil&&r!=nil&&l.Val==r.Val{
		if !isSameTree(l.Left,r.Left){
			return false
		}
		if !isSameTree(l.Right,r.Right){
			return false
		}
		return true
	}
	return false
}
// @lc code=end

