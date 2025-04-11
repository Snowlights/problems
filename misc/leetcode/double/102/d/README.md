### 题目  

<p>给你一个有 <code>n</code> 个节点的 <strong>有向带权</strong> 图，节点编号为 <code>0</code> 到 <code>n - 1</code> 。图中的初始边用数组 <code>edges</code> 表示，其中 <code>edges[i] = [from<sub>i</sub>, to<sub>i</sub>, edgeCost<sub>i</sub>]</code> 表示从 <code>from<sub>i</sub></code> 到 <code>to<sub>i</sub></code> 有一条代价为 <code>edgeCost<sub>i</sub></code> 的边。</p>

<p>请你实现一个 <code>Graph</code> 类：</p>

<ul>
	<li><code>Graph(int n, int[][] edges)</code> 初始化图有 <code>n</code> 个节点，并输入初始边。</li>
	<li><code>addEdge(int[] edge)</code> 向边集中添加一条边，其中<strong> </strong><code>edge = [from, to, edgeCost]</code> 。数据保证添加这条边之前对应的两个节点之间没有有向边。</li>
	<li><code>int shortestPath(int node1, int node2)</code> 返回从节点 <code>node1</code> 到 <code>node2</code> 的路径<strong> 最小</strong> 代价。如果路径不存在，返回 <code>-1</code> 。一条路径的代价是路径中所有边代价之和。</li>
</ul>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/01/11/graph3drawio-2.png" style="width: 621px; height: 191px;"/></p>

<pre><strong>输入：</strong>
[&#34;Graph&#34;, &#34;shortestPath&#34;, &#34;shortestPath&#34;, &#34;addEdge&#34;, &#34;shortestPath&#34;]
[[4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]], [3, 2], [0, 3], [[1, 3, 4]], [0, 3]]
<b>输出：</b>
[null, 6, -1, null, 6]

<strong>解释：</strong>
Graph g = new Graph(4, [[0, 2, 5], [0, 1, 2], [1, 2, 1], [3, 0, 3]]);
g.shortestPath(3, 2); // 返回 6 。从 3 到 2 的最短路径如第一幅图所示：3 -&gt; 0 -&gt; 1 -&gt; 2 ，总代价为 3 + 2 + 1 = 6 。
g.shortestPath(0, 3); // 返回 -1 。没有从 0 到 3 的路径。
g.addEdge([1, 3, 4]); // 添加一条节点 1 到节点 3 的边，得到第二幅图。
g.shortestPath(0, 3); // 返回 6 。从 0 到 3 的最短路径为 0 -&gt; 1 -&gt; 3 ，总代价为 2 + 4 = 6 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 100</code></li>
	<li><code>0 &lt;= edges.length &lt;= n * (n - 1)</code></li>
	<li><code>edges[i].length == edge.length == 3</code></li>
	<li><code>0 &lt;= from<sub>i</sub>, to<sub>i</sub>, from, to, node1, node2 &lt;= n - 1</code></li>
	<li><code>1 &lt;= edgeCost<sub>i</sub>, edgeCost &lt;= 10<sup>6</sup></code></li>
	<li>图中任何时候都不会有重边和自环。</li>
	<li>调用 <code>addEdge</code> 至多 <code>100</code> 次。</li>
	<li>调用 <code>shortestPath</code> 至多 <code>100</code> 次。</li>
</ul>
 
### 思路  

## 方法一：Dijkstra

定义 $\textit{start}$ 为起点，$\textit{dis}[i]$ 表示从 $\textit{start}$ 到 $i$ 的最短路的长度。
初始时 $\textit{dis}[\textit{start}]=0$，其余 $\textit{dis}[i]$ 为 $\infty$。  
首先，从 $\textit{start}$ 出发，更新邻居的最短路。  
下一步，寻找除去 $\textit{start}$ 的 $\textit{dis}$ 的最小值，设这个点为 $x$，那么 $\textit{dis}[x]$ 就已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。  
证明：由于除去起点的其余 $\textit{dis}[i]$ 都不低于 $\textit{dis}[x]$，且图中边权都非负，那么从另外一个点 $y$ 去更新 $\textit{dis}[x]$ 时，是无法让 $\textit{dis}[x]$ 变得更小的（因为 $\textit{dis}[x]$ 是当前最小），因此 $\textit{dis}[x]$ 已经是从 $\textit{start}$ 到 $x$ 的最短路的长度了。  
由于在寻找 $\textit{dis}$ 的最小值时，需要排除在前面的循环中找到的 $x$（因为已经更新 $x$ 到其它点的最短路了，无需反复更新），可以用一个 $\textit{vis}$ 数组标记这些 $x$。  
以上，通过**数学归纳法**，可以证明每次找到的未被标记的 $\textit{dis}$ 的最小值就是最短路。  
由于输入的图最坏是**稠密图**，所以采用邻接矩阵实现。

```go 
const inf = math.MaxInt / 2 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	g := make([][]int, n) // 邻接矩阵
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			g[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
		}
	}
	for _, e := range edges {
		g[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边）
	}
	return g
}

func (g Graph) AddEdge(e []int) {
	g[e[0]][e[1]] = e[2] // 添加一条边（输入保证这条边之前不存在）
}

// 朴素 Dijkstra 算法
func (g Graph) ShortestPath(start, end int) int {
	n := len(g)
	dis := make([]int, n) // 从 start 出发，到各个点的最短路，如果不存在则为无穷大
	for i := range dis {
		dis[i] = inf
	}
	dis[start] = 0
	vis := make([]bool, n)
	for {
		// 找到当前最短路，去更新它的邻居的最短路，
		// 根据数学归纳法，dis[x] 一定是最短路长度
		x := -1
		for i, b := range vis {
			if !b && (x < 0 || dis[i] < dis[x]) {
				x = i
			}
		}
		if x < 0 || dis[x] == inf { // 所有从 start 能到达的点都被更新了
			return -1
		}
		if x == end { // 找到终点，提前退出
			return dis[x]
		}
		vis[x] = true // 标记，在后续的循环中无需反复更新 x 到其余点的最短路长度
		for y, w := range g[x] {
			dis[y] = min(dis[y], dis[x]+w) // 更新最短路长度
		}
	}
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析  

- 时间复杂度：$O(qn^2)$，其中 $q$ 为 $\texttt{shortestPath}$ 的调用次数。每次求最短路的时间复杂度为 $O(n^2)$，在本题的输入下，这比堆的实现要快。
- 空间复杂度：$O(n^2)$。


## 方法二：Floyd

Floyd 本质是动态规划。由于这个动态规划的状态定义不是很好想出来，所以我就直接描述算法了：  
定义 $d[k][i][j]$ 表示从 $i$ 到 $j$ 的最短路长度，并且从 $i$ 到 $j$ 的路径上的中间节点（不含 $i$ 和 $j$）的编号至多为 $k$。  
分类讨论：
- 如果 $i$ 到 $j$ 的路径上的节点编号没有 $k$，那么按照定义 $d[k][i][j] = d[k-1][i][j]$。
- 如果 $i$ 到 $j$ 的路径上的节点编号有 $k$，那么可以视作先从 $i$ 到 $k$，再从 $k$ 到 $j$。由于 $i$ 到 $k$ 和 $k$ 到 $j$ 的中间节点都没有 $k$，所以有 $d[k][i][j] = d[k-1][i][k] + d[k-1][k][j]$。
取最小值，得 

$$
d[k][i][j] =\min(d[k-1][i][j], d[k-1][i][k] + d[k-1][k][j])
$$

初始值 $d[0][i][j]$ 为原图中 $i$ 到 $j$ 的边长，如果不存在则为 $\infty$。最终 $i$ 到 $j$ 的最短路长度为 $d[k-1][i][j]$。  
代码实现时，第一个维度可以优化掉，即

$$
d[i][j] =\min(d[[i][j], d[i][k] + d[k][j])
$$
> 对于 $\texttt{addEdge}$ 操作，记 $x=\textit{from},y=\textit{to}$ 如果 $\textit{edgeCode} \ge d[x][y]$，则无法更新任何点对的最短路。
> 否则枚举所有 $d[i][j]$，尝试看看能否更新成更小，即 $i---x-y---j$ 是否更短：
> 
> $$
> d[i][j] = \min(d[i][j], d[i][x] + \textit{edgeCode} + d[y][j])
> $$
> 
> 由于当 $i=x$ 或 $j=y$ 时，需要用到 $d[i][i]$ 这样的值，所以初始化的时候，$d[i][i]$ 要置为 $0$。

```go  
const inf = math.MaxInt / 3 // 防止更新最短路时加法溢出

type Graph [][]int

func Constructor(n int, edges [][]int) Graph {
	d := make([][]int, n) // 邻接矩阵
	for i := range d {
		d[i] = make([]int, n)
		for j := range d[i] {
			if j != i {
				d[i][j] = inf // 初始化为无穷大，表示 i 到 j 没有边
			}
		}
	}
	for _, e := range edges {
		d[e[0]][e[1]] = e[2] // 添加一条边（输入保证没有重边和自环）
	}
	for k := range d {
		for i := range d {
			for j := range d {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}
	return d
}

func (d Graph) AddEdge(e []int) {
	x, y, w := e[0], e[1], e[2]
	if w >= d[x][y] { // 无需更新
		return
	}
	for i := range d {
		for j := range d {
			d[i][j] = min(d[i][j], d[i][x]+w+d[y][j])
		}
	}
}

func (d Graph) ShortestPath(start, end int) int {
	ans := d[start][end]
	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n^3 + qn^2)$，其中 $q$ 为 $\texttt{addEdge}$ 的调用次数。
- 空间复杂度：$O(n^2)$。