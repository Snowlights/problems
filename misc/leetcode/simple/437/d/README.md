#### 题目

<p>给你一个大小为 <code>n x m</code> 的二维整数矩阵 <code>grid</code>，其中每个元素的值为 <code>0</code>、<code>1</code> 或 <code>2</code>。</p>

<p><strong>V 形对角线段</strong> 定义如下：</p>

<ul>
	<li>线段从&nbsp;<code>1</code> 开始。</li>
	<li>后续元素按照以下无限序列的模式排列：<code>2, 0, 2, 0, ...</code>。</li>
	<li>该线段：
	<ul>
		<li>起始于某个对角方向（左上到右下、右下到左上、右上到左下或左下到右上）。</li>
		<li>沿着相同的对角方向继续，保持&nbsp;<strong>序列模式&nbsp;</strong>。</li>
		<li>在保持&nbsp;<strong>序列模式&nbsp;</strong>的前提下，最多允许&nbsp;<strong>一次顺时针 90 度转向&nbsp;</strong>另一个对角方向。</li>
	</ul>
	</li>
</ul>

<p><img alt="" src="https://pic.leetcode.cn/1739609732-jHpPma-length_of_longest3.jpg" style="width: 481px; height: 202px;" /></p>

<p>返回最长的&nbsp;<strong>V 形对角线段&nbsp;</strong>的&nbsp;<strong>长度&nbsp;</strong>。如果不存在有效的线段，则返回 0。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[2,2,1,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://pic.leetcode.cn/1739609768-rhePxN-matrix_1-2.jpg" style="width: 201px; height: 192px;" /></p>

<p>最长的 V 形对角线段长度为 5，路径如下：<code>(0,2) → (1,3) → (2,4)</code>，在 <code>(2,4)</code> 处进行&nbsp;<strong>顺时针 90 度转向&nbsp;</strong>，继续路径为 <code>(3,3) → (4,2)</code>。</p>
</div>

<p><strong>示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[2,2,2,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://pic.leetcode.cn/1739609774-nYJElV-matrix_2.jpg" style="width: 201px; height: 201px;" /></p>

<p>最长的 V 形对角线段长度为 4，路径如下：<code>(2,3) → (3,2)</code>，在 <code>(3,2)</code> 处进行&nbsp;<strong>顺时针 90 度转向&nbsp;</strong>，继续路径为 <code>(2,1) → (1,0)</code>。</p>
</div>

<p><strong>示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1,2,2,2,2],[2,2,2,2,0],[2,0,0,0,0],[0,0,2,2,2],[2,0,0,2,0]]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://pic.leetcode.cn/1739609780-tlkdUW-matrix_3.jpg" style="width: 201px; height: 201px;" /></p>

<p>最长的 V 形对角线段长度为 5，路径如下：<code>(0,0) → (1,1) → (2,2) → (3,3) → (4,4)</code>。</p>
</div>

<p><strong>示例 4：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">1</span></p>

<p><strong>解释：</strong></p>

<p>最长的 V 形对角线段长度为 1，路径如下：<code>(0,0)</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == grid.length</code></li>
	<li><code>m == grid[i].length</code></li>
	<li><code>1 &lt;= n, m &lt;= 500</code></li>
	<li><code>grid[i][j]</code> 的值为 <code>0</code>、<code>1</code> 或 <code>2</code>。</li>
</ul>

#### 思路

想象有一个人在网格图上移动，按照题目要求，移动经过的值必须为 $1,2,0,2,0,\cdots$。

由于移动路径不会形成环（否则路径会有多个 $1$，或者说拐弯了不止一次），我们可以写一个递归，来模拟人在网格图上的移动。

定义 $\textit{dfs}(i,j,k,\textit{canTurn},\textit{target})$ 表示在如下约束下的最长移动步数。
- **上一步**在 $(i,j)$。\n- 移动方向为 $\textit{DIRS}[k]$，其中 $\textit{DIRS}$ 是一个长为 $4$ 的方向数组。
- 是否可以右转，用布尔值 $\textit{canTurn}$ 表示。
- 当前位置的目标值必须等于 $\textit{target}$。

**递归边界**：
- 出界，返回 $0$。
- 如果 $\textit{grid}[i'][j']\ne \textit{target}$，返回 $0$。其中 $(i',j')$ 是从 $(i,j)$ 向 $\textit{DIRS}[k]$ 方向移动一步后的位置。

**递归入口**：
- 如果 $\textit{grid}[i][j]=1$，那么枚举 $k=0,1,2,3$，递归 $\textit{dfs}(i,j,k,\texttt{true},2)$。其中 $2$ 是因为下一步的值必须是 $2$。

计算所有 $\textit{dfs}(i,j,k,\texttt{true},2)+1$ 的最大值，即为答案。

⚠**注意**：$\textit{target}$ 无需记忆化，因为知道 $(i,j)$ 就间接知道 $\textit{target}$ 是多少，代码只是为了方便实现，额外传入了 $\textit{target}$

``` 
var DIRS = [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func lenOfVDiagonal(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	memo := make([][][4][2]int, m)
	for i := range memo {
		memo[i] = make([][4][2]int, n)
	}

	var dfs func(int, int, int, int, int) int
	dfs = func(i, j, k, canTurn, target int) (res int) {
		i += DIRS[k][0]
		j += DIRS[k][1]
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != target {
			return
		}
		p := &memo[i][j][k][canTurn]
		if *p > 0 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		res = dfs(i, j, k, canTurn, 2-target) // 直行
		if canTurn == 1 {
			res = max(res, dfs(i, j, (k+1)%4, 0, 2-target)) // 右转
		}
		return res + 1 // 算上当前位置
	}

	for i, row := range grid {
		for j, x := range row {
			if x == 1 {
				for k := range 4 { // 枚举起始方向
					ans = max(ans, dfs(i, j, k, 1, 2)+1)
				}
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。保存多少状态，就需要多少空间。

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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
