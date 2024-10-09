#### 题目

<p>给你一个整数 <code>n</code> 和一个二维整数数组 <code>queries</code>。</p>

<p>有 <code>n</code> 个城市，编号从 <code>0</code> 到 <code>n - 1</code>。初始时，每个城市 <code>i</code> 都有一条<strong>单向</strong>道路通往城市 <code>i + 1</code>（ <code>0 &lt;= i &lt; n - 1</code>）。</p>

<p><code>queries[i] = [u<sub>i</sub>, v<sub>i</sub>]</code> 表示新建一条从城市 <code>u<sub>i</sub></code> 到城市 <code>v<sub>i</sub></code> 的<strong>单向</strong>道路。每次查询后，你需要找到从城市 <code>0</code> 到城市 <code>n - 1</code> 的<strong>最短路径</strong>的<strong>长度</strong>。</p>

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

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= n &lt;= 500</code></li>
	<li><code>1 &lt;= queries.length &lt;= 500</code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>0 &lt;= queries[i][0] &lt; queries[i][1] &lt; n</code></li>
	<li><code>1 &lt; queries[i][1] - queries[i][0]</code></li>
	<li>查询中没有重复的道路。</li>
</ul>

#### 思路

## 方法一：BFS

暴力。每次加边后重新跑一遍 BFS，求出从 $0$ 到 $n-1$ 的最短路。

### 细节

为避免反复创建 $\textit{vis}$ 数组，可以在 $\textit{vis}$ 中保存当前节点是第几次询问访问的。

```
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	g := make([][]int, n-1)
	for i := range g {
		g[i] = append(g[i], i+1)
	}

	vis := make([]int, n-1)
	bfs := func(i int) int {
		q := []int{0}
		for step := 1; ; step++ {
			tmp := q
			q = nil
			for _, x := range tmp {
				for _, y := range g[x] {
					if y == n-1 {
						return step
					}
					if vis[y] != i {
						vis[y] = i
						q = append(q, y)
					}
				}
			}
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		g[q[0]] = append(g[q[0]], q[1])
		ans[i] = bfs(i + 1)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(q(n+q))$，其中 $q$ 是 $\textit{queries}$ 的长度。每次 BFS 的时间是 $\mathcal{O}(n+q)$。
- 空间复杂度：$\mathcal{O}(n+q)$。


## 方法二：DP

定义 $f[i]$ 为从 $0$ 到 $i$ 的最短路。
用 $\textit{from}[i]$ 记录额外添加的边的终点是 $i$，起点列表是 $\textit{from}[i]$。

我们可以从 $i-1$ 到 $i$，也可以从 $\textit{from}[i][j]$ 到 $i$，这些位置作为转移来源，用其 $f$ 值 $+1$ 更新 $f[i]$ 的最小值。

初始值：$f[i] = i$。
答案：$f[n-1]$。

### 细节

设添加的边为 $l\rightarrow r$，只有当 $f[l]+1 < f[r]$ 时才更新 DP。

``` 
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	from := make([][]int, n)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		f[i] = i
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		from[r] = append(from[r], l)
		if f[l]+1 < f[r] {
			f[r] = f[l] + 1
			for i := r + 1; i < n; i++ {
				f[i] = min(f[i], f[i-1]+1)
				for _, j := range from[i] {
					f[i] = min(f[i], f[j]+1)
				}
			}
		}
		ans[qi] = f[n-1]
	}
	return ans
}
```

#### 复杂度分析

= 时间复杂度：$\mathcal{O}(q(n+q))$，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n+q)$。

## 题单

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