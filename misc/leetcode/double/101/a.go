package double_101

import (
	"math"
	"math/bits"
	"sort"
)

// 1
func minNumber(nums1, nums2 []int) int {
	var mask1, mask2 uint
	for _, x := range nums1 {
		mask1 |= 1 << x
	}
	for _, x := range nums2 {
		mask2 |= 1 << x
	}
	if m := mask1 & mask2; m > 0 {
		return bits.TrailingZeros(m)
	}
	x, y := bits.TrailingZeros(mask1), bits.TrailingZeros(mask2)
	return min(x*10+y, y*10+x)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// 2
func maximumCostSubstring(s, chars string, vals []int) (ans int) {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}
	f := 0
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 3
func makeSubKSumEqual(arr []int, k int) (ans int64) {
	k = gcd(k, len(arr))
	g := make([][]int, k)
	for i, x := range arr {
		g[i%k] = append(g[i%k], x)
	}
	for _, b := range g {
		sort.Ints(b)
		for _, x := range b {
			ans += int64(abs(x - b[len(b)/2]))
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 4
func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	ans := math.MaxInt32
	dis := make([]int, n)                // dis[i] 表示从 start 到 i 的最短路长度
	for start := 0; start < n; start++ { // 枚举每个起点跑 BFS
		for j := range dis {
			dis[j] = -1
		}
		dis[start] = 0
		type pair struct{ x, fa int }
		q := []pair{{start, -1}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, fa := p.x, p.fa
			for _, y := range g[x] {
				if dis[y] < 0 { // 第一次遇到
					dis[y] = dis[x] + 1
					q = append(q, pair{y, x})
				} else if y != fa { // 第二次遇到
					ans = min(ans, dis[x]+dis[y]+1)
				}
			}
		}
	}
	if ans == math.MaxInt32 { // 无环图
		return -1
	}
	return ans
}
