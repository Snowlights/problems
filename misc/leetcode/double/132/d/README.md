### 题目

<p>有一条无限长的数轴，原点在 0 处，沿着 x 轴 <strong>正</strong>&nbsp;方向无限延伸。</p>

<p>给你一个二维数组&nbsp;<code>queries</code>&nbsp;，它包含两种操作：</p>

<ol>
	<li>操作类型 1 ：<code>queries[i] = [1, x]</code>&nbsp;。在距离原点 <code>x</code>&nbsp;处建一个障碍物。数据保证当操作执行的时候，位置 <code>x</code>&nbsp;处 <strong>没有</strong>&nbsp;任何障碍物。</li>
	<li>操作类型 2 ：<code>queries[i] = [2, x, sz]</code>&nbsp;。判断在数轴范围&nbsp;<code>[0, x]</code>&nbsp;内是否可以放置一个长度为&nbsp;<code>sz</code>&nbsp;的物块，这个物块需要&nbsp;<strong>完全</strong>&nbsp;放置在范围&nbsp;<code>[0, x]</code>&nbsp;内。如果物块与任何障碍物有重合，那么这个物块&nbsp;<strong>不能</strong>&nbsp;被放置，但物块可以与障碍物刚好接触。注意，你只是进行查询，并&nbsp;<strong>不是</strong>&nbsp;真的放置这个物块。每个查询都是相互独立的。</li>
</ol>

<p>请你返回一个 boolean 数组<code>results</code>&nbsp;，如果第&nbsp;<code>i</code> 个操作类型 2 的操作你可以放置物块，那么&nbsp;<code>results[i]</code>&nbsp;为&nbsp;<code>true</code>&nbsp;，否则为 <code>false</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>queries = [[1,2],[2,3,3],[2,3,1],[2,2,2]]</span></p>

<p><span class="example-io"><b>输出：</b>[false,true,true]</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/04/22/example0block.png" style="padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; width: 309px; height: 129px;" /></strong></p>

<p>查询 0 ，在&nbsp;<code>x = 2</code>&nbsp;处放置一个障碍物。在&nbsp;<code>x = 3</code>&nbsp;之前任何大小不超过 2 的物块都可以被放置。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>queries = </span>[[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]<!-- notionvc: 4a471445-5af1-4d72-b11b-94d351a2c8e9 --></p>

<p><b>输出：</b>[true,true,false]</p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/04/22/example1block.png" style="padding: 10px; background: rgb(255, 255, 255); border-radius: 0.5rem; width: 310px; height: 130px;" /></strong></p>

<ul>
	<li>查询 0 在&nbsp;<code>x = 7</code>&nbsp;处放置一个障碍物。在&nbsp;<code>x = 7</code>&nbsp;之前任何大小不超过 7 的物块都可以被放置。</li>
	<li>查询 2 在&nbsp;<code>x = 2</code>&nbsp;处放置一个障碍物。现在，在&nbsp;<code>x = 7</code>&nbsp;之前任何大小不超过 5 的物块可以被放置，<code>x = 2</code>&nbsp;之前任何大小不超过 2 的物块可以被放置。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= queries.length &lt;= 15 * 10<sup>4</sup></code></li>
	<li><code>2 &lt;= queries[i].length &lt;= 3</code></li>
	<li><code>1 &lt;= queries[i][0] &lt;= 2</code></li>
	<li><code>1 &lt;= x, sz &lt;= min(5 * 10<sup>4</sup>, 3 * queries.length)</code></li>
	<li>输入保证操作 1 中，<code>x</code>&nbsp;处不会有障碍物。</li>
	<li>输入保证至少有一个操作类型 2 。</li>
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