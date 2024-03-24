### 题目

<p>Alice 有一棵 <code>n</code> 个节点的树，节点编号为 <code>0</code> 到 <code>n - 1</code> 。树用一个长度为 <code>n - 1</code> 的二维整数数组 <code>edges</code> 表示，其中 <code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> ，表示树中节点 <code>a<sub>i</sub></code> 和 <code>b<sub>i</sub></code> 之间有一条边。</p>

<p>Alice 想要 Bob 找到这棵树的根。她允许 Bob 对这棵树进行若干次 <strong>猜测</strong> 。每一次猜测，Bob 做如下事情：</p>

<ul>
	<li>选择两个 <strong>不相等</strong> 的整数 <code>u</code> 和 <code>v</code> ，且树中必须存在边 <code>[u, v]</code> 。</li>
	<li>Bob 猜测树中 <code>u</code> 是 <code>v</code> 的 <strong>父节点</strong> 。</li>
</ul>

<p>Bob 的猜测用二维整数数组 <code>guesses</code> 表示，其中 <code>guesses[j] = [u<sub>j</sub>, v<sub>j</sub>]</code> 表示 Bob 猜 <code>u<sub>j</sub></code> 是 <code>v<sub>j</sub></code> 的父节点。</p>

<p>Alice 非常懒，她不想逐个回答 Bob 的猜测，只告诉 Bob 这些猜测里面 <strong>至少</strong> 有 <code>k</code> 个猜测的结果为 <code>true</code> 。</p>

<p>给你二维整数数组 <code>edges</code> ，Bob 的所有猜测和整数 <code>k</code> ，请你返回可能成为树根的 <strong>节点数目</strong> 。如果没有这样的树，则返回 <code>0</code>。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2022/12/19/ex-1.png" style="width: 727px; height: 250px;"/></p>

<pre><b>输入：</b>edges = [[0,1],[1,2],[1,3],[4,2]], guesses = [[1,3],[0,1],[1,0],[2,4]], k = 3
<b>输出：</b>3
<b>解释：</b>
根为节点 0 ，正确的猜测为 [1,3], [0,1], [2,4]
根为节点 1 ，正确的猜测为 [1,3], [1,0], [2,4]
根为节点 2 ，正确的猜测为 [1,3], [1,0], [2,4]
根为节点 3 ，正确的猜测为 [1,0], [2,4]
根为节点 4 ，正确的猜测为 [1,3], [1,0]
节点 0 ，1 或 2 为根时，可以得到 3 个正确的猜测。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2022/12/19/ex-2.png" style="width: 600px; height: 303px;"/></p>

<pre><b>输入：</b>edges = [[0,1],[1,2],[2,3],[3,4]], guesses = [[1,0],[3,4],[2,1],[3,2]], k = 1
<b>输出：</b>5
<b>解释：</b>
根为节点 0 ，正确的猜测为 [3,4]
根为节点 1 ，正确的猜测为 [1,0], [3,4]
根为节点 2 ，正确的猜测为 [1,0], [2,1], [3,4]
根为节点 3 ，正确的猜测为 [1,0], [2,1], [3,2], [3,4]
根为节点 4 ，正确的猜测为 [1,0], [2,1], [3,2]
任何节点为根，都至少有 1 个正确的猜测。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>edges.length == n - 1</code></li>
	<li><code>2 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= guesses.length <= 10<sup>5</sup></code></li>
	<li><code>0 <= a<sub>i</sub>, b<sub>i</sub>, u<sub>j</sub>, v<sub>j</sub> <= n - 1</code></li>
	<li><code>a<sub>i</sub> != b<sub>i</sub></code></li>
	<li><code>u<sub>j</sub> != v<sub>j</sub></code></li>
	<li><code>edges</code> 表示一棵有效的树。</li>
	<li><code>guesses[j]</code> 是树中的一条边。</li>
	<li><code>guesses</code> 是唯一的。</li>
	<li><code>0 <= k <= guesses.length</code></li>
</ul>

### 思路

如果只求以 $0$ 为根时的猜对次数 $\textit{cnt}_0$，那么把 $\textit{guesses}$ 转成哈希表，DFS 一次这棵树就可以算出来。如果要枚举以每个点为根时的猜对次数，暴力做法就太慢了，怎么优化呢？注意到，如果节点 $x$ 和 $y$ 之间有边，那么从「以 $x$ 为根的树」变成「以 $y$ 为根的树」，就只有 $[x,y]$ 和 $[y,x]$ 这两个猜测的正确性变了，其余猜测的正确性不变。因此，从 $0$ 出发，再次 DFS 这棵树，从节点 $x$ 递归到节点 $y$ 时：

- 如果有猜测 $[x,y]$，那么猜对次数减一；
- 如果有猜测 $[y,x]$，那么猜对次数加一。\
- DFS 的同时，统计猜对次数 $\ge k$ 的节点个数，即为答案。

这个套路叫做「\ DP」。

```go  
func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	type pair struct {
		x, y int
	}
	g, h := make([][]int, len(edges)+1), make(map[pair]int)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	for _, e := range guesses {
		h[pair{e[0], e[1]}] = 1
	}

	cnt := 0
	var dfs func(fa, x int)
	dfs = func(fa, x int) {
		for _, to := range g[x] {
			if to != fa {
				if h[pair{x, to}] == 1 {
					cnt++
				}
				dfs(x, to)
			}
		}
	}
	dfs(-1, 0)
	var reDfs func(fa, x, cnt int)
	reDfs = func(fa, x, cnt int) {
		if cnt >= k {
			ans++
		}
		for _, to := range g[x] {
			if to != fa {
				reDfs(x, to, cnt-h[pair{x, to}]+h[pair{to, x}])
			}
		}
	}
	reDfs(-1, 0, cnt)
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $\textit{edges}$ 的长度加一，$m$ 为 $\textit{guesses}$ 的长度。
- 空间复杂度：$O(n+m)$。
