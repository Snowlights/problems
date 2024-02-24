### 题目

<p>给你一个整数&nbsp;<code>n</code>&nbsp;和一个下标从&nbsp;<strong>0</strong>&nbsp;开始的字符串数组&nbsp;<code>words</code>&nbsp;，和一个下标从&nbsp;<strong>0</strong>&nbsp;开始的数组&nbsp;<code>groups</code>&nbsp;，两个数组长度都是&nbsp;<code>n</code>&nbsp;。</p>

<p>两个长度相等字符串的 <strong>汉明距离</strong>&nbsp;定义为对应位置字符&nbsp;<strong>不同</strong>&nbsp;的数目。</p>

<p>你需要从下标&nbsp;<code>[0, 1, ..., n - 1]</code>&nbsp;中选出一个&nbsp;<strong>最长子序列</strong>&nbsp;，将这个子序列记作长度为 <code>k</code> 的&nbsp;<code>[i<sub>0</sub>, i<sub>1</sub>, ..., i<sub>k - 1</sub>]</code>&nbsp;，它需要满足以下条件：</p>

<ul>
	<li><strong>相邻</strong> 下标对应的 <code>groups</code> 值 <strong>不同</strong>。即，对于所有满足&nbsp;<code>0 &lt; j + 1 &lt; k</code>&nbsp;的&nbsp;<code>j</code>&nbsp;都有&nbsp;<code>groups[i<sub>j</sub>] != groups[i<sub>j + 1</sub>]</code>&nbsp;。</li>
	<li>对于所有&nbsp;<code>0 &lt; j + 1 &lt; k</code>&nbsp;的下标&nbsp;<code>j</code>&nbsp;，都满足&nbsp;<code>words[i<sub>j</sub>]</code> 和&nbsp;<code>words[i<sub>j + 1</sub>]</code>&nbsp;的长度 <strong>相等</strong>&nbsp;，且两个字符串之间的 <strong>汉明距离</strong>&nbsp;为 <code>1</code>&nbsp;。</li>
</ul>

<p>请你返回一个字符串数组，它是下标子序列&nbsp;<strong>依次</strong>&nbsp;对应&nbsp;<code>words</code>&nbsp;数组中的字符串连接形成的字符串数组。如果有多个答案，返回任意一个。</p>

<p><strong>子序列</strong>&nbsp;指的是从原数组中删掉一些（也可能一个也不删掉）元素，剩余元素不改变相对位置得到的新的数组。</p>

<p><b>注意：</b><code>words</code>&nbsp;中的字符串长度可能&nbsp;<strong>不相等</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>n = 3, words = ["bab","dab","cab"], groups = [1,2,2]
<b>输出：</b>["bab","cab"]
<b>解释：</b>一个可行的子序列是 [0,2] 。
- groups[0] != groups[2]
- words[0].length == words[2].length 且它们之间的汉明距离为 1 。
所以一个可行的答案是 [words[0],words[2]] = ["bab","cab"] 。
另一个可行的子序列是 [0,1] 。
- groups[0] != groups[1]
- words[0].length = words[1].length 且它们之间的汉明距离为 1 。
所以另一个可行的答案是 [words[0],words[1]] = ["bab","dab"] 。
符合题意的最长子序列的长度为 2 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>n = 4, words = ["a","b","c","d"], groups = [1,2,3,4]
<b>输出：</b>["a","b","c","d"]
<b>解释：</b>我们选择子序列 [0,1,2,3] 。
它同时满足两个条件。
所以答案为 [words[0],words[1],words[2],words[3]] = ["a","b","c","d"] 。
它是所有下标子序列里最长且满足所有条件的。
所以它是唯一的答案。
</pre>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>1 &lt;= n == words.length == groups.length &lt;= 1000</code></li>
	<li><code>1 &lt;= words[i].length &lt;= 10</code></li>
	<li><code>1 &lt;= groups[i] &lt;= n</code></li>
	<li><code>words</code>&nbsp;中的字符串&nbsp;<strong>互不相同</strong>&nbsp;。</li>
	<li><code>words[i]</code> 只包含小写英文字母。</li>
</ul>


### 思路

## 子序列 DP 的思考套路

- 子序列 + 不考虑相邻元素：选或不选。代表题目：[494. 目标和（0-1 背包）](https://leetcode.cn/problems/target-sum/)
- 子序列 + 考虑相邻元素：枚举选哪个。代表题目：[300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)

## 本题思路

本题属于「子序列 + 考虑相邻元素」，用枚举选哪个解决，状态定义类似最长递增子序列。

定义 $f[i]$ 表示从 $i$ 到 $n-1$ 中，我们选出的最长子序列的长度。定义成后缀是为了方便后面输出具体方案。

初始值 $f[i]=1$，表示选择它自己作为子序列。

如果 $\textit{groups}[j] \ne \textit{groups}[i]$ 并且 $\textit{words}[j]$ 和 $\textit{words}[i]$ 满足题目要求，并且 $f[j]+1 > f[i]$，那么更新

$$
f[i] = f[j] + 1
$$

并且记录转移来源 $\textit{from}[i] = j$。

那么最长子序列的长度就是 $\max(f)$。

## 如何输出方案

设 $\textit{mx}$ 是 $\max(f)$ 的下标，即 $f[\textit{mx}]=\max(f)$。

从 $\textit{mx}$ 开始不断循环，每次把 $\textit{words}[mx]$ 加入答案，然后更新

$$
mx = \textit{from}[mx]
$$

表示顺着转移来源往右走。

当找到了 $\max(f)$ 个字符串时停止循环。

```go [sol-Go]
func ok(s, t string) (diff bool) {
	if len(s) != len(t) {
		return
	}
	for i := range s {
		if s[i] != t[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return
}

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	f := make([]int, n)
	from := make([]int, n)
	mx := n - 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if f[j] > f[i] && groups[j] != groups[i] && ok(words[i], words[j]) {
				f[i] = f[j]
				from[i] = j
			}
		}
		f[i]++ // 加一写在这里
		if f[i] > f[mx] {
			mx = i
		}
	}

	ans := make([]string, f[mx])
	for i := range ans {
		ans[i] = words[mx]
		mx = from[mx]
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2L)$，其中 $L$ 为 $\textit{words}[i]$ 的长度，至多为 $10$。
- 空间复杂度：$\mathcal{O}(n)$。
