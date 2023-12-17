### 题目

<p>给你两个下标从 <strong>0</strong> 开始的整数数组 <code>nums1</code> 和 <code>nums2</code> ，它们分别含有 <code>n</code> 和 <code>m</code> 个元素。</p>

<p>请你计算以下两个数值：</p>

<ul>
	<li>统计 <code>0 <= i < n</code> 中的下标 <code>i</code> ，满足 <code>nums1[i]</code> 在 <code>nums2</code> 中 <strong>至少</strong> 出现了一次。</li>
	<li>统计 <code>0 <= i < m</code> 中的下标 <code>i</code> ，满足 <code>nums2[i]</code> 在 <code>nums1</code> 中 <strong>至少</strong> 出现了一次。</li>
</ul>

<p>请你返回一个长度为 <code>2</code> 的整数数组<em> </em><code>answer</code> ，<strong>按顺序</strong> 分别为以上两个数值。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [4,3,2,3,1], nums2 = [2,2,5,2,3,6]
<b>输出：</b>[3,4]
<b>解释：</b>分别计算两个数值：
- nums1 中下标为 1 ，2 和 3 的元素在 nums2 中至少出现了一次，所以第一个值为 3 。
- nums2 中下标为 0 ，1 ，3 和 4 的元素在 nums1 中至少出现了一次，所以第二个值为 4 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums1 = [3,4,2,3], nums2 = [1,5]
<b>输出：</b>[0,0]
<b>解释：</b>两个数组中没有公共元素，所以两个值都为 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums1.length</code></li>
	<li><code>m == nums2.length</code></li>
	<li><code>1 <= n, m <= 100</code></li>
	<li><code>1 <= nums1[i], nums2[i] <= 100</code></li>
</ul>

### 思路

转换为哈希表

```go  
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	h, c := make(map[int]bool), make(map[int]bool)
	for _, v := range nums1 {
		h[v] = true
	}
	ans := make([]int, 2)
	for _, v := range nums2 {
		if h[v] {
			ans[1]++
		}
		c[v] = true
	}
	for _, v := range nums1 {
		if c[v] {
			ans[0]++
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums1}$ 的长度、 $m$ 为 $\textit{nums2}$ 的长度。。
- 空间复杂度：$\mathcal{O}(n+m)$。
