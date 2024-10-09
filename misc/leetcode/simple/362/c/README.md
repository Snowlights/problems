#### 题目

<p>给你一个大小为 <code>3 * 3</code>&nbsp;，下标从 <strong>0</strong>&nbsp;开始的二维整数矩阵&nbsp;<code>grid</code>&nbsp;，分别表示每一个格子里石头的数目。网格图中总共恰好有&nbsp;<code>9</code>&nbsp;个石头，一个格子里可能会有 <strong>多个</strong>&nbsp;石头。</p>
<p>每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。</p>
<p>请你返回每个格子恰好有一个石头的 <strong>最少移动次数</strong>&nbsp;。</p>
<p>&nbsp;</p>
<p><strong class="example">示例 1：</strong></p>
<p><img alt src="https://assets.leetcode.com/uploads/2023/08/23/example1-3.svg" style="width: 401px; height: 281px;" /></p>
<pre>
<b>输入：</b>grid = [[1,1,0],[1,1,1],[1,2,1]]
<b>输出：</b>3
<b>解释：</b>让每个格子都有一个石头的一个操作序列为：
1 - 将一个石头从格子 (2,1) 移动到 (2,2) 。
2 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
3 - 将一个石头从格子 (1,2) 移动到 (0,2) 。
总共需要 3 次操作让每个格子都有一个石头。
让每个格子都有一个石头的最少操作次数为 3 。
</pre>
<p><strong class="example">示例 2：</strong></p>
<p><img alt src="https://assets.leetcode.com/uploads/2023/08/23/example2-2.svg" style="width: 401px; height: 281px;" /></p>
<pre>
<b>输入：</b>grid = [[1,3,0],[1,0,0],[1,0,3]]
<b>输出：</b>4
<b>解释：</b>让每个格子都有一个石头的一个操作序列为：
1 - 将一个石头从格子 (0,1) 移动到 (0,2) 。
2 - 将一个石头从格子 (0,1) 移动到 (1,1) 。
3 - 将一个石头从格子 (2,2) 移动到 (1,2) 。
4 - 将一个石头从格子 (2,2) 移动到 (2,1) 。
总共需要 4 次操作让每个格子都有一个石头。
让每个格子都有一个石头的最少操作次数为 4 。
</pre>
<p>&nbsp;</p>
<p><strong>提示：</strong></p>
<ul>
<li><code>grid.length == grid[i].length == 3</code></li>
<li><code>0 &lt;= grid[i][j] &lt;= 9</code></li>
<li><code>grid</code>&nbsp;中元素之和为&nbsp;<code>9</code> 。</li>
</ul>

#### 思路

一个通用的做法是**最小费用最大流**。即使是 $10\times 10$ 的网格也可以做。  
建图规则如下： 
- 从每个大于 $1$ 的格子向每个等于 $0$ 的格子连边，容量为 $1$，费用为两个格子之间的曼哈顿距离。
- 从超级源点向每个大于 $1$ 的格子连边，容量为格子的值减一（即移走的石子数），费用为 $0$。
- 从每个等于 $0$ 的格子向超级汇点连边，容量 $1$（即移入的石子数），费用为 $0$。
  
答案为最大流时，对应的最小费用。

```go 
func minimumMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	src := m * n   // 超级源点
	dst := src + 1 // 超级汇点
	type edge struct{ to, rid, cap, cost int }
	g := make([][]edge, m*n+2)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], edge{to, len(g[to]), cap, cost})
		g[to] = append(g[to], edge{from, len(g[from]) - 1, 0, -cost})
	}
	for x, row := range grid {
		for y, v := range row {
			if v > 1 {
				addEdge(src, x*n+y, v-1, 0)
				for i, r := range grid {
					for j, w := range r {
						if w == 0 {
							addEdge(x*n+y, i*n+j, 1, abs(x-i)+abs(y-j))
						}
					}
				}
			} else if v == 0 {
				addEdge(x*n+y, dst, 1, 0)
			}
		}
	}

	// 下面是最小费用最大流模板
	const inf int = 1e9
	dist := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	spfa := func() bool {
		for i := range dist {
			dist[i] = 1e9
		}
		dist[src] = 0
		inQ := make([]bool, len(g))
		inQ[src] = true
		q := []int{src}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				if newD := dist[v] + e.cost; newD < dist[w] {
					dist[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						q = append(q, w)
						inQ[w] = true
					}
				}
			}
		}
		return dist[dst] < inf
	}
	ek := func() (maxFlow, minCost int) {
		for spfa() {
			// 沿 st-end 的最短路尽量增广
			minF := inf
			for v := dst; v != src; {
				p := fa[v]
				if c := g[p.v][p.i].cap; c < minF {
					minF = c
				}
				v = p.v
			}
			for v := dst; v != src; {
				p := fa[v]
				e := &g[p.v][p.i]
				e.cap -= minF
				g[v][e.rid].cap += minF
				v = p.v
			}
			maxFlow += minF
			minCost += dist[dst] * minF
		}
		return
	}
	_, cost := ek()
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。