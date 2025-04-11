#### 题目

<p data-pm-slice="1 3 []">给你一个大小为 <code>n x m</code>&nbsp;的二维矩阵&nbsp;<code>grid</code>&nbsp;，以及一个长度为 <code>n</code>&nbsp;的整数数组&nbsp;<code>limits</code>&nbsp;，和一个整数&nbsp;<code>k</code>&nbsp;。你的目标是从矩阵 <code>grid</code> 中提取出&nbsp;<strong>至多</strong> <code>k</code>&nbsp;个元素，并计算这些元素的最大总和，提取时需满足以下限制<b>：</b></p>

<ul data-spread="false">
	<li>
	<p>从 <code>grid</code>&nbsp;的第 <code>i</code> 行提取的元素数量不超过 <code>limits[i]</code> 。</p>
	</li>
</ul>

<p data-pm-slice="1 1 []">返回最大总和。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1,2],[3,4]], limits = [1,2], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>7</span></p>

<p><b>解释：</b></p>

<ul>
	<li>从第 2 行提取至多 2 个元素，取出 4 和 3 。</li>
	<li>至多提取 2 个元素时的最大总和&nbsp;<code>4 + 3 = 7</code>&nbsp;。</li>
</ul>
</div>

<p><b>示例 2：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b></span><span class="example-io">grid = [[5,3,7],[8,2,6]], limits = [2,2], k = 3</span></p>

<p><span class="example-io"><b>输出：</b></span><span class="example-io">21</span></p>

<p><b>解释：</b></p>

<ul>
	<li>从第 1&nbsp;行提取至多 2 个元素，取出 7 。</li>
	<li>从第 2 行提取至多 2 个元素，取出&nbsp;8 和 6 。</li>
	<li>至多提取 3&nbsp;个元素时的最大总和 <code>7 + 8 + 6 = 21</code>&nbsp;。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>n == grid.length == limits.length</code></li>
	<li><code>m == grid[i].length</code></li>
	<li><code>1 &lt;= n, m &lt;= 500</code></li>
	<li><code>0 &lt;= grid[i][j] &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= limits[i] &lt;= m</code></li>
	<li><code>0 &lt;= k &lt;= min(n * m, sum(limits))</code></li>
</ul>

#### 思路

本题没有负数，所以每排都可以取恰好 $\textit{limits}[i]$ 个数。

为了最大化元素和，每排取最大的 $\textit{limits}[i]$ 个数。

然后再取这些数中最大的 $k$ 个数，求和即为答案。

```
func maxSum(grid [][]int, limits []int, k int) (ans int64) {
	a := []int{}
	for i, row := range grid {
		slices.SortFunc(row, func(a, b int) int { return b - a })
		a = append(a, row[:limits[i]]...)
	}
	slices.SortFunc(a, func(a, b int) int { return b - a })
	for _, x := range a[:k] {
		ans += int64(x)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn\log (mn))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。用快速选择可以做到 $\mathcal{O}(mn)$，见 C++ 代码。
- 空间复杂度：$\mathcal{O}(mn)$，或者 $\textit{limits}[i]$ 之和。

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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)