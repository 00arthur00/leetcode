/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	last := dummy
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			last.Next, l2 = &ListNode{Val: l2.Val}, l2.Next
		} else {
			last.Next, l1 = &ListNode{Val: l1.Val}, l1.Next
		}
		last = last.Next
	}

	for l1 != nil {
		last.Next, l1 = &ListNode{Val: l1.Val}, l1.Next
		last = last.Next
	}

	for l2 != nil {
		last.Next, l2 = &ListNode{Val: l2.Val}, l2.Next
		last = last.Next
	}
	return dummy.Next
}

