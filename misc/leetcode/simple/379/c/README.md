#### 题目

<p>给你两个下标从 <code>0</code> 开始的整数数组 <code>nums1</code> 和 <code>nums2</code> ，它们的长度都是偶数<code> n</code> 。</p>

<p>你必须从 <code>nums1</code> 中移除 <code>n / 2</code> 个元素，同时从 <code>nums2</code> 中也移除 <code>n / 2</code> 个元素。移除之后，你将 <code>nums1</code> 和 <code>nums2</code> 中剩下的元素插入到集合 <code>s</code> 中。</p>

<p>返回集合 <code>s</code>可能的<strong> 最多 </strong>包含多少元素。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1,2,1,2], nums2 = [1,1,1,1]
<strong>输出：</strong>2
<strong>解释：</strong>从 nums1 和 nums2 中移除两个 1 。移除后，数组变为 nums1 = [2,2] 和 nums2 = [1,1] 。因此，s = {1,2} 。
可以证明，在移除之后，集合 s 最多可以包含 2 个元素。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1,2,3,4,5,6], nums2 = [2,3,2,3,2,3]
<strong>输出：</strong>5
<strong>解释：</strong>从 nums1 中移除 2、3 和 6 ，同时从 nums2 中移除两个 3 和一个 2 。移除后，数组变为 nums1 = [1,4,5] 和 nums2 = [2,3,2] 。因此，s = {1,2,3,4,5} 。
可以证明，在移除之后，集合 s 最多可以包含 5 个元素。 
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [1,1,2,2,3,3], nums2 = [4,4,5,5,6,6]
<strong>输出：</strong>6
<strong>解释：</strong>从 nums1 中移除 1、2 和 3 ，同时从 nums2 中移除 4、5 和 6 。移除后，数组变为 nums1 = [1,2,3] 和 nums2 = [4,5,6] 。因此，s = {1,2,3,4,5,6} 。
可以证明，在移除之后，集合 s 最多可以包含 6 个元素。 </pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums1.length == nums2.length</code></li>
	<li><code>1 <= n <= 2 * 10<sup>4</sup></code></li>
	<li><code>n</code>是偶数。</li>
	<li><code>1 <= nums1[i], nums2[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

从「添加」的角度思考

设 $\textit{nums}_1$ 中有 $n_1$ 个不同元素，
$\textit{nums}_2$ 中有 $n_2$ 个不同元素，它们的交集有 $\textit{common}$ 个元素。考虑怎么从两个数组中选择不同元素，添加到集合 $s$ 中，使 $s$ 的大小最大：

- 对于 $\textit{nums}_1$，优先选**不在**交集中的元素，这可以选 $n_1-\textit{common}$ 个，但不能超过题目规定的 $n/2$，所以至多选 $c_1 = \min(n_1-\textit{common}, n/2)$ 个不在交集中的元素。
- 对于 $\textit{nums}_2$，优先选**不在**交集中的元素，同理，至多选 $c_2 = \min(n_2-\textit{common}, n/2)$ 个。
- 由于都和 $n/2$ 取最小值，所以 $c_1 + c_2 \le n/2 + n/2 = n$。
- 如果 $c_1 + c_2 < n$，那么还可以再选 $n-c_1-c_2$ 个数，且这些数只能从交集中选，所以不能超过 $\textit{common}$ 个，所以还可以再选 $\min(n-c_1-c_2, \textit{common})$ 个数。

最终答案为

$$
\begin{aligned}
&c_1 + c_2 + \min(n-c_1-c_2, \textit{common})\\
=\ &\min(n, c_1 + c_2 + \textit{common})
\end{aligned}

$$

```go [sol]
func maximumSetSize(nums1 []int, nums2 []int) int {
	s1, s2 := make(map[int]bool), make(map[int]bool)
	n, common := len(nums1), 0
	for _, v := range nums1 {
		s1[v] = true
	}
	for _, v := range nums2 {
		if s2[v] {
			continue
		}
		s2[v] = true
		if s1[v] {
			common++
		}
	}

	c1 := min(n/2, len(s1)-common)
	c2 := min(n/2, len(s2)-common)
	return min(n, c1+c2+common)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums1}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
