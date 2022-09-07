package _00_300

import (
	"container/heap"
	"math"
	"sort"
)

// 200
func numIslands(grid [][]byte) int {

	ans, n, m := 0, len(grid), len(grid[0])
	dir := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var dfs func([][]byte, int, int)
	dfs = func(g [][]byte, x, y int) {
		g[x][y] = '0'
		for _, d := range dir {
			xx, yy := d[0]+x, d[1]+y
			if 0 <= xx && xx < n && 0 <= yy && yy < m && g[xx][yy] == '1' {
				dfs(g, xx, yy)
			}
		}
	}
	for i, v := range grid {
		for j, d := range v {
			if d == '1' {
				ans++
				dfs(grid, i, j)
			}
		}
	}

	return ans
}

// 202
func isHappy(n int) bool {

	if n == 1 {
		return true
	}

	exitMap := make(map[int]bool)
	res := n
	for res != 1 {
		res = calculate(res)
		if exitMap[res] {
			return false
		}
		exitMap[res] = true
	}

	return true
}

func calculate(n int) int {

	res := 0
	for n > 0 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	return res
}

// 205
func isIsomorphic(s string, t string) bool {

	sh, th := make(map[byte]byte), make(map[byte]byte)
	for i, v := range s {
		val, ok := sh[byte(v)]
		if ok && val != t[i] {
			return false
		}
		val, ok = th[byte(t[i])]
		if ok && val != byte(v) {
			return false
		}
		sh[byte(v)] = t[i]
		th[byte(t[i])] = byte(v)
	}
	return true
}

// 206
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre, cur, next *ListNode
	cur, next = head, head.Next
	for cur.Next != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	cur.Next = pre
	return cur
}

// 209
func minSubArrayLen(target int, nums []int) int {

	left, sum, ans := 0, 0, math.MaxInt32
	for r, v := range nums {
		sum += v
		if sum >= target {
			ans = min(ans, r-left+1)
		}
		for left <= r && sum-nums[left] >= target {
			sum -= nums[left]
			left++
			ans = min(ans, r-left+1)
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 213
func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 217
func containsDuplicate(nums []int) bool {
	h := make(map[int]bool)
	for _, v := range nums {
		if h[v] {
			return true
		}
		h[v] = true
	}
	return false
}

// 218
type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
	n := len(buildings)
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	idx := 0
	h := hp{}
	for _, boundary := range boundaries {
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}

		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}
	return
}
