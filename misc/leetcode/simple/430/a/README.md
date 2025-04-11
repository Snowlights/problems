#### 题目

<p>给你一个由&nbsp;<b>非负&nbsp;</b>整数组成的 <code>m x n</code> 矩阵 <code>grid</code>。</p>

<p>在一次操作中，你可以将任意元素 <code>grid[i][j]</code> 的值增加 1。</p>

<p>返回使 <code>grid</code> 的所有列&nbsp;<strong>严格递增&nbsp;</strong>所需的&nbsp;<strong>最少&nbsp;</strong>操作次数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">grid = [[3,2],[1,3],[3,4],[0,1]]</span></p>

<p><strong>输出:</strong> <span class="example-io">15</span></p>

<p><strong>解释:</strong></p>
<ul>
	<li>为了让第 <code>0</code>&nbsp;列严格递增，可以对 <code>grid[1][0]</code> 执行 3 次操作，对 <code>grid[2][0]</code> 执行 2 次操作，对 <code>grid[3][0]</code> 执行 6 次操作。</li>
	<li>为了让第 <code>1</code>&nbsp;列严格递增，可以对 <code>grid[3][1]</code> 执行 4 次操作。</li>
</ul>
<img alt="" src="https://assets.leetcode.com/uploads/2024/11/10/firstexample.png" style="width: 200px; height: 347px;" /></div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">grid = [[3,2,1],[2,1,0],[1,2,3]]</span></p>

<p><strong>输出:</strong> <span class="example-io">12</span></p>

<p><strong>解释:</strong></p>
<ul>
	<li>为了让第 <code>0</code>&nbsp;列严格递增，可以对 <code>grid[1][0]</code> 执行 2 次操作，对 <code>grid[2][0]</code> 执行 4 次操作。</li>
	<li>为了让第 <code>1</code>&nbsp;列严格递增，可以对 <code>grid[1][1]</code> 执行 2 次操作，对 <code>grid[2][1]</code> 执行 2 次操作。</li>
	<li>为了让第 <code>2</code>&nbsp;列严格递增，可以对 <code>grid[1][2]</code> 执行 2 次操作。</li>
</ul>
<img alt="" src="https://assets.leetcode.com/uploads/2024/11/10/secondexample.png" style="width: 300px; height: 257px;" /></div>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>
<ul>
	<li><code>m == grid.length</code></li>
	<li><code>n == grid[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 50</code></li>
	<li><code>0 &lt;= grid[i][j] &lt; 2500</code></li>
</ul>

#### 思路

由于我们只能变大，不能变小，那么第一个数肯定不需要变。

对于后面的每个数 $x$：
- 如果它比前一个数 $\textit{pre}$ 大，那么不变。
- 否则，$x$ 至少要增大到 $\textit{pre}+1$，才能保持严格递增。增大到恰好等于 $\textit{pre}+1$ 是最优的（不然后面的数需要变得更大），操作 $\textit{pre}+1-x$ 次。

```
func minimumOperations(nums []int) int {
	seen := map[int]bool{}
	for i, x := range slices.Backward(nums) {
		if seen[x] {
			return i/3 + 1
		}
		seen[x] = true
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid} 的行数和列数。。
- 空间复杂度：$\mathcal{O}(1)$。


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)