### 题目

<p>给你三个整数 <code>start</code> ，<code>finish</code> 和 <code>limit</code> 。同时给你一个下标从 <strong>0</strong> 开始的字符串 <code>s</code> ，表示一个 <strong>正</strong> 整数。</p>

<p>如果一个 <strong>正</strong> 整数 <code>x</code> 末尾部分是 <code>s</code> （换句话说，<code>s</code> 是 <code>x</code> 的 <strong>后缀</strong>），且 <code>x</code> 中的每个数位至多是 <code>limit</code> ，那么我们称 <code>x</code> 是 <strong>强大的</strong> 。</p>

<p>请你返回区间 <code>[start..finish]</code> 内强大整数的 <strong>总数目</strong> 。</p>

<p>如果一个字符串 <code>x</code> 是 <code>y</code> 中某个下标开始（<strong>包括</strong> <code>0</code> ），到下标为 <code>y.length - 1</code> 结束的子字符串，那么我们称 <code>x</code> 是 <code>y</code> 的一个后缀。比方说，<code>25</code> 是 <code>5125</code> 的一个后缀，但不是 <code>512</code> 的后缀。</p>

<p> </p>

1. <p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>start = 1, finish = 6000, limit = 4, s = "124"
<b>输出：</b>5
<b>解释：</b>区间 [1..6000] 内的强大数字为 124 ，1124 ，2124 ，3124 和 4124 。这些整数的各个数位都 <= 4 且 "124" 是它们的后缀。注意 5124 不是强大整数，因为第一个数位 5 大于 4 。
这个区间内总共只有这 5 个强大整数。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>start = 15, finish = 215, limit = 6, s = "10"
<b>输出：</b>2
<b>解释：</b>区间 [15..215] 内的强大整数为 110 和 210 。这些整数的各个数位都 <= 6 且 "10" 是它们的后缀。
这个区间总共只有这 2 个强大整数。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>start = 1000, finish = 2000, limit = 4, s = "3000"
<b>输出：</b>0
<b>解释：</b>区间 [1000..2000] 内的整数都小于 3000 ，所以 "3000" 不可能是这个区间内任何整数的后缀。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= start <= finish <= 10<sup>15</sup></code></li>
	<li><code>1 <= limit <= 9</code></li>
	<li><code>1 <= s.length <= floor(log<sub>10</sub>(finish)) + 1</code></li>
	<li><code>s</code> 数位中每个数字都小于等于 <code>limit</code> 。</li>
	<li><code>s</code> 不包含任何前导 0 。</li>
</ul>

### 思路

定义 $\textit{dfs}(i,\textit{limitLow},\textit{limitHigh})$ 表示构造第 $i$ 位及其之后数位的合法方案数， 其余参数的含义为：

- $\textit{limitHigh}$ 表示当前是否受到了 $\textit{finish}$ 的约束（我们要构造的数字不能超过 $\textit{finish}$）。若为真，则第 $i$ 位填入的数字至多为 $\textit{finish}[i]$，否则至多为 $9$，这个数记作 $\textit{hi}$。如果在受到约束的情况下填了 $\textit{finish}[i]$，那么后续填入的数字仍会受到 $\textit{finish}$ 的约束。例如 $\textit{finish}=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{limitLow}$ 表示当前是否受到了 $\textit{start}$ 的约束（我们要构造的数字不能低于 $\textit{start}$）。若为真，则第 $i$ 位填入的数字至少为 $\textit{start}[i]$，否则至少为 $0$，这个数记作 $\textit{lo}$。如果在受到约束的情况下填了 $\textit{start}[i]$，那么后续填入的数字仍会受到 $\textit{start}$ 的约束。

枚举第 $i$ 位填什么数字。
如果 $i< n - |s|$，那么可以填 $[\textit{lo}, \min(\textit{hi}, \textit{limit})]$ 内的数，否则只能填 $s[i-(n-|s|)]$。这里 $|s|$ 表示 $s$ 的长度。
递归终点：$\textit{dfs}(n,*,*)=1$，表示成功构造出一个合法数字。
递归入口：$\textit{dfs}(0, \texttt{true}, \texttt{true})$，表示：

- 从最高位开始枚举。
- 一开始要受到 $\textit{start}$ 和 $\textit{finish}$ 的约束（否则就可以随意填了，这肯定不行）。

### 答疑

**问**：记忆化三个状态有点麻烦，能不能只记忆化 $i$ 这个状态？
**答**：是可以的。比如 $\textit{finish}=234$，第一位填 $2$，第二位填 $3$，后面无论怎么递归，都不会再次递归到第一位填 $2$，第二位填 $3$ 的情况，所以不需要记录。对于 $\textit{start}$ 也同理。

根据这个例子，我们可以只记录不受到 $\textit{limitLow}$ 或 $\textit{limitHigh}$ 约束时的状态 $i$。相当于记忆化的是 $(i,\texttt{false},\texttt{false})$ 这个状态，因为其它状态只会递归访问一次。

```go [sol]
func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
	low, high := strconv.FormatInt(start, 10), strconv.FormatInt(finish, 10)
	low = strings.Repeat("0", len(high)-len(low)) + low
	memo, n := make([]int, len(high)), len(high)
	diff := n - len(s)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int, bool, bool) int
	dfs = func(i int, lowLimit, highLimit bool) int {
		if i == n {
			return 1
		}

		var res int
		if !lowLimit && !highLimit {
			if memo[i] != -1 {
				return memo[i]
			}
			defer func() {
				memo[i] = res
			}()
		}

		l, h := 0, 9
		if lowLimit {
			l = int(low[i] - '0')
		}
		if highLimit {
			h = int(high[i] - '0')
		}

		if i < diff {
			for d := l; d <= min(h, limit); d++ {
				res += dfs(i+1, lowLimit && d == l, highLimit && d == h)
			}
		} else {
			x := int(s[i-diff] - '0')
			if l <= x && x <= min(h, limit) {
				res += dfs(i+1, lowLimit && x == l, highLimit && x == h)
			}
		}
		return res
	}

	return int64(dfs(0, true, true))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD)$，其中 $n=\mathcal{O}(\log \textit{finish})$，$D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以动态规划的时间复杂度为 $\mathcal{O}(nD)$。
- 空间复杂度：$\mathcal{O}(n)$。即状态个数。

## 题单：数位 DP

- [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/)（[题解](https://leetcode.cn/problems/number-of-digit-one/solution/by-endlesscheng-h9ua/)）
- [面试题 17.06. 2 出现的次数](https://leetcode.cn/problems/number-of-2s-in-range-lcci/)（[题解](https://leetcode.cn/problems/number-of-2s-in-range-lcci/solution/by-endlesscheng-x4mf/)）
- [357. 统计各位数字都不同的数字个数](https://leetcode.cn/problems/count-numbers-with-unique-digits/)
- [600. 不含连续 1 的非负整数](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/)（[题解](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/solution/by-endlesscheng-1egu/)）
- [788. 旋转数字](https://leetcode.cn/problems/rotated-digits/)
- [902. 最大为 N 的数字组合](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/) 1990
- [2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers/) 2120
- [1012. 至少有 1 位重复的数字](https://leetcode.cn/problems/numbers-with-repeated-digits/)（[题解](https://leetcode.cn/problems/numbers-with-repeated-digits/solution/by-endlesscheng-c5vg/)）2230
- [2827. 范围中美丽整数的数目](https://leetcode.cn/problems/number-of-beautiful-integers-in-the-range/) 2324
- [2719. 统计整数数目](https://leetcode.cn/problems/count-of-integers/) 2355
- [2801. 统计范围内的步进数字数目](https://leetcode.cn/problems/count-stepping-numbers-in-range/) 2367
- [1397. 找到所有好字符串](https://leetcode.cn/problems/find-all-good-strings/) 2667
- [1215. 步进数](https://leetcode.cn/problems/stepping-numbers/)（会员题）1675
- [248. 中心对称数 III](https://leetcode.cn/problems/strobogrammatic-number-iii/)（会员题）
- [1067. 范围内的数字计数](https://leetcode.cn/problems/digit-count-in-range/)（会员题）2025
- [1088. 易混淆数 II](https://leetcode.cn/problems/confusing-number-ii/)（会员题）2077
