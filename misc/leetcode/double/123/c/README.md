### 题目

<p>给你一个长度为 <code>n</code> 的数组 <code>nums</code> 和一个 <strong>正</strong> 整数 <code>k</code> 。</p>

<p>如果 <code>nums</code> 的一个子数组中，第一个元素和最后一个元素 <strong>差的绝对值恰好</strong> 为 <code>k</code> ，我们称这个子数组为 <strong>好</strong> 的。换句话说，如果子数组 <code>nums[i..j]</code> 满足 <code>|nums[i] - nums[j]| == k</code> ，那么它是一个好子数组。</p>

<p>请你返回 <code>nums</code> 中 <strong>好</strong> 子数组的 <strong>最大</strong> 和，如果没有好子数组，返回<em> </em><code>0</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,3,4,5,6], k = 1
<b>输出：</b>11
<b>解释：</b>好子数组中第一个元素和最后一个元素的差的绝对值必须为 1 。好子数组有 [1,2] ，[2,3] ，[3,4] ，[4,5] 和 [5,6] 。最大子数组和为 11 ，对应的子数组为 [5,6] 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [-1,3,2,4,5], k = 3
<b>输出：</b>11
<b>解释：</b>好子数组中第一个元素和最后一个元素的差的绝对值必须为 3 。好子数组有 [-1,3,2] 和 [2,4,5] 。最大子数组和为 11 ，对应的子数组为 [2,4,5] 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [-1,-2,-3,-4], k = 2
<b>输出：</b>-6
<b>解释：</b>好子数组中第一个元素和最后一个元素的差的绝对值必须为 2 。好子数组有 [-1,-2,-3] 和 [-2,-3,-4] 。最大子数组和为 -6 ，对应的子数组为 [-1,-2,-3] 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>-10<sup>9</sup> <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>1 <= k <= 10<sup>9</sup></code></li>
</ul>

### 思路

为方便描述，把 $\textit{nums}$ 简称为 $a$。
子数组 $a[i..j]$ 的元素和为

$$
s[j+1]-s[i]

$$

枚举 $j$，我们需要找到最小的 $s[i]$，满足 $|a[i]-a[j]|=k$，即 $a[i] = a[j]-k$ 或 $a[i]=a[j]+k$。
定义哈希表 $\textit{minS}$，键为 $a[i]$，值为相同 $a[i]$ 下的 $s[i]$ 的最小值。
子数组最后一个数为 $a[j]$ 时，子数组的最大元素和为

$$
\begin{aligned}
& s[j+1]-\textit{minS}[a[i]]\\
=\ &s[j+1]-\min(\textit{minS}[a[j]-k],\textit{minS}[a[j]+k])
\end{aligned}

$$

```go [sol]
func maximumSubarraySum(nums []int, k int) int64 {
	ans := math.MinInt
	minS := map[int]int{}
	sum := 0
	for _, x := range nums {
		s, ok := minS[x+k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x-k]
		if ok {
			ans = max(ans, sum+x-s)
		}

		s, ok = minS[x]
		if !ok || sum < s {
			minS[x] = sum
		}

		sum += x
	}
	if ans == math.MinInt {
		return 0
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
