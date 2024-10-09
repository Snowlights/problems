#### 题目

<p>给你两个下标从 <strong>0</strong> 开始的字符串 <code>source</code> 和 <code>target</code> ，它们的长度均为 <code>n</code> 并且由 <strong>小写 </strong>英文字母组成。</p>

<p>另给你两个下标从 <strong>0</strong> 开始的字符串数组 <code>original</code> 和 <code>changed</code> ，以及一个整数数组 <code>cost</code> ，其中 <code>cost[i]</code> 代表将字符串 <code>original[i]</code> 更改为字符串 <code>changed[i]</code> 的成本。</p>

<p>你从字符串 <code>source</code> 开始。在一次操作中，<strong>如果 </strong>存在 <strong>任意</strong> 下标 <code>j</code> 满足 <code>cost[j] == z</code>  、<code>original[j] == x</code> 以及 <code>changed[j] == y</code> ，你就可以选择字符串中的 <strong>子串</strong> <code>x</code> 并以 <code>z</code> 的成本将其更改为 <code>y</code> 。 你可以执行 <strong>任意数量 </strong>的操作，但是任两次操作必须满足<strong> 以下两个 </strong>条件 <strong>之一</strong> ：</p>

<ul>
	<li>在两次操作中选择的子串分别是 <code>source[a..b]</code> 和 <code>source[c..d]</code> ，满足 <code>b < c</code>  <strong>或</strong> <code>d < a</code> 。换句话说，两次操作中选择的下标<strong> 不相交 </strong>。</li>
	<li>在两次操作中选择的子串分别是 <code>source[a..b]</code> 和 <code>source[c..d]</code> ，满足 <code>a == c</code> <strong>且</strong> <code>b == d</code> 。换句话说，两次操作中选择的下标<strong> 相同 </strong>。</li>
</ul>

<p>返回将字符串 <code>source</code> 转换为字符串 <code>target</code> 所需的<strong> 最小 </strong>成本。如果不可能完成转换，则返回 <code>-1</code> 。</p>

<p><strong>注意</strong>，可能存在下标 <code>i</code> 、<code>j</code> 使得 <code>original[j] == original[i]</code> 且 <code>changed[j] == changed[i]</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>source = "abcd", target = "acbe", original = ["a","b","c","c","e","d"], changed = ["b","c","b","e","b","e"], cost = [2,5,5,1,2,20]
<strong>输出：</strong>28
<strong>解释：</strong>将 "abcd" 转换为 "acbe"，执行以下操作：
- 将子串 source[1..1] 从 "b" 改为 "c" ，成本为 5 。
- 将子串 source[2..2] 从 "c" 改为 "e" ，成本为 1 。
- 将子串 source[2..2] 从 "e" 改为 "b" ，成本为 2 。
- 将子串 source[3..3] 从 "d" 改为 "e" ，成本为 20 。
产生的总成本是 5 + 1 + 2 + 20 = 28 。 
可以证明这是可能的最小成本。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>source = "abcdefgh", target = "acdeeghh", original = ["bcd","fgh","thh"], changed = ["cde","thh","ghh"], cost = [1,3,5]
<strong>输出：</strong>9
<strong>解释：</strong>将 "abcdefgh" 转换为 "acdeeghh"，执行以下操作：
- 将子串 source[1..3] 从 "bcd" 改为 "cde" ，成本为 1 。
- 将子串 source[5..7] 从 "fgh" 改为 "thh" ，成本为 3 。可以执行此操作，因为下标 [5,7] 与第一次操作选中的下标不相交。
- 将子串 source[5..7] 从 "thh" 改为 "ghh" ，成本为 5 。可以执行此操作，因为下标 [5,7] 与第一次操作选中的下标不相交，且与第二次操作选中的下标相同。
产生的总成本是 1 + 3 + 5 = 9 。
可以证明这是可能的最小成本。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>source = "abcdefgh", target = "addddddd", original = ["bcd","defgh"], changed = ["ddd","ddddd"], cost = [100,1578]
<strong>输出：</strong>-1
<strong>解释：</strong>无法将 "abcdefgh" 转换为 "addddddd" 。
如果选择子串 source[1..3] 执行第一次操作，以将 "abcdefgh" 改为 "adddefgh" ，你无法选择子串 source[3..7] 执行第二次操作，因为两次操作有一个共用下标 3 。
如果选择子串 source[3..7] 执行第一次操作，以将 "abcdefgh" 改为 "abcddddd" ，你无法选择子串 source[1..3] 执行第二次操作，因为两次操作有一个共用下标 3 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= source.length == target.length <= 1000</code></li>
	<li><code>source</code>、<code>target</code> 均由小写英文字母组成</li>
	<li><code>1 <= cost.length == original.length == changed.length <= 100</code></li>
	<li><code>1 <= original[i].length == changed[i].length <= source.length</code></li>
	<li><code>original[i]</code>、<code>changed[i]</code> 均由小写英文字母组成</li>
	<li><code>original[i] != changed[i]</code></li>
	<li><code>1 <= cost[i] <= 10<sup>6</sup></code></li>
</ul>

#### 思路

1. 把每个字符串转换成一个整数编号，这一步可以用**字典树**完成。见 [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)。
2. 建图，从 $\textit{original}[i]$ 向 $\textit{changed}[i]$ 连边，边权为 $\textit{cost}[i]$。
3. 用 Floyd 算法求图中任意两点最短路，得到 $\textit{dis}$ 矩阵，原理请看[【图解】带你发明 Floyd 算法！](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/)这里得到的 $\textit{dis}[i][j]$ 表示编号为 $i$ 的子串，通过若干次替换操作变成编号为 $j$ 的子串的最小成本。
4. 动态规划。定义 $\textit{dfs}(i)$ 表示从 $\textit{source}[i]$ 开始向后修改的最小成本。
5. 如果 $\textit{source}[i] = \textit{target}[i]$，可以不修改，$\textit{dfs}(i) = \textit{dfs}(i+1)$。
6. 也可以从 $\textit{source}[i]$ 开始向后修改，利用字典树**快速判断** $\textit{source}$ 和 $\textit{target}$ 的下标从 $i$ 到 $j$ 的子串是否在 $\textit{original}$ 和 $\textit{changed}$ 中，如果在就用 $\textit{dis}[x][y] + \textit{dfs}(j+1)$ 更新 $\textit{dfs}(i)$ 的最小值，其中 $x$ 和 $y$ 分别是 $\textit{source}$ 和 $\textit{target}$ 的这段子串对应的编号。
7. 递归边界 $\textit{dfs}(n) = 0$。
8. 递归入口 $\textit{dfs}(0)$，即为答案。如果答案是无穷大则返回 $-1$。

```go  [sol]
func minimumCost(source, target string, original, changed []string, cost []int) int64 {
	const inf = math.MaxInt / 2
	type node struct {
		son [26]*node
		sid int // 字符串的编号
	}
	root := &node{}
	sid := 0
	put := func(s string) int {
		o := root
		for _, b := range s {
			b -= 'a'
			if o.son[b] == nil {
				o.son[b] = &node{sid: -1}
			}
			o = o.son[b]
		}
		if o.sid < 0 {
			o.sid = sid
			sid++
		}
		return o.sid
	}

	// 初始化距离矩阵
	m := len(cost)
	dis := make([][]int, m*2)
	for i := range dis {
		dis[i] = make([]int, m*2)
		for j := range dis[i] {
			if j != i {
				dis[i][j] = inf
			}
		}
	}
	for i, c := range cost {
		x := put(original[i])
		y := put(changed[i])
		dis[x][y] = min(dis[x][y], c)
	}

	// Floyd 求任意两点最短路
	for k := 0; k < sid; k++ {
		for i := 0; i < sid; i++ {
			if dis[i][k] == inf {
				continue
			}
			for j := 0; j < sid; j++ {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	n := len(source)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i >= n {
			return 0
		}
		ptr := &memo[i]
		if *ptr != -1 {
			return *ptr
		}
		res := inf
		if source[i] == target[i] {
			res = dfs(i + 1) // 不修改 source[i]
		}
		p, q := root, root
		for j := i; j < n; j++ {
			p = p.son[source[j]-'a']
			q = q.son[target[j]-'a']
			if p == nil || q == nil {
				break
			}
			if p.sid >= 0 && q.sid >= 0 {
				// 修改从 i 到 j 的这一段
				res = min(res, dis[p.sid][q.sid]+dfs(j+1))
			}
		}
		*ptr = res
		return res
	}
	ans := dfs(0)
	if ans == inf {
		return -1
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2+mn+m^3)$，其中 $n$ 为 $\textit{source}$ 的长度，$m$ 为 $\textit{cost}$ 的长度。DP 需要 $\mathcal{O}(n^2)$ 的时间，把 $2m$ 个长度至多为 $n$ 的字符串插入字典树需要 $\mathcal{O}(mn)$ 的时间，Floyd 需要 $\mathcal{O}(m^3)$ 的时间。
- 空间复杂度：$\mathcal{O}(n+mn+m^2)$。DP 需要 $\mathcal{O}(n)$ 的空间，把 $2m$ 个长度至多为 $n$ 的字符串插入字典树需要 $\mathcal{O}(mn)$ 的空间，Floyd 需要 $\mathcal{O}(m^2)$ 的空
