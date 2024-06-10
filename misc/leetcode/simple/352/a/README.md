#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个整数 <code>threshold</code> 。</p>

<p>请你从 <code>nums</code> 的子数组中找出以下标 <code>l</code> 开头、下标 <code>r</code> 结尾 <code>(0 <= l <= r < nums.length)</code> 且满足以下条件的 <strong>最长子数组</strong> ：</p>

<ul>
	<li><code>nums[l] % 2 == 0</code></li>
	<li>对于范围 <code>[l, r - 1]</code> 内的所有下标 <code>i</code> ，<code>nums[i] % 2 != nums[i + 1] % 2</code></li>
	<li>对于范围 <code>[l, r]</code> 内的所有下标 <code>i</code> ，<code>nums[i] <= threshold</code></li>
</ul>

<p>以整数形式返回满足题目要求的最长子数组的长度。</p>

<p><strong>注意：子数组</strong> 是数组中的一个连续非空元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [3,2,5,4], threshold = 5
<strong>输出：</strong>3
<strong>解释：</strong>在这个示例中，我们选择从 l = 1 开始、到 r = 3 结束的子数组 => [2,5,4] ，满足上述条件。
因此，答案就是这个子数组的长度 3 。可以证明 3 是满足题目要求的最大长度。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1,2], threshold = 2
<strong>输出：</strong>1
<strong>解释：</strong>
在这个示例中，我们选择从 l = 1 开始、到 r = 1 结束的子数组 => [2] 。
该子数组满足上述全部条件。可以证明 1 是满足题目要求的最大长度。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>nums = [2,3,4,5], threshold = 4
<strong>输出：</strong>3
<strong>解释：</strong>
在这个示例中，我们选择从 l = 0 开始、到 r = 2 结束的子数组 => [2,3,4] 。 
该子数组满足上述全部条件。
因此，答案就是这个子数组的长度 3 。可以证明 3 是满足题目要求的最大长度。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 100 </code></li>
	<li><code>1 <= nums[i] <= 100 </code></li>
	<li><code>1 <= threshold <= 100</code></li>
</ul>

#### 思路

题目的约束实际上把数组划分成了若干段，每段都满足要求，且互不相交。

那么遍历一遍，计算每一段的长度，取最大值，即为答案。

```go
func longestAlternatingSubarray(a []int, threshold int) (ans int) {
	for i, n := 0, len(a); i < n; {
		if a[i]%2 > 0 || a[i] > threshold {
			i++
		} else {
			i0 := i
			for i++; i < n && a[i] <= threshold && a[i]%2 != a[i-1]%2; i++ {}
			ans = max(ans, i-i0)
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。注意 $i$ 是全局变量，只会增加，不会减少。所以二重循环至多执行 $\mathcal{O}(n)$ 次。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
