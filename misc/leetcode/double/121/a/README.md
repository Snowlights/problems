### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。</p>

<p>如果一个前缀 <code>nums[0..i]</code> 满足对于 <code>1 <= j <= i</code> 的所有元素都有 <code>nums[j] = nums[j - 1] + 1</code> ，那么我们称这个前缀是一个 <strong>顺序前缀</strong> 。特殊情况是，只包含 <code>nums[0]</code> 的前缀也是一个 <strong>顺序前缀</strong> 。</p>

<p>请你返回 <code>nums</code> 中没有出现过的 <strong>最小</strong> 整数 <code>x</code> ，满足 <code>x</code> 大于等于 <strong>最长</strong> 顺序前缀的和。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,3,2,5]
<b>输出：</b>6
<b>解释：</b>nums 的最长顺序前缀是 [1,2,3] ，和为 6 ，6 不在数组中，所以 6 是大于等于最长顺序前缀和的最小整数。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [3,4,5,1,12,14,13]
<b>输出：</b>15
<b>解释：</b>nums 的最长顺序前缀是 [3,4,5] ，和为 12 ，12、13 和 14 都在数组中，但 15 不在，所以 15 是大于等于最长顺序前缀和的最小整数。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 50</code></li>
	<li><code>1 <= nums[i] <= 50</code></li>
</ul>

### 思路

1. 求出最长前缀，同时维护最长前缀的元素和 $\textit{sum}$。
2. 不断增加 $\textit{sum}$，直到 $\textit{sum}$ 不在 $\textit{nums}$ 中。
3. 返回 $\textit{sum}$。

```go [sol]
func missingInteger(nums []int) int {
	x, h := nums[0], make(map[int]bool)
	flag := true
	for i, v := range nums {
		if flag && i > 0 && v - nums[i-1] == 1 {
			x += v
		} else if i > 0 {
			flag = false
		}
		h[v] = true
	}
	for h[x] {
		x++
	}
	return x
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
