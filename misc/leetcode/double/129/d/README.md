### 题目

<p>给你 3 个正整数&nbsp;<code>zero</code>&nbsp;，<code>one</code>&nbsp;和&nbsp;<code>limit</code>&nbsp;。</p>

<p>一个 <span data-keyword="binary-array">二进制数组</span> <code>arr</code> 如果满足以下条件，那么我们称它是 <strong>稳定的</strong> ：</p>

<ul>
	<li>0 在&nbsp;<code>arr</code>&nbsp;中出现次数 <strong>恰好</strong>&nbsp;为<strong>&nbsp;</strong><code>zero</code>&nbsp;。</li>
	<li>1 在&nbsp;<code>arr</code>&nbsp;中出现次数 <strong>恰好</strong>&nbsp;为&nbsp;<code>one</code>&nbsp;。</li>
	<li><code>arr</code> 中每个长度超过 <code>limit</code>&nbsp;的 <span data-keyword="subarray-nonempty">子数组</span> 都 <strong>同时</strong> 包含 0 和 1 。</li>
</ul>

<p>请你返回 <strong>稳定</strong>&nbsp;二进制数组的 <em>总</em> 数目。</p>

<p>由于答案可能很大，将它对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<b>取余</b>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>zero = 1, one = 1, limit = 2</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>两个稳定的二进制数组为&nbsp;<code>[1,0]</code> 和&nbsp;<code>[0,1]</code>&nbsp;，两个数组都有一个 0 和一个 1 ，且没有子数组长度大于 2 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">zero = 1, one = 2, limit = 1</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong></p>

<p>唯一稳定的二进制数组是&nbsp;<code>[1,0,1]</code>&nbsp;。</p>

<p>二进制数组&nbsp;<code>[1,1,0]</code> 和&nbsp;<code>[0,1,1]</code>&nbsp;都有长度为 2 且元素全都相同的子数组，所以它们不稳定。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>zero = 3, one = 3, limit = 2</span></p>

<p><span class="example-io"><b>输出：</b>14</span></p>

<p><strong>解释：</strong></p>

<p>所有稳定的二进制数组包括&nbsp;<code>[0,0,1,0,1,1]</code>&nbsp;，<code>[0,0,1,1,0,1]</code>&nbsp;，<code>[0,1,0,0,1,1]</code>&nbsp;，<code>[0,1,0,1,0,1]</code>&nbsp;，<code>[0,1,0,1,1,0]</code>&nbsp;，<code>[0,1,1,0,0,1]</code>&nbsp;，<code>[0,1,1,0,1,0]</code>&nbsp;，<code>[1,0,0,1,0,1]</code>&nbsp;，<code>[1,0,0,1,1,0]</code>&nbsp;，<code>[1,0,1,0,0,1]</code>&nbsp;，<code>[1,0,1,0,1,0]</code>&nbsp;，<code>[1,0,1,1,0,0]</code>&nbsp;，<code>[1,1,0,0,1,0]</code>&nbsp;和&nbsp;<code>[1,1,0,1,0,0]</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= zero, one, limit &lt;= 1000</code></li>
</ul>

### 思路

## 方法一：记忆化搜索

看示例 3，$\textit{zero}=3,\ \textit{one}=3,\ \textit{limit}=2$。
先解释 $\textit{limit}$，意思是数组中至多有连续 $\textit{limit}$ 个 $0$，且至多有连续 $\textit{limit}$ 个 $1$。

考虑稳定数组的最后一个位置填 $0$ 还是 $1$：
- 填 $0$，问题变成剩下 $2$ 个 $0$ 和 $3$ 个 $1$ 怎么填。
- 填 $1$，问题变成剩下 $3$ 个 $0$ 和 $2$ 个 $1$ 怎么填。
- 这两个都是和原问题相似的子问题。

看上去，定义 $\textit{dfs}(i,j)$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数？但这样定义不方便计算 $\textit{limit}$ 带来的影响。
改成定义 $\textit{dfs}(i,j,k)$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数，其中第 $i+j$ 个位置要填 $k$，其中 $k$ 为 $0$ 或 $1$。

如果 $k=0$，考虑第 $i+j-1$ 个位置要填什么：
- 填 $0$，问题变成 $\textit{dfs}(i-1,j,0)$。
- 填 $1$，问题变成 $\textit{dfs}(i-1,j,1)$。

但是，$\textit{dfs}(i-1,j,0)$ 包含了「最后连续 $\textit{limit}$ 个位置填 $0$」的方案，如果在这个方案末尾再加一个 $0$，就有连续 $\textit{limit}+1$ 个 $0$ 了，这是不允许的。

$\textit{dfs}(i-1,j,0)$ 中的「最后连续 $\textit{limit}$ 个位置填 $0$」的方案有多少个？
因为 $\textit{dfs}$ 的定义是稳定数组的方案数，只包含合法方案，所以在最后连续 $\textit{limit}$ 个位置填 $0$ 的情况下，倒数第 $\textit{limit}+1$ 个位置一定要填 $1$，这有 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 种方案。对于 $\textit{dfs}(i,j,0)$ 来说，这 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 个方案就是不合法方案了，要减去，得

$$
\textit{dfs}(i,j,0) = \textit{dfs}(i-1,j,0) + \textit{dfs}(i-1,j,1) - \textit{dfs}(i-\textit{limit}-1,j,1)
$$

同理得

$$
\textit{dfs}(i,j,1) = \textit{dfs}(i,j-1,0) + \textit{dfs}(i,j-1,1) - \textit{dfs}(i,j-\textit{limit}-1,0)
$$

递归边界 1：如果 $i<0$ 或者 $j<0$，返回 $0$。也可以在计算 $\textit{dfs}(i-\textit{limit}-1,j,1)$ 前判断 $i>\textit{limit}$，在计算 $\textit{dfs}(i,j-\textit{limit}-1,0)$ 前判断 $j>\textit{limit}$。
递归边界 2：如果 $i=0$，那么当 $k=1$ 且 $j\le \textit{limit}$ 的情况下返回 $1$，否则返回 $0$；如果 $j=0$，那么当 $k=0$ 且 $i\le \textit{limit}$ 的情况下返回 $1$，否则返回 $0$。
递归入口：$\textit{dfs}(\textit{zero},\textit{one},0)+\textit{dfs}(\textit{zero},\textit{one},1)$，即为答案。

```
func numberOfStableArrays(zero, one, limit int) int {
	const mod int = 1e9 + 7
	dp := make([][][]int, zero+1)
	for i := range dp {
		dp[i] = make([][]int, one+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(z, o, pre int) int
	dfs = func(z, o, pre int) int {
		if z == 0 {
			if pre == 1 && o <= limit {
				return 1
			}
			return 0
		}
		if o == 0 {
			if pre == 0 && z <= limit {
				return 1
			}
			return 0
		}

		res := 0
		dv := &dp[z][o][pre]
		if *dv >= 0 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
		if pre == 0 {
			res = (dfs(z-1, o, 0) + dfs(z-1, o, 1) + mod) % mod
			if z > limit {
				res = (res - dfs(z-limit-1, o, 1) + mod) % mod
			}
		} else {
			res = (dfs(z, o-1, 0) + dfs(z, o-1, 1) + mod) % mod
			if o > limit {
				res = (res - dfs(z, o-limit-1, 0) + mod) % mod
			}
		}
		return res
	}

	return (dfs(zero, one, 1) + dfs(zero, one, 0) + mod) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\textit{zero}\cdot \textit{one})$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\textit{zero}\cdot \textit{one})$。
- 空间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。有多少个状态，$\textit{memo}$ 数组的大小就是多少。


## 方法二：递推

定义 $f[i][j][k]$ 表示用 $i$ 个 $0$ 和 $j$ 个 $1$ 构造稳定数组的方案数，其中第 $i+j$ 个位置要填 $k$，其中 $k$ 为 $0$ 或 $1$。
状态转移方程：

$$
\begin{aligned}
&f[i][j][0] = f[i-1][j][0] + f[i-1][j][1] - f[i-\textit{limit}-1][j][1]\\
&f[i][j][1] = f[i][j-1][0] + f[i][j-1][1] - f[i][j-\textit{limit}-1][0]
\end{aligned}
$$

如果 $i\le \textit{limit}$ 则 $f[i-\textit{limit}-1][j][1]$ 视作 $0$，如果 $j\le \textit{limit}$ 则 $f[i][j-\textit{limit}-1][0]$ 视作 $0$。
初始值：$f[i][0][0] = f[0][j][1] = 1$，其中 $1\le i \le \min(\textit{limit}, \textit{zero}),\ 1\le j \le \min(\textit{limit}, \textit{one})$。翻译自递归边界。
答案：$f[\textit{zero}][\textit{one}][0] + f[\textit{zero}][\textit{one}][1]$。翻译自递归入口。

```
func numberOfStableArrays(zero, one, limit int) (ans int) {
	const mod = 1_000_000_007
	f := make([][][2]int, zero+1)
	for i := range f {
		f[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(limit, zero); i++ {
		f[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		f[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
			if i > limit {
				// + mod 保证答案非负
				f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
			}
			f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
			if j > limit {
				f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
			}
		}
	}
	return (f[zero][one][0] + f[zero][one][1]) % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\textit{zero}\cdot \textit{one})$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\textit{zero}\cdot \textit{one})$。
- 空间复杂度：$\mathcal{O}(\textit{zero}\cdot \textit{one})$。有多少个状态，$\textit{memo}$ 数组的大小就是多少。


## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)