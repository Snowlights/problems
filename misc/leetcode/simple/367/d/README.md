#### 题目

<p>给你一个下标从 <strong>0</strong> 开始、大小为 <code>n * m</code> 的二维整数矩阵 <code><font face="monospace">grid</font></code><font face="monospace"> ，定义一个下标从 <strong>0</strong> 开始、大小为 <code>n * m</code> 的的二维矩阵</font> <code>p</code>。如果满足以下条件，则称 <code>p</code> 为 <code>grid</code> 的 <strong>乘积矩阵</strong> ：</p>

<ul>
	<li>对于每个元素 <code>p[i][j]</code> ，它的值等于除了 <code>grid[i][j]</code> 外所有元素的乘积。乘积对 <code>12345</code> 取余数。</li>
</ul>

<p>返回 <code>grid</code> 的乘积矩阵。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>grid = [[1,2],[3,4]]
<strong>输出：</strong>[[24,12],[8,6]]
<strong>解释：</strong>p[0][0] = grid[0][1] * grid[1][0] * grid[1][1] = 2 * 3 * 4 = 24
p[0][1] = grid[0][0] * grid[1][0] * grid[1][1] = 1 * 3 * 4 = 12
p[1][0] = grid[0][0] * grid[0][1] * grid[1][1] = 1 * 2 * 4 = 8
p[1][1] = grid[0][0] * grid[0][1] * grid[1][0] = 1 * 2 * 3 = 6
所以答案是 [[24,12],[8,6]] 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>grid = [[12345],[2],[1]]
<strong>输出：</strong>[[2],[0],[0]]
<strong>解释：</strong>p[0][0] = grid[0][1] * grid[0][2] = 2 * 1 = 2
p[0][1] = grid[0][0] * grid[0][2] = 12345 * 1 = 12345. 12345 % 12345 = 0 ，所以 p[0][1] = 0
p[0][2] = grid[0][0] * grid[0][1] = 12345 * 2 = 24690. 24690 % 12345 = 0 ，所以 p[0][2] = 0
所以答案是 [[2],[0],[0]] 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == grid.length&nbsp;&lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= m == grid[i].length&nbsp;&lt;= 10<sup>5</sup></code></li>
	<li><code>2 &lt;= n * m &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= grid[i][j] &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

**核心思想**：把矩阵拉成一维的，我们需要算出每个数左边所有数的乘积，以及右边所有数的乘积，这都可以用递推得到。

先算出从 $\textit{grid}[i][j]$ 的下一个元素开始，到最后一个元素 $\textit{grid}[n-1][m-1]$ 的乘积，记作 $\textit{suf}[i][j]$。这可以从最后一行最后一列开始，倒着遍历得到。

然后算出从第一个元素 $\textit{grid}[0][0]$ 开始，到 $\textit{grid}[i][j]$ 的上一个元素的乘积，记作 $\textit{pre}[i][j]$。这可以从第一行第一列开始，正着遍历得到。

那么

$$
p[i][j] = \textit{pre}[i][j]\cdot \textit{suf}[i][j]
$$

代码实现时，可以先初始化 $p[i][j]=\textit{suf}[i][j]$，然后把 $\textit{pre}[i][j]$ 乘到 $\textit{p}[i][j]$ 中，就得到了答案。这样 $\textit{pre}$ 和 $\textit{suf}$ 就可以压缩成一个变量。

关于取模的原理，见文末的「算法小课堂」。

```go [sol-Go]
func constructProductMatrix(grid [][]int) [][]int {
	const mod = 12345
	n, m := len(grid), len(grid[0])
	p := make([][]int, n)
	suf := 1 // 后缀乘积
	for i := n - 1; i >= 0; i-- {
		p[i] = make([]int, m)
		for j := m - 1; j >= 0; j-- {
			p[i][j] = suf // p[i][j] 先初始化成后缀乘积
			suf = suf * grid[i][j] % mod
		}
	}

	pre := 1 // 前缀乘积
	for i, row := range grid {
		for j, x := range row {
			p[i][j] = p[i][j] * pre % mod // 然后再乘上前缀乘积
			pre = pre * x % mod
		}
	}
	return p
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 和 $m$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 练习：前后缀分解（右边的数字为难度分）

- [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/) 和本题几乎一样
- [2256. 最小平均差](https://leetcode.cn/problems/minimum-average-difference/) 1395
- [2483. 商店的最少代价](https://leetcode.cn/problems/minimum-penalty-for-a-shop/) 1495
- [2420. 找到所有好下标](https://leetcode.cn/problems/find-all-good-indices/) 1695
- [2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/) 2219
- [2484. 统计回文子序列数目](https://leetcode.cn/problems/count-palindromic-subsequences/) 2223
- [2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/) 2432
- [2552. 统计上升四元组](https://leetcode.cn/problems/count-increasing-quadruplets/) 2433
- [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)

