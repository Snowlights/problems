package main

import (
	"container/heap"
	"sort"
)

func minOperations(nums []int, k int) int {
	ans, hp := 0, &hp2{nums}
	heap.Init(hp)
	for hp.IntSlice[0] < k {
		v := heap.Pop(hp).(int)
		hp.IntSlice[0] += v * 2
		heap.Fix(hp, 0)
		ans++
	}
	return ans
}

type hp2 struct{ sort.IntSlice }

func (h hp2) Less(i, j int) bool  { return h.IntSlice[i] < h.IntSlice[j] }
func (h *hp2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
