#### 题目

<p>给你一个由正整数构成的二维矩阵 <code>grid</code>。</p>

<p>你需要从矩阵中选择<strong> 一个或多个 </strong>单元格，选中的单元格应满足以下条件：</p>

<ul>
	<li>所选单元格中的任意两个单元格都不会处于矩阵的 <strong>同一行</strong>。</li>
	<li>所选单元格的值 <strong>互不相同</strong>。</li>
</ul>

<p>你的得分为所选单元格值的<strong>总和</strong>。</p>

<p>返回你能获得的<strong> 最大 </strong>得分。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1,2,3],[4,3,2],[1,1,1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">8</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/29/grid1drawio.png" /></p>

<p>选择上图中用彩色标记的单元格，对应的值分别为 1、3 和 4 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[8,7,6],[8,3,2]]</span></p>

<p><strong>输出：</strong> <span class="example-io">15</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/29/grid8_8drawio.png" style="width: 170px; height: 114px;" /></p>

<p>选择上图中用彩色标记的单元格，对应的值分别为 7 和 8 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= grid.length, grid[i].length &lt;= 10</code></li>
	<li><code>1 &lt;= grid[i][j] &lt;= 100</code></li>
</ul>

#### 思路

![lc3276.png](https://pic.leetcode.cn/1725186494-jnwWda-lc3276.png)

看示例 1，$\textit{grid}$ 的最大值为 $4$。

讨论 $4$ **选或不选**：

- 不选 $4$，问题变成在 $[1,3]$ 中选数字，每行至多一个数且没有相同元素时，所选元素之和的最大值。
- 选 $4$，**枚举**选哪一行的 $4$（本例中只能选第二行），问题变成在 $[1,3]$ 中选数字，第二行不能选数字，每行至多一个数且没有相同元素时，所选元素之和的最大值。

于是，我们需要在递归的过程中，跟踪两个信息：

- $i$：当前要在 $[1,i]$ 中选数字。
- $j$：已选的行号集合为 $j$。后续所选数字所在的行号不能在 $j$ 中。

因此，定义状态为 $\textit{dfs}(i,j)$，表示在 $[1,i]$ 中选数字，所选数字所处的行号不能在集合 $j$ 中，每行至多一个数且没有相同元素时，所选元素之和的最大值。

讨论元素 $i$ **选或不选**：

- 不选 $i$，问题变成在 $[1,i-1]$ 中选数字，所选数字所处的行号不能在集合 $j$ 中，每行至多一个数且没有相同元素时，所选元素之和的最大值，即 $\textit{dfs}(i-1,j)$。
- 选 $i$，**枚举**选第 $k$ 行的 $i$（提前预处理 $i$ 所处的行号），问题变成在 $[1,i-1]$ 中选数字（注意 $i$ 只能选一次），所选数字所处的行号不能在集合 $j\cup \{k\}$ 中，每行至多一个数且没有相同元素时，所选元素之和的最大值，即 $\textit{dfs}(i-1,j\cup \{k\})$。

两种情况取最大值，即为 $\textit{dfs}(i,j)$。

递归边界：$\textit{dfs}(0,j) = 0$。

递归入口：$\textit{dfs}(\textit{mx},\varnothing)$。一开始没有选行号。

代码用到了一些位运算技巧，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

## 记忆化搜索

```
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	n := len(allNums)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 1<<len(grid))
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		// 不选 x
		res := dfs(i-1, j)
		// 枚举选第 k 行的 x
		x := allNums[i]
		for t, lb := pos[x], 0; t > 0; t ^= lb {
			lb = t & -t    // lb = 1<<k，其中 k 是行号
			if j&lb == 0 { // 没选过第 k 行的数
				res = max(res, dfs(i-1, j|lb)+x)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n-1, 0)
}
```

## 递推

```
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	allNums := make([]int, 0, len(pos))
	for x := range pos {
		allNums = append(allNums, x)
	}

	f := make([][]int, len(allNums)+1)
	for i := range f {
		f[i] = make([]int, 1<<len(grid))
	}
	for i, x := range allNums {
		for j, v := range f[i] {
			f[i+1][j] = v // 不选 x
			for t, lb := pos[x], 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[i+1][j] = max(f[i+1][j], f[i][j|lb]+x) // 选第 k 行的 x
				}
			}
		}
	}
	return f[len(allNums)][0]
}
```

## 空间压缩

去掉 $f$ 的第一个维度。

```
func maxScore(grid [][]int) int {
	pos := map[int]int{}
	for i, row := range grid {
		for _, x := range row {
			// 保存所有包含 x 的行号（压缩成二进制数）
			pos[x] |= 1 << i
		}
	}

	f := make([]int, 1<<len(grid))
	for x, posMask := range pos {
		for j := range f {
			for t, lb := posMask, 0; t > 0; t ^= lb {
				lb = t & -t    // lb = 1<<k，其中 k 是行号
				if j&lb == 0 { // 没选过第 k 行的数
					f[j] = max(f[j], f[j|lb]+x)
				}
			}
		}
	}
	return f[0]
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn2^m)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn + 2^m)$ 或 $\mathcal{O}(k + 2^m)$，其中 $k$ 是 $\textit{grid}$ 中的不同元素的个数。

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
