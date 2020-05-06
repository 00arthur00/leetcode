/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root==nil{
		return nil
	}
	leftMost := root
	cur:= leftMost
	var pre *Node

	processChild :=func(childNode *Node){
		if childNode==nil{
			return
		}
		if pre!=nil{
			pre.Next = childNode
		}else{
			leftMost = childNode
		}
		pre = childNode
	}

	for leftMost!=nil{
		pre = nil
		cur = leftMost

		leftMost = nil

		for cur!=nil{
			processChild(cur.Left)
			processChild(cur.Right)
			cur = cur.Next
		}
	}
	return root
}
// @lc code=end

