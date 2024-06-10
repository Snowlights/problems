#### 题目

<p>给你两个由正整数和 <code>0</code> 组成的数组 <code>nums1</code> 和 <code>nums2</code> 。</p>

<p>你必须将两个数组中的<strong> 所有</strong> <code>0</code> 替换为 <strong>严格</strong> 正整数，并且满足两个数组中所有元素的和 <strong>相等</strong> 。</p>

<p>返回 <strong>最小</strong> 相等和 ，如果无法使两数组相等，则返回 <code>-1</code><em> </em>。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [3,2,0,1,0], nums2 = [6,5,0]
<strong>输出：</strong>12
<strong>解释：</strong>可以按下述方式替换数组中的 0 ：
- 用 2 和 4 替换 nums1 中的两个 0 。得到 nums1 = [3,2,2,1,4] 。
- 用 1 替换 nums2 中的一个 0 。得到 nums2 = [6,5,1] 。
两个数组的元素和相等，都等于 12 。可以证明这是可以获得的最小相等和。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums1 = [2,0,2,0], nums2 = [1,4]
<strong>输出：</strong>-1
<strong>解释：</strong>无法使两个数组的和相等。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums1.length, nums2.length <= 10<sup>5</sup></code></li>
	<li><code>0 <= nums1[i], nums2[i] <= 10<sup>6</sup></code></li>
</ul>

#### 思路

把 $0$ 看成 $1$，设 $s_1$ 为 $\textit{nums}_1$ 的元素和，
$s_2$ 为 $\textit{nums}_2$ 的元素和。这是元素和的最小值。

分类讨论：

- 如果 $\textit{nums}_1$ 中有 $0$，但 $\textit{nums}_2$ 中没有 $0$，并且 $s_1 > s_2$，那么 $s_2$ 无法增加，返回 $-1$。
- 如果 $\textit{nums}_2$ 中有 $0$，但 $\textit{nums}_1$ 中没有 $0$，并且 $s_1 < s_2$，那么 $s_1$ 无法增加，返回 $-1$。
- 如果 $\textit{nums}_1$ 和 $\textit{nums}_1$ 中都没有 $0$，并且 $s_1 \ne s_2$，那么无法相等，返回 $-1$。
- 否则，答案为 $\max(s_1,s_2)$。

```go  
func minSum(nums1, nums2 []int) int64 {
	s1 := int64(0)
	zero1 := false
	for _, x := range nums1 {
		if x == 0 {
			zero1 = true
			s1++
		} else {
			s1 += int64(x)
		}
	}

	s2 := int64(0)
	zero2 := false
	for _, x := range nums2 {
		if x == 0 {
			zero2 = true
			s2++
		} else {
			s2 += int64(x)
		}
	}

	if zero1 && !zero2 && s1 > s2 ||
		!zero1 && zero2 && s1 < s2 ||
		!zero1 && !zero2 && s1 != s2 {
		return -1
	}
	return max(s1, s2)
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
