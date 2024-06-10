#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。</p>

<p>请你从所有满足&nbsp;<code>i &lt; j &lt; k</code> 的下标三元组 <code>(i, j, k)</code> 中，找出并返回下标三元组的最大值。如果所有满足条件的三元组的值都是负数，则返回 <code>0</code> 。</p>

<p><strong>下标三元组</strong> <code>(i, j, k)</code> 的值等于 <code>(nums[i] - nums[j]) * nums[k]</code> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [12,6,1,2,7]
<strong>输出：</strong>77
<strong>解释：</strong>下标三元组 (0, 2, 4) 的值是 (nums[0] - nums[2]) * nums[4] = 77 。
可以证明不存在值大于 77 的有序下标三元组。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,10,3,4,19]
<strong>输出：</strong>133
<strong>解释：</strong>下标三元组 (1, 2, 4) 的值是 (nums[1] - nums[2]) * nums[4] = 133 。
可以证明不存在值大于 133 的有序下标三元组。 
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3]
<strong>输出：</strong>0
<strong>解释：</strong>唯一的下标三元组 (0, 1, 2) 的值是一个负数，(nums[0] - nums[1]) * nums[2] = -3 。因此，答案是 0 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>

#### 思路

## 方法一：枚举 j

枚举 $j$，为了让 $(\\textit{nums}[i] - \\textit{nums}[j]) * \\textit{nums}[k]$ 尽量大，我们需要知道 $j$ 左侧元素的最大值，和 $j$ 右侧元素的最大值。  
也就是 $\\textit{nums}$ 的前缀最大值 $\\textit{preMax}$ 和后缀最大值 $\\textit{sufMax}$，这都可以用递推预处理出来：
- $\\textit{preMax}[i] = \\max(\\textit{preMax}[i-1], \\textit{nums}[i])$\r\n- $\\textit{sufMax}[i] = \\max(\\textit{sufMax}[i+1], \\textit{nums}[i])$  
  
代码实现时，可以只预处理 $\\textit{sufMax}$ 数组，$\\textit{preMax}$ 可以在计算答案的同时算出来。

```go  
func maximumTripletValue(nums []int) int64 {
	ans := 0
	n := len(nums)
	sufMax := make([]int, n+1)
	for i := n - 1; i > 1; i-- {
		sufMax[i] = max(sufMax[i+1], nums[i])
	}
	preMax := 0
	for j, x := range nums {
		ans = max(ans, (preMax-x)*sufMax[j+1])
		preMax = max(preMax, x)
	}
	return int64(ans)
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：枚举 k

枚举 $k$，我们需要知道 $k$ 左边 $\\textit{nums}[i] - \\textit{nums}[j]$ 的最大值。

类似 [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)，我们可以在遍历的过程中，维护 $\\textit{nums}[i]$ 的最大值 $\\textit{preMax}$，同时维护 $\\textit{preMax}$ 减当前元素的最大值 $\\textit{maxDiff}$，这就是 $k$ 左边 $\\textit{nums}[i] - \\textit{nums}[j]$ 的最大值。

```go  
func maximumTripletValue(nums []int) int64 {
	ans, maxDiff, preMax := 0, 0, 0
	for _, x := range nums {
		ans = max(ans, maxDiff*x)
		maxDiff = max(maxDiff, preMax-x)
		preMax = max(preMax, x)
	}
	return int64(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
