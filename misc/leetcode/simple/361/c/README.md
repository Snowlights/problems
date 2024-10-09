#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，以及整数 <code>modulo</code> 和整数 <code>k</code> 。</p>

<p>请你找出并统计数组中 <strong>趣味子数组</strong> 的数目。</p>

<p>如果 <strong>子数组</strong> <code>nums[l..r]</code> 满足下述条件，则称其为 <strong>趣味子数组</strong> ：</p>

<ul>
	<li>在范围 <code>[l, r]</code> 内，设 <code>cnt</code> 为满足 <code>nums[i] % modulo == k</code> 的索引 <code>i</code> 的数量。并且 <code>cnt % modulo == k</code> 。</li>
</ul>

<p>以整数形式表示并返回趣味子数组的数目。<em> </em></p>

<p><span><strong>注意：</strong>子数组是数组中的一个连续非空的元素序列。</span></p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [3,2,4], modulo = 2, k = 1
<strong>输出：</strong>3
<strong>解释：</strong>在这个示例中，趣味子数组分别是： 
子数组 nums[0..0] ，也就是 [3] 。 
- 在范围 [0, 0] 内，只存在 1 个下标 i = 0 满足 nums[i] % modulo == k 。
- 因此 cnt = 1 ，且 cnt % modulo == k 。
子数组 nums[0..1] ，也就是 [3,2] 。
- 在范围 [0, 1] 内，只存在 1 个下标 i = 0 满足 nums[i] % modulo == k 。
- 因此 cnt = 1 ，且 cnt % modulo == k 。
子数组 nums[0..2] ，也就是 [3,2,4] 。
- 在范围 [0, 2] 内，只存在 1 个下标 i = 0 满足 nums[i] % modulo == k 。
- 因此 cnt = 1 ，且 cnt % modulo == k 。
可以证明不存在其他趣味子数组。因此，答案为 3 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [3,1,9,6], modulo = 3, k = 0
<strong>输出：</strong>2
<strong>解释：</strong>在这个示例中，趣味子数组分别是： 
子数组 nums[0..3] ，也就是 [3,1,9,6] 。
- 在范围 [0, 3] 内，只存在 3 个下标 i = 0, 2, 3 满足 nums[i] % modulo == k 。
- 因此 cnt = 3 ，且 cnt % modulo == k 。
子数组 nums[1..1] ，也就是 [1] 。
- 在范围 [1, 1] 内，不存在下标满足 nums[i] % modulo == k 。
- 因此 cnt = 0 ，且 cnt % modulo == k 。
可以证明不存在其他趣味子数组，因此答案为 2 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5 </sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>1 <= modulo <= 10<sup>9</sup></code></li>
	<li><code>0 <= k < modulo</code></li>
</ul>

#### 思路

对于本题，由于需要统计 $\textit{cnt}$，我们可以把满足 $\textit{nums}[i]\bmod \textit{modulo} = k$ 的 $\textit{nums}[i]$ 视作 $1$，不满足则视作 $0$。
如此转换后，算出 $\textit{nums}$ 的前缀和数组 $s$，那么题目中的 $\textit{cnt}\bmod \textit{modulo} = k$ 等价于

$$
(s[r+1]-s[l])\bmod \textit{modulo} = k

$$

上式等价于

$$
s[l]\bmod \textit{modulo} = (s[r+1]-k)\bmod \textit{modulo}

$$

根据上式，我们可以一边枚举 $r$，一边用一个哈希表统计有多少个 $s[r+1]\bmod k$。这样可以快速知道有多少个 $(s[r+1]-k)\bmod \textit{modulo}$，也就是 $s[l]\bmod \textit{modulo}$ 的个数，把个数加到答案中。
代码实现时，前缀和数组可以优化成一个变量 $s$。

```go
func countInterestingSubarrays(nums []int, mod, k int) (ans int64) {
	cnt := map[int]int{0: 1} // 把 s[0]=0 算进去
	s := 0
	for _, x := range nums {
		if x%mod == k {
			s = (s + 1) % mod // 这里取模，下面 cnt[s]++ 就不需要取模了
		}
		ans += int64(cnt[(s-k+mod)%mod]) // 避免减法出现负数
		cnt[s]++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
