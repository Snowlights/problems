### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。如果 <code>nums</code> 中长度为 <code>m</code> 的子数组 <code>s</code> 满足以下条件，我们称它是一个交替子序列：</p>

<ul>
	<li><code>m</code> 大于 <code>1</code> 。</li>
	<li><code>s<sub>1</sub> = s<sub>0</sub> + 1</code> 。</li>
	<li>下标从 0 开始的子数组 <code>s</code> 与数组 <code>[s<sub>0</sub>, s<sub>1</sub>, s<sub>0</sub>, s<sub>1</sub>,...,s<sub>(m-1) % 2</sub>]</code> 一样。也就是说，<code>s<sub>1</sub> - s<sub>0</sub> = 1</code> ，<code>s<sub>2</sub> - s<sub>1</sub> = -1</code> ，<code>s<sub>3</sub> - s<sub>2</sub> = 1</code> ，<code>s<sub>4</sub> - s<sub>3</sub> = -1</code> ，以此类推，直到 <code>s[m - 1] - s[m - 2] = (-1)<sup>m</sup></code> 。</li>
</ul>

<p>请你返回 <code>nums</code> 中所有 <strong>交替</strong> 子数组中，最长的长度，如果不存在交替子数组，请你返回 <code>-1</code> 。</p>

<p>子数组是一个数组中一段连续 <strong>非空</strong> 的元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [2,3,4,3,4]
<b>输出：</b>4
<b>解释：</b>交替子数组有 [3,4] ，[3,4,3] 和 [3,4,3,4] 。最长的子数组为 [3,4,3,4] ，长度为4 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [4,5,6]
<b>输出：</b>2
<strong>解释：</strong>[4,5] 和 [5,6] 是仅有的两个交替子数组。它们长度都为 2 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
</ul>
 
### 思路 

直接暴力向后寻找

```go 
func alternatingSubarray(a []int) (ans int) {
	ans = -1
	for i := range a {
		j, base := i + 1, 1
		for j < len(a) {
			if a[j] - a[j-1] != base {
				break
			}
			base *= -1
			j++
		}
		if j - i > 1 {
			ans = max(ans, j - i)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(1)$ 。仅用到部分变量。