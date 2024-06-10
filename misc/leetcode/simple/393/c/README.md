#### 题目
              
<p>给你一个整数数组 <code>coins</code> 表示不同面额的硬币，另给你一个整数 <code>k</code> 。</p>

<p>你有无限量的每种面额的硬币。但是，你<strong> 不能 </strong>组合使用不同面额的硬币。</p>

<p>返回使用这些硬币能制造的<strong> 第 </strong><code>k<sup>th</sup></code><strong> 小</strong> 金额。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="
    border-color: var(--border-tertiary);
    border-left-width: 2px;
    color: var(--text-secondary);
    font-size: .875rem;
    margin-bottom: 1rem;
    margin-top: 1rem;
    overflow: visible;
    padding-left: 1rem;">
<p><strong>输入：</strong> <span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;">coins = [3,6,9], k = 3</span></p>

<p><strong>输出：</strong> <span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;">9</span></p>

<p><strong>解释：</strong>给定的硬币可以制造以下金额：<br />
3元硬币产生3的倍数：3, 6, 9, 12, 15等。<br />
6元硬币产生6的倍数：6, 12, 18, 24等。<br />
9元硬币产生9的倍数：9, 18, 27, 36等。<br />
所有硬币合起来可以产生：3, 6, <u><strong>9</strong></u>, 12, 15等。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="
    border-color: var(--border-tertiary);
    border-left-width: 2px;
    color: var(--text-secondary);
    font-size: .875rem;
    margin-bottom: 1rem;
    margin-top: 1rem;
    overflow: visible;
    padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;">coins = [5,2], k = 7</span></p>

<p><strong>输出：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;">12</span></p>

<p><strong>解释：</strong>给定的硬币可以制造以下金额：<br />
5元硬币产生5的倍数：5, 10, 15, 20等。<br />
2元硬币产生2的倍数：2, 4, 6, 8, 10, 12等。<br />
所有硬币合起来可以产生：2, 4, 5, 6, 8, 10, <u><strong>12</strong></u>, 14, 15等。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= coins.length &lt;= 15</code></li>
	<li><code>1 &lt;= coins[i] &lt;= 25</code></li>
	<li><code>1 &lt;= k &lt;= 2 * 10<sup>9</sup></code></li>
	<li><code>coins</code> 包含两两不同的整数。</li>
</ul>

#### 思路

由于 $k$ 越大答案越大，有单调性，可以**二分答案**。不了解二分答案的同学请看 [一图掌握二分答案！四种写法！](https://leetcode.cn/problems/h-index-ii/solution/tu-jie-yi-tu-zhang-wo-er-fen-da-an-si-ch-d15k/)

问题变成：

- 可以组合出多少个不超过 $m$ 的金额？我们可以比较个数与 $k$ 的大小，来缩小二分区间。

对于面额为 $x=\textit{coins}[i]$ 的硬币，我们可以用它组合出 $\left\lfloor\dfrac{m}{x}\right\rfloor$ 个不同的金额。

比如 $\textit{coins}=[4,6],\ m=13$，用 $4$ 可以组合出 $4,8,12$ 共 $3$ 个不同的金额，用 $6$ 可以组合出 $6,12$ 共 $2$ 个不同的金额。其中 $12$ 是重复的，需要去掉。所以一共可以组合出 $3+2-1=4$ 个不同的不超过 $m=13$ 的金额。

一般地，如果只有两种面额为 $x$ 和 $y$ 的硬币，则可以组合出

$$
\left\lfloor\dfrac{m}{x}\right\rfloor + \left\lfloor\dfrac{m}{y}\right\rfloor - \left\lfloor\dfrac{m}{\texttt{LCM}(x,y)}\right\rfloor
$$

个不同的不超过 $m$ 的金额。其中 $\texttt{LCM}(x,y)$ 是 $x$ 和 $y$ 的最小公倍数。这是**容斥原理**在 $n=2$ 的情况。

我们需要枚举 $\textit{coins}$ 的所有**非空子集**，设子集大小为 $k$，子集元素的最小公倍数为 $\textit{lcm}$，那么这个子集对个数的贡献为

$$
(-1)^{k-1} \left\lfloor\dfrac{m}{\textit{lcm}}\right\rfloor
$$

累加所有非空子集的贡献，即为不同的不超过 $m$ 的金额个数。

- 开区间二分下界：$k-1$ 一定无法满足要求。
- 开区间二分上界：$\min(\textit{coins})\cdot k$ 一定可以满足要求。

``` go
func findKthSmallest(coins []int, k int) int64 {
	ans := sort.Search(slices.Max(coins)*k, func(m int) bool {
		cnt := 0
	next:
		for i := uint(1); i < 1<<len(coins); i++ { // 枚举所有非空子集
			lcmRes := 1 // 计算子集 LCM
			for j := i; j > 0; j &= j - 1 {
				lcmRes = lcm(lcmRes, coins[bits.TrailingZeros(j)])
				if lcmRes > m { // 太大了
					continue next
				}
			}
			c := m / lcmRes
			if bits.OnesCount(i)%2 == 0 {
				c = -c
			}
			cnt += c
		}
		return cnt >= k
	})
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n2^n\log (mk)\log M)$，其中 $n$ 为 $\textit{coins}$ 的长度，$m=\min(coins),\ M=\max(\textit{coins})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

