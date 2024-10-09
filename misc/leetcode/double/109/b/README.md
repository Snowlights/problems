### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的字符串 <code>s</code> ，将 <code>s</code> 中的元素重新 <b>排列</b> 得到新的字符串 <code>t</code> ，它满足：</p>

<ul>
	<li>所有辅音字母都在原来的位置上。更正式的，如果满足 <code>0 &lt;= i &lt; s.length</code> 的下标 <code>i</code> 处的 <code>s[i]</code> 是个辅音字母，那么 <code>t[i] = s[i]</code> 。</li>
	<li>元音字母都必须以他们的 <strong>ASCII</strong> 值按 <strong>非递减</strong> 顺序排列。更正式的，对于满足 <code>0 &lt;= i &lt; j &lt; s.length</code> 的下标 <code>i</code> 和 <code>j</code>  ，如果 <code>s[i]</code> 和 <code>s[j]</code> 都是元音字母，那么 <code>t[i]</code> 的 ASCII 值不能大于 <code>t[j]</code> 的 ASCII 值。</li>
</ul>

<p>请你返回结果字母串。</p>

<p>元音字母为 <code>&#39;a&#39;</code> ，<code>&#39;e&#39;</code> ，<code>&#39;i&#39;</code> ，<code>&#39;o&#39;</code> 和 <code>&#39;u&#39;</code> ，它们可能是小写字母也可能是大写字母，辅音字母是除了这 5 个字母以外的所有字母。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>s = &#34;lEetcOde&#34;
<b>输出：</b>&#34;lEOtcede&#34;
<b>解释：</b>&#39;E&#39; ，&#39;O&#39; 和 &#39;e&#39; 是 s 中的元音字母，&#39;l&#39; ，&#39;t&#39; ，&#39;c&#39; 和 &#39;d&#39; 是所有的辅音。将元音字母按照 ASCII 值排序，辅音字母留在原地。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>s = &#34;lYmpH&#34;
<b>输出：</b>&#34;lYmpH&#34;
<b>解释：</b>s 中没有元音字母（s 中都为辅音字母），所以我们返回 &#34;lYmpH&#34; 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> 只包含英语字母表中的 <strong>大写 </strong>和 <strong>小写 </strong>字母。</li>
</ul>
 
### 思路 

模拟、排序、拼接返回

```go 
func sortVowels(s string) (ans string) {
	in := func(bb byte) bool{
		return bb == 'a' || bb == 'e' || bb == 'i' || bb == 'o' || bb=='u'||
			bb == 'A' || bb == 'E' || bb == 'I' || bb == 'O' || bb=='U'
	}
	a, b := make([]byte, 0), []byte(s)
	for _, v := range b {
		if in(v) {
			a = append(a, v)
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	j := 0
	for i, v := range b {
		if in(v) {
			b[i] = a[j]
			j++
		}
	}

	return string(b)
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{s}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$ 。