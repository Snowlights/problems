### 题目

<p>给你一棵 <code>n</code> 个节点的 <strong>无向</strong> 树，节点编号为 <code>0</code> 到 <code>n - 1</code> ，树的根节点在节点 <code>0</code> 处。同时给你一个长度为 <code>n - 1</code> 的二维整数数组 <code>edges</code> ，其中 <code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>]</code> 表示树中节点 <code>a<sub>i</sub></code> 和 <code>b<sub>i</sub></code> 之间有一条边。</p>

<p>给你一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>cost</code> ，其中 <code>cost[i]</code> 是第 <code>i</code> 个节点的 <b>开销</b> 。</p>

<p>你需要在树中每个节点都放置金币，在节点 <code>i</code> 处的金币数目计算方法如下：</p>

<ul>
	<li>如果节点 <code>i</code> 对应的子树中的节点数目小于 <code>3</code> ，那么放 <code>1</code> 个金币。</li>
	<li>否则，计算节点 <code>i</code> 对应的子树内 <code>3</code> 个不同节点的开销乘积的 <strong>最大值</strong> ，并在节点 <code>i</code> 处放置对应数目的金币。如果最大乘积是 <b>负数</b> ，那么放置 <code>0</code> 个金币。</li>
</ul>

<p>请你返回一个长度为 <code>n</code> 的数组<em> </em><code>coin</code> ，<code>coin[i]</code>是节点 <code>i</code> 处的金币数目。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012641.png" style="width: 600px; height: 233px;" /></p>

<pre>
<b>输入：</b>edges = [[0,1],[0,2],[0,3],[0,4],[0,5]], cost = [1,2,3,4,5,6]
<b>输出：</b>[120,1,1,1,1,1]
<b>解释：</b>在节点 0 处放置 6 * 5 * 4 = 120 个金币。所有其他节点都是叶子节点，子树中只有 1 个节点，所以其他每个节点都放 1 个金币。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012614.png" style="width: 800px; height: 374px;" /></p>

<pre>
<b>输入：</b>edges = [[0,1],[0,2],[1,3],[1,4],[1,5],[2,6],[2,7],[2,8]], cost = [1,4,2,3,5,7,8,-4,2]
<b>输出：</b>[280,140,32,1,1,1,1,1,1]
<b>解释：</b>每个节点放置的金币数分别为：
- 节点 0 处放置 8 * 7 * 5 = 280 个金币。
- 节点 1 处放置 7 * 5 * 4 = 140 个金币。
- 节点 2 处放置 8 * 2 * 2 = 32 个金币。
- 其他节点都是叶子节点，子树内节点数目为 1 ，所以其他每个节点都放 1 个金币。
</pre>

<p><strong class="example">示例 3：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/09/screenshot-2023-11-10-012513.png" style="width: 300px; height: 277px;" /></p>

<pre>
<b>输入：</b>edges = [[0,1],[0,2]], cost = [1,2,-2]
<b>输出：</b>[0,1,1]
<b>解释：</b>节点 1 和 2 都是叶子节点，子树内节点数目为 1 ，各放置 1 个金币。节点 0 处唯一的开销乘积是 2 * 1 * -2 = -4 。所以在节点 0 处放置 0 个金币。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= n <= 2 * 10<sup>4</sup></code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 <= a<sub>i</sub>, b<sub>i</sub> < n</code></li>
	<li><code>cost.length == n</code></li>
	<li><code>1 <= |cost[i]| <= 10<sup>4</sup></code></li>
	<li><code>edges</code> 一定是一棵合法的树。</li>
</ul>

### 思路

首先来看如下问题：
给你一个长度 $\ge 3$ 的数组，返回**三数之积**的最大值，如果最大值是负数，返回 $0$。

- 如果只考虑正数，显然最大的三个数的乘积是最大的。
- 如果把负数也考虑进来，那么最小的两个负数乘最大的正数也可能是最大值。

这两种情况再和 $0$ 取最大值，即为返回值。
由上诉讨论可知，只需要知道 $5$ 个数，就能算出三数之积的最大值。
每棵子树只需要返回它最小的 $2$ 个 $\textit{cost}$ 值和最大的 $3$ 个 $\textit{cost}$ 值就行，
其余数字可以舍弃。对于一棵子树，把它的所有儿子子树的返回值与当前节点的 $\textit{cost}$ 值排序后，
按照三数之积的方法，就得到了当前节点的答案。

```go [sol]
func placedCoins(edges [][]int, cost []int) []int64 {
	g := make(map[int][]int)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}

	ans := make([]int64, len(cost))
	for i := range ans {
		ans[i] = 1
	}
	var dfs func(x, fa int) []int
	dfs = func(x, fa int) []int {
		a := []int{cost[x]}
		for _, to := range g[x] {
			if to == fa {
				continue
			}
			a = append(a, dfs(to, x)...)
		}
		sort.Ints(a)
		n := len(a)
		if len(a) < 3 {
			ans[x] = 1
		} else {
			ans[x] = int64(max(a[0] * a[1] * a[n-1], a[n-3] * a[n-2] * a[n-1], 0))
		}
		if len(a) > 5 {
			a = append(a[:2], a[n-3:]...)
		}
		return a
	}
	dfs(0, -1)
	return ans
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上，如果手动维护最小的两个数和最大的三个数可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。
