/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
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
//利用next的节点做层次遍历
func connect(root *Node) *Node {
	if root==nil{
		return nil
	}

	leftMost := root
	for leftMost.Left!=nil{
		head := leftMost

		for head!=nil{
			head.Left.Next =  head.Right
			if head.Next!=nil{
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		leftMost = leftMost.Left
	}
	return root
}
// @lc code=end

