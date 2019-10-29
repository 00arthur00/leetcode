/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	basePoint := hasCircle(head)
	if basePoint == nil {
		return basePoint
	}
	for head != basePoint {
		head = head.Next
		basePoint = basePoint.Next
	}
	return basePoint
}
func hasCircle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, quick := head, head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
		if slow == quick {
			return slow
		}
	}
	return nil
}

// @lc code=end

