### 题目

<p>给你一个大小为 <code>m x n</code>&nbsp;的二维整数数组&nbsp;<code>grid</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;。</p>

<p>你的任务是统计满足以下 <strong>条件</strong> 且从左上格子&nbsp;<code>(0, 0)</code>&nbsp;出发到达右下格子&nbsp;<code>(m - 1, n - 1)</code>&nbsp;的路径数目：</p>

<ul>
	<li>每一步你可以向右或者向下走，也就是如果格子存在的话，可以从格子&nbsp;<code>(i, j)</code>&nbsp;走到格子&nbsp;<code>(i, j + 1)</code>&nbsp;或者格子&nbsp;<code>(i + 1, j)</code>&nbsp;。</li>
	<li>路径上经过的所有数字&nbsp;<code>XOR</code>&nbsp;异或值必须 <strong>等于</strong>&nbsp;<code>k</code>&nbsp;。</li>
</ul>

<p>请你返回满足上述条件的路径总数。</p>

<p>由于答案可能很大，请你将答案对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong> 后返回。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[2, 1, 5], [7, 10, 0], [12, 6, 4]], k = 11</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><b>解释：</b></p>

<p>3 条路径分别为：</p>

<ul>
	<li><code>(0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2)</code></li>
	<li><code>(0, 0) → (1, 0) → (1, 1) → (1, 2) → (2, 2)</code></li>
	<li><code>(0, 0) → (0, 1) → (1, 1) → (2, 1) → (2, 2)</code></li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1, 3, 3, 3], [0, 3, 3, 2], [3, 0, 1, 1]], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>5</span></p>

<p><b>解释：</b></p>

<p>5 条路径分别为：</p>

<ul>
	<li><code>(0, 0) → (1, 0) → (2, 0) → (2, 1) → (2, 2) → (2, 3)</code></li>
	<li><code>(0, 0) → (1, 0) → (1, 1) → (2, 1) → (2, 2) → (2, 3)</code></li>
	<li><code>(0, 0) → (1, 0) → (1, 1) → (1, 2) → (1, 3) → (2, 3)</code></li>
	<li><code>(0, 0) → (0, 1) → (1, 1) → (1, 2) → (2, 2) → (2, 3)</code></li>
	<li><code>(0, 0) → (0, 1) → (0, 2) → (1, 2) → (2, 2) → (2, 3)</code></li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1, 1, 1, 2], [3, 0, 3, 2], [3, 0, 2, 2]], k = 10</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= m == grid.length &lt;= 300</code></li>
	<li><code>1 &lt;= n == grid[r].length &lt;= 300</code></li>
	<li><code>0 &lt;= grid[r][c] &lt; 16</code></li>
	<li><code>0 &lt;= k &lt; 16</code></li>
</ul>

### 思路

## 方法一：记忆化搜索

本题由于值域范围小，可以把路径的 XOR 值作为 DP 的第三个参数，即定义 $\textit{dfs}(i,j,x)$ 表示从左上角 $(0,0)$ 到当前位置 $(i,j)$，路径 XOR 值为 $x$ 的方案数。

设从左上角 $(0,0)$ 到 $(i,j-1)$ 的路径 XOR 值为 $y$，那么必须满足

$$
y\oplus \textit{grid}[i][j] = x
$$

即

$$
y = x \oplus \textit{grid}[i][j]
$$

其中 $\oplus$ 表示异或运算。

分类讨论怎么到达 $(i,j)$：

- 如果是从左边过来，根据上文的公式，有 $\textit{dfs}(i,j,x) = \textit{dfs}(i,j-1,x\oplus \textit{grid}[i][j])$。
- 如果是从上边过来，则 $\textit{dfs}(i,j,x) = \textit{dfs}(i-1,j,x\oplus \textit{grid}[i][j])$。

两条路径互斥，根据加法原理，有

$$
\textit{dfs}(i,j,x) = \textit{dfs}(i,j-1,x\oplus \textit{grid}[i][j]) + \textit{dfs}(i-1,j,x\oplus \textit{grid}[i][j])
$$

> 倒序递归是为了方便后面 1:1 翻译成正序的递推。

递归边界：

- $\textit{dfs}(-1,j,x)=\textit{dfs}(i,-1,x)=0$。出界，方案数为 $0$。
- $\textit{dfs}(0,0,\textit{grid}[0][0])=1$，其余 $\textit{dfs}(0,0,x)=0$。左上角 $(0,0)$ 到它自己的方案数是 $1$，路径 XOR 值为 $\textit{grid}[0][0]$。

递归入口：$\textit{dfs}(m-1,n-1,k)$。

### 细节

设 $\textit{mx}$ 为 $\textit{grid}[i][j]$ 的最大值，其二进制长度为 $L$。

那么 $2^L-1$ 就是 XOR 能取到的最大值。

如果 $k > 2^L-1$，那么直接返回 $0$。

> 注：也可以用所有 $\textit{grid}[i][j]$ 的 OR，作为 XOR 可以取到的最大值。

```
func countPathsWithXorValue(grid [][]int, k int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}
	u := 1 << bits.Len(uint(mx))
	if k >= u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, u)
			for x := range memo[i][j] {
				memo[i][j][x] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, x int) int {
		if i < 0 || j < 0 {
			return 0
		}
		val := grid[i][j]
		if i == 0 && j == 0 {
			if x == val {
				return 1
			}
			return 0
		}
		p := &memo[i][j][x]
		if *p != -1 {
			return *p
		}
		*p = (dfs(i, j-1, x^val) + dfs(i-1, j, x^val)) % mod
		return *p
	}
	return dfs(m-1, n-1, k)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U=\max(\textit{grid[i][j]})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mnU)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mnU)$。
- 空间复杂度：$\mathcal{O}(mnU)$。保存多少状态，就需要多少空间。

## 方法二：递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j+1][x]$ 的定义和 $\textit{dfs}(i,j,x)$ 的定义是一样的，都表示从左上角 $(0,0)$ 到当前位置 $(i,j)$，路径 XOR 值为 $x$ 的方案数。这里 $+1$ 是为了把 $\textit{dfs}(-1,j,x)$ 和 $\textit{dfs}(i,-1,x)$ 也翻译过来，这样我们可以把 $f[0][j][x]$ 和 $f[i][0][x]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][j+1][x] = f[i+1][j][x\oplus \textit{grid}[i][j]] + f[i][j+1][x\oplus \textit{grid}[i][j]]
$$

初始值：

- $f[1][1][\textit{grid}[0][0]]=1$，翻译自递归边界 $\textit{dfs}(0,0,\textit{grid}[0][0])=1$。
- 其余为 $0$。

答案为 $f[m][n][k]$，翻译自递归入口 $\textit{dfs}(m-1,n-1,k)$。

```
func countPathsWithXorValue(grid [][]int, k int) int {
	const mod = 1_000_000_007
	mx := 0
	for _, row := range grid {
		mx = max(mx, slices.Max(row))
	}
	u := 1 << bits.Len(uint(mx))
	if k >= u {
		return 0
	}

	m, n := len(grid), len(grid[0])
	f := make([][][]int, m+1)
	for i := range f {
		f[i] = make([][]int, n+1)
		for j := range f[i] {
			f[i][j] = make([]int, u)
		}
	}
	f[1][1][grid[0][0]] = 1
	for i, row := range grid {
		for j, val := range row {
			for x := range u {
				f[i+1][j+1][x] += (f[i+1][j][x^val] + f[i][j+1][x^val]) % mod
			}
		}
	}
	return f[m][n][k]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mnU)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U=\max(\textit{grid[i][j]})$。
- 空间复杂度：$\mathcal{O}(mnU)$。


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