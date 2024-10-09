package _600_1700

// 1667
// select user_id, CONCAT(UPPER(left(name, 1)), LOWER(RIGHT(name, length(name) - 1))) as name
// from Users
// order by user_id

// 1669
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	idx, cur := 0, list1
	a--
	for idx < a {
		cur = cur.Next
		idx++
	}
	next := cur.Next
	cur.Next = list2
	for cur.Next != nil {
		cur = cur.Next
	}
	for idx < b {
		next = next.Next
		idx++
	}
	cur.Next = next

	return list1
}
