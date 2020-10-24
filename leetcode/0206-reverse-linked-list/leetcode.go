package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var res *ListNode
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}
