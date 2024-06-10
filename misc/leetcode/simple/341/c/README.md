#### 题目

<p>给你一个字符串 <code>word</code> ，你可以向其中任何位置插入 "a"、"b" 或 "c" 任意次，返回使 <code>word</code> <strong>有效</strong> 需要插入的最少字母数。</p>

<p>如果字符串可以由 "abc" 串联多次得到，则认为该字符串 <strong>有效</strong> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>word = "b"
<strong>输出：</strong>2
<strong>解释：</strong>在 "b" 之前插入 "a" ，在 "b" 之后插入 "c" 可以得到有效字符串 "<strong>a</strong>b<strong>c</strong>" 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>word = "aaa"
<strong>输出：</strong>6
<strong>解释：</strong>在每个 "a" 之后依次插入 "b" 和 "c" 可以得到有效字符串 "a<strong>bc</strong>a<strong>bc</strong>a<strong>bc</strong>" 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>word = "abc"
<strong>输出：</strong>0
<strong>解释：</strong>word 已经是有效字符串，不需要进行修改。 
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= word.length <= 50</code></li>
	<li><code>word</code> 仅由字母 "a"、"b" 和 "c" 组成。</li>
</ul>

#### 思路

下文将 $\textit{word}$ 简记为 $s$。

# 方法一：考虑相邻字母

对于两个相邻字符 $x$ 和 $y$（$x$ 在 $y$ 左侧），使 $s$ 有效的话需要插入

$$
y-x-1

$$

个字母。
考虑到这可能是个负数，可以通过如下技巧转换在 $[0,2]$ 内：

$$
(y-x-1+3)\bmod 3

$$

- 例如 $x=\text{`a'},y=\text{`c'}$，则有 $(\text{`c'}-\text{`a'}+2)\bmod 3 = 1$，意思是需要补一个字母 $\text{`b'}$。
- 例如 $x=\text{`c'},y=\text{`a'}$，则有 $(\text{`a'}-\text{`c'}+2)\bmod 3 = 0$，无需补字母。

最后补齐开头的 $s[0]-\text{`a'}$，和结尾的 $\text{`c'}-s[n-1]$。这俩可以合并为 $s[0]-s[n-1]+2$。

```go  
func addMinimum(s string) int {
	ans := int(s[0]) - int(s[len(s)-1]) + 2
	for i := 1; i < len(s); i++ {
		ans += (int(s[i]) - int(s[i-1]) + 2) % 3
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。

# 方法二：考虑 abc 的周期数

计算出 $\text{`abc'}$ 的周期数 $t$，那么有效字符串的长度为 $3t$，需要插入的字符个数为 $3t-n$。
对于两个相邻字符 $x$ 和 $y$（$x$ 在 $y$ 左侧），如果 $x<y$，那么 $x$ 和 $y$ 可以在同一个 $\text{`abc'}$ 周期内，否则一定不在。
所以 $t$ 就是 $x\ge y$ 的次数。

```go
func addMinimum(s string) int {
	t := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] >= s[i] { // 必须生成一个新的 abc
			t++
		}
	}
	return t*3 - len(s)
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
- 空间复杂度：$O(1)$。仅用到若干额外变量。
