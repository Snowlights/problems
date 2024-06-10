#### 题目

<p>给你一个由 <strong>小写英文字母</strong> 组成的字符串 <code>s</code> ，你可以对其执行一些操作。在一步操作中，你可以用其他小写英文字母 <strong>替换</strong>&nbsp; <code>s</code> 中的一个字符。</p>

<p>请你执行 <strong>尽可能少的操作</strong> ，使 <code>s</code> 变成一个 <strong>回文串</strong> 。如果执行 <strong>最少</strong> 操作次数的方案不止一种，则只需选取 <strong>字典序最小</strong> 的方案。</p>

<p>对于两个长度相同的字符串 <code>a</code> 和 <code>b</code> ，在 <code>a</code> 和 <code>b</code> 出现不同的第一个位置，如果该位置上 <code>a</code> 中对应字母比 <code>b</code> 中对应字母在字母表中出现顺序更早，则认为 <code>a</code> 的字典序比 <code>b</code> 的字典序要小。</p>

<p>返回最终的回文字符串。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "egcfe"
<strong>输出：</strong>"efcfe"
<strong>解释：</strong>将 "egcfe" 变成回文字符串的最小操作次数为 1 ，修改 1 次得到的字典序最小回文字符串是 "efcfe"，只需将 'g' 改为 'f' 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "abcd"
<strong>输出：</strong>"abba"
<strong>解释：</strong>将 "abcd" 变成回文字符串的最小操作次数为 2 ，修改 2 次得到的字典序最小回文字符串是 "abba" 。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "seven"
<strong>输出：</strong>"neven"
<strong>解释：</strong>将 "seven" 变成回文字符串的最小操作次数为 1 ，修改 1 次得到的字典序最小回文字符串是 "neven" 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 1000</code></li>
	<li><code>s</code> 仅由小写英文字母组成</li>
</ul>

#### 思路

贪心替换

对于两个中心对称的字母 $x=s[i]$ 和 $y=s[n-1-i]$，如果 $x\ne y$，那么只需要修改一次，就可以让这两个字母相同：把 $x$ 改成 $y$ 或者把 $y$ 改成 $x$。

- 如果 $x>y$，那么把 $x$ 修改成 $y$ 更好，这样字典序更小。
- 如果 $x<y$，那么把 $y$ 修改成 $x$ 更好，这样字典序更小。

代码实现时可以把 $x=y$ 的情况合并到 $x<y$ 中，从而少写一个 `else if` 的判断逻辑。

```go 
func makeSmallestPalindrome(s string) (ans string) {
	sb := []byte(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if sb[i] == sb[n-1-i] {
			continue
		}
		if sb[i] > sb[n-1-i] {
			sb[i] = sb[n-1-i]
		} else if sb[i] < sb[n-1-i] {
			sb[n-1-i] = sb[i]
		}
	}

	return string(sb)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
