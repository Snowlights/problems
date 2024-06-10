#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个整数 <code>k</code> 。</p>

<p>你可以对数组执行 <strong>至多</strong> <code>k</code> 次操作：</p>

<ul>
	<li>从数组中选择一个下标 <code>i</code> ，将 <code>nums[i]</code> <strong>增加</strong> 或者 <strong>减少</strong> <code>1</code> 。</li>
</ul>

<p>最终数组的频率分数定义为数组中众数的 <strong>频率</strong> 。</p>

<p>请你返回你可以得到的 <strong>最大</strong> 频率分数。</p>

<p>众数指的是数组中出现次数最多的数。一个元素的频率指的是数组中这个元素的出现次数。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,6,4], k = 3
<b>输出：</b>3
<b>解释：</b>我们可以对数组执行以下操作：
- 选择 i = 0 ，将 nums[0] 增加 1 。得到数组 [2,2,6,4] 。
- 选择 i = 3 ，将 nums[3] 减少 1 ，得到数组 [2,2,6,3] 。
- 选择 i = 3 ，将 nums[3] 减少 1 ，得到数组 [2,2,6,2] 。
元素 2 是最终数组中的众数，出现了 3 次，所以频率分数为 3 。
3 是所有可行方案里的最大频率分数。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [1,4,4,2,4], k = 0
<b>输出：</b>3
<b>解释：</b>我们无法执行任何操作，所以得到的频率分数是原数组中众数的频率 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>0 <= k <= 10<sup>14</sup></code></li>
</ul>

## 前置知识：中位数贪心

为方便描述，将 $\textit{nums}$ 简记为 $a$。

**定理**：将 $a$ 的所有元素变为 $a$ 的**中位数**是最优的。

**证明**：设 $a$ 的长度为 $n$，设要将所有 $a[i]$ 变为 $x$。假设 $a$ 已经从小到大排序。首先，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之外，那么 $x$ 向区间方向移动可以使距离和变小；同时，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之内，无论如何移动 $x$，它到 $a[0]$ 和 $a[n-1]$ 的距离和都是一个定值 $a[n-1]-a[0]$，那么去掉 $a[0]$ 和 $a[n-1]$ 这两个最左最右的数，问题规模缩小。不断缩小问题规模，如果最后剩下 $1$ 个数，那么 $x$ 就取它；如果最后剩下 $2$ 个数，那么 $x$ 取这两个数之间的任意值都可以（包括这两个数）。因此 $x$ 可以取 $a[n/2]$。

#### 思路

把数组排序后，要变成一样的数必然在一个连续子数组中，那么用**滑动窗口**来做，枚举子数组的右端点 $\textit{right}$，然后维护子数组的左端点 $\textit{left}$。

根据中位数贪心，最优做法是把子数组内的元素都变成子数组的中位数，操作次数如果超过 $k$，就必须移动左端点。

求出数组的前缀和，就可以 $\mathcal{O}(1)$ 算出操作次数了，具体请

<p><img alt="" src="https://pic.leetcode.cn/1679808210-FVsAou-t3.png" /></p>

```go  
func maxFrequencyScore(nums []int, k int64) int {

	sort.Ints(nums)
	pre := make([]int, len(nums)+1)
	for i, v := range nums {
		pre[i+1] = pre[i] + v
	}

	cost := func(l, m, r int) int {
		left := (m-l)*nums[m] - (pre[m] - pre[l])
		right := pre[r+1] - pre[m+1] - nums[m]*(r-m)
		return left + right
	}
	l, ans := 0, 1
	for r := range nums {
		for int64(cost(l, (l+r)>>1, r)) > k {
			l++
		}
		ans = max(ans, r-l+1)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。
