#### 题目

<p>给你一个二维 <strong>二进制 </strong>数组 <code>grid</code>。你需要找到 3 个<strong> 不重叠</strong>、面积 <strong>非零</strong> 、边在水平方向和竖直方向上的矩形，并且满足 <code>grid</code> 中所有的 1 都在这些矩形的内部。</p>

<p>返回这些矩形面积之和的<strong> 最小 </strong>可能值。</p>

<p><strong>注意</strong>，这些矩形可以相接。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1,0,1],[1,1,1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/05/14/example0rect21.png" style="padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; width: 280px; height: 198px;" /></p>

<ul>
	<li>位于 <code>(0, 0)</code> 和 <code>(1, 0)</code> 的 1 被一个面积为 2 的矩形覆盖。</li>
	<li>位于 <code>(0, 2)</code> 和 <code>(1, 2)</code> 的 1 被一个面积为 2 的矩形覆盖。</li>
	<li>位于 <code>(1, 1)</code> 的 1 被一个面积为 1 的矩形覆盖。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1,0,1,0],[0,1,0,1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/05/14/example1rect2.png" style="padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; width: 356px; height: 198px;" /></p>

<ul>
	<li>位于 <code>(0, 0)</code> 和 <code>(0, 2)</code> 的 1 被一个面积为 3 的矩形覆盖。</li>
	<li>位于 <code>(1, 1)</code> 的 1 被一个面积为 1 的矩形覆盖。</li>
	<li>位于 <code>(1, 3)</code> 的 1 被一个面积为 1 的矩形覆盖。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= grid.length, grid[i].length &lt;= 30</code></li>
	<li><code>grid[i][j]</code> 是 0 或 1。</li>
	<li>输入保证 <code>grid</code> 中至少有三个 1 。</li>
</ul>

#### 思路

## 方法一：暴力枚举

一共有如下六种情况。

![w403d.png](https://pic.leetcode.cn/1719114413-gJmraG-w403d.png)

暴力枚举分割线的位置，划分成三个区域，每个区域对应周赛第二题，见 [题解](https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i/solutions/2819335/bian-li-pythonjavacgo-by-endlesscheng-6po1/)。

代码实现时，只需实现上面三种，下面三种可以通过把 $\textit{grid}$ 顺时针旋转 90° 得到。

```
func minimumArea(a [][]int, l, r int) int {
	left, right := len(a[0]), 0
	top, bottom := len(a), 0
	for i, row := range a {
		for j, x := range row[l:r] {
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

func minimumSum(grid [][]int) int {
	ans := math.MaxInt
	f := func(a [][]int) {
		m, n := len(a), len(a[0])
		if m >= 3 {
			for i := 1; i < m; i++ {
				for j := i + 1; j < m; j++ {
					// 图片上左
					area := minimumArea(a[:i], 0, n)
					area += minimumArea(a[i:j], 0, n)
					area += minimumArea(a[j:], 0, n)
					ans = min(ans, area)
				}
			}
		}
		if m >= 2 && n >= 2 {
			for i := 1; i < m; i++ {
				for j := 1; j < n; j++ {
					// 图片上中
					area := minimumArea(a[:i], 0, n)
					area += minimumArea(a[i:], 0, j)
					area += minimumArea(a[i:], j, n)
					ans = min(ans, area)
					// 图片上右
					area = minimumArea(a[:i], 0, j)
					area += minimumArea(a[:i], j, n)
					area += minimumArea(a[i:], 0, n)
					ans = min(ans, area)
				}
			}
		}
	}
	f(grid)
	f(rotate(grid))
	return ans
}

// 顺时针旋转矩阵 90°
func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((mn)^2)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。


## 方法二：用 DP 预处理

定义 $f[i+1][j+1]$ 表示包含「左上角为 $(0,0)$ 右下角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。

定义 $\textit{border}[i+1][j+1]$ 包含三个数，分别表示上述最小矩形的上边界、左边界和右边界。

从上到下，从左到右遍历 $\textit{grid}$。设当前遍历到 $\textit{grid}[i]$ 这一排，其中最左边的 $1$ 和最右边的 $1$ 的列号分别为 $\textit{left}$ 和 $\textit{right}$，分类讨论：

- 如果 $\textit{grid}[i]$ 这一行全为 $0$，那么「左上角为 $(0,0)$ 右下角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积，等于「左上角为 $(0,0)$ 右下角为 $(i-1,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积，即 $f[i+1][j+1] = f[i][j+1],\ \textit{border}[i+1][j+1]=\textit{border}[i][j+1]$。
- 如果 $\textit{grid}[i]$ 这一行包含 $1$，且上面的全为 $0$，那么 $f[i+1][j+1] = \textit{right}-\textit{left}+1,\ \textit{border}[i+1][j+1]=(i,\textit{left},\textit{right})$。
- 如果 $\textit{grid}[i]$ 这一行包含 $1$，且上面也包含 $1$，那么最小矩形：
    - 上边界是 $\textit{border}[i][j+1]$ 的上边界 $t$。
    - 左边界是 $\textit{border}[i][j+1]$ 的左边界与 $\textit{left}$ 的最小值 $l$。
    - 右边界是 $\textit{border}[i][j+1]$ 的右边界与 $\textit{right}$ 的最大值 $r$。
    - 下边界是 $i$。
    - $f[i+1][j+1] = (r - l + 1) \cdot (i - t + 1)$。
    - $\textit{border}[i+1][j+1]=(t,l,r)$。

代码实现时，$\textit{border}$ 可以用一个长为 $n$ 的数组滚动计算。

按照上述方法，预处理：

- 包含「左上角为 $(0,0)$ 右下角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。
- 包含「左下角为 $(m-1,0)$ 右上角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。
- 包含「右下角为 $(m-1,n-1)$ 左上角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。
- 包含「右上角为 $(0,n-1)$ 左下角为 $(i,j)$ 的子矩形」中的所有 $1$ 的最小矩形面积。

这样就可以快速计算出方法一图中的上中和上右两种情况。

对于方法一图中的上左情况，如何计算中间区域的最小矩形面积？

预处理每行最左最右 $1$ 的列号。在枚举两条分割线的同时，维护中间区域的最左最右 $1$ 的列号，以及最上最下的 $1$ 的行号。

```
func minimumArea(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	// f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	type data struct{ top, left, right int }
	border := make([]data, n)
	for j := range border {
		border[j].top = -1 // 无
	}

	for i, row := range a {
		left, right := -1, 0
		for j, x := range row {
			if x > 0 {
				if left < 0 {
					left = j
				}
				right = j
			}
			preB := border[j]
			if left < 0 { // 这一排目前全是 0
				f[i+1][j+1] = f[i][j+1] // 等于上面的结果
			} else if preB.top < 0 { // 这一排有 1，上面全是 0
				f[i+1][j+1] = right - left + 1
				border[j] = data{i, left, right}
			} else { // 这一排有 1，上面也有 1
				l, r := min(preB.left, left), max(preB.right, right)
				f[i+1][j+1] = (r - l + 1) * (i - preB.top + 1)
				border[j] = data{preB.top, l, r}
			}
		}
	}
	return f
}

func minimumSum(grid [][]int) int {
	ans := math.MaxInt
	f := func(a [][]int) {
		m, n := len(a), len(a[0])
		type pair struct{ l, r int }
		lr := make([]pair, m) // 每一行最左最右 1 的列号
		for i, row := range a {
			l, r := -1, 0
			for j, x := range row {
				if x > 0 {
					if l < 0 {
						l = j
					}
					r = j
				}
			}
			lr[i] = pair{l, r}
		}

		// lt[i+1][j+1] = 包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		lt := minimumArea(a)
		a = rotate(a)
		// lb[i][j+1] = 包含【左下角为 (m-1,0) 右上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		lb := rotate(rotate(rotate(minimumArea(a))))
		a = rotate(a)
		// rb[i][j] = 包含【右下角为 (m-1,n-1) 左上角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		rb := rotate(rotate(minimumArea(a)))
		a = rotate(a)
		// rt[i+1][j] = 包含【右上角为 (0,n-1) 左下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		rt := rotate(minimumArea(a))

		if m >= 3 {
			for i := 1; i < m; i++ {
				left, right, top, bottom := n, 0, m, 0
				for j := i + 1; j < m; j++ {
					if p := lr[j-1]; p.l >= 0 {
						left = min(left, p.l)
						right = max(right, p.r)
						top = min(top, j-1)
						bottom = j - 1
					}
					// 图片上左
					area := lt[i][n] // minimumArea(a[:i], 0, n)
					area += (right - left + 1) * (bottom - top + 1) // minimumArea(a[i:j], 0, n)
					area += lb[j][n] // minimumArea(a[j:], 0, n)
					ans = min(ans, area)
				}
			}
		}

		if m >= 2 && n >= 2 {
			for i := 1; i < m; i++ {
				for j := 1; j < n; j++ {
					// 图片上中
					area := lt[i][n] // minimumArea(a[:i], 0, n)
					area += lb[i][j] // minimumArea(a[i:], 0, j)
					area += rb[i][j] // minimumArea(a[i:], j, n)
					ans = min(ans, area)
					// 图片上右
					area = lt[i][j]  // minimumArea(a[:i], 0, j)
					area += rt[i][j] // minimumArea(a[:i], j, n)
					area += lb[i][n] // minimumArea(a[i:], 0, n)
					ans = min(ans, area)
				}
			}
		}
	}
	f(grid)
	f(rotate(grid))
	return ans
}

// 顺时针旋转矩阵 90°
func rotate(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, m)
	}
	for i, row := range a {
		for j, x := range row {
			b[j][m-1-i] = x
		}
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)