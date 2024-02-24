### 题目

<p>给你一个整数 <code>n</code> 。</p>

<p>如果一个字符串 <code>s</code> 只包含小写英文字母，<strong>且</strong> 将 <code>s</code> 的字符重新排列后，新字符串包含 <strong>子字符串</strong> <code>"leet"</code> ，那么我们称字符串 <code>s</code> 是一个 <strong>好</strong> 字符串。</p>

<p>比方说：</p>

<ul>
	<li>字符串 <code>"lteer"</code> 是好字符串，因为重新排列后可以得到 <code>"leetr"</code> 。</li>
	<li><code>"letl"</code> 不是好字符串，因为无法重新排列并得到子字符串 <code>"leet"</code> 。</li>
</ul>

<p>请你返回长度为 <code>n</code> 的好字符串 <strong>总</strong> 数目。</p>

<p>由于答案可能很大，将答案对<strong> </strong><code>10<sup>9</sup> + 7</code> <strong>取余</strong> 后返回。</p>

<p><strong>子字符串</strong> 是一个字符串中一段连续的字符序列。</p>

<div class="notranslate" style="all: initial;"> </div>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>n = 4
<b>输出：</b>12
<b>解释：</b>总共有 12 个字符串重新排列后包含子字符串 "leet" ："eelt" ，"eetl" ，"elet" ，"elte" ，"etel" ，"etle" ，"leet" ，"lete" ，"ltee" ，"teel" ，"tele" 和 "tlee" 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>n = 10
<b>输出：</b>83943898
<b>解释：</b>长度为 10 的字符串重新排列后包含子字符串 "leet" 的方案数为 526083947580 。所以答案为 526083947580 % (10<sup>9</sup> + 7) = 83943898 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
</ul>

### 思路

正难则反。总共有 $26^n$ 个字符串，减去不含 `leet` 的字符串个数，就得到了答案。

不含 `leet` 的字符串，需要至少满足如下三个条件中的一个：

1. 不含 `l`。
2. 不含 `t`。
3. 不含 `e` 或者恰好包含一个 `e`。

分类讨论。

#### 至少满足一个条件

1. 不含 `l`：每个位置可以填 $25$ 种字母，方案数为 $25^n$。
2. 不含 `t`：同上，方案数为 $25^n$。
3. 不含 `e` 或者恰好包含一个 `e`：不含 `e` 同上，方案数为 $25^n$；恰好包含一个 `e`，先从 $n$ 个位置中选一个填 `e`，然后剩下 $n-1$ 个位置不能包含 `e`，方案数为 $n\cdot 25^{n-1}$。加起来就是 $25^n + n\cdot 25^{n-1}$。

直接加起来，就是 $(3\cdot 25+n)\cdot 25^{n-1}$，但这样就重复统计了「至少满足两个条件」的情况，要减去。

#### 至少满足两个条件

1. 不含 `l` 和 `t`：每个位置可以填 $24$ 种字母，方案数为 $24^n$。
2. 不含 `l` 且 `e` 的个数不足两个：同「满足一个条件」中 3 的分析，额外不能填 `l`，方案数为 $24^n + n\cdot 24^{n-1}$。
3. 不含 `t` 且 `e` 的个数不足两个：同上，方案数为 $24^n + n\cdot 24^{n-1}$。

直接加起来，就是 $(3\cdot 24+2n)\cdot 24^{n-1}$，但这样就重复统计了「满足三个条件」的情况，要减去。

#### 满足三个条件

同「满足一个条件」中 3 的分析，额外不能填 `l` 和 `t`，方案数为 $23^n + n\cdot 23^{n-1}$。

#### 总结

不含 `leet` 的字符串的个数为「至少满足一个条件」减去「至少满足两个条件」加上「满足三个条件」，这就是**容斥原理**。

最后用 $26^n$ 减去不含 `leet` 的字符串的个数，得到答案：

$$
26^n - (3\cdot 25+n)\cdot 25^{n-1} + (3\cdot 24+2n)\cdot 24^{n-1} - (23+n)\cdot 23^{n-1}

$$

```go
func stringCount(n int) int {
	mod := int(1e9 + 7)
	pow := func(x, n int) int {
		res := 1
		for ; n > 0; n /= 2 {
			if n%2 > 0 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	return ((pow(26, n)-
		pow(25, n-1)*(75+n)+
		pow(24, n-1)*(72+n*2)-
		pow(23, n-1)*(23+n))%mod + mod) % mod
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(1)$。
