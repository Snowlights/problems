### 题目

<p>给你两个长度都为 <code>n</code>&nbsp;的整数数组&nbsp;<code>arr</code> 和&nbsp;<code>brr</code>&nbsp;以及一个整数&nbsp;<code>k</code>&nbsp;。你可以对 <code>arr</code>&nbsp;执行以下操作任意次：</p>

<ul>
	<li>将&nbsp;<code>arr</code>&nbsp;分割成若干个&nbsp;<strong>连续的</strong>&nbsp;子数组，并将这些子数组按任意顺序重新排列。这个操作的代价为&nbsp;<code>k</code>&nbsp;。</li>
	<li>
	<p>选择 <code>arr</code>&nbsp;中的任意一个元素，将它增加或者减少一个正整数&nbsp;<code>x</code>&nbsp;。这个操作的代价为 <code>x</code>&nbsp;。</p>
	</li>
</ul>

<p>请你返回将 <code>arr</code>&nbsp;变为 <code>brr</code>&nbsp;的 <strong>最小</strong>&nbsp;总代价。</p>

<p><strong>子数组</strong>&nbsp;是一个数组中一段连续 <strong>非空</strong>&nbsp;的元素序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>arr = [-7,9,5], brr = [7,-2,-5], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>13</span></p>

<p><b>解释：</b></p>

<ul>
	<li>将&nbsp;<code>arr</code>&nbsp;分割成两个连续子数组：<code>[-7]</code> 和&nbsp;<code>[9, 5]</code>&nbsp;然后将它们重新排列成&nbsp;<code>[9, 5, -7]</code>&nbsp;，代价为 2 。</li>
	<li>将&nbsp;<code>arr[0]</code>&nbsp;减小 2 ，数组变为&nbsp;<code>[7, 5, -7]</code>&nbsp;，操作代价为 2 。</li>
	<li>将&nbsp;<code>arr[1]</code>&nbsp;减小 7 ，数组变为&nbsp;<code>[7, -2, -7]</code>&nbsp;，操作代价为 7 。</li>
	<li>将&nbsp;<code>arr[2]</code>&nbsp;增加 2 ，数组变为&nbsp;<code>[7, -2, -5]</code>&nbsp;，操作代价为 2 。</li>
</ul>

<p>将两个数组变相等的总代价为&nbsp;<code>2 + 2 + 7 + 2 = 13</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>arr = [2,1], brr = [2,1], k = 0</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><b>解释：</b></p>

<p>由于数组已经相等，不需要进行任何操作，所以总代价为 0 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= arr.length == brr.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= k &lt;= 2 * 10<sup>10</sup></code></li>
	<li><code>-10<sup>5</sup> &lt;= arr[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>5</sup> &lt;= brr[i] &lt;= 10<sup>5</sup></code></li>
</ul>

### 思路

把 $\textit{arr}$ 简记为 $a$，$\textit{brr}$ 简记为 $b$。

如果不使用操作一，那么答案为所有 $|a[i]-b[i]|$ 之和。

如果使用操作一，那么直接把 $a$ 分成 $n$ 个长为 $1$ 的子数组，这样 $a$ 就可以随意排了。

最优配对方式是最小的 $a[i]$ 与最小的 $b[i]$ 一对，次小的 $a[i]$ 与次小的 $b[i]$ 一对。用交换论证法可以证明这样做是最优的。

```
func minCost(a, b []int, k int64) int64 {
	ans2 := int64(0)
	for i, x := range a {
		ans2 += int64(abs(x - b[i]))
	}

	slices.Sort(a)
	slices.Sort(b)
	ans1 := k
	for i, x := range a {
		ans1 += int64(abs(x - b[i]))
	}

	return min(ans1, ans2)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $a$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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