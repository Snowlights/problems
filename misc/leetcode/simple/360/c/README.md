#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的数组 <code>nums</code> ，它包含 <strong>非负</strong> 整数，且全部为 <code>2</code> 的幂，同时给你一个整数 <code>target</code> 。</p>

<p>一次操作中，你必须对数组做以下修改：</p>

<ul>
	<li>选择数组中一个元素 <code>nums[i]</code> ，满足 <code>nums[i] > 1</code> 。</li>
	<li>将 <code>nums[i]</code> 从数组中删除。</li>
	<li>在 <code>nums</code> 的 <strong>末尾</strong> 添加 <strong>两个</strong> 数，值都为 <code>nums[i] / 2</code> 。</li>
</ul>

<p>你的目标是让 <code>nums</code> 的一个 <strong>子序列</strong> 的元素和等于 <code>target</code> ，请你返回达成这一目标的 <strong>最少操作次数</strong> 。如果无法得到这样的子序列，请你返回 <code>-1</code> 。</p>

<p>数组中一个 <strong>子序列</strong> 是通过删除原数组中一些元素，并且不改变剩余元素顺序得到的剩余数组。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><b>输入：</b>nums = [1,2,8], target = 7
<b>输出：</b>1
<b>解释：</b>第一次操作中，我们选择元素 nums[2] 。数组变为 nums = [1,2,4,4] 。
这时候，nums 包含子序列 [1,2,4] ，和为 7 。
无法通过更少的操作得到和为 7 的子序列。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>nums = [1,32,1,2], target = 12
<b>输出：</b>2
<b>解释：</b>第一次操作中，我们选择元素 nums[1] 。数组变为 nums = [1,1,2,16,16] 。
第二次操作中，我们选择元素 nums[3] 。数组变为 nums = [1,1,2,16,8,8] 。
这时候，nums 包含子序列 [1,1,2,8] ，和为 12 。
无法通过更少的操作得到和为 12 的子序列。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre><b>输入：</b>nums = [1,32,1], target = 35
<b>输出：</b>-1
<b>解释：</b>无法得到和为 35 的子序列。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 1000</code></li>
	<li><code>1 <= nums[i] <= 2<sup>30</sup></code></li>
	<li><code>nums</code> 只包含非负整数，且均为 2 的幂。</li>
	<li><code>1 <= target < 2<sup>31</sup></code></li>
</ul>

#### 思路

由于可以把一个数一分为二，所以整个数组可以全部变成 $1$。因此如果 $\textit{nums}$ 的元素和小于 $\textit{target}$，则无解，返回 $-1$。否则一定有解。
然后从低位到高位贪心：

- 如果 $\textit{target}$ 的第 $i$ 位是 $0$，跳过。
- 如果 $\textit{target}$ 的第 $i$ 位是 $1$，那么先看看所有 $\le 2^i$ 的元素和 $s$ 能否 $\ge \textit{target}\&  \textit{mask}$，其中 $\textit{mask}=2^{i+1}-1$。如果能，那么无需操作。
- 如果不能，那么就需要把一个更大的数不断地一分为二，直到分解出 $2^i$ 为止。

```go  
func minOperations(nums []int, target int) (ans int) {
	s := 0
	cnt := [31]int{}
	for _, v := range nums {
		s += v
		cnt[bits.TrailingZeros(uint(v))]++
	}
	if s < target {
		return -1
	}
	s = 0
	for i := 0; i < 31; i++ {
		s += cnt[i] << i
		mask := 1<<(i+1) - 1
		if s < target&mask {
			for j := i + 1; ; j++ {
				if cnt[j] > 0 {
					ans += j - i
					cnt[j]--
					s += 1 << j // 从 i 到 j 都至少有一个 1，可以提前加到 s 中
					break
				}
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\log \textit{target})$。
