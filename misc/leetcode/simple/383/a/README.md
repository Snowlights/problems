#### 题目

<p>边界上有一只蚂蚁，它有时向 <strong>左 </strong>走，有时向 <strong>右 </strong>走。</p>

<p>给你一个 <strong>非零</strong> 整数数组 <code>nums</code> 。蚂蚁会按顺序读取 <code>nums</code> 中的元素，从第一个元素开始直到结束。每一步，蚂蚁会根据当前元素的值移动：</p>

<ul>
	<li>如果 <code>nums[i] < 0</code> ，向 <strong>左</strong> 移动<!-- notionvc: 55fee232-4fc9-445f-952a-f1b979415864 --> <code>-nums[i]</code>单位。</li>
	<li>如果 <code>nums[i] > 0</code> ，向 <strong>右</strong> 移动 <code>nums[i]</code>单位。</li>
</ul>

<p>返回蚂蚁 <strong>返回 </strong>到边界上的次数。</p>

<p><strong>注意：</strong></p>

<ul>
	<li>边界两侧有无限的空间。</li>
	<li>只有在蚂蚁移动了 <code>|nums[i]|</code> 单位后才检查它是否位于边界上。换句话说，如果蚂蚁只是在移动过程中穿过了边界，则不会计算在内。<!-- notionvc: 5ff95338-8634-4d02-a085-1e83c0be6fcd --></li>
</ul>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,3,-5]
<strong>输出：</strong>1
<strong>解释：</strong>第 1 步后，蚂蚁距边界右侧 2 单位远<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->。
第 2 步后，蚂蚁距边界右侧 5 单位远<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->。
第 3 步后，蚂蚁位于边界上。
所以答案是 1 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [3,2,-3,-4]
<strong>输出：</strong>0
<strong>解释：</strong>第 1 步后，蚂蚁距边界右侧 3 单位远<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->。
第 2 步后，蚂蚁距边界右侧 5 单位远<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->。
第 3 步后，蚂蚁距边界右侧 2 单位远<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->。
第 4 步后，蚂蚁距边界左侧 2 单位远<!-- notionvc: 61ace51c-559f-4bc6-800f-0a0db2540433 -->。
蚂蚁从未返回到边界上，所以答案是 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 100</code></li>
	<li><code>-10 <= nums[i] <= 10</code></li>
	<li><code>nums[i] != 0</code></li>
</ul>

#### 思路

计算前缀和中0出现的次数

```go [sol]
func returnToBoundaryCount(nums []int) int {
	ans, cur := 0, 0
	for _, v := range nums {
		cur += v
		if cur == 0 {
			ans++
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。
