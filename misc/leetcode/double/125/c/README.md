### 题目

<p>给你一棵无根带权树，树中总共有 <code>n</code> 个节点，分别表示 <code>n</code> 个服务器，服务器从 <code>0</code> 到 <code>n - 1</code> 编号。同时给你一个数组 <code>edges</code> ，其中 <code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>, weight<sub>i</sub>]</code> 表示节点 <code>a<sub>i</sub></code> 和 <code>b<sub>i</sub></code> 之间有一条双向边，边的权值为 <code>weight<sub>i</sub></code> 。再给你一个整数 <code>signalSpeed</code> 。</p>

<p>如果两个服务器 <code>a</code> ，<code>b</code> 和 <code>c</code> 满足以下条件，那么我们称服务器 <code>a</code> 和 <code>b</code> 是通过服务器 <code>c</code> <strong>可连接的</strong> ：</p>

<ul>
	<li><code>a < b</code> ，<code>a != c</code> 且 <code>b != c</code> 。</li>
	<li>从 <code>c</code> 到 <code>a</code> 的距离是可以被 <code>signalSpeed</code> 整除的。</li>
	<li>从 <code>c</code> 到 <code>b</code> 的距离是可以被 <code>signalSpeed</code> 整除的。</li>
	<li>从 <code>c</code> 到 <code>b</code> 的路径与从 <code>c</code> 到 <code>a</code> 的路径没有任何公共边。</li>
</ul>

<p>请你返回一个长度为 <code>n</code> 的整数数组 <code>count</code> ，其中 <code>count[i]</code> 表示通过服务器 <code>i</code> <strong>可连接</strong> 的服务器对的 <strong>数目</strong> 。</p>

<p> </p>

<p><b>示例 1：</b></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/01/21/example22.png" style="width: 438px; height: 243px; padding: 10px; background: #fff; border-radius: .5rem;" /></p>

<pre>
<b>输入：</b>edges = [[0,1,1],[1,2,5],[2,3,13],[3,4,9],[4,5,2]], signalSpeed = 1
<b>输出：</b>[0,4,6,6,4,0]
<b>解释：</b>由于 signalSpeed 等于 1 ，count[c] 等于所有从 c 开始且没有公共边的路径对数目。
在输入图中，count[c] 等于服务器 c 左边服务器数目乘以右边服务器数目。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/01/21/example11.png" style="width: 495px; height: 484px; padding: 10px; background: #fff; border-radius: .5rem;" /></p>

<pre>
<b>输入：</b>edges = [[0,6,3],[6,5,3],[0,3,1],[3,2,7],[3,1,6],[3,4,2]], signalSpeed = 3
<b>输出：</b>[2,0,0,0,0,0,2]
<b>解释：</b>通过服务器 0 ，有 2 个可连接服务器对(4, 5) 和 (4, 6) 。
通过服务器 6 ，有 2 个可连接服务器对 (4, 5) 和 (0, 5) 。
所有服务器对都必须通过服务器 0 或 6 才可连接，所以其他服务器对应的可连接服务器对数目都为 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= n <= 1000</code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 3</code></li>
	<li><code>0 <= a<sub>i</sub>, b<sub>i</sub> < n</code></li>
	<li><code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>, weight<sub>i</sub>]</code><!-- notionvc: a2623897-1bb1-4c07-84b6-917ffdcd83ec --></li>
	<li><code>1 <= weight<sub>i</sub> <= 10<sup>6</sup></code></li>
	<li><code>1 <= signalSpeed <= 10<sup>6</sup></code></li>
	<li>输入保证 <code>edges</code> 构成一棵合法的树。</li>
</ul>

### 思路

把 $c$ 作为树根计算。

![b125C.png](https://pic.leetcode.cn/1709427910-nOCIAc-b125C.png)

```go [sol]
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	n := len(edges) + 1
	ans := make([]int, n)
	type pair struct {
		to, wt int
	}
	g := make([][]pair, n)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], pair{v[1], v[2]})
		g[v[1]] = append(g[v[1]], pair{v[0], v[2]})
	}

	for i, gt := range g {
		var cnt, total int
		var dfs func(fa, x, sum int)
		dfs = func(fa, x, sum int) {
			if sum % signalSpeed == 0 {
				cnt++
			}
			for _, v := range g[x] {
				if v.to == fa {
					continue
				}
				dfs(x, v.to, sum + v.wt)
			}
		}
		for _, v := range gt {
			cnt = 0
			dfs(i, v.to, v.wt)
			ans[i] += total * cnt
			total += cnt
		}
	}

	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{edges}$ 的长度加一。
- 空间复杂度：$\mathcal{O}(n)$。
