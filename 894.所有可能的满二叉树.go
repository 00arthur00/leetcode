/*
 * @lc app=leetcode.cn id=894 lang=golang
 *
 * [894] 所有可能的满二叉树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var memo = make(map[int][]*TreeNode)

//递归:
//每个满二叉树含有3个或更多结点，在其根结点处有2个子结点。这些子结点 left 和 right 本身就是满二叉树。没有满二叉树具有正偶数个结点
//因此，对于 N \geq 3N≥3，我们可以设定如下的递归策略：\text{FBT}(N) =FBT(N)= [对于所有的 xx，所有的树的左子结点来自 \text{FBT}(x)FBT(x) 而右子结点来自 \text{FBT}(N-1-x)FBT(N−1−x)]。
func allPossibleFBT(N int) []*TreeNode {
	if _, ok := memo[N]; !ok {
		ans := make([]*TreeNode, 0)
		if N == 1 {
			ans = append(ans, &TreeNode{})
		} else if N%2 == 1 {
			for x := 0; x < N; x++ {
				for _, lt := range allPossibleFBT(x) {
					for _, rt := range allPossibleFBT(N - x - 1) {
						root := &TreeNode{}
						root.Left = lt
						root.Right = rt
						ans = append(ans, root)
					}
				}
			}
		}
		memo[N] = ans
	}
	return memo[N]
}
