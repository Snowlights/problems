### 题目  

<p>现有一个含 <code>n</code> 个顶点的 <strong>双向</strong> 图，每个顶点按从 <code>0</code> 到 <code>n - 1</code> 标记。图中的边由二维整数数组 <code>edges</code> 表示，其中 <code>edges[i] = [u<sub>i</sub>, v<sub>i</sub>]</code> 表示顶点 <code>u<sub>i</sub></code> 和 <code>v<sub>i</sub></code> 之间存在一条边。每对顶点最多通过一条边连接，并且不存在与自身相连的顶点。</p>

<p>返回图中 <strong>最短</strong> 环的长度。如果不存在环，则返回 <code>-1</code> 。</p>

<p><strong>环</strong> 是指以同一节点开始和结束，并且路径中的每条边仅使用一次。</p>

<p> </p>

<p><strong>示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/01/04/cropped.png" style="width: 387px; height: 331px;"/>
<pre><strong>输入：</strong>n = 7, edges = [[0,1],[1,2],[2,0],[3,4],[4,5],[5,6],[6,3]]
<strong>输出：</strong>3
<strong>解释：</strong>长度最小的循环是：0 -&gt; 1 -&gt; 2 -&gt; 0 
</pre>

<p><strong>示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/01/04/croppedagin.png" style="width: 307px; height: 307px;"/>
<pre><strong>输入：</strong>n = 4, edges = [[0,1],[0,2]]
<strong>输出：</strong>-1
<strong>解释：</strong>图中不存在循环
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 1000</code></li>
	<li><code>1 &lt;= edges.length &lt;= 1000</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 &lt;= u<sub>i</sub>, v<sub>i</sub> &lt; n</code></li>
	<li><code>u<sub>i</sub> != v<sub>i</sub></code></li>
	<li>不存在重复的边</li>
</ul>
 
### 思路  

![b101_t4_cut.png](https://pic.leetcode.cn/1680363054-UnoCDM-b101_t4_cut.png)

### 答疑

**问**：为什么说发现一个已经入队的点，就说明有环？
**答**：这说明到同一个点有两条不同的路径，这两条路径组成了一个环。

```go 
func findShortestCycle(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	ans := math.MaxInt32
	dis := make([]int, n)                // dis[i] 表示从 start 到 i 的最短路长度
	for start := 0; start < n; start++ { // 枚举每个起点跑 BFS
		for j := range dis {
			dis[j] = -1
		}
		dis[start] = 0
		type pair struct{ x, fa int }
		q := []pair{{start, -1}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, fa := p.x, p.fa
			for _, y := range g[x] {
				if dis[y] < 0 { // 第一次遇到
					dis[y] = dis[x] + 1
					q = append(q, pair{y, x})
				} else if y != fa { // 第二次遇到
					ans = min(ans, dis[x]+dis[y]+1)
				}
			}
		}
	}
	if ans == math.MaxInt32 { // 无环图
		return -1
	}
	return ans
}
```

### 复杂度分析  

- 时间复杂度：$O(nm)$，其中 $m$ 为 $\textit{edges}$ 的长度。每次 BFS 需要 $O(m)$ 的时间。
- 空间复杂度：$O(n+m)$。