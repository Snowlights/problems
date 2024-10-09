package _900_2000

// 1971
func validPath(n int, edges [][]int, source int, destination int) bool {
	g := make(map[int][]int)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}
	vis := make(map[int]bool)
	q := []int{source}
	vis[source] = true
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == destination {
				return true
			}
			for _, to := range g[v] {
				if !vis[to] {
					q = append(q, to)
					vis[to] = true
				}
			}
		}
	}
	return false
}
