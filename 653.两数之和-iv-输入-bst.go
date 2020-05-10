/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
 *
 * https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (54.98%)
 * Likes:    130
 * Dislikes: 0
 * Total Accepted:    15K
 * Total Submissions: 27.3K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * 给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
 * 
 * 案例 1:
 * 
 * 
 * 输入: 
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 * 
 * Target = 9
 * 
 * 输出: True
 * 
 * 
 * 
 * 
 * 案例 2:
 * 
 * 
 * 输入: 
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 * 
 * Target = 28
 * 
 * 输出: False
 * 
 * 
 * 
 * 
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
//dfs+hash
func findTarget(root *TreeNode, k int) bool {
	memo:=make(map[int]struct{})
	return find(root,memo,k)
}
func find(root *TreeNode, memo map[int]struct{}, k int) bool{
	if root==nil{
		return false
	}
	t:=k-root.Val
	if _,ok:=memo[t];ok{
		return true
	}
	memo[root.Val] = struct{}{}
	if find(root.Left,memo,k){
		return true
	}
	return find(root.Right,memo,k)
}
//dfs
func findTarget1(root *TreeNode, k int) bool {
	if root==nil{
		return false
	}
	return helper(root,root,k)
}
func helper(root,cur *TreeNode, t int) bool{
	if cur == nil{
		return false
	}
	if helper(root,cur.Left,t){
		return true
	}
	if t-cur.Val!=cur.Val&&exist(root,t-cur.Val){
		return true
	}
	return helper(root,cur.Right,t)
}
func exist(root *TreeNode, t int) bool{
	if root==nil{
		return false
	}
	if t==root.Val{
		return true
	}
	if t>root.Val{
		return exist(root.Right,t)
	}
	return exist(root.Left,t)
}
// @lc code=end

