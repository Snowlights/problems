### 题目  

<p>给你一个只包含 <strong>非负</strong>&nbsp;整数的数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>我们定义满足 <code>l &lt;= r</code>&nbsp;的子数组&nbsp;<code>nums[l..r]</code>&nbsp;的分数为&nbsp;<code>nums[l] AND nums[l + 1] AND ... AND nums[r]</code>&nbsp;，其中&nbsp;<strong>AND</strong>&nbsp;是按位与运算。</p>

<p>请你将数组分割成一个或者更多子数组，满足：</p>

<ul>
	<li><strong>每个</strong> 元素都&nbsp;<strong>只</strong>&nbsp;属于一个子数组。</li>
	<li>子数组分数之和尽可能<strong>&nbsp;小</strong>&nbsp;。</li>
</ul>

<p>请你在满足以上要求的条件下，返回<strong>&nbsp;最多</strong>&nbsp;可以得到多少个子数组。</p>

<p>一个 <strong>子数组</strong>&nbsp;是一个数组中一段连续的元素。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,0,2,0,1,2]
<b>输出：</b>3
<strong>解释：</strong>我们可以将数组分割成以下子数组：
- [1,0] 。子数组分数为 1 AND 0 = 0 。
- [2,0] 。子数组分数为 2 AND 0 = 0 。
- [1,2] 。子数组分数为 1 AND 2 = 0 。
分数之和为 0 + 0 + 0 = 0 ，是我们可以得到的最小分数之和。
在分数之和为 0 的前提下，最多可以将数组分割成 3 个子数组。所以返回 3 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [5,7,1,3]
<b>输出：</b>1
<b>解释：</b>我们可以将数组分割成一个子数组：[5,7,1,3] ，分数为 1 ，这是可以得到的最小总分数。
在总分数为 1 的前提下，最多可以将数组分割成 1 个子数组。所以返回 1 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>

### 思路

我们需要：
- 先满足子数组分数之和尽量小；
- 再满足分割出的子数组尽量多。

## 提示 1

AND 的性质是，参与 AND 的数越多，结果越小。

## 提示 2

子数组的 AND，不会低于整个 $\\textit{nums}$ 数组的 AND。

## 提示 3

如果 $\\textit{nums}$ 数组的 AND（记作 $a$）是大于 $0$ 的，那么只能分割出一个数组（即 $\\textit{nums}$ 数组）。根据提示 2，如果分割出超过一个数组，那么得分至少为 $2a$，而这是大于 $a$ 的，不满足「子数组分数之和尽量小」的要求。所以在 $a>0$ 的情况下，答案为 $1$。  
如果 $\\textit{nums}$ 数组的 AND 是等于 $0$ 的，那么可以分割出尽量多的 AND 等于 $0$ 的子数组。怎么分？从左到右遍历数组，只要发现 AND 等于 $0$ 就立刻分割。如果不立刻分割，由于 AND 的数越多越能为 $0$，现在多分了一个数，后面就要少分一个数，可能后面就不能为 $0$ 了。注意，如果最后剩下的一段子数组的 AND 大于 $0$，这一段可以直接合并到前一个 AND 为 $0$ 的子数组中。

```go  
func maxSubarrays(nums []int) (ans int) {
	and := -1 // -1 就是 111...1，和任何数 AND 都等于那个数
	for _, x := range nums {
		and &= x
		if and == 0 {
			ans++ // 分割
			and = -1
		}
	}
	return max(ans, 1) // 如果 ans=0 说明所有数的 and>0，答案为 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$ 。