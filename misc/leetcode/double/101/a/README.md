### 题目  

给你两个只包含 1 到 9 之间数字的数组 <code>nums1</code> 和 <code>nums2</code> ，每个数组中的元素 <strong>互不相同</strong> ，请你返回 <strong>最小</strong> 的数字，两个数组都 <strong>至少</strong> 包含这个数字的某个数位。
<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums1 = [4,1,3], nums2 = [5,7]
<b>输出：</b>15
<b>解释：</b>数字 15 的数位 1 在 nums1 中出现，数位 5 在 nums2 中出现。15 是我们能得到的最小数字。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums1 = [3,5,2,6], nums2 = [3,1,7]
<b>输出：</b>3
<b>解释：</b>数字 3 的数位 3 在两个数组中都出现了。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length, nums2.length &lt;= 9</code></li>
	<li><code>1 &lt;= nums1[i], nums2[i] &lt;= 9</code></li>
	<li>每个数组中，元素 <strong>互不相同</strong> 。</li>
</ul>
 
### 思路  

如果 $\textit{nums}_1$ 与 $\textit{nums}_2$ 有交集，那么答案就是交集的最小值。  
否则取两个数组的最小值，记作 $x$ 和 $y$，那么答案就是 $\min(10x+y, 10y+x)$。

```go 
func minNumber(nums1, nums2 []int) int {
	var mask1, mask2 uint
	for _, x := range nums1 {
		mask1 |= 1 << x
	}
	for _, x := range nums2 {
		mask2 |= 1 << x
	}
	if m := mask1 & mask2; m > 0 {
		return bits.TrailingZeros(m)
	}
	x, y := bits.TrailingZeros(mask1), bits.TrailingZeros(mask2)
	return min(x*10+y, y*10+x)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
```

### 复杂度分析  

- 时间复杂度：$O(n+m)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$m$ 为 $\textit{nums}_2$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。