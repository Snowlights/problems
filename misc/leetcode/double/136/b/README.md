#### 题目

<p>给你一个&nbsp;<code>m x n</code>&nbsp;的二进制矩阵&nbsp;<code>grid</code>&nbsp;。</p>

<p>如果矩阵中一行或者一列从前往后与从后往前读是一样的，那么我们称这一行或者这一列是 <strong>回文</strong> 的。</p>

<p>你可以将 <code>grid</code>&nbsp;中任意格子的值 <strong>翻转</strong>&nbsp;，也就是将格子里的值从 <code>0</code>&nbsp;变成 <code>1</code>&nbsp;，或者从 <code>1</code>&nbsp;变成 <code>0</code>&nbsp;。</p>

<p>请你返回 <strong>最少</strong>&nbsp;翻转次数，使得矩阵 <strong>要么</strong>&nbsp;所有行是 <strong>回文的</strong>&nbsp;，要么所有列是 <strong>回文的</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1,0,0],[0,0,0],[0,0,1]]</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><b>解释：</b></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/07/screenshot-from-2024-07-08-00-20-10.png" style="width: 420px; height: 108px;" /></p>

<p>将高亮的格子翻转，得到所有行都是回文的。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = </span>[[0,1],[0,1],[0,0]]</p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/07/screenshot-from-2024-07-08-00-31-23.png" style="width: 300px; height: 100px;" /></p>

<p>将高亮的格子翻转，得到所有列都是回文的。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1],[0]]</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><strong>解释：</strong></p>

<p>所有行已经是回文的。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m * n &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>0 &lt;= grid[i][j] &lt;= 1</code></li>
</ul>

#### 思路

先计算所有行变成回文最少需要翻转多少次。

也就是对于每一行 $\textit{row}$，计算这一行变成回文最少需要翻转多少次。

也就是累加 $\textit{row}[j]\ne \textit{row}[n-1-j]$ 的个数，其中 $0\le j \le \lfloor n/2 \rfloor$。

对于列，统计方式同理。

两种情况取最小值，即为答案。

```
func minFlips(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	diffRow := 0
	for _, row := range grid {
		for j := 0; j < n/2; j++ {
			if row[j] != row[n-1-j] {
				diffRow++
			}
		}
	}

	diffCol := 0
	for j := 0; j < n; j++ {
		for i, row := range grid[:m/2] {
			if row[j] != grid[m-1-i][j] {
				diffCol++
			}
		}
	}

	return min(diffRow, diffCol)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略 `zip` 的开销。

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