package interview

// 面试题 04-01
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	g := make([][]int, n)
	vis := make(map[int]bool)

	for _, v := range graph {
		g[v[0]] = append(g[v[0]], v[1])
	}
	q := []int{start}
	vis[start] = true
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == target {
				return true
			}
			for _, to := range g[v] {
				if !vis[to] {
					q = append(q, to)
				}
			}
		}
	}
	return false
}
