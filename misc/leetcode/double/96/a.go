package double_96

import (
	"container/heap"
	"sort"
)

// 1
func getCommon(nums1 []int, nums2 []int) int {
	h := make(map[int]bool)
	for _, v := range nums1 {
		h[v] = true
	}
	for _, v := range nums2 {
		if h[v] {
			return v
		}
	}
	return -1
}

// 2
func minOperations(nums1, nums2 []int, k int) (ans int64) {
	sum := 0
	for i, x := range nums1 {
		x -= nums2[i]
		if k > 0 {
			if x%k != 0 {
				return -1
			}
			sum += x / k
			if x > 0 {
				ans += int64(x / k)
			}
		} else if x != 0 {
			return -1
		}
	}
	if sum != 0 {
		return -1
	}
	return
}

// 3
func maxScore(nums1, nums2 []int, k int) int64 {
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	sum := 0
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].y > a[j].y })

	h := hp{nums2[:k]} // 复用内存
	for i, p := range a[:k] {
		sum += p.x
		h.IntSlice[i] = p.x
	}
	ans := sum * a[k-1].y
	heap.Init(&h)
	for _, p := range a[k:] {
		if p.x > h.IntSlice[0] {
			sum += p.x - h.replace(p.x)
			ans = max(ans, sum*p.y)
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}
func (h hp) replace(v int) int  { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(&h, 0); return top }
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 4
func isReachable(targetX int, targetY int) bool {
	g := gcd(targetX, targetY)
	return g&(g-1) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
