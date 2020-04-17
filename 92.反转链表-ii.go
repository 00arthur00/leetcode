/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var pre, cur *ListNode = nil, head
	for m > 1 {
		pre = cur
		cur = cur.Next
		m--
		n--
	}

	cont, tail := pre, cur

	for n > 0 {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		n--
	}

	if cont == nil {
		head = pre
	} else {
		cont.Next = pre

	}
	tail.Next = cur

	return head
}
// @lc code=end

