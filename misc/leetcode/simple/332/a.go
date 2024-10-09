package simple_332

import (
	"container/heap"
	"math"
	"sort"
	"strings"
)

// 1
func pickGifts(gifts []int, k int) (ans int64) {
	h := &hp{gifts}
	heap.Init(h)
	for ; k > 0 && gifts[0] > 1; k-- {
		gifts[0] = int(math.Sqrt(float64(gifts[0])))
		heap.Fix(h, 0)
	}
	for _, x := range gifts {
		ans += int64(x)
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Pop() (_ interface{}) { return }
func (hp) Push(interface{})     {}

// 2
func vowelStrings(words []string, queries [][]int) []int {
	sum := make([]int, len(words)+1)
	for i, w := range words {
		sum[i+1] = sum[i]
		if strings.Contains("aeiou", w[:1]) && strings.Contains("aeiou", w[len(w)-1:]) {
			sum[i+1]++
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sum[q[1]+1] - sum[q[0]]
	}
	return ans
}

// 3
func minCapability(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		f0, f1 := 0, 0
		for _, x := range nums {
			if x <= mx {
				f0, f1 = f1, max(f1, f0+1)
			} else {
				f0 = f1
			}
		}
		return f1 >= k
	})
}

func max(a, b int) int {
	if b > a {
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
