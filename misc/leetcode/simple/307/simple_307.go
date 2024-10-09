package simple_307

import (
	"bytes"
	"container/heap"
	"sort"
)

// 1
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {

	ans := 0
	for idx := range energy {
		i, e := energy[idx], experience[idx]
		if i >= initialEnergy {
			ans += i + 1 - initialEnergy
			initialEnergy = i + 1
		}
		if e >= initialExperience {
			ans += e + 1 - initialExperience
			initialExperience = e + 1
		}
		initialEnergy -= i
		initialExperience += e
	}

	return ans
}

// 2
func largestPalindromic(num string) string {
	h := make(map[byte]int)
	ans := []byte{}
	for _, v := range num {
		h[byte(v)]++
	}
	if h['0'] == len(num) {
		return "0"
	}

	for s := '9'; s > '0'; s-- {
		ans = append(ans, bytes.Repeat([]byte{byte(s)}, h[byte(s)]/2)...)
	}
	if len(ans) > 0 && h['0'] > 0 {
		ans = append(ans, bytes.Repeat([]byte{byte('0')}, h[byte('0')]/2)...)
	}
	idx := len(ans) - 1

	for s := '9'; s >= '0'; s-- {
		if h[byte(s)]&1 == 1 {
			ans = append(ans, byte(s))
			break
		}
	}

	for ; idx >= 0; idx-- {
		ans = append(ans, ans[idx])
	}
	return string(ans)
}

// 3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {

	sTop := make(map[*TreeNode]*TreeNode)
	q := []*TreeNode{root}
	var s *TreeNode
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				sTop[v.Left] = v
				q = append(q, v.Left)
			}
			if v.Right != nil {
				sTop[v.Right] = v
				q = append(q, v.Right)
			}
			if v.Val == start {
				s = v
			}
		}
	}
	vis, ans := make(map[*TreeNode]bool), -1
	vis[s] = true
	q = []*TreeNode{s}
	for ; len(q) > 0; ans++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil && !vis[v.Left] {
				q = append(q, v.Left)
				vis[v.Left] = true
			}
			if v.Right != nil && !vis[v.Right] {
				q = append(q, v.Right)
				vis[v.Right] = true
			}
			if p, ok := sTop[v]; ok && !vis[p] {
				q = append(q, p)
				vis[p] = true
			}
		}
	}

	return ans
}

// 4
func kSum(nums []int, k int) int64 {
	n := len(nums)
	sum := 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
		} else {
			nums[i] = -x
		}
	}
	sort.Ints(nums)
	h := &hp{{sum, 0}}
	for ; k > 1; k-- {
		p := heap.Pop(h).(pair)
		if p.i < n {
			heap.Push(h, pair{p.sum - nums[p.i], p.i + 1}) // 保留 nums[p.i-1]
			if p.i > 0 {
				heap.Push(h, pair{p.sum - nums[p.i] + nums[p.i-1], p.i + 1}) // 不保留 nums[p.i-1]
			}
		}
	}
	return int64((*h)[0].sum)
}

type pair struct{ sum, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].sum > h[j].sum }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
