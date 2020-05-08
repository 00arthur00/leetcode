/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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

func binaryTreePaths(root *TreeNode) []string {
	if root==nil{
		return nil
	}
	res:=make([]string,0)
	helper(root,"",&res)
	return res
}
func helper(root *TreeNode,path string,res *[]string) {
	if root==nil{
		return
	}
	path =  path+strconv.Itoa(root.Val)
	if root.Left==nil&&root.Right==nil{
		*res = append(*res,path)
		return 
	}
	path += "->"

	helper(root.Left,path,res)
	helper(root.Right,path,res)
}
// @lc code=end

