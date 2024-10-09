package _00_700

import (
	"fmt"
	"sort"
	"strings"
)

// 678
func checkValidString(s string) bool {

	l, c := 0, 0
	for _, v := range s {
		switch v {
		case '(':
			l++
		case ')':
			if l > 0 {
				l--
				continue
			}
			if c > 0 {
				c--
			} else {
				return false
			}
		case '*':
			c++
		}
	}
	if l > c {
		return false
	}

	r, c := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case ')':
			r++
		case '(':
			if l > 0 {
				r--
				continue
			}
			if c > 0 {
				c--
			} else {
				return false
			}
		case '*':
			c++
		}
	}
	if r > c {
		return false
	}
	return true
}

// 680
func validPalindrome(str string) bool {
	s, e := 0, len(str)-1
	for s < e {
		if str[s] != str[e] {
			break
		}
		s++
		e--
	}

	isPartition := func(str string) bool {
		s, e := 0, len(str)-1
		for s <= e {
			if str[s] != str[e] {
				return false
			}
			s++
			e--
		}
		return true
	}

	if s != e {
		return isPartition(str[:s]+str[s+1:]) || isPartition(str[:e]+str[e+1:])
	}
	return true
}

// 684
func findRedundantConnection(edges [][]int) []int {

	indegree, g := make(map[int]int), make(map[int][]int)
	h := make(map[int]bool)
	for _, edge := range edges {
		indegree[edge[0]]++
		indegree[edge[1]]++
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
		h[edge[0]] = true
		h[edge[1]] = true
	}

	bfs := func(from, abandon int) int {
		vis := make(map[int]bool)
		q := []int{from}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, to := range g[v] {
					if v == from && to == abandon {
						continue
					}
					if !vis[to] {
						q = append(q, to)
						vis[to] = true
					}
				}
			}
		}
		return len(vis)
	}

	for i := len(edges) - 1; i >= 0; i-- {
		v := edges[i]
		if a, b := indegree[v[0]]-1, indegree[v[1]]-1; a == 0 || b == 0 {
			continue
		}
		if bfs(v[0], v[1]) == len(h) {
			return v
		}
	}

	return nil
}

//func findRedundantConnection(edges [][]int) []int {
//
//	fa := make([]int, len(edges)+1)
//	for i := range fa {
//		fa[i] = i
//	}
//	var find func(int) int
//	find = func(from int) int {
//		if fa[from] != from {
//			fa[from] = find(fa[from])
//		}
//		return fa[from]
//	}
//	union := func(from, to int) bool {
//		f, t := find(from), find(to)
//		if f == t {
//			return true
//		}
//		fa[f] = t
//		return false
//	}
//
//	for _, e := range edges {
//		if union(e[0], e[1]) {
//			return e
//		}
//	}
//
//	return nil
//}

// 685
func findRedundantDirectedConnection(edges [][]int) []int {

	// [[2,1],[3,1],[4,2],[1,4]]
	fa, parent := make([]int, len(edges)+1), make([]int, len(edges)+1)
	for i := range fa {
		fa[i] = i
		parent[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			fa[from] = find(fa[from])
		}
		return fa[from]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}
	var conflictEdge, cycleEdge []int
	for _, e := range edges {
		from, to := e[0], e[1]
		if parent[to] != to { // to 有两个父节点
			conflictEdge = e
		} else {
			parent[to] = from
			if find(from) == find(to) { // from 和 to 已连接
				cycleEdge = e
			} else {
				merge(from, to)
			}
		}
	}

	if conflictEdge == nil {
		return cycleEdge
	}

	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

// 686
func repeatedStringMatch(a string, b string) int {
	m, n, exist := len(a), len(b), make([]bool, 26)
	for _, ch := range a {
		exist[ch-'a'] = true
	}
	for _, ch := range b {
		if !(exist[ch-'a']) {
			return -1
		}
	}

	ret := n / m
	str := strings.Repeat(a, ret)
	for i := 0; i < 3; i++ {
		if strings.Contains(str, b) {
			return ret + i
		}
		str += a
	}
	return -1
}

// 688
func knightProbability(n int, k int, row int, column int) float64 {
	dir := [][]int{
		{-2, -1}, {-2, 1}, {2, -1}, {2, 1},
		{-1, -2}, {-1, 2}, {1, -2}, {1, 2},
	}

	dp := make([][][]float64, k+1)
	for i := range dp {
		dp[i] = make([][]float64, n)
		for j := range dp[i] {
			dp[i][j] = make([]float64, n)
			for k := range dp[i][j] {
				if i == 0 {
					dp[i][j][k] = 1
				} else {
					for _, d := range dir {
						if x, y := j+d[0], k+d[1]; 0 <= x && x < n && 0 <= y && y < n {
							dp[i][j][k] += dp[i-1][x][y] / 8
						}
					}
				}
			}
		}
	}
	return dp[k][row][column]
}

// 690
type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	h, vis := make(map[int]*Employee, len(employees)), make(map[int]bool)
	for _, v := range employees {
		h[v.Id] = v
	}
	ans := 0
	q := []*Employee{h[id]}
	vis[id] = true
	for len(q) > 0 {
		fmt.Println(q)
		tmp := q
		q = nil
		for _, e := range tmp {
			ans += e.Importance
			for _, sub := range e.Subordinates {
				if !vis[sub] {
					q = append(q, h[sub])
					vis[sub] = true
				}
			}
		}
	}

	return ans
}

// 692
func topKFrequent(words []string, k int) []string {

	type node struct {
		word  string
		count int
	}
	nodeList, nodeMap := make([]*node, 0, len(words)), map[string]*node{}
	for _, word := range words {
		n, ok := nodeMap[word]
		if !ok {
			n = &node{word: word}
			nodeList = append(nodeList, n)
			nodeMap[word] = n
		}
		n.count++
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].count > nodeList[j].count || (nodeList[i].count == nodeList[j].count && nodeList[j].word > nodeList[i].word)
	})

	ans := make([]string, 0, k)
	for _, v := range nodeList[:k] {
		ans = append(ans, v.word)
	}
	return ans
}

// 695
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		grid[i][j] = 0
		tmp := 1
		for _, d := range dir {
			x, y := i+d[0], j+d[1]
			if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
				tmp += dfs(x, y)
			}
		}
		return tmp
	}
	for i, v := range grid {
		for j, vv := range v {
			if vv == 1 {
				ans = max(ans, dfs(i, j))
			}
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
