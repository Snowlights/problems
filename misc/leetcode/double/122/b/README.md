### 题目

<p>给你一个下标从 <strong>0</strong> 开始且全是 <strong>正</strong> 整数的数组 <code>nums</code> 。</p>

<p>一次 <b>操作</b> 中，如果两个 <strong>相邻</strong> 元素在二进制下数位为 <strong>1</strong> 的数目 <strong>相同</strong> ，那么你可以将这两个元素交换。你可以执行这个操作 <strong>任意次</strong> （<strong>也可以 0 次</strong>）。</p>

<p>如果你可以使数组变有序，请你返回 <code>true</code> ，否则返回 <code>false</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [8,4,2,30,15]
<b>输出：</b>true
<b>解释：</b>我们先观察每个元素的二进制表示。 2 ，4 和 8 分别都只有一个数位为 1 ，分别为 "10" ，"100" 和 "1000" 。15 和 30 分别有 4 个数位为 1 ："1111" 和 "11110" 。
我们可以通过 4 个操作使数组有序：
- 交换 nums[0] 和 nums[1] 。8 和 4 分别只有 1 个数位为 1 。数组变为 [4,8,2,30,15] 。
- 交换 nums[1] 和 nums[2] 。8 和 2 分别只有 1 个数位为 1 。数组变为 [4,2,8,30,15] 。
- 交换 nums[0] 和 nums[1] 。4 和 2 分别只有 1 个数位为 1 。数组变为 [2,4,8,30,15] 。
- 交换 nums[3] 和 nums[4] 。30 和 15 分别有 4 个数位为 1 ，数组变为 [2,4,8,15,30] 。
数组变成有序的，所以我们返回 true 。
注意我们还可以通过其他的操作序列使数组变得有序。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,3,4,5]
<b>输出：</b>true
<b>解释：</b>数组已经是有序的，所以我们返回 true 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [3,16,8,4,2]
<b>输出：</b>false
<b>解释：</b>无法通过操作使数组变为有序。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 100</code></li>
	<li><code>1 <= nums[i] <= 2<sup>8</sup></code></li>
</ul>

### 思路

**适用场景**：按照题目要求，数组会被分割成若干组，每一组的判断/处理逻辑是相同的。

**核心思想**：
- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的工作（排序）。
- 内层循环负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。这个写法是所有写法中最不容易出 bug 的。

```go [sol]
func canSortArray(nums []int) bool {
	i, n := 0, len(nums)

	for i < n {
		s := i
		for i + 1 < n && bits.OnesCount(uint(nums[i])) ==
			bits.OnesCount(uint(nums[i+1])) {
			i++
		}
		i++
		sort.Ints(nums[s:i])
	}

	return sort.IntsAreSorted(nums)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n \log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
