/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
 *
 * https://leetcode-cn.com/problems/trim-a-binary-search-tree/description/
 *
 * algorithms
 * Easy (65.31%)
 * Likes:    203
 * Dislikes: 0
 * Total Accepted:    11.9K
 * Total Submissions: 18.2K
 * Testcase Example:  '[1,0,2]\n1\n2'
 *
 * 给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L)
 * 。你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。
 * 
 * 示例 1:
 * 
 * 
 * 输入: 
 * ⁠   1
 * ⁠  / \
 * ⁠ 0   2
 * 
 * ⁠ L = 1
 * ⁠ R = 2
 * 
 * 输出: 
 * ⁠   1
 * ⁠     \
 * ⁠      2
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: 
 * ⁠   3
 * ⁠  / \
 * ⁠ 0   4
 * ⁠  \
 * ⁠   2
 * ⁠  /
 * ⁠ 1
 * 
 * ⁠ L = 1
 * ⁠ R = 3
 * 
 * 输出: 
 * ⁠     3
 * ⁠    / 
 * ⁠  2   
 * ⁠ /
 * ⁠1
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
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val>=L&&root.Val<=R{
		root.Left =trimBST(root.Left,L,root.Val)
		root.Right = trimBST(root.Right,root.Val,R)
		return root
	}
	if root.Val<L{
		return trimBST(root.Right,L,R)
	}
	return trimBST(root.Left,L,R)
}
// @lc code=end

