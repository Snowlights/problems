#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的字符串&nbsp;<code>s</code>&nbsp;和一个整数&nbsp;<code>k</code>。</p>

<p>你需要执行以下分割操作，直到字符串&nbsp;<code>s&nbsp;</code>变为&nbsp;<strong>空</strong>：</p>

<ul>
	<li>选择&nbsp;<code>s</code>&nbsp;的最长<strong>前缀</strong>，该前缀最多包含&nbsp;<code>k&nbsp;</code>个&nbsp;<strong>不同&nbsp;</strong>字符。</li>
	<li><strong>删除&nbsp;</strong>这个前缀，并将分割数量加一。如果有剩余字符，它们在&nbsp;<code>s</code>&nbsp;中保持原来的顺序。</li>
</ul>

<p>执行操作之 <strong>前</strong> ，你可以将&nbsp;<code>s</code>&nbsp;中&nbsp;<strong>至多一处 </strong>下标的对应字符更改为另一个小写英文字母。</p>

<p>在最优选择情形下改变至多一处下标对应字符后，用整数表示并返回操作结束时得到的最大分割数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "accca", k = 2
<strong>输出：</strong>3
<strong>解释：</strong>在此示例中，为了最大化得到的分割数量，可以将 s[2] 改为 'b'。
s 变为 "acbca"。
按照以下方式执行操作，直到 s 变为空：
- 选择最长且至多包含 2 个不同字符的前缀，"<em><strong>ac</strong></em>bca"。
- 删除该前缀，s 变为 "bca"。现在分割数量为 1。
- 选择最长且至多包含 2 个不同字符的前缀，"<em><strong>bc</strong></em>a"。
- 删除该前缀，s 变为 "a"。现在分割数量为 2。
- 选择最长且至多包含 2 个不同字符的前缀，"<strong><em>a</em></strong>"。
- 删除该前缀，s 变为空。现在分割数量为 3。
因此，答案是 3。
可以证明，分割数量不可能超过 3。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "aabaab", k = 3
<strong>输出：</strong>1
<strong>解释：</strong>在此示例中，为了最大化得到的分割数量，可以保持 s 不变。
按照以下方式执行操作，直到 s 变为空： 
- 选择最长且至多包含 3 个不同字符的前缀，"<em><strong>aabaab</strong></em>"。
- 删除该前缀，s 变为空。现在分割数量为 1。
因此，答案是 1。
可以证明，分割数量不可能超过 1。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "xxyz", k = 1
<strong>输出：</strong>4
<strong>解释：</strong>在此示例中，为了最大化得到的分割数量，可以将 s[1] 改为 'a'。
s 变为 "xayz"。
按照以下方式执行操作，直到 s 变为空：
- 选择最长且至多包含 1 个不同字符的前缀，"<em><strong>x</strong></em>ayz"。
- 删除该前缀，s 变为 "ayz"。现在分割数量为 1。
- 选择最长且至多包含 1 个不同字符的前缀，"<em><strong>a</strong></em>yz"。
- 删除该前缀，s 变为 "yz"，现在分割数量为 2。
- 选择最长且至多包含 1 个不同字符的前缀，"<em><strong>y</strong></em>z"。
- 删除该前缀，s 变为 "z"。现在分割数量为 3。
- 选择最且至多包含 1 个不同字符的前缀，"<em>z</em>"。
- 删除该前缀，s 变为空。现在分割数量为 4。
因此，答案是 4。
可以证明，分割数量不可能超过 4。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>4</sup></code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
	<li><code>1 &lt;= k &lt;= 26</code></li>
</ul>

#### 思路

定义 $\textit{dfs}(i,\textit{mask}, \textit{changed})$ 表示当前遍历到 $s[i]$，
当前这一段的字符集合是 $\textit{mask}$，是否已经修改了字符（$\textit{changed}$）， 后续可以得到的最大分割数。
分类讨论：
- 如果不改 $s[i]$：
  - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $s[i]$ 必须划分到下一段子串中。答案为 $\textit{dfs}(i+1, \{s[i]\},\textit{changed}) + 1$。
  - 如果 $s[i]$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $s[i]$ 必须在当前这一段中。答案为 $\textit{dfs}(i+1, \textit{mask}\cup \{s[i]\},\textit{changed})$。
-如果 $\textit{changed}=\texttt{false}$，那么可以改 $s[i]$，枚举改成第 $j$ 个字母。
  - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小超过了 $k$，那么 $j$ 必须划分到下一段子串中。答案为 $\textit{dfs}(i+1, \{j\},\texttt{true}) + 1$。
  - 如果 $j$ 加到 $\textit{mask}$ 后，集合的大小没有超过 $k$，那么 $j$ 必须在当前这一段中。答案为 $\textit{dfs}(i+1, \textit{mask}\cup \{j\},\texttt{true})$。

这些情况取最大值，就得到了 $\textit{dfs}(i,\textit{mask}, \textit{changed})$。\n\n递归边界：$\textit{dfs}(n,*,*) = 1$。注意当 $i>0$ 时，$\textit{mask}\ne 0$，表示一段子串的字符集合。
所以递归到 $i=n$ 时，$\textit{mask}$ 就是最后一段的字符集合了，返回 $1$。递归入口：$\textit{dfs}(0,0,\texttt{false})$，也就是答案。

```go [sol]
func maxPartitionsAfterOperations(s string, k int) int {
	n := len(s)
	type args struct {
		i, mask int
		changed bool
	}
	memo := map[args]int{}
	var dfs func(int, int, bool) int
	dfs = func(i, mask int, changed bool) (res int) {
		if i == n {
			return 1
		}

		a := args{i, mask, changed}
		if v, ok := memo[a]; ok { // 之前计算过
			return v
		}

		// 不改 s[i]
		bit := 1 << (s[i] - 'a')
		newMask := mask | bit
		if bits.OnesCount(uint(newMask)) > k {
			// 分割出一个子串，这个子串的最后一个字母在 i-1
			// s[i] 作为下一段的第一个字母，也就是 bit 作为下一段的 mask 的初始值
			res = dfs(i+1, bit, changed) + 1
		} else { // 不分割
			res = dfs(i+1, newMask, changed)
		}

		if !changed {
			// 枚举把 s[i] 改成 a,b,c,...,z
			for j := 0; j < 26; j++ {
				newMask := mask | 1<<j
				if bits.OnesCount(uint(newMask)) > k {
					// 分割出一个子串，这个子串的最后一个字母在 i-1
					// j 作为下一段的第一个字母，也就是 1<<j 作为下一段的 mask 的初始值
					res = max(res, dfs(i+1, 1<<j, true)+1)
				} else { // 不分割
					res = max(res, dfs(i+1, newMask, true))
				}
			}
		}

		memo[a] = res // 记忆化
		return res
	}
	return dfs(0, 0, false)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。 由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。分类讨论：
  - 如果 $\textit{mask}$ 之前没有修改，这样的状态有 $\mathcal{O}(n|\Sigma|)$ 个，单个状态的计算时间为 $\mathcal{O}(|\Sigma|)$，即枚举修改的时间。
  - 如果 $\textit{mask}$ 之前有修改，这样的状态有 $\mathcal{O}(n|\Sigma|^2)$ 个，单个状态的计算时间为 $\mathcal{O}(1)$，只能不修改。
  - 所以时间复杂度为 $\mathcal{O}(n|\Sigma|^2)$。
- 空间复杂度：$\mathcal{O}(n|\Sigma|^2)$。
