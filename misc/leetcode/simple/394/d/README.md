#### 题目

<p>给你一个 <code>n</code>&nbsp;个节点的无向带权图，节点编号为 <code>0</code>&nbsp;到 <code>n - 1</code>&nbsp;。图中总共有 <code>m</code>&nbsp;条边，用二维数组&nbsp;<code>edges</code>&nbsp;表示，其中&nbsp;<code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>, w<sub>i</sub>]</code>&nbsp;表示节点 <code>a<sub>i</sub></code> 和&nbsp;<code>b<sub>i</sub></code>&nbsp;之间有一条边权为&nbsp;<code>w<sub>i</sub></code>&nbsp;的边。</p>

<p>对于节点 <code>0</code>&nbsp;为出发点，节点 <code>n - 1</code>&nbsp;为结束点的所有最短路，你需要返回一个长度为 <code>m</code>&nbsp;的 <strong>boolean</strong>&nbsp;数组&nbsp;<code>answer</code>&nbsp;，如果&nbsp;<code>edges[i]</code>&nbsp;<strong>至少</strong>&nbsp;在其中一条最短路上，那么&nbsp;<code>answer[i]</code>&nbsp;为&nbsp;<code>true</code>&nbsp;，否则&nbsp;<code>answer[i]</code>&nbsp;为&nbsp;<code>false</code>&nbsp;。</p>

<p>请你返回数组&nbsp;<code>answer</code>&nbsp;。</p>

<p><b>注意</b>，图可能不连通。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/03/05/graph35drawio-1.png" style="height: 129px; width: 250px;" /></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 6, edges = [[0,1,4],[0,2,1],[1,3,2],[1,4,3],[1,5,1],[2,3,1],[3,5,3],[4,5,2]]</span></p>

<p><span class="example-io"><b>输出：</b>[true,true,true,false,true,true,true,false]</span></p>

<p><strong>解释：</strong></p>

<p>以下为节点 0 出发到达节点 5 的 <strong>所有</strong>&nbsp;最短路：</p>

<ul>
	<li>路径&nbsp;<code>0 -&gt; 1 -&gt; 5</code>&nbsp;：边权和为&nbsp;<code>4 + 1 = 5</code>&nbsp;。</li>
	<li>路径&nbsp;<code>0 -&gt; 2 -&gt; 3 -&gt; 5</code>&nbsp;：边权和为&nbsp;<code>1 + 1 + 3 = 5</code>&nbsp;。</li>
	<li>路径&nbsp;<code>0 -&gt; 2 -&gt; 3 -&gt; 1 -&gt; 5</code>&nbsp;：边权和为&nbsp;<code>1 + 1 + 2 + 1 = 5</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/03/05/graphhhh.png" style="width: 185px; height: 136px;" /></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 4, edges = [[2,0,1],[0,1,1],[0,3,4],[3,2,2]]</span></p>

<p><span class="example-io"><b>输出：</b>[true,false,false,true]</span></p>

<p><strong>解释：</strong></p>

<p>只有一条从节点 0 出发到达节点 3 的最短路&nbsp;<code>0 -&gt; 2 -&gt; 3</code>&nbsp;，边权和为&nbsp;<code>1 + 2 = 3</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>m == edges.length</code></li>
	<li><code>1 &lt;= m &lt;= min(5 * 10<sup>4</sup>, n * (n - 1) / 2)</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt; n</code></li>
	<li><code>a<sub>i</sub> != b<sub>i</sub></code></li>
	<li><code>1 &lt;= w<sub>i</sub> &lt;= 10<sup>5</sup></code></li>
	<li>图中没有重边。</li>
</ul>

#### 思路

首先用 Dijkstra 算法（堆优化版本）计算出起点 $0$ 到所有节点的最短路长度 $\textit{dis}$。
如果 $\textit{dis}[n-1]=\infty$，说明无法从起点 $0$ 到终点 $n-1$，答案全为 $\textit{false}$。
否则，我们可以从终点 $n-1$ 出发，倒着 DFS 或 BFS，设当前在点 $y$，邻居为 $x$，边权为 $w$，如果满足

$$
\textit{dis}[x] + w = \textit{dis}[y]
$$

则说明 $x\textit{-}y$ 这条边在从 $0$ 到 $n-1$ 的最短路上。

#### 答疑

**问**：为什么在求出最短路后，不能从起点 $0$ 出发去寻找在最短路上的边？
**答**：从起点 $0$ 出发，当发现 $\textit{dis}[x] + w = \textit{dis}[y]$ 时，这仅仅意味着 $x\textit{-}y$ 这条边在从 $0$ 出发的最短路上，
但这条最短路不一定通向终点 $n-1$（比如图是一棵树，这条边不通往 $n-1$）。而从终点 $n-1$ 出发倒着寻找，就能保证符合等式的边在通向终点的最短路上。

```
func findAnswer(n int, edges [][]int) []bool {
	ans := make([]bool, len(edges))

	const (
		inf int = 1e18
	)
	type edge struct{ i, to, wt int }
	dijkstra := func(g [][]edge, start int) []int {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0
		// 虽然可以用 dist 来判断是否需要 relax，但是对于一些变形题，用 vis 是最稳的
		vis := make([]bool, n)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = -1
		}
		h := vdHeap{{start, 0}}
		for len(h) > 0 {
			p := h.pop()
			v := p.v
			if vis[v] { // p.dis > dist[v]
				continue
			}
			vis[v] = true
			for _, e := range g[v] {
				w := e.to
				if newD := dist[v] + int(e.wt); newD < dist[w] {
					dist[w] = newD
					fa[w] = v
					h.push(vdPair{w, dist[w]})
				}
			}
		}
		return dist
	}
	g := make([][]edge, n)
	for i, v := range edges {
		g[v[0]] = append(g[v[0]], edge{
			i:  i,
			to: v[1],
			wt: v[2],
		})
		g[v[1]] = append(g[v[1]], edge{
			i:  i,
			to: v[0],
			wt: v[2],
		})
	}
	dis := dijkstra(g, 0)

	if dis[n-1] == inf {
		return ans
	}

	var dfs func(x, fa int)
	dfs = func(x, fa int) {
		for _, to := range g[x] {
			if to.to == fa {
				continue
			}
			if dis[x]-to.wt != dis[to.to] {
				continue
			}
			ans[to.i] = true
			dfs(to.to, x)
		}
	}
	dfs(n-1, n)
	return ans
}

type vdPair struct {
	v   int
	dis int
}
type vdHeap []vdPair

func (h vdHeap) Len() int              { return len(h) }
func (h vdHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h vdHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *vdHeap) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *vdHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *vdHeap) push(v vdPair)        { heap.Push(h, v) }
func (h *vdHeap) pop() vdPair          { return heap.Pop(h).(vdPair) }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
