### 题目

<p>给你一个长度为 <code>n</code> 的数组 <code>original</code> 和一个长度为 <code>n x 2</code> 的二维数组 <code>bounds</code>，其中 <code>bounds[i] = [u<sub>i</sub>, v<sub>i</sub>]</code>。</p>

<p>你需要找到长度为 <code>n</code>&nbsp;且满足以下条件的&nbsp;<strong>可能的&nbsp;</strong>数组 <code>copy</code> 的数量：</p>

<ol>
	<li>对于 <code>1 &lt;= i &lt;= n - 1</code>&nbsp;，都有&nbsp;<code>(copy[i] - copy[i - 1]) == (original[i] - original[i - 1])</code>&nbsp;。</li>
	<li>对于 <code>0 &lt;= i &lt;= n - 1</code>&nbsp;，都有&nbsp;<code>u<sub>i</sub> &lt;= copy[i] &lt;= v<sub>i</sub></code><sub>&nbsp;</sub>。</li>
</ol>

<p>返回满足这些条件的数组数目。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">original = [1,2,3,4], bounds = [[1,2],[2,3],[3,4],[4,5]]</span></p>

<p><strong>输出：</strong><span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>可能的数组为：</p>

<ul>
	<li><code>[1, 2, 3, 4]</code></li>
	<li><code>[2, 3, 4, 5]</code></li>
</ul>
</div>

<p><strong class="example">示例 2</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">original = [1,2,3,4], bounds = [[1,10],[2,9],[3,8],[4,7]]</span></p>

<p><strong>输出：</strong><span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>可能的数组为：</p>

<ul>
	<li><code>[1, 2, 3, 4]</code></li>
	<li><code>[2, 3, 4, 5]</code></li>
	<li><code>[3, 4, 5, 6]</code></li>
	<li><code>[4, 5, 6, 7]</code></li>
</ul>
</div>

<p><strong class="example">示例 3</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">original = [1,2,1,2], bounds = [[1,1],[2,3],[3,3],[2,3]]</span></p>

<p><strong>输出：</strong><span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<p>没有可行的数组。</p>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>2 &lt;= n == original.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= original[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>bounds.length == n</code></li>
	<li><code>bounds[i].length == 2</code></li>
	<li><code>1 &lt;= bounds[i][0] &lt;= bounds[i][1] &lt;= 10<sup>9</sup></code></li>
</ul>

### 思路

题目要求 $\textit{copy}[i] - \textit{copy}[i - 1] = \textit{original}[i] - \textit{original}[i - 1]$，那么有

$$
\begin{aligned}
\textit{copy}[1] - \textit{copy}[0] &= \textit{original}[1] - \textit{original}[0]     \\
\textit{copy}[2] - \textit{copy}[1] &= \textit{original}[2] - \textit{original}[1]     \\
\textit{copy}[3] - \textit{copy}[2] &= \textit{original}[3] - \textit{original}[2]     \\
&\ \ \vdots \\
\textit{copy}[i] - \textit{copy}[i - 1] &= \textit{original}[i] - \textit{original}[i - 1] \\
\end{aligned}
$$

累加等号左边的所有项，累加等号右边的所有项，得

$$
\textit{copy}[i] - \textit{copy}[0] = \textit{original}[i] - \textit{original}[0]
$$

移项得

$$
\textit{copy}[i] = \textit{copy}[0] + \textit{original}[i] - \textit{original}[0]
$$

换句话说，确定了 $\textit{copy}[0]$，那么整个数组也就确定了。所以 $\textit{copy}[0]$ 的取值范围（整数集合）的大小就是答案。

题目要求

$$
u_i\le \textit{copy}[i] \le v_i
$$

设 $d_i = \textit{original}[i] - \textit{original}[0]$，用 $\textit{copy}[i] = \textit{copy}[0] + d_i$ 替换上式中的 $\textit{copy}[i]$，得

$$
u_i\le \textit{copy}[0] + d_i \le v_i
$$

移项得

$$
u_i - d_i \le \textit{copy}[0] \le v_i - d_i
$$

所以我们可以得到 $n$ 个关于 $\textit{copy}[0]$ 的不等式，或者说区间：

$$
\begin{aligned}
& [u_0,v_0] \\
& [u_1 - d_1 ,v_1-d_1] \\
& [u_2 - d_2 ,v_2-d_2] \\
& \ \ \vdots \\
& [u_{n-1} - d_{n-1} ,v_{n-1}-d_{n-1}] \\
& \end{aligned}
$$

这些区间的**交集**，即为 $\textit{copy}[0]$ 能取到的值。

交集的大小即为答案。如果交集为空，返回 $0$。

怎么算交集的范围？所有区间左端点取最大值，右端点取最小值。

```
func countArrays(original []int, bounds [][]int) int {
	mn, mx := math.MinInt, math.MaxInt
	for i, b := range bounds {
		d := original[i] - original[0]
		mn = max(mn, b[0]-d) // 计算区间交集
		mx = min(mx, b[1]-d)
	}
	return max(mx-mn+1, 0) // 注意交集可能是空的
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{original}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)