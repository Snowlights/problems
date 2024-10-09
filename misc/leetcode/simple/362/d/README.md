#### 题目

<p>给你两个长度都为 <code>n</code>&nbsp;的字符串&nbsp;<code>s</code> 和&nbsp;<code>t</code>&nbsp;。你可以对字符串 <code>s</code>&nbsp;执行以下操作：</p>
<ul>
<li>将 <code>s</code>&nbsp;长度为 <code>l</code>&nbsp;（<code>0 &lt; l &lt; n</code>）的 <strong>后缀字符串</strong>&nbsp;删除，并将它添加在 <code>s</code>&nbsp;的开头。<br/>
比方说，<code>s = 'abcd'</code>&nbsp;，那么一次操作中，你可以删除后缀&nbsp;<code>'cd'</code>&nbsp;，并将它添加到&nbsp;<code>s</code>&nbsp;的开头，得到&nbsp;<code>s = 'cdab'</code>&nbsp;。</li>
</ul>
<p>给你一个整数&nbsp;<code>k</code>&nbsp;，请你返回&nbsp;<strong>恰好</strong> <code>k</code>&nbsp;次操作将<em>&nbsp;</em><code>s</code> 变为<em>&nbsp;</em><code>t</code>&nbsp;的方案数。</p>
<p>由于答案可能很大，返回答案对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong>&nbsp;后的结果。</p>
<p>&nbsp;</p>
<p><strong class="example">示例 1：</strong></p>
<pre>
<b>输入：</b>s = "abcd", t = "cdab", k = 2
<b>输出：</b>2
<b>解释：</b>
第一种方案：
第一次操作，选择 index = 3 开始的后缀，得到 s = "dabc" 。
第二次操作，选择 index = 3 开始的后缀，得到 s = "cdab" 。

第二种方案：
第一次操作，选择 index = 1 开始的后缀，得到 s = "bcda" 。
第二次操作，选择 index = 1 开始的后缀，得到 s = "cdab" 。
</pre>
<p><strong class="example">示例 2：</strong></p>
<pre>
<b>输入：</b>s = "ababab", t = "ababab", k = 1
<b>输出：</b>2
<b>解释：</b>
第一种方案：
选择 index = 2 开始的后缀，得到 s = "ababab" 。

第二种方案：
选择 index = 4 开始的后缀，得到 s = "ababab" 。
</pre>
<p>&nbsp;</p>
<p><strong>提示：</strong></p>
<ul>
<li><code>2 &lt;= s.length &lt;= 5 * 10<sup>5</sup></code></li>
<li><code>1 &lt;= k &lt;= 10<sup>15</sup></code></li>
<li><code>s.length == t.length</code></li>
<li><code>s</code> 和&nbsp;<code>t</code>&nbsp;都只包含小写英文字母。</li>
</ul>

#### 思路

根据题意，我们可以把 $s$ 看成是一个首尾相连的字符串（循环字符串），只要 $s+s$ 中包含 $t$，那么可以从 $s$ 变成 $t$。  
计算 $t$ 在循环字符串 $s$ 中的出现次数，记作 $c$，这可以用 KMP 等字符串匹配算法解决，即寻找 $t$ 在 $s+s$ 中的出现次数，注意当 $s=t$ 时要少统计一次。  
定义 $f[i][0]$ 表示 $i$ 次操作后等于 $t$ 的方案数，$f[i][1]$ 表示 $i$ 次操作后不等于 $t$ 的方案数。  
初始值：
- 如果 $s=t$，那么 $f[0][0]=1,f[0][1]=0$。
- 如果 $s\ne t$，那么 $f[0][0]=0,f[0][1]=1$。
  
分类讨论，如果操作后等于 $t$：
- 如果上一步也是 $t$，我们有 $c-1$ 种操作方案；
- 如果上一步不是 $t$，我们有 $c$ 种操作方案。
  
如果操作后不等于 $t$：
- 如果上一步是 $t$，我们有 $n-c$ 种操作方案；
- 如果上一步不是 $t$，我们有 $n-1-c$ 种操作方案。
  
所以有

$$
\begin{aligned}\nf[i][0] &= f[i-1][0]\cdot (c-1) + f[i-1][1]\cdot c\\\nf[i][1] &= f[i-1][0]\cdot (n-c) + f[i-1][1]\cdot (n-1-c)\n\end{aligned}
$$

上式可以改写成如下矩阵形式

$$
\begin{bmatrix}\nf[i][0] \\\nf[i][1] \\\n\end{bmatrix}\n=\n\begin{bmatrix}\nc-1 & c \\\nn-c & n-1-c \\\n\end{bmatrix}\n\cdot\n\begin{bmatrix}\nf[i-1][0] \\\nf[i-1][1] \\\n\end{bmatrix}
$$

进而得到

$$
\begin{bmatrix}\nf[k][0] \\\nf[k][1] \\\n\end{bmatrix}\n=\n\begin{bmatrix}\nc-1 & c \\\nn-c & n-1-c \\\n\end{bmatrix} ^ k\n\cdot\n\begin{bmatrix}\nf[0][0]\\\nf[0][1] \\\n\end{bmatrix}
$$

利用**矩阵快速幂**（参考 [70. 爬楼梯的官方题解的方法二](https://leetcode.cn/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode-solution/)），可以得到 $f[k][0]$，即为本题答案。

```go 
const mod = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, m)
	}
	return a
}

func newIdentityMatrix(n int) matrix {
	a := make(matrix, n)
	for i := range a {
		a[i] = make([]int, n)
		a[i][i] = 1
	}
	return a
}

func (a matrix) mul(b matrix) matrix {
	c := newMatrix(len(a), len(b[0]))
	for i, row := range a {
		for j := range b[0] {
			for k, v := range row {
				c[i][j] = (c[i][j] + v*b[k][j]) % mod
			}
			if c[i][j] < 0 {
				c[i][j] += mod
			}
		}
	}
	return c
}

func (a matrix) pow(n int64) matrix {
	res := newIdentityMatrix(len(a))
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res.mul(a)
		}
		a = a.mul(a)
	}
	return res
}

func numberOfWays(s, t string, k int64) int {
	n := len(s)
	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	kmpSearch := func(text, pattern string) (cnt int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				if i-lenP+1 < n {
					cnt++
				}
				c = match[c-1]
			}
		}
		return
	}
	// 上面都是模板
	// 下面是本题代码
	c := kmpSearch(s+s, t)
	m := matrix{
		{c - 1, c},
		{n - c, n - 1 - c},
	}.pow(k)
	if s == t {
		return m[0][0]
	}
	return m[0][1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+\log k)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。