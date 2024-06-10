#### 题目

<p>给你一个下标从 <strong>0 </strong>开始的 <strong>正整数</strong> 数组 <code>nums</code> 和一个 <strong>正整数</strong> <code>limit</code> 。</p>

<p>在一次操作中，你可以选择任意两个下标 <code>i</code> 和 <code>j</code>，<strong>如果</strong> 满足 <code>|nums[i] - nums[j]| &lt;= limit</code> ，则交换 <code>nums[i]</code> 和 <code>nums[j]</code> 。</p>

<p>返回执行任意次操作后能得到的 <strong>字典序最小的数组</strong><em> </em>。</p>

<p>如果在数组 <code>a</code> 和数组 <code>b</code> 第一个不同的位置上，数组 <code>a</code> 中的对应字符比数组 <code>b</code> 中的对应字符的字典序更小，则认为数组 <code>a</code> 就比数组 <code>b</code> 字典序更小。例如，数组 <code>[2,10,3]</code> 比数组 <code>[10,2,3]</code> 字典序更小，下标 <code>0</code> 处是两个数组第一个不同的位置，且 <code>2 &lt; 10</code> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,5,3,9,8], limit = 2
<strong>输出：</strong>[1,3,5,8,9]
<strong>解释：</strong>执行 2 次操作：
- 交换 nums[1] 和 nums[2] 。数组变为 [1,3,5,9,8] 。
- 交换 nums[3] 和 nums[4] 。数组变为 [1,3,5,8,9] 。
即便执行更多次操作，也无法得到字典序更小的数组。
注意，执行不同的操作也可能会得到相同的结果。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,7,6,18,2,1], limit = 3
<strong>输出：</strong>[1,6,7,18,1,2]
<strong>解释：</strong>执行 3 次操作：
- 交换 nums[1] 和 nums[2] 。数组变为 [1,6,7,18,2,1] 。
- 交换 nums[0] 和 nums[4] 。数组变为 [2,6,7,18,1,1] 。
- 交换 nums[0] 和 nums[5] 。数组变为 [1,6,7,18,1,2] 。
即便执行更多次操作，也无法得到字典序更小的数组。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,7,28,19,10], limit = 3
<strong>输出：</strong>[1,7,28,19,10]
<strong>解释：</strong>[1,7,28,19,10] 是字典序最小的数组，因为不管怎么选择下标都无法执行操作。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= limit &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

把 $\textit{nums}[i]$ 及其下标 $i$ 绑在一起排序（也可以单独排序下标），然后把 $\textit{nums}$ 分成若干段，每一段都是递增的且相邻元素之差不超过 $\textit{limit}$，那么这一段可以随意排序。

```go  
func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	slices.SortFunc(ids, func(i, j int) int { return nums[i] - nums[j] })

	ans := make([]int, n)
	for i := 0; i < n; {
		st := i
		for i++; i < n && nums[ids[i]]-nums[ids[i-1]] <= limit; i++ {
		}
		subIds := slices.Clone(ids[st:i])
		slices.Sort(subIds)
		for j, idx := range subIds {
			ans[idx] = nums[ids[st+j]]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。