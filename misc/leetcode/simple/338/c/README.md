#### 题目  

<p>给你一个正整数数组 <code>nums</code> 。</p>

<p>同时给你一个长度为 <code>m</code> 的整数数组 <code>queries</code> 。第 <code>i</code> 个查询中，你需要将 <code>nums</code> 中所有元素变成 <code>queries[i]</code> 。你可以执行以下操作 <strong>任意</strong> 次：</p>

<ul>
	<li>将数组里一个元素 <strong>增大</strong> 或者 <strong>减小</strong> <code>1</code> 。</li>
</ul>

<p>请你返回一个长度为 <code>m</code> 的数组<em> </em><code>answer</code> ，其中<em> </em><code>answer[i]</code>是将 <code>nums</code> 中所有元素变成 <code>queries[i]</code> 的 <strong>最少</strong> 操作次数。</p>

<p><strong>注意</strong>，每次查询后，数组变回最开始的值。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [3,1,6,8], queries = [1,5]
<b>输出：</b>[14,10]
<b>解释：</b>第一个查询，我们可以执行以下操作：
- 将 nums[0] 减小 2 次，nums = [1,1,6,8] 。
- 将 nums[2] 减小 5 次，nums = [1,1,1,8] 。
- 将 nums[3] 减小 7 次，nums = [1,1,1,1] 。
第一个查询的总操作次数为 2 + 5 + 7 = 14 。
第二个查询，我们可以执行以下操作：
- 将 nums[0] 增大 2 次，nums = [5,1,6,8] 。
- 将 nums[1] 增大 4 次，nums = [5,5,6,8] 。
- 将 nums[2] 减小 1 次，nums = [5,5,5,8] 。
- 将 nums[3] 减小 3 次，nums = [5,5,5,5] 。
第二个查询的总操作次数为 2 + 4 + 1 + 3 = 10 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [2,9,6,3], queries = [10]
<b>输出：</b>[20]
<b>解释：</b>我们可以将数组中所有元素都增大到 10 ，总操作次数为 8 + 1 + 4 + 7 = 20 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums.length</code></li>
	<li><code>m == queries.length</code></li>
	<li><code>1 &lt;= n, m &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i], queries[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

![t3.png](https://pic.leetcode.cn/1679808210-FVsAou-t3.png)

```go 
func minOperations(nums, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1) // 前缀和
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	ans := make([]int64, len(queries))
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		left := q*j - sum[j]               // 蓝色面积
		right := sum[n] - sum[j] - q*(n-j) // 绿色面积
		ans[i] = int64(left + right)
	}
	return ans
}
```

#### 复杂度分析  

- 时间复杂度：$O((n+q)\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$O(n)$。返回值不计入。