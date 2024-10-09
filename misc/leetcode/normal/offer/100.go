package offer

// 106
func isBipartite(graph [][]int) bool {
	n := len(graph)
	mark := make([]int, n)
	for i := range mark {
		mark[i] = -1
	}
	for i := 0; i < n; i++ {
		if mark[i] != -1 {
			continue
		}
		queue := []int{i}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range graph[u] {
				if mark[v] != -1 {
					if mark[u]^mark[v] != 1 {
						return false
					}
				} else {
					mark[v] = mark[u] ^ 1
					queue = append(queue, v)
				}
			}
		}
	}
	return true
}

// offer 110
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	g := make(map[int][]int)
	for i, gra := range graph {
		g[i] = append(g[i], gra...)
	}
	ans := make([][]int, 0)
	q := [][]int{{0}}
	for len(q) > 0 {
		val := q[0]
		q = q[1:]
		if val[len(val)-1] == n-1 {
			ans = append(ans, val)
			continue
		}
		for _, v := range g[val[len(val)-1]] {
			q = append(q, append(append([]int{}, val...), v))
		}
	}
	return ans
}

// offer 111
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	type edge struct {
		to  string
		val float64
	}

	g, h := make(map[string][]edge), make(map[string]bool)
	for i, e := range equations {
		from, to := e[0], e[1]
		h[from], h[to] = true, true
		g[from] = append(g[from], edge{
			to:  to,
			val: values[i],
		})
		g[to] = append(g[to], edge{
			to:  from,
			val: 1 / values[i],
		})
	}

	ans := make([]float64, len(queries))
	for i, query := range queries {
		if !h[query[0]] || !h[query[1]] {
			ans[i] = -1
			continue
		}
		q, vis := []edge{{
			to:  query[0],
			val: 1,
		}}, make(map[string]bool)
		vis[query[0]] = true
		for len(q) > 0 {
			tmp := q
			q = nil

			for _, v := range tmp {
				if v.to == query[1] {
					ans[i] = v.val
					q = nil
					break
				}
				for _, to := range g[v.to] {
					if !vis[to.to] {
						vis[to.to] = true
						q = append(q, edge{
							to:  to.to,
							val: v.val * to.val,
						})
					}
				}
			}
		}
		if !vis[query[1]] {
			ans[i] = -1
		}
	}

	return ans
}

// offer 113
func findOrder(numCourses int, prerequisites [][]int) []int {
	g, indegree := make(map[int][]int), make(map[int]int)
	for _, course := range prerequisites {
		from, to := course[1], course[0]
		g[from] = append(g[from], to)
		indegree[to]++
	}
	q, ans, vis := []int{}, []int{}, make(map[int]bool)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
			vis[i] = true
			ans = append(ans, i)
		}
	}

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, to := range g[v] {
				if indegree[to]--; indegree[to] == 0 {
					q = append(q, to)
					vis[to] = true
					ans = append(ans, to)
				}
			}
		}
	}
	if len(vis) != numCourses {
		return nil
	}

	return ans
}

// offer 116
func findCircleNum(isConnected [][]int) int {

	n, group := len(isConnected), len(isConnected)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			from = find(fa[from])
		}
		return fa[from]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
		group--
	}

	for from, v := range isConnected {
		for to, val := range v {
			if val == 1 {
				merge(from, to)
			}
		}
	}

	return group
}

// offer 118
func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			from = find(fa[from])
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

	for _, v := range edges {
		from, to := find(v[0]), find(v[1])
		if from == to {
			return v
		}
		merge(from, to)
	}

	return nil
}
