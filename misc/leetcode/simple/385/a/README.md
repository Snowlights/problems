#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的字符串数组 <code>words</code> 。</p>

<p>定义一个 <strong>布尔 </strong>函数 <code>isPrefixAndSuffix</code> ，它接受两个字符串参数 <code>str1</code> 和 <code>str2</code> ：</p>

<ul>
	<li>当 <code>str1</code> 同时是 <code>str2</code> 的前缀（<span data-keyword="string-prefix">prefix</span>）和后缀（<span data-keyword="string-suffix">suffix</span>）时，<code>isPrefixAndSuffix(str1, str2)</code> 返回 <code>true</code>，否则返回 <code>false</code>。</li>
</ul>

<p>例如，<code>isPrefixAndSuffix("aba", "ababa")</code> 返回 <code>true</code>，因为 <code>"aba"</code> 既是 <code>"ababa"</code> 的前缀，也是 <code>"ababa"</code> 的后缀，但是 <code>isPrefixAndSuffix("abc", "abcd")</code> 返回<code> false</code>。</p>

<p>以整数形式，返回满足 <code>i &lt; j</code> 且 <code>isPrefixAndSuffix(words[i], words[j])</code> 为 <code>true</code> 的下标对 <code>(i, j)</code> 的<strong> 数量 </strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>words = ["a","aba","ababa","aa"]
<strong>输出：</strong>4
<strong>解释：</strong>在本示例中，计数的下标对包括：
i = 0 且 j = 1 ，因为 isPrefixAndSuffix("a", "aba") 为 true 。
i = 0 且 j = 2 ，因为 isPrefixAndSuffix("a", "ababa") 为 true 。
i = 0 且 j = 3 ，因为 isPrefixAndSuffix("a", "aa") 为 true 。
i = 1 且 j = 2 ，因为 isPrefixAndSuffix("aba", "ababa") 为 true 。
因此，答案是 4 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>words = ["pa","papa","ma","mama"]
<strong>输出：</strong>2
<strong>解释：</strong>在本示例中，计数的下标对包括：
i = 0 且 j = 1 ，因为 isPrefixAndSuffix("pa", "papa") 为 true 。
i = 2 且 j = 3 ，因为 isPrefixAndSuffix("ma", "mama") 为 true 。
因此，答案是 2 。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>words = ["abab","ab"]
<strong>输出：</strong>0
<strong>解释：</strong>在本示例中，唯一有效的下标对是 i = 0 且 j = 1 ，但是 isPrefixAndSuffix("abab", "ab") 为 false 。
因此，答案是 0 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 50</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 10</code></li>
	<li><code>words[i]</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

## 方法一：Z 函数 + 字典树

对于两个字符串 $s$ 和 $t$，设 $s$ 的长度为 $m$。如果 $s$ 既是 $t$ 的前缀又是 $t$ 的后缀，那么对于 $t$ 来说， 它的长为 $m$ 的前后缀必须相等。  
怎么快速判断 $t$ 的某个长度的前后缀相等？我们可以用 Z 函数解决。  
由于 $z[i]$ 的定义是后缀 $t[i:]$ 与 $t$ 的最长公共前缀的长度，所以只要 $z[i] = n-1-i$，那么 $t[i:]$ 和与其等长的 $t$ 的前缀是相等的。  
如果 $z[i] = n-1-i$ 成立，我们只需要判断 $s$ 是否为 $t$ 的前缀。  
枚举 $t=\textit{words}[j]$，怎么统计有多少个 $s=\textit{words}[i]$ 是 $t$ 的前缀？  
这可以用**字典树**解决，在遍历 $\textit{words}$ 的同时，维护每个字符串的出现次数。当我们遍历 $t$ 时，同时遍历字典树上的对应节点，并把 $t$ 插入字典树。具体请看代码。

```go [sol]
func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n
	return z
}

func countPrefixSuffixPairs(words []string) (ans int64) {
	type node struct {
		son [26]*node
		cnt int
	}
	root := &node{}
	for _, t := range words {
		z := calcZ(t)
		cur := root
		for i, c := range t {
			c -= 'a'
			if cur.son[c] == nil {
				cur.son[c] = &node{}
			}
			cur = cur.son[c]
			if z[len(t)-1-i] == i+1 { // t[:i+1] == t[len(t)-1-i:]
				ans += int64(cur.cnt)
			}
		}
		cur.cnt++
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有 $\textit{words}[i]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(L)$。

## 方法二：只用字典树

把每个字符串 $s$ 视作一个 pair 列表：

$$[(s[0],s[n-1]), (s[1],s[n-2]), (s[2],s[n-3]), \cdots, (s[n-1], s[0])]
$$

只要这个 pair 列表是另一个字符串 $t$ 的 pair 列表的前缀，那么 $s$ 就是 $t$ 的前后缀。

```go [sol]
func countPrefixSuffixPairs(words []string) (ans int64) {
	type pair struct{ x, y byte }
	type node struct {
		son map[pair]*node
		cnt int
	}
	root := &node{son: map[pair]*node{}}
	for _, s := range words {
		cur := root
		for i := range s {
			p := pair{s[i], s[len(s)-1-i]}
			if cur.son[p] == nil {
				cur.son[p] = &node{son: map[pair]*node{}}
			}
			cur = cur.son[p]
			ans += int64(cur.cnt)
		}
		cur.cnt++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L)$，其中 $L$ 为所有 $\textit{words}[i]$ 的长度之和。
- 空间复杂度：$\mathcal{O}(L)$。