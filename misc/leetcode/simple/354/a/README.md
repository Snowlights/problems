#### 题目

<p>给你一个下标从 <strong>1</strong> 开始、长度为 <code>n</code> 的整数数组 <code>nums</code> 。</p>

<p>对 <code>nums</code> 中的元素 <code>nums[i]</code> 而言，如果 <code>n</code> 能够被 <code>i</code> 整除，即 <code>n % i == 0</code> ，则认为 <code>num[i]</code> 是一个 <strong>特殊元素</strong> 。</p>

<p>返回 <code>nums</code> 中所有 <strong>特殊元素</strong> 的 <strong>平方和</strong> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [1,2,3,4]
<strong>输出：</strong>21
<strong>解释：</strong>nums 中共有 3 个特殊元素：nums[1] ，因为 4 被 1 整除；nums[2] ，因为 4 被 2 整除；以及 nums[4] ，因为 4 被 4 整除。 
因此，nums 中所有元素的平方和等于 nums[1] * nums[1] + nums[2] * nums[2] + nums[4] * nums[4] = 1 * 1 + 2 * 2 + 4 * 4 = 21 。  
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [2,7,1,19,18,3]
<strong>输出：</strong>63
<strong>解释：</strong>nums 中共有 4 个特殊元素：nums[1] ，因为 6 被 1 整除；nums[2] ，因为 6 被 2 整除；nums[3] ，因为 6 被 3 整除；以及 nums[6] ，因为 6 被 6 整除。 
因此，nums 中所有元素的平方和等于 nums[1] * nums[1] + nums[2] * nums[2] + nums[3] * nums[3] + nums[6] * nums[6] = 2 * 2 + 7 * 7 + 1 * 1 + 3 * 3 = 63 。 </pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length == n <= 50</code></li>
	<li><code>1 <= nums[i] <= 50</code></li>
</ul>

#### 思路

按照题目遍历即可

```go
func sumOfSquares(nums []int) (ans int) {
	for i, x := range nums {
		if len(nums)%(i+1) == 0 {
			ans += x * x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。