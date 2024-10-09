package main

import (
	"container/heap"
	"sort"
)

func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	sort.Slice(items, func(i, j int) bool { return items[i][0] > items[j][0] })
	ans, totalProfit := 0, 0
	vis := map[int]bool{}
	duplicate := hp{} // 重复类别的利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit
			if !vis[category] {
				vis[category] = true
			} else { // 重复类别
				heap.Push(&duplicate, profit)
			}
		} else if !vis[category] && duplicate.Len() > 0 {
			vis[category] = true
			totalProfit += profit - heap.Pop(&duplicate).(int) // 选一个重复类别中的最小利润替换
		} // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
