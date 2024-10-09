package simple_327

import (
	"container/heap"
	"sort"
)

// 1
func maximumCount(nums []int) int {
	neg, pos := 0, 0
	for _, v := range nums {
		if v < 0 {
			neg++
		}
		if v > 0 {
			pos++
		}
	}

	return max(neg, pos)
}

// 2
//func maxKelements(nums []int, k int) int64 {
//	h := &hp2{nums}
//	heap.Init(h)
//	ans := 0
//	for ; k > 0; k-- {
//		val := heap.Pop(h).(int)
//		ans += val
//		heap.Push(h, (val+2)/3)
//	}
//	return int64(ans)
//}
//
//type hp2 struct{ sort.IntSlice }
//
//func (h hp2) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
//func (h *hp2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
//func (h *hp2) Pop() interface{} {
//	a := h.IntSlice
//	v := a[len(a)-1]
//	h.IntSlice = a[:len(a)-1]
//	return v
//}

// 3
func isItPossible(word1, word2 string) bool {
	c1 := map[rune]int{}
	for _, c := range word1 {
		c1[c]++
	}
	c2 := map[rune]int{}
	for _, c := range word2 {
		c2[c]++
	}
	for x, c := range c1 {
		for y, d := range c2 {
			if y == x { // 无变化
				if len(c1) == len(c2) {
					return true
				}
			} else if len(c1)-b2i(c == 1)+b2i(c1[y] == 0) ==
				len(c2)-b2i(d == 1)+b2i(c2[x] == 0) { // 计算变化量
				return true
			}
		}
	}
	return false
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

// 4
func findCrossingTime(n, k int, time [][]int) (cur int) {
	sort.SliceStable(time, func(i, j int) bool {
		a, b := time[i], time[j]
		return a[0]+a[2] < b[0]+b[2]
	})
	waitL, waitR := make(hp, k), hp{}
	for i := range waitL {
		waitL[i].i = k - 1 - i
	}
	workL, workR := hp2{}, hp2{}
	for n > 0 {
		for workL.Len() > 0 && workL[0].t <= cur {
			heap.Push(&waitL, heap.Pop(&workL)) // 左边完成放箱
		}
		for workR.Len() > 0 && workR[0].t <= cur {
			heap.Push(&waitR, heap.Pop(&workR)) // 右边完成搬箱
		}
		if waitR.Len() > 0 && waitR[0].t <= cur { // 右边过桥
			p := heap.Pop(&waitR).(pair)
			cur += time[p.i][2]
			heap.Push(&workL, pair{p.i, cur + time[p.i][3]}) // 放箱，记录完成时间
		} else if waitL.Len() > 0 && waitL[0].t <= cur { // 左边过桥
			p := heap.Pop(&waitL).(pair)
			cur += time[p.i][0]
			heap.Push(&workR, pair{p.i, cur + time[p.i][1]}) // 搬箱，记录完成时间
			n--
		} else if workL.Len() == 0 { // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
			cur = workR[0].t
		} else if workR.Len() == 0 {
			cur = workL[0].t
		} else {
			cur = min(workL[0].t, workR[0].t)
		}
	}
	for workR.Len() > 0 {
		p := heap.Pop(&workR).(pair)       // 右边完成搬箱
		cur = max(p.t, cur) + time[p.i][2] // 过桥
	}
	return cur // 最后一个过桥的时间
}

type pair struct{ i, t int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type hp2 []pair

func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
