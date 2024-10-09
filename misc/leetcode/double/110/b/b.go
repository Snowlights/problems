package main

import (
	. "problems/testutil/leetcode"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) (ans *ListNode) {
	gcd := func(a, b int) int {
		for a%b != 0 {
			a, b = b, a%b
		}
		return b
	}

	cur, next := head, head.Next
	for cur != nil && next != nil {
		val := gcd(cur.Val, next.Val)
		node := &ListNode{
			Val:  val,
			Next: next,
		}
		cur.Next = node
		cur = next
		next = next.Next
	}

	return head
}
