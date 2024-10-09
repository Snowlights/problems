package _300_1400

import (
	"bytes"
	"math"
	"sort"
	"strings"
)

// 1300
func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 0; i < len(arr); i++ {
		prefix[i+1] = prefix[i] + arr[i]
	}
	ans, tmp := 0, math.MaxInt32
	for i := 0; i < 1e5; i++ {
		idx := sort.SearchInts(arr, i)
		val := prefix[idx] + (n-idx)*i
		if diff := abs(target - val); tmp > diff {
			tmp = diff
			ans = i
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 1311
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	g := make(map[int][]int)
	for i, v := range friends {
		for _, vv := range v {
			g[i] = append(g[i], vv)
			g[vv] = append(g[vv], i)
		}
	}

	idx, q, vis := 1, []int{id}, make(map[int]bool)
	vis[id] = true
	for idx <= level {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, to := range g[v] {
				if !vis[to] {
					q = append(q, to)
					vis[to] = true
				}
			}
		}
		idx++
	}
	type pair struct {
		watch string
		count int
	}
	ans, ansM := make([]*pair, 0), make(map[string]*pair)
	for _, v := range q {
		for _, vv := range watchedVideos[v] {
			_, ok := ansM[vv]
			if !ok {
				ansM[vv] = &pair{
					watch: vv,
					count: 0,
				}
				ans = append(ans, ansM[vv])
			}
			ansM[vv].count++
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].count < ans[j].count || (ans[i].count == ans[j].count && ans[j].watch > ans[i].watch)
	})
	res := make([]string, 0, len(ans))
	for _, v := range ans {
		res = append(res, v.watch)
	}
	return res
}

// 1314
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			pre[i][j] = mat[i-1][j-1] + pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1]
		}
	}

	ans := make([][]int, m)
	for i := range mat {
		ans[i] = make([]int, n)
		for j := range mat[i] {
			left, right := max(0, j-k), min(n, j+k+1)
			top, button := max(0, i-k), min(m, i+k+1)
			ans[i][j] = pre[button][right] - pre[button][left] - pre[top][right] + pre[top][left]
		}
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 1319
func makeConnected(n int, connections [][]int) (ans int) {
	if len(connections) < n-1 {
		return -1
	}

	graph := make([][]int, n)
	for _, c := range connections {
		v, w := c[0], c[1]
		graph[v] = append(graph[v], w)
		graph[w] = append(graph[w], v)
	}

	vis := make([]bool, n)
	var dfs func(int)
	dfs = func(from int) {
		vis[from] = true
		for _, to := range graph[from] {
			if !vis[to] {
				dfs(to)
			}
		}
	}
	for i, v := range vis {
		if !v {
			ans++
			dfs(i)
		}
	}
	return ans - 1
}

// 1324
func printVertically(s string) []string {
	parts, ml := strings.Split(s, " "), 0
	for i := 0; i < len(parts); i++ {
		ml = max(ml, len(parts[i]))
	}
	ans := make([][]byte, ml)
	for i := range ans {
		ans[i] = bytes.Repeat([]byte(" "), len(parts))
	}
	for i := range parts {
		for j := range parts[i] {
			ans[j][i] = parts[i][j]
		}
	}

	res := make([]string, 0)
	for _, v := range ans {
		res = append(res, strings.TrimRight(string(v), " "))
	}
	return res
}
