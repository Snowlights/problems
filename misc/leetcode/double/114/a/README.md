### 题目

<p>给你一个正整数数组&nbsp;<code>nums</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;。</p>

<p>一次操作中，你可以将数组的最后一个元素删除，将该元素添加到一个集合中。</p>

<p>请你返回收集元素&nbsp;<code>1, 2, ..., k</code>&nbsp;需要的&nbsp;<strong>最少操作次数</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,1,5,4,2], k = 2
<b>输出：</b>4
<b>解释：</b>4 次操作后，集合中的元素依次添加了 2 ，4 ，5 和 1 。此时集合中包含元素 1 和 2 ，所以答案为 4 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [3,1,5,4,2], k = 5
<b>输出：</b>5
<b>解释：</b>5 次操作后，集合中的元素依次添加了 2 ，4 ，5 ，1 和 3 。此时集合中包含元素 1 到 5 ，所以答案为 5 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [3,2,5,3,1], k = 3
<b>输出：</b>4
<b>解释：</b>4 次操作后，集合中的元素依次添加了 1 ，3 ，5 和 2 。此时集合中包含元素 1 到 3  ，所以答案为 4 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 50</code></li>
	<li><code>1 &lt;= nums[i] &lt;= nums.length</code></li>
	<li><code>1 &lt;= k &lt;= nums.length</code></li>
	<li>输入保证你可以收集到元素&nbsp;<code>1, 2, ..., k</code> 。</li>
</ul>

### 思路

倒序遍历。由于元素范围在 $[1,50]$，我们可以用一个 $64$ 位整数表示集合，只要集合中有 $1$ 到 $k$ 就立刻返回。

```go  
func minOperations(nums []int, k int) int {
	all := 2<<k - 2 // 1~k
	set := 0
	for i := len(nums) - 1; ; i-- {
		set |= 1 << nums[i]
		if set&all == all {
			return len(nums) - i
		}
	}
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$ 。
