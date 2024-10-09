package _200_1300

import (
	"sort"
)

// 1202
func smallestStringWithSwaps(s string, pairs [][]int) string {
	n := len(s)
	fa := make([]int, n)
	rank := make([]int, n)
	for i := range fa {
		fa[i] = i
		rank[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	union := func(x, y int) {
		fx, fy := find(x), find(y)
		if fx == fy {
			return
		}
		if rank[fx] < rank[fy] {
			fx, fy = fy, fx
		}
		rank[fx] += rank[fy]
		fa[fy] = fx
	}

	for _, p := range pairs {
		union(p[0], p[1])
	}

	groups := map[int][]byte{}
	for i := range s {
		f := find(i)
		groups[f] = append(groups[f], s[i])
	}

	for _, bytes := range groups {
		sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	}

	ans := make([]byte, n)
	for i := range ans {
		f := find(i)
		ans[i] = groups[f][0]
		groups[f] = groups[f][1:]
	}
	return string(ans)
}

func smallestStringWithSwapsV2(s string, pairs [][]int) string {

	g := make(map[int][]int)
	for _, v := range pairs {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}
	hMap := make([][]int, 0)
	vis := make(map[int]bool)
	for i := range s {
		if !vis[i] {
			vis[i] = true
			q := []int{i}
			ans := []int{i}
			for len(q) > 0 {
				tmp := q
				q = nil
				for _, v := range tmp {
					for _, to := range g[v] {
						if !vis[to] {
							q = append(q, to)
							vis[to] = true
							ans = append(ans, to)
						}
					}
				}
			}
			hMap = append(hMap, ans)
		}
	}

	ans := make([]byte, len(s))
	for _, v := range hMap {
		b := make([]byte, 0, len(v))
		for _, v := range v {
			b = append(b, s[v])
		}
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		sort.Ints(v)
		for i, vv := range v {
			ans[vv] = b[i]
		}
	}

	return string(ans)
}

// 1222
func queensAttacktheKing(queens [][]int, king []int) [][]int {
	dir := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	ans := [][]int{}
	queenMap := make(map[[2]int]bool)
	for _, v := range queens {
		queenMap[[2]int{v[0], v[1]}] = true
	}
	for _, d := range dir {
		s, e := king[0], king[1]
		for 0 <= s && s < 8 && 0 <= e && e < 8 {
			if queenMap[[2]int{s, e}] {
				ans = append(ans, []int{s, e})
				break
			}
			s, e = s+d[0], e+d[1]
		}
	}

	return ans
}
