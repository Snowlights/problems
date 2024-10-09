package simple_331

import (
	"container/heap"
	"math"
	"sort"
)

// 1
func pickGifts(gifts []int, k int) int64 {
	hp := &hp{
		gifts,
	}
	heap.Init(hp)
	for k > 0 {
		hp.IntSlice[0] = int(math.Sqrt(float64(hp.IntSlice[0])))
		heap.Fix(hp, 0)
		k--
	}
	ans := 0
	for _, v := range gifts {
		ans += v
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}

// 2
func vowelStrings(words []string, queries [][]int) []int {
	h := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	prefix := make([]int, len(words)+1)
	for i, v := range words {
		if h[v[0]] && h[v[len(v)-1]] {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = prefix[q[1]+1] - prefix[q[0]]
	}
	return ans
}

// 3
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		f0, f1 := 0, 0
		for _, v := range nums {
			if v <= mx {
				f0, f1 = f1, max(f0+1, f1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 4
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}

	mn, a := math.MaxInt32, []int{}
	for x, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		mn = min(mn, x)
		for c = abs(c) / 2; c > 0; c-- {
			a = append(a, x)
		}
	}

	sort.Ints(a) // 也可以用快速选择
	for _, x := range a[:len(a)/2] {
		ans += int64(min(x, mn*2))
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
