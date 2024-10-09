#### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;，请你返回&nbsp;<code>nums</code>&nbsp;中有多少个<span data-keyword="subarray-nonempty">子数组</span>满足：子数组中所有元素按位&nbsp;<code>AND</code>&nbsp;的结果为 <code>k</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,1,1], k = 1</span></p>

<p><span class="example-io"><b>输出：</b>6</span></p>

<p><strong>解释：</strong></p>

<p>所有子数组都只含有元素 1 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,1,2], k = 1</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><b>解释：</b></p>

<p>按位&nbsp;<code>AND</code>&nbsp;值为 1 的子数组包括：<code>[<u><strong>1</strong></u>,1,2]</code>, <code>[1,<u><strong>1</strong></u>,2]</code>, <code>[<u><strong>1,1</strong></u>,2]</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>按位&nbsp;<code>AND</code>&nbsp;值为 2 的子数组包括：<code>[1,<b><u>2</u></b>,3]</code>, <code>[1,<u><strong>2,3</strong></u>]</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i], k &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

怎么计算子数组的 AND？

首先，我们有如下 $\mathcal{O}(n^2)$ 的暴力算法：
从左到右正向遍历 $\textit{nums}$，对于 $x=\textit{nums}[i]$，从 $i-1$ 开始倒着遍历 $\textit{nums}[j]$，更新 $\textit{nums}[j]=\textit{nums}[j]\&x$。
- $i=1$ 时，我们会把 $\textit{nums}[0]$ 到 $\textit{nums}[1]$ 的 AND 记录在 $\textit{nums}[0]$ 中。
- $i=2$ 时，我们会把 $\textit{nums}[1]$ 到 $\textit{nums}[2]$ 的 AND 记录在 $\textit{nums}[1]$ 中，$\textit{nums}[0]$ 到 $\textit{nums}[2]$ 的 AND 记录在 $\textit{nums}[0]$ 中。
- $i=3$ 时，我们会把 $\textit{nums}[2]$ 到 $\textit{nums}[3]$ 的 AND 记录在 $\textit{nums}[2]$ 中；$\textit{nums}[1]$ 到 $\textit{nums}[3]$ 的 AND 记录在 $\textit{nums}[1]$ 中；$\textit{nums}[0]$ 到 $\textit{nums}[3]$ 的 AND 记录在 $\textit{nums}[0]$ 中。
- 按照该算法，可以计算出所有子数组的 AND。注意单个元素也算子数组。

下面来优化该算法。

前置知识：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

把二进制数看成集合，两个数的 AND 就是两个集合的**交集**。
对于两个二进制数 $a$ 和 $b$，如果 $a\&b = a$，从集合的角度上看，$a$ 对应的集合是 $b$ 对应的集合的子集。或者说，$b$ 对应的集合是 $a$ 对应的集合的**超集**。
据此我们可以提出如下优化：
仍然是从左到右正向遍历 $\textit{nums}$，对于 $x=\textit{nums}[i]$，从 $i-1$ 开始倒着遍历 $\textit{nums}[j]$：

- 如果 $\textit{nums}[j]\&x\ne\textit{nums}[j]$，说明 $\textit{nums}[j]$ 可以变小（求交集后，集合元素只会减少不会变多），更新 $\textit{nums}[j]=\textit{nums}[j]\&x$。
- 否则 $\textit{nums}[j]\&x=\textit{nums}[j]$，从集合的角度看，此时 $x$ 不仅是 $\textit{nums}[j]$ 的超集，同时也是 $\textit{nums}[k]\ (k<j)$ 的超集（因为前面的循环保证了每个集合都是其左侧相邻集合的超集），在 $A\subseteq B$ 的前提下，$A\cap B=A$，所以后续的循环都不会改变元素值，**退出内层循环**。

## 写法一：二分查找

由于每个元素都是其右侧元素的子集，所以从 $\textit{nums}[0]$ 到 $\textit{nums}[i]$ 的元素值是非降的。既然是有序数组，我们可以在 $[0,i]$ 中**二分查找** $k$，做法见 [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)。
设左闭右开区间 $[\textit{left},\textit{right})$ 是 $\textit{nums}[j]=k$ 的 $j$ 的范围。把左闭右开区间的长度 $\textit{right}-\textit{left}$ 加入答案。


```
func countSubarrays(nums []int, k int) (ans int64) {
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		a := nums[:i+1]
		ans += int64(sort.SearchInts(a, k+1) - sort.SearchInts(a, k))
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U + n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于 $2^{29}-1<10^9<2^{30}-1$，二进制数对应集合的大小不会超过 $29$，因此在 AND 运算下，每个数字至多可以减小 $29$ 次。总体上看（除去二分），二重循环的总循环次数等于每个数字可以减小的次数之和，即 $O(n\log U)$。
- 空间复杂度：$\mathcal{O}(1)$。


## 写法二：三指针

由于元素值只会减少，所以当 $i$ 增大时，$\textit{left}$ 和 $\textit{right}$ 不会减少，有了单调性的保证，上面的二分查找可以改成 [滑动窗口](https://www.bilibili.com/video/BV1hd4y1r7Gq/)：
- 当 $\textit{left}\le i$ 且 $\textit{nums}[\textit{left}] < k$ 时，把 $\textit{left}$ 加一。
- 当 $\textit{right}\le i$ 且 $\textit{nums}[\textit{right}] \le k$ 时，把 $\textit{right}$ 加一。
- 把左闭右开区间的长度 $\textit{right}-\textit{left}$ 加入答案。

```
func countSubarrays(nums []int, k int) (ans int64) {
	left, right := 0, 0
	for i, x := range nums {
		for j := i - 1; j >= 0 && nums[j]&x != nums[j]; j-- {
			nums[j] &= x
		}
		for left <= i && nums[left] < k {
			left++
		}
		for right <= i && nums[right] <= k {
			right++
		}
		ans += int64(right - left)
	}
	return
}
```


#### 复杂度分析
- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。由于 $\textit{left}$ 和 $\textit{right}$ 只会增大，不会减少，所以 $\textit{left}$ 和 $\textit{right}$ 的移动次数之和为 $\mathcal{O}(n)$。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)