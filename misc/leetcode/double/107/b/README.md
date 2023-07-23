### 题目

<p>给你三个整数 <code>x</code> ，<code>y</code> 和 <code>z</code> 。</p>

<p>这三个整数表示你有 <code>x</code> 个 <code>"AA"</code> 字符串，<code>y</code> 个 <code>"BB"</code> 字符串，和 <code>z</code> 个 <code>"AB"</code> 字符串。你需要选择这些字符串中的部分字符串（可以全部选择也可以一个都不选择），将它们按顺序连接得到一个新的字符串。新字符串不能包含子字符串 <code>"AAA"</code> 或者 <code>"BBB"</code> 。</p>

<p>请你返回新字符串的最大可能长度。</p>

<p><strong>子字符串</strong> 是一个字符串中一段连续 <strong>非空</strong> 的字符序列。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><b>输入：</b>x = 2, y = 5, z = 1
<b>输出：</b>12
<strong>解释： </strong>我们可以按顺序连接 "BB" ，"AA" ，"BB" ，"AA" ，"BB" 和 "AB" ，得到新字符串 "BBAABBAABBAB" 。
字符串长度为 12 ，无法得到一个更长的符合题目要求的字符串。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>x = 3, y = 2, z = 2
<b>输出：</b>14
<b>解释：</b>我们可以按顺序连接 "AB" ，"AB" ，"AA" ，"BB" ，"AA" ，"BB" 和 "AA" ，得到新字符串 "ABABAABBAABBAA" 。
字符串长度为 14 ，无法得到一个更长的符合题目要求的字符串。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= x, y, z <= 50</code></li>
</ul>

### 思路

如果没有 AB，那么 AA 和 BB 只能交替连接，答案为

$$
(\min(x,y)\cdot 2 + [x\ne y])\cdot 2

$$

如果有 AB，它可以与自身连接，且只能插在 BB 和 AA 之间，即 BB + (ABABAB...) + AA。

> 或者接在后缀 BB 之后，或者加到前缀 AA 之前。

所以 AB 不会改变 AA 和 BB 交替连接的上限。
所以答案为

$$
(\min(x,y)\cdot 2 + [x\ne y] + z)\cdot 2

$$

* 记忆化搜索
  定义 $\textit{dfs}(x,y,z,k)$，其中 $x,y,z$ 为 AA/BB/AB 的使用数量，$k=0,1,2$ 表示上一个字符串是 AA/BB/AB，此时可以构造出的字符串的最大可能长度
  分类讨论：

- AA 后面只能接 BB；
- BB 后面可以接 AA 或 AB；
- AB 后面可以接 AA 或 AB。

```go  
func longestStringV1(x, y, z int) int {
	ans := min(x, y) * 2
	if x != y {
		ans++
	}
	return (ans + z) * 2
}

func min(a, b int) int { if b < a { return b }; return a }

func longestStringV2(x, y, z int) (ans int) {

	dp := make([][][][3]int, x+1)
	for i := range dp {
		dp[i] = make([][][3]int, y+1)
		for j := range dp[i] {
			dp[i][j] = make([][3]int, z+1)
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}

	var dfs func(i, j, k, pre int) int
	dfs = func(i, j, k, pre int) int {
		dv := &dp[i][j][k][pre]
		if *dv >= 0 {
			return *dv
		}
		res := 2
		defer func() {
			*dv = res
		}()

		switch pre {
		case 0:
			if j+1 <= y {
				res += dfs(i, j+1, k, 1)
			}
		case 1:
			cur := res
			if i+1 <= x {
				res = max(res, cur+dfs(i+1, j, k, 0))
			}
			if k+1 <= z {
				res = max(res, cur+dfs(i, j, k+1, 2))
			}
		case 2:
			cur := res
			if i+1 <= x {
				res = max(res, cur+dfs(i+1, j, k, 0))
			}
			if k+1 <= z {
				res = max(res, cur+dfs(i, j, k+1, 2))
			}
		}
		return res
	}

	return max(dfs(1, 0, 0, 0), max(dfs(0, 1, 0, 1), dfs(0, 0, 1, 2)))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

### 复杂度分析

> 公式法

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$ 。

> 记忆化搜索

- 时间复杂度：$\mathcal{O}(3xyz)$。
- 空间复杂度：$\mathcal{O}(3xyz)$ 。
