#### 题目  

<p>现有一棵无向、无根的树，树中有 <code>n</code> 个节点，按从 <code>0</code> 到 <code>n - 1</code> 编号。给你一个整数 <code>n</code> 和一个长度为 <code>n - 1</code> 的二维整数数组 <code>edges</code> ，其中 <code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> 表示树中节点 <code>a<sub>i</sub></code> 和 <code>b<sub>i</sub></code> 之间存在一条边。</p>

<p>每个节点都关联一个价格。给你一个整数数组 <code>price</code> ，其中 <code>price[i]</code> 是第 <code>i</code> 个节点的价格。</p>

<p>给定路径的 <strong>价格总和</strong> 是该路径上所有节点的价格之和。</p>

<p>另给你一个二维整数数组 <code>trips</code> ，其中 <code>trips[i] = [start<sub>i</sub>, end<sub>i</sub>]</code> 表示您从节点 <code>start<sub>i</sub></code> 开始第 <code>i</code> 次旅行，并通过任何你喜欢的路径前往节点 <code>end<sub>i</sub></code> 。</p>

<p>在执行第一次旅行之前，你可以选择一些 <strong>非相邻节点</strong> 并将价格减半。</p>

<p>返回执行所有旅行的最小价格总和。</p>

<p> </p>

<p><strong>示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/03/16/diagram2.png" style="width: 541px; height: 181px;"/>
<pre><strong>输入：</strong>n = 4, edges = [[0,1],[1,2],[1,3]], price = [2,2,10,6], trips = [[0,3],[2,1],[2,3]]
<strong>输出：</strong>23
<strong>解释：
</strong>上图表示将节点 2 视为根之后的树结构。第一个图表示初始树，第二个图表示选择节点 0 、2 和 3 并使其价格减半后的树。
第 1 次旅行，选择路径 [0,1,3] 。路径的价格总和为 1 + 2 + 3 = 6 。
第 2 次旅行，选择路径 [2,1] 。路径的价格总和为 2 + 5 = 7 。
第 3 次旅行，选择路径 [2,1,3] 。路径的价格总和为 5 + 2 + 3 = 10 。
所有旅行的价格总和为 6 + 7 + 10 = 23 。可以证明，23 是可以实现的最小答案。</pre>

<p><strong>示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/03/16/diagram3.png" style="width: 456px; height: 111px;"/>
<pre><strong>输入：</strong>n = 2, edges = [[0,1]], price = [2,2], trips = [[0,0]]
<strong>输出：</strong>1
<strong>解释：</strong>
上图表示将节点 0 视为根之后的树结构。第一个图表示初始树，第二个图表示选择节点 0 并使其价格减半后的树。 
第 1 次旅行，选择路径 [0] 。路径的价格总和为 1 。 
所有旅行的价格总和为 1 。可以证明，1 是可以实现的最小答案。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 50</code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt;= n - 1</code></li>
	<li><code>edges</code> 表示一棵有效的树</li>
	<li><code>price.length == n</code></li>
	<li><code>price[i]</code> 是一个偶数</li>
	<li><code>1 &lt;= price[i] &lt;= 1000</code></li>
	<li><code>1 &lt;= trips.length &lt;= 100</code></li>
	<li><code>0 &lt;= start<sub>i</sub>, end<sub>i</sub> &lt;= n - 1</code></li>
</ul>
 
#### 思路  

## 方法一：暴力 DFS 每条路径

#### 提示 1

如果知道每个点总共被经过多少次，就可以仿照 [337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/) 计算答案了（下面会细说）。  
注意到数据范围比较小，可以对每个 $\textit{trips}[i]$ 都跑一遍 DFS，把从 $\textit{start}$ 到 $\textit{end}$ 的路径上的点 $x$ 的经过次数 $\textit{cnt}[x]$ 都加一。
这一技巧在之前的双周赛中出现过，见 [2467. 树上最大得分和路径](https://leetcode.cn/problems/most-profitable-path-in-a-tree/)。

#### 提示 2

既然知道了每个点会被经过多少次，把 $\textit{price}[i]$ 更新成 $\textit{price}[i]\cdot \textit{cnt}[i]$，问题就变成计算减半后的 $\textit{price}[i]$ 之和的最小值。  
随便选一个节点出发 DFS，比如节点 $0$。对于节点 $x$ 及其儿子 $y$，分类讨论：
- 如果 $\textit{price}[x]$ 不变，那么 $\textit{price}[y]$ 可以减半，也可以不变，取这两种情况的最小值；
- 如果 $\textit{price}[x]$ 减半，那么 $\textit{price}[y]$ 只能不变。

因此子树 $x$ 需要返回两个值：
- $\textit{price}[x]$ 不变时的子树 $x$ 的最小价值总和；
- $\textit{price}[x]$ 减半时的子树 $x$ 的最小价值总和。

答案就是根节点不变/减半的最小值。

```go 
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	cnt := make([]int, n)
	for _, t := range trips {
		end := t[1]
		var dfs func(int, int) bool
		dfs = func(x, fa int) bool {
			if x == end { // 到达终点（注意树只有唯一的一条简单路径）
				cnt[x]++ // 统计从 start 到 end 的路径上的点经过了多少次
				return true // 找到终点
			}
			for _, y := range g[x] {
				if y != fa && dfs(y, x) {
					cnt[x]++ // 统计从 start 到 end 的路径上的点经过了多少次
					return true
				}
			}
			return false // 未找到终点
		}
		dfs(t[0], -1)
	}

	// 类似 337. 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/
	var dfs func(int, int) (int, int)
	dfs = func(x, fa int) (int, int) {
		notHalve := price[x] * cnt[x] // x 不变
		halve := notHalve / 2 // x 减半
		for _, y := range g[x] {
			if y != fa {
				nh, h := dfs(y, x) // 计算 y 不变/减半的最小价值总和
				notHalve += min(nh, h) // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
				halve += nh // x 减半，那么 y 只能不变
			}
		}
		return notHalve, halve
	}
	nh, h := dfs(0, -1)
	return min(nh, h)
}

func min(a, b int) int { if a > b { return b }; return a }
```

#### 复杂度分析  

- 时间复杂度：$O(nm)$，其中 $m$ 为 $\textit{trips}$ 的长度。
- 空间复杂度：$O(n)$。


## 方法二：Tarjan 离线 LCA + 树上差分

核心思路：利用**树上差分**打标记，再通过一次 DFS 算出 $\textit{cnt}$ 值。  
从 $x=\textit{start}$ 到 $y=\textit{end}$ 的路径可以视作从 $x$ 向上到某个点「拐弯」，再向下到达 $y$。 （拐弯的点也可能就是 $x$ 或 $y$）  
这个拐弯的点就是 $x$ 和 $y$ 的 $\textit{lca}$（最近公共祖先）。  
把路径视作 $x-\textit{lca}'-\textit{lca}-y$，其中 $\textit{lca}'$ 是 $\textit{lca}$ 的儿子。由于更新的是点，
拆分成 $x-\textit{lca}'$ 和 $y-\textit{lca}$。那么自底向上更新差分 $\textit{diff}$ 值：
- 对于 $x-\textit{lca}'$，更新 $\textit{diff}[x]$ 加一，$\textit{diff}[\textit{lca}]$ 减一；
- 对于 $y-\textit{lca}$，更新 $\textit{diff}[y]$ 加一，$\textit{diff}[\textit{father}[\textit{lca}]]$ 减一，这里 $\textit{father}[i]$ 表示 $i$ 的父节点。

最近公共祖先，用 **Tarjan 离线算法**计算，解释见代码注释。  
然后 DFS，在递归的「归」的过程中累加 $\textit{diff}$，计算出 $\textit{cnt}$ 值。

```go  
func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建树
	}

	qs := make([][]int, n)
	for _, t := range trips {
		x, y := t[0], t[1]
		qs[x] = append(qs[x], y) // 路径端点分组
		if x != y {
			qs[y] = append(qs[y], x)
		}
	}

	// 并查集模板
	pa := make([]int, n)
	for i := range pa {
		pa[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if pa[x] != x {
			pa[x] = find(pa[x])
		}
		return pa[x]
	}

	diff := make([]int, n)
	father := make([]int, n)
	color := make([]int8, n)
	var tarjan func(int, int)
	tarjan = func(x, fa int) {
	father[x] = fa
		color[x] = 1 // 递归中
		for _, y := range g[x] {
			if color[y] == 0 { // 未递归
				tarjan(y, x)
				pa[y] = x // 相当于把 y 的子树节点全部 merge 到 x
			}
		}
		for _, y := range qs[x] {
			// color[y] == 2 意味着 y 所在子树已经遍历完
			// 也就意味着 y 已经 merge 到它和 x 的 lca 上了
			if y == x || color[y] == 2 { // 从 y 向上到达 lca 然后拐弯向下到达 x
				diff[x]++
				diff[y]++
				lca := find(y)
				diff[lca]--
				if f := father[lca]; f >= 0 {
					diff[f]--
				}
			}
		}
		color[x] = 2 // 递归结束
	}
	tarjan(0, -1)

	var dfs func(int, int) (int, int, int)
	dfs = func(x, fa int) (notHalve, halve, cnt int) {
		cnt = diff[x]
		for _, y := range g[x] {
			if y != fa {
				nh, h, c := dfs(y, x)  // 计算 y 不变/减半的最小价值总和
				notHalve += min(nh, h) // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
				halve += nh            // x 减半，那么 y 只能不变
				cnt += c               // 自底向上累加差分值
			}
		}
		notHalve += price[x] * cnt  // x 不变
		halve += price[x] * cnt / 2 // x 减半
		return
	}
	nh, h, _ := dfs(0, -1)
	return min(nh, h)
}

func min(a, b int) int { if a > b { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$O(n+m\alpha)$，其中 $m$ 为 $\textit{trips}$ 的长度，$\alpha$ 为并查集的常数，可视作 $O(1)$。
- 空间复杂度：$O(n+m)