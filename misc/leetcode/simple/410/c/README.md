#### 题目

<p>给你一个长度为&nbsp;<code>n</code>&nbsp;的&nbsp;<strong>正</strong>&nbsp;整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>如果两个&nbsp;<strong>非负</strong>&nbsp;整数数组&nbsp;<code>(arr1, arr2)</code>&nbsp;满足以下条件，我们称它们是&nbsp;<strong>单调</strong>&nbsp;数组对：</p>

<ul>
	<li>两个数组的长度都是&nbsp;<code>n</code>&nbsp;。</li>
	<li><code>arr1</code>&nbsp;是单调<strong>&nbsp;非递减</strong>&nbsp;的，换句话说&nbsp;<code>arr1[0] &lt;= arr1[1] &lt;= ... &lt;= arr1[n - 1]</code>&nbsp;。</li>
	<li><code>arr2</code>&nbsp;是单调 <strong>非递增</strong>&nbsp;的，换句话说&nbsp;<code>arr2[0] &gt;= arr2[1] &gt;= ... &gt;= arr2[n - 1]</code>&nbsp;。</li>
	<li>对于所有的&nbsp;<code>0 &lt;= i &lt;= n - 1</code>&nbsp;都有&nbsp;<code>arr1[i] + arr2[i] == nums[i]</code>&nbsp;。</li>
</ul>

<p>请你返回所有 <strong>单调</strong>&nbsp;数组对的数目。</p>

<p>由于答案可能很大，请你将它对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,3,2]</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p>单调数组对包括：</p>

<ol>
	<li><code>([0, 1, 1], [2, 2, 1])</code></li>
	<li><code>([0, 1, 2], [2, 2, 0])</code></li>
	<li><code>([0, 2, 2], [2, 1, 0])</code></li>
	<li><code>([1, 2, 2], [1, 1, 0])</code></li>
</ol>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [5,5,5,5]</span></p>

<p><span class="example-io"><b>输出：</b>126</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length &lt;= 2000</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 50</code></li>
</ul>

#### 思路

## 方法一：前缀和优化 DP

看示例 1，$\textit{nums}=[2,3,2]$。

如果 $\textit{arr}_1[2]=2$，那么 $\textit{arr}_2[2]=\textit{nums}[2] - \textit{arr}_1[2]=2-2= 0$。考虑枚举 $\textit{arr}_1[1]$ 是多少：

- 如果 $\textit{arr}_1[1]=0$，那么问题变成计算下标 $0$ 到 $1$ 中的单调数组对的个数，且 $\textit{arr}_1[1]=0$。
- 如果 $\textit{arr}_1[1]=1$，那么问题变成计算下标 $0$ 到 $1$ 中的单调数组对的个数，且 $\textit{arr}_1[1]=1$。
- 如果 $\textit{arr}_1[1]=2$，那么问题变成计算下标 $0$ 到 $1$ 中的单调数组对的个数，且 $\textit{arr}_1[1]=2$。
- 注意 $\textit{arr}_1[1]$ 不能超过 $\textit{arr}_1[2]$，且 $\textit{arr}_2[1] = \textit{nums}[1] - \textit{arr}_1[1]$ 不能低于 $\textit{arr}_2[2]$。

累加这些方案数，我们就得到了下标 $0$ 到 $2$ 中的单调数组对的个数，且 $\textit{arr}_1[2]=2$。

上面的讨论说明本题的**子问题**是「下标 $0$ 到 $i$ 中的单调数组对的个数，且 $\textit{arr}_1[i]=j$」，将其记作 $f[i][j]$。

考虑枚举 $\textit{arr}_1[i-1] = k$，那么必须满足 $\textit{arr}_1[i-1]\le \textit{arr}_1[i]$ 且 $\textit{arr}_2[i-1]\ge \textit{arr}_2[i]$，即 $k\le j$ 且 $\textit{nums}[i-1]-k\ge \textit{nums}[i] - j$。

整理得

$$
k \le \min(j,\textit{nums}[i-1] - \textit{nums}[i] + j) = j + \min(\textit{nums}[i-1] - \textit{nums}[i], 0)
$$

⚠**注意**：无需判断 $k\le \textit{nums}[i-1]$，因为 $\textit{nums}[i-1] - \textit{nums}[i] + j = \textit{nums}[i-1] - (\textit{nums}[i] - j)\le \textit{nums}[i-1]$，由数学归纳法可以证明 $k\le \textit{nums}[i-1]$。

记 $\textit{maxK} = j + \min(\textit{nums}[i-1] - \textit{nums}[i], 0)$，那么有

$$
f[i][j] =
\begin{cases}
0, & \textit{maxK} < 0     \\
\sum\limits_{k=0}^{\textit{maxK}} f[i-1][k], &  \textit{maxK} \ge 0     \\
\end{cases}
$$

其中和式可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化。

设 $f[i-1]$ 的前缀和 $s[j] = \sum\limits_{k=0}^{j} f[i-1][k]$，状态转移方程化简为

$$
f[i][j] =
\begin{cases}
0, & \textit{maxK} < 0     \\
s[\textit{maxK}], &  \textit{maxK} \ge 0     \\
\end{cases}
$$

初始值：$f[0][j] = 1$，其中 $j=0,1,2,\cdots,\textit{nums}[0]$。

答案：枚举 $\textit{arr}_1[n-1] = 0,1,2,\cdots,\textit{nums}[n-1]$，累加得 $\sum\limits_{j=0}^{\textit{nums}[n-1]} f[n-1][j]$。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1Cf421v7Ky/)，欢迎点赞关注！

```
func countOfPairs(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Max(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	s := make([]int, m+1)

	for j := 0; j <= nums[0]; j++ {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		s[0] = f[i-1][0]
		for k := 1; k <= m; k++ {
			s[k] = s[k-1] + f[i-1][k] // f[i-1] 的前缀和
		}
		for j := 0; j <= nums[i]; j++ {
			maxK := j + min(nums[i-1]-nums[i], 0)
			if maxK >= 0 {
				f[i][j] = s[maxK] % mod
			}
		}
	}

	for _, v := range f[n-1][:nums[n-1]+1] {
		ans += v
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(nm)$。

### 空间优化

进一步分析，$\textit{maxK} \ge 0$ 即

$$
j + \min(\textit{nums}[i-1] - \textit{nums}[i], 0) \ge 0
$$

变形得

$$
j\ge \max(\textit{nums}[i] - \textit{nums}[i-1], 0)
$$

记 $j_0 = \max(\textit{nums}[i] - \textit{nums}[i-1], 0)$，那么我们只需要计算在 $j$ 区间 $[j_0, \textit{nums}[i]]$ 中的 $f[i][j]$，其余 $f[i][j]$ 均为 $0$。

设 $s[j] = \sum\limits_{k=0}^{j} f[i-1][k]$，状态转移方程化简为

$$
f[i][j] =
\begin{cases}
0, & j < j_0     \\
s[j-j_0], &  j\ge j_0    \\
\end{cases}
$$

代码实现时，$f[i][j]$ 可以用一维数组滚动计算：先把前缀和直接保存在 $f$ 数组中，然后倒序更新 $f[j] = f[j-j_0]$（倒序更新的理由同 0-1 背包）。

此外，由于在计算答案时只考虑 $j\le \textit{nums}[n-1]$ 的状态，所以 $f$ 数组只需要开 $\textit{nums}[n-1]+1$ 的大小。

```
func countOfPairs(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := nums[n-1]
	f := make([]int, m+1)
	for j := range f[:min(nums[0], m)+1] {
		f[j] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= m; j++ {
			f[j] += f[j-1] // 计算前缀和
		}
		j0 := max(nums[i]-nums[i-1], 0)
		for j := min(nums[i], m); j >= j0; j-- {
			f[j] = f[j-j0] % mod
		}
		clear(f[:min(j0, m+1)])
	}
	for _, v := range f {
		ans += v
	}
	return ans % mod
}
```

### 进一步优化

如果 $j_0 > \min(\textit{nums}[i],m)$，那么后面计算出的 $f[j]$ 均为 $0$，我们可以直接返回 $0$。

此外，前缀和只需要计算到 $j=\min(\textit{nums}[i],m) - j_0$ 为止。

```
func countOfPairs(nums []int) (ans int) {
	const mod = 1_000_000_007
	n := len(nums)
	m := nums[n-1]
	f := make([]int, m+1)
	for j := range f[:min(nums[0], m)+1] {
		f[j] = 1
	}
	for i := 1; i < n; i++ {
		j0 := max(nums[i]-nums[i-1], 0)
		m2 := min(nums[i], m)
		if j0 > m2 {
			return 0
		}
		for j := 1; j <= m2-j0; j++ {
			f[j] = (f[j] + f[j-1]) % mod // 计算前缀和
		}
		copy(f[j0:m2+1], f)
		clear(f[:j0])
	}
	for _, v := range f {
		ans += v
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $\textit{nums}$ 的长度，$m=\textit{nums}[n-1]$。
- 空间复杂度：$\mathcal{O}(m)$。

## 方法二：组合数学

首先来看一个特殊的情况：所有 $\textit{nums}[i]$ 都相同。

例如示例 2，$\textit{nums}=[5,5,5,5]$，在所有元素都相同的情况下，只要 $\textit{arr}_1$ 是单调非递减的，那么 $\textit{arr}_2$ 就一定是单调非递增的。

问题变成：

- 计算长为 $n=4$ 的单调非递减数组个数，数组元素范围是 $[0,5]$。

![w410d-C.png](https://pic.leetcode.cn/1723373547-JXaMKp-w410d-C.png)

考虑一条从左下角走到右上角的路径，每次只能向右或向上走。向右走时，把之前向上走的次数（或者说当前高度）作为数组元素值。如上图，对应的数组为 $[2,3,3,4]$。

由于路径与单调非递减数组是一一对应的，所以问题变成路径个数是多少。

要向上走 $5$ 步，向右走 $4$ 步，一共走 $5+4=9$ 步。选择其中 $4$ 步向右走，于是问题变成从 $9$ 步中选 $4$ 步的方案数，即

$$
C(9,4) = 126
$$

以 $m=\textit{nums}[n-1]$ 为基准，如果所有元素都等于 $m$，那么问题等价于从 $m+n$ 步中选 $n$ 步的方案数，即

$$
C(m+n,n)
$$

回到原问题，来看看 $\textit{nums}[i]$ 会如何影响路径个数。

为方便描述，下文把 $\textit{arr}_1$ 记作 $a$。

如果 $a[i-1] = x,\ a[i] = y$，那么必须满足 $x\le y$ 且 $\textit{nums}[i-1]-x\ge \textit{nums}[i]-y$，即

$$
y \ge \max(x, x+ \textit{nums}[i]-\textit{nums}[i-1])
$$

分类讨论：

- 如果 $\textit{nums}[i]-\textit{nums}[i-1]\le 0$，那么上式相当于 $y\ge x$。由于我们要计算的就是单调非递减的数组个数，所以这不会影响上面得出的 $C(m+n,m)$ 的结论。
- 如果 $\textit{nums}[i]-\textit{nums}[i-1]> 0$，那么上式相当于 $y\ge x + \textit{nums}[i]-\textit{nums}[i-1]$。这意味着 $a[i]$ 填的数字必须比 $a[i-1]$ 填的数字多至少 $d=\textit{nums}[i]-\textit{nums}[i-1]$。用路径来理解，就是在 $i$ 这个位置向右走之前，要「**强制性**」地向上走 $d$ 步。

剩下的 $m+n-d$ 步可以随意安排向右走还是向上走。于是问题变成从 $m+n-d$ 步中选 $n$ 步向右走的方案数，即

$$
C(m+n-d,n)
$$

一般地，设 $d_i=\max(\textit{nums}[i]-\textit{nums}[i-1],0)$，计算

$$
m = \textit{nums}[n-1] - (d_1 + d_2 + \cdots + d_{n-1})
$$

那么答案为

$$
\begin{cases}
0, & m < 0     \\
C(m+n,n), & m \ge 0     \\
\end{cases}
$$

由于 $C(m+n,n) = C(m+n,m)$，答案也可以是 $C(m+n,m)$。

### 答疑

**问**：为什么要以 $\textit{nums}[n-1]$ 为基准？

**答**：从路径的角度上看，$\textit{nums}[n-1]$ 与 $n$ 一起，决定了我们总共要走多少步。至于 $\textit{nums}$ 中的其他元素，影响的是 $a[i]$ 与 $a[i-1]$ 的差值，相当于在某些位置强行向上走若干步，这不会改变「总共要走 $\textit{nums}[n-1]+n$ 步」的事实。此外，如果以 $\textit{arr}_2$，也就是「单调非递增的数组个数」计算答案，也可以以 $\textit{nums}[0]$ 为基准，计算方法是类似的，感兴趣的读者可以自行推导，从而加深对方法二的理解。

```
const mod = 1_000_000_007
const mx = 3001 // MAX_N + MAX_M = 2000 + 1000 = 3000

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

func comb(n, m int) int {
	return f[n] * invF[m] % mod * invF[n-m] % mod
}

func countOfPairs(nums []int) int {
	n := len(nums)
	m := nums[n-1]
	for i := 1; i < n; i++ {
		m -= max(nums[i]-nums[i-1], 0)
		if m < 0 {
			return 0
		}
	}
	return comb(m+n, n)
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

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。忽略预处理的时间和空间。
- 空间复杂度：$\mathcal{O}(1)$。


- 时间复杂度：$\mathcal{O}(n+q)$，其中 $q$ 是 $\textit{queries}$ 的长度。注意每个点只会被合并一次，在后面的循环中会被 `i = find(l)` 以及 `i = find(i + 1)` 跳过。由于数组的特殊性，每次合并的复杂度为 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

## 分类题单

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
