package shunfeng

import (
	"sort"
	"strings"
)

// 1
func hasCycle(graph string) bool {
	g, gMap := make(map[string][]string), make(map[string]bool)

	for _, v := range strings.Split(graph, ",") {
		p := strings.Split(v, "->")
		g[p[0]] = append(g[p[0]], p[1])
		gMap[p[0]] = true
		gMap[p[1]] = true
	}

	for k, _ := range gMap {
		vis, q := make(map[string]bool), []string{}
		q = append(q, k)
		turn := 0
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				if v == k && turn > 0 {
					return true
				}
				for _, vv := range g[v] {
					if vis[vv] {
						continue
					}
					q = append(q, vv)
					vis[vv] = true
				}
			}
			turn++
		}
	}

	return false
}

// 2
func minRemainingSpace(N []int, V int) int {
	sort.Ints(N)
	dp := make([]int, V+1)
	dp[0] = 1
	for _, v := range N {
		for i := V; i >= v; i-- {
			dp[i] |= dp[i-v]
		}
	}
	for i := V; i >= 0; i-- {
		if dp[i] == 1 {
			return V - i
		}
	}
	return V
}

// 3
func findMaxCI(nums []int) int {

	pre, day, ans := nums[0], 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > pre {
			day++
			ans = max(ans, day)
		} else {
			day = 1
		}
		pre = nums[i]
	}
	return ans
}

// 4
func isPointInPolygon(x float64, y float64, c []float64) bool {
	x1, y1, x2, y2 := c[0], c[1], c[2], c[3]
	t := (x-x1)*(y2-y1) - (x2-x1)*(y-y1)
	tt := t
	for i := 5; i < len(c); i += 2 {
		x1, y1, x2, y2 = c[i-3], c[i-2], c[i-1], c[i]
		u := (x-x1)*(y2-y1) - (x2-x1)*(y-y1)
		if u*t < 0 {
			return false
		}
		t = u
	}

	return t*tt > 0
}

// 5
func isCompliance(distance [][]int, n int) bool {

	type node struct {
		n, l int
	}
	g := make(map[int][]*node)
	for i, d := range distance {
		for j, v := range d {
			if i == j {
				continue
			}
			g[i] = append(g[i], &node{
				n: j,
				l: v,
			})
		}
	}

	var dfs func(int)
	vis, ans := make(map[int]bool), 0
	dfs = func(i int) {
		vis[i] = true
		for _, node := range g[i] {
			if !vis[node.n] && node.l <= 2 {
				dfs(node.n)
			}
		}
	}
	for i := 0; i < n; i++ {
		if vis[i] {
			continue
		}
		ans++
		dfs(i)
	}
	return ans <= n
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
