### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的<span data-keyword="prime">质数</span>数组&nbsp;<code>nums</code>&nbsp;。你的任务是返回一个长度为 <code>n</code>&nbsp;的数组 <code>ans</code>&nbsp;，对于每个下标 <code>i</code>&nbsp;，以下<strong>&nbsp;条件</strong>&nbsp;均成立：</p>

<ul>
	<li><code>ans[i] OR (ans[i] + 1) == nums[i]</code></li>
</ul>

<p>除此以外，你需要 <strong>最小化</strong>&nbsp;结果数组里每一个&nbsp;<code>ans[i]</code>&nbsp;。</p>

<p>如果没法找到符合 <strong>条件</strong>&nbsp;的&nbsp;<code>ans[i]</code>&nbsp;，那么&nbsp;<code>ans[i] = -1</code>&nbsp;。</p>

<p><strong>质数</strong>&nbsp;指的是一个大于 1 的自然数，且它只有 1 和自己两个因数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,3,5,7]</span></p>

<p><span class="example-io"><b>输出：</b>[-1,1,4,3]</span></p>

<p><b>解释：</b></p>

<ul>
	<li>对于&nbsp;<code>i = 0</code>&nbsp;，不存在&nbsp;<code>ans[0]</code>&nbsp;满足&nbsp;<code>ans[0] OR (ans[0] + 1) = 2</code>&nbsp;，所以&nbsp;<code>ans[0] = -1</code>&nbsp;。</li>
	<li>对于&nbsp;<code>i = 1</code>&nbsp;，满足 <code>ans[1] OR (ans[1] + 1) = 3</code>&nbsp;的最小&nbsp;<code>ans[1]</code>&nbsp;为&nbsp;<code>1</code>&nbsp;，因为&nbsp;<code>1 OR (1 + 1) = 3</code>&nbsp;。</li>
	<li>对于&nbsp;<code>i = 2</code>&nbsp;，满足 <code>ans[2] OR (ans[2] + 1) = 5</code>&nbsp;的最小 <code>ans[2]</code>&nbsp;为&nbsp;<code>4</code>&nbsp;，因为&nbsp;<code>4 OR (4 + 1) = 5</code>&nbsp;。</li>
	<li>对于&nbsp;<code>i = 3</code>&nbsp;，满足&nbsp;<code>ans[3] OR (ans[3] + 1) = 7</code>&nbsp;的最小&nbsp;<code>ans[3]</code>&nbsp;为&nbsp;<code>3</code>&nbsp;，因为&nbsp;<code>3 OR (3 + 1) = 7</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [11,13,31]</span></p>

<p><span class="example-io"><b>输出：</b>[9,12,15]</span></p>

<p><b>解释：</b></p>

<ul>
	<li>对于&nbsp;<code>i = 0</code>&nbsp;，满足&nbsp;<code>ans[0] OR (ans[0] + 1) = 11</code> 的最小&nbsp;<code>ans[0]</code>&nbsp;为&nbsp;<code>9</code>&nbsp;，因为&nbsp;<code>9 OR (9 + 1) = 11</code>&nbsp;。</li>
	<li>对于&nbsp;<code>i = 1</code>&nbsp;，满足&nbsp;<code>ans[1] OR (ans[1] + 1) = 13</code>&nbsp;的最小&nbsp;<code>ans[1]</code>&nbsp;为&nbsp;<code>12</code>&nbsp;，因为&nbsp;<code>12 OR (12 + 1) = 13</code>&nbsp;。</li>
	<li>对于&nbsp;<code>i = 2</code>&nbsp;，满足&nbsp;<code>ans[2] OR (ans[2] + 1) = 31</code>&nbsp;的最小&nbsp;<code>ans[2]</code>&nbsp;为&nbsp;<code>15</code>&nbsp;，因为&nbsp;<code>15 OR (15 + 1) = 31</code>&nbsp;。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>2 &lt;= nums[i] &lt;= 1000</code></li>
	<li><code>nums[i]</code>&nbsp;是一个质数。</li>
</ul>


### 思路

本文中出现的数字均为二进制数。

例如 $x=100111$，那么 $x\ |\ (x+1) = 100111\ |\ 101000 = 101111$。

可以发现，$x\ |\ (x+1)$ 的本质是把二进制最右边的 $0$ 置为 $1$。

反过来，如果我们知道了 $x\ |\ (x+1)$ 的结果 $101111$，那么对应的 $x$ 只能是这些：

- $100111$。
- $101011$。
- $101101$。
- $101110$。

其中最小的是 $100111$，也就是把 $101111$ 最右边的 $0$ 的右边的 $1$ 置为 $0$。

由于 $x\ |\ (x+1)$ 最低位一定是 $1$（讨论 $x$ 是奇数和偶数的情况可以证明），所以如果 $\textit{nums}[i]$ 是偶数（质数中只有 $2$），那么无解，答案为 $-1$。

## 写法一

把 $101111$ 取反，得 $010000$，其 $\text{lowbit}=10000$，右移一位得 $1000$。把 $101111$ 与 $1000$ 异或，即可得到 $100111$。

关于 $\text{lowbit}$ 的原理，请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)


```
func minBitwiseArray(nums []int) []int {
	for i, x := range nums {
		if x == 2 {
			nums[i] = -1
		} else {
			t := ^x
			nums[i] ^= t & -t >> 1
		}
	}
	return nums
}
```

## 写法二

把 $101111$ 加一，得 $110000$，再 AND $101111$ 取反后的值 $010000$，可以得到方法一中的 $\text{lowbit}=10000$。

```
func minBitwiseArray(nums []int) []int {
	for i, x := range nums {
		if x == 2 {
			nums[i] = -1
		} else {
			nums[i] ^= (x + 1) &^ x >> 1
		}
	}
	return nums
}
```

#### 复杂度分析

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
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)