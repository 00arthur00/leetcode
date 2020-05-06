/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, sum int) [][]int {
	current := sum
	path:=make([]int,0)
	result := make([][]int,0)
	helper(root, current,path, &result)
	return result
}
func helper(root *TreeNode, current int, path []int, result *[][]int){
	if root ==nil{
		return
	}
	path =  append(path, root.Val)

	left := current - root.Val
	if  left== 0 && root.Left==nil&&root.Right==nil{
		resultNode:=make([]int,len(path))
		copy(resultNode,path)
		*result = append(*result,resultNode)
		return
	}

	helper(root.Left, left, path, result)
	helper(root.Right,left, path, result)

}
// @lc code=end

