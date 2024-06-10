#### 题目

<p>给你一棵 <strong>树</strong>（即，一个连通、无向且无环的图），<strong>根</strong> 节点为 <code>0</code> ，由编号从 <code>0</code> 到 <code>n - 1</code> 的 <code>n</code> 个节点组成。这棵树用一个长度为 <code>n</code> 、下标从 <strong>0</strong> 开始的数组 <code>parent</code> 表示，其中 <code>parent[i]</code> 为节点 <code>i</code> 的父节点，由于节点 <code>0</code> 为根节点，所以 <code>parent[0] == -1</code> 。</p>

<p>另给你一个长度为 <code>n</code> 的字符串 <code>s</code> ，其中 <code>s[i]</code> 是分配给 <code>i</code> 和 <code>parent[i]</code> 之间的边的字符。<code>s[0]</code> 可以忽略。</p>

<p>找出满足 <code>u < v</code> ，且从 <code>u</code> 到 <code>v</code> 的路径上分配的字符可以 <strong>重新排列</strong> 形成 <strong>回文</strong> 的所有节点对 <code>(u, v)</code> ，并返回节点对的数目。</p>

<p>如果一个字符串正着读和反着读都相同，那么这个字符串就是一个 <strong>回文</strong> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/07/15/treedrawio-8drawio.png" style="width: 281px; height: 181px;"/></p>

<pre><strong>输入：</strong>parent = [-1,0,0,1,1,2], s = "acaabc"
<strong>输出：</strong>8
<strong>解释：</strong>符合题目要求的节点对分别是：
- (0,1)、(0,2)、(1,3)、(1,4) 和 (2,5) ，路径上只有一个字符，满足回文定义。
- (2,3)，路径上字符形成的字符串是 "aca" ，满足回文定义。
- (1,5)，路径上字符形成的字符串是 "cac" ，满足回文定义。
- (3,5)，路径上字符形成的字符串是 "acac" ，可以重排形成回文 "acca" 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>parent = [-1,0,0,0,0], s = "aaaaa"
<strong>输出：</strong>10
<strong>解释：</strong>任何满足 u < v 的节点对 (u,v) 都符合题目要求。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == parent.length == s.length</code></li>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
	<li>对于所有 <code>i >= 1</code> ，<code>0 <= parent[i] <= n - 1</code> 均成立</li>
	<li><code>parent[0] == -1</code></li>
	<li><code>parent</code> 表示一棵有效的树</li>
	<li><code>s</code> 仅由小写英文数字组成</li>
</ul>

#### 思路

#### 提示 1

回文串等价于至多一个字母出现奇数次，其余字母出现偶数次。

#### 提示 2

用一个长为 $26$ 的二进制数来压缩存储每个字母的奇偶性。一条边都可以看成是 `1<<(s[i]-'a')`。那么路径所对应的二进制数，就是路径上的所有边的异或和（因为异或就是模 $2$ 剩余系中的加法，刚好可以表示奇偶性）。只有 $27$ 个二进制数符合要求：

- $0$，表示每个字母都出现偶数次。
- $2^0,2^1,\cdots,2^{25}$，表示第 $i$ 个字母出现奇数次，其余字母出现偶数次。

#### 提示 3

设 $v$ 和 $w$ 的最近公共祖先为 $lca$，设从根到 $x$ 的路径异或和为 $XOR_{v}$。
$v$ 到 $w$ 的路径可以看成是 $v-\textit{lca}-w$，其中 $\textit{lca}$ 到 $v$ 的路径异或和，等于根到 $\textit{v}$ 的异或和，
再异或上根到 $\textit{lca}$ 的异或和。（从根到 $\textit{lca}$ 的边异或了两次，等于 $0$ 抵消掉。）
所以 $v-\textit{lca}-w$ 异或和为

$$
(XOR_{v} \oplus XOR_{lca}) \oplus (XOR_{w} \oplus XOR_{lca})

$$

$XOR_{lca}$ 异或了两次，抵消掉，所以上式为

$$
XOR_{v} \oplus XOR_{w}

$$

把所有 $\textit{XOR}$ 求出来，就变成判断这 $n-1$ 个数当中：

- 两数异或和是否为 $0$？这意味着路径上的每个字母都出现偶数次。
- 两数异或和是否为 $2$ 的幂？这意味着路径上恰好有个字母出现奇数次，其余字母出现偶数次。
- 特殊情况：$XOR_{v}=0$ 或者 $XOR_{v}$ 为 $2$ 的幂，表示从根到 $v$ 的路径符合要求，我们可以异或上一条「空路径」对应的异或值，即 $0$，就转换成了上面两数异或和的情况。

这可以用类似两数之和的思路解决，用哈希表记录 $\textit{XOR}_{v}$ 的个数，设当前算出的异或和为 $x$，去哈希表中找 $x$ 的出现次数以及 $x\oplus 2^i$ 的出现次数。

```go
func countPalindromePaths(parent []int, s string) int64 {
	n := len(parent)
	type pair struct{ to, wt int }
	g := make([][]pair, n)
	for i := 1; i < n; i++ {
		p := parent[i]
		g[p] = append(g[p], pair{i, 1 << (s[i] - 'a')})
	}

	ans := 0
	cnt := map[int]int{0: 1} // 一条「空路径」
	var dfs func(int, int)
	dfs = func(v, xor int) {
		for _, e := range g[v] {
			x := xor ^ e.wt
			ans += cnt[x] // x ^ x = 0
			for i := 0; i < 26; i++ {
				ans += cnt[x^(1<<i)] // x ^ (x^(1<<i)) = 1<<i
			}
			cnt[x]++
			dfs(e.to, x)
		}
	}
	dfs(0, 0)
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$
