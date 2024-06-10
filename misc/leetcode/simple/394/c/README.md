#### 题目

<p>给你一个大小为 <code>m x n</code>&nbsp;的二维矩形&nbsp;<code>grid</code>&nbsp;。每次 <strong>操作</strong>&nbsp;中，你可以将 <strong>任一</strong> 格子的值修改为 <strong>任意</strong>&nbsp;非负整数。完成所有操作后，你需要确保每个格子&nbsp;<code>grid[i][j]</code>&nbsp;的值满足：</p>

<ul>
	<li>如果下面相邻格子存在的话，它们的值相等，也就是&nbsp;<code>grid[i][j] == grid[i + 1][j]</code>（如果存在）。</li>
	<li>如果右边相邻格子存在的话，它们的值不相等，也就是&nbsp;<code>grid[i][j] != grid[i][j + 1]</code>（如果存在）。</li>
</ul>

<p>请你返回需要的 <strong>最少</strong>&nbsp;操作数目。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1,0,2],[1,0,2]]</span></p>

<p><b>输出：</b>0</p>

<p><b>解释：</b></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/04/15/examplechanged.png" style="width: 254px; height: 186px;padding: 10px; background: #fff; border-radius: .5rem;" /></strong></p>

<p>矩阵中所有格子已经满足要求。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1,1,1],[0,0,0]]</span></p>

<p><b>输出：</b>3</p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/03/27/example21.png" style="width: 254px; height: 186px;padding: 10px; background: #fff; border-radius: .5rem;" /></strong></p>

<p>将矩阵变成&nbsp;<code>[[1,0,1],[1,0,1]]</code>&nbsp;，它满足所有要求，需要 3 次操作：</p>

<ul>
	<li>将&nbsp;<code>grid[1][0]</code>&nbsp;变为 1 。</li>
	<li>将&nbsp;<code>grid[0][1]</code> 变为 0 。</li>
	<li>将&nbsp;<code>grid[1][2]</code>&nbsp;变为 1 。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>grid = [[1],[2],[3]]</span></p>

<p><b>输出：</b>2</p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/03/31/changed.png" style="width: 86px; height: 277px;padding: 10px; background: #fff; border-radius: .5rem;" /></p>

<p>这个矩阵只有一列，我们可以通过 2 次操作将所有格子里的值变为 1 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n, m &lt;= 1000</code></li>
	<li><code>0 &lt;= grid[i][j] &lt;= 9</code></li>
</ul>

#### 思路

记忆化搜索

反向思考，计算**保留多少个数不变**。用 $mn$ 减去最多保留不变的元素个数，就是最少需要修改的元素个数。
首先统计每一列的每种元素的出现次数到 $\textit{cnt}$ 数组中，其中 $\textit{cnt}[i][x]$ 表示第 $i$ 列元素 $x$ 的出现次数。
对于第 $i$ 列，我们需要知道上一列都变成了什么数，比如上一列都变成了 $1$，那么第 $i$ 列就必须都变成不等于 $1$ 的数，比如 $2$，那么第 $i$ 列就可以保留 $\textit{cnt}[i][2]$ 个元素不变。
枚举每一列变成哪个数，就可以枚举到所有的情况。
定义 $\textit{dfs}(i,j)$ 表示考虑前 $i$ 列，且第 $i+1$ 列都变成 $j$ 时的最大保留不变元素个数。
用「枚举选哪个」思考。枚举第 $i$ 列的元素都变成 $k$，则有

$$
\textit{dfs}(i,j) = \max\limits_{k} \textit{dfs}(i-1,k) + \textit{cnt}[i][k]
$$

在 $[0,9]$ 中枚举不等于 $j$ 的 $k$，取最大值即为 $\textit{dfs}(i,j)$。
递归边界：$\textit{dfs}(-1,j) = 0$。
答案：$mn - \textit{dfs}(n-1,10)$。注意初始 $j$ 是一个不在 $[0,9]$ 中的数。

``` go
func minimumOperations(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	cnt := make([][10]int, m)
	for _, v := range grid {
		for j, vv := range v {
			cnt[j][vv]++
		}
	}

	memo := make([][11]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, pre int) int
	dfs = func(i, pre int) int {
		if i == m {
			return 0
		}
		dv := &memo[i][pre]
		if *dv >= 0 {
			return *dv
		}
		res := 0
		defer func() {
			*dv = res
		}()

		for j := 0; j < 10; j++ {
			if j == pre {
				continue
			}
			res = max(res, dfs(i+1, j)+cnt[i][j])
		}

		return res
	}

	return n*m - dfs(0, 10)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn + nU^2)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数，$U$ 为元素种类数，即 $10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nU)$，单个状态的计算时间为 $\mathcal{O}(U)$，所以动态规划的时间复杂度为 $\mathcal{O}(nU^2)$。
- 空间复杂度：$\mathcal{O}(mn + nU)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

