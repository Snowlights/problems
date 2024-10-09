### 题目

<p>一个公司在全国有 <code>n</code>&nbsp;个分部，它们之间有的有道路连接。一开始，所有分部通过这些道路两两之间互相可以到达。</p>

<p>公司意识到在分部之间旅行花费了太多时间，所以它们决定关闭一些分部（<b>也可能不关闭任何分部</b>），同时保证剩下的分部之间两两互相可以到达且最远距离不超过&nbsp;<code>maxDistance</code>&nbsp;。</p>

<p>两个分部之间的 <strong>距离</strong> 是通过道路长度之和的 <strong>最小值</strong>&nbsp;。</p>

<p>给你整数&nbsp;<code>n</code>&nbsp;，<code>maxDistance</code>&nbsp;和下标从 <strong>0</strong>&nbsp;开始的二维整数数组&nbsp;<code>roads</code>&nbsp;，其中&nbsp;<code>roads[i] = [u<sub>i</sub>, v<sub>i</sub>, w<sub>i</sub>]</code>&nbsp;表示一条从&nbsp;<code>u<sub>i</sub></code>&nbsp;到&nbsp;<code>v<sub>i</sub></code>&nbsp;长度为&nbsp;<code>w<sub>i</sub></code>的&nbsp;<strong>无向</strong>&nbsp;道路。</p>

<p>请你返回关闭分部的可行方案数目，满足每个方案里剩余分部之间的最远距离不超过<em>&nbsp;</em><code>maxDistance</code>。</p>

<p><strong>注意</strong>，关闭一个分部后，与之相连的所有道路不可通行。</p>

<p><b>注意</b>，两个分部之间可能会有多条道路。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/08/example11.png" style="width: 221px; height: 191px;" /></p>

<pre>
<b>输入：</b>n = 3, maxDistance = 5, roads = [[0,1,2],[1,2,10],[0,2,10]]
<b>输出：</b>5
<b>解释：</b>可行的关闭分部方案有：
- 关闭分部集合 [2] ，剩余分部为 [0,1] ，它们之间的距离为 2 。
- 关闭分部集合 [0,1] ，剩余分部为 [2] 。
- 关闭分部集合 [1,2] ，剩余分部为 [0] 。
- 关闭分部集合 [0,2] ，剩余分部为 [1] 。
- 关闭分部集合 [0,1,2] ，关闭后没有剩余分部。
总共有 5 种可行的关闭方案。
</pre>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/11/08/example22.png" style="width: 221px; height: 241px;" /></p>

<pre>
<b>输入：</b>n = 3, maxDistance = 5, roads = [[0,1,20],[0,1,10],[1,2,2],[0,2,2]]
<b>输出：</b>7
<b>解释：</b>可行的关闭分部方案有：
- 关闭分部集合 [] ，剩余分部为 [0,1,2] ，它们之间的最远距离为 4 。
- 关闭分部集合 [0] ，剩余分部为 [1,2] ，它们之间的距离为 2 。
- 关闭分部集合 [1] ，剩余分部为 [0,2] ，它们之间的距离为 2 。
- 关闭分部集合 [0,1] ，剩余分部为 [2] 。
- 关闭分部集合 [1,2] ，剩余分部为 [0] 。
- 关闭分部集合 [0,2] ，剩余分部为 [1] 。
- 关闭分部集合 [0,1,2] ，关闭后没有剩余分部。
总共有 7 种可行的关闭方案。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>n = 1, maxDistance = 10, roads = []
<b>输出：</b>2
<b>解释：</b>可行的关闭分部方案有：
- 关闭分部集合 [] ，剩余分部为 [0] 。
- 关闭分部集合 [0] ，关闭后没有剩余分部。
总共有 2 种可行的关闭方案。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10</code></li>
	<li><code>1 &lt;= maxDistance &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= roads.length &lt;= 1000</code></li>
	<li><code>roads[i].length == 3</code></li>
	<li><code>0 &lt;= u<sub>i</sub>, v<sub>i</sub> &lt;= n - 1</code></li>
	<li><code>u<sub>i</sub> != v<sub>i</sub></code></li>
	<li><code>1 &lt;= w<sub>i</sub> &lt;= 1000</code></li>
	<li>一开始所有分部之间通过道路互相可以到达。</li>
</ul>

### 思路

根据 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/) 中「枚举集合」的技巧，
我们可以枚举保留哪些节点，在这些节点之间连边。

连边后，根据[【模板讲解】带你发明 Floyd 算法：从记忆化搜索到递推](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/) 求出任意两点之间最短路，
如果保留的节点之间的最短路均不超过 $\textit{maxDistance}$，则把答案加一。

代码实现时，可以先预处理原图的邻接矩阵 $g$，这样无需每次枚举子集都要遍历 $\textit{roads}$，直接拷贝一份 $g$ 即可。

```go  
func numberOfSets(n, maxDistance int, roads [][]int) (ans int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, n)
		for j := range g[i] {
			if j != i { // g[i][i] = 0
				g[i][j] = math.MaxInt / 2 // 防止加法溢出
			}
		}
	}
	for _, e := range roads {
		x, y, wt := e[0], e[1], e[2]
		g[x][y] = min(g[x][y], wt)
		g[y][x] = min(g[y][x], wt)
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
next:
	for s := 0; s < 1<<n; s++ { // 枚举子集
		for i, row := range g {
			if s>>i&1 == 0 {
				continue
			}
			copy(f[i], row)
		}

		// Floyd
		for k := range f {
			if s>>k&1 == 0 {
				continue
			}
			for i := range f {
				if s>>i&1 == 0 {
					continue
				}
				for j := range f {
					f[i][j] = min(f[i][j], f[i][k]+f[k][j])
				}
			}
		}

		for i, di := range f {
			if s>>i&1 == 0 {
				continue
			}
			for j, dij := range di {
				if s>>j&1 > 0 && dij > maxDistance {
					continue next
				}
			}
		}
		ans++
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(m + 2^nn^3)$，其中 $m$ 为 $\textit{roads}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 相关题目

#### Floyd
- [2642. 设计可以求最短路径的图类](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/) 1811
- [1334. 阈值距离内邻居最少的城市](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) 1855
- [2101. 引爆最多的炸弹](https://leetcode.cn/problems/detonate-the-maximum-bombs/) 1880

#### 二进制枚举

- [78. 子集](https://leetcode.cn/problems/subsets/)
- [77. 组合](https://leetcode.cn/problems/combinations/)
- [1286. 字母组合迭代器](https://leetcode.cn/problems/iterator-for-combination/) 1591
- [2397. 被列覆盖的最多行数](https://leetcode.cn/problems/maximum-rows-covered-by-columns/) 1719
- [2212. 射箭比赛中的最大得分](https://leetcode.cn/problems/maximum-points-in-an-archery-competition/) 1869
- [1601. 最多可达成的换楼请求数目](https://leetcode.cn/problems/maximum-number-of-achievable-transfer-requests/) 2119


