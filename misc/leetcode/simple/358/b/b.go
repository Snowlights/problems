package main

import . "problems/testutil/leetcode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

// 2. 两数相加：自己和自己相加
func double(l1 *ListNode) *ListNode {
	dummy := &ListNode{} // 哨兵节点，作为新链表的头节点的前一个节点
	cur := dummy
	carry := 0 // 进位
	for l1 != nil {
		carry += l1.Val * 2                   // 节点值和进位加在一起
		cur.Next = &ListNode{Val: carry % 10} // 每个节点保存一个数位
		carry /= 10                           // 新的进位
		cur = cur.Next                        // 下一个节点
		l1 = l1.Next                          // 下一个节点
	}
	if carry != 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

func doubleIt(head *ListNode) *ListNode {
	head = reverseList(head)
	res := double(head) // 反转后，就变成【2. 两数相加】了
	return reverseList(res)
}
