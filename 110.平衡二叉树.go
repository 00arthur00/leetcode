/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {

	if root==nil{
		return true
	}
	left,ok:=depth(root.Left)
	if !ok{
		return false
	}
	right,ok:=depth(root.Right)
	if !ok{
		return false
	}
	if left >= right && left-right <=1{
		return true
	}

	if right >= left &&right-left<=1{
		return true
	}
	return false
}

func depth(root *TreeNode) (int,bool){
	if root ==nil{
		return 0,true
	}
	left, ok:=depth(root.Left)
	if !ok{
		return -1, false
	}
	left++

	right, ok:=depth(root.Right)
	if !ok{
		return -1, false
	}
	right++ 
	if left >= right && left-right <=1{
		return left,true
	}

	if right >= left &&right-left<=1{
		return right,true
	}
	return right, false
} 
// @lc code=end

