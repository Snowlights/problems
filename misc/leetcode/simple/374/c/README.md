#### 题目

<p>给你一个字符串&nbsp;<code>word</code>&nbsp;和一个整数 <code>k</code>&nbsp;。</p>

<p>如果&nbsp;<code>word</code>&nbsp;的一个子字符串 <code>s</code>&nbsp;满足以下条件，我们称它是 <strong>完全字符串：</strong></p>

<ul>
	<li><code>s</code>&nbsp;中每个字符 <strong>恰好</strong>&nbsp;出现 <code>k</code>&nbsp;次。</li>
	<li>相邻字符在字母表中的顺序 <strong>至多</strong>&nbsp;相差&nbsp;<code>2</code>&nbsp;。也就是说，<code>s</code>&nbsp;中两个相邻字符&nbsp;<code>c1</code> 和&nbsp;<code>c2</code>&nbsp;，它们在字母表中的位置相差<strong>&nbsp;至多</strong>&nbsp;为 <code>2</code> 。</li>
</ul>

<p>请你返回 <code>word</code>&nbsp;中 <strong>完全</strong>&nbsp;子字符串的数目。</p>

<p><strong>子字符串</strong>&nbsp;指的是一个字符串中一段连续 <strong>非空</strong>&nbsp;的字符序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>word = "igigee", k = 2
<b>输出：</b>3
<b>解释：</b>完全子字符串需要满足每个字符恰好出现 2 次，且相邻字符相差至多为 2 ：<em><strong>igig</strong></em>ee, igig<strong style="font-style: italic;">ee</strong>, <em><strong>igigee</strong>&nbsp;。</em>
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>word = "aaabbbccc", k = 3
<b>输出：</b>6
<b>解释：</b>完全子字符串需要满足每个字符恰好出现 3 次，且相邻字符相差至多为 2 ：<em><strong>aaa</strong></em>bbbccc, aaa<em><strong>bbb</strong></em>ccc, aaabbb<em><strong>ccc</strong></em>, <em><strong>aaabbb</strong></em>ccc, aaa<em><strong>bbbccc</strong></em>, <em><strong>aaabbbccc </strong></em>。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 10<sup>5</sup></code></li>
	<li><code>word</code>&nbsp;只包含小写英文字母。</li>
	<li><code>1 &lt;= k &lt;= word.length</code></li>
</ul>

#### 思路

「相邻字母相差至多为 $2$」这个约束把 $\textit{word}$ 划分成了多个子串 $s$，每个子串分别处理。可以用 [分组循环](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/) 找到每个子串 $s$。

对于每个子串，由于每个字符恰好出现 $k$ 次，我们可以枚举有 $m$ 种字符，这样问题就变成了：

- 长度固定为 $m\cdot k$ 的**滑动窗口**，判断每种字符是否都出现了恰好 $k$ 次。
- 
```go  
func f(s string, k int) (res int) {
	for m := 1; m <= 26 && k*m <= len(s); m++ {
		cnt := [26]int{}
		check := func() {
			for i := range cnt {
				if cnt[i] > 0 && cnt[i] != k {
					return
				}
			}
			res++
		}
		for right, in := range s {
			cnt[in-'a']++
			if left := right + 1 - k*m; left >= 0 {
				check()
				cnt[s[left]-'a']--
			}
		}
	}
	return
}

func countCompleteSubstrings(word string, k int) (ans int) {
	for i, n := 0, len(word); i < n; {
		st := i
		for i++; i < n && abs(int(word[i])-int(word[i-1])) <= 2; i++ {
		}
		ans += f(word[st:i], k)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写英文字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。忽略切片开销。