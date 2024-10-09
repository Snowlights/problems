#### 题目  

<p>给你一个仅由小写英文字母组成的字符串 <code>s</code> 。在一步操作中，你可以完成以下行为：</p>

<ul>
	<li>选则 <code>s</code> 的任一非空子字符串，可能是整个字符串，接着将字符串中的每一个字符替换为英文字母表中的前一个字符。例如，&#39;b&#39; 用 &#39;a&#39; 替换，&#39;a&#39; 用 &#39;z&#39; 替换。</li>
</ul>

<p>返回执行上述操作 <strong>恰好一次</strong> 后可以获得的 <strong>字典序最小</strong> 的字符串。</p>

<p><strong>子字符串</strong> 是字符串中的一个连续字符序列。</p>
现有长度相同的两个字符串 <code>x</code> 和 字符串 <code>y</code> ，在满足 <code>x[i] != y[i]</code> 的第一个位置 <code>i</code> 上，如果  <code>x[i]</code> 在字母表中先于 <code>y[i]</code> 出现，则认为字符串 <code>x</code> 比字符串 <code>y</code> <strong>字典序更小</strong> 。

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>s = &#34;cbabc&#34;
<strong>输出：</strong>&#34;baabc&#34;
<strong>解释：</strong>我们选择从下标 0 开始、到下标 1 结束的子字符串执行操作。 
可以证明最终得到的字符串是字典序最小的。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>s = &#34;acbbc&#34;
<strong>输出：</strong>&#34;abaab&#34;
<strong>解释：</strong>我们选择从下标 1 开始、到下标 4 结束的子字符串执行操作。
可以证明最终得到的字符串是字典序最小的。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>s = &#34;leetcode&#34;
<strong>输出：</strong>&#34;kddsbncd&#34;
<strong>解释：</strong>我们选择整个字符串执行操作。
可以证明最终得到的字符串是字典序最小的。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 3 * 10<sup>5</sup></code></li>
	<li><code>s</code> 仅由小写英文字母组成</li>
</ul>
 
#### 思路  

根据题意，把 $\texttt{a}$ 替换成 $\texttt{z}$ 会让字典序变大，所以子串里面是不能包含 $\texttt{a}$ 的。

替换其它字符可以让字典序变小。

那么从左到右找到第一个不等于 $\texttt{a}$ 的字符 $s[i]$，继续向后遍历，每个字符都减一，直到 $s$ 末尾或者遇到了 $\texttt{a}$。

特别地，如果 $s$ 全为 $\texttt{a}$，由于题目要求选择的子串是非空的，且必须操作一次，那么就把最后一个 $\texttt{a}$ 改成 $\texttt{z}$。

```go 
func smallestString(s string) (ans string) {
	body, i := []byte(s), 0
	for i < len(s) && body[i] == 'a' {
		i++
	}
	flag := false
	for i < len(s) && body[i] != 'a' {
		body[i]--
		flag = true
		i++
	}

	if !flag {
		body[len(body)-1]--
		if body[len(body)-1] < 'a' {
			body[len(body)-1] = 'z'
		}
	}

	return string(body)
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$，如果可以直接修改 $s$ 则为 $\mathcal{O}(1)$ 空间（C++）。
