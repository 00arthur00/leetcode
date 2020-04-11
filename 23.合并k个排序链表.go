/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummy:=&ListNode{}
	tail :=dummy
	for len(lists) > 1 {
		index:=findMinIndex(lists)
		tail.Next = &ListNode{Val:lists[index][0].Val}
		lists[index] = list[index].Next
		if list[index]==nil{
			list = append(list[0:index-1],lists[index+1]...)
		}
	}
	tail.Next = lists[0]
	return dummy.Next
}
func findMinIndex(ls []*ListNode) int {
	index := 0
	min := ls[0].Val
	for i := 1; i < len(ls); i++ {
		ls[i].Val < min{
			min=ls[i].Val
			index=i
		}
	}
	return index
}
