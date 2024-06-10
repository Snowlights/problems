#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。</p>

<p>如果下标三元组 <code>(i, j, k)</code> 满足下述全部条件，则认为它是一个 <strong>山形三元组</strong> ：</p>

<ul>
	<li><code>i < j < k</code></li>
	<li><code>nums[i] < nums[j]</code> 且 <code>nums[k] < nums[j]</code></li>
</ul>

<p>请你找出 <code>nums</code> 中 <strong>元素和最小</strong> 的山形三元组，并返回其 <strong>元素和</strong> 。如果不存在满足条件的三元组，返回 <code>-1</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [8,6,1,5,3]
<strong>输出：</strong>9
<strong>解释：</strong>三元组 (2, 3, 4) 是一个元素和等于 9 的山形三元组，因为： 
- 2 < 3 < 4
- nums[2] < nums[3] 且 nums[4] < nums[3]
这个三元组的元素和等于 nums[2] + nums[3] + nums[4] = 9 。可以证明不存在元素和小于 9 的山形三元组。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [5,4,8,7,10,2]
<strong>输出：</strong>13
<strong>解释：</strong>三元组 (1, 3, 5) 是一个元素和等于 13 的山形三元组，因为： 
- 1 < 3 < 5 
- nums[1] < nums[3] 且 nums[5] < nums[3]
这个三元组的元素和等于 nums[1] + nums[3] + nums[5] = 13 。可以证明不存在元素和小于 13 的山形三元组。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [6,5,4,3,4,5]
<strong>输出：</strong>-1
<strong>解释：</strong>可以证明 nums 中不存在山形三元组。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>8</sup></code></li>
</ul>

#### 思路

遇到这种三元组的题目，一个通常的做法是枚举中间的数。

知道了 $\textit{nums}[j]$，只需要知道 $j$ 左边的最小值和右边的最小值，就知道了三元组的和的最小值。
左右的最小值可以递推算出来。定义 $\textit{suf}[i]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 的最小值（后缀最小值），则有

$$
\textit{suf}[i] = \min(\textit{suf}[i+1], \textit{nums}[i])

$$

前缀最小值 $\textit{pre}$ 的计算方式同理，可以和答案一起算，所以只需要一个变量。
那么答案就是

$$
\textit{pre} + \textit{nums}[j] + \textit{suf}[j+1]

$$

的最小值。

```go
func minimumSum(nums []int) int {
	n := len(nums)
	suf := make([]int, n)
	suf[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans := math.MaxInt
	pre := nums[0]
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] {
			ans = min(ans, pre+nums[j]+suf[j+1])
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
