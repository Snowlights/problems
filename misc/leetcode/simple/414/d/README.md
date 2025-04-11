#### 题目

<p>给你一个&nbsp;<code>50 x 50</code>&nbsp;的国际象棋棋盘，棋盘上有 <strong>一个</strong>&nbsp;马和一些兵。给你两个整数&nbsp;<code>kx</code> 和&nbsp;<code>ky</code>&nbsp;，其中&nbsp;<code>(kx, ky)</code>&nbsp;表示马所在的位置，同时还有一个二维数组&nbsp;<code>positions</code>&nbsp;，其中&nbsp;<code>positions[i] = [x<sub>i</sub>, y<sub>i</sub>]</code>&nbsp;表示第 <code>i</code>&nbsp;个兵在棋盘上的位置。</p>

<p>Alice 和 Bob 玩一个回合制游戏，Alice 先手。玩家的一次操作中，可以执行以下操作：</p>

<ul>
	<li>玩家选择一个仍然在棋盘上的兵，然后移动马，通过 <strong>最少</strong>&nbsp;的 <strong>步数</strong> 吃掉这个兵。<strong>注意</strong>&nbsp;，玩家可以选择&nbsp;<strong>任意</strong>&nbsp;一个兵，<strong>不一定</strong>&nbsp;要选择从马的位置出发&nbsp;<strong>最少</strong>&nbsp;移动步数的兵。</li>
	<li><span>在马吃兵的过程中，马 <strong>可能</strong>&nbsp;会经过一些其他兵的位置，但这些兵 <strong>不会</strong>&nbsp;被吃掉。<strong>只有</strong>&nbsp;选中的兵在这个回合中被吃掉。</span></li>
</ul>

<p>Alice 的目标是 <strong>最大化</strong>&nbsp;两名玩家的 <strong>总</strong>&nbsp;移动次数，直到棋盘上不再存在兵，而 Bob 的目标是 <strong>最小化</strong>&nbsp;总移动次数。</p>

<p>假设两名玩家都采用 <strong>最优</strong>&nbsp;策略，请你返回 Alice 可以达到的 <strong>最大</strong>&nbsp;总移动次数。</p>

<p>在一次&nbsp;<strong>移动</strong>&nbsp;中，如下图所示，马有 8 个可以移动到的位置，每个移动位置都是沿着坐标轴的一个方向前进 2 格，然后沿着垂直的方向前进 1 格。</p>

<p><img src="https://assets.leetcode.com/uploads/2024/08/01/chess_knight.jpg" style="width: 275px; height: 273px;" /></p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>kx = 1, ky = 1, positions = [[0,0]]</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/08/16/gif3.gif" style="width: 275px; height: 275px;" /></p>

<p>马需要移动 4 步吃掉&nbsp;<code>(0, 0)</code>&nbsp;处的兵。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>kx = 0, ky = 2, positions = [[1,1],[2,2],[3,3]]</span></p>

<p><span class="example-io"><b>输出：</b>8</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/08/16/gif4.gif" style="width: 320px; height: 320px;" /></strong></p>

<ul>
	<li>Alice 选择&nbsp;<code>(2, 2)</code>&nbsp;处的兵，移动马吃掉它需要 2 步：<code>(0, 2) -&gt; (1, 4) -&gt; (2, 2)</code>&nbsp;。</li>
	<li>Bob 选择&nbsp;<code>(3, 3)</code>&nbsp;处的兵，移动马吃掉它需要 2 步：<code>(2, 2) -&gt; (4, 1) -&gt; (3, 3)</code>&nbsp;。</li>
	<li>Alice 选择&nbsp;<code>(1, 1)</code>&nbsp;处的兵，移动马吃掉它需要 4 步：<code>(3, 3) -&gt; (4, 1) -&gt; (2, 2) -&gt; (0, 3) -&gt; (1, 1)</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>kx = 0, ky = 0, positions = [[1,2],[2,4]]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>Alice 选择&nbsp;<code>(2, 4)</code>&nbsp;处的兵，移动马吃掉它需要 2 步：<code>(0, 0) -&gt; (1, 2) -&gt; (2, 4)</code>&nbsp;。注意，<code>(1, 2)</code>&nbsp;处的兵不会被吃掉。</li>
	<li>Bob 选择&nbsp;<code>(1, 2)</code>&nbsp;处的兵，移动马吃掉它需要 1 步：<code>(2, 4) -&gt; (1, 2)</code>&nbsp;。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= kx, ky &lt;= 49</code></li>
	<li><code>1 &lt;= positions.length &lt;= 15</code></li>
	<li><code>positions[i].length == 2</code></li>
	<li><code>0 &lt;= positions[i][0], positions[i][1] &lt;= 49</code></li>
	<li><code>positions[i]</code>&nbsp;两两互不相同。</li>
	<li>输入保证对于所有&nbsp;<code>0 &lt;= i &lt; positions.length</code>&nbsp;，都有&nbsp;<code>positions[i] != [kx, ky]</code>&nbsp;。</li>
</ul>

#### 思路

## 引入

考虑以下两种情况：

- Alice 吃掉第一个兵，Bob 吃掉第二个兵，Alice 吃掉第三个兵，现在轮到 Bob 操作。
- Alice 吃掉第二个兵，Bob 吃掉第一个兵，Alice 吃掉第三个兵，现在轮到 Bob 操作。

这两种情况，Bob 面临的局面是完全一样的：前三个兵都被吃掉，马现在在第三个兵的位置。

有重复的子问题，就可以用记忆化搜索/递推解决。

## 状态和状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前马在第 $i$ 个兵的位置。特别地，把 $(\textit{kx}, \textit{ky})$ 当作第 $n+1$ 个兵的位置。
- $\textit{mask}$：剩余没有被吃掉的兵的集合。

因此，定义状态为 $\textit{dfs}(i,\textit{mask})$，表示当前马在第 $i$ 个兵的位置，且剩余没有被吃掉的兵的集合为 $\textit{mask}$ 的情况下，继续游戏，两名玩家的总移动次数的最大值。

> 注意题目要计算的是两名玩家的总移动次数，不是 Alice 一个人的总移动次数。

接下来，思考如何从一个状态转移到另一个状态。

设从 $(x,y)$ 移动到第 $i$ 个兵的最小步数为 $\textit{dis}[i][x][y]$，这可以用网格图 BFS 算出来：反向思考，计算从第 $i$ 个兵的位置出发，通过「马走日」移动到 $(x,y)$ 的最小步数。

设当前位置为 $(x,y) = \textit{positions}[i]$，考虑枚举吃掉第 $j$ 个兵：

- 如果第 $j$ 个兵在集合 $\textit{mask}$ 中，把马移动 $\textit{dis}[j][x][y]$ 步，吃掉第 $j$ 个兵。现在问题变成当前马在第 $j$ 个兵的位置，且剩余没有被吃掉的兵的集合为 $\textit{mask}\setminus \{j\}$ 的情况下，继续游戏，两名玩家的总移动次数的最大值，即 $\textit{dfs}(j,\textit{mask}\setminus \{j\})$。

如果当前是 Alice 操作，则有

$$
\textit{dfs}(i,\textit{mask}) = \max_{j\in \textit{mask}}  \textit{dfs}(j,\textit{mask}\setminus \{j\}) + \textit{dis}[j][x][y]
$$

如果当前是 Bob 操作，则有

$$
\textit{dfs}(i,\textit{mask}) = \min_{j\in \textit{mask}}  \textit{dfs}(j,\textit{mask}\setminus \{j\}) + \textit{dis}[j][x][y]
$$

如何判断当前是谁在操作？

可以添加一个状态 $\textit{curPlayer}$，也可以用已被吃掉的兵的集合 $\complement_U \textit{mask}$ 的大小的奇偶性来判断，其中全集 $\textit{U}={0,1,2,\cdots,n-1}$。

- 如果吃掉了 $0,2,4,\cdots$ 个兵，那么当前轮到 Alice 操作。
- 如果吃掉了 $1,3,5,\cdots$ 个兵，那么当前轮到 Bob 操作。

递归边界：$\textit{dfs}(i,\varnothing) = 0$。所有兵都被吃掉，游戏结束。

递归入口：$\textit{dfs}(n,U)$。即答案。

代码中用到了一些位运算技巧，原理见 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)


## 记忆化搜索

```
func maxMoves(kx, ky int, positions [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	n := len(positions)
	// 计算马到兵的步数，等价于计算兵到其余格子的步数
	dis := make([][50][50]int, n)
	for i, pos := range positions {
		d := &dis[i]
		for j := range d {
			for k := range d[j] {
				d[j][k] = -1
			}
		}
		px, py := pos[0], pos[1]
		d[px][py] = 0
		q := []pair{{px, py}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, dir := range dirs {
					x, y := p.x+dir.x, p.y+dir.y
					if 0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0 {
						d[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
	}

	positions = append(positions, []int{kx, ky})
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, 1<<n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	u := 1<<n - 1
	var dfs func(int, int) int
	dfs = func(i, mask int) int {
		if mask == 0 {
			return 0
		}
		p := &memo[i][mask]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 0
		x, y := positions[i][0], positions[i][1]
		if bits.OnesCount(uint(u^mask))%2 == 0 { // Alice 操作
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				res = max(res, dfs(j, mask^1<<j)+dis[j][x][y])
			}
		} else { // Bob 操作
			res = math.MaxInt
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				res = min(res, dfs(j, mask^1<<j)+dis[j][x][y])
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(n, u)
}
```

## 递推

注意要先把 $\textit{mask}$ 小的状态算出来，再算 $\textit{mask}$ 大的。所以外层循环 $\textit{mask}$，内层循环 $i$。

```
func maxMoves(kx, ky int, positions [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}
	n := len(positions)
	// 计算马到兵的步数，等价于计算兵到其余格子的步数
	dis := make([][50][50]int, n)
	for i, pos := range positions {
		d := &dis[i]
		for j := range d {
			for k := range d[j] {
				d[j][k] = -1
			}
		}
		px, py := pos[0], pos[1]
		d[px][py] = 0
		q := []pair{{px, py}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, dir := range dirs {
					x, y := p.x+dir.x, p.y+dir.y
					if 0 <= x && x < 50 && 0 <= y && y < 50 && d[x][y] < 0 {
						d[x][y] = step
						q = append(q, pair{x, y})
					}
				}
			}
		}
	}

	positions = append(positions, []int{kx, ky})
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for mask := 1; mask < 1<<n; mask++ {
		for i, p := range positions {
			x, y := p[0], p[1]
			odd := bits.OnesCount(uint(u^mask))%2 > 0
			if odd {
				f[mask][i] = math.MaxInt
			}
			op := func(a, b int) int {
				if odd {
					return min(a, b)
				}
				return max(a, b)
			}
			for s := uint(mask); s > 0; s &= s - 1 {
				j := bits.TrailingZeros(s)
				f[mask][i] = op(f[mask][i], f[mask^1<<j][j]+dis[j][x][y])
			}
		}
	}
	return f[u][n]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nL^2 + n^2 2^n)$，其中 $n$ 是 $\textit{positions}$ 的长度，$L=50$。
- 空间复杂度：$\mathcal{O}(nL^2 + n2^n)$。

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