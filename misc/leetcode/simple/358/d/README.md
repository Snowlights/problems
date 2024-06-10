#### 题目

<p>给你一个长度为 <code>n</code> 的正整数数组 <code>nums</code> 和一个整数 <code>k</code> 。</p>

<p>一开始，你的分数为 <code>1</code> 。你可以进行以下操作至多 <code>k</code> 次，目标是使你的分数最大：</p>

<ul>
	<li>选择一个之前没有选过的 <strong>非空</strong> 子数组 <code>nums[l, ..., r]</code> 。</li>
	<li>从 <code>nums[l, ..., r]</code> 里面选择一个 <strong>质数分数</strong> 最高的元素 <code>x</code> 。如果多个元素质数分数相同且最高，选择下标最小的一个。</li>
	<li>将你的分数乘以 <code>x</code> 。</li>
</ul>

<p><code>nums[l, ..., r]</code> 表示 <code>nums</code> 中起始下标为 <code>l</code> ，结束下标为 <code>r</code> 的子数组，两个端点都包含。</p>

<p>一个整数的 <strong>质数分数</strong> 等于 <code>x</code> 不同质因子的数目。比方说， <code>300</code> 的质数分数为 <code>3</code> ，因为 <code>300 = 2 * 2 * 3 * 5 * 5</code> 。</p>

<p>请你返回进行至多 <code>k</code> 次操作后，可以得到的 <strong>最大分数</strong> 。</p>

<p>由于答案可能很大，请你将结果对 <code>10<sup>9 </sup>+ 7</code> 取余后返回。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><b>输入：</b>nums = [8,3,9,3,8], k = 2
<b>输出：</b>81
<b>解释：</b>进行以下操作可以得到分数 81 ：
- 选择子数组 nums[2, ..., 2] 。nums[2] 是子数组中唯一的元素。所以我们将分数乘以 nums[2] ，分数变为 1 * 9 = 9 。
- 选择子数组 nums[2, ..., 3] 。nums[2] 和 nums[3] 质数分数都为 1 ，但是 nums[2] 下标更小。所以我们将分数乘以 nums[2] ，分数变为 9 * 9 = 81 。
81 是可以得到的最高得分。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>nums = [19,12,14,6,10,18], k = 3
<b>输出：</b>4788
<b>解释：</b>进行以下操作可以得到分数 4788 ：
- 选择子数组 nums[0, ..., 0] 。nums[0] 是子数组中唯一的元素。所以我们将分数乘以 nums[0] ，分数变为 1 * 19 = 19 。
- 选择子数组 nums[5, ..., 5] 。nums[5] 是子数组中唯一的元素。所以我们将分数乘以 nums[5] ，分数变为 19 * 18 = 342 。
- 选择子数组 nums[2, ..., 3] 。nums[2] 和 nums[3] 质数分数都为 2，但是 nums[2] 下标更小。所以我们将分数乘以 nums[2] ，分数变为  342 * 14 = 4788 。
4788 是可以得到的最高的分。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length == n <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>5</sup></code></li>
	<li><code>1 <= k <= min(n * (n + 1) / 2, 10<sup>9</sup>)</code></li>
</ul>

#### 思路

贪心地说，先考虑 $\textit{nums}$ 中最大的元素 $x$，我们需要知道：有多少个子数组可以取 $x$ 作为质数分数最高的元素。
我们可以先把 $[1,10^5]$ 中的每个数的质数分数（不同质因子的数目）预处理出来，记作数组 $\textit{omega}$。
然后用单调栈求出每个 $i$ 左侧质数分数【大于等于】$\textit{omega}[\textit{nums}[i]]$ 的最近的数的下标 $\textit{left}[i]$（不存在则为 $-1$），以及右侧质数分数【大于】$\textit{omega}[\textit{nums}[i]]$ 的最近的数的下标 $\textit{right}[i]$（不存在则为 $n$）。
请注意，我们不能在 $i$ 左侧包含质数分数和 $\textit{omega}[\textit{nums}[i]]$ 一样的数，否则不满足题目「如果多个元素质数分数相同且最高，选择下标最小的一个」的要求。所以对于左侧用【大于等于】。
那么子数组的左端点可以在 $(\textit{left}[i],i]$ 内，子数组的右端点可以在 $[i,\textit{right}[i])$ 内。
根据 [乘法原理](https://baike.baidu.com/item/%E4%B9%98%E6%B3%95%E5%8E%9F%E7%90%86/7538447)，一共有

$$
\textit{tot} = (i-\textit{left}[i])\cdot (\textit{right}[i]-i)

$$

个子数组以 $\textit{nums}[i]$ 作为这个子数组的质数分数。
设 $k'=\min(k, \textit{tot})$，则 $\textit{nums}[i]$ 对答案的贡献为

$$
\textit{nums}[i] ^ {k'}

$$

上式可以用快速幂计算，具体请看 [50. Pow(x, n)](https://leetcode.cn/problems/powx-n/)。
根据上式，按照 $\textit{nums}$ 从大到小的顺序计算贡献，一边计算一边更新剩余操作次数 $k$。

```go  
const mod int = 1e9 + 7

// 预处理 omega
const mx int = 1e5
var omega [mx + 1]int8
func init() {
	for i := 2; i <= mx; i++ {
		if omega[i] == 0 { // i 是质数
			for j := i; j <= mx; j += i {
				omega[j]++ // i 是 j 的一个质因子
			}
		}
	}
}

func maximumScore(nums []int, k int) int {
	n := len(nums)
	left := make([]int, n)  // 质数分数 >= omega[nums[i]] 的左侧最近元素下标
	right := make([]int, n) // 质数分数 >  omega[nums[i]] 的右侧最近元素下标
	for i := range right {
		right[i] = n
	}
	st := []int{-1}
	for i, v := range nums {
		for len(st) > 1 && omega[nums[st[len(st)-1]]] < omega[v] {
			right[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return nums[id[i]] > nums[id[j]] })

	ans := 1
	for _, i := range id {
		tot := (i - left[i]) * (right[i] - i)
		if tot >= k {
			ans = ans * pow(nums[i], k) % mod
			break
		}
		ans = ans * pow(nums[i], tot) % mod
		k -= tot // 更新剩余操作次数
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。这里忽略预处理 $\textit{omega}$ 的时间和空间。
- 空间复杂度：$\mathcal{O}(n)$。
