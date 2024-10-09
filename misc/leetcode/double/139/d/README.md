#### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的二维整数数组&nbsp;<code>coordinates</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;，其中&nbsp;<code>0 &lt;= k &lt; n</code>&nbsp;。</p>

<p><code>coordinates[i] = [x<sub>i</sub>, y<sub>i</sub>]</code>&nbsp;表示二维平面里一个点&nbsp;<code>(x<sub>i</sub>, y<sub>i</sub>)</code>&nbsp;。</p>

<p>如果一个点序列&nbsp;<code>(x<sub>1</sub>, y<sub>1</sub>)</code>, <code>(x<sub>2</sub>, y<sub>2</sub>)</code>, <code>(x<sub>3</sub>, y<sub>3</sub>)</code>, ..., <code>(x<sub>m</sub>, y<sub>m</sub>)</code>&nbsp;满足以下条件，那么我们称它是一个长度为 <code>m</code>&nbsp;的 <strong>上升序列</strong>&nbsp;：</p>

<ul>
	<li>对于所有满足&nbsp;<code>1 &lt;= i &lt; m</code> 的&nbsp;<code>i</code>&nbsp;都有&nbsp;<code>x<sub>i</sub> &lt; x<sub>i + 1</sub></code>&nbsp;且&nbsp;<code>y<sub>i</sub> &lt; y<sub>i + 1</sub></code>&nbsp;。</li>
	<li>对于所有&nbsp;<code>1 &lt;= i &lt;= m</code>&nbsp;的&nbsp;<code>i</code>&nbsp;对应的点&nbsp;<code>(x<sub>i</sub>, y<sub>i</sub>)</code>&nbsp;都在给定的坐标数组里。</li>
</ul>

<p>请你返回包含坐标 <code>coordinates[k]</code>&nbsp;的 <strong>最长上升路径</strong>&nbsp;的长度。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>coordinates = [[3,1],[2,2],[4,1],[0,0],[5,3]], k = 1</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<p><code>(0, 0)</code>&nbsp;，<code>(2, 2)</code>&nbsp;，<code>(5, 3)</code><!-- notionvc: 082cee9e-4ce5-4ede-a09d-57001a72141d -->&nbsp;是包含坐标 <code>(2, 2)</code>&nbsp;的最长上升路径。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>coordinates = [[2,1],[7,0],[5,6]], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><b>解释：</b></p>

<p><code>(2, 1)</code>&nbsp;，<code>(5, 6)</code>&nbsp;是包含坐标 <code>(5, 6)</code>&nbsp;的最长上升路径。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == coordinates.length &lt;= 10<sup>5</sup></code></li>
	<li><code>coordinates[i].length == 2</code></li>
	<li><code>0 &lt;= coordinates[i][0], coordinates[i][1] &lt;= 10<sup>9</sup></code></li>
	<li><code>coordinates</code>&nbsp;中的元素 <strong>互不相同</strong>&nbsp;。<!-- notionvc: 6e412fc2-f9dd-4ba2-b796-5e802a2b305a --><!-- notionvc: c2cf5618-fe99-4909-9b4c-e6b068be22a6 --></li>
	<li><code>0 &lt;= k &lt;= n - 1</code></li>
</ul>

#### 思路

设 $\textit{kx} = \textit{coordinates}[k][0],\ \textit{ky} = \textit{coordinates}[k][1]$。

仔细读题：我们要选的并不是 $\textit{coordinates}$ 的子序列，所以可以排序。

按照 $x$ 从小到大排序，问题变成计算 $y$ 的 LIS。

⚠**注意**：对于 $x$ 相同的点，要按照 $y$ **从大到小**排序。这可以保证在计算 LIS 时，对于相同的 $x$，我们至多选一个 $y$。

然后选择 $x<\textit{kx}$ 且 $y < \textit{ky}$ 或者 $x>\textit{kx}$ 且 $y > \textit{ky}$ 的 $y$ 计算 LIS。

关于 LIS 的二分算法，请看[【基础算法精讲 20】](https://www.bilibili.com/video/BV1ub411Q7sB/)。
```
func maxPathLength(coordinates [][]int, k int) int {
	kx, ky := coordinates[k][0], coordinates[k][1]
	sort.Slice(coordinates, func(i, j int) bool {
		a, b := coordinates[i], coordinates[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	g := []int{}
	for _, p := range coordinates {
		x, y := p[0], p[1]
		if x < kx && y < ky || x > kx && y > ky {
			j := sort.SearchInts(g, y)
			if j < len(g) {
				g[j] = y
			} else {
				g = append(g, y)
			}
		}
	}
	return len(g) + 1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{coordinates}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)