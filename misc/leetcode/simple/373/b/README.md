#### 题目

<p>给你一个字符串 <code>s</code> 和一个正整数 <code>k</code> 。</p>

<p>用 <code>vowels</code> 和 <code>consonants</code> 分别表示字符串中元音字母和辅音字母的数量。</p>

<p>如果某个字符串满足以下条件，则称其为 <strong>美丽字符串</strong> ：</p>

<ul>
	<li><code>vowels == consonants</code>，即元音字母和辅音字母的数量相等。</li>
	<li><code>(vowels * consonants) % k == 0</code>，即元音字母和辅音字母的数量的乘积能被 <code>k</code> 整除。</li>
</ul>

<p>返回字符串 <code>s</code> 中 <strong>非空美丽子字符串</strong> 的数量。</p>

<p>子字符串是字符串中的一个连续字符序列。</p>

<p>英语中的<strong> 元音字母 </strong>为 <code>'a'</code>、<code>'e'</code>、<code>'i'</code>、<code>'o'</code> 和 <code>'u'</code> 。</p>

<p>英语中的<strong> 辅音字母 </strong>为除了元音字母之外的所有字母。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "baeyh", k = 2
<strong>输出：</strong>2
<strong>解释：</strong>字符串 s 中有 2 个美丽子字符串。
- 子字符串 "b<em><strong>aeyh</strong></em>"，vowels = 2（["a","e"]），consonants = 2（["y","h"]）。
可以看出字符串 "aeyh" 是美丽字符串，因为 vowels == consonants 且 vowels * consonants % k == 0 。
- 子字符串 "<em><strong>baey</strong></em>h"，vowels = 2（["a","e"]），consonants = 2（["b","y"]）。
可以看出字符串 "baey" 是美丽字符串，因为 vowels == consonants 且 vowels * consonants % k == 0 。
可以证明字符串 s 中只有 2 个美丽子字符串。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "abba", k = 1
<strong>输出：</strong>3
<strong>解释：</strong>字符串 s 中有 3 个美丽子字符串。
- 子字符串 "<strong><em>ab</em></strong>ba"，vowels = 1（["a"]），consonants = 1（["b"]）。
- 子字符串 "ab<strong><em>ba</em></strong>"，vowels = 1（["a"]），consonants = 1（["b"]）。
- 子字符串 "<em><strong>abba</strong></em>"，vowels = 2（["a","a"]），consonants = 2（["b","b"]）。
可以证明字符串 s 中只有 3 个美丽子字符串。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "bcdf", k = 1
<strong>输出：</strong>0
<strong>解释：</strong>字符串 s 中没有美丽子字符串。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>1 &lt;= k &lt;= 1000</code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

把元音视作 $1$，辅音视作 $-1$。

「元音字母和辅音字母的数量相等」就等价于：找到一个和为 $0$ 的连续子数组。注意这种子数组的长度一定是偶数，因为元音辅音数量相等。

设子数组的长度为 $L$。由于元音辅音数量相等，所以元音辅音数量都等于 $\dfrac{L}{2}$，所以「元音字母和辅音字母的数量的乘积能被 $k$ 整除」等价于

$$
\left(\dfrac{L}{2}\right)^2 \bmod k = 0
$$

这等价于

$$
L^2 \bmod (4k) = 0
$$

这个平方很烦人，如果能去掉平方就好做了。

## 来点数学

我们来研究下，如果一个数 $L$ 的平方能被 $n$ 整除，意味着什么？

假设 $n$ 是一个质数，例如 $3$，那么 $L$ 必须包含质因子 $3$，此时题目约束就变成了：$L$ 是 $3$ 的倍数。我们把平方去掉了！

如果 $n$ 是一个质数 $p$ 的 $e$ 次幂呢？分类讨论：

- 如果 $e$ 是偶数，比如 $n=3^4$，那么 $L$ 必须包含因子 $3^2$，才能使得 $L^2$ 能被 $n$ 整除。此时题目约束就变成了：$L$ 是 $3^2$ 的倍数。
- 如果 $e$ 是奇数，比如 $n=3^5$，那么 $L$ 必须包含因子 $3^3$，才能使得 $L^2$ 能被 $n$ 整除。此时题目约束就变成了：$L$ 是 $3^3$ 的倍数。

所以 $L$ 必须包含因子 $p^r$，其中 $r=\left\lceil\dfrac{e}{2}\right\rceil = \left\lfloor\dfrac{e+1}{2}\right\rfloor$。

如果 $n$ 可以分解出多个质因子，只需要把每个质因子及其幂次按照上面的方法处理，把结果相乘，就得到 $L$ 必须是什么数的倍数了。

这样就把平方去掉了。

## 套路：前缀和+哈希表

把 $4k$ 按照上述方法计算，设 $L$ 必须是 $k'$ 的倍数。

现在问题变成，有多少个和为 $0$ 的子数组，其长度是 $k'$ 的倍数？

设子数组的下标范围为 $[j,i)$，那么其长度 $L=i-j$，则有

$$
(i-j)\bmod k' = 0
$$

等价于

$$
i \bmod k' = j\bmod k'
$$

对于前缀和来说（定义见最上面贴的题解），子数组和为 $0$ 相当于 $s[i]-s[j] = 0$，即

$$
s[i] = s[j]
$$

我们需要**同时满足**这两个条件（下标模 $k'$ 相等，$s$ 值相等），这可以一并用哈希表解决。

哈希表的 key 是一个 pair：$(i\bmod k', s[i])$，哈希表的 value 是这个 pair 的出现次数。

代码实现时，前缀和数组可以用一个变量表示。

代码实现时，可以把 aeiou 压缩成一个二进制数，从而快速判断字母是否为元音，原理请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

#### 答疑

**问**：为什么哈希表要在一开始插入一个 $(k'-1, 0)$？

**答**：前缀和的第一项是 $0$，由于代码中是从下标 $0$ 开始算第二个前缀和，所以相当于 $s[-1] = 0$，而 $k'-1$ 和 $-1$ 关于 $k'$ 同余，所以插入 $(k'-1, 0)$。


```go  
func beautifulSubstrings(s string, k int) (ans int) {
	m := 1
	k *= 4
	for p := 2; p*p <= k; p++ {
		if k%p > 0 {
			continue
		}
		e := 1
		for k /= p; k%p == 0; k /= p {
			e++
		}
		m *= pow(p, (e+1)/2)
	}
	if k > 1 {
		m *= k
	}

	k = m

	sum := make([]int, len(s)+1)
	for i, v := range s {
		sum[i+1] = sum[i]
		if strings.Contains("aeiou", string(v)) {
			sum[i+1]++
		} else {
			sum[i+1]--
		}
	}

	pos := make([]map[int]int, len(sum)*2+10)
	for i := range pos {
		pos[i] = map[int]int{}
	}

	for i, s := range sum {
		s += len(sum)
		r := pos[s][i%k]
		ans += r
		pos[s][i%k]++
	}
	return
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x
		}
		x = x * x
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + \sqrt k)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
