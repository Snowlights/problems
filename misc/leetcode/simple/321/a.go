package simple_321

// 1
func pivotInteger(n int) int {
	pre, suf := 0, n*(n+1)/2
	for i := 1; i <= n; i++ {
		pre += i
		if pre == suf {
			return i
		}
		suf -= i
	}
	return -1
}

// 2

func appendCharacters(s string, t string) int {
	m, n := len(s), len(t)
	i, j := 0, 0
	for i < m && j < n {
		if s[i] == t[j] {
			i++
			j++
			continue
		}
		i++
	}
	return n - j
}

// 3
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	h := make([]*ListNode, 0, 1000)
	cur := head
	for cur != nil {
		h = append(h, cur)
		pre := cur
		cur = cur.Next
		pre.Next = nil
	}

	newHead, val := &ListNode{}, 0
	for i := len(h) - 1; i >= 0; i-- {
		if h[i].Val >= val {
			h[i].Next = newHead.Next
			newHead.Next = h[i]
			val = h[i].Val
		}
	}

	return newHead.Next
}

// 4
func countSubarrays(nums []int, k int) int {
	h, pos := make(map[int]int), 0
	for i, v := range nums {
		if v == k {
			pos = i
			break
		}
	}

	h[0]++
	c, ans := 0, 0
	for _, v := range nums[pos+1:] {
		if v > k {
			c++
		} else {
			c--
		}
		h[c]++
	}
	c = 0
	ans += h[0] + h[1]
	for i := pos - 1; i >= 0; i-- {
		if nums[i] > k {
			c--
		} else {
			c++
		}
		ans += h[c] + h[c+1]
	}

	return ans
}
