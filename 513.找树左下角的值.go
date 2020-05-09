/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	maxLevel:=-1
	ans:=root.Val
	helper(root,0,&maxLevel,&ans)
	return ans
}
func helper(root *TreeNode, currentLevel int, maxLevel *int, ans *int){
	if root==nil{
		return 
	}
	helper(root.Left,currentLevel+1,maxLevel,ans)
	if root.Left!=nil&&currentLevel+1>*maxLevel{
		*ans = root.Left.Val
		*maxLevel = currentLevel+1
	}else if root.Right!=nil && currentLevel+1 > *maxLevel{
		*ans = root.Right.Val
		*maxLevel = currentLevel+1
	}
	helper(root.Right,currentLevel+1,maxLevel,ans)
}
// @lc code=end

