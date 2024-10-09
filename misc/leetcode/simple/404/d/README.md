#### 题目

<p>给你两棵 <strong>无向</strong>&nbsp;树，分别有&nbsp;<code>n</code> 和&nbsp;<code>m</code>&nbsp;个节点，节点编号分别为&nbsp;<code>0</code>&nbsp;到&nbsp;<code>n - 1</code>&nbsp;和&nbsp;<code>0</code>&nbsp;到&nbsp;<code>m - 1</code>&nbsp;。给你两个二维整数数组&nbsp;<code>edges1</code> 和&nbsp;<code>edges2</code>&nbsp;，长度分别为&nbsp;<code>n - 1</code> 和&nbsp;<code>m - 1</code>&nbsp;，其中&nbsp;<code>edges1[i] = [a<sub>i</sub>, b<sub>i</sub>]</code>&nbsp;表示在第一棵树中节点&nbsp;<code>a<sub>i</sub></code> 和&nbsp;<code>b<sub>i</sub></code>&nbsp;之间有一条边，<code>edges2[i] = [u<sub>i</sub>, v<sub>i</sub>]</code>&nbsp;表示在第二棵树中节点&nbsp;<code>u<sub>i</sub></code> 和&nbsp;<code>v<sub>i</sub></code>&nbsp;之间有一条边。</p>

<p>你必须在第一棵树和第二棵树中分别选一个节点，并用一条边连接它们。</p>

<p>请你返回添加边后得到的树中，<strong>最小直径</strong>&nbsp;为多少。</p>

<p>一棵树的 <strong>直径</strong>&nbsp;指的是树中任意两个节点之间的最长路径长度。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b><img alt="" src="https://assets.leetcode.com/uploads/2024/04/22/example11-transformed.png" style="width: 1000px; height: 494px;" /></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>edges1 = [[0,1],[0,2],[0,3]], edges2 = [[0,1]]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<p>将第一棵树中的节点 0 与第二棵树中的任意节点连接，得到一棵直径为 3 得树。</p>
</div>

<p><strong class="example">示例 2：<img alt="" src="https://assets.leetcode.com/uploads/2024/04/22/example211.png" style="width: 1000px; height: 492px;" /></strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>edges1 = [[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]], edges2 = [[0,1],[0,2],[0,3],[2,4],[2,5],[3,6],[2,7]]</span></p>

<p><span class="example-io"><b>输出：</b>5</span></p>

<p><strong>解释：</strong></p>

<p>将第一棵树中的节点 0 和第二棵树中的节点 0 连接，可以得到一棵直径为 5 的树。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n, m &lt;= 10<sup>5</sup></code></li>
	<li><code>edges1.length == n - 1</code></li>
	<li><code>edges2.length == m - 1</code></li>
	<li><code>edges1[i].length == edges2[i].length == 2</code></li>
	<li><code>edges1[i] = [a<sub>i</sub>, b<sub>i</sub>]</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt; n</code></li>
	<li><code>edges2[i] = [u<sub>i</sub>, v<sub>i</sub>]</code></li>
	<li><code>0 &lt;= u<sub>i</sub>, v<sub>i</sub> &lt; m</code></li>
	<li>输入保证&nbsp;<code>edges1</code> 和&nbsp;<code>edges2</code>&nbsp;分别表示一棵合法的树。</li>
</ul>

#### 思路

**前置知识**：[树形 DP：树的直径【基础算法精讲 23】](https://www.bilibili.com/video/BV17o4y187h1/)

设 $d_1,d_2$ 分别为两棵树的直径长度。答案有三种情况：

- 第一棵树的直径特别长。那么连边后，新树的直径仍然为第一棵树的直径，答案为 $d_1$。
- 第二棵树的直径特别长。那么连边后，新树的直径仍然为第二棵树的直径，答案为 $d_2$。
- 新树的直径经过添加的边。假设我们连接了第一棵树的节点 $x_1$ 与第二棵树的节点 $x_2$，那么新树的直径相当于以下三部分之和：
  - $x_1$ 到第一棵树最远点的距离。由直径的定义可知，选 $x_1$ 为第一棵树的直径中点，可以让 $x_1$ 到第一棵树最远点的距离**最小**。
  - $x_2$ 到第二棵树最远点的距离。由直径的定义可知，选 $x_2$ 为第二棵树的直径中点，可以让 $x_2$ 到第二棵树最远点的距离**最小**。
  - $x_1\text{-}x_2$ 这条边的长度，即 $1$。
  - 三部分之和为

$$
\begin {aligned}
& \left\lceil\dfrac{d_1}{2}\right\rceil + \left\lceil\dfrac{d_2}{2}\right\rceil + 1   \\
={} & \left\lfloor\dfrac{d_1+1}{2}\right\rfloor + \left\lfloor\dfrac{d_2+1}{2}\right\rfloor + 1       \\
\end {aligned}
$$

![w404d.png](https://pic.leetcode.cn/1719719529-GSlAHr-w404d.png)

三种情况取最大值，答案为

$$
\max\left\{d_1,d_2, \left\lfloor\dfrac{d_1+1}{2}\right\rfloor + \left\lfloor\dfrac{d_2+1}{2}\right\rfloor + 1  \right\}
$$

```
func diameter(edges [][]int) (res int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	var dfs func(int, int) int
	dfs = func(x, fa int) (maxLen int) {
		for _, y := range g[x] {
			if y != fa {
				subLen := dfs(y, x) + 1
				res = max(res, maxLen+subLen)
				maxLen = max(maxLen, subLen)
			}
		}
		return
	}
	dfs(0, -1)
	return
}

func minimumDiameterAfterMerge(edges1, edges2 [][]int) int {
	d1 := diameter(edges1)
	d2 := diameter(edges2)
	return max(d1, d2, (d1+1)/2+(d2+1)/2+1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $\textit{edges}_1$ 的长度，$m$ 是 $\textit{edges}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(n+m)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)