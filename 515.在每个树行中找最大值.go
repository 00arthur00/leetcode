/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
//计算层数,必须用先序遍历
func largestValues(root *TreeNode) []int {
	ans:=make([]int, 1)
	if root==nil{
		return nil
	}
	level:=1
	ans[0]=root.Val
	helper(root, level,&ans)
	return ans
}
func helper(root *TreeNode, level int, ans *[]int){
	if root==nil{
		return
	}
	if level>len(*ans){
		*ans = append(*ans,-1<<31)
	}
	if root.Val > (*ans)[level-1]{
		(*ans)[level-1] = root.Val
	}
	helper(root.Left, level+1,ans)
	helper(root.Right, level+1,ans)
}
// @lc code=end

