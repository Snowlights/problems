#### 题目

<p>给你三个整数&nbsp;<code>a</code>&nbsp;，<code>b</code>&nbsp;和&nbsp;<code>n</code>&nbsp;，请你返回&nbsp;<code>(a XOR x) * (b XOR x)</code>&nbsp;的&nbsp;<strong>最大值</strong>&nbsp;且 <code>x</code>&nbsp;需要满足 <code>0 &lt;= x &lt; 2<sup>n</sup></code>。</p>

<p>由于答案可能会很大，返回它对&nbsp;<code>10<sup>9 </sup>+ 7</code>&nbsp;<strong>取余</strong>&nbsp;后的结果。</p>

<p><strong>注意</strong>，<code>XOR</code>&nbsp;是按位异或操作。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>a = 12, b = 5, n = 4
<b>输出：</b>98
<b>解释：</b>当 x = 2 时，(a XOR x) = 14 且 (b XOR x) = 7 。所以，(a XOR x) * (b XOR x) = 98 。
98 是所有满足 0 &lt;= x &lt; 2<sup>n </sup>中 (a XOR x) * (b XOR x) 的最大值。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>a = 6, b = 7 , n = 5
<b>输出：</b>930
<b>解释：</b>当 x = 25 时，(a XOR x) = 31 且 (b XOR x) = 30 。所以，(a XOR x) * (b XOR x) = 930 。
930 是所有满足 0 &lt;= x &lt; 2<sup>n </sup>中 (a XOR x) * (b XOR x) 的最大值。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>a = 1, b = 6, n = 3
<b>输出：</b>12
<b>解释： </b>当 x = 5 时，(a XOR x) = 4 且 (b XOR x) = 3 。所以，(a XOR x) * (b XOR x) = 12 。
12 是所有满足 0 &lt;= x &lt; 2<sup>n </sup>中 (a XOR x) * (b XOR x) 的最大值。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= a, b &lt; 2<sup>50</sup></code></li>
	<li><code>0 &lt;= n &lt;= 50</code></li>
</ul>

#### 思路

先来看 $a$ 和 $b$ 均小于 $2^n$ 的情况。

举例说明：

```py
a = 11101
b = 00001
```

其中 $a$ 和 $b$ 同为 $0$ 的比特位，$x$ 可以取 $1$，这样异或后都是 $1$。同为 $1$ 的比特位，$x$ 可以取 $0$，这样异或后也都是 $1$。

所以重点讨论 $a$ 和 $b$ 的同一个比特位一个是 $1$ 另一个是 $0$ 的情况。

可以发现，无论 $x$ 取 $0$ 还是取 $1$，总有一个是 $1$ 另一个是 $0$，换句话说，$1$ 的个数是不变的。我把这样的比特位叫做「可分配」的位。

设 $\textit{ax} = a \oplus x,\ \textit{bx} = b \oplus x$，其中 $\oplus$ 表示异或运算。

如果 $x=00010$，则有

```py
ax = 11111
bx = 00011
```

乘积为 $31\cdot 3 = 93$，不够大。如何分配这些 $1$，使得 $\textit{ax}\cdot \textit{bx}$ 尽量大呢？

注意 $\textit{ax}$ 和 $\textit{bx}$ 的低两位一定是 $1$。其余位无论 $x$ 怎么取，一定满足总有一个是 $1$ 另一个是 $0$。也就是说，$\textit{ax}+\textit{bx}$ 是一个**定值**！

根据**基本不等式（均值定理）**，在和为定值的前提下，要让乘积尽量大，应当让 $\textit{ax}$ 与 $\textit{bx}$ 尽量接近。

所以，最优的分配方式是把剩下「可分配」的位中的最高位分给 $\textit{ax}$，其余位分给 $\textit{bx}$，即

```py
ax = 10011
bx = 01111
```

此时乘积最大，为 $19\cdot 15 = 285$。

#### 情况二

然后来看 $a$ 和 $b$ 至少有一个 $\ge 2^n$ 的情况。

不妨设 $a\ge b$。

对于高于或等于 $n$ 的比特位，我们是无法修改的，这会对「分配」产生什么影响呢？

分类讨论：

- 如果高于或等于 $n$ 的比特位，$a$ 和 $b$ 是一样的，那么问题转换成情况一。
- 否则高于或等于 $n$ 的比特位，满足 $a>b$，由于这些比特位无法修改，所以无论怎么分配，永远有 $\textit{ax} > \textit{bx}$，所以对于可修改的部分（低于 $n$ 的比特位），应当把「可分配」的位**全部**分给 $\textit{bx}$，让 $\textit{bx}$ 尽量大，从而使 $\textit{ax}$ 和 $\textit{bx}$ 尽量接近，得到最大的 $\textit{ax}\cdot \textit{bx}$。

```go  
func maximumXorProduct(A, B int64, n int) int {
	const mod = 1_000_000_007
	a, b := int(A), int(B)
	if a < b {
		a, b = b, a
	}

	mask := 1<<n - 1
	ax := a &^ mask // 第 n 位及其左边，无法被 x 影响，先算出来
	bx := b &^ mask
	a &= mask
	b &= mask // 低于第 n 位，能被 x 影响

	left := a ^ b      // a XOR x 和 b XOR x 一个是 0 一个是 1
	one := mask ^ left // a XOR x 和 b XOR x 均为 1
	ax |= one          // 先加到结果中
	bx |= one

	// 现在要把 left 分配到 ax 和 bx 中
	// 根据基本不等式（均值定理），分配后应当使 ax 和 bx 尽量接近，乘积才能尽量大
	if left > 0 && ax == bx { // a &^ mask = b &^ mask
		// 尽量均匀分配，例如把 1111 分成 1000 和 0111
		highBit := 1 << (bits.Len(uint(left)) - 1)
		ax |= highBit
		left ^= highBit
	}
	//（注意最上面保证了 a>=b）如果 a &^ mask 更大，则应当全部分给 bx
	bx |= left
	return ax % mod * (bx % mod) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
