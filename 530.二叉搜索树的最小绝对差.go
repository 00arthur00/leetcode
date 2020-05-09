/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {
	var pre *TreeNode
	min:=1<<31
	inorder(root,&pre, &min)
	return min
}
func inorder(root *TreeNode, pre **TreeNode, min *int){
	if root==nil{
		return 
	}
	inorder(root.Left, pre, min)
	if *pre != nil{
		diff:= root.Val-(*pre).Val
		if diff <0{
			diff = - diff
		}
		if diff < *min{
			*min = diff
		}
	}
	*pre = root
	inorder(root.Right, pre, min)
}
// @lc code=end

