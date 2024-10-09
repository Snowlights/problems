### 题目

<p>给你一个下标从 <strong>0</strong> 开始的数组 <code>words</code> ，它包含 <code>n</code> 个字符串。</p>

<p>定义 <strong>连接</strong> 操作 <code>join(x, y)</code> 表示将字符串 <code>x</code> 和 <code>y</code> 连在一起，得到 <code>xy</code> 。如果 <code>x</code> 的最后一个字符与 <code>y</code> 的第一个字符相等，连接后两个字符中的一个会被 <strong>删除</strong> 。</p>

<p>比方说 <code>join("ab", "ba") = "aba"</code> ， <code>join("ab", "cde") = "abcde"</code> 。</p>

<p>你需要执行 <code>n - 1</code> 次 <strong>连接</strong> 操作。令 <code>str<sub>0</sub> = words[0]</code> ，从 <code>i = 1</code> 直到 <code>i = n - 1</code> ，对于第 <code>i</code> 个操作，你可以执行以下操作之一：</p>

<ul>
	<li>令 <code>str<sub>i</sub> = join(str<sub>i - 1</sub>, words[i])</code></li>
	<li>令 <code>str<sub>i</sub> = join(words[i], str<sub>i - 1</sub>)</code></li>
</ul>

<p>你的任务是使 <code>str<sub>n - 1</sub></code> 的长度<strong> 最小 </strong>。</p>

<p>请你返回一个整数，表示 <code>str<sub>n - 1</sub></code> 的最小长度。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>words = ["aa","ab","bc"]
<b>输出：</b>4
<strong>解释：</strong>这个例子中，我们按以下顺序执行连接操作，得到 <code>str<sub>2</sub></code> 的最小长度：
<code>str<sub>0</sub> = "aa"</code>
<code>str<sub>1</sub> = join(str<sub>0</sub>, "ab") = "aab"
</code><code>str<sub>2</sub> = join(str<sub>1</sub>, "bc") = "aabc"</code> 
<code>str<sub>2</sub></code> 的最小长度为 4 。</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>words = ["ab","b"]
<b>输出：</b>2
<b>解释：</b>这个例子中，str<sub>0</sub> = "ab"，可以得到两个不同的 str<sub>1</sub>：
join(str<sub>0</sub>, "b") = "ab" 或者 join("b", str<sub>0</sub>) = "bab" 。
第一个字符串 "ab" 的长度最短，所以答案为 2 。
</pre>

<p><strong>示例 3：</strong></p>

<pre><b>输入：</b>words = ["aaa","c","aba"]
<b>输出：</b>6
<b>解释：</b>这个例子中，我们按以下顺序执行连接操作，得到 <code>str<sub>2</sub> 的最小长度：</code>
<code>str<sub>0</sub> = "</code>aaa"
<code>str<sub>1</sub> = join(str<sub>0</sub>, "c") = "aaac"</code>
<code>str<sub>2</sub> = join("aba", str<sub>1</sub>) = "abaaac"</code>
<code>str<sub>2</sub></code> 的最小长度为 6 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= words.length <= 1000</code></li>
	<li><code>1 <= words[i].length <= 50</code></li>
	<li><code>words[i]</code> 中只包含小写英文字母。</li>
</ul>

### 思路

因为每次操作我们只关心 `str` 的第一个以及最后一个字符，因此维护 $f(i, j, k)$ 表示已经连接了前 $i$ 个字符串，
此时 `str` 的第一个字符是 $j$，最后一个字符是 $k$ 的最小长度。设第 $i$ 个字符串长度为 $l_i$，第一个字符为 $x_i$，最后一个字符为 $y_i$。
我们可以枚举把它接在开头还是末尾，以及原来开头/末尾的字符是哪个。

* 如果接在开头，则

  $$
  f(i, x_i, k) = \min\begin{cases}
  f(i - 1, y_i, k) + l_i - 1 & \\
  f(i - 1, j, k) + l_i & j \ne y_i\end{cases}

  $$

  * 如果接在末尾，则

    $$
    f(i, j, y_i) = \min\begin{cases}
    f(i - 1, j, x_i) + l_i - 1 & \\
    f(i - 1, j, k) + l_i & k \ne x_i \end{cases}

    $$

    初值 $f(1, x_i, y_i) = l_i$，
    答案就是 $\min (f(n, j, k))$。

```go  
func minimizeConcatenatedLength(a []string) (ans int) {

	n := len(a)
	dp := make([][26][26]int, n+1)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i == n {
			return 0
		}

		dv := &dp[i][j][k]
		if *dv >= 0 {
			return *dv
		}
		res := 0
		defer func() {
			*dv = res
		}()

		cur := a[i]
		if int(cur[len(cur)-1]-'a') == j {
			res = dfs(i+1, int(cur[0]-'a'), k) + len(cur) - 1
		} else {
			res = dfs(i+1, int(cur[0]-'a'), k) + len(cur)
		}
		if int(cur[0]-'a') == k {
			res = min(res, dfs(i+1, j, int(cur[len(cur)-1]-'a'))+len(cur)-1)
		} else {
			res = min(res, dfs(i+1, j, int(cur[len(cur)-1]-'a'))+len(cur))
		}

		return res
	}

	return dfs(1, int(a[0][0]-'a'), int(a[0][len(a[0])-1]-'a')) + len(a[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $|\Sigma|$ 表示字母表大小。
- 空间复杂度：$\mathcal{O}(n|\Sigma|^2)$ 。
