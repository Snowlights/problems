#### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;，请你将&nbsp;<code>s</code> 分成&nbsp;<code>k</code>&nbsp;个<strong>&nbsp;子字符串</strong>&nbsp;，使得每个 <strong>子字符串</strong>&nbsp;变成&nbsp;<strong>半回文串</strong>&nbsp;需要修改的字符数目最少。</p>

<p>请你返回一个整数，表示需要修改的 <strong>最少</strong>&nbsp;字符数目。</p>

<p><strong>注意：</strong></p>

<ul>
	<li>如果一个字符串从左往右和从右往左读是一样的，那么它是一个 <strong>回文串</strong>&nbsp;。</li>
	<li>如果长度为 <code>len</code>&nbsp;的字符串存在一个满足&nbsp;<code>1 &lt;= d &lt; len</code>&nbsp;的正整数&nbsp;<code>d</code>&nbsp;，<code>len % d == 0</code>&nbsp;成立且所有对 <code>d</code>&nbsp;做除法余数相同的下标对应的字符连起来得到的字符串都是 <strong>回文串</strong>&nbsp;，那么我们说这个字符串是 <strong>半回文串</strong>&nbsp;。比方说&nbsp;<code>"aa"</code>&nbsp;，<code>"aba"</code> ，<code>"adbgad"</code>&nbsp;和&nbsp;<code>"abab"</code>&nbsp;都是 <strong>半回文串</strong>&nbsp;，而&nbsp;<code>"a"</code>&nbsp;，<code>"ab"</code>&nbsp;和&nbsp;<code>"abca"</code>&nbsp;不是。</li>
	<li><strong>子字符串</strong>&nbsp;指的是一个字符串中一段连续的字符序列。</li>
</ul>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>s = "abcac", k = 2
<b>输出：</b>1
<b>解释：</b>我们可以将 s 分成子字符串 "ab" 和 "cac" 。子字符串 "cac" 已经是半回文串。如果我们将 "ab" 变成 "aa" ，它也会变成一个 d = 1 的半回文串。
该方案是将 s 分成 2 个子字符串的前提下，得到 2 个半回文子字符串需要的最少修改次数。所以答案为 1 。</pre>

<p><strong class="example">示例 2:</strong></p>

<pre>
<b>输入：</b>s = "abcdef", k = 2
<b>输出：</b>2
<b>解释：</b>我们可以将 s 分成子字符串 "abc" 和 "def" 。子字符串 "abc" 和 "def" 都需要修改一个字符得到半回文串，所以我们总共需要 2 次字符修改使所有子字符串变成半回文串。
该方案是将 s 分成 2 个子字符串的前提下，得到 2 个半回文子字符串需要的最少修改次数。所以答案为 2 。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>s = "aabbaa", k = 3
<b>输出：</b>0
<b>解释：</b>我们可以将 s 分成子字符串 "aa" ，"bb" 和 "aa" 。
字符串 "aa" 和 "bb" 都已经是半回文串了。所以答案为 0 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= s.length &lt;= 200</code></li>
	<li><code>1 &lt;= k &lt;= s.length / 2</code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
</ul>

#### 思路

## 预处理

预处理 $s$ 的每个【长度至少为 $2$ 的】子串修改成半回文串至少要修改多少次，记到数组 $\textit{modify}$ 中，$\textit{modify}[i][j]$ 对应从 $s[i]$ 到 $s[j]$ 的子串。  
对于每个子串，枚举其因子 $d$（注意 $d$ 要小于子串长度）。例如子串 $t$ 的长度为 $6$，$d=2$ 时，我们需要判断
- $t[0]+ t[2]+ t[4]$ 组成的字符串，改成回文串需要修改多少个字母。
- $t[1]+ t[3]+ t[5]$ 组成的字符串，改成回文串需要修改多少个字母。  
  
所有修改次数相加，就是因子为 $d$ 时的修改次数。取所有修改次数的最小值，就是这个子串的最小修改次数。

## 记忆化搜索

按照【划分型 DP】的套路，定义 $\textit{dfs}(i,j)$ 表示把 $s[0]$ 到 $s[j]$ 的字符串分成 $i+1$ 个子字符串，使得每个子字符串变成半回文串需要修改的最少字符数目。  
枚举 $s[j]$ 结束的子串在 $s[L]$ 处开始，那么有

$$
\textit{dfs}(i,j) = \min_{L=2i}^{j-1} \textit{dfs}(i-1,L-1) + \textit{modify}[L][j]
$$

注意 $L$ 从 $2i$ 开始（为后面的子串预留空间），到 $j-1$ 结束（因为子串长度至少为 $2$）。

递归边界：$i=0$ 时，只剩下一个子串，返回 $\textit{modify}[0][j]$。  
递归入口：$\textit{dfs}(k-1,n-1)$，即为答案。

```go  
const mx = 200
var divisors [mx + 1][]int
func init() {
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}
}

func calc(s string) int {
	n := len(s)
	res := n
	for _, d := range divisors[n] {
		cnt := 0
		for i0 := 0; i0 < d; i0++ {
			t := []byte{}
			for i := i0; i < n; i += d {
				t = append(t, s[i])
			}
			for i, m := 0, len(t); i < m/2; i++ {
				v, w := t[i], t[m-1-i]
				if v != w {
					cnt++
				}
			}
		}
		res = min(res, cnt)
	}
	return res
}

func minimumChanges(s string, k int) (ans int) {
	n := len(s)
	modify := make([][]int, n-1)
	for l := range modify {
		modify[l] = make([]int, n)
		for r := l + 1; r < n; r++ { // 半回文串长度至少为 2
			modify[l][r] = calc(s[l : r+1])
		}
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return modify[0][j]
		}
		p := &memo[j][i]
		if *p != -1 {
			return *p
		}
		res := n
		for L := i * 2; L < j; L++ {
			res = min(res, dfs(i-1, L-1)+modify[L][j])
		}
		*p = res
		return res
	}
	return dfs(k-1, n-1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3\log n)$，其中 $n$ 为 $s$ 的长度。时间主要在预处理上，有 $\mathcal{O}(n^2)$ 个子串，平均每个子串有 $\mathcal{O}(\log n)$ 个因子，每个因子需要 $\mathcal{O}(n)$ 的时间计算修改次数。
- 空间复杂度：$\mathcal{O}(n^2)$。
