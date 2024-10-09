### 题目

<p>给你一个二进制数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>你可以对数组执行以下操作&nbsp;<strong>任意</strong>&nbsp;次（也可以 0 次）：</p>

<ul>
	<li>选择数组中 <strong>任意</strong>&nbsp;一个下标 <code>i</code>&nbsp;，并将从下标 <code>i</code>&nbsp;开始一直到数组末尾 <strong>所有</strong>&nbsp;元素 <strong>反转</strong>&nbsp;。</li>
</ul>

<p><b>反转</b>&nbsp;一个元素指的是将它的值从 0 变 1 ，或者从 1 变 0 。</p>

<p>请你返回将 <code>nums</code>&nbsp;中所有元素变为 1 的 <strong>最少</strong>&nbsp;操作次数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [0,1,1,0,1]</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong><br />
我们可以执行以下操作：</p>

<ul>
	<li>选择下标&nbsp;<code>i = 1</code>&nbsp;执行操作，得到<span class="example-io">&nbsp;<code>nums = [0,<u><strong>0</strong></u>,<u><strong>0</strong></u>,<u><strong>1</strong></u>,<u><strong>0</strong></u>]</code>&nbsp;。</span></li>
	<li>选择下标&nbsp;<code>i = 0</code>&nbsp;执行操作，得到<span class="example-io">&nbsp;<code>nums = [<u><strong>1</strong></u>,<u><strong>1</strong></u>,<u><strong>1</strong></u>,<u><strong>0</strong></u>,<u><strong>1</strong></u>]</code>&nbsp;。</span></li>
	<li>选择下标&nbsp;<code>i = 4</code>&nbsp;执行操作，得到<span class="example-io">&nbsp;<code>nums = [1,1,1,0,<u><strong>0</strong></u>]</code>&nbsp;。</span></li>
	<li>选择下标&nbsp;<code>i = 3</code>&nbsp;执行操作，得到<span class="example-io">&nbsp;<code>nums = [1,1,1,<u><strong>1</strong></u>,<u><strong>1</strong></u>]</code>&nbsp;。</span></li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,0,0,0]</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong><br />
我们可以执行以下操作：</p>

<ul>
	<li>选择下标&nbsp;<code>i = 1</code>&nbsp;执行操作，得到<span class="example-io">&nbsp;<code>nums = [1,<u><strong>1</strong></u>,<u><strong>1</strong></u>,<u><strong>1</strong></u>]</code>&nbsp;。</span></li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 1</code></li>
</ul>

### 思路

由于 $\textit{nums}[i]$ 会被其左侧元素的操作影响，所以我们先从最左边的 $\textit{nums}[0]$ 开始思考。

分类讨论：

- 如果 $\textit{nums}[0]=1$，无需反转，问题变成剩下 $n-1$ 个数如何操作。接下来考虑 $\textit{nums}[1]$。
- 如果 $\textit{nums}[0]=0$，反转次数加一，问题变成剩下 $n-1$ 个数（在修改次数是奇数的情况下）如何操作。接下来考虑 $\textit{nums}[1]$。

对后续元素来说，由于反转偶数次等于没反转，所以只需考虑反转次数的奇偶性。

一般地，设反转次数的奇偶性为 $k$，分类讨论：

- 如果 $\textit{nums}[i]\ne k$，无需反转，接下来考虑 $\textit{nums}[i+1]$。
- 如果 $\textit{nums}[i]=k$，反转次数加一，接下来考虑 $\textit{nums}[i+1]$。

```
func minOperations(nums []int) (ans int) {
	for _, x := range nums {
		if x == ans%2 {
			ans++
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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