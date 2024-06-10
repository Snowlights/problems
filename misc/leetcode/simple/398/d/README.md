#### 题目

<p>给你有一个 <strong>非负</strong>&nbsp;整数&nbsp;<code>k</code>&nbsp;。有一个无限长度的台阶，<strong>最低</strong>&nbsp;一层编号为 0 。</p>

<p>虎老师有一个整数&nbsp;<code>jump</code>&nbsp;，一开始值为 0 。虎老师从台阶 1 开始，虎老师可以使用 <strong>任意</strong>&nbsp;次操作，目标是到达第&nbsp;<code>k</code> 级台阶。假设虎老师位于台阶 <code>i</code> ，一次 <strong>操作</strong> 中，虎老师可以：</p>

<ul>
	<li>向下走一级到&nbsp;<code>i - 1</code>&nbsp;，但该操作&nbsp;<strong>不能</strong>&nbsp;连续使用，如果在台阶第 0 级也不能使用。</li>
	<li>向上走到台阶&nbsp;<code>i + 2<sup>jump</sup></code>&nbsp;处，然后&nbsp;<code>jump</code>&nbsp;变为&nbsp;<code>jump + 1</code>&nbsp;。</li>
</ul>

<p>请你返回虎老师到达台阶 <code>k</code>&nbsp;处的总方案数。</p>

<p><b>注意</b>&nbsp;，虎老师可能到达台阶 <code>k</code>&nbsp;处后，通过一些操作重新回到台阶 <code>k</code>&nbsp;处，这视为不同的方案。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>k = 0</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>2 种到达台阶 0 的方案为：</p>

<ul>
	<li>虎老师从台阶&nbsp;1 开始。
	<ul>
		<li>执行第一种操作，从台阶 1 向下走到台阶 0 。</li>
	</ul>
	</li>
	<li>虎老师从台阶 1 开始。
	<ul>
		<li>执行第一种操作，从台阶 1 向下走到台阶 0 。</li>
		<li>执行第二种操作，向上走 2<sup>0</sup>&nbsp;级台阶到台阶 1 。</li>
		<li>执行第一种操作，从台阶 1 向下走到台阶 0 。</li>
	</ul>
	</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>k = 1</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p>4 种到达台阶 1 的方案为：</p>

<ul>
	<li>虎老师从台阶 1 开始，已经到达台阶 1 。</li>
	<li>虎老师从台阶 1 开始。
	<ul>
		<li>执行第一种操作，从台阶 1 向下走到台阶 0 。</li>
		<li>执行第二种操作，向上走 2<sup>0</sup>&nbsp;级台阶到台阶 1 。</li>
	</ul>
	</li>
	<li>虎老师从台阶 1 开始。
	<ul>
		<li>执行第二种操作，向上走 2<sup>0</sup>&nbsp;级台阶到台阶 2 。</li>
		<li>执行第一种操作，向下走 1 级台阶到台阶 1 。</li>
	</ul>
	</li>
	<li>虎老师从台阶 1 开始。
	<ul>
		<li>执行第一种操作，从台阶 1 向下走到台阶 0 。</li>
		<li>执行第二种操作，向上走&nbsp;2<sup>0</sup>&nbsp;级台阶到台阶 1 。</li>
		<li>执行第一种操作，向下走 1 级台阶到台阶 0 。</li>
		<li>执行第二种操作，向上走 2<sup>1</sup>&nbsp;级台阶到台阶 2 。</li>
		<li>执行第一种操作，向下走&nbsp;1 级台阶到台阶 1 。</li>
	</ul>
	</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= k &lt;= 10<sup>9</sup></code></li>
</ul>


#### 思路

## 方法一：记忆化搜索

根据题意，我们需要三个参数来表示当前的状态：

- $i$：当前位于台阶 $i$。
- $j$：已经使用了 $j$ 次第二种操作。
- $\textit{preDown}$：上一次操作是否使用了操作一。

将其定义为 $\textit{dfs}(i,j,\textit{preDown})$，表示在该状态下，有多少种方案可以到达台阶 $k$。

枚举当前使用哪个操作：

- 使用操作一：前提是 $\textit{preDown}=\texttt{false}$ 且 $i>0$。使用操作一后，要解决的子问题是 $\textit{dfs}(i-1,j,\texttt{true})$，加入返回值中。
- 使用操作二：要解决的子问题是 $\textit{dfs}(i+2^j,j+1,\texttt{false})$，加入返回值中。
- 此外，如果 $i=k$，可以把返回值加一。

递归边界：如果 $i>k+1$，由于操作一不能连续使用，无法到达 $k$，返回 $0$。

递归入口：$\textit{dfs}(1,0,\texttt{false})$，即答案。

```
const mx = 31
var c [mx][mx]int

func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

func waysToReachStair(k int) (ans int) {
	for j := bits.Len(uint(max(k-1, 0))); 1<<j-k <= j+1; j++ {
		ans += c[j+1][1<<j-k]
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

有多少个状态？

对于 $j$ 来说，由于 $i+1\le k$ 的限制，有 $\mathcal{O}(\log k)$ 个不同的 $j$。

对于固定的 $j$ 和 $\textit{preDown}$，有多少个不同的 $i$？

考虑 $j$ 对于 $i$ 的贡献，由于 $j$ 是一点一点增大的，所以 $i$ 一定包含 $2^0,2^1,2^2,\cdots,2^{j-1}$。从 $i$ 中减掉这些数的和，剩下的 $|i-1|$ 就是使用操作一的次数（注意 $i$ 从 $1$ 开始）。由于操作一不能连续使用，所以操作一的次数不超过 $j+1$。所以对于固定的 $j$ 和 $\textit{preDown}$，有 $\mathcal{O}(j)$ 个不同的 $i$。

综上所述，总共有 $\mathcal{O}(1+2+\cdots + \log k) = \mathcal{O}(\log^2 k)$ 个状态。

- 时间复杂度：$\mathcal{O}(\log^2 k)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(\log^2 k)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(\log^2 k)$。
- 空间复杂度：$\mathcal{O}(\log^2 k)$。保存多少状态，就需要多少空间。

上面的状态个数分析，也引出了下面的组合数学做法。

## 方法二：组合数学

假设使用了 $m$ 次操作一，$j$ 次操作二，那么有

$$
1 + 2^0 + 2^1 + 2^2 + \cdots + 2^{j-1} - m = k
$$

即

$$
m =  2^j - k
$$

注意上式当 $j=0$ 时也是成立的。

由于操作一不能连续使用，我们需要在这 $j$ 次操作二前后，以及相邻两次操作二的空隙中，插入 $m$ 个操作一，所以方案数等于从 $j+1$ 个物品中选出 $m$ 个物品的方案数，即

$$
\binom {j+1} m
$$

枚举 $j$，最终答案为

$$
\sum_{j=0}^{29} \binom {j+1} m
$$

其中 $0\le m \le j+1$。根据题目的数据范围，$j$ 至多枚举到 $29$。

```
const mx = 31
var c [mx][mx]int
func init() {
	for i := 0; i < mx; i++ {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

func waysToReachStair(k int) (ans int) {
	for j := bits.Len(uint(max(k-1, 0))); 1<<j-k <= j+1; j++ {
		ans += c[j+1][1<<j-k]
	}
	return
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。预处理的时间和空间不计入。
- 空间复杂度：$\mathcal{O}(1)$。


## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)