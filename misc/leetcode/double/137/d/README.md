#### 题目

<p>给你一个&nbsp;<code>m x n</code>&nbsp;的二维整数数组&nbsp;<code>board</code>&nbsp;，它表示一个国际象棋棋盘，其中&nbsp;<code>board[i][j]</code>&nbsp;表示格子 <code>(i, j)</code>&nbsp;的 <strong>价值</strong>&nbsp;。</p>

<p>处于 <strong>同一行</strong>&nbsp;或者 <strong>同一列</strong>&nbsp;车会互相 <strong>攻击</strong>&nbsp;。你需要在棋盘上放三个车，确保它们两两之间都&nbsp;<b>无法互相攻击</b>&nbsp;。</p>

<p>请你返回满足上述条件下，三个车所在格子 <strong>值</strong>&nbsp;之和 <strong>最大</strong>&nbsp;为多少。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>board = </span>[[-3,1,1,1],[-3,1,-3,1],[-3,2,1,1]]</p>

<p><b>输出：</b>4</p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/08/08/rooks2.png" style="width: 294px; height: 450px;" /></p>

<p>我们可以将车分别放在格子&nbsp;<code>(0, 2)</code>&nbsp;，<code>(1, 3)</code>&nbsp;和&nbsp;<code>(2, 1)</code>&nbsp;处，价值之和为&nbsp;<code>1 + 1 + 2 = 4</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>board = [[1,2,3],[4,5,6],[7,8,9]]</span></p>

<p><span class="example-io"><b>输出：</b>15</span></p>

<p><strong>解释：</strong></p>

<p>我们可以将车分别放在格子&nbsp;<code>(0, 0)</code>&nbsp;，<code>(1, 1)</code>&nbsp;和&nbsp;<code>(2, 2)</code>&nbsp;处，价值之和为&nbsp;<code>1 + 5 + 9 = 15</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>board = [[1,1,1],[1,1,1],[1,1,1]]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<p>我们可以将车分别放在格子&nbsp;<code>(0, 2)</code>&nbsp;，<code>(1, 1)</code>&nbsp;和&nbsp;<code>(2, 0)</code>&nbsp;处，价值之和为&nbsp;<code>1 + 1 + 1 = 3</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= m == board.length &lt;= 500</code></li>
	<li><code>3 &lt;= n == board[i].length &lt;= 500</code></li>
	<li><code>-10<sup>9</sup> &lt;= board[i][j] &lt;= 10<sup>9</sup></code></li>
</ul>


#### 思路

类似 [3128. 直角三角形](https://leetcode.cn/problems/right-triangles/)，有三个元素的题目，**枚举中间**比较容易。

枚举中间的车在第 $i$ 排，那么需要知道第 $0$ 到 $i-1$ 排的车放在哪最合适，以及第 $i+1$ 到 $m-1$ 排的车放在哪最合适。
> 为什么枚举中间更容易？比较一下，如果枚举第三个车，那么需要知道第一个车和第二个车的位置关系，尤其是这两个车不能同一行。但枚举第二个车，就**自动保证**了第一个车和第三个车不会同行，我们只需关注三车不能同列。

用**前后缀分解**解决。
为了保证三个车在不同的列上，我们需要计算：
- 第 $0$ 到 $i-1$ 排的车，放在三个不同的列上的最大、次大、第三大的格子值分别是多少，以及具体放在哪列。记作 $\textit{pre}[i-1]$。
- 第 $i+1$ 到 $m-1$ 排的车，放在三个不同的列上的最大、次大、第三大的格子值分别是多少，以及具体放在哪列。记作 $\textit{suf}[i+1]$。

**枚举中间**的车在第 $i$ 排第 $j$ 列，暴力枚举 $\textit{pre}[i-1]$ 和 $\textit{suf}[i+1]$ 的最大、次大、第三大的组合，只要三个车在不同的列上，就用格子值之和更新答案的最大值。
代码实现时，可以只预处理 $\textit{suf}$，对于 $\textit{pre}$，可以在枚举中间的车的同时计算。

```
func maximumValueSum(board [][]int) int64 {
	m := len(board)
	type pair struct{ x, j int }
	suf := make([][3]pair, m)
	p := [3]pair{} // 最大、次大、第三大
	for i := range p {
		p[i].x = math.MinInt
	}
	update := func(row []int) {
		for j, x := range row {
			if x > p[0].x {
				if p[0].j != j { // 如果相等，仅更新最大
					if p[1].j != j { // 如果相等，仅更新最大和次大
						p[2] = p[1]
					}
					p[1] = p[0]
				}
				p[0] = pair{x, j}
			} else if x > p[1].x && j != p[0].j {
				if p[1].j != j { // 如果相等，仅更新次大
					p[2] = p[1]
				}
				p[1] = pair{x, j}
			} else if x > p[2].x && j != p[0].j && j != p[1].j {
				p[2] = pair{x, j}
			}
		}
	}
	for i := m - 1; i > 1; i-- {
		update(board[i])
		suf[i] = p
	}

	ans := math.MinInt
	for i := range p {
		p[i].x = math.MinInt // 重置，计算 pre
	}
	for i, row := range board[:m-2] {
		update(row)
		for j, x := range board[i+1] { // 第二个车
			for _, p := range p { // 第一个车
				for _, q := range suf[i+2] { // 第三个车
					if p.j != j && q.j != j && p.j != q.j { // 没有同列的车
						ans = max(ans, p.x+x+q.x)
						break
					}
				}
			}
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{board}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m)$。

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