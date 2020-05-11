/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var front *ListNode
func isPalindrome(head *ListNode) bool {
	if head==nil{
		return true
	}
	front = head
	return check(head)
}
func check(head *ListNode) bool{
	if head==nil{
		return true
	}
	if false==check(head.Next){
		return false
	}
	if front.Val!=head.Val{
		return false
	}
	front = front.Next
	return true
}
// @lc code=end

