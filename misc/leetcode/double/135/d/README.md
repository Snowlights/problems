#### 题目

<p>给你一个大小为 <code>n x n</code>&nbsp;的二维矩阵&nbsp;<code>grid</code>&nbsp;，一开始所有格子都是白色的。一次操作中，你可以选择任意下标为&nbsp;<code>(i, j)</code>&nbsp;的格子，并将第&nbsp;<code>j</code>&nbsp;列中从最上面到第&nbsp;<code>i</code>&nbsp;行所有格子改成黑色。</p>

<p>如果格子 <code>(i, j)</code>&nbsp;为白色，且左边或者右边的格子至少一个格子为黑色，那么我们将 <code>grid[i][j]</code>&nbsp;加到最后网格图的总分中去。</p>

<p>请你返回执行任意次操作以后，最终网格图的 <strong>最大</strong>&nbsp;总分数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[0,0,0,0,0],[0,0,3,0,0],[0,1,0,0,0],[5,0,0,3,0],[0,0,0,0,2]]</span></p>

<p><span class="example-io"><b>输出：</b>11</span></p>

<p><strong>解释：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/05/11/one.png" style="width: 300px; height: 200px;" />
<p>第一次操作中，我们将第 1 列中，最上面的格子到第 3 行的格子染成黑色。第二次操作中，我们将第 4 列中，最上面的格子到最后一行的格子染成黑色。最后网格图总分为&nbsp;<code>grid[3][0] + grid[1][2] + grid[3][3]</code>&nbsp;等于 11 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[10,9,0,0,15],[7,1,0,8,0],[5,20,0,11,0],[0,0,0,1,2],[8,12,1,10,3]]</span></p>

<p><span class="example-io"><b>输出：</b>94</span></p>

<p><strong>解释：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/05/11/two-1.png" style="width: 300px; height: 200px;" />
<p>我们对第 1 ，2 ，3 列分别从上往下染黑色到第 1 ，4， 0 行。最后网格图总分为&nbsp;<code>grid[0][0] + grid[1][0] + grid[2][1] + grid[4][1] + grid[1][3] + grid[2][3] + grid[3][3] + grid[4][3] + grid[0][4]</code>&nbsp;等于 94 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;=&nbsp;n == grid.length &lt;= 100</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>0 &lt;= grid[i][j] &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

## 一、从超时做法说起

从 $\mathcal{O}(n^4)$ 的 DP 开始。

要想知道第 $j$ 列的哪些行的 $\textit{grid}[i][j]$ 被计入总分，我们需要知道：

- 第 $j+1$ 列有多少个黑色格子（下文简称为黑格）。
- 第 $j$ 列有多少个黑格。
- 第 $j-1$ 列有多少个黑格。

定义 $\textit{dfs}(j,\textit{cur},\textit{pre})$ 表示考虑第 $0$ 列到第 $j$ 列，其中第 $j+1$ 列有 $\textit{pre}$ 个黑格，第 $j$ 列有 $\textit{cur}$ 个黑格，在第 $0$ 列到第 $j$ 列中能得到最大总分。

枚举第 $j-1$ 列有 $\textit{nxt}$ 个黑格，问题变成：

- 考虑第 $0$ 列到第 $j-1$ 列，其中第 $j$ 列有 $\textit{cur}$ 个黑格，第 $j-1$ 列有 $\textit{nxt}$ 个黑格，在第 $0$ 列到第 $j-1$ 列中能得到最大总分，即 $\textit{dfs}(j-1,\textit{nxt},\textit{cur})$。

定义 $s$ 为 $\textit{grid}[\textit{cur}][j]$ 到 $\textit{grid}[\max(\textit{nxt},\textit{pre})-1][j]$ 的元素和，如果 $\max(\textit{nxt},\textit{pre}) \le \textit{cur}$ 则 $s=0$。

用 $\textit{dfs}(j-1,\textit{nxt},\textit{cur}) + s$ 更新 $\textit{dfs}(j,\textit{cur},\textit{pre})$ 的最大值。

递归边界：$j=0$ 时返回 $s$。

递归入口：$\max\limits_{i=0}^{n} \textit{dfs}(n-1,i,0)$。枚举第 $n-1$ 列有 $i$ 个黑格。

#### 复杂度分析（超时）

- 时间复杂度：$\mathcal{O}(n^4)$，其中 $n$ 是 $\textit{grid}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^3)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^4)$。
- 空间复杂度：$\mathcal{O}(n^3)$。

## 二、记忆化搜索

如果不枚举 $\textit{nxt}$，而是枚举 $\textit{cur}$ 呢？

我们的原则是，在从右往左递归的过程中，只把第 $j$ 列或者第 $j+1$ 列的格子计入总分，不考虑第 $j-1$ 列的格子。

如何**不重不漏**地统计呢？

![lc3225-cut.png](https://pic.leetcode.cn/1721570150-lvUEyU-lc3225-cut.png)

定义 $\textit{dfs}(j,\textit{pre},\textit{dec})$ 表示考虑第 $0$ 列到第 $j$ 列，其中：

- 第 $j+1$ 列有 $\textit{pre}$ 个黑格；
- 第 $j+1$ 列和第 $j+2$ 列的黑格个数的大小关系用布尔值 $\textit{dec}$ 表示，只有当第 $j+1$ 列的黑格个数小于第 $j+2$ 列的黑格个数时 $\textit{dec}$ 才为 $\texttt{true}$。

在上述约束下，第 $0$ 列到第 $j$ 列中能得到最大总分。

枚举第 $j$ 列有 $\textit{cur}$ 个黑格，按照上图中的四种情况计算。

递归边界：$j=-1$ 时返回 $0$。

递归入口：$\max\limits_{i=0}^{n} \textit{dfs}(n-2,i,0)$。枚举第 $n-1$ 列有 $i$ 个黑格。注意第 $n-1$ 列的格子会在 $j=n-2$ 中计入。


```
func maximumScore(grid [][]int) (ans int64) {
	n := len(grid)
	// 每列的前缀和（从上到下）
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	memo := make([][][2]int64, n-1)
	for i := range memo {
		memo[i] = make([][2]int64, n+1)
		for j := range memo[i] {
			memo[i][j] = [2]int64{-1, -1} // -1 表示没有计算过
		}
	}
	var dfs func(int, int, int) int64
	dfs = func(j, pre, dec int) int64 {
		if j < 0 {
			return 0
		}
		p := &memo[j][pre][dec]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := int64(0)
		// 枚举第 j 列有 cur 个黑格
		for cur := 0; cur <= n; cur++ {
			if cur == pre { // 情况一：相等
				// 没有可以计入总分的格子
				res = max(res, dfs(j-1, cur, 0))
			} else if cur < pre { // 情况二：右边黑格多
				// 第 j 列的第 [cur, pre) 行的格子可以计入总分
				res = max(res, dfs(j-1, cur, 1)+colSum[j][pre]-colSum[j][cur])
			} else if dec == 0 { // 情况三：cur > pre >= 第 j+2 列的黑格个数
				// 第 j+1 列的第 [pre, cur) 行的格子可以计入总分
				res = max(res, dfs(j-1, cur, 0)+colSum[j+1][cur]-colSum[j+1][pre])
			} else if pre == 0 { // 情况四（凹形）：cur > pre < 第 j+2 列的黑格个数
				// 此时第 j+2 列全黑最优（递归过程中一定可以枚举到这种情况）
				// 第 j+1 列全白是最优的，所以只需考虑 pre=0 的情况
				// 由于第 j+1 列在 dfs(j+1) 的情况二中已经统计过，这里不重复统计
				res = max(res, dfs(j-1, cur, 0))
			}
		}
		*p = res // 记忆化
		return res
	}

	// 枚举第 n-1 列有 i 个黑格
	for i := 0; i <= n; i++ {
		ans = max(ans, dfs(n-2, i, 0))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 是 $\textit{grid}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^3)$。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 三、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

$f[j+1][\textit{pre}][\textit{dec}]$ 的定义和 $\textit{dfs}(j,\textit{pre},\textit{dec})$ 的定义是一样的。注意这里是 $j+1$ 不是 $j$，因为要避免 $j=-1$ 时出现负数下标。

初始值 $f[0][\textit{pre}][\textit{dec}]=0$，翻译自递归边界。

答案为 $\max\limits_{i=0}^{n} f[n-1][i][0]$，翻译自递归入口。

```
func maximumScore(grid [][]int) (ans int64) {
	n := len(grid)
	// 每列的前缀和（从上到下）
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	f := make([][][2]int64, n)
	for j := range f {
		f[j] = make([][2]int64, n+1)
	}
	for j := 0; j < n-1; j++ {
		// pre 表示第 j+1 列的黑格个数
		for pre := 0; pre <= n; pre++ {
			// dec=1 意味着第 j+1 列的黑格个数 (pre) < 第 j+2 列的黑格个数
			for dec := 0; dec < 2; dec++ {
				res := int64(0)
				// 枚举第 j 列有 cur 个黑格
				for cur := 0; cur <= n; cur++ {
					if cur == pre { // 情况一：相等
						// 没有可以计入总分的格子
						res = max(res, f[j][cur][0])
					} else if cur < pre { // 情况二：右边黑格多
						// 第 j 列的第 [cur, pre) 行的格子可以计入总分
						res = max(res, f[j][cur][1]+colSum[j][pre]-colSum[j][cur])
					} else if dec == 0 { // 情况三：cur > pre >= 第 j+2 列的黑格个数
						// 第 j+1 列的第 [pre, cur) 行的格子可以计入总分
						res = max(res, f[j][cur][0]+colSum[j+1][cur]-colSum[j+1][pre])
					} else if pre == 0 { // 情况四（凹形）：cur > pre < 第 j+2 列的黑格个数
						// 此时第 j+2 列全黑最优（递归过程中一定可以枚举到这种情况）
						// 第 j+1 列全白是最优的，所以只需考虑 pre=0 的情况
						// 由于第 j+1 列在 dfs(j+1) 的情况二中已经统计过，这里不重复统计
						res = max(res, f[j][cur][0])
					}
				}
				f[j+1][pre][dec] = res
			}
		}
	}

	for _, row := range f[n-1] {
		ans = max(ans, row[0])
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 是 $\textit{grid}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 四、时间优化

把最内层的枚举 $\textit{cur}$ 的循环优化掉。

单独计算 $\textit{pre}=0$ 的状态。下面讨论 $\textit{pre}>0$ 的情况。

### 1) dec = 1

先说 $\textit{dec}=1$ 的状态。

我们相当于计算的是 $f[j][\textit{pre}][0]$ 与下式（情况二）的最大值：

$$
\begin{aligned}
& \max\limits_{\textit{cur}=0}^{\textit{pre}-1} \{f[j][\textit{cur}][1] + \textit{colSum}[j][\textit{pre}] - \textit{colSum}[j][\textit{cur}]\}      \\
={} & \textit{colSum}[j][\textit{pre}] +   \max\limits_{\textit{cur}=0}^{\textit{pre}-1} \{f[j][\textit{cur}][1] - \textit{colSum}[j][\textit{cur}]\}      \\
\end{aligned}
$$

其中

$$
\max\limits_{\textit{cur}=0}^{\textit{pre}-1} \{f[j][\textit{cur}][1] - \textit{colSum}[j][\textit{cur}]\}
$$

可以一边**从小到大**枚举 $\textit{pre}$，一边用一个变量 $\textit{preMax}$ 维护。

### 2) dec = 0

然后说 $\textit{dec}=0$ 的状态。

除了上面 $\textit{dec}=1$ 要计算的，$\textit{dec}=0$ 也要计算外，还需要计算下式（情况三）的最大值：

$$
\begin{aligned}
& \max\limits_{\textit{cur}=pre+1}^{n} \{ f[j][\textit{cur}][0] + \textit{colSum}[j + 1][\textit{cur}]  - \textit{colSum}[j + 1][\textit{pre}] \}      \\
={} & - \textit{colSum}[j + 1][\textit{pre}] +   \max\limits_{\textit{cur}=pre+1}^{n} \{ f[j][\textit{cur}][0] + \textit{colSum}[j + 1][\textit{cur}]  \}       \\
\end{aligned}
$$

其中

$$
\max\limits_{\textit{cur}=pre+1}^{n} \{ f[j][\textit{cur}][0] + \textit{colSum}[j + 1][\textit{cur}]  \}
$$

可以一边**从大到小**枚举 $\textit{pre}$，一边用一个变量 $\textit{sufMax}$ 维护。

```
func maximumScore(grid [][]int) (ans int64) {
	n := len(grid)
	colSum := make([][]int64, n)
	for j := range colSum {
		colSum[j] = make([]int64, n+1)
		for i, row := range grid {
			colSum[j][i+1] = colSum[j][i] + int64(row[j])
		}
	}

	f := make([][][2]int64, n)
	for j := range f {
		f[j] = make([][2]int64, n+1)
	}
	for j := 0; j < n-1; j++ {
		// 单独计算 pre=0 的情况
		for cur, s := range colSum[j+1] {
			f[j+1][0][0] = max(f[j+1][0][0], f[j][cur][0]+s)
			f[j+1][0][1] = max(f[j+1][0][1], f[j][cur][0])
		}

		// 用前缀最大值优化
		preMax := f[j][0][1] - colSum[j][0]
		for pre := 1; pre <= n; pre++ {
			f[j+1][pre][0] = max(f[j][pre][0], preMax+colSum[j][pre])
			f[j+1][pre][1] = f[j+1][pre][0]
			preMax = max(preMax, f[j][pre][1]-colSum[j][pre])
		}

		// 用后缀最大值优化
		sufMax := f[j][n][0] + colSum[j+1][n]
		for pre := n - 1; pre > 0; pre-- {
			f[j+1][pre][0] = max(f[j+1][pre][0], sufMax-colSum[j+1][pre])
			sufMax = max(sufMax, f[j][pre][0]+colSum[j+1][pre])
		}
	}

	for _, row := range f[n-1] {
		ans = max(ans, row[0])
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $\textit{grid}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。

注：用滚动数组可以把空间复杂度优化至 $\mathcal{O}(n)$。


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)