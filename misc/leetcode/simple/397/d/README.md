#### 题目

<p>给你一个数组 <code>nums</code> ，它是 <code>[0, 1, 2, ..., n - 1]</code> 的一个<span data-keyword="permutation">排列</span> 。对于任意一个 <code>[0, 1, 2, ..., n - 1]</code> 的排列 <code>perm</code> ，其 <strong>分数 </strong>定义为：</p>

<p><code>score(perm) = |perm[0] - nums[perm[1]]| + |perm[1] - nums[perm[2]]| + ... + |perm[n - 1] - nums[perm[0]]|</code></p>

<p>返回具有<strong> 最低</strong> 分数的排列 <code>perm</code> 。如果存在多个满足题意且分数相等的排列，则返回其中<span data-keyword="lexicographically-smaller-array">字典序最小</span>的一个。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [1,0,2]</span></p>

<p><strong>输出：</strong><span class="example-io">[0,1,2]</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/04/04/example0gif.gif" style="width: 235px; height: 235px;" /></strong></p>

<p>字典序最小且分数最低的排列是 <code>[0,1,2]</code>。这个排列的分数是 <code>|0 - 0| + |1 - 2| + |2 - 1| = 2</code> 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [0,2,1]</span></p>

<p><strong>输出：</strong><span class="example-io">[0,2,1]</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/04/04/example1gif.gif" style="width: 235px; height: 235px;" /></strong></p>

<p>字典序最小且分数最低的排列是 <code>[0,2,1]</code>。这个排列的分数是 <code>|0 - 1| + |2 - 2| + |1 - 0| = 2</code> 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n == nums.length &lt;= 14</code></li>
	<li><code>nums</code> 是 <code>[0, 1, 2, ..., n - 1]</code> 的一个排列。</li>
</ul>


#### 思路

为方便描述，将 $\textit{nums}$ 数组简记为 $a$，将下标 $[0,n-1]$ 的排列 $\textit{perm}$ 简记为 $p$。

例如 $p=[0,2,1]$，对于分数的计算式子，要相减的两个数的对应关系是：

- $p_0$ 对应 $a_{p_1}$。
- $p_1$ 对应 $a_{p_2}$。
- $p_2$ 对应 $a_{p_0}$。

把 $p$ 循环左移一位，得到 $p'=[2,1,0]$，我们在计算得分时，要相减的两个数的对应关系是：

- $p'_0$ 对应 $a_{p'_1}$，即原来的 $p_1$ 对应 $a_{p_2}$。
- $p'_1$ 对应 $a_{p'_2}$，即原来的 $p_2$ 对应 $a_{p_0}$。
- $p'_2$ 对应 $a_{p'_0}$，即原来的 $p_0$ 对应 $a_{p_1}$。

可以发现，在循环左移后，对应关系是不变的。换句话说，**循环左移后，分数是不变的**。

**推论**：题目要计算的字典序最小的排列 $p$，一定满足 $p_0 = 0$。

**证明**：反证法，假设 $p_0 > 0$，我们可以将 $p$ 不断循环左移，直到 $p_0 =0$。由于循环左移后分数不变，我们得到了一个字典序更小的答案，矛盾，故原命题成立。

## 一、寻找子问题

假设 $n=6$。

回顾枚举 [46. 全排列](https://leetcode.cn/problems/permutations/) 的思路：

- $p_0 = 0$。
- 枚举 $p_1$ 选哪个，它不能是之前选过的数 $0$，可以是 $1,2,3,4,5$，假设 $p_1 = 2$。
- 枚举 $p_2$ 选哪个，它不能是之前选过的数 $0,2$，可以是 $1,3,4,5$，假设 $p_2 = 4$。
- 枚举 $p_3$ 选哪个，它不能是之前选过的数 $0,2,4$，可以是 $1,3,5$，假设 $p_3 = 1$。
- 枚举 $p_4$ 选哪个，它不能是之前选过的数 $0,1,2,4$，可以是 $3,5$，假设 $p_4 = 5$。
- 枚举 $p_5$ 选哪个，它不能是之前选过的数 $0,1,2,4,5$，只能是 $3$，所以 $p_5 = 3$。
- 生成的排列为 $p = [0,2,4,1,5,3]$。

在这个过程中，我们需要知道：

- 哪些数不能选。这可以用一个集合 $S$ 存储选过的数。枚举 $p_i$ 的值，然后把 $p_i$ 加入集合 $S$。注意剩下能选的数变少了，要解决的问题规模也变小了。

除此以外，为了计算分数：

- 对于 $p_i$，我们还需要知道上一个数 $p_{i-1}$ 是多少。设 $j = p_{i-1}$。

在上面的例子中，对于 $p_4$ 而言，之前选过的数的集合 $S=\{0,1,2,4\}$，上一个选的数是 $j=1$，那么枚举：

- $p_4=3$，把 $|1-a_3|$ 加入分数，接下来要解决的问题是：在 $S=\{0,1,2,3,4\},\ j=3$ 的情况下，剩余数字能得到的最低分数。
- $p_4=5$，把 $|1-a_5|$ 加入分数，接下来要解决的问题是：在 $S=\{0,1,2,4,5\},\ j=5$ 的情况下，剩余数字能得到的最低分数。

由于这些都是**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「枚举选哪个」。

请注意上述过程中，会产生**重复子问题**，例如：

- 目前生成的排列是 $p = [0,2,4,1,\text{\_},\text{\_}]$，现在递归到倒数第二个位置，那么 $S=\{0,1,2,4\},\ j=1$。
- 目前生成的排列是 $p = [0,4,2,1,\text{\_},\text{\_}]$，现在递归到倒数第二个位置，那么 $S=\{0,1,2,4\},\ j=1$。

这样的重复子问题，是本题可以用 DP 优化的关键。换句话说，前面的排列具体长啥样，我们并不关心，**我们记录的是无序的集合，不是有序的列表**。

## 二、状态定义与状态转移方程

按照上面的讨论，定义 $\textit{dfs}(S,j)$ 表示在之前选过的数的集合为 $S$，上一个选的数是 $j$ 的情况下，剩余数字能得到的最低分数。

考虑当前数字选什么：

- 枚举 $k=1,2,\cdots,n-1$ 且 $k\notin S$。注意 $0$ 一定在 $S$ 中，无需枚举。
- 要解决的问题变成：在之前选过的数的集合为 $S \cup \{k\}$，上一个选的数是 $k$ 的情况下，剩余数字能得到的最低分数。这个分数加上 $|j-a_k|$，更新 $\textit{dfs}(S,j)$ 的最小值。

即

$$
\textit{dfs}(S,j) = \min_{k} \textit{dfs}(S \cup \{k\}, k) + |j - a_k|
$$

递归边界：$\textit{dfs}(U,j)=|j - a_0|$，其中 $U=\{0,1,2,\cdots n-1\}$。

递归入口：$\textit{dfs}(\{0\}, 0)$，也就是答案。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

由于有重复子问题，整个递归过程中会有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(S,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

```
func findPermutation(a []int) []int {
	n := len(a)
	u := 1<<n - 1
	f := make([][]int, 1<<n)
	g := make([][]int, 1<<n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MaxInt
		}
		g[i] = make([]int, n)
	}
	for j := range f[u] {
		f[u][j] = abs(j - a[0])
		g[u][j] = -1
	}
	for s := u - 2; s > 0; s -= 2 { // 注意偶数不含 0，是无效状态
		for j := 0; j < n; j++ {
			if s>>j&1 == 0 { // 无效状态，因为 j 一定在 s 中
				continue
			}
			for k := 1; k < n; k++ {
				if s>>k&1 > 0 { // k 之前填过
					continue
				}
				v := f[s|1<<k][k] + abs(j-a[k])
				if v < f[s][j] {
					f[s][j] = v
					g[s][j] = k // 记录该状态下填了哪个数
				}
			}
		}
	}

	ans := make([]int, 0, n)
	for s, j := 0, 0; j >= 0; {
		ans = append(ans, j)
		s |= 1 << j
		j = g[s][j]
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2 2^n)$，其中 $n$ 为 $\textit{a}$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(n2^n)$，单个状态的计算时间为 $\mathcal{O}(n)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2 2^n)$。
- 空间复杂度：$\mathcal{O}(n2^n)$。保存多少状态，就需要多少空间。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)