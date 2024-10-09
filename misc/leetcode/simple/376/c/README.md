#### 题目

<p>给你一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。</p>

<p>你可以对 <code>nums</code> 执行特殊操作 <strong>任意次</strong> （也可以 <strong>0</strong> 次）。每一次特殊操作中，你需要 <strong>按顺序</strong> 执行以下步骤：</p>

<ul>
	<li>从范围 <code>[0, n - 1]</code> 里选择一个下标 <code>i</code> 和一个 <strong>正</strong> 整数 <code>x</code> 。</li>
	<li>将 <code>|nums[i] - x|</code> 添加到总代价里。</li>
	<li>将 <code>nums[i]</code> 变为 <code>x</code> 。</li>
</ul>

<p>如果一个正整数正着读和反着读都相同，那么我们称这个数是<strong> 回文数</strong> 。比方说，<code>121</code> ，<code>2552</code> 和 <code>65756</code> 都是回文数，但是 <code>24</code> ，<code>46</code> ，<code>235</code> 都不是回文数。</p>

<p>如果一个数组中的所有元素都等于一个整数 <code>y</code> ，且 <code>y</code> 是一个小于 <code>10<sup>9</sup></code> 的 <strong>回文数</strong> ，那么我们称这个数组是一个 <strong>等数数组 </strong>。</p>

<p>请你返回一个整数，表示执行任意次特殊操作后使 <code>nums</code> 成为 <strong>等数数组</strong> 的 <strong>最小</strong> 总代价。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [1,2,3,4,5]
<b>输出：</b>6
<b>解释：</b>我们可以将数组中所有元素变为回文数 3 得到等数数组，数组变成 [3,3,3,3,3] 需要执行 4 次特殊操作，代价为 |1 - 3| + |2 - 3| + |4 - 3| + |5 - 3| = 6 。
将所有元素变为其他回文数的总代价都大于 6 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [10,12,13,14,15]
<b>输出：</b>11
<b>解释：</b>我们可以将数组中所有元素变为回文数 11 得到等数数组，数组变成 [11,11,11,11,11] 需要执行 5 次特殊操作，代价为 |10 - 11| + |12 - 11| + |13 - 11| + |14 - 11| + |15 - 11| = 11 。
将所有元素变为其他回文数的总代价都大于 11 。
</pre>

<p><strong class="example">示例 3 ：</strong></p>

<pre>
<b>输入：</b>nums = [22,33,22,33,22]
<b>输出：</b>22
<b>解释：</b>我们可以将数组中所有元素变为回文数 22 得到等数数组，数组变为 [22,22,22,22,22] 需要执行 2 次特殊操作，代价为 |33 - 22| + |33 - 22| = 22 。
将所有元素变为其他回文数的总代价都大于 22 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>


## 前置知识：中位数贪心

为方便描述，将 $\textit{nums}$ 简记为 $a$。

**定理**：将 $a$ 的所有元素变为 $a$ 的**中位数**是最优的。

**证明**：设 $a$ 的长度为 $n$，设要将所有 $a[i]$ 变为 $x$。假设 $a$ 已经从小到大排序。首先，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之外，那么 $x$ 向区间方向移动可以使距离和变小；同时，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之内，无论如何移动 $x$，它到 $a[0]$ 和 $a[n-1]$ 的距离和都是一个定值 $a[n-1]-a[0]$，那么去掉 $a[0]$ 和 $a[n-1]$ 这两个最左最右的数，问题规模缩小。不断缩小问题规模，如果最后剩下 $1$ 个数，那么 $x$ 就取它；如果最后剩下 $2$ 个数，那么 $x$ 取这两个数之间的任意值都可以（包括这两个数）。因此 $x$ 可以取 $a[n/2]$。

本题回文数可能取不到中位数，我们可以找离中位数最近的数。

#### 思路

首先预处理出 $10^9$ 内的回文数，这可以通过枚举回文数的左半部分得到。

> 注：也可以多预处理一个 $10^9+1$，不过考虑到 $10^9-1$ 也是回文数，这一步可以省略。

然后二分找离 $\textit{nums}$ 的中位数最近的回文数（中位数左右两侧都要找），作为要变成的数字。

```go  
func minimumCost(nums []int) int64 {
	pal := make([]int, 0, 109999)
	// 哨兵。根据题目来定，也可以设置成 -2e9 等
	pal = append(pal, 0)
	initPalindromeNumber := func() {
		const mx int = 1e9
		// 严格按顺序从小到大生成所有回文数
	outer:
		for base := 1; ; base *= 10 {
			// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
			for i := base; i < base*10; i++ {
				x := i
				for t := i / 10; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
			// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
		}

		// 哨兵。根据 mx 调整，如果 mx 是 2e9 的话要写成 mx+2
		pal = append(pal, mx+1)
	}
	initPalindromeNumber()

	sort.Ints(nums)
	n := len(nums)
	i := sort.SearchInts(pal, nums[n/2])
	cost := func(i int) int{
		res := 0
		for _, v := range nums {
			res += abs(v-i)
		}
		return res
	}
	return int64(min(cost(pal[i-1]), cost(pal[i])))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + \log U)$ 或 $\mathcal{O}(n + \log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U$ 为 $10^9+1$ 内的回文数个数。忽略预处理的时间。
- 空间复杂度：$\mathcal{O}(1)$。忽略预处理的空间。

## 思考题

能否直接找离某个数最近的回文数？

这题是 [564. 寻找最近的回文数](https://leetcode.cn/problems/find-the-closest-palindrome/)，注意数据范围。

本题也可以从中位数出发，向左右暴力找回文数。在本题的数据范围下，最大的回文数间隔只有 $10011001 - 10000001 = 11000$。

## 生成回文数

- [2081. k 镜像数字的和](https://leetcode.cn/problems/sum-of-k-mirror-numbers/)

## 中位数贪心（右边数字为难度分）

- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- [2033. 获取单值网格的最小操作数](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/) 1672
- [2448. 使数组相等的最小开销](https://leetcode.cn/problems/minimum-cost-to-make-array-equal/) 2005
- [2607. 使子数组元素和相等](https://leetcode.cn/problems/make-k-subarray-sums-equal/) 2071
- [1703. 得到连续 K 个 1 的最少相邻交换次数](https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/) 2467
