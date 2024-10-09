#### 题目

<p>给你一个大小为 4 的整数数组 <code>a</code> 和一个大小 <strong>至少</strong>为 4 的整数数组 <code>b</code>。</p>

<p>你需要从数组 <code>b</code> 中选择四个下标 <code>i<sub>0</sub></code>, <code>i<sub>1</sub></code>, <code>i<sub>2</sub></code>, 和 <code>i<sub>3</sub></code>，并满足 <code>i<sub>0</sub> &lt; i<sub>1</sub> &lt; i<sub>2</sub> &lt; i<sub>3</sub></code>。你的得分将是 <code>a[0] * b[i<sub>0</sub>] + a[1] * b[i<sub>1</sub>] + a[2] * b[i<sub>2</sub>] + a[3] * b[i<sub>3</sub>]</code> 的值。</p>

<p>返回你能够获得的 <strong>最大 </strong>得分。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">a = [3,2,5,6], b = [2,-6,4,-5,-3,2,-7]</span></p>

<p><strong>输出：</strong> <span class="example-io">26</span></p>

<p><strong>解释：</strong><br />
选择下标 0, 1, 2 和 5。得分为 <code>3 * 2 + 2 * (-6) + 5 * 4 + 6 * 2 = 26</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">a = [-1,4,5,-2], b = [-5,-1,-3,-2,-4]</span></p>

<p><strong>输出：</strong> <span class="example-io">-1</span></p>

<p><strong>解释：</strong><br />
选择下标 0, 1, 3 和 4。得分为 <code>(-1) * (-5) + 4 * (-1) + 5 * (-2) + (-2) * (-4) = -1</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>a.length == 4</code></li>
	<li><code>4 &lt;= b.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>5</sup> &lt;= a[i], b[i] &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

## 一、寻找子问题

我们要从 $b$ 中选一个长为 $4$ 的子序列，与 $a$ 运算。

看示例 1，$a=[3,2,5,6],\ b = [2,-6,4,-5,-3,2,-7]$。

考虑从右往左选 $b$ 数字，分类讨论：

- 如果不选 $b[6]$，那么需要解决的问题为：从 $b[0]$ 到 $b[5]$ 选 $4$ 个数，与 $a[0]$ 到 $a[3]$ 运算，结果的最大值。
- 如果选 $b[6]$，那么需要解决的问题为：从 $b[0]$ 到 $b[5]$ 选 $3$ 个数，与 $a[0]$ 到 $a[2]$ 运算，结果的最大值。

由于选或不选都会把原问题变成一个**和原问题相似的、规模更小的子问题**，所以可以用**递归**解决。

> 注 1：从右往左思考，主要是为了方便把递归翻译成递推。从左往右思考也是可以的。
>
> 注 2：动态规划有「**选或不选**」和「**枚举选哪个**」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「**选或不选**」。

## 二、状态定义与状态转移方程

根据上面的讨论，我们需要在递归过程中跟踪以下信息：

- $i$：当前要考虑 $b[i]$ 选或不选。
- $j$：如果选 $b[i]$，那么与 $a[j]$ 相乘。

因此，定义状态为 $\textit{dfs}(i,j)$，表示从 $b[0]$ 到 $b[i]$ 选 $j+1$ 个数，与 $a[0]$ 到 $a[j]$ 运算，结果的最大值。

接下来，思考如何从一个状态转移到另一个状态。

考虑 $b[i]$ 数字，分类讨论：

- 如果不选 $b[i]$，那么需要解决的问题为：从 $b[0]$ 到 $b[i-1]$ 选 $j+1$ 个数，与 $a[0]$ 到 $a[j]$ 运算，结果的最大值，即 $\textit{dfs}(i-1,j)$。
- 如果选 $b[i]$，那么需要解决的问题为：从 $b[0]$ 到 $b[i-1]$ 选 $j$ 个数，与 $a[0]$ 到 $a[j-1]$ 运算，结果的最大值，即 $\textit{dfs}(i-1,j-1) + a[j]\cdot b[i]$。

这两种情况取最大值，就得到了 $\textit{dfs}(i,j)$，即

$$
\textit{dfs}(i,j) = \max(\textit{dfs}(i-1,j),\textit{dfs}(i-1,j-1) + a[j]\cdot b[i])
$$

**递归边界**：$\textit{dfs}(i,-1)=0,\ \textit{dfs}(-1,j\ge 0)=-\infty$。用 $-\infty$ 表示不合法的状态，保证 $\max$ 不会取到不合法的状态。

**递归入口**：$\textit{dfs}(n-1,3)$，也就是答案。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$ 或者 $-1$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$ 或者 $-1$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。本题由于可以算出负数，可以把初始值设为一个算不出的值 $-\infty$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

具体请看视频讲解 [动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

⚠**注意**：代码先判断了 $j<0$，然后判断 $i<0$。如果你想反过来写，要在 $i<0$ 的逻辑里面额外判断 $j<0$ 的情况。

```
func maxScore(a, b []int) int64 {
	n := len(b)
	memo := make([][4]int64, n)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = math.MinInt64 // 表示没有计算过
		}
	}
	var dfs func(int, int) int64
	dfs = func(i, j int) int64 {
		if j < 0 { // 选完了
			return 0
		}
		if i < 0 { // j >= 0，没选完
			return math.MinInt64 / 2 // 防止溢出
		}
		p := &memo[i][j]
		if *p == math.MinInt64 { // 需要计算，并记忆化
			*p = max(dfs(i-1, j), dfs(i-1, j-1)+int64(a[j])*int64(b[i]))
		}
		return *p
	}
	return dfs(n-1, 3)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m=4$ 是 $a$ 的长度，$n$ 是 $b$ 的长度。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(mn)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以总的时间复杂度为 $\mathcal{O}(mn)$。
- 空间复杂度：$\mathcal{O}(mn)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i+1][j+1]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示从 $b[0]$ 到 $b[i]$ 选 $j+1$ 个数，与 $a[0]$ 到 $a[j]$ 运算，结果的最大值。这里 $+1$ 是为了把 $\textit{dfs}(-1,j)$ 和 $\textit{dfs}(i,-1)$ 也翻译过来，这样我们可以把 $f[0][j]$ 和 $f[i][0]$ 作为初始值。

相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i+1][j+1] = \max(f[i][j+1],f[i][j] + a[j]\cdot b[i])
$$

初始值 $f[i][0]=0,\ f[0][j>0]=-\infty$，翻译自递归边界 $\textit{dfs}(i,-1)=0,\ \textit{dfs}(-1,j\ge 0)=-\infty$。

答案为 $f[n][4]$，翻译自递归入口 $\textit{dfs}(n-1,3)$。

```
func maxScore(a, b []int) int64 {
	n := len(b)
	f := make([][5]int64, n+1)
	for j := 1; j < 5; j++ {
		f[0][j] = math.MinInt64 / 2
	}
	for i, y := range b {
		for j, x := range a {
			f[i+1][j+1] = max(f[i][j+1], f[i][j]+int64(x)*int64(y))
		}
	}
	return f[n][4]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m=4$ 是 $a$ 的长度，$n$ 是 $b$ 的长度。
- 空间复杂度：$\mathcal{O}(mn)$。

## 五、空间优化

像 [0-1 背包](https://www.bilibili.com/video/BV16Y411v7Y6/) 那样，去掉第一个维度，把 $f[i+1]$ 和 $f[i]$ 保存到**同一个数组**中。

状态转移方程类似 0-1 背包，$j$ 要倒序枚举。

```
func maxScore(a, b []int) int64 {
	f := [5]int64{}
	for j := 1; j < 5; j++ {
		f[j] = math.MinInt64 / 2
	}
	for _, y := range b {
		for j := 3; j >= 0; j-- {
			f[j+1] = max(f[j+1], f[j]+int64(a[j])*int64(y))
		}
	}
	return f[4]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(mn)$，其中 $m=4$ 是 $a$ 的长度，$n$ 是 $b$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)