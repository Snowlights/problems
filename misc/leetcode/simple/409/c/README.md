#### 题目

<p>给你一个整数 <code>n</code> 和一个二维整数数组 <code>queries</code>。</p>

<p>有 <code>n</code> 个城市，编号从 <code>0</code> 到 <code>n - 1</code>。初始时，每个城市 <code>i</code> 都有一条<strong>单向</strong>道路通往城市 <code>i + 1</code>（ <code>0 &lt;= i &lt; n - 1</code>）。</p>

<p><code>queries[i] = [u<sub>i</sub>, v<sub>i</sub>]</code> 表示新建一条从城市 <code>u<sub>i</sub></code> 到城市 <code>v<sub>i</sub></code> 的<strong>单向</strong>道路。每次查询后，你需要找到从城市 <code>0</code> 到城市 <code>n - 1</code> 的<strong>最短路径</strong>的<strong>长度</strong>。</p>

<p>所有查询中不会存在两个查询都满足 <code>queries[i][0] &lt; queries[j][0] &lt; queries[i][1] &lt; queries[j][1]</code>。</p>

<p>返回一个数组 <code>answer</code>，对于范围 <code>[0, queries.length - 1]</code> 中的每个 <code>i</code>，<code>answer[i]</code> 是处理完<strong>前</strong> <code>i + 1</code> 个查询后，从城市 <code>0</code> 到城市 <code>n - 1</code> 的最短路径的<em>长度</em>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 5, queries = [[2, 4], [0, 2], [0, 4]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[3, 2, 1]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/28/image8.jpg" style="width: 350px; height: 60px;" /></p>

<p>新增一条从 2 到 4 的道路后，从 0 到 4 的最短路径长度为 3。</p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/28/image9.jpg" style="width: 350px; height: 60px;" /></p>

<p>新增一条从 0 到 2 的道路后，从 0 到 4 的最短路径长度为 2。</p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/28/image10.jpg" style="width: 350px; height: 96px;" /></p>

<p>新增一条从 0 到 4 的道路后，从 0 到 4 的最短路径长度为 1。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 4, queries = [[0, 3], [0, 2]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[1, 1]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/28/image11.jpg" style="width: 300px; height: 70px;" /></p>

<p>新增一条从 0 到 3 的道路后，从 0 到 3 的最短路径长度为 1。</p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/28/image12.jpg" style="width: 300px; height: 70px;" /></p>

<p>新增一条从 0 到 2 的道路后，从 0 到 3 的最短路径长度仍为 1。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>3 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>0 &lt;= queries[i][0] &lt; queries[i][1] &lt; n</code></li>
	<li><code>1 &lt; queries[i][1] - queries[i][0]</code></li>
	<li>查询中不存在重复的道路。</li>
	<li>不存在两个查询都满足 <code>i != j</code> 且 <code>queries[i][0] &lt; queries[j][0] &lt; queries[i][1] &lt; queries[j][1]</code>。</li>
</ul>

#### 思路

把目光放在**边**上。
初始有 $n-1$ 条边，我们在 $0\rightarrow 1$ 这条边上，目标是到达 $(n-2)\rightarrow (n-1)$ 这条边，并把这条边走完。
处理 $\textit{queries}$ 之前，需要走 $n-1$ 条边。

![w409c-1.jpg](https://pic.leetcode.cn/1722747389-ZsMpqd-w409c-1.jpg)

连一条从 $2$ 到 $4$ 的边，意味着什么？

相当于把 $2\rightarrow 3$ 这条边和 $3\rightarrow 4$ 这条边合并成一条边。现在从起点到终点需要 $3$ 条边。

![w409c-2.jpg](https://pic.leetcode.cn/1722747344-UibNQD-w409c.jpg)

连一条从 $0$ 到 $2$ 的边，意味着什么？

相当于把 $0\rightarrow 1$ 这条边和 $1\rightarrow 2$ 这条边合并成一条边。现在从起点到终点需要 $2$ 条边。

用**并查集**实现边的合并。初始化一个大小为 $n-1$ 的并查集，并查集中的节点 $i$ 表示题目的边 $i \rightarrow (i+1)$。

连一条从 $L$ 到 $R$ 的边，相当于把并查集中的节点 $L,L+1,L+2\cdots, R-2$ 合并到并查集中的节点 $R-1$ 上。 合并的同时，维护并查集连通块个数。
答案就是每次合并后的并查集连通块个数。

```
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	fa := make([]int, n-1)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	ans := make([]int, len(queries))
	cnt := n - 1 // 并查集连通块个数
	for qi, q := range queries {
		l, r := q[0], q[1]-1
		fr := find(r)
		for i := find(l); i < r; i = find(i + 1) {
			fa[i] = fr
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
} 
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $q$ 是 $\textit{queries}$ 的长度。注意每个点只会被合并一次，在后面的循环中会被 `i = find(l)` 以及 `i = find(i + 1)` 跳过。由于数组的特殊性，每次合并的复杂度为 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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
