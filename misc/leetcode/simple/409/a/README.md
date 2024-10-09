#### 题目

<p>给你一个 <code>n x n</code> 的二维数组 <code>grid</code>，它包含范围 <code>[0, n<sup>2</sup> - 1]</code> 内的<strong>不重复</strong>元素。</p>

<p>实现 <code>neighborSum</code> 类：</p>

<ul>
	<li><code>neighborSum(int [][]grid)</code> 初始化对象。</li>
	<li><code>int adjacentSum(int value)</code> 返回在 <code>grid</code> 中与 <code>value</code> 相邻的元素之<strong>和</strong>，相邻指的是与 <code>value</code> 在上、左、右或下的元素。</li>
	<li><code>int diagonalSum(int value)</code> 返回在 <code>grid</code> 中与 <code>value</code> 对角线相邻的元素之<strong>和</strong>，对角线相邻指的是与 <code>value</code> 在左上、右上、左下或右下的元素。</li>
</ul>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/24/design.png" style="width: 400px; height: 248px;" /></p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong></p>

<p>["neighborSum", "adjacentSum", "adjacentSum", "diagonalSum", "diagonalSum"]</p>

<p>[[[[0, 1, 2], [3, 4, 5], [6, 7, 8]]], [1], [4], [4], [8]]</p>

<p><strong>输出：</strong> [null, 6, 16, 16, 4]</p>

<p><strong>解释：</strong></p>

<p><strong class="example"><img alt="" src="https://assets.leetcode.com/uploads/2024/06/24/designexample0.png" style="width: 250px; height: 249px;" /></strong></p>

<ul>
	<li>1 的相邻元素是 0、2 和 4。</li>
	<li>4 的相邻元素是 1、3、5 和 7。</li>
	<li>4 的对角线相邻元素是 0、2、6 和 8。</li>
	<li>8 的对角线相邻元素是 4。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong></p>

<p>["neighborSum", "adjacentSum", "diagonalSum"]</p>

<p>[[[[1, 2, 0, 3], [4, 7, 15, 6], [8, 9, 10, 11], [12, 13, 14, 5]]], [15], [9]]</p>

<p><strong>输出：</strong> [null, 23, 45]</p>

<p><strong>解释：</strong></p>

<p><strong class="example"><img alt="" src="https://assets.leetcode.com/uploads/2024/06/24/designexample2.png" style="width: 300px; height: 300px;" /></strong></p>

<ul>
	<li>15 的相邻元素是 0、10、7 和 6。</li>
	<li>9 的对角线相邻元素是 4、12、14 和 15。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= n == grid.length == grid[0].length &lt;= 10</code></li>
	<li><code>0 &lt;= grid[i][j] &lt;= n<sup>2</sup> - 1</code></li>
	<li>所有 <code>grid[i][j]</code> 值均不重复。</li>
	<li><code>adjacentSum</code> 和 <code>diagonalSum</code> 中的 <code>value</code> 均在范围 <code>[0, n<sup>2</sup> - 1]</code> 内。</li>
	<li>最多会调用 <code>adjacentSum</code> 和 <code>diagonalSum</code> 总共 <code>2 * n<sup>2</sup></code> 次。</li>
</ul>

#### 思路

用一个长为 $8$ 的数组存放偏移向量，前 $4$ 个表示上下左右四个方向，后 $4$ 个表示斜向的四个方向。用一个大小为 $n^2\times 2$ 的数组 $s$ 保存元素和，其中 $s[v][0]$ 为 $\texttt{adjacentSum}(v)$ 的结果，$s[v][1]$ 为 $\texttt{diagonalSum}(v)$ 的结果。这可以在初始化时，遍历 $\textit{grid}[i][j]$ 以及偏移向量，累加每个元素的相邻元素之和计算出来。\n\n> 注：也可以在 $\textit{grid}$ 外面加一圈 $0$，这样无需判断下标越界。

```
var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, 1}, {-1, -1}, {1, -1}}

type neighborSum [][2]int

func Constructor(grid [][]int) neighborSum {
	n := len(grid)
	s := make(neighborSum, n*n)
	for i, row := range grid {
		for j, v := range row {
			for k, d := range dirs {
				x, y := i+d.x, j+d.y
				if 0 <= x && x < n && 0 <= y && y < n {
					s[v][k/4] += grid[x][y]
				}
			}
		}
	}
	return s
}

func (s neighborSum) AdjacentSum(value int) int {
	return s[value][0]
}

func (s neighborSum) DiagonalSum(value int) int {
	return s[value][1]
}
```

#### 复杂度分析

- 时间复杂度：初始化 $\mathcal{O}(n^2)$，其余 $\mathcal{O}(1)$，其中 $n$ 为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：初始化 $\mathcal{O}(n^2)$，其余 $\mathcal{O}(1)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)