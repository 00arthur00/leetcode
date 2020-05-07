/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
 type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	return BSTIterator{
		stack: stack,
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	val:=node.Val
	this.stack = this.stack[:len(this.stack)-1]
	if node.Right != nil {
		this.stack = append(this.stack, node.Right)
		node = node.Right
		for node.Left!=nil{
			this.stack = append(this.stack, node.Left)
			node = node.Left
		}
	}
	return val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.stack) == 0 {
		return false
	}
	return true
}



/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

