package main

func minEdgeReversals(n int, edges [][]int) []int {
	type pair struct {
		to int
		f  bool
	}
	g := make([][]pair, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], pair{
			to: e[1],
			f:  false,
		})
		g[e[1]] = append(g[e[1]], pair{
			to: e[0],
			f:  true,
		})
	}

	ans := make([]int, n)
	var f func(int, int)  // int64
	f = func(v, fa int) { // sum 表示以 0 为根时的子树 v 中的节点到 v 的距离之和
		for _, w := range g[v] {
			if w.to != fa {
				if w.f {
					ans[0]++
				}
				f(w.to, v)
			}
		}
		return
	}
	f(0, -1)

	var reRoot func(v, fa int)
	reRoot = func(v, fa int) {
		for _, w := range g[v] {
			if w.to != fa {
				ans[w.to] = ans[v]
				if w.f {
					ans[w.to]--
				} else {
					ans[w.to]++
				}
				reRoot(w.to, v)
			}
		}
	}
	reRoot(0, -1)
	return ans
}
