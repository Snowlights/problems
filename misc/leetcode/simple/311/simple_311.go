package simple_311

// 1
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}

// 2
func longestContinuousSubstring(s string) int {
	st, e, ans := 0, 1, 0
	for e < len(s) {
		for e < len(s) && s[e]-s[e-1] == 1 {
			e++
		}
		ans = max(ans, e-st)
		st = e
		e++
	}
	return max(ans, e-st)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {

	q := []*TreeNode{root}
	idx := 0
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
		}

		if idx%2 == 1 {
			s, e := 0, len(tmp)-1
			for s < e {
				tmp[s].Val, tmp[e].Val = tmp[e].Val, tmp[s].Val
				s++
				e--
			}
		}

		idx++
	}

	return root
}

// 4
func sumPrefixScores(words []string) []int {
	h := make(map[string]int)
	for _, w := range words {
		s, e := 1, len(w)
		for s <= e {
			h[w[:s]]++
			s++
		}
	}
	ans := make([]int, len(words))
	for i, w := range words {
		s, e := 1, len(w)
		for s <= e {
			ans[i] += h[w[:s]]
			s++
		}
	}
	return ans
}
