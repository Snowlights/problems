#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的 <code>m x n</code> 整数矩阵 <code>grid</code> 。你一开始的位置在 <strong>左上角</strong> 格子 <code>(0, 0)</code> 。</p>

<p>当你在格子 <code>(i, j)</code> 的时候，你可以移动到以下格子之一：</p>

<ul>
	<li>满足 <code>j < k <= grid[i][j] + j</code> 的格子 <code>(i, k)</code> （向右移动），或者</li>
	<li>满足 <code>i < k <= grid[i][j] + i</code> 的格子 <code>(k, j)</code> （向下移动）。</li>
</ul>

<p>请你返回到达 <strong>右下角</strong> 格子 <code>(m - 1, n - 1)</code> 需要经过的最少移动格子数，如果无法到达右下角格子，请你返回 <code>-1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/01/25/ex1.png" style="width: 271px; height: 171px;"/></p>

<pre><b>输入：</b>grid = [[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]
<b>输出：</b>4
<b>解释：</b>上图展示了到达右下角格子经过的 4 个格子。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/01/25/ex2.png" style="width: 271px; height: 171px;"/></p>

<pre><b>输入：</b>grid = [[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]]
<b>输出：</b>3
<strong>解释：</strong>上图展示了到达右下角格子经过的 3 个格子。
</pre>

<p><strong>示例 3：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/01/26/ex3.png" style="width: 181px; height: 81px;"/></p>

<pre><b>输入：</b>grid = [[2,1,0],[1,0,0]]
<b>输出：</b>-1
<b>解释：</b>无法到达右下角格子。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 <= m, n <= 10<sup>5</sup></code></li>
	<li><code>1 <= m * n <= 10<sup>5</sup></code></li>
	<li><code>0 <= grid[i][j] < m * n</code></li>
	<li><code>grid[m - 1][n - 1] == 0</code></li>
</ul>

#### 思路

暴力做法是从 $(0,0)$ 出发，向右/向下尝试移动到每个满足要求的格子。假设移动到了 $(i,j)$，
那么问题就变成从 $(i,j)$ 出发，到右下角的最少移动格子数。这是一个和原问题相似的子问题，启发我们用递归来思考。
看了上面的视频后，你知道递归中有重叠子问题，从而可以用记忆化搜索来优化，而记忆化搜索又可以翻译成如下的递推。
定义 $f[i][j]$ 表示从 $(i,j)$ 出发，到右下角最少移动格子数。
设 $g=\textit{grid}[i][j]$，有

$$
f[i][j] = \min\left\{\min_{k=j+1}^{j+g} f[i][k], \min_{k=i+1}^{i+g} f[k][j]\right\} + 1

$$

$i$ 和 $j$ 均倒序遍历。答案为 $f[0][0]$。
但这样做时间复杂度是 $O(mn(m+n))$ 的，无法接受。
由于有「区间查询」、「单点更新」这两个操作，我们可以用线段树来优化。
但还有更「轻量级」的做法。
对于 $\min\limits_{k=j+1}^{j+g} f[i][k]$ 来说，在倒序遍历 $j$ 时，$k$ 的左边界 $j+1$ 是在**单调减小**的，
我们可以用一个 $f$ 值底小顶大的单调栈来维护 $f[i][k]$ 及其下标 $k$。由于是倒序遍历，
单调栈中的下标是底大顶小的，从那么在单调栈上二分查找最大的不超过 $j+g$ 的下标 $k$，
对应的 $f[i][k]$ 就是 $[j+1, j+g]$ 范围内的最小值。
对于 $\min\limits_{k=i+1}^{i+g} f[k][j]$ 也同理，每一列都需要一个单调栈。

```go  
func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colSt := make([][]pair, n) // 每列的单调栈
	for i := m - 1; i >= 0; i-- {
		st := []pair{} // 当前行的单调栈
		for j := n - 1; j >= 0; j-- {
			st2 := colSt[j]
			mn = math.MaxInt32
			if i == m-1 && j == n-1 { // 特殊情况：已经是终点
				mn = 0
			} else if g := grid[i][j]; g > 0 {
				// 在单调栈上二分
				k := j + g
				k = sort.Search(len(st), func(i int) bool { return st[i].i <= k })
				if k < len(st) {
					mn = min(mn, st[k].x)
				}
				k = i + g
				k = sort.Search(len(st2), func(i int) bool { return st2[i].i <= k })
				if k < len(st2) {
					mn = min(mn, st2[k].x)
				}
			}

			if mn < math.MaxInt32 {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(st) > 0 && mn <= st[len(st)-1].x {
					st = st[:len(st)-1]
				}
				st = append(st, pair{mn, j})
				for len(st2) > 0 && mn <= st2[len(st2)-1].x {
					st2 = st2[:len(st2)-1]
				}
				colSt[j] = append(st2, pair{mn, i})
			}
		}
	}
	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt32 {
		return -1
	}
	return
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$O(mn\log(mn))$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$O(mn)$。
