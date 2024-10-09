### 题目

<p>给你一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的数组 <code>nums</code> ，数组中的元素为 <strong>互不相同</strong> 的正整数。请你返回让 <code>nums</code> 成为递增数组的 <strong>最少右移</strong> 次数，如果无法得到递增数组，返回 <code>-1</code> 。</p>

<p>一次 <strong>右移</strong> 指的是同时对所有下标进行操作，将下标为 <code>i</code> 的元素移动到下标 <code>(i + 1) % n</code> 处。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,4,5,1,2]
<b>输出：</b>2
<b>解释：</b>
第一次右移后，nums = [2,3,4,5,1] 。
第二次右移后，nums = [1,2,3,4,5] 。
现在 nums 是递增数组了，所以答案为 2 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [1,3,5]
<b>输出：</b>0
<b>解释：</b>nums 已经是递增数组了，所以答案为 0 。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>nums = [2,1,4]
<b>输出：</b>-1
<b>解释：</b>无法将数组变为递增数组。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 100</code></li>
	<li><code>1 <= nums[i] <= 100</code></li>
	<li><code>nums</code> 中的整数互不相同。</li>
</ul>

### 思路

由于右移 $n$ 次就变回原数组了，所以答案至多为 $n-1$。
我们可以不断右移，每次右移前，先判断数组是否有序，如果有序就直接返回右移次数，否则就继续右移。
如果循环中途没有返回，最后返回 $-1$，表示无法得到递增数组。

```go
import "sort"

func minimumRightShifts(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	for i := 1; i <= len(nums); i++ {
		nums = append([]int{nums[len(nums)-1]}, nums[:len(nums)-1]...)
		if sort.IntsAreSorted(nums) {
			return i
		}
	}

	return -1
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n)$ 。
