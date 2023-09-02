### 题目

<p>给你一个下标从 <strong>0</strong> 开始长度为 <code>n</code> 的数组 <code>nums</code> 。</p>

<p>每一秒，你可以对数组执行以下操作：</p>

<ul>
	<li>对于范围在 <code>[0, n - 1]</code> 内的每一个下标 <code>i</code> ，将 <code>nums[i]</code> 替换成 <code>nums[i]</code> ，<code>nums[(i - 1 + n) % n]</code> 或者 <code>nums[(i + 1) % n]</code> 三者之一。</li>
</ul>

<p><strong>注意</strong>，所有元素会被同时替换。</p>

<p>请你返回将数组 <code>nums</code> 中所有元素变成相等元素所需要的 <strong>最少</strong> 秒数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [1,2,1,2]
<b>输出：</b>1
<b>解释：</b>我们可以在 1 秒内将数组变成相等元素：
- 第 1 秒，将每个位置的元素分别变为 [nums[3],nums[1],nums[3],nums[3]] 。变化后，nums = [2,2,2,2] 。
1 秒是将数组变成相等元素所需要的最少秒数。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [2,1,3,3,2]
<b>输出：</b>2
<b>解释：</b>我们可以在 2 秒内将数组变成相等元素：
- 第 1 秒，将每个位置的元素分别变为 [nums[0],nums[2],nums[2],nums[2],nums[3]] 。变化后，nums = [2,3,3,3,3] 。
- 第 2 秒，将每个位置的元素分别变为 [nums[1],nums[1],nums[2],nums[3],nums[4]] 。变化后，nums = [3,3,3,3,3] 。
2 秒是将数组变成相等元素所需要的最少秒数。
</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>nums = [5,5,5,5]
<b>输出：</b>0
<b>解释：</b>不需要执行任何操作，因为一开始数组中的元素已经全部相等。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n == nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

### 思路

### 提示 1

最终所有元素一定变成了一个在 $\textit{nums}$ 中的数。
枚举这个数。

### 提示 2

考虑把数字 $x$「扩散」到其它位置，那么每一秒 $x$ 都可以向左右扩散一位。
多个相同数字 $x$ 同时扩散，那么扩散完整个数组的耗时，就取决于相距最远的两个相邻的 $x$。
假设这两个 $x$ 的下标分别为 $i$ 和 $j$，且 $i<j$，那么耗时为：

$$
\left\lfloor\dfrac{j-i}{2}\right\rfloor

$$

取所有耗时的最小值，即为答案。

### 提示 3

统计所有相同数字的下标，记到一个哈希表 $\textit{pos}$ 中。
本题数组可以视作是环形的，假设最左边的 $x$ 的下标是 $i$，只需要在 $\textit{pos}[x]$ 列表末尾添加一个 $i+n$，就可以转换成非环形数组处理了。

```go
func minimumSeconds(a []int) (ans int) {
	m := make(map[int][]int)
	for i, v := range a {
		m[v] = append(m[v], i)
	}

	ans = len(a)
	for _, v := range m {
		tmp := 0
		for i := 1; i < len(v); i++ {
			tmp = max(tmp, v[i]-v[i-1])
		}
		tmp = max(tmp, len(a)-v[len(v)-1]+v[0])
		ans = min(ans, tmp/2)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
