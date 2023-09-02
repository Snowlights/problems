### 题目

<p>给你一个字符串 <code>s</code> ，一个字符 <strong>互不相同</strong> 的字符串 <code>chars</code> 和一个长度与 <code>chars</code> 相同的整数数组 <code>vals</code> 。</p>

<p><strong>子字符串的开销</strong> 是一个子字符串中所有字符对应价值之和。空字符串的开销是 <code>0</code> 。</p>

<p><strong>字符的价值</strong> 定义如下：</p>

<ul>
	<li>如果字符不在字符串 <code>chars</code> 中，那么它的价值是它在字母表中的位置（下标从 <strong>1</strong> 开始）。

<ul>
	<li>比方说，<code>&#39;a&#39;</code> 的价值为 <code>1</code> ，<code>&#39;b&#39;</code> 的价值为 <code>2</code> ，以此类推，<code>&#39;z&#39;</code> 的价值为 <code>26</code> 。</li>
</ul>
</li>
<li>否则，如果这个字符在 <code>chars</code> 中的位置为 <code>i</code> ，那么它的价值就是 <code>vals[i]</code> 。</li>
</ul>

<p>请你返回字符串 <code>s</code> 的所有子字符串中的最大开销。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>s = "adaa", chars = "d", vals = [-1000]
<b>输出：</b>2
<b>解释：</b>字符 "a" 和 "d" 的价值分别为 1 和 -1000 。
最大开销子字符串是 "aa" ，它的开销为 1 + 1 = 2 。
2 是最大开销。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>s = "abc", chars = "abc", vals = [-1,-1,-1]
<b>输出：</b>0
<b>解释：</b>字符 "a" ，"b" 和 "c" 的价值分别为 -1 ，-1 和 -1 。
最大开销子字符串是 "" ，它的开销为 0 。
0 是最大开销。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= s.length <= 10<sup>5</sup></code></li>
	<li><code>s</code> 只包含小写英文字母。</li>
	<li><code>1 <= chars.length <= 26</code></li>
	<li><code>chars</code> 只包含小写英文字母，且 <strong>互不相同</strong> 。</li>
	<li><code>vals.length == chars.length</code></li>
	<li><code>-1000 <= vals[i] <= 1000</code></li>
</ul>

### 思路

定义 $f[i]$ 为以 $a[i]$ 结尾的最大子数组和，考虑是否接在以 $a[i-1]$ 结尾的最大子数组之后：

- 接：$f[i] = f[i-1] + a[i]$
- 不接：$f[i] = a[i]$

取最大值，则有

$$
f[i] = \max(f[i-1],0) + a[i]

$$

答案为 $\max(f)$。

```go  
func maximumCostSubstring(s, chars string, vals []int) (ans int) {
	mapping := [26]int{}
	for i := range mapping {
		mapping[i] = i + 1
	}
	for i, c := range chars {
		mapping[c-'a'] = vals[i]
	}
	f := 0
	for _, c := range s {
		f = max(f, 0) + mapping[c-'a']
		ans = max(ans, f)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```
### 复杂度分析

- 时间复杂度：$O(n+|\Sigma|)$，其中 $n$ 为 $\textit{nums}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$O(|\Sigma|)$。
