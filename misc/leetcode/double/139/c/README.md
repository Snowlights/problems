### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和一个 <strong>正</strong>&nbsp;整数&nbsp;<code>k</code>&nbsp;。</p>

<p>定义长度为 <code>2 * x</code>&nbsp;的序列 <code>seq</code>&nbsp;的 <strong>值</strong>&nbsp;为：</p>

<ul>
	<li><code>(seq[0] OR seq[1] OR ... OR seq[x - 1]) XOR (seq[x] OR seq[x + 1] OR ... OR seq[2 * x - 1])</code>.</li>
</ul>

<p>请你求出 <code>nums</code>&nbsp;中所有长度为 <code>2 * k</code>&nbsp;的 <span data-keyword="subsequence-array">子序列</span> 的 <strong>最大值</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,6,7], k = 1</span></p>

<p><span class="example-io"><b>输出：</b>5</span></p>

<p><strong>解释：</strong></p>

<p>子序列&nbsp;<code>[2, 7]</code>&nbsp;的值最大，为&nbsp;<code>2 XOR 7 = 5</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [4,2,5,6,7], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>子序列&nbsp;<code>[4, 5, 6, 7]</code>&nbsp;的值最大，为&nbsp;<code>(4 OR 5) XOR (6 OR 7) = 2</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 400</code></li>
	<li><code>1 &lt;= nums[i] &lt; 2<sup>7</sup></code></li>
	<li><code>1 &lt;= k &lt;= nums.length / 2</code></li>
</ul>


### 思路

想象有一根分割线，把 $\textit{nums}$ 分成左右两部分，左部和右部分别计算所有长为 $k$ 的子序列的 OR 都**有哪些值**。

比如左部计算出的 OR 有 $[2,3,5]$，右部计算出的 OR 有 $[1,3,6]$，两两组合计算 XOR，取其中最大值作为答案。

把 OR 理解成一个类似加法的东西，本题是一个二维的 0-1 背包。不了解 0-1 背包的同学请看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

考虑计算右部，定义 $f[i][j][x]$ 表示从 $\textit{nums}[i]$ 到 $\textit{nums}[n-1]$ 中选出 $j$ 个数，这些数的 OR 是否等于 $x$。

设 $v=\textit{nums}[i]$，用**刷表法**转移：

- 不选 $v$，那么 $f[i][j][x] = f[i+1][j][x]$。
- 选 $v$，如果 $f[i+1][j][x]=\texttt{true}$，那么 $f[i][j][x|v]=\texttt{true}$。

初始值 $f[n][0][0]=\texttt{true}$。什么也不选，OR 等于 $0$。

对于每个 $i$，由于我们只需要 $f[i][k]$ 中的数据，把 $f[i][k]$ 复制到 $\textit{suf}[i]$ 中。这样做无需创建三维数组，空间复杂度更小。

代码实现时，$f$ 的第一个维度可以优化掉。

对于左部 $\textit{pre}$ 的计算也同理。

最后，枚举 $i=k-1,k,k+1,\cdots,n-k-1$，两两组合 $\textit{pre}[i]$ 和 $\textit{suf}[i+1]$ 中的元素计算 XOR，取其中最大值作为答案。


```
func maxValue(nums []int, k int) (ans int) {
	const mx = 1 << 7
	n := len(nums)
	suf := make([][mx]bool, n)
	f := make([][mx]bool, k+1)
	f[0][0] = true
	for i := n - 1; i >= k; i-- {
		v := nums[i]
		for j := k - 1; j >= 0; j-- {
			for x, hasX := range f[j] {
				if hasX {
					f[j+1][x|v] = true
				}
			}
		}
		suf[i] = f[k]
	}

	pre := make([][mx]bool, k+1)
	pre[0][0] = true
	for i, v := range nums[:n-k] {
		for j := k - 1; j >= 0; j-- {
			for x, hasX := range pre[j] {
				if hasX {
					pre[j+1][x|v] = true
				}
			}
		}
		if i < k-1 {
			continue
		}
		for x, hasX := range pre[k] {
			if hasX {
				for y, hasY := range suf[i+1] {
					if hasY {
						ans = max(ans, x^y)
					}
				}
			}
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nkU)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U$ 是 $\textit{nums}$ 所有元素的 OR，本题至多为 $2^7-1$。
- 空间复杂度：$\mathcal{O}(n+k)U)$。


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