/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, quick := head, head.Next
	for slow != nil && quick != nil {
		if slow == quick {
			return true
		}
		slow = slow.Next
		if quick.Next == nil {
			break
		}
		quick = quick.Next.Next
	}
	return false
}

// @lc code=end

