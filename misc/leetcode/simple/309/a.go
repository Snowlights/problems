package simple_309

import (
	"container/heap"
	"sort"
)

// 1
func checkDistances(s string, distance []int) bool {
	h := make(map[int][]int)
	for i, d := range s {
		val := d - 'a'
		h[int(val)] = append(h[int(val)], i)
	}
	for i, d := range distance {
		val, ok := h[int(i)]
		if !ok {
			continue
		}
		if val[1]-val[0]-1 != d {
			return false
		}
	}
	return true
}

// 2
func numberOfWays(startPos int, endPos int, k int) int {

	const mod = 1e9 + 7
	type pair struct {
		pos, step int
	}
	dp := map[pair]int{}

	var dfs func(pos, step int) int
	dfs = func(pos, step int) int {
		if abs(pos-endPos) > step {
			return 0
		}

		if step == 0 {
			return 1
		}
		p := pair{
			pos, step,
		}
		if val, ok := dp[p]; ok {
			return val
		}

		ans := (dfs(pos-1, step-1) + dfs(pos+1, step-1)) % mod
		dp[p] = ans
		return ans
	}

	return dfs(startPos, k)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 3
func longestNiceSubarray(nums []int) int {
	ans := 1
	for i := 1; i < len(nums); i++ {
		j, base := i-1, nums[i]
		for j >= 0 && base&nums[j] == 0 {
			base |= nums[j]
			j--
		}
		ans = max(ans, i-j)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 4
func mostBooked(n int, meetings [][]int) (ans int) {
	cnt := make([]int, n)
	idle := hp{make([]int, n)}
	for i := 0; i < n; i++ {
		idle.IntSlice[i] = i
	}
	using := hp2{}
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	for _, m := range meetings {
		st, end := m[0], m[1]
		for len(using) > 0 && using[0].end <= st {
			heap.Push(&idle, heap.Pop(&using).(pair).i) // 维护在 st 时刻空闲的会议室
		}
		var i int
		if idle.Len() == 0 {
			p := heap.Pop(&using).(pair) // 没有可用的会议室，那么弹出一个最早结束的会议室
			end += p.end - st            // 更新当前会议的结束时间
			i = p.i
		} else {
			i = heap.Pop(&idle).(int)
		}
		cnt[i]++
		heap.Push(&using, pair{end, i}) // 使用一个会议室
	}
	for i, c := range cnt {
		if c > cnt[ans] {
			ans = i
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type pair struct{ end, i int }
type hp2 []pair

func (h hp2) Len() int { return len(h) }
func (h hp2) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.end < b.end || a.end == b.end && a.i < b.i
}
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
