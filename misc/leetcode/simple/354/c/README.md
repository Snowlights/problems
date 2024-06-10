#### 题目

<p>如果元素 <code>x</code> 在长度为 <code>m</code> 的整数数组 <code>arr</code> 中满足 <code>freq(x) * 2 > m</code> ，那么我们称 <code>x</code> 是 <strong>支配元素</strong> 。其中 <code>freq(x)</code> 是 <code>x</code> 在数组 <code>arr</code> 中出现的次数。注意，根据这个定义，数组 <code>arr</code> <strong>最多</strong> 只会有 <strong>一个</strong> 支配元素。</p>

<p>给你一个下标从 <strong>0</strong> 开始长度为 <code>n</code> 的整数数组 <code>nums</code> ，数据保证它含有一个支配元素。</p>

<p>你需要在下标 <code>i</code> 处将 <code>nums</code> 分割成两个数组 <code>nums[0, ..., i]</code> 和 <code>nums[i + 1, ..., n - 1]</code> ，如果一个分割满足以下条件，我们称它是 <strong>合法</strong> 的：</p>

<ul>
	<li><code>0 <= i < n - 1</code></li>
	<li><code>nums[0, ..., i]</code> 和 <code>nums[i + 1, ..., n - 1]</code> 的支配元素相同。</li>
</ul>

<p>这里， <code>nums[i, ..., j]</code> 表示 <code>nums</code> 的一个子数组，它开始于下标 <code>i</code> ，结束于下标 <code>j</code> ，两个端点都包含在子数组内。特别地，如果 <code>j < i</code> ，那么 <code>nums[i, ..., j]</code> 表示一个空数组。</p>

<p>请你返回一个 <strong>合法分割</strong> 的 <strong>最小</strong> 下标。如果合法分割不存在，返回 <code>-1</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [1,2,2,2]
<b>输出：</b>2
<b>解释：</b>我们将数组在下标 2 处分割，得到 [1,2,2] 和 [2] 。
数组 [1,2,2] 中，元素 2 是支配元素，因为它在数组中出现了 2 次，且 2 * 2 > 3 。
数组 [2] 中，元素 2 是支配元素，因为它在数组中出现了 1 次，且 1 * 2 > 1 。
两个数组 [1,2,2] 和 [2] 都有与 nums 一样的支配元素，所以这是一个合法分割。
下标 2 是合法分割中的最小下标。</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [2,1,3,1,1,1,7,1,2,1]
<b>输出：</b>4
<b>解释：</b>我们将数组在下标 4 处分割，得到 [2,1,3,1,1] 和 [1,7,1,2,1] 。
数组 [2,1,3,1,1] 中，元素 1 是支配元素，因为它在数组中出现了 3 次，且 3 * 2 > 5 。
数组 [1,7,1,2,1] 中，元素 1 是支配元素，因为它在数组中出现了 3 次，且 3 * 2 > 5 。
两个数组 [2,1,3,1,1] 和 [1,7,1,2,1] 都有与 nums 一样的支配元素，所以这是一个合法分割。
下标 4 是所有合法分割中的最小下标。</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>nums = [3,3,3,3,7,2,2]
<b>输出：</b>-1
<b>解释：</b>没有合法分割。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>nums</code> 有且只有一个支配元素。</li>
</ul>

#### 思路

首先证明：分割出的两个数组的支配元素就是原数组的支配元素。
设这两个数组的支配元素为 $y$（题目要求支配元素相同），那么对于第一个数组有

$$
\text{freq}_1(y) \cdot 2 > i+1

$$

对于第二个数组有

$$
\text{freq}_2(y) \cdot 2 > n-i-1

$$

由于这两个数组合并之后就是原数组，所以

$$
\text{freq}(y) \cdot 2 = \text{freq}_1(y) \cdot 2 + \text{freq}_2(y) \cdot 2 > (i+1) + (n-i-1) = n

$$

上式表明，$y$ 就是原数组的支配元素，证毕。

#### 算法

首先求出众数 $\textit{mode}$ 及其出现次数 $\textit{total}$。
然后枚举 $i$，一边枚举一边统计 $\text{freq}_1(\textit{mode})$，那么 $\text{freq}_2(\textit{mode}) =\textit{total} -\text{freq}_1(\textit{mode})$。
只要满足 $\text{freq}_1(\textit{mode}) \cdot 2 > i+1$ 且 $\text{freq}_2(\textit{mode}) \cdot 2 > n-i-1$，就返回 $i$。
如果没有这样的 $i$，返回 $-1$。

```go
func minimumIndex(nums []int) int {
	// 也可以用摩尔投票法实现
	freq := map[int]int{}
	mode := nums[0]
	for _, x := range nums {
		freq[x]++
		if freq[x] > freq[mode] {
			mode = x
		}
	}

	total := freq[mode]
	freq1 := 0
	for i, x := range nums {
		if x == mode {
			freq1++
		}
		if freq1*2 > i+1 && (total-freq1)*2 > len(nums)-i-1 {
			return i
		}
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。用摩尔投票法可以做到 $\mathcal{O}(1)$ 额外空间。
