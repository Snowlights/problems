#### 题目

<p>给你一个整数 <code>k</code> 和一个整数 <code>x</code> 。</p>

<p>令 <code>s</code> 为整数 <code>num</code> 的下标从 <strong>1</strong> 开始的二进制表示。我们说一个整数 <code>num</code> 的 <strong>价值</strong> 是满足 <code>i % x == 0</code> 且 <code><font face="monospace">s[i]</font></code> 是 <strong>设置位</strong> 的 <code>i</code> 的数目。</p>

<p>请你返回<strong> 最大</strong> 整数<em> </em><code>num</code> ，满足从 <code>1</code> 到 <code>num</code> 的所有整数的 <strong>价值</strong> 和小于等于 <code>k</code> 。</p>

<p><b>注意：</b></p>

<ul>
	<li>一个整数二进制表示下 <strong>设置位</strong> 是值为 <code>1</code> 的数位。</li>
	<li>一个整数的二进制表示下标从右到左编号，比方说如果 <code>s == 11100</code> ，那么 <code>s[4] == 1</code> 且 <code>s[2] == 0</code> 。</li>
</ul>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>k = 9, x = 1
<b>输出：</b>6
<b>解释：</b>数字 1 ，2 ，3 ，4 ，5 和 6 二进制表示分别为 "1" ，"10" ，"11" ，"100" ，"101" 和 "110" 。
由于 x 等于 1 ，每个数字的价值分别为所有设置位的数目。
这些数字的所有设置位数目总数是 9 ，所以前 6 个数字的价值和为 9 。
所以答案为 6 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>k = 7, x = 2
<b>输出：</b>9
<b>解释：</b>由于 x 等于 2 ，我们检查每个数字的偶数位。
2 和 3 在二进制表示下的第二个数位为设置位，所以它们的价值和为 2 。
6 和 7 在二进制表示下的第二个数位为设置位，所以它们的价值和为 2 。
8 和 9 在二进制表示下的第四个数位为设置位但第二个数位不是设置位，所以它们的价值和为 2 。
数字 1 ，4 和 5 在二进制下偶数位都不是设置位，所以它们的价值和为 0 。
10 在二进制表示下的第二个数位和第四个数位都是设置位，所以它的价值为 2 。
前 9 个数字的价值和为 6 。
前 10 个数字的价值和为 8，超过了 k = 7 ，所以答案为 9 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= k <= 10<sup>15</sup></code></li>
	<li><code>1 <= x <= 8</code></li>
</ul>

#### 思路

## 方法一：二分答案 + 数位 DP

由于 $\textit{num}$ 越大，价值和也越大，有单调性，可以二分答案。
于是问题转换成：计算从 $1$ 到 $\textit{num}$ 的价值和，判断其是否 $\le k$。
思路和 [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/) 是一样的，
设 $\textit{num}$ 的二进制长度为 $n$，，额外在 $d=1$ 时判断当前下标是否为 $x$ 的倍数，如果是， 把 $\textit{cnt}_1$ 加一。

最后还剩一个问题：二分的上界取多少合适？
$1$ 到 $\textit{num}$ 中的数字 $v$，除以 $2^{x-1}$ 后如果是奇数，就说明 $v$ 至少包含一个我们需要的 $1$。而 $1$ 到 $\textit{num}$ 中
有 $\left\lfloor\dfrac{\textit{num}+1}{2}\right\rfloor$ 个奇数，所以取 $\textit{num}$ 为

$$
(k+1) \cdot 2^x

$$

可以保证价值和至少是 $k+1$，一定不满足二分判定。

```go [sol]
func findMaximumNumber(k int64, x int) int64 {
	return int64(sort.Search(int(1e18), func(i int) bool {
		i++
		return countDigitOne(i, x) > int(k)
	}))
}

func countDigitOne(n, x int) int {
	s := strconv.FormatInt(int64(n), 2)
	n = len(s)

	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	btoi := func(b bool) int{
		if b {
			return 1
		}
		return 0
	}
	var dfs func(i, pre int, limit bool) int
	dfs = func(i, pre int, limit bool) (res int) {
		if i == n {
			return pre
		}

		if !limit {
			dv := &memo[i][pre]
			if *dv != -1 {
				return *dv
			}
			defer func() {
				*dv = res
			}()
		}

		low, high := 0, 1
		if limit {
			high = int(s[i] - '0')
		}

		for ; low <= high; low++ {
			res += dfs(i+1, pre + btoi(low == 1 && (n-i) % x == 0), limit && low == high)
		}
		return res
	}untDigitOne(long num) {\n        this.num = num;\n        int m = 64 - Long.numberOfLeadingZeros(num);\n        memo = new long[m][m + 1];\n        for (long[] row : memo) {\n            Arrays.fill(row, -1);\n
	return dfs(0, 0, true)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((x + \log k)^3)$。$\textit{num}$ 的二进制长度为 $\mathcal{O}(x + \log k)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}((x + \log k)^2)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}((x + \log k)^2)$。再算上 $\mathcal{O}(x + \log k)$ 的二分次数，总的时间复杂度为 $\mathcal{O}((x + \log k)^3)$。
- 空间复杂度：$\mathcal{O}((x + \log k)^2)$。

## 方法二：二分答案+数学公式

举例说明，$1$ 到 $\textit{num}=6$ 的二进制表示如下：

$$
\begin{aligned}
&001\\
&010\\
&011\\
&100\\
&101\\
&110
\end{aligned}

$$

考虑计算 $1$ 到 $\textit{num}$ 的每个比特位上有多少个 $1$。
最低位：相当于计算 $1$ 到 $\textit{num}$ 中有多少个奇数，这有 $\left\lfloor\dfrac{\textit{num}+1}{2}\right\rfloor = 3$ 个。
次低位：我们可以把每个数都除以 $2$（右移 $1$ 位），忽略 $00$，得到：

$$
\begin{aligned}
&01\\
&01\\
&10\\
&10\\
&11
\end{aligned}

$$

同样地，考虑其中的奇数，发现 $1$ 到 $\left\lfloor\dfrac{\textit{num}}{2}\right\rfloor-1$ 中的每个奇数都出现了 $2$ 次，
但是 $\left\lfloor\dfrac{\textit{num}}{2}\right\rfloor$ 只出现了 $1$ 次。所以次低位中有 $3$ 个 $1$。
从右往左第三位：我们可以把每个数都除以 $2^2=4$（右移 $2$ 位）。同样地，考虑其中的奇数，发现 $1$ 到 $\left\lfloor\dfrac{\textit{num}}{4}\right\rfloor-1$ 中的每个奇数都出现了 $4$ 次，
但是 $\left\lfloor\dfrac{\textit{num}}{4}\right\rfloor=1$ 出现了 $3$ 次。
这个 $3$ 怎么算？我们可以把 $\textit{num}=6$ 和 $2^2-1 = 3$ 取 AND，得到 $2$，
说明把 $\left\lfloor\dfrac{\textit{num}}{4}\right\rfloor$ 作为前缀的二进制数，
最低位和次低位是 $00,01,10$，这有 $2+1=3$ 个。一般地，从低到高第 $i$ 位（$i$ 从 $0$ 开始）的 $1$ 的个数，我们分两部分计算：

1. 设 $n = \left\lfloor\dfrac{\textit{num}}{2^i}\right\rfloor$，从 $1$ 到 $n-1$ 有 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个奇数，每个奇数作为 $1$ 到 $\textit{num}$ 的二进制数的前缀，出现了 $2^i$ 次。所以一共有 $\left\lfloor\dfrac{n}{2}\right\rfloor\cdot 2^i$ 个 $1$。
2. $n$ 单独作为 $1$ 到 $\textit{num}$ 的二进制数的前缀，出现了 $ (\textit{num} \& (2^i-1))+1 $ 次。如果 $n$ 是奇数，那么它为第 $i$ 位贡献了 $({num} \& (2^i-1))+1$ 个 $1$，否则为第 $i$ 位贡献了 $0$ 个 $1$。

最后，按照题目要求，只统计 $x$ 的倍数位置上的 $1$ 的个数。

```go [sol]
func findMaximumNumber(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		num++
		res := 0
		i := x - 1
		for n := num >> i; n > 0; n >>= x {
			res += n / 2 << i
			if n%2 > 0 {
				mask := 1<<i - 1
				res += num&mask + 1
			}
			i += x
		}
		return res > int(k)
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}\left(\dfrac{\log^2 k}{x}\right)$。$\textit{num}$ 的二进制长度为 $m=\mathcal{O}(x + \log k)$，我们统计了其中的 $\mathcal{O}(m/x)$ 个比特位，所以每次二分需要 $\mathcal{O}\left(\dfrac{\log k}{x}\right)$ 的时间。再算上 $\mathcal{O}(x + \log k)$ 的二分次数，总的时间复杂度为 $\mathcal{O}\left(\dfrac{\log^2 k}{x}\right)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法三：逐位构造

为方便计算，考察 $1$ 到 $\textit{num}-1$ 的价值和。从高到低构建 $\textit{num}$ 的每个比特位。设当前枚举到 $\textit{num}$ 的从低到高的第 $i$ 个比特位（$i$ 从 $0$ 开始），设 $i$ 左边有 $\textit{pre}_1$ 个编号是 $x$ 的倍数且填了 $1$ 的比特位。如果 $\textit{num}$ 的第 $i$ 个比特位填 $1$，那么价值和会增加多少呢？例如 $i=2$，那么：

- 我们新增了 $4$ 个小于 $\textit{num}$ 的二进制数，它们在 $i=2$ 这个比特位上是 $0$，并且分别以 $00,01,10,11$ 结尾，并且这 $4$ 个二进制数在 $i$ 左边都有 $\textit{pre}_1$ 个编号是 $x$ 的倍数且填了 $1$ 的比特位。这一部分产生了 $4\cdot \textit{pre}_1$ 的价值和。
- $i$ 右边有 $\left\lfloor\dfrac{i}{x}\right\rfloor$ 个编号是 $x$ 的倍数的比特位，这些比特位一共有 $\left\lfloor\dfrac{i}{x}\right\rfloor\cdot 2^{i-1}$ 个 $1$。例如 $x=1$ 时，$00,01,10,11$ 中一共有 $4$ 个 $1$。

两者相加，在 $\textit{num}$ 的从低到高的第 $i$ 个比特位上填 $1$，会让 $1$ 到 $\textit{num}-1$ 的价值和增加

$$
\textit{cnt} = \textit{pre}_1\cdot 2^i + \left\lfloor\dfrac{i}{x}\right\rfloor\cdot 2^{i-1}

$$

如果 $\textit{cnt}\le k$，那么这个比特位可以填 $1$。由于我们是从高到低考虑的，能填 $1$ 就填 $1$，这会让答案尽量大。然后把 $k$ 减少 $\textit{cnt}$。
最后还剩一个问题：$i$ 从哪个比特位开始枚举？\n\n由方法一可知，$1$ 到 $(k+1) \cdot 2^x$ 的价值和至少是 $k+1$，所以我们至多考虑 $1$ 到 $(k+1) \cdot 2^x - 1$ 的价值和，
所以 $i$ 应当初始化为 $(k+1) \cdot 2^x$ 的最高位，即 $(k+1) \cdot 2^x$ 的二进制长度减一。

```go [sol]
func findMaximumNumber(K int64, x int) int64 {
	k := int(K)
	num, pre1 := 0, 0
	for i := bits.Len(uint((k+1)<<x)) - 1; i >= 0; i-- {
		cnt := pre1<<i + i/x<<i>>1
		if cnt > k {
			continue
		}
		k -= cnt
		num |= 1 << i
		if (i+1)%x == 0 {
			pre1++
		}
	}
	return int64(num - 1)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(x+\log k)$。
- 空间复杂度：$\mathcal{O}(1)$。
