#### 题目

<p>给你一个大小为&nbsp;<code>n x n</code> 的整数方阵 <code>grid</code>。返回一个经过如下调整的矩阵：</p>

<ul>
	<li><strong>左下角三角形</strong>（包括中间对角线）的对角线按&nbsp;<strong>非递增顺序&nbsp;</strong>排序。</li>
	<li><strong>右上角三角形&nbsp;</strong>的对角线按&nbsp;<strong>非递减顺序&nbsp;</strong>排序。</li>
</ul>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1,7,3],[9,8,2],[4,5,6]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[[8,2,3],[9,6,7],[4,5,1]]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/12/29/4052example1drawio.png" style="width: 461px; height: 181px;" /></p>

<p>标有黑色箭头的对角线（左下角三角形）应按非递增顺序排序：</p>

<ul>
	<li><code>[1, 8, 6]</code> 变为 <code>[8, 6, 1]</code>。</li>
	<li><code>[9, 5]</code> 和 <code>[4]</code> 保持不变。</li>
</ul>

<p>标有蓝色箭头的对角线（右上角三角形）应按非递减顺序排序：</p>

<ul>
	<li><code>[7, 2]</code> 变为 <code>[2, 7]</code>。</li>
	<li><code>[3]</code> 保持不变。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[0,1],[1,2]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[[2,1],[1,0]]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/12/29/4052example2adrawio.png" style="width: 383px; height: 141px;" /></p>

<p>标有黑色箭头的对角线必须按非递增顺序排序，因此 <code>[0, 2]</code> 变为 <code>[2, 0]</code>。其他对角线已经符合要求。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">grid = [[1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[[1]]</span></p>

<p><strong>解释：</strong></p>

<p>只有一个元素的对角线已经符合要求，因此无需修改。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>grid.length == grid[i].length == n</code></li>
	<li><code>1 &lt;= n &lt;= 10</code></li>
	<li><code>-10<sup>5</sup> &lt;= grid[i][j] &lt;= 10<sup>5</sup></code></li>
</ul>


#### 思路

对于每条对角线，行号 $i$ 减列号 $j$ 是一个定值。（回想一下 [51. N 皇后](https://leetcode.cn/problems/n-queens/) 的写法）

设 $k=i-j+n$，那么右上角那条对角线的 $k=1$，左下角那条对角线的 $k=m+n-1$。（本题 $m=n$）
于是枚举 $k=1,2,3,\ldots,m+n-1$，就相当于在从右上到左下，一条一条地枚举对角线。
由于 $i = k+j-n$，知道 $j$ 就知道 $i$，所以我们只需要计算出每条对角线的 $j$ 的最小值和最大值，就可以开始遍历对角线了。

- 由于 $j=i-k+n$，当 $i=0$ 的时候 $j$ 取到最小值 $n-k$，但这个数不能是负数，所以最小的 $j$ 是 $\max(n-k,0)$。
- 由于 $j=i-k+n$，当 $i=m-1$ 的时候 $j$ 取到最大值 $m + n - 1 - k$，但这个数不能超过 $n-1$，所以最大的 $j$ 是 $\min(m + n - 1 - k, n - 1)$。

然后就可以模拟了：

1. 把对角线的元素存入列表 $a$ 中。
2. 如果最小的 $j$ 大于 $0$，说明我们在主对角线右上，升序排序；否则降序排序。
3. 把 $a$ 按顺序原样放回对角线中。


```
func sortMatrix(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	// 第一排在右上，最后一排在左下
	// 每排从左上到右下
	// 令 k=i-j+n，那么右上角 k=1，左下角 k=m+n-1
	for k := 1; k < m+n; k++ {
		// 核心：计算 j 的最小值和最大值
		minJ := max(n-k, 0)       // i=0 的时候，j=n-k，但不能是负数
		maxJ := min(m+n-1-k, n-1) // i=m-1 的时候，j=m+n-1-k，但不能超过 n-1
		a := []int{}
		for j := minJ; j <= maxJ; j++ {
			a = append(a, grid[k+j-n][j]) // 根据 k 的定义得 i=k+j-n
		}
		if minJ > 0 { // 右上角三角形
			slices.Sort(a)
		} else { // 左下角三角形（包括中间对角线）
			slices.SortFunc(a, func(a, b int) int { return b - a })
		}
		for j := minJ; j <= maxJ; j++ {
			grid[k+j-n][j] = a[j-minJ]
		}
	}
	return grid
}
```

#### 复杂度分析

-  时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{grid}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(m+n)$。


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