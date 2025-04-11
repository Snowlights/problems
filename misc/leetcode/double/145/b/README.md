### 题目

<p>Bob 被困在了一个地窖里，他需要破解 <code>n</code>&nbsp;个锁才能逃出地窖，每一个锁都需要一定的 <strong>能量</strong>&nbsp;才能打开。每一个锁需要的能量存放在一个数组&nbsp;<code>strength</code>&nbsp;里，其中&nbsp;<code>strength[i]</code>&nbsp;表示打开第 <code>i</code>&nbsp;个锁需要的能量。</p>

<p>Bob 有一把剑，它具备以下的特征：</p>

<ul>
	<li>一开始剑的能量为 0 。</li>
	<li>剑的能量增加因子&nbsp;<code><font face="monospace">X</font></code>&nbsp;一开始的值为 1 。</li>
	<li>每分钟，剑的能量都会增加当前的&nbsp;<code>X</code>&nbsp;值。</li>
	<li>打开第 <code>i</code>&nbsp;把锁，剑的能量需要到达 <strong>至少</strong>&nbsp;<code>strength[i]</code>&nbsp;。</li>
	<li>打开一把锁以后，剑的能量会变回 0 ，<code>X</code>&nbsp;的值会增加一个给定的值 <code>K</code>&nbsp;。</li>
</ul>

<p>你的任务是打开所有 <code>n</code>&nbsp;把锁并逃出地窖，请你求出需要的 <strong>最少</strong>&nbsp;分钟数。</p>

<p>请你返回 Bob<strong>&nbsp;</strong>打开所有 <code>n</code>&nbsp;把锁需要的 <strong>最少</strong>&nbsp;时间。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>strength = [3,4,1], K = 1</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><b>解释：</b></p>

<table style="border: 1px solid black;">
	<tbody>
		<tr>
			<th style="border: 1px solid black;">时间</th>
			<th style="border: 1px solid black;">能量</th>
			<th style="border: 1px solid black;">X</th>
			<th style="border: 1px solid black;">操作</th>
			<th style="border: 1px solid black;">更新后的 X</th>
		</tr>
		<tr>
			<td style="border: 1px solid black;">0</td>
			<td style="border: 1px solid black;">0</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">什么也不做</td>
			<td style="border: 1px solid black;">1</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">打开第 3&nbsp;把锁</td>
			<td style="border: 1px solid black;">2</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">什么也不做</td>
			<td style="border: 1px solid black;">2</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">4</td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">打开第 2 把锁</td>
			<td style="border: 1px solid black;">3</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">4</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">打开第 1 把锁</td>
			<td style="border: 1px solid black;">3</td>
		</tr>
	</tbody>
</table>

<p>无法用少于 4 分钟打开所有的锁，所以答案为 4 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>strength = [2,5,4], K = 2</span></p>

<p><span class="example-io"><b>输出：</b>5</span></p>

<p><b>解释：</b></p>

<table style="border: 1px solid black;">
	<tbody>
		<tr>
			<th style="border: 1px solid black;">时间</th>
			<th style="border: 1px solid black;">能量</th>
			<th style="border: 1px solid black;">X</th>
			<th style="border: 1px solid black;">操作</th>
			<th style="border: 1px solid black;">更新后的 X</th>
		</tr>
		<tr>
			<td style="border: 1px solid black;">0</td>
			<td style="border: 1px solid black;">0</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">什么也不做</td>
			<td style="border: 1px solid black;">1</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">什么也不做</td>
			<td style="border: 1px solid black;">1</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">2</td>
			<td style="border: 1px solid black;">1</td>
			<td style="border: 1px solid black;">打开第 1 把锁</td>
			<td style="border: 1px solid black;">3</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">什么也不做</td>
			<td style="border: 1px solid black;">3</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">4</td>
			<td style="border: 1px solid black;">6</td>
			<td style="border: 1px solid black;">3</td>
			<td style="border: 1px solid black;">打开第 2 把锁</td>
			<td style="border: 1px solid black;">5</td>
		</tr>
		<tr>
			<td style="border: 1px solid black;">5</td>
			<td style="border: 1px solid black;">5</td>
			<td style="border: 1px solid black;">5</td>
			<td style="border: 1px solid black;">打开第 3 把锁</td>
			<td style="border: 1px solid black;">7</td>
		</tr>
	</tbody>
</table>

<p>无法用少于 5 分钟打开所有的锁，所以答案为 5 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == strength.length</code></li>
	<li><code>1 &lt;= n &lt;= 8</code></li>
	<li><code>1 &lt;= K &lt;= 10</code></li>
	<li><code>1 &lt;= strength[i] &lt;= 10<sup>6</sup></code></li>
</ul>

### 思路

## 方法一：排列型回溯+最优性剪枝

排列型回溯，枚举所有开锁顺序（全排列），具体见[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/)。

设当前开了 $i$ 把锁，那么 $x=1+k\cdot i$。

枚举开第 $j$ 把锁，需要的时间为 $\left\lceil\dfrac{strength[j]}{x}\right\rceil$。

最优性剪枝：由于用时 $\textit{time}$ 在递归过程中只会增大，如果发现 $\textit{time}\ge \textit{ans}$，那么直接返回，不再往下递归。

关于上取整的计算，当 $a$ 和 $b$ 均为正整数时，我们有

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
$$

讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1YeqHYSEXv/?t=3m54s)，欢迎点赞关注~

```
func findMinimumTime(strength []int, k int) int {
	ans := math.MaxInt
	n := len(strength)
	done := make([]bool, n)
	var dfs func(int, int)
	dfs = func(i, time int) {
		// 最优性剪枝：答案不可能变小
		if time >= ans {
			return
		}
		if i == n {
			ans = time
			return
		}
		x := 1 + k*i
		for j, d := range done {
			if !d {
				done[j] = true // 已开锁
				dfs(i+1, time+(strength[j]-1)/x+1)
				done[j] = false // 恢复现场
			}
		}
	}
	dfs(0, 0)
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n!)$，其中 $n$ 是 $\textit{strength}$ 的长度。证明见[【基础算法精讲 16】](https://www.bilibili.com/video/BV1mY411D7f6/)。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：状压 DP

如果你没做过状压 DP，请先完成 [526. 优美的排列](https://leetcode.cn/problems/beautiful-arrangement/)，并阅读我的题解 [教你一步步思考状压 DP：从记忆化搜索到递推](https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/)。

定义 $\textit{dfs}(i)$ 表示尚未开锁的下标集合为 $i$ 时，打开所有锁需要的最少时间。

当前的 $x = 1 + k\cdot (n-|i|)$，其中 $|i|$ 是集合 $i$ 的大小。

枚举开第 $j$ 把锁，那么问题变成：尚未开锁的下标集合为 $i\setminus\{j\}$ 时，打开所有锁需要的最少时间，即 $\textit{dfs}(i\setminus\{j\})$。

取最小值，得

$$
\textit{dfs}(i) = \min\limits_{j\in i}  \textit{dfs}(i\setminus\{j\}) + \left\lceil\dfrac{strength[j]}{x}\right\rceil
$$

递归边界：$\textit{dfs}(\varnothing) = 0$。

递归入口：$\textit{dfs}(U)$，其中全集 $U=\{0,1,2,\ldots,n-1\}$。

代码实现时，用二进制表示集合，用位运算实现集合操作，具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

### 记忆化搜索

```
func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	memo := make([]int, 1<<n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		x := 1 + k*(n-bits.OnesCount(uint(i)))
		res := math.MaxInt
		for j, s := range strength {
			if i>>j&1 > 0 {
				res = min(res, dfs(i^1<<j)+(s-1)/x+1)
			}
		}
		*p = res // 记忆化
		return res
	}
	return dfs(1<<n - 1)
}
```

### 递推

```
func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	m := 1 << n
	f := make([]int, m)
	for i := 1; i < m; i++ {
		x := 1 + k*(n-bits.OnesCount(uint(i)))
		f[i] = math.MaxInt
		for j, s := range strength {
			if i>>j&1 > 0 {
				f[i] = min(f[i], f[i^1<<j]+(s-1)/x+1)
			}
		}
	}
	return f[m-1]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n2^n)$，其中 $n$ 是 $\textit{strength}$ 的长度。
- 空间复杂度：$\mathcal{O}(2^n)$。

## 方法三：最小费用流

如果 $n=100$ 呢？

创建一个**二分图**，左部为锁的编号 $i=0,1,2,\ldots,n-1$，右部表示这个锁是第 $j=0,1,2,\ldots,n-1$ 次开的。

问题相当于计算这个二分图的**最大带权匹配**。

建图：

- 从 $i$ 向 $j$ 连边，容量为 $1$，费用为 $\left\lceil\dfrac{\textit{strength}[i]}{1 + k\cdot j}\right\rceil$。
- 从超级源点 $S=2n$ 向每个 $i$ 连边，容量为 $1$，费用 $0$。
- 从每个 $j$ 向超级汇点 $T=2n+1$ 连边，容量为 $1$，费用为 $0$。

计算从 $S$ 到 $T$ 的最小费用流，满流时的费用即为答案。

```
func findMinimumTime(strength []int, k int) int {
	n := len(strength)
	S := n * 2
	T := S + 1

	// rid 为反向边在邻接表中的下标
	type neighbor struct{ to, rid, cap, cost int }
	g := make([][]neighbor, T+1)
	addEdge := func(from, to, cap, cost int) {
		g[from] = append(g[from], neighbor{to, len(g[to]), cap, cost})
		g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0, -cost})
	}
	for i, s := range strength {
		// 枚举这个锁是第几次开的
		for j := range n {
			x := 1 + k*j
			addEdge(i, n+j, 1, (s-1)/x+1)
		}
		addEdge(S, i, 1, 0)
	}
	for i := n; i < n*2; i++ {
		addEdge(i, T, 1, 0)
	}

	// 下面是最小费用最大流模板
	dis := make([]int, len(g))
	type vi struct{ v, i int }
	fa := make([]vi, len(g))
	inQ := make([]bool, len(g))
	spfa := func() bool {
		for i := range dis {
			dis[i] = math.MaxInt
		}
		dis[S] = 0
		inQ[S] = true
		q := []int{S}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			inQ[v] = false
			for i, e := range g[v] {
				if e.cap == 0 {
					continue
				}
				w := e.to
				newD := dis[v] + e.cost
				if newD < dis[w] {
					dis[w] = newD
					fa[w] = vi{v, i}
					if !inQ[w] {
						inQ[w] = true
						q = append(q, w)
					}
				}
			}
		}
		// 循环结束后所有 inQ[v] 都为 false，无需重置
		return dis[T] < math.MaxInt
	}

	minCost := 0
	for spfa() {
		// 沿 st-end 的最短路尽量增广
		// 特别地，如果建图时所有边的容量都设为 1，那么 minF 必然为 1，下面第一个 for 循环可以省略
		minF := math.MaxInt
		for v := T; v != S; {
			p := fa[v]
			minF = min(minF, g[p.v][p.i].cap)
			v = p.v
		}
		for v := T; v != S; {
			p := fa[v]
			e := &g[p.v][p.i]
			e.cap -= minF
			g[v][e.rid].cap += minF
			v = p.v
		}
		minCost += dis[T] * minF
	}
	return minCost
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 是 $\textit{strength}$ 的长度。由于二分图的特殊性，算法跑至多 $n$ 次 $\mathcal{O}(n^2)$ 的 SPFA 就结束了。
- 空间复杂度：$\mathcal{O}(n^2)$。

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