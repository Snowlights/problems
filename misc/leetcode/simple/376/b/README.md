#### 题目

<p>给你一个长度为 <code>n</code> 的整数数组 <code>nums</code>，以及一个正整数 <code>k</code> 。</p>

<p>将这个数组划分为一个或多个长度为 <code>3</code> 的子数组，并满足以下条件：</p>

<ul>
	<li><code>nums</code> 中的 <strong>每个 </strong>元素都必须 <strong>恰好 </strong>存在于某个子数组中。</li>
	<li>子数组中<strong> 任意 </strong>两个元素的差必须小于或等于 <code>k</code> 。</li>
</ul>

<p>返回一个<em> </em><strong>二维数组 </strong>，包含所有的子数组。如果不可能满足条件，就返回一个空数组。如果有多个答案，返回 <strong>任意一个</strong> 即可。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,3,4,8,7,9,3,5,1], k = 2
<strong>输出：</strong>[[1,1,3],[3,4,5],[7,8,9]]
<strong>解释：</strong>可以将数组划分为以下子数组：[1,1,3]，[3,4,5] 和 [7,8,9] 。
每个子数组中任意两个元素的差都小于或等于 2 。
注意，元素的顺序并不重要。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,3,3,2,7,3], k = 3
<strong>输出：</strong>[]
<strong>解释：</strong>无法划分数组满足所有条件。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums.length</code></li>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
	<li><code>n</code> 是 <code>3</code> 的倍数</li>
	<li><code>1 <= nums[i] <= 10<sup>5</sup></code></li>
	<li><code>1 <= k <= 10<sup>5</sup></code></li>
</ul>

#### 思路

注意本题的子数组不要求是连续的。

既然元素的顺序并不重要，我们应当尽量把相近的数字都放在一起。

所以把数组排序后，从小到大三个三个地切分即可。

> 注：题目保证数组长度是 $3$ 的倍数。

```go  
func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 2; i < len(nums); i+=3 {
		if nums[i] - nums[i-2] <= k {
			ans = append(ans, nums[i-2:i+1])
		}
	}
	if len(ans) * 3 != len(nums) {
		return nil
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。
