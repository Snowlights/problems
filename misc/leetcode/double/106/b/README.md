### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的字符串 <code>s</code> ，这个字符串只包含 <code>0</code> 到 <code>9</code> 的数字字符。</p>

<p>如果一个字符串 <code>t</code> 中至多有一对相邻字符是相等的，那么称这个字符串是 <strong>半重复的</strong> 。</p>

<p>请你返回 <code>s</code> 中最长 <strong>半重复</strong> 子字符串的长度。</p>

<p>一个 <strong>子字符串</strong> 是一个字符串中一段连续 <strong>非空</strong> 的字符。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>s = &#34;52233&#34;
<b>输出：</b>4
<b>解释：</b>最长半重复子字符串是 &#34;5223&#34; ，子字符串从 i = 0 开始，在 j = 3 处结束。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>s = &#34;5494&#34;
<b>输出：</b>4
<b>解释：</b>s 就是一个半重复字符串，所以答案为 4 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>s = &#34;1111111&#34;
<b>输出：</b>2
<b>解释：</b>最长半重复子字符串是 &#34;11&#34; ，子字符串从 i = 0 开始，在 j = 1 处结束。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 50</code></li>
	<li><code>&#39;0&#39; &lt;= s[i] &lt;= &#39;9&#39;</code></li>
</ul>
 
### 思路  

移动右指针 $\textit{right}$，并统计相邻相同的情况出现了多少次，记作 $\textit{cnt}$

如果 $\textit{cnt}>1$，则不断移动左指针 $\textit{left}$ 直到 $s[\textit{left}]=s[\textit{left}-1]$，此时将一对相同的字符移到窗口之外。然后将 $\textit{same}$ 置为 $1$。

然后统计子串长度 $\textit{right}-\textit{left}+1$ 的最大值。


```go 

func longestSemiRepetitiveSubstring(s string) (ans int) {
	l, cnt := 0, 0
	ans = 1
	for r := 1; r < len(s); r++ {
		if s[r] == s[r-1] {
			cnt++
			if cnt > 1 {
				l++
				for s[l] != s[l-1] {
					l++
				}
				cnt--
			}
		}
		ans = max(ans, r-l+1)
	}

	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。注意 $\textit{left}$ 只会增加不会减少，所以二重循环的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。