package main

const mx int = 1e5 + 1

var np = [mx]bool{1: true} // 质数=false 非质数=true
func init() {
	for i := 2; i < mx; i++ {
		if !np[i] {
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func countPaths(n int, edges [][]int) (ans int64) {
	fa := make([]int, n+1) // n+1
	size := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	sizeF := func(x int) int {
		return size[find(x)]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
		size[to] += size[from]
	}
	g := make([][]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		if np[x] && np[y] {
			merge(x, y)
		}
	}

	for x := 1; x <= n; x++ {
		if np[x] { // 跳过非质数
			continue
		}
		sum := 1
		for _, to := range g[x] { // 质数 x 把这棵树分成了若干个连通块
			// 求这个相邻的非质数
			if np[to] {
				// 原来的路径端点数量为 sum，新增路径数量为 size(v) * sum
				ans += int64(sizeF(to) * sum)
				sum += sizeF(to)
			}
		}
	}
	return
}
