#### 题目

<p>给你一个&nbsp;<code>m x n</code>&nbsp;的二进制矩形&nbsp;<code>grid</code>&nbsp;和一个整数&nbsp;<code>health</code>&nbsp;表示你的健康值。</p>

<p>你开始于矩形的左上角&nbsp;<code>(0, 0)</code>&nbsp;，你的目标是矩形的右下角&nbsp;<code>(m - 1, n - 1)</code>&nbsp;。</p>

<p>你可以在矩形中往上下左右相邻格子移动，但前提是你的健康值始终是 <b>正数</b>&nbsp;。</p>

<p>对于格子&nbsp;<code>(i, j)</code>&nbsp;，如果&nbsp;<code>grid[i][j] = 1</code>&nbsp;，那么这个格子视为 <strong>不安全</strong>&nbsp;的，会使你的健康值减少 1 。</p>

<p>如果你可以到达最终的格子，请你返回&nbsp;<code>true</code>&nbsp;，否则返回 <code>false</code>&nbsp;。</p>

<p><b>注意</b>&nbsp;，当你在最终格子的时候，你的健康值也必须为<strong>&nbsp;正数</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><b>解释：</b></p>

<p>沿着下图中灰色格子走，可以安全到达最终的格子。</p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_1drawio.png" style="width: 301px; height: 121px;" /></div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3</span></p>

<p><span class="example-io"><b>输出：</b>false</span></p>

<p><b>解释：</b></p>

<p>健康值最少为 4 才能安全到达最后的格子。</p>
<img alt="" src="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_2drawio.png" style="width: 361px; height: 161px;" /></div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><b>解释：</b></p>

<p>沿着下图中灰色格子走，可以安全到达最终的格子。</p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/08/04/3868_examples_3drawio.png" style="width: 181px; height: 121px;" /></p>

<p>任何不经过格子&nbsp;<code>(1, 1)</code>&nbsp;的路径都是不安全的，因为你的健康值到达最终格子时，都会小于等于 0 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 50</code></li>
<li><code>2 <= m * n</code></li>
	<li><code>1 &lt;= health &lt;= m + n</code></li>
	<li><code>grid[i][j]</code>&nbsp;要么是 0 ，要么是 1 。</li>
</ul>

#### 思路

本题和 [2290. 到达角落需要移除障碍物的最小数目](https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/) 是一样的。~~曾经的困难题，现在的中等题~~

从 $(i,j)$ 移动到与其相邻的格子 $(x,y)$，视作一条从 $(i,j)$ 到 $(x,y)$ 的有向边，边权为 $\textit{grid}[x][y]$。

问题变成计算从起点到终点的最短路。

这可以用 Dijkstra 算法解决，原理请看 [Dijkstra 算法介绍](https://leetcode.cn/problems/network-delay-time/solution/liang-chong-dijkstra-xie-fa-fu-ti-dan-py-ooe8/)。

由于本题的边权只有 $0$ 和 $1$，也可以用 **0-1 BFS** 解决。

0-1 BFS 本质是对 Dijkstra 算法的优化。因为边权只有 $0$ 和 $1$，我们可以把最小堆换成**双端队列**，遇到 $0$ 边权就加入**队首**，遇到 $1$ 边权就加入**队尾**，这样可以保证队首总是最小的，就不需要最小堆了。


```
func findSafeWalk(grid [][]int, health int) bool {
	type pair struct{ x, y int }
	dirs := []pair{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	m, n := len(grid), len(grid[0])
	dis := make([][]int, m)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			dis[i][j] = math.MaxInt
		}
	}
	dis[0][0] = grid[0][0]
	q := [2][]pair{{{}}} // 两个 slice 头对头来实现 deque
	for len(q[0]) > 0 || len(q[1]) > 0 {
		var p pair
		if len(q[0]) > 0 {
			p, q[0] = q[0][len(q[0])-1], q[0][:len(q[0])-1]
		} else {
			p, q[1] = q[1][0], q[1][1:]
		}
		for _, d := range dirs {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x < m && 0 <= y && y < n {
				g := grid[x][y]
				if dis[p.x][p.y]+g < dis[x][y] {
					dis[x][y] = dis[p.x][p.y] + g
					q[g] = append(q[g], pair{x, y})
				}
			}
		}
	}
	return dis[m-1][n-1] < health
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。每个点至多入队两次。
- 空间复杂度：$\mathcal{O}(mn)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)