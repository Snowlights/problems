#### 题目  

<p>给你一个 <strong>正</strong> 整数数组 <code>nums</code> 。</p>

<p>将 <code>nums</code> 分成两个数组：<code>nums1</code> 和 <code>nums2</code> ，并满足下述条件：</p>

<ul>
	<li>数组 <code>nums</code> 中的每个元素都属于数组 <code>nums1</code> 或数组 <code>nums2</code> 。</li>
	<li>两个数组都 <strong>非空</strong> 。</li>
	<li>分区值 <strong>最小</strong> 。</li>
</ul>

<p>分区值的计算方法是 <code>|max(nums1) - min(nums2)|</code> 。</p>

<p>其中，<code>max(nums1)</code> 表示数组 <code>nums1</code> 中的最大元素，<code>min(nums2)</code> 表示数组 <code>nums2</code> 中的最小元素。</p>

<p>返回表示分区值的整数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,2,4]
<strong>输出：</strong>1
<strong>解释：</strong>可以将数组 nums 分成 nums1 = [1,2] 和 nums2 = [3,4] 。
- 数组 nums1 的最大值等于 2 。
- 数组 nums2 的最小值等于 3 。
分区值等于 |2 - 3| = 1 。
可以证明 1 是所有分区方案的最小值。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [100,1,10]
<strong>输出：</strong>9
<strong>解释：</strong>可以将数组 nums 分成 nums1 = [10] 和 nums2 = [100,1] 。 
- 数组 nums1 的最大值等于 10 。 
- 数组 nums2 的最小值等于 1 。 
分区值等于 |10 - 1| = 9 。 
可以证明 9 是所有分区方案的最小值。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

答案的理论最小值，是任意两个元素的差的绝对值的最小值。能否取到这个最小值呢？

是可以的，把数组排序后，最小值必然对应两个相邻元素，设其为 $\textit{nums}[i-1]$ 和 $\textit{nums}[i]$。

把不超过 $\textit{nums}[i-1]$ 的数分到第一个数组中，把不低于 $\textit{nums}[i]$ 的数分到第二个数组中，即可满足题目要求。


```go 
func findValueOfPartition(a []int) (ans int) {
	sort.Ints(a)
	ans = math.MaxInt32
	for i := 1; i < len(a); i++ {
		ans = min(ans, a[i]-a[i-1])
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。