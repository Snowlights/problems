### 题目

<p>给你一个下标从 <strong>0</strong>&nbsp;开始的字符串&nbsp;<code>word</code>&nbsp;。</p>

<p>一次操作中，你可以选择&nbsp;<code>word</code>&nbsp;中任意一个下标 <code>i</code>&nbsp;，将&nbsp;<code>word[i]</code> 修改成任意一个小写英文字母。</p>

<p>请你返回消除 <code>word</code>&nbsp;中所有相邻 <strong>近似相等</strong>&nbsp;字符的 <strong>最少</strong>&nbsp;操作次数。</p>

<p>两个字符&nbsp;<code>a</code> 和&nbsp;<code>b</code>&nbsp;如果满足&nbsp;<code>a == b</code>&nbsp;或者&nbsp;<code>a</code> 和&nbsp;<code>b</code>&nbsp;在字母表中是相邻的，那么我们称它们是 <strong>近似相等</strong>&nbsp;字符。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>word = "aaaaa"
<b>输出：</b>2
<b>解释：</b>我们将 word 变为 "a<em><strong>c</strong></em>a<em><strong>c</strong></em>a" ，该字符串没有相邻近似相等字符。
消除 word 中所有相邻近似相等字符最少需要 2 次操作。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>word = "abddez"
<b>输出：</b>2
<b>解释：</b>我们将 word 变为 "<em><strong>y</strong></em>bd<em><strong>o</strong></em>ez" ，该字符串没有相邻近似相等字符。
消除 word 中所有相邻近似相等字符最少需要 2 次操作。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>word = "zyxyxyz"
<b>输出：</b>3
<b>解释：</b>我们将 word 变为 "z<em><strong>a</strong></em>x<em><strong>a</strong></em>x<em><strong>a</strong></em>z" ，该字符串没有相邻近似相等字符。
消除 word 中所有相邻近似相等字符最少需要 3 次操作
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 100</code></li>
	<li><code>word</code>&nbsp;只包含小写英文字母。</li>
</ul>

### 思路

下文把 $\textit{word}$ 简记为 $s$。
从左到右遍历 $s$，如果发现 $s[i-1]$ 和 $s[i]$ 近似相等，应当改 $s[i-1]$ 还是 $s[i]$？  
如果改 $s[i-1]$，那么 $s[i]$ 和 $s[i+1]$ 是可能近似相等的，但如果改 $s[i]$，就可以避免 $s[i]$ 和 $s[i+1]$ 近似相等。  
所以每次发现两个相邻字母近似相等，就改右边那个。

```go  
func removeAlmostEqualCharacters(word string) int {
	i, n, ans := 1, len(word), 0
	for i < n {
		if abs(int(word[i])-int(word[i-1])) <= 1 {
			i += 2
			ans++
		} else {
			i++
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
