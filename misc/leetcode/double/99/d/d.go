package main

func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	type pair struct {
		x, y int
	}
	g, h := make([][]int, len(edges)+1), make(map[pair]int)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	for _, e := range guesses {
		h[pair{e[0], e[1]}] = 1
	}

	cnt := 0
	var dfs func(fa, x int)
	dfs = func(fa, x int) {
		for _, to := range g[x] {
			if to != fa {
				if h[pair{x, to}] == 1 {
					cnt++
				}
				dfs(x, to)
			}
		}
	}
	dfs(-1, 0)
	var reDfs func(fa, x, cnt int)
	reDfs = func(fa, x, cnt int) {
		if cnt >= k {
			ans++
		}
		for _, to := range g[x] {
			if to != fa {
				reDfs(x, to, cnt-h[pair{x, to}]+h[pair{to, x}])
			}
		}
	}
	reDfs(-1, 0, cnt)
	return
}
