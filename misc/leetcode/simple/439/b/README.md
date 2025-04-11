#### 题目

<p>给你一个字符串 <code>s</code> 和一个整数 <code>k</code>。</p>

<p>在一次操作中，你可以将任意位置的字符替换为字母表中相邻的字符（字母表是循环的，因此&nbsp;<code>'z'</code>&nbsp;的下一个字母是&nbsp;<code>'a'</code>）。例如，将 <code>'a'</code> 替换为下一个字母结果是 <code>'b'</code>，将 <code>'a'</code> 替换为上一个字母结果是 <code>'z'</code>；同样，将 <code>'z'</code> 替换为下一个字母结果是 <code>'a'</code>，替换为上一个字母结果是 <code>'y'</code>。</p>

<p>返回在进行&nbsp;<strong>最多</strong> <code>k</code> 次操作后，<code>s</code> 的&nbsp;<strong>最长回文子序列&nbsp;</strong>的长度。</p>

<p><strong>子序列&nbsp;</strong>是一个&nbsp;<strong>非空&nbsp;</strong>字符串，可以通过删除原字符串中的某些字符（或不删除任何字符）并保持剩余字符的相对顺序得到。</p>

<p><strong>回文&nbsp;</strong>是正着读和反着读都相同的字符串。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">s = "abced", k = 2</span></p>

<p><strong>输出:</strong> <span class="example-io">3</span></p>

<p><strong>解释:</strong></p>

<ul>
	<li>将 <code>s[1]</code> 替换为下一个字母，得到 <code>"acced"</code>。</li>
	<li>将 <code>s[4]</code> 替换为上一个字母，得到 <code>"accec"</code>。</li>
</ul>

<p>子序列 <code>"ccc"</code> 形成一个长度为 3 的回文，这是最长的回文子序列。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">s = "aaazzz", k = 4</span></p>

<p><strong>输出:</strong> 6</p>

<p><strong>解释:</strong></p>

<ul>
	<li>将 <code>s[0]</code> 替换为上一个字母，得到 <code>"zaazzz"</code>。</li>
	<li>将 <code>s[4]</code> 替换为下一个字母，得到 <code>"zaazaz"</code>。</li>
	<li>将 <code>s[3]</code> 替换为下一个字母，得到 <code>"zaaaaz"</code>。</li>
</ul>

<p>整个字符串形成一个长度为 6 的回文。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 200</code></li>
	<li><code>1 &lt;= k &lt;= 200</code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

**前置题目**：[516. 最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence/)，请看 [区间 DP【基础算法精讲 22】](https://www.bilibili.com/video/BV1Gs4y1E7EU/)。

本题在 516 题的基础上，增加一个参数 $k$，表示剩余操作次数。

## 写法一：记忆化搜索

定义 $\textit{dfs}(i,j,k)$ 表示执行最多 $k$ 次操作后，子串 $s[i]$ 到 $s[j]$ 中的最长回文子序列的长度。

和 516 题一样，分类讨论：

- 不选 $s[i]$，那么问题变成执行最多 $k$ 次操作后，子串 $s[i+1]$ 到 $s[j]$ 中的最长回文子序列的长度，即 $\textit{dfs}(i+1,j,k)$。
- 不选 $s[j]$，那么问题变成执行最多 $k$ 次操作后，子串 $s[i]$ 到 $s[j-1]$ 中的最长回文子序列的长度，即 $\textit{dfs}(i,j-1,k)$。
- 操作，把 $s[i]$ 和 $s[j]$ 都变成一样的，设操作 $\textit{op}$ 次且 $\textit{op}\le k$，那么问题变成执行最多 $k-\textit{op}$ 次操作后，子串 $s[i+1]$ 到 $s[j-1]$ 中的最长回文子序列的长度，即 $\textit{dfs}(i+1,j-1,k-\textit{op})+2$。其中 $+2$ 是因为 $s[i]$ 和 $s[j]$ 变成相同的了。

$\textit{op}$ 怎么算？设 $d = |s[i]-s[j]|$，那么 $\textit{op}=\min(d,26-d)$，前者没有跨过 $\texttt{z}$ 到 $\texttt{a}$，后者跨过。

三种情况取最大值，得

$$
\textit{dfs}(i,j,k) = \max(\textit{dfs}(i+1,j,k), \textit{dfs}(i,j-1,k), \textit{dfs}(i+1,j-1,k-\textit{op})+2)
$$

**递归边界**：$\textit{dfs}(i+1,i,k) = 0,\ \textit{dfs}(i,i,k) = 1$。注意只有一个字母的子串一定是回文串（回文子序列）。

**递归入口**：$\textit{dfs}(0,n-1,k)$，即答案。

```
func longestPalindromicSubsequence(s string, k int) int {
	n := len(s)
	memo := make([][][]int, n)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, k+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i >= j {
			return j - i + 1 // i=j+1 时返回 0，i=j 时返回 1
		}
		p := &memo[i][j][k]
		if *p != -1 {
			return *p
		}
		res := max(dfs(i+1, j, k), dfs(i, j-1, k))
		d := abs(int(s[i]) - int(s[j]))
		op := min(d, 26-d)
		if op <= k {
			res = max(res, dfs(i+1, j-1, k-op)+2)
		}
		*p = res
		return res
	}
	return dfs(0, n-1, k)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2k)$，其中 $n$ 是 $s$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n^2k)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(n^2k)$。
- 空间复杂度：$\mathcal{O}(n^2k)$。保存多少状态，就需要多少空间。

## 写法二：递推

和 516 题一样，1:1 翻译成递推。

为了提高访问缓存的效率，把 $k$ 放到第一个维度。这样第一个维度我们只会访问 $f[k-13],f[k-12],\ldots,f[k]$，减少 cache miss。

```
func longestPalindromicSubsequence(s string, K int) int {
	n := len(s)
	f := make([][][]int, K+1)
	for k := range f {
		f[k] = make([][]int, n)
		for j := range f[k] {
			f[k][j] = make([]int, n)
		}
		for i := n - 1; i >= 0; i-- {
			f[k][i][i] = 1
			for j := i + 1; j < n; j++ {
				res := max(f[k][i+1][j], f[k][i][j-1])
				d := abs(int(s[i]) - int(s[j]))
				op := min(d, 26-d)
				if op <= k {
					res = max(res, f[k-op][i+1][j-1]+2)
				}
				f[k][i][j] = res
			}
		}
	}
	return f[K][0][n-1]
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2k)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2k)$。

**注**：$f$ 的第一维度可以用滚动数组优化，这样空间复杂度可以优化到 $\mathcal{O}(n^2|\Sigma|)$，其中 $|\Sigma|=26$ 是字符集合的大小。


## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)