#### 题目

<p>给你一个 <code>m x n</code> 的二维数组 <code>grid</code>，数组由&nbsp;<strong>正整数</strong> 组成。</p>

<p>你的任务是以&nbsp;<strong>之字形&nbsp;</strong>遍历 <code>grid</code>，同时跳过每个&nbsp;<strong>交替&nbsp;</strong>的单元格。</p>

<p>之字形遍历的定义如下：</p>

<ul>
	<li>从左上角的单元格 <code>(0, 0)</code> 开始。</li>
	<li>在当前行中向 <strong>右</strong> 移动，直到到达该行的末尾。</li>
	<li>下移到下一行，然后在该行中向&nbsp;<strong>左</strong><em>&nbsp;</em>移动，直到到达该行的开头。</li>
	<li>继续在行间交替向右和向左移动，直到所有行都被遍历完。</li>
</ul>

<p><strong>注意：</strong>在遍历过程中，必须跳过每个&nbsp;<strong>交替&nbsp;</strong>的单元格。</p>

<p>返回一个整数数组 <code>result</code>，其中包含按&nbsp;<strong>顺序&nbsp;</strong>记录的、且跳过交替单元格后的之字形遍历中访问到的单元格值。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1,2],[3,4]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[1,4]</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/11/23/4012_example0.png" style="width: 200px; height: 200px;" /></strong></p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[2,1],[2,1],[2,1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[2,1,2]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/11/23/4012_example1.png" style="width: 200px; height: 240px;" /></p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1,2,3],[4,5,6],[7,8,9]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[1,3,5,7,9]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/11/23/4012_example2.png" style="width: 260px; height: 250px;" /></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n == grid.length &lt;= 50</code></li>
	<li><code>2 &lt;= m == grid[i].length &lt;= 50</code></li>
	<li><code>1 &lt;= grid[i][j] &lt;= 2500</code></li>
</ul>

#### 思路

用一个变量 $\textit{ok}$ 表示当前数字能否加入答案。
初始值为 $\texttt{true}$，每遍历一个数就取反，这样我们可以选一个数，跳过一个数，选一个数，跳过一个数，……
对于下标为奇数的行，倒序遍历（或者将其反转）。

``` 
func zigzagTraversal(grid [][]int) (ans []int) {
	ok := true
	for i, row := range grid {
		if i%2 > 0 {
			slices.Reverse(row)
		}
		for _, x := range row {
			if ok {
				ans = append(ans, x)
			}
			ok = !ok
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

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