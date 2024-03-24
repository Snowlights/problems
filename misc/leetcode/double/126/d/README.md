### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的整数数组&nbsp;<code>nums</code>&nbsp;和一个 <strong>正</strong>&nbsp;整数&nbsp;<code>k</code>&nbsp;。</p>

<p>一个整数数组的 <strong>能量</strong>&nbsp;定义为和 <strong>等于</strong>&nbsp;<code>k</code>&nbsp;的子序列的数目。</p>

<p>请你返回 <code>nums</code>&nbsp;中所有子序列的 <strong>能量和</strong>&nbsp;。</p>

<p>由于答案可能很大，请你将它对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;"> nums = [1,2,3], k = 3 </span></p>

<p><strong>输出：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;"> 6 </span></p>

<p><strong>解释：</strong></p>

<p>总共有&nbsp;<code>5</code>&nbsp;个能量不为 0 的子序列：</p>

<ul>
	<li>子序列&nbsp;<code>[<em><strong>1</strong></em>,<em><strong>2</strong></em>,<em><strong>3</strong></em>]</code> 有&nbsp;<code>2</code>&nbsp;个和为&nbsp;<code>3</code>&nbsp;的子序列：<code>[1,2,<strong><em>3</em></strong>]</code> 和 <code>[<strong><em>1</em></strong>,<strong><em>2</em></strong>,3]</code>&nbsp;。</li>
	<li>子序列&nbsp;<code>[<em><strong>1</strong></em>,2,<em><strong>3</strong></em>]</code>&nbsp;有 <code>1</code>&nbsp;个和为&nbsp;<code>3</code>&nbsp;的子序列：<code>[1,2,<strong><em>3</em></strong>]</code>&nbsp;。</li>
	<li>子序列&nbsp;<code>[1,<em><strong>2</strong></em>,<em><strong>3</strong></em>]</code> 有&nbsp;<code>1</code>&nbsp;个和为&nbsp;<code>3</code>&nbsp;的子序列：<code>[1,2,<strong><em>3</em></strong>]</code>&nbsp;。</li>
	<li>子序列&nbsp;<code>[<em><strong>1</strong></em>,<em><strong>2</strong></em>,3]</code>&nbsp;有&nbsp;<code>1</code>&nbsp;个和为&nbsp;<code>3</code>&nbsp;的子序列：<code>[<strong><em>1</em></strong>,<strong><em>2</em></strong>,3]</code>&nbsp;。</li>
	<li>子序列&nbsp;<code>[1,2,<em><strong>3</strong></em>]</code>&nbsp;有&nbsp;<code>1</code>&nbsp;个和为&nbsp;<code>3</code>&nbsp;的子序列：<code>[1,2,<strong><em>3</em></strong>]</code>&nbsp;。</li>
</ul>

<p>所以答案为&nbsp;<code>2 + 1 + 1 + 1 + 1 = 6</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;"> nums = [2,3,3], k = 5 </span></p>

<p><strong>输出：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;"> 4 </span></p>

<p><strong>解释：</strong></p>

<p>总共有&nbsp;<code>3</code>&nbsp;个能量不为 0 的子序列：</p>

<ul>
	<li>子序列&nbsp;<code>[<em><strong>2</strong></em>,<em><strong>3</strong></em>,<em><strong>3</strong></em>]</code>&nbsp;有 2 个子序列和为&nbsp;<code>5</code>&nbsp;：<code>[<strong><em>2</em></strong>,3,<strong><em>3</em></strong>]</code> 和&nbsp;<code>[<strong><em>2</em></strong>,<strong><em>3</em></strong>,3]</code>&nbsp;。</li>
	<li>子序列&nbsp;<code>[<em><strong>2</strong></em>,3,<em><strong>3</strong></em>]</code>&nbsp;有 1 个子序列和为&nbsp;<code>5</code>&nbsp;：<code>[<strong><em>2</em></strong>,3,<strong><em>3</em></strong>]</code>&nbsp;。</li>
	<li>子序列&nbsp;<code>[<em><strong>2</strong></em>,<em><strong>3</strong></em>,3]</code>&nbsp;有 1 个子序列和为 <code>5</code>&nbsp;：<code>[<strong><em>2</em></strong>,<strong><em>3</em></strong>,3]</code>&nbsp;。</li>
</ul>

<p>所以答案为&nbsp;<code>2 + 1 + 1 = 4</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;"> nums = [1,2,3], k = 7 </span></p>

<p><strong>输出：</strong> <span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;"> 0 </span></p>

<p><strong>解释：</strong>不存在和为 <code>7</code>&nbsp;的子序列，所以 <code>nums</code>&nbsp;的能量和为&nbsp;<code>0</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= k &lt;= 100</code></li>
</ul>

### 思路

## 方法一：二维 0-1 背包

#### 观察

如果和为 $k$ 的子序列 $S$ 的长度是 $c$，那么有多少个子序列 $T$，会包含 $S$ 呢？即 $S$ 是 $T$ 的子序列。

例如示例 1，子序列 $S=[3]$，出现在子序列 $[1,2,3],[1,3],[2,3],[3]$ 中，这 $4$ 个子序列都可以是 $T$。

除了 $3$ 以外的每个数，都可以选择在/不在包含 $[3]$ 的子序列 $T$ 中。

所以有 $2^{n-c}$ 个子序列 $T$。

这意味着 $S$ 对答案的贡献是 $2^{n-c}$。

#### 思路

枚举和为 $k$ 的子序列的长度 $c$，问题变成：

- 有 $n$ 个物品，每个物品的体积是 $\textit{nums}[i]$，问恰好装满容量为 $k$ 的背包，且选择 $c$ 个物品的方案数。

这可以用二维 0-1 背包解决。不了解 0-1 背包的同学请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

定义 $f[i][j][c]$ 表示考虑前 $i$ 个物品，所选物品体积和是 $j$，选了 $c$ 个物品的方案数。

考虑第 $i$ 个物品选或不选：

- 不选：$f[i+1][j][c] = f[i][j][c]$。
- 选：$f[i+1][j][c] = f[i][j-\textit{nums}[i]][c-1]$。

两者相加，得

$$
f[i+1][j][c] = f[i][j][c] + f[i][j-\textit{nums}[i]][c-1]
$$

初始值：$f[0][0][0] = 1$。

答案：$\sum\limits_{c=1}^{n} f[n][k][c] \cdot 2^{n-c}$。

代码实现时，第一个维度可以优化掉。

```go [sol-Go]
func sumOfPower(nums []int, k int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for i, x := range nums {
		for j := k; j >= x; j-- {
			for c := i + 1; c > 0; c-- {
				f[j][c] = (f[j][c] + f[j-x][c-1]) % mod
			}
		}
	}
	pow2 := 1
	for i := n; i > 0; i-- {
		ans = (ans + f[k][i]*pow2) % mod
		pow2 = pow2 * 2 % mod
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

## 方法二：一维 0-1 背包

定义 $f[i][j]$ 表示考虑前 $i$ 个数，所选元素和是 $j$ 时的能量和。

对于 $\textit{nums}[i]$，它只有三种可能：

- 不在子序列 $S$，也不在子序列 $T$ 中：$f[i+1][j] = f[i][j]$。
- 不在子序列 $S$，在子序列 $T$ 中：$f[i+1][j] = f[i][j]$。
- 在子序列 $S$ 中：$f[i+1][j] = f[i][j-\textit{nums}[i]]$。

三者相加，得

$$
f[i+1][j] = f[i][j]\cdot 2 + f[i][j-\textit{nums}[i]]
$$

初始值：$f[0][0] = 1$。

答案：$f[n][k]$。

代码实现时，第一个维度可以优化掉。

```go [sol-Go]
func sumOfPower(nums []int, k int) int {
	const mod = 1_000_000_007
	f := make([]int, k+1)
	f[0] = 1
	for _, x := range nums {
		for j := k; j >= 0; j-- {
			if j >= x {
				f[j] = (f[j]*2 + f[j-x]) % mod
			} else {
				f[j] = f[j] * 2 % mod
			}
		}
	}
	return f[k]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。

## 题单：0-1 背包

- [2915. 和为目标值的最长子序列的长度](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/) 1659
- [416. 分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/)
- [494. 目标和](https://leetcode.cn/problems/target-sum/)
- [2787. 将一个数字表示成幂的和的方案数](https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/) 1818
- [474. 一和零](https://leetcode.cn/problems/ones-and-zeroes/)（二维）
- [1049. 最后一块石头的重量 II](https://leetcode.cn/problems/last-stone-weight-ii/) 2092
- [879. 盈利计划](https://leetcode.cn/problems/profitable-schemes/) 2204
- [956. 最高的广告牌](https://leetcode.cn/problems/tallest-billboard/) 2381
- [2518. 好分区的数目](https://leetcode.cn/problems/number-of-great-partitions/) 2415
- [2742. 给墙壁刷油漆](https://leetcode.cn/problems/painting-the-walls/) 2425
- [2291. 最大股票收益](https://leetcode.cn/problems/maximum-profit-from-trading-stocks/)（会员题）
- [2431. 最大限度地提高购买水果的口味](https://leetcode.cn/problems/maximize-total-tastiness-of-purchased-fruits/)（会员题）
