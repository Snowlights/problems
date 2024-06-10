#### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和一个 <strong>非负</strong>&nbsp;整数&nbsp;<code>k</code>&nbsp;。</p>

<p>一次操作中，你可以选择任一下标&nbsp;<code>i</code>&nbsp;，然后将&nbsp;<code>nums[i]</code>&nbsp;加&nbsp;<code>1</code>&nbsp;或者减&nbsp;<code>1</code>&nbsp;。</p>

<p>请你返回将 <code>nums</code>&nbsp;<strong>中位数</strong>&nbsp;变为 <code>k</code>&nbsp;所需要的 <strong>最少</strong>&nbsp;操作次数。</p>

<p>一个数组的 <strong>中位数</strong>&nbsp;指的是数组按 <strong>非递减</strong> 顺序排序后最中间的元素。如果数组长度为偶数，我们选择中间两个数的较大值为中位数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,5,6,8,5], k = 4</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><b>解释：</b>我们将&nbsp;<code>nums[1]</code> 和&nbsp;<code>nums[4]</code>&nbsp;减 <code>1</code>&nbsp;得到&nbsp;<code>[2, 4, 6, 8, 4]</code>&nbsp;。现在数组的中位数等于&nbsp;<code>k</code>&nbsp;。所以答案为 2 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,5,6,8,5], k = 7</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><b>解释：</b>我们将&nbsp;<code>nums[1]</code>&nbsp;增加 1 两次，并且将&nbsp;<code>nums[2]</code>&nbsp;增加 1 一次，得到&nbsp;<code>[2, 7, 7, 8, 5]</code>&nbsp;。结果数组的中位数等于&nbsp;<code>k</code>&nbsp;。所以答案为 3 。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3,4,5,6], k = 4</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><b>解释：</b>数组中位数已经等于 <code>k</code>&nbsp;了，所以不需要进行任何操作。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

把 $\textit{nums}$ 从小到大排序后，中位数为 $\textit{nums}[m]$，其中 $m=\left\lfloor\dfrac{n}{2}\right\rfloor$，$n$ 为 $\textit{nums}$ 的长度。

我们需要把中位数左边的数都变成 $\le k$ 的，右边的数都变成 $\ge k$ 的。

分类讨论：
- 如果 $\textit{nums}[m] > k$，要把下标在 $[0,m]$ 中的大于 $k$ 的数都变成 $k$，由于下标在 $[m+1,n-1]$ 中的数已经大于 $k$（因为数组是有序的），所以下标在 $[m+1,n-1]$ 中的数无需操作。
- 如果 $\textit{nums}[m] < k$，要把下标在 $[m,n-1]$ 中的小于 $k$ 的数都变成 $k$，由于下标在 $[0,m-1]$ 中的数已经小于 $k$（因为数组是有序的），所以下标在 $[0,m-1]$ 中的数无需操作。

累加元素的变化量，即为答案。

``` go
func minOperationsToMakeMedianK(nums []int, k int) (ans int64) {
	sort.Ints(nums)
	m := len(nums) / 2
	if nums[m] > k {
		for i := m; i >= 0 && nums[i] > k; i-- {
			ans += int64(nums[i] - k)
		}
	} else {
		for i := m; i < len(nums) && nums[i] < k; i++ {
			ans += int64(k - nums[i])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。用快速选择算法可以做到期望 $\mathcal{O}(n)$ 的时间复杂度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

