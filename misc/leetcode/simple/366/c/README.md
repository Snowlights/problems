#### 题目

<p>给你两个下标从 <strong>0</strong> 开始的二进制字符串 <code>s1</code> 和 <code>s2</code> ，两个字符串的长度都是 <code>n</code> ，再给你一个正整数 <code>x</code> 。</p>

<p>你可以对字符串 <code>s1</code> 执行以下操作 <strong>任意次</strong> ：</p>

<ul>
	<li>选择两个下标 <code>i</code> 和 <code>j</code> ，将 <code>s1[i]</code> 和 <code>s1[j]</code> 都反转，操作的代价为 <code>x</code> 。</li>
	<li>选择满足 <code>i < n - 1</code> 的下标 <code>i</code> ，反转 <code>s1[i]</code> 和 <code>s1[i + 1]</code> ，操作的代价为 <code>1</code> 。</li>
</ul>

<p>请你返回使字符串 <code>s1</code> 和 <code>s2</code> 相等的 <strong>最小</strong> 操作代价之和，如果无法让二者相等，返回 <code>-1</code> 。</p>

<p><strong>注意</strong> ，反转字符的意思是将 <code>0</code> 变成 <code>1</code> ，或者 <code>1</code> 变成 <code>0</code> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><b>输入：</b>s1 = "1100011000", s2 = "0101001010", x = 2
<b>输出：</b>4
<b>解释：</b>我们可以执行以下操作：
- 选择 i = 3 执行第二个操作。结果字符串是 s1 = "110<em><strong>11</strong></em>11000" 。
- 选择 i = 4 执行第二个操作。结果字符串是 s1 = "1101<em><strong>00</strong></em>1000" 。
- 选择 i = 0 和 j = 8 ，执行第一个操作。结果字符串是 s1 = "<em><strong>0</strong></em>1010010<em><strong>1</strong></em>0" = s2 。
总代价是 1 + 1 + 2 = 4 。这是最小代价和。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>s1 = "10110", s2 = "00011", x = 4
<b>输出：</b>-1
<b>解释：</b>无法使两个字符串相等。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == s1.length == s2.length</code></li>
	<li><code>1 <= n, x <= 500</code></li>
	<li><code>s1</code> 和 <code>s2</code> 只包含字符 <code>'0'</code> 和 <code>'1'</code> 。</li>
</ul>

#### 思路

## 方法一：O(n^2) DP

#### 什么时候返回 -1？

注意到，反转两个字符：

- 如果都是 $0$ 或者都是 $1$，那么反转后 $1$ 的个数会加 $2$ 或者减 $2$；
- 如果一个是 $0$ 另一个是 $1$，那么反转后 $1$ 的个数不变。

所以，无论是哪种操作，都不会改变 $s_1$ 中的 $1$ 的个数的奇偶性。
那么只要 $s_1$ 和 $s_2$ 的 $1$ 的个数一个是奇数一个是偶数，就直接返回 $-1$。否则，哪怕只用第二种操作，都一定可以让 $s_1=s_2$。

#### 然后呢？

先想 DP，再想贪心。毕竟 DP 就是对暴力搜索的优化。（或者看数据范围比较小，也可以往 DP 上想。）
考虑 $s_1$ 和 $s_2$ 的最后一对字符（也可以考虑第一对字符）：

- 如果相同，那么无需修改。
- 如果不同：
  - 选择第一种操作，相当于后面可以「免费」反转一个字符。
  - 选择第二种操作，那么下一个字符要把 $0$ 看作 $1$，把 $1$ 看作 $0$。

所以除了知道当前下标 $i$，还需要知道免费反转次数 $j$，以及上一个字符是否选择了第二种操作 $\textit{preRev}$。
定义 $\textit{dfs}(i,j,\textit{preRev})$，参数含义如上，返回值是在这种状态下的最小操作代价之和。
分类讨论：

- 如果 `(s1[i] == s2[i]) == (not pre_rev)`，表示 $s_1[i]$ 和 $s_2[i]$ 是相等的，无需操作，返回 $\textit{dfs}(i - 1, j, \text{false})$。
- 否则：
  - 选择第一种操作：$\textit{dfs}(i - 1, j + 1, \text{false}) + x$。
  - 选择第二种操作：$\textit{dfs}(i - 1, j, \text{true}) + 1$。
  - 如果 $j>0$，免费反转一次：$\textit{dfs}(i - 1, j-1, \text{false})$。
  - 这三种情况取最小值。

递归边界，当 $i<0$ 时：

- 如果 $j>0$ 或者 $\textit{preRev}$ 为真，那么不合法，返回 $\infty$。
- 否则返回 $0$。

递归入口：$\textit{dfs}(n-1,0,\text{false})$，即答案。

```go
func minOperations(s1, s2 string, x int) int {
	if strings.Count(s1, "1")%2 != strings.Count(s2, "1")%2 {
		return -1
	}
	n := len(s1)
	memo := make([][][2]int, n)
	for i := range memo {
		memo[i] = make([][2]int, n+1)
		for j := range memo[i] {
			memo[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, preRev int) int {
		if i < 0 {
			if j > 0 || preRev > 0 {
				return 1e9
			}
			return 0
		}
		p := &memo[i][j][preRev]
		if *p != -1 {
			return *p
		}
		if s1[i] == s2[i] == (preRev == 0) { // 无需反转
			return dfs(i-1, j, 0)
		}
		res := min(dfs(i-1, j+1, 0)+x, dfs(i-1, j, 1)+1)
		if j > 0 { // 可以免费反转
			res = min(res, dfs(i-1, j-1, 0))
		}
		*p = res
		return res
	}
	return dfs(n-1, 0, 0)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s_1$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。
- 空间复杂度：$\mathcal{O}(n^2)$。

## 方法二：O(n) DP

把 $s_1[i]\ne s_2[i]$ 的下标统计到数组 $p$ 中。如果 $p$ 的长度是奇数，那么和上面一样，无法操作，返回 $-1$。
设 $f[i]$ 表示修改 $p$ 的前 $i$ 个位置的最小花费，那么对于 $p[i]$，有两种方案：
第一种操作：花费 $x$，那么对于 $p[i]$ 相当于花费 $\dfrac{x}{2}$，因此

$$
f[i] = f[i-1] + \dfrac{x}{2}

$$

注意这个转移一定会发生偶数次，因为 $p$ 的长度是偶数，并且第二种操作每次反转两个数，所以第一种操作一定会反转偶数个下标。
第二种操作：需要不断用相邻的位置操作，把 $p[i]$ 和 $p[i-1]$ 都反转，那么需要操作 $p[i]-p[i-1]$ 次，因此

$$
f[i] = f[i-2] + p[i]-p[i-1]

$$

两者取最小值，即

$$
f[i] = \min(f[i-1] + \dfrac{x}{2}, f[i-2] + p[i]-p[i-1])

$$

代码实现时，为了方便处理 $\dfrac{x}{2}$，可以先在计算过程中把所有数都乘 $2$，最后返回答案的时候再除以 $2$。
初始值 $f[-1]=0$，$f[0]=x$，答案为 $f[m-1]$，这里 $m$ 是数组 $p$ 的长度。你可以把 $f$ 数组的下标都加一来避免负数，也可以用两个变量滚动计算。

```go
func minOperations(s1, s2 string, x int) int {
	if s1 == s2 {
		return 0
	}
	p := []int{}
	for i, c := range s1 {
		if byte(c) != s2[i] {
			p = append(p, i)
		}
	}
	if len(p)%2 > 0 {
		return -1
	}
	f0, f1 := 0, x
	for i := 1; i < len(p); i++ {
		f0, f1 = f1, min(f1+x, f0+(p[i]-p[i-1])*2)
	}
	return f1 / 2
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s_1$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
