package main

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	ans := make([]int, n)
	type pair struct {
		to, wt int
	}
	g := make([][]pair, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], pair{v[1], v[2]})
		g[v[1]] = append(g[v[1]], pair{v[0], v[2]})
	}

	for i, gt := range g {
		var cnt, total int
		var dfs func(fa, x, sum int)
		dfs = func(fa, x, sum int) {
			if sum%signalSpeed == 0 {
				cnt++
			}
			for _, v := range g[x] {
				if v.to == fa {
					continue
				}
				dfs(x, v.to, sum+v.wt)
			}
		}
		for _, v := range gt {
			cnt = 0
			dfs(i, v.to, v.wt)
			ans[i] += total * cnt
			total += cnt
		}
	}

	return ans
}
