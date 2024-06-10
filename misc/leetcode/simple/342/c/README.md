#### 题目

<p>给你一个长度为 <code>n</code> 的整数数组 <code>nums</code> ，请你求出每个长度为 <code>k</code> 的子数组的 <b>美丽值</b> 。</p>

<p>一个子数组的 <strong>美丽值</strong> 定义为：如果子数组中第 <code>x</code> <strong>小整数</strong> 是 <strong>负数</strong> ，那么美丽值为第 <code>x</code> 小的数，否则美丽值为 <code>0</code> 。</p>

<p>请你返回一个包含<em> </em><code>n - k + 1</code> 个整数的数组，<strong>依次</strong> 表示数组中从第一个下标开始，每个长度为 <code>k</code> 的子数组的<strong> 美丽值</strong> 。</p>

<ul>
	<li>
	<p>子数组指的是数组中一段连续 <strong>非空</strong> 的元素序列。</p>
	</li>
</ul>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [1,-1,-3,-2,3], k = 3, x = 2
<b>输出：</b>[-1,-2,-2]
<b>解释：</b>总共有 3 个 k = 3 的子数组。
第一个子数组是 <code>[1, -1, -3]</code> ，第二小的数是负数 -1 。
第二个子数组是 <code>[-1, -3, -2]</code> ，第二小的数是负数 -2 。
第三个子数组是 <code>[-3, -2, 3] ，第二小的数是负数 -2 。</code></pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [-1,-2,-3,-4,-5], k = 2, x = 2
<b>输出：</b>[-1,-2,-3,-4]
<b>解释：</b>总共有 4 个 k = 2 的子数组。
<code>[-1, -2] 中第二小的数是负数 -1 。</code>
<code>[-2, -3] 中第二小的数是负数 -2 。</code>
<code>[-3, -4] 中第二小的数是负数 -3 。</code>
<code>[-4, -5] 中第二小的数是负数 -4 。</code></pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>nums = [-3,1,2,-3,0,-3], k = 2, x = 1
<b>输出：</b>[-3,0,-3,-3,-3]
<b>解释：</b>总共有 5 个 k = 2 的子数组。
<code>[-3, 1] 中最小的数是负数 -3 。</code>
<code>[1, 2] 中最小的数不是负数，所以美丽值为 0 。</code>
<code>[2, -3] 中最小的数是负数 -3 。</code>
<code>[-3, 0] 中最小的数是负数 -3 。</code>
<code>[0, -3] 中最小的数是负数 -3 。</code></pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums.length </code></li>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= k <= n</code></li>
	<li><code>1 <= x <= k </code></li>
	<li><code>-50 <= nums[i] <= 50 </code></li>
</ul>

#### 思路

滑动窗口。由于值域很小，所以借鉴**计数排序**，用一个 $\textit{cnt}$ 数组维护窗口内每个数的出现次数。然后遍历 $\textit{cnt}$ 去求第 $x$ 小的数。
什么是第 $x$ 小的数？
设它是 $\textit{num}$，那么 $<\textit{num}$ 的数有 $<x$ 个，$\le\textit{num}$ 的数有 $\ge x$ 个，就说明 $\textit{num}$ 是第 $x$ 小的数。

```go
func getSubarrayBeauty(nums []int, k, x int) []int {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] { // 先往窗口内添加 k-1 个数
		cnt[num+bias]++
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+bias]++ // 进入窗口（保证窗口有恰好 k 个数）
		left := x
		for j, c := range cnt[:bias] { // 暴力枚举负数范围 [-50,-1]
			left -= c
			if left <= 0 { // 找到美丽值
				ans[i] = j - bias
				break
			}
		}
		cnt[nums[i]+bias]-- // 离开窗口
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=50$。
- 空间复杂度：$\mathcal{O}(U)$。
