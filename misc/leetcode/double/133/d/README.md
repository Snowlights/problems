### 题目

<p>给你一个整数&nbsp;<code>n</code>&nbsp;和一个二维数组&nbsp;<code>requirements</code>&nbsp;，其中&nbsp;<code>requirements[i] = [end<sub>i</sub>, cnt<sub>i</sub>]</code> <span class="text-only" data-eleid="10" style="white-space: pre;">表示这个要求中的末尾下标和 <strong>逆序对</strong> 的数目。</span></p>

<p>整数数组 <code>nums</code>&nbsp;中一个下标对&nbsp;<code>(i, j)</code>&nbsp;如果满足以下条件，那么它们被称为一个 <strong>逆序对</strong>&nbsp;：</p>

<ul>
	<li><code>i &lt; j</code>&nbsp;且&nbsp;<code>nums[i] &gt; nums[j]</code></li>
</ul>

<p>请你返回&nbsp;<code>[0, 1, 2, ..., n - 1]</code>&nbsp;的&nbsp;<span data-keyword="permutation">排列</span> <code>perm</code>&nbsp;的数目，满足对 <strong>所有</strong>&nbsp;的&nbsp;<code>requirements[i]</code>&nbsp;都有&nbsp;<code>perm[0..end<sub>i</sub>]</code>&nbsp;恰好有&nbsp;<code>cnt<sub>i</sub></code>&nbsp;个逆序对。</p>

<p>由于答案可能会很大，将它对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;<strong>取余</strong>&nbsp;后返回。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 3, requirements = [[2,2],[0,0]]</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>两个排列为：</p>

<ul>
	<li><code>[2, 0, 1]</code>
	<ul>
		<li>前缀&nbsp;<code>[2, 0, 1]</code>&nbsp;的逆序对为&nbsp;<code>(0, 1)</code> 和&nbsp;<code>(0, 2)</code>&nbsp;。</li>
		<li>前缀&nbsp;<code>[2]</code>&nbsp;的逆序对数目为 0 个。</li>
	</ul>
	</li>
	<li><code>[1, 2, 0]</code>
	<ul>
		<li>前缀&nbsp;<code>[1, 2, 0]</code>&nbsp;的逆序对为&nbsp;<code>(0, 2)</code> 和&nbsp;<code>(1, 2)</code>&nbsp;。</li>
		<li>前缀&nbsp;<code>[1]</code>&nbsp;的逆序对数目为 0 个。</li>
	</ul>
	</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 3, requirements = [[2,2],[1,1],[0,0]]</span></p>

<p><b>输出：</b>1</p>

<p><strong>解释：</strong></p>

<p>唯一满足要求的排列是&nbsp;<code>[2, 0, 1]</code>&nbsp;：</p>

<ul>
	<li>前缀&nbsp;<code>[2, 0, 1]</code>&nbsp;的逆序对为&nbsp;<code>(0, 1)</code> 和&nbsp;<code>(0, 2)</code>&nbsp;。</li>
	<li>前缀&nbsp;<code>[2, 0]</code>&nbsp;的逆序对为&nbsp;<code>(0, 1)</code>&nbsp;。</li>
	<li>前缀&nbsp;<code>[2]</code>&nbsp;的逆序对数目为 0 。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 2, requirements = [[0,0],[1,0]]</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><b>解释：</b></p>

<p>唯一满足要求的排列为&nbsp;<code>[0, 1]</code>&nbsp;：</p>

<ul>
	<li>前缀&nbsp;<code>[0]</code>&nbsp;的逆序对数目为 0 。</li>
	<li>前缀&nbsp;<code>[0, 1]</code>&nbsp;的逆序对为&nbsp;<code>(0, 1)</code>&nbsp;。</li>
</ul>
</div>

<div id="gtx-trans" style="position: absolute; left: 198px; top: 756px;">
<div class="gtx-trans-icon">&nbsp;</div>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 300</code></li>
	<li><code>1 &lt;= requirements.length &lt;= n</code></li>
	<li><code>requirements[i] = [end<sub>i</sub>, cnt<sub>i</sub>]</code></li>
	<li><code>0 &lt;= end<sub>i</sub> &lt;= n - 1</code></li>
	<li><code>0 &lt;= cnt<sub>i</sub> &lt;= 400</code></li>
	<li>输入保证至少有一个&nbsp;<code>i</code>&nbsp;满足&nbsp;<code>end<sub>i</sub> == n - 1</code>&nbsp;。</li>
	<li>输入保证所有的&nbsp;<code>end<sub>i</sub></code>&nbsp;互不相同。</li>
</ul>


### 思路

## 一、寻找子问题

看示例 1，要求构造逆序对为 $2$ 的排列。

讨论最后一个数填什么：

- 填 $2$，那么前面不会有比 $2$ 大的数，这意味着剩下的 $n-1=2$ 个数的逆序对等于 $2-0=2$。
- 填 $1$，那么前面一定有 $1$ 个比 $1$ 大的数，这意味着剩下的 $n-1=2$ 个数的逆序对等于 $2-1=1$。
- 填 $0$，那么前面一定有 $2$ 个比 $0$ 大的数，这意味着剩下的 $n-1=2$ 个数的逆序对等于 $2-2=0$。

这些问题都是**和原问题相似的、规模更小的子问题**，可以用**递归**解决。

> 注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「枚举选哪个」。

## 二、状态定义与状态转移方程

因为要解决的问题都形如「前 $i$ 个数的逆序对为 $j$ 时的排列个数」，所以用它作为本题的状态定义 $\textit{dfs}(i,j)$。

直接考虑第 $i$ 个数 $\textit{perm}[i]$ 和前面 $\textit{perm}[0]$ 到 $\textit{perm}[i-1]$ 可以组成的逆序对的个数：

- 前面有 $k$ 个数比 $\textit{perm}[i]$ 大，组成了 $k$ 个逆序对，那么问题变成前 $i-1$ 个数的逆序对为 $j-k$ 时的排列个数，即 $\textit{dfs}(i-1,j-k)$。

累加得

$$
\textit{dfs}(i,j) = \sum_{k=0}^{\min(i,j)}\textit{dfs}(i-1,j-k)
$$

其中 $\min(i,j)$ 是因为前面只有 $i$ 个数，至多和 $\textit{perm}[i]$ 组成 $i$ 个逆序对。

⚠**注意**：我们**不需要知道每个位置具体填了什么数**。无论之前填了什么数，只要 $\textit{perm}[i]$ 填的是剩余元素的**最大值**，那么 $k$ 就是 $0$；只要 $\textit{perm}[i]$ 填的是剩余元素的**次大值**，那么 $k$ 就是 $1$；依此类推。

除此以外，设 $\textit{req}[i]$ 是前 $i$ 个数的逆序对个数（没有要求就是 $-1$），如果 $\textit{req}[i-1]\ge 0$，则无需枚举 $k$，分类讨论：

- 如果 $j<\textit{req}[i-1]$ 或者 $j-i>\textit{req}[i-1]$，则无法满足，$\textit{dfs}(i,j) = 0$。
- 否则 $\textit{dfs}(i,j) = \textit{dfs}(i-1,\textit{req}[i-1])$。

**递归边界**：$\textit{dfs}(0,0)=1$，此时找到了一个符合要求的排列。

**递归入口**：$\textit{dfs}(n-1,\textit{req}[n-1])$，也就是答案。

根据题意，$\textit{req}[0]$ 一定为 $0$。代码实现时，可以在递归之前判断 $\textit{req}[0] > 0$ 的情况，如果满足则直接返回 $0$。

## 三、递归搜索 + 保存递归返回值 = 记忆化搜索

考虑到整个递归过程中有大量重复递归调用（递归入参相同）。由于递归函数没有副作用，同样的入参无论计算多少次，算出来的结果都是一样的，因此可以用**记忆化搜索**来优化：

- 如果一个状态（递归入参）是第一次遇到，那么可以在返回前，把状态及其结果记到一个 $\textit{memo}$ 数组中。
- 如果一个状态不是第一次遇到（$\textit{memo}$ 中保存的结果不等于 $\textit{memo}$ 的初始值），那么可以直接返回 $\textit{memo}$ 中保存的结果。

**注意**：$\textit{memo}$ 数组的**初始值**一定不能等于要记忆化的值！例如初始值设置为 $0$，并且要记忆化的 $\textit{dfs}(i,j)$ 也等于 $0$，那就没法判断 $0$ 到底表示第一次遇到这个状态，还是表示之前遇到过了，从而导致记忆化失效。一般把初始值设置为 $-1$。

> Python 用户可以无视上面这段，直接用 `@cache` 装饰器。

```
func numberOfPermutations(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, m+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		defer func() { *p = res }()
		if r := req[i-1]; r >= 0 {
			if j < r || j-i > r {
				return 0
			}
			return dfs(i-1, r)
		}
		for k := 0; k <= min(i, j); k++ {
			res += dfs(i-1, j-k)
		}
		return res % mod
	}
	return dfs(n-1, req[n-1])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\min(n,m))$，其中 $m=\max(\textit{cnt}_i)$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数等于 $\mathcal{O}(nm)$，单个状态的计算时间为 $\mathcal{O}(\min(n,m))$，所以动态规划的时间复杂度为 $\mathcal{O}(nm\min(n,m))$。
- 空间复杂度：$\mathcal{O}(nm)$。保存多少状态，就需要多少空间。

## 四、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。

具体来说，$f[i][j]$ 的定义和 $\textit{dfs}(i,j)$ 的定义是一样的，都表示前 $i$ 个数的逆序对为 $j$ 时的排列个数。

如果 $\textit{req}[i-1]<0$，相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i][j] = \sum_{k=0}^{\min(i,j)}f[i-1][j-k]
$$

如果 $r=\textit{req}[i-1]\ge 0$，递推式为

$$
f[i][j] =
\begin{cases}
f[i-1][r],\ &r\le j \le i+r\\
0,\ &\text{otherwise}\\
\end{cases}
$$

初始值 $f[0][0]=1$，翻译自递归边界 $\textit{dfs}(0,0)=1$。

答案为 $f[n-1][\textit{req}[n-1]]$，翻译自递归入口 $\textit{dfs}(n-1,\textit{req}[n-1])$。

```
func numberOfPermutations(n int, requirements [][]int) int {
	const mod = 1_000_000_007
	req := make([]int, n)
	for i := 1; i < n; i++ {
		req[i] = -1
	}
	for _, p := range requirements {
		req[p[0]] = p[1]
	}
	if req[0] > 0 {
		return 0
	}

	m := slices.Max(req)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	f[0][0] = 1
	for i := 1; i < n; i++ {
		mx := m
		if req[i] >= 0 {
			mx = req[i]
		}
		if r := req[i-1]; r >= 0 {
			for j := r; j <= min(i+r, mx); j++ {
				f[i][j] = f[i-1][r]
			}
		} else {
			for j := 0; j <= mx; j++ {
				for k := 0; k <= min(i, j); k++ {
					f[i][j] = (f[i][j] + f[i-1][j-k]) % mod
				}
			}
		}
	}
	return f[n-1][req[n-1]]
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm\min(n,m))$，其中 $m=\max(\textit{cnt}_i)$。
- 空间复杂度：$\mathcal{O}(nm)$。

## 五、终极优化

#### 1) 时间优化

和式

$$
\sum_{k=0}^{\min(i,j)}f[i-1][j-k]
$$

可以用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 优化。

#### 2) 空间优化

观察上面的状态转移方程，在计算 $f[i]$ 时，只会用到 $f[i-1]$，不会用到比 $i-1$ 更早的状态。

因此可以去掉第一个维度，反复利用同一个长为 $m+1$ 的一维数组。

代码实现时，前缀和可以直接保存在 $f$ 中。


```
func maximumLength(nums []int, k int) int {
	fs := map[int][]int{}
	records := make([]struct{ mx, mx2, num int }, k+1)
	for _, x := range nums {
		if fs[x] == nil {
			fs[x] = make([]int, k+1)
		}
		f := fs[x]
		for j := k; j >= 0; j-- {
			f[j]++
			if j > 0 {
				p := records[j-1]
				m := p.mx
				if x == p.num {
					m = p.mx2
				}
				f[j] = max(f[j], m+1)
			}

			// records[j] 维护 fs[.][j] 的 mx,mx2,num
			v := f[j]
			p := &records[j]
			if v > p.mx {
				if x != p.num {
					p.num = x
					p.mx2 = p.mx
				}
				p.mx = v
			} else if x != p.num && v > p.mx2 {
				p.mx2 = v
			}
		}
	}
	return records[k].mx
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(nk)$。

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