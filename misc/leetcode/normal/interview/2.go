package interview

// 面试题 02.01
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	h := make(map[int]bool)
	cur := head
	pre := &ListNode{Next: head}
	for cur != nil {
		if h[cur.Val] {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		h[cur.Val] = true
		pre = cur
		cur = cur.Next
	}
	return head
}

// 面试题 02.02
func kthToLast(head *ListNode, k int) int {
	pre, cur := head, head
	for k > 0 {
		k--
		cur = cur.Next
	}
	for cur != nil {
		pre = pre.Next
		cur = cur.Next
	}
	return pre.Val
}

// 面试题 02.03
func deleteNode(node *ListNode) {
	*node = *node.Next
}

// 面试题 02.04
func partition(head *ListNode, x int) *ListNode {
	a, b := make([]*ListNode, 0), make([]*ListNode, 0)
	cur := head
	for cur != nil {
		if cur.Val < x {
			a = append(a, cur)
			cur = cur.Next
			a[len(a)-1].Next = nil
		} else {
			b = append(b, cur)
			cur = cur.Next
			b[len(b)-1].Next = nil
		}
	}

	newHead := &ListNode{}
	cur = newHead
	for len(a) > 0 {
		now := a[0]
		a = a[1:]
		cur.Next = now
		cur = cur.Next
	}
	for len(b) > 0 {
		now := b[0]
		b = b[1:]
		cur.Next = now
		cur = cur.Next
	}
	return newHead.Next
}

// 面试题 02.05
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur, f := &ListNode{}, 0
	p := cur
	for l1 != nil && l2 != nil {
		node := &ListNode{}
		val := l1.Val + l2.Val + f
		f = val / 10
		node.Val = val % 10
		cur.Next = node
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		node := &ListNode{}
		val := l1.Val + f
		f = val / 10
		node.Val = val % 10
		cur.Next = node
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		node := &ListNode{}
		val := l2.Val + f
		f = val / 10
		node.Val = val % 10
		cur.Next = node
		cur = cur.Next
		l2 = l2.Next
	}
	if f == 1 {
		cur.Next = &ListNode{Val: 1}
	}

	return p.Next
}

// 面试题 02.06
func isPalindrome(head *ListNode) bool {
	a := make([]*ListNode, 0)
	for head != nil {
		a = append(a, head)
		head = head.Next
	}
	for i := 0; i < len(a); i++ {
		if a[i].Val != a[len(a)-1-i].Val {
			return false
		}
	}
	return true
}

// 面试题 02.07
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	h := make(map[*ListNode]bool)
	for cur := headA; cur != nil; cur = cur.Next {
		h[cur] = true
	}
	for cur := headB; cur != nil; cur = cur.Next {
		if h[cur] {
			return cur
		}
	}
	return nil
}

// 面试题 02.08
func detectCycle(head *ListNode) *ListNode {
	h := make(map[*ListNode]bool)
	for cur := head; cur != nil; cur = cur.Next {
		if h[cur] {
			return cur
		}
		h[cur] = true
	}
	return nil
}
