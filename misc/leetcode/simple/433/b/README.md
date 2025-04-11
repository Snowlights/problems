#### 题目

<p>给你一个整数数组 <code>nums</code> 和一个正整数 <code>k</code>，返回所有长度最多为 <code>k</code> 的 <strong>子序列</strong> 中&nbsp;<strong>最大值&nbsp;</strong>与&nbsp;<strong>最小值&nbsp;</strong>之和的总和。</p>

<p><strong>非空子序列</strong>&nbsp;是指从另一个数组中删除一些或不删除任何元素（且不改变剩余元素的顺序）得到的数组。</p>

<p>由于答案可能非常大，请返回对 <code>10<sup>9</sup> + 7</code> 取余数的结果。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3], k = 2</span></p>

<p><strong>输出：</strong> 24</p>

<p><strong>解释：</strong></p>

<p>数组 <code>nums</code> 中所有长度最多为 2 的子序列如下：</p>

<table style="border: 1px solid black; border-collapse: collapse;">
	<thead>
		<tr>
			<th style="border: 1px solid black;">子序列</th>
			<th style="border: 1px solid black;">最小值</th>
			<th style="border: 1px solid black;">最大值</th>
			<th style="border: 1px solid black;">和</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td style="border: 1px solid black;"><code>[1]</code></td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">2</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;"><code>[2]</code></td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">4</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;"><code>[3]</code></td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">6</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;"><code>[1, 2]</code></td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">3</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;"><code>[1, 3]</code></td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">4</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;"><code>[2, 3]</code></td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">5</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;"><strong>总和</strong></td>
			<td style="border: 1px solid black;">&nbsp;</td>
			<td style="border: 1px solid black;">&nbsp;</td>
			<td style="border: 1px solid black;">24</td>
		</tr>
	</tbody>
</table>

<p>因此，输出为 24。</p>
</div>

<p><strong>示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [5,0,6], k = 1</span></p>

<p><strong>输出：</strong> 22</p>

<p><strong>解释：</strong></p>

<p>对于长度恰好为 1 的子序列，最小值和最大值均为元素本身。因此，总和为 <code>5 + 5 + 0 + 0 + 6 + 6 = 22</code>。</p>
</div>

<p><strong>示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,1,1], k = 2</span></p>

<p><strong>输出：</strong> 12</p>

<p><strong>解释：</strong></p>

<p>子序列 <code>[1, 1]</code> 和 <code>[1]</code> 各出现 3 次。对于所有这些子序列，最小值和最大值均为 1。因此，总和为 12。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= k &lt;= min(100, nums.length)</code></li>
</ul>

#### 思路

由于我们只关心子序列中的最值，所以元素顺序不影响答案，可以先把数组排序。

排序后，假设 $\textit{nums}[i]$ 是子序列的最大值，这样的子序列有多少个？

我们可以从下标 $[0,i-1]$ 中选至多 $\min(k-1,i)$ 个数，作为子序列的其他元素。这样的选法总共有

$$
\sum_{j=0}^{\min(k-1,i)} \binom i j
$$

个。

由于 $k$ 很小，直接暴力计算上式。

于是 $\textit{nums}[i]$ 对答案的贡献为

$$
\textit{nums}[i] \cdot \sum_{j=0}^{\min(k-1,i)} \binom i j
$$

同理可以枚举 $\textit{nums}[i]$ 作为子序列的最小值，从右边选数字，做法同上。

**技巧**：根据对称性，$\textit{nums}[n-1-i]$ 作为最小值时，子序列个数和 $\textit{nums}[i]$ 作为最大值的个数是一样的，所以二者可以一同计算。

## 代码实现

1. 关于组合数，我们需要预处理阶乘及其逆元，然后利用公式 $C(n,m) = \dfrac{n!}{m!(n-m)!}$ 计算。
2. 关于逆元的知识点，见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)，包含费马小定理的数学证明。

```
const mod = 1_000_000_007
const mx = 100_000

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
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

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func minMaxSums(nums []int, k int) (ans int) {
	slices.Sort(nums)
	for i, x := range nums {
		s := 0
		for j := range min(k, i+1) {
			s += comb(i, j)
		}
		ans = (ans + s%mod*(x+nums[len(nums)-1-i])) % mod
	}
	return
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n\log n + nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 优化

对于和式

$$
s_i = \sum_{j=0}^{\min(k-1,i)} \binom i j
$$

考虑怎么通过 $s_i$ 递推得到 $s_{i+1}$：

- 如果 $i< k-1$，那么我们相当于在 $s_i$，也就是包含至多 $i$ 个数的子序列的基础上，多了一个数。这个数可以选也可以不选，所以有
  $$
  s_{i+1} = 2\cdot s_i
  $$
- 如果 $i \ge k-1$，那么我们相当于在 $s_i$，也就是包含至多 $k-1$ 个数的子序列的基础上，多了一个数。这个数可以选也可以不选，如果选，那么前面的子序列不能包含恰好 $k-1$ 个数，要减掉，也就是减去从 $i$ 个数中选出 $k-1$ 个数的方案数，所以有
  $$
  s_{i+1} = 2\cdot s_i - \binom {i} {k-1}
  $$

这样可以 $\mathcal{O}(1)$ 递推和式，不需要 $\mathcal{O}(k)$ 暴力去算。

初始值 $s_0 = 1$。

代码实现时，$s_i$ 可以优化成一个变量。

```
const mod = 1_000_000_007
const mx = 100_000

var f [mx]int    // f[i] = i!
var invF [mx]int // invF[i] = i!^-1

func init() {
	f[0] = 1
	for i := 1; i < mx; i++ {
		f[i] = f[i-1] * i % mod
	}

	invF[mx-1] = pow(f[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invF[i-1] = invF[i] * i % mod
	}
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

func comb(n, m int) int {
	if m > n {
		return 0
	}
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func minMaxSums(nums []int, k int) (ans int) {
	slices.Sort(nums)
	s := 1
	for i, x := range nums {
		ans = (ans + s*(x+nums[len(nums)-1-i])) % mod
		s = (s*2 - comb(i, k-1) + mod) % mod
	}
	return
}
```

#### 复杂度分析

忽略预处理的时间和空间。

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。


## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)