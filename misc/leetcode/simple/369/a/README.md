#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个整数 <code>k</code> 。</p>

<p><code>nums</code> 中的 <strong>K-or</strong> 是一个满足以下条件的非负整数：</p>

<ul>
	<li>只有在 <code>nums</code> 中，至少存在 <code>k</code> 个元素的第 <code>i</code> 位值为 1 ，那么 K-or 中的第 <code>i</code> 位的值才是 1 。</li>
</ul>

<p>返回 <code>nums</code> 的 <strong>K-or</strong> 值。</p>

<p><strong>注意</strong> ：对于整数 <code>x</code> ，如果 <code>(2<sup>i</sup> AND x) == 2<sup>i</sup></code> ，则 <code>x</code> 中的第 <code>i</code> 位值为 1 ，其中 <code>AND</code> 为按位与运算符。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [7,12,9,8,9,15], k = 4
<strong>输出：</strong>9
<strong>解释：</strong>nums[0]、nums[2]、nums[4] 和 nums[5] 的第 0 位的值为 1 。
nums[0] 和 nums[5] 的第 1 位的值为 1 。
nums[0]、nums[1] 和 nums[5] 的第 2 位的值为 1 。
nums[1]、nums[2]、nums[3]、nums[4] 和 nums[5] 的第 3 位的值为 1 。
只有第 0 位和第 3 位满足数组中至少存在 k 个元素在对应位上的值为 1 。因此，答案为 2^0 + 2^3 = 9 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,12,1,11,4,5], k = 6
<strong>输出：</strong>0
<strong>解释：</strong>因为 k == 6 == nums.length ，所以数组的 6-or 等于其中所有元素按位与运算的结果。因此，答案为 2 AND 12 AND 1 AND 11 AND 4 AND 5 = 0 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [10,8,5,9,11,6,8], k = 1
<strong>输出：</strong>15
<strong>解释：</strong>因为 k == 1 ，数组的 1-or 等于其中所有元素按位或运算的结果。因此，答案为 10 OR 8 OR 5 OR 9 OR 11 OR 6 OR 8 = 15 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 50</code></li>
	<li><code>0 <= nums[i] < 2<sup>31</sup></code></li>
	<li><code>1 <= k <= nums.length</code></li>
</ul>

#### 思路

模拟

```go  
func findKOr(nums []int, k int) (ans int) {
	for i := 0; i < 31; i++ {
		cnt1 := 0
		for _, x := range nums {
			cnt1 += x >> i & 1
		}
		if cnt1 >= k {
			ans |= 1 << i
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n \log U)$，其中 $n$ 为 $\textit{nums}$ 的长度。 $U = max(nums)$
- 空间复杂度：$\mathcal{O}(n)$。
