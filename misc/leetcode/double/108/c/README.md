### 题目  

<p>给你一个二进制字符串 <code>s</code> ，你需要将字符串分割成一个或者多个 <strong>子字符串</strong>  ，使每个子字符串都是 <strong>美丽</strong> 的。</p>

<p>如果一个字符串满足以下条件，我们称它是 <strong>美丽</strong> 的：</p>

<ul>
	<li>它不包含前导 0 。</li>
	<li>它是 <code>5</code> 的幂的 <strong>二进制</strong> 表示。</li>
</ul>

<p>请你返回分割后的子字符串的 <strong>最少</strong> 数目。如果无法将字符串 <code>s</code> 分割成美丽子字符串，请你返回 <code>-1</code> 。</p>

<p>子字符串是一个字符串中一段连续的字符序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>s = &#34;1011&#34;
<b>输出：</b>2
<b>解释：</b>我们可以将输入字符串分成 [&#34;101&#34;, &#34;1&#34;] 。
- 字符串 &#34;101&#34; 不包含前导 0 ，且它是整数 5<sup>1</sup> = 5 的二进制表示。
- 字符串 &#34;1&#34; 不包含前导 0 ，且它是整数 5<sup>0</sup> = 1 的二进制表示。
最少可以将 s 分成 2 个美丽子字符串。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>s = &#34;111&#34;
<b>输出：</b>3
<b>解释：</b>我们可以将输入字符串分成 [&#34;1&#34;, &#34;1&#34;, &#34;1&#34;] 。
- 字符串 &#34;1&#34; 不包含前导 0 ，且它是整数 5<sup>0</sup> = 1 的二进制表示。
最少可以将 s 分成 3 个美丽子字符串。
</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>s = &#34;0&#34;
<b>输出：</b>-1
<b>解释：</b>无法将给定字符串分成任何美丽子字符串。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 15</code></li>
	<li><code>s[i]</code> 要么是 <code>&#39;0&#39;</code> 要么是 <code>&#39;1&#39;</code> 。</li>
</ul>
 
### 思路  

由于字符串的长度为 $n$，所以二进制的值是低于 $2^n$ 的。那么可以预处理 $2^n$ 以内的 $5$ 的幂的二进制表示，记作 $\textit{pow}_5$。由于测试数据很多，可以用全局变量，预处理 $2^{15}$ 以内的 $5$ 的幂（有 $7$ 个）。  
定义 $\textit{dfs}(i)$ 表示划分从 $s[i]$ 开始的后缀，最少要划分成多少段。  
枚举 $\textit{pow}_5$ 中的字符串 $t$，设其长度为 $m$，如果从 $s[i]$ 到 $s[i+m-1]$ 与 $t$ 相等，那么有  

$$
\textit{dfs}(i) = \textit{dfs}(i+m) + 1
$$

所有情况取最小值。  
如果 $s[i]=\texttt{0}$，或者不存在这样的 $t$，那么 $\textit{dfs}(i)=\infty$。  
递归边界：$\textit{dfs}(n)=0$。  
递归入口：$\textit{dfs}(0)$。  
注意在比较字符串时，由于 $\textit{pow}_5$ 中字符串的公共前缀很短，很容易就失配了，所以除了匹配上的那次是 $\mathcal{O}(n)$ 的，其余匹配都可以视作是 $\mathcal{O}(1)$ 的。

```go 
package main

import "strconv"

var pow5 []string

func init() {
	for p5 := 1; p5 < 1<<15; p5 *= 5 {
		pow5 = append(pow5, strconv.FormatUint(uint64(p5), 2))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumBeautifulSubstrings(s string) (ans int) {

	n := len(s)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		f[i] = n + 1
		if s[i] == '0' {
			continue
		}
		for _, t := range pow5 {
			if i+len(t) > n {
				break
			}
			if s[i:i+len(t)] == t {
				f[i] = min(f[i], f[i+len(t)]+1)
			}
		}
	}
	if f[0] > n {
		return -1
	}
	return f[0]
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2)$。注意在比较字符串时，由于 $\textit{pow}_5$ 中字符串的公共前缀很短，很容易就失配了，所以除了匹配上的那次是 $\mathcal{O}(n)$ 的，其余匹配都可以视作是 $\mathcal{O}(1)$ 的。
- 空间复杂度：$\mathcal{O}(n)$。