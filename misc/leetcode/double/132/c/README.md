### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和一个 <strong>非负</strong>&nbsp;整数&nbsp;<code>k</code>&nbsp;。如果一个整数序列&nbsp;<code>seq</code>&nbsp;满足在范围下标范围&nbsp;<code>[0, seq.length - 2]</code>&nbsp;中存在 <strong>不超过</strong>&nbsp;<code>k</code>&nbsp;个下标 <code>i</code>&nbsp;满足&nbsp;<code>seq[i] != seq[i + 1]</code>&nbsp;，那么我们称这个整数序列为&nbsp;<strong>好</strong>&nbsp;序列。</p>

<p>请你返回 <code>nums</code>&nbsp;中&nbsp;<strong>好</strong> <span data-keyword="subsequence-array">子序列</span>&nbsp;的最长长度</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,1,1,3], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p>最长好子序列为&nbsp;<code>[<em><strong>1</strong></em>,<em><strong>2</strong></em>,<strong><em>1</em></strong>,<em><strong>1</strong></em>,3]</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3,4,5,1], k = 0</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>最长好子序列为&nbsp;<code>[<strong><em>1</em></strong>,2,3,4,5,<strong><em>1</em></strong>]</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 500</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= k &lt;= min(nums.length, 25)</code></li>
</ul>


### 思路

遍历到 $x=\textit{nums}[i]$ 时，我们需要维护以 $x$ 结尾的、至多包含 $j$ 个不同相邻元素的子序列的最大长度，定义为 $f[x][j]$，初始全为 $0$。

对于 $x$，有三种决策：

1. 不选：$f[x][j]$ 不变。
2. 选，且和子序列的前一个数一样，或者作为子序列的第一个数：$f[x][j]$ 增加 $1$。
3. 选，且和子序列的前一个数不一样：设前一个数为 $y$，我们需要知道最大的 $f[y][j-1]$。

对于第三种决策，暴力枚举 $y$ 就太慢了（可以通过第三题但无法通过本题）。我们可以维护 $f[\cdot][j-1]$ 中的最大值 $\textit{mx}$、最大值对应的数字 $\textit{num}$，以及 $f[\textit{num}_2][j-1]$ 中的最大值 $\textit{mx}_2$，其中 $\textit{num}_2\ne \textit{num}$。

于是：

- 如果 $x\ne \textit{num}$，那么最大的 $f[y][j-1]$ 就是 $\textit{mx}$。
- 如果 $x = \textit{num}$，那么最大的 $f[y][j-1]$ 就是 $\textit{mx}_2$。

把最大的 $f[y][j-1]$ 记作 $m$，则 $f[x][j]$ 更新为

$$
\max(f[x][j] + 1, m + 1)
$$

对于不同的 $j$，我们需要维护对应的 $\textit{mx},\textit{mx}_2,\textit{num}$。用一个长为 $k+1$ 的数组 $\textit{records}$ 记录。

由于在计算 $f[x][j]$ 时会用到 $\textit{records}[j-1]$，然后更新 $\textit{records}[j]$，可以倒序枚举 $j$，以避免使用覆盖后的数据。

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