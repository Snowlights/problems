#### 题目

<p>给你三个字符串 <code>s1</code>、<code>s2</code> 和 <code>s3</code>。 你可以根据需要对这三个字符串执行以下操作 <strong>任意次数</strong> <!-- notionvc: b5178de7-3318-4129-b7d9-726b47e90621 -->。</p>

<p>在每次操作中，你可以选择其中一个长度至少为 <code>2</code> 的字符串 <!-- notionvc: 3342ac46-33c8-4010-aacd-e58678ce31ef --> 并删除其 <strong>最右位置上</strong> 的字符。</p>

<p>如果存在某种方法能够使这三个字符串相等，请返回使它们相等所需的 <strong>最小</strong> 操作次数；否则，返回 <code>-1</code>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s1 = "abc"，s2 = "abb"，s3 = "ab"
<strong>输出：</strong>2
<strong>解释：</strong>对 s1 和 s2 进行一次操作后，可以得到三个相等的字符串。
可以证明，不可能用少于两次操作使它们相等。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s1 = "dac"，s2 = "bac"，s3 = "cac"
<strong>输出：</strong>-1
<strong>解释：</strong>因为 s1 和 s2 的最左位置上的字母<!-- notionvc: 47239f7c-eec1-49f8-af79-c206ec88cb07 -->不相等，所以无论进行多少次操作，它们都不可能相等。因此答案是 -1 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s1.length, s2.length, s3.length &lt;= 100</code></li>
	<li><code>s1</code>、<code>s2</code> 和 <code>s3</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

设 $\textit{lcp}$ 为三个字符串的最长公共前缀的长度。

如果 $\textit{lcp}=0$，无法操作成一样的，返回 $-1$。

否则返回三个字符串的长度之和，减去剩下的长度 $3\cdot \textit{lcp}$。

```go  
func findMinimumOperations(s1, s2, s3 string) int {
	n := min(len(s1), len(s2), len(s3))
	i := 0
	for i < n && s2[i] == s1[i] && s3[i] == s1[i] {
		i++
	}
	if i == 0 {
		return -1
	}
	return len(s1) + len(s2) + len(s3) - i*3
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为三个字符串中的最短字符串的长度。
- 空间复杂度：$\mathcal{O}(1)$。

