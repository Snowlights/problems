### 题目

<p>给你一个下标从 <strong>0</strong> 开始的数组 <code>words</code> ，数组中包含 <strong>互不相同</strong> 的字符串。</p>

<p>如果字符串 <code>words[i]</code> 与字符串 <code>words[j]</code> 满足以下条件，我们称它们可以匹配：</p>

<ul>
	<li>字符串 <code>words[i]</code> 等于 <code>words[j]</code> 的反转字符串。</li>
	<li><code>0 <= i < j < words.length</code></li>
</ul>

<p>请你返回数组 <code>words</code> 中的 <strong>最大</strong> 匹配数目。</p>

<p>注意，每个字符串最多匹配一次。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>words = ["cd","ac","dc","ca","zz"]
<b>输出：</b>2
<strong>解释：</strong>在此示例中，我们可以通过以下方式匹配 2 对字符串：
- 我们将第 0 个字符串与第 2 个字符串匹配，因为 word[0] 的反转字符串是 "dc" 并且等于 words[2]。
- 我们将第 1 个字符串与第 3 个字符串匹配，因为 word[1] 的反转字符串是 "ca" 并且等于 words[3]。
可以证明最多匹配数目是 2 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>words = ["ab","ba","cc"]
<b>输出：</b>1
<b>解释：</b>在此示例中，我们可以通过以下方式匹配 1 对字符串：
- 我们将第 0 个字符串与第 1 个字符串匹配，因为 words[1] 的反转字符串 "ab" 与 words[0] 相等。
可以证明最多匹配数目是 1 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>words = ["aa","ab"]
<b>输出：</b>0
<strong>解释：</strong>这个例子中，无法匹配任何字符串。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= words.length <= 50</code></li>
	<li><code>words[i].length == 2</code></li>
	<li><code>words</code> 包含的字符串互不相同。</li>
	<li><code>words[i]</code> 只包含小写英文字母。</li>
</ul>

### 思路

由于题目保证 $\textit{words}$ 中的字符串互不相同，所以可以遍历 $\textit{words}$， 对于 $w=\\textit{words}[i]$：
- 如果前面遇到了 $w$ 反转的字符串，那么就找到了一个匹配。
- 如果前面没有遇到，那么把 $w$ 加入一个哈希表（或者数组），方便后面快速判断是否有对应的字符串。

```go  
func maximumNumberOfStringPairs(words []string) (ans int) {
	vis := [26][26]bool{}
	for _, s := range words {
		x, y := s[0]-'a', s[1]-'a'
		if vis[y][x] {
			ans++
		} else {
			vis[x][y] = true
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{words}$ 的长度。字符串长度视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(n)$。
