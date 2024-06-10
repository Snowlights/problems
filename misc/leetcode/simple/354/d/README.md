#### 题目

<p>给你一个字符串 <code>word</code> 和一个字符串数组 <code>forbidden</code> 。</p>

<p>如果一个字符串不包含 <code>forbidden</code> 中的任何字符串，我们称这个字符串是 <strong>合法</strong> 的。</p>

<p>请你返回字符串 <code>word</code> 的一个 <strong>最长合法子字符串</strong> 的长度。</p>

<p><strong>子字符串</strong> 指的是一个字符串中一段连续的字符，它可以为空。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>word = "cbaaaabc", forbidden = ["aaa","cb"]
<b>输出：</b>4
<b>解释：</b>总共有 9 个合法子字符串："c" ，"b" ，"a" ，"ba" ，"aa" ，"bc" ，"baa" ，"aab" 和 "aabc" 。最长合法子字符串的长度为 4 。
其他子字符串都要么包含 "aaa" ，要么包含 "cb" 。</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>word = "leetcode", forbidden = ["de","le","e"]
<strong>输出：</strong>4
<b>解释：</b>总共有 11 个合法子字符串："l" ，"t" ，"c" ，"o" ，"d" ，"tc" ，"co" ，"od" ，"tco" ，"cod" 和 "tcod" 。最长合法子字符串的长度为 4 。
所有其他子字符串都至少包含 "de" ，"le" 和 "e" 之一。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= word.length <= 10<sup>5</sup></code></li>
	<li><code>word</code> 只包含小写英文字母。</li>
	<li><code>1 <= forbidden.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= forbidden[i].length <= 10</code></li>
	<li><code>forbidden[i]</code> 只包含小写英文字母。</li>
</ul>

#### 思路

#### 提示 1

$\textit{forbidden}[i]$ 的长度不超过 $10$。

#### 提示 2

[同向双指针]。

#### 提示 3

初始化子串左端点 $\textit{left}=0$，枚举子串右端点 $\textit{right}$。
对于示例 2，只要 $\textit{right}\ge 1$，那么合法子串是不能包含 $\texttt{le}$ 的，所以左端点 $\textit{left}$ 必须向右移，不可能再回到 $0$（否则就包含 $\texttt{le}$ 了）。
因为左端点只会向右移动，不会向左移动，这样的**单调性**保证了算法的效率。
当 $\textit{right}$ 右移到一个新的字母时，**枚举**以该字母为右端点的 $\textit{forbidden}[i]$ 的最短长度。
如果发现子串 $\textit{word}[i]$ 到 $\textit{word}[\textit{right}]$ 在 $\textit{forbidden}$ 中（用哈希表实现），
那么更新 $\textit{left}=i+1$ 并结束枚举，从而避免合法子串包含 $\textit{forbidden}$ 中的字符串。枚举结束后，
更新答案为合法子串长度 $\textit{right}-\textit{left}+1$ 的最大值。

```go
func longestValidSubstring(word string, forbidden []string) (ans int) {
	has := make(map[string]bool, len(forbidden))
	for _, s := range forbidden {
		has[s] = true
	}

	left := 0
	for right := range word {
		for i := right; i >= left && i > right-10; i-- {
			if has[word[i:right+1]] {
				left = i + 1 // 当子串右端点 >= right 时，合法子串一定不能包含 word[i]
				break
			}
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L+nM^2)$，其中 $L$ 为所有 $\textit{forbidden}[i]$ 的长度之和，$n$ 为 $\textit{word}$ 的长度，$M=10$ 表示 $\textit{forbidden}[i]$ 的最长长度。请注意，在哈希表中查询一个长为 $M$ 的字符串的时间是 $\mathcal{O}(M)$，每次移动右指针会执行至多 $M$ 次这样的查询。
- 空间复杂度：$\mathcal{O}(L)$。
