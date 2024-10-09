package _00_900

import "sort"

// 875
func minEatingSpeed(piles []int, h int) int {
	return sort.Search(1e9, func(i int) bool {
		tmp := 0
		if i == 0 {
			return false
		}
		for _, p := range piles {
			tmp += p / i
			if p%i > 0 {
				tmp++
			}
		}
		return tmp <= h
	})
}

// 876
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, next := head, head
	for next != nil && next.Next != nil {
		pre = pre.Next
		next = next.Next.Next
	}
	return pre
}

// 886
const (
	UNCOLORED, RED, GREEN = 0, 1, 2
)

func possibleBipartition(n int, dislikes [][]int) bool {

	color := make([]int, n+1)
	g := make(map[int][]int)
	for _, dis := range dislikes {
		g[dis[0]] = append(g[dis[0]], dis[1])
		g[dis[1]] = append(g[dis[1]], dis[0])
	}

	for i := 1; i <= n; i++ {
		if color[i] == UNCOLORED {
			color[i] = RED
			q := []int{i}
			for len(q) > 0 {
				tmp := q
				q = nil
				for _, v := range tmp {
					c := RED
					if color[v] == RED {
						c = GREEN
					}
					for _, vv := range g[v] {
						if color[vv] == UNCOLORED {
							color[vv] = c
							q = append(q, vv)
						} else if color[vv] != c {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

func buildGraph(dislikes [][]int, n int, graph [][]bool) {
	for _, d := range dislikes {
		graph[d[0]][d[1]] = true
		graph[d[1]][d[0]] = true
	}
}

// 891
func sumSubseqWidths(nums []int) int {
	const (
		mod int = 1e9 + 7
	)
	ans, n := 0, len(nums)

	pow := make([]int, n+1)
	pow[0] = 1
	for i := 1; i <= n; i++ {
		pow[i] = (pow[i-1] * 2) % mod
	}

	sort.Ints(nums)
	for i, v := range nums {
		ans = (ans + (pow[i]-pow[n-1-i])*v) % mod
	}

	return ans
}

// 893
func mincostTickets(days []int, costs []int) int {
	sort.Ints(days)
	dp := make([]int, days[len(days)-1]+30)
	for _, d := range days {
		dp[d] = 1
	}
	for d := 1; d < days[len(days)-1]+30; d++ {
		if dp[d] > 0 {
			dp[d] = min(dp[max(d-1, 0)]+costs[0], min(
				dp[max(d-7, 0)]+costs[1], dp[max(d-30, 0)]+costs[2]))
		} else {
			dp[d] = dp[d-1]
		}
	}
	return dp[days[len(days)-1]+29]
}

// 894
func allPossibleFBT(n int) []*TreeNode {

	ans := make([]*TreeNode, 0)
	if n%2 == 0 {
		return ans
	}

	if n == 1 {
		return append(ans, &TreeNode{})
	}

	for i := 1; i < n; i += 2 {
		l, r := allPossibleFBT(i), allPossibleFBT(n-1-i)
		for _, v := range l {
			for _, v2 := range r {
				ans = append(ans, &TreeNode{Left: v, Right: v2})
			}
		}
	}

	return ans
}

// 895
type FreqStack struct {
	h     map[int]int
	stack [][]int
}

func Constructor() FreqStack {
	return FreqStack{
		h:     map[int]int{},
		stack: [][]int{},
	}
}

func (this *FreqStack) Push(val int) {
	c := this.h[val]
	if c == len(this.stack) {
		this.stack = append(this.stack, []int{val})
	} else {
		this.stack[c] = append(this.stack[c], val)
	}
	this.h[val]++
}

func (this *FreqStack) Pop() int {
	sLen := len(this.stack) - 1
	l := len(this.stack[sLen]) - 1
	val := this.stack[sLen][l]

	if l == 0 {
		this.stack = this.stack[:sLen]
	} else {
		this.stack[sLen] = this.stack[sLen][:l]
	}
	this.h[val]--
	return val
}

// 896
func isMonotonic(nums []int) bool {
	return sort.IsSorted(sort.IntSlice(nums)) || sort.IsSorted(sort.Reverse(sort.IntSlice(nums)))
}

// 898
func subarrayBitwiseORs(arr []int) int {
	cur, ans := make(map[int]bool), make(map[int]bool)
	for _, v := range arr {
		cur2 := map[int]bool{v: true}
		for k := range cur {
			cur2[k|v] = true
		}
		cur = cur2
		for k := range cur {
			ans[k|v] = true
		}
	}

	return len(ans)
}
