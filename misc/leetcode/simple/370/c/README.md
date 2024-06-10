#### 题目

<p>有一棵 <code>n</code>&nbsp;个节点的无向树，节点编号为 <code>0</code>&nbsp;到 <code>n - 1</code>&nbsp;，根节点编号为 <code>0</code>&nbsp;。给你一个长度为 <code>n - 1</code>&nbsp;的二维整数数组&nbsp;<code>edges</code>&nbsp;表示这棵树，其中&nbsp;<code>edges[i] = [a<sub>i</sub>, b<sub>i</sub>]</code>&nbsp;表示树中节点&nbsp;<code>a<sub>i</sub></code>&nbsp;和&nbsp;<code>b<sub>i</sub></code>&nbsp;有一条边。</p>

<p>同时给你一个长度为 <code>n</code>&nbsp;下标从 <strong>0</strong>&nbsp;开始的整数数组&nbsp;<code>values</code>&nbsp;，其中&nbsp;<code>values[i]</code>&nbsp;表示第 <code>i</code>&nbsp;个节点的值。</p>

<p>一开始你的分数为 <code>0</code>&nbsp;，每次操作中，你将执行：</p>

<ul>
	<li>选择节点&nbsp;<code>i</code>&nbsp;。</li>
	<li>将&nbsp;<code>values[i]</code>&nbsp;加入你的分数。</li>
	<li>将&nbsp;<code>values[i]</code>&nbsp;变为&nbsp;<code>0</code>&nbsp;。</li>
</ul>

<p>如果从根节点出发，到任意叶子节点经过的路径上的节点值之和都不等于 0 ，那么我们称这棵树是 <strong>健康的</strong>&nbsp;。</p>

<p>你可以对这棵树执行任意次操作，但要求执行完所有操作以后树是&nbsp;<strong>健康的</strong>&nbsp;，请你返回你可以获得的 <strong>最大分数</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/10/11/graph-13-1.png" style="width: 515px; height: 443px;" /></p>

<pre>
<b>输入：</b>edges = [[0,1],[0,2],[0,3],[2,4],[4,5]], values = [5,2,5,2,1,1]
<b>输出：</b>11
<b>解释：</b>我们可以选择节点 1 ，2 ，3 ，4 和 5 。根节点的值是非 0 的。所以从根出发到任意叶子节点路径上节点值之和都不为 0 。所以树是健康的。你的得分之和为 values[1] + values[2] + values[3] + values[4] + values[5] = 11 。
11 是你对树执行任意次操作以后可以获得的最大得分之和。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/10/11/graph-14-2.png" style="width: 522px; height: 245px;" /></p>

<pre>
<b>输入：</b>edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]], values = [20,10,9,7,4,3,5]
<b>输出：</b>40
<b>解释：</b>我们选择节点 0 ，2 ，3 和 4 。
- 从 0 到 4 的节点值之和为 10 。
- 从 0 到 3 的节点值之和为 10 。
- 从 0 到 5 的节点值之和为 3 。
- 从 0 到 6 的节点值之和为 5 。
所以树是健康的。你的得分之和为 values[0] + values[2] + values[3] + values[4] = 40 。
40 是你对树执行任意次操作以后可以获得的最大得分之和。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>edges.length == n - 1</code></li>
	<li><code>edges[i].length == 2</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt; n</code></li>
	<li><code>values.length == n</code></li>
	<li><code>1 &lt;= values[i] &lt;= 10<sup>9</sup></code></li>
	<li>输入保证&nbsp;<code>edges</code>&nbsp;构成一棵合法的树。</li>
</ul>

#### 思路

正难则反，先把所有 $\textit{values}[i]$ 加到答案中，然后考虑哪些 $\textit{values}[i]$ 不能选（撤销，不加入答案）。  
设当前节点为 $x$，计算以 $x$ 为根的子树是健康时，失去的最小分数。那么答案就是 $\textit{values}$ 的元素和，减去「以 $0$ 为根的子树是健康时，**失去**的最小分数」。  
用「**选或不选**」分类讨论：
- 第一种情况：失去 $\textit{values}[x]$，也就是不加入答案，那么 $x$ 的所有子孙节点都可以加入答案，失去的最小分数就是 $\textit{values}[x]$。
- 第二种情况：$\textit{values}[x]$ 加入答案，问题变成「以 $y$ 为根的子树是健康时，失去的最小分数」，这里 $y$ 是 $x$ 的儿子。如果有多个儿子，累加失去的最小分数。

这两种情况取最小值。注意第一种情况是不会往下递归的，所以当我们递归到叶子的时候，叶子一定不能加入答案，此时直接返回 $\textit{values}[x]$。  
代码实现时，为了方便判断 $x$ 是否为叶子节点，可以假设还有一条 $0$ 到 $-1$ 的边，这样不会误把根节点 $0$ 当作叶子。

```go  
func maximumScoreAfterOperations(edges [][]int, values []int) int64 {
	g := make([][]int, len(values))
	g[0] = append(g[0], -1) // 避免误把根节点当作叶子
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	total := 0
	// dfs(x, fa) 计算以 x 为根的子树是健康时，失去的最小分数
	var dfs func(int, int) int
	dfs = func(x, fa int) int {
		total += values[x]
		if len(g[x]) == 1 { // x 是叶子
			return values[x]
		}
		loss := 0 // 第二种情况
		for _, y := range g[x] {
			if y != fa {
				loss += dfs(y, x) // 计算以 y 为根的子树是健康时，失去的最小分数
			}
		}
		return min(values[x], loss) // 两种情况取最小值
	}
	return int64(total - dfs(0, -1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{values}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$
