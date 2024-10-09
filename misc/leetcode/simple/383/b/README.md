#### 题目

p>给你一个下标从 <strong>0</strong> 开始的字符串 <code>word</code> 和一个整数 <code>k</code> 。</p>

<p>在每一秒，你必须执行以下操作：</p>

<ul>
	<li>移除 <code>word</code> 的前 <code>k</code> 个字符。</li>
	<li>在 <code>word</code> 的末尾添加 <code>k</code> 个任意字符。</li>
</ul>

<p><strong>注意 </strong>添加的字符不必和移除的字符相同。但是，必须在每一秒钟都执行 <strong>两种 </strong>操作。</p>

<p>返回将 <code>word</code> 恢复到其 <strong>初始 </strong>状态所需的 <strong>最短 </strong>时间（该时间必须大于零）。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>word = "abacaba", k = 3
<strong>输出：</strong>2
<strong>解释：</strong>
第 1 秒，移除 word 的前缀 "aba"，并在末尾添加 "bac" 。因此，word 变为 "cababac"。
第 2 秒，移除 word 的前缀 "cab"，并在末尾添加 "aba" 。因此，word 变为 "abacaba" 并恢复到始状态。
可以证明，2 秒是 word 恢复到其初始状态所需的最短时间。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>word = "abacaba", k = 4
<strong>输出：</strong>1
<strong>解释：
</strong>第 1 秒，移除 word 的前缀 "abac"，并在末尾添加 "caba" 。因此，word 变为 "abacaba" 并恢复到初始状态。
可以证明，1 秒是 word 恢复到其初始状态所需的最短时间。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>word = "abcbabcd", k = 2
<strong>输出：</strong>4
<strong>解释：</strong>
每一秒，我们都移除 word 的前 2 个字符，并在 word 末尾添加相同的字符。
4 秒后，word 变为 "abcbabcd" 并恢复到初始状态。
可以证明，4 秒是 word 恢复到其初始状态所需的最短时间。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= word.length <= 50</code></li>
	<li><code>1 <= k <= word.length</code></li>
	<li><code>word</code>仅由小写英文字母组成。</li>
</ul>

#### 思路

如果只操作一次，就能让 $s$ 恢复成其初始值，意味着什么？  
由于可以往 $s$ 的末尾添加任意字符，这意味着只要 $s[k:]$ 是 $s$ 的前缀，就能让 $s$ 恢复成其初始值，
其中 $s[k:]$ 表示从 $s[k]$ 开始的后缀。  
例如示例 2 的 $s=\text{abacaba},\ k=4$，由于后缀 $s[4:]=\text{aba}$ 是 $s$ 的前缀，所以只需操作一次。 
如果操作一次不行，我们就看 $s[2k:]$ 是否为 $s$ 的前缀。依此类推。  
如果任意非空 $s[xk:]$（$x>0$）都不是 $s$ 的前缀（例如示例 3），那么只能操作 $\left\lceil\dfrac{n}{k}\right\rceil$ 次，
把 $s$ 的字符全部删除，由于可以添加任意字符，我们可以直接生成一个新的 $s$。  
我们通过计算 $s$ 后缀与 $s$ 的 LCP（最长公共前缀）长度，即 [z 函数（扩展 KMP）](https://oi-wiki.org/string/z-func/)来判断，如果 LCP 长度大于等于后缀长度，就说明对应操作可以让 $s$ 恢复成其初始值

```go [sol]
func minimumTimeToInitialState(s string, k int) int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		if i%k == 0 && z[i] >= n-i {
			return i / k
		}
	}
	return (n-1)/k + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
