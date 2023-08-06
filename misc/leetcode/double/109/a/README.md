### 题目

<p>给你一个整数数组 <code>nums</code> ，如果它是数组 <code>base[n]</code> 的一个排列，我们称它是个 <strong>好</strong> 数组。</p>

<p><code>base[n] = [1, 2, ..., n - 1, n, n]</code> （换句话说，它是一个长度为 <code>n + 1</code> 且包含 <code>1</code> 到 <code>n - 1</code> 恰好各一次，包含 <code>n</code>  两次的一个数组）。比方说，<code>base[1] = [1, 1]</code> ，<code>base[3] = [1, 2, 3, 3]</code> 。</p>

<p>如果数组是一个好数组，请你返回 <code>true</code> ，否则返回 <code>false</code> 。</p>

<p><strong>注意：</strong>数组的排列是这些数字按任意顺序排布后重新得到的数组。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [2, 1, 3]
<b>输出：</b>false
<b>解释：</b>因为数组的最大元素是 3 ，唯一可以构成这个数组的 base[n] 对应的 n = 3 。但是 base[3] 有 4 个元素，但数组 nums 只有 3 个元素，所以无法得到 base[3] = [1, 2, 3, 3] 的排列，所以答案为 false 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [1, 3, 3, 2]
<b>输出：</b>true
<b>解释：因为</b>数组的最大元素是 3 ，唯一可以构成这个数组的 base[n] 对应的 n = 3 ，可以看出数组是 base[3] = [1, 2, 3, 3] 的一个排列（交换 nums 中第二个和第四个元素）。所以答案为 true 。</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>nums = [1, 1]
<b>输出：</b>true
<b>解释：</b>因为数组的最大元素是 1 ，唯一可以构成这个数组的 base[n] 对应的 n = 1，可以看出数组是 base[1] = [1, 1] 的一个排列。所以答案为 true 。</pre>

<p><strong>示例 4：</strong></p>

<pre><b>输入：</b>nums = [3, 4, 4, 1, 2, 1]
<b>输出：</b>false
<b>解释：</b>因为数组的最大元素是 4 ，唯一可以构成这个数组的 base[n] 对应的 n = 4 。但是 base[n] 有 5 个元素而 nums 有 6 个元素。所以答案为 false 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 100</code></li>
	<li><code>1 <= num[i] <= 200</code></li>
</ul>

### 思路

首先，一旦出现超过 $n$ 的数，直接返回 `false`。
其次，只要小于 $n$ 的数，没有出现超过一次，那么就说明这些数都出现恰好一次，且 $n$ 恰好出现两次。

```go

func isGood(nums []int) bool {
	n := len(nums) - 1
	vis := make([]bool, n+1)
	for _, v := range nums {
		if v > n || v < n && vis[v] {
			return false
		}
		vis[v] = true
	}
	return true
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 。
