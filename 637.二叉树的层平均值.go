/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val float64
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func averageOfLevels(root *TreeNode) []float64 {
	sum, count := make([]float64,0),make([]float64,0) 
	helper(root,1,&sum,&count)

	for i,_:=range sum{
		sum[i] = sum[i]/count[i]
	}
	return sum
}
func helper(root*TreeNode, level int,sum,count *[]float64){
	if root==nil{
		return 
	}
	if level>len(*sum){
		*sum = append(*sum,0)
		*count = append(*count,0)
	}
	(*sum)[level-1]+=float64(root.Val)
	(*count)[level-1]++
	helper(root.Left,level+1,sum,count)
	helper(root.Right,level+1,sum,count)
}
// @lc code=end

