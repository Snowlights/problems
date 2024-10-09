#### 题目

<p>给你一个整数数组 <code>nums</code> 。数组 <code>nums</code> 的<strong> 唯一性数组</strong> 是一个按元素从小到大排序的数组，包含了 <code>nums</code> 的所有<span data-keyword="subarray-nonempty">非空子数组中</span>不同元素的个数。</p>

<p>换句话说，这是由所有 <code>0 &lt;= i &lt;= j &lt; nums.length</code> 的 <code>distinct(nums[i..j])</code> 组成的递增数组。</p>

<p>其中，<code>distinct(nums[i..j])</code> 表示从下标 <code>i</code> 到下标 <code>j</code> 的子数组中不同元素的数量。</p>

<p>返回 <code>nums</code> <strong>唯一性数组 </strong>的 <strong>中位数 </strong>。</p>

<p><strong>注意</strong>，数组的 <strong>中位数 </strong>定义为有序数组的中间元素。如果有两个中间元素，则取值较小的那个。<!-- notionvc: 7e0f5178-4273-4a82-95ce-3395297921dc --></p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [1,2,3]</span></p>

<p><strong>输出：</strong><span class="example-io">1</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code> 的唯一性数组为 <code>[distinct(nums[0..0]), distinct(nums[1..1]), distinct(nums[2..2]), distinct(nums[0..1]), distinct(nums[1..2]), distinct(nums[0..2])]</code>，即 <code>[1, 1, 1, 2, 2, 3]</code> 。唯一性数组的中位数为 1 ，因此答案是 1 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [3,4,3,4,5]</span></p>

<p><strong>输出：</strong><span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code> 的唯一性数组为 <code>[1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3]</code> 。唯一性数组的中位数为 2 ，因此答案是 2 。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [4,3,5,4]</span></p>

<p><strong>输出：</strong><span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code> 的唯一性数组为 <code>[1, 1, 1, 1, 2, 2, 2, 3, 3, 3]</code> 。唯一性数组的中位数为 2 ，因此答案是 2 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

## 提示 1：二分答案

一共有 $m = \dfrac{n(n+1)}{2}$ 个非空连续子数组，中位数是其中第 $k = \left\lceil\dfrac{m}{2}\right\rceil$ 个。

二分中位数 $\textit{upper}$，问题变成：
- $\texttt{distinct}$ 值 $\le \textit{upper}$ 的子数组有多少个？

设子数组的个数为 $\textit{cnt}$，如果 $\textit{cnt} < k$ 说明二分的 $\textit{upper}$ 小了，更新二分左边界 $\textit{left}$，否则更新二分右边界 $\textit{right}$。

## 提示 2：滑动窗口

怎么计算 $\texttt{distinct}$ 值 $\le \textit{upper}$ 的子数组个数？

由于子数组越长，不同元素个数（$\texttt{distinct}$ 值）不会变小，这样的**单调性**可以让我们滑窗。
用一个哈希表 $\textit{freq}$ 统计窗口（子数组）内的元素及其出现次数。
枚举窗口右端点 $r$，把 $\textit{nums}[r]$ 加入 $\textit{freq}$。如果发现 $\textit{freq}$ 的大小超过 $\textit{upper}$，就不断增大窗口左端点 $l$，直到 $\textit{freq}$ 的大小 $\le \textit{upper}$ 为止。
此时右端点为 $r$，左端点为 $l,l+1,l+2,\cdots,r$ 的子数组都是满足要求的（$\texttt{distinct}$ 值 $\le \textit{upper}$），一共有 $r-l+1$ 个，加入统计的子数组个数 $\textit{cnt}$。

## 其它细节

开区间二分左边界：$0$，一定不满足要求，因为没有子数组的 $\texttt{distinct}$ 值是 $0$。
开区间二分右边界：$n$ 或者 $\textit{nums}$ 中的不同元素个数，一定满足要求，因为所有子数组的 $\texttt{distinct}$ 值不超过 $n$。

## 答疑

**问**：为什么二分出来的答案，一定是某个子数组的 $\texttt{distinct}$ 值？有没有可能，二分出来的答案不是任何子数组的 $\texttt{distinct}$ 值？
**答**：反证法。如果答案 $d$ 不是任何子数组的 $\texttt{distinct}$ 值，那么 $\texttt{distinct}$ 值 $\le d$ 和 $\le d-1$ 算出来的子数组个数是一样的。也就是说 $d-1$ 同样满足要求，即 `check(d - 1) == true`，这与循环不变量相矛盾。


```
func medianOfUniquenessArray(nums []int) int {
	n := len(nums)
	k := (n*(n+1)/2 + 1) / 2
	ans := 1 + sort.Search(n-1, func(upper int) bool {
		upper++
		cnt := 0
		l := 0
		freq := map[int]int{}
		for r, in := range nums {
			freq[in]++
			for len(freq) > upper {
				out := nums[l]
				freq[out]--
				if freq[out] == 0 {
					delete(freq, out)
				}
				l++
			}
			cnt += r - l + 1
			if cnt >= k {
				return true
			}
		}
		return false
	})
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。二分 $\mathcal{O}(\log n)$ 次，每次会跑一个 $\mathcal{O}(n)$ 的滑动窗口。
- 空间复杂度：$\mathcal{O}(n)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)