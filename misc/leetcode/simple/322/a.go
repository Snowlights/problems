package simple_322

import (
	"math"
	"sort"
)

// 1
func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}
	for i := range sentence {
		if sentence[i] == ' ' && sentence[i-1] != sentence[i+1] {
			return false
		}
	}
	return true
}

// 2
func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	ans := 0
	tmp := skill[0] + skill[len(skill)-1]
	for i := 0; i < len(skill)/2; i++ {
		if val := skill[i] + skill[len(skill)-1-i]; val != tmp {
			return -1
		}
		ans += skill[i] * skill[len(skill)-1-i]
	}

	return int64(ans)
}

// 3
func minScore(n int, roads [][]int) int {
	type pair struct {
		to, val int
	}
	g := make(map[int][]*pair)
	for _, r := range roads {
		from, to, w := r[0], r[1], r[2]
		g[from] = append(g[from], &pair{
			to:  to,
			val: w,
		})
		g[to] = append(g[to], &pair{
			to:  from,
			val: w,
		})
	}

	ans := math.MaxInt64
	var dfs func(int)
	vis := make(map[int]bool)
	dfs = func(i int) {
		for _, v := range g[i] {
			ans = min(ans, v.val)
			if !vis[v.to] {
				vis[i] = true
				dfs(v.to)
			}
		}
	}
	dfs(1)
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 4
func magnificentSets(n int, edges [][]int) int {

	g := make(map[int][]int)
	for _, v := range edges {
		from, to := v[0]-1, v[1]-1
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	var bfs func(int) int
	bfs = func(i int) int {
		vis := make([]bool, n)
		q, depth := []int{i}, 0
		vis[i] = true
		for len(q) > 0 {
			tmp := q
			depth++
			q = nil
			for _, v := range tmp {
				for _, to := range g[v] {
					if !vis[to] {
						q = append(q, to)
						vis[to] = true
					}
				}
			}
		}
		return depth
	}

	colors := make([]int8, n)
	var nodes []int
	ans := 0
	var isBipartite func(int, int8) bool
	isBipartite = func(x int, c int8) bool { // 二分图判定
		nodes = append(nodes, x)
		colors[x] = c
		for _, y := range g[x] {
			if colors[y] == c || colors[y] == 0 && !isBipartite(y, -c) {
				return false
			}
		}
		return true
	}

	for i, c := range colors {
		if c == 0 {
			nodes = nil
			if !isBipartite(i, 1) { // 如果不是二分图（有奇环），则无法分组
				return -1
			}
			// 否则一定可以分组
			maxDepth := 0
			for _, x := range nodes { // 枚举连通块的每个点，作为起点 BFS，求最大深度
				maxDepth = max(maxDepth, bfs(x))
			}
			ans += maxDepth
		}
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
