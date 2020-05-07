/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root==nil{
		return nil
	}
	pVal:=p.Val
	qVal:=q.Val
	node:=root
	for node!=nil{
		parentVal:=node.Val
		if pVal>parentVal&&qVal>parentVal{
			node = node.Right
			continue
		}
		if pVal<parentVal&&qVal<parentVal{
			node =node.Left
			continue
		}
		break		
	}
	return node
}
// @lc code=end

