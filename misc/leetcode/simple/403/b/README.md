#### 题目

<p>给你一个二维 <strong>二进制 </strong>数组 <code>grid</code>。请你找出一个边在水平方向和竖直方向上、面积 <strong>最小</strong> 的矩形，并且满足 <code>grid</code> 中所有的 1 都在矩形的内部。</p>

<p>返回这个矩形可能的 <strong>最小 </strong>面积。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[0,1,0],[1,0,1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">6</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/05/08/examplerect0.png" style="padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; width: 279px; height: 198px;" /></p>

<p>这个最小矩形的高度为 2，宽度为 3，因此面积为 <code>2 * 3 = 6</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[0,0],[1,0]]</span></p>

<p><strong>输出：</strong> <span class="example-io">1</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/05/08/examplerect1.png" style="padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; width: 204px; height: 201px;" /></p>

<p>这个最小矩形的高度和宽度都是 1，因此面积为 <code>1 * 1 = 1</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= grid.length, grid[i].length &lt;= 1000</code></li>
	<li><code>grid[i][j]</code> 是 0 或 1。</li>
	<li>输入保证 <code>grid</code> 中至少有一个 1 。</li>
</ul>

#### 思路

设最左、最右、最上、最下的 $1$ 的行号/列号分别为 $\textit{left},\textit{right},\textit{top},\textit{bottom}$，则答案为

$$
(\textit{right} - \textit{left} + 1) \cdot (\textit{bottom} - \textit{top} + 1)
$$

```
func minimumArea(grid [][]int) int {
	left, right := len(grid[0]), 0
	top, bottom := len(grid), 0
	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				left = min(left, j)
				right = max(right, j)
				top = min(top, i)
				bottom = i
			}
		}
	}
	return (right - left + 1) * (bottom - top + 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)