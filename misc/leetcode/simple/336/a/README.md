#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的字符串数组 <code>words</code> 和两个整数：<code>left</code> 和 <code>right</code> 。</p>

<p>如果字符串以元音字母开头并以元音字母结尾，那么该字符串就是一个 <strong>元音字符串</strong> ，其中元音字母是 <code>&#39;a&#39;</code>、<code>&#39;e&#39;</code>、<code>&#39;i&#39;</code>、<code>&#39;o&#39;</code>、<code>&#39;u&#39;</code> 。</p>

<p>返回<em> </em><code>words[i]</code> 是元音字符串的数目，其中<em> </em><code>i</code> 在闭区间 <code>[left, right]</code> 内。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>words = [&#34;are&#34;,&#34;amy&#34;,&#34;u&#34;], left = 0, right = 2
<strong>输出：</strong>2
<strong>解释：</strong>
- &#34;are&#34; 是一个元音字符串，因为它以 &#39;a&#39; 开头并以 &#39;e&#39; 结尾。
- &#34;amy&#34; 不是元音字符串，因为它没有以元音字母结尾。
- &#34;u&#34; 是一个元音字符串，因为它以 &#39;u&#39; 开头并以 &#39;u&#39; 结尾。
在上述范围中的元音字符串数目为 2 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>words = [&#34;hey&#34;,&#34;aeo&#34;,&#34;mu&#34;,&#34;ooo&#34;,&#34;artro&#34;], left = 1, right = 4
<strong>输出：</strong>3
<strong>解释：</strong>
- &#34;aeo&#34; 是一个元音字符串，因为它以 &#39;a&#39; 开头并以 &#39;o&#39; 结尾。
- &#34;mu&#34; 不是元音字符串，因为它没有以元音字母开头。
- &#34;ooo&#34; 是一个元音字符串，因为它以 &#39;o&#39; 开头并以 &#39;o&#39; 结尾。
- &#34;artro&#34; 是一个元音字符串，因为它以 &#39;a&#39; 开头并以 &#39;o&#39; 结尾。
在上述范围中的元音字符串数目为 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 1000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 10</code></li>
	<li><code>words[i]</code> 仅由小写英文字母组成</li>
	<li><code>0 &lt;= left &lt;= right &lt; words.length</code></li>
</ul>
 
#### 思路  

遍历 $[\textit{left},\textit{right}]$ 内的元素，按照要求判断。

```go 
func vowelStrings(words []string, left, right int) (ans int) {
	for _, s := range words[left : right+1] {
		if strings.Contains("aeiou", s[:1]) && strings.Contains("aeiou", s[len(s)-1:]) {
			ans++
		}
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{words}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。