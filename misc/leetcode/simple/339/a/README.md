#### 题目  

<p>给你一个仅由 <code>0</code> 和 <code>1</code> 组成的二进制字符串 <code>s</code> 。<span style=""> </span><span style=""> </span></p>

<p>如果子字符串中 <strong>所有的<span style=""> </span></strong><code><span style="">0</span></code><strong><span style=""> </span>都在 </strong><code>1</code><strong> 之前</strong> 且其中 <code>0</code> 的数量等于 <code>1</code> 的数量，则认为 <code>s</code> 的这个子字符串是平衡子字符串。请注意，空子字符串也视作平衡子字符串。<span style=""> </span></p>

<p>返回 <span style=""> </span><code>s</code> 中最长的平衡子字符串长度。</p>

<p>子字符串是字符串中的一个连续字符序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>s = &#34;01000111&#34;
<strong>输出：</strong>6
<strong>解释：</strong>最长的平衡子字符串是 &#34;000111&#34; ，长度为 6 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>s = &#34;00111&#34;
<strong>输出：</strong>4
<strong>解释：</strong>最长的平衡子字符串是 &#34;0011&#34; ，长度为 <span style=""> </span>4 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>s = &#34;111&#34;
<strong>输出：</strong>0
<strong>解释：</strong>除了空子字符串之外不存在其他平衡子字符串，所以答案为 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 50</code></li>
	<li><code>&#39;0&#39; &lt;= s[i] &lt;= &#39;1&#39;</code></li>
</ul>
 
#### 思路  

记录上一段连续相同字符个数 $\textit{pre}$，以及当前连续相同字符个数 $\textit{cur}$。  
如果当前字符是 $1$，那么上一段的字符是 $0$，这两段可以组成一个 $01$ 串，由于 $0$ 和 $1$ 的个数需要相等，那么当前这个 $01$ 串的最大长度就是 

$$
2\cdot \min(\textit{pre}, \textit{cur})
$$

取所有长度的最大值，即为答案。更新答案的时候，可以在当前字符是 $1$，且下一个字符是 $0$ 的位置上更新（或者 $1$ 在最末尾的时候）。

```go 
func findTheLongestBalancedSubstring(s string) (ans int) {
	pre, cur := 0, 0
	for i, c := range s {
		cur++
		if i == len(s)-1 || byte(c) != s[i+1] {
			if c == '1' {
				ans = max(ans, min(pre, cur)*2)
			}
			pre = cur
			cur = 0
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

#### 复杂度分析  

- 时间复杂度：$O(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。