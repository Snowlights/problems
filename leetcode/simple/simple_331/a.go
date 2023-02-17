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
