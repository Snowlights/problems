#### 题目

<p>给你一个 <code>m x n</code> 的网格。一个机器人从网格的左上角 <code>(0, 0)</code> 出发，目标是到达网格的右下角 <code>(m - 1, n - 1)</code>。在任意时刻，机器人只能向右或向下移动。</p>

<p>网格中的每个单元格包含一个值 <code>coins[i][j]</code>：</p>

<ul>
	<li>如果 <code>coins[i][j] &gt;= 0</code>，机器人可以获得该单元格的金币。</li>
	<li>如果 <code>coins[i][j] &lt; 0</code>，机器人会遇到一个强盗，强盗会抢走该单元格数值的&nbsp;<strong>绝对值</strong> 的金币。</li>
</ul>

<p>机器人有一项特殊能力，可以在行程中&nbsp;<strong>最多感化&nbsp;</strong>2个单元格的强盗，从而防止这些单元格的金币被抢走。</p>

<p><strong>注意：</strong>机器人的总金币数可以是负数。</p>

<p>返回机器人在路径上可以获得的&nbsp;<strong>最大金币数&nbsp;</strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">coins = [[0,1,-1],[1,-2,3],[2,-3,4]]</span></p>

<p><strong>输出：</strong> <span class="example-io">8</span></p>

<p><strong>解释：</strong></p>

<p>一个获得最多金币的最优路径如下：</p>

<ol>
	<li>从 <code>(0, 0)</code> 出发，初始金币为 <code>0</code>（总金币 = <code>0</code>）。</li>
	<li>移动到 <code>(0, 1)</code>，获得 <code>1</code> 枚金币（总金币 = <code>0 + 1 = 1</code>）。</li>
	<li>移动到 <code>(1, 1)</code>，遇到强盗抢走 <code>2</code> 枚金币。机器人在此处使用一次感化能力，避免被抢（总金币 = <code>1</code>）。</li>
	<li>移动到 <code>(1, 2)</code>，获得 <code>3</code> 枚金币（总金币 = <code>1 + 3 = 4</code>）。</li>
	<li>移动到 <code>(2, 2)</code>，获得 <code>4</code> 枚金币（总金币 = <code>4 + 4 = 8</code>）。</li>
</ol>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">coins = [[10,10,10],[10,10,10]]</span></p>

<p><strong>输出：</strong> <span class="example-io">40</span></p>

<p><strong>解释：</strong></p>

<p>一个获得最多金币的最优路径如下：</p>

<ol>
	<li>从 <code>(0, 0)</code> 出发，初始金币为 <code>10</code>（总金币 = <code>10</code>）。</li>
	<li>移动到 <code>(0, 1)</code>，获得 <code>10</code> 枚金币（总金币 = <code>10 + 10 = 20</code>）。</li>
	<li>移动到 <code>(0, 2)</code>，再获得 <code>10</code> 枚金币（总金币 = <code>20 + 10 = 30</code>）。</li>
	<li>移动到 <code>(1, 2)</code>，获得 <code>10</code> 枚金币（总金币 = <code>30 + 10 = 40</code>）。</li>
</ol>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>m == coins.length</code></li>
	<li><code>n == coins[i].length</code></li>
	<li><code>1 &lt;= m, n &lt;= 500</code></li>
	<li><code>-1000 &lt;= coins[i][j] &lt;= 1000</code></li>
</ul>

#### 思路

请先完成不允许感化的版本：[LCR 166. 珠宝的最高价值](https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/)，[讲解](https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/solution/jiao-ni-yi-bu-bu-si-kao-dpcong-hui-su-da-epvl/)。

本题相当于可以不选路径上的至多 $2$ 个数。
**多一个约束，就多一个参数。**
额外增加一个参数 $k$，定义 $\textit{dfs}(i,j,k)$ 表示从 $(i,j)$ 走到 $(0,0)$，在剩余不选次数为 $k$ 的情况下，可以获得的最大金币数。
用「选或不选」分类讨论：
- 选：$\textit{dfs}(i,j,k) = \max(\textit{dfs}(i - 1, j, k), \textit{dfs}(i, j - 1, k)) + \textit{coins}[i][j]$。
- 不选：如果 $k>0$ 且 $\textit{coins}[i][j]<0$，则可以不选，$\textit{dfs}(i,j,k) = \max(\textit{dfs}(i - 1, j, k-1), \textit{dfs}(i, j - 1, k-1))$。

两种情况取最大值。

**递归边界**：
- $\textit{dfs}(-1,j,k)=\textit{dfs}(i,-1,k)=-\infty$。用 $-\infty$ 表示不合法的状态，从而保证 $\max$ 不会取到不合法的状态。
- $\textit{dfs}(0,0,0)=\textit{coins}[0][0]$。
- $\textit{dfs}(0,0,k>0)=\max(\textit{coins}[0][0],0)$。

**递归入口**：$\textit{dfs}(m-1,n-1,2)$，这是原问题，也是答案。
⚠**注意**：由于答案可能是负数，所以记忆化数组的初始值不能是 $-1$，可以初始化成 $-\infty$。

```
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	memo := make([][][3]int, m)
	for i := range memo {
		memo[i] = make([][3]int, n)
		for j := range memo[i] {
			for k := range memo[i][j] {
				memo[i][j][k] = math.MinInt
			}
		}
	}
	var dfs func(int, int, int) int
	dfs = func(i, j, k int) int {
		if i < 0 || j < 0 {
			return math.MinInt
		}
		x := coins[i][j]
		if i == 0 && j == 0 {
			if k == 0 {
				return x
			}
			return max(x, 0)
		}
		p := &memo[i][j][k]
		if *p != math.MinInt { // 之前计算过
			return *p
		}
		res := max(dfs(i-1, j, k), dfs(i, j-1, k)) + x // 选
		if x < 0 && k > 0 {
			res = max(res, dfs(i-1, j, k-1), dfs(i, j-1, k-1)) // 不选
		}
		*p = res // 记忆化
		return res
	}
	return dfs(m-1, n-1, 2)
}
```

``` 
func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	f := make([][][3]int, m+1)
	for i := range f {
		f[i] = make([][3]int, n+1)
	}
	for j := range f[0] {
		f[0][j] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
	}
	for i, row := range coins {
		f[i+1][0] = [3]int{math.MinInt / 2, math.MinInt / 2, math.MinInt / 2}
		for j, x := range row {
			if i == 0 && j == 0 {
				f[1][1][0] = x
				f[1][1][1] = max(x, 0)
				f[1][1][2] = max(x, 0)
				continue
			}
			f[i+1][j+1][0] = max(f[i+1][j][0], f[i][j+1][0]) + x
			f[i+1][j+1][1] = max(f[i+1][j][1]+x, f[i][j+1][1]+x, f[i+1][j][0], f[i][j+1][0])
			f[i+1][j+1][2] = max(f[i+1][j][2]+x, f[i][j+1][2]+x, f[i+1][j][1], f[i][j+1][1])
		}
	}
	return f[m][n][2]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m$ 和 $n$ 分别为 $\textit{coins}$ 的行数和列数。
- 空间复杂度：$\mathcal{O}(mn)$。


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 思路

如果固定子串的左端点，那么**子串越长，字典序越大**。
单个子串的长度不能超过多少？
由于其余 $k-1$ 个子串必须是非空的，取长度为 $1$，其余子串的长度之和**至少**为 $k-1$。
所以单个子串的长度**至多**为 $n-(k-1)$。
暴力枚举子串的左端点，计算满足长度上限的最大子串。
注意特判 $k=1$ 的情况，此时无法分割，答案就是 $s$。


```
func answerString(s string, k int) (ans string) {
	if k == 1 {
		return s
	}
	n := len(s)
	for i := range n {
		ans = max(ans, s[i:min(i+n-k+1, n)])
	}
	return
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n-k))$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n-k)$ 或 $\mathcal{O}(1)$。Go 的切片不会发生拷贝，只需要 $\mathcal{O}(1)$ 额外空间。

注：利用**字符串哈希+二分长度**或者**后缀数组**，可以快速比较 $s$ 任意两个子串的大小，做到 $\mathcal{O}(n\log (n-k))$ 或者 $\mathcal{O}(n\log n)$。


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