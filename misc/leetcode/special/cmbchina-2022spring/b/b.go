package main

func numFlowers(roads [][]int) (ans int) {

	n := len(roads)
	g := make([][]int, n+1)
	for _, r := range roads {
		from, to := r[0], r[1]
		g[from] = append(g[from], to)
		g[to] = append(g[to], from)
	}

	for i := range g {
		ans = max(ans, len(g[i]))
	}
	ans++
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
