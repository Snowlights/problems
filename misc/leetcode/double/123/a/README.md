### 题目

<p>给你一个下标从 <strong>0</strong> 开始长度为 <code>3</code> 的整数数组 <code>nums</code> ，需要用它们来构造三角形。</p>

<ul>
	<li>如果一个三角形的所有边长度相等，那么这个三角形称为 <strong>equilateral</strong> 。</li>
	<li>如果一个三角形恰好有两条边长度相等，那么这个三角形称为 <strong>isosceles</strong> 。</li>
	<li>如果一个三角形三条边的长度互不相同，那么这个三角形称为 <strong>scalene</strong> 。</li>
</ul>

<p>如果这个数组无法构成一个三角形，请你返回字符串 <code>"none"</code> ，否则返回一个字符串表示这个三角形的类型。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,3,3]
<b>输出：</b>"equilateral"
<b>解释：</b>由于三条边长度相等，所以可以构成一个等边三角形，返回 "equilateral" 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [3,4,5]
<b>输出：</b>"scalene"
<b>解释：</b>
nums[0] + nums[1] = 3 + 4 = 7 ，大于 nums[2] = 5<code> 。</code>
nums[0] + nums[2] = 3 + 5 = 8 ，大于 nums[1] = 4 。
nums[1] + nums[2] = 4 + 5 = 9 ，大于 nums[0] = 3 。
由于任意两边纸盒都大于第三边，所以可以构成一个三角形。
因为三条边的长度互不相等，所以返回 "scalene" 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>nums.length == 3</code></li>
	<li><code>1 <= nums[i] <= 100</code></li>
</ul>

### 思路

按照题意模拟即可

```go [sol]
func triangleType(nums []int) string {
	sort.Ints(nums)
	if nums[0]+nums[1] <= nums[2] {
		return "none"
	} else if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[1] == nums[2] {
		return "isosceles"
	}
	return "scalene"
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
