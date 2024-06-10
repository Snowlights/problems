#### 题目

<p>给你一个由 <strong>正整数</strong> 组成、大小为 <code>m x n</code> 的矩阵 <code>grid</code>。你可以从矩阵中的任一单元格移动到另一个位于正下方或正右侧的任意单元格（不必相邻）。从值为 <code>c1</code> 的单元格移动到值为 <code>c2</code> 的单元格的得分为 <code>c2 - c1</code> 。</p>

<p>你可以从<strong> 任一</strong> 单元格开始，并且必须至少移动一次。</p>

<p>返回你能得到的 <strong>最大 </strong>总得分。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/03/14/grid1.png" style="width: 240px; height: 240px;" />
<div class="example-block">
<p><strong>输入：</strong><span class="example-io">grid = [[9,5,7,3],[8,9,6,1],[6,7,14,3],[2,5,3,1]]</span></p>

<p><strong>输出：</strong><span class="example-io">9</span></p>

<p><strong>解释：</strong>从单元格 <code>(0, 1)</code> 开始，并执行以下移动：<br />
- 从单元格 <code>(0, 1)</code> 移动到 <code>(2, 1)</code>，得分为 <code>7 - 5 = 2</code> 。<br />
- 从单元格 <code>(2, 1)</code> 移动到 <code>(2, 2)</code>，得分为 <code>14 - 7 = 7</code> 。<br />
总得分为 <code>2 + 7 = 9</code> 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/04/08/moregridsdrawio-1.png" style="width: 180px; height: 116px;" /></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">grid = [[4,3,2],[3,2,1]]</span></p>

<p><strong>输出：</strong><span class="example-io">-1</span></p>

<p><strong>解释：</strong>从单元格 <code>(0, 0)</code> 开始，执行一次移动：从 <code>(0, 0)</code> 到 <code>(0, 1)</code> 。得分为 <code>3 - 4 = -1</code> 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>2 &lt;= m, n &lt;= 1000</code></li>
	<li><code>4 &lt;= m * n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= grid[i][j] &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

把 $\textit{grid}[i][j]$ 视作**海拔高度**，题目要计算的是重力势能的变化量之和，也就是终点和起点的海拔高度之差。

枚举终点海拔高度，那么**起点海拔高度越小越好**，且起点只能在 $(i,j)$ 的左上方向（可以是正左或者正上）。

按照[【图解】一张图秒懂二维前缀和！](https://leetcode.cn/problems/range-sum-query-2d-immutable/solution/tu-jie-yi-zhang-tu-miao-dong-er-wei-qian-84qp/)的思路，定义 $f[i+1][j+1]$ 表示左上角在 $(0,0)$，右下角在 $(i,j)$ 的子矩阵的**最小值**。

类似二维前缀和，我们有

$$
f[i+1][j+1] = \min(f[i+1][j],f[i][j+1], \textit{grid}[i][j])
$$

注意起点终点不能重合，如果终点在 $(i,j)$，那么起点的最小值为

$$
\min(f[i+1][j],f[i][j+1])
$$

取终点起点之差的最大值，即为答案。

``` go
func maxScore(grid [][]int) int {
	ans := math.MinInt
	colMin := make([]int, len(grid[0]))
	for i := range colMin {
		colMin[i] = math.MaxInt
	}
	for _, row := range grid {
		preMin := math.MaxInt
		for j, x := range row {
			ans = max(ans, x-min(preMin, colMin[j]))
			colMin[j] = min(colMin[j], x)
			preMin = min(preMin, colMin[j])
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(A\cdot n)$，其中 $n$ 为 $s$ 的长度，$A$ 为 $n$ 的因子个数。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。其中 $|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)