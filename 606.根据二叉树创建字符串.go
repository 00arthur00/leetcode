/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
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
func tree2str(t *TreeNode) string {
	if t==nil{
		return ""
	}
	//case1, 左右都空,直接返回相应val
	if t.Left==nil&&t.Right==nil{
		return strconv.Itoa(t.Val)
	}
	//左边不空,右边空
	if t.Right==nil{
		return strconv.Itoa(t.Val)+"("+tree2str(t.Left)+")"
	}
	//右边空(返回"")或者左右都不空
	return strconv.Itoa(t.Val)+"("+tree2str(t.Left)+")("+tree2str(t.Right)+")"
	
}
// @lc code=end

