#### 题目

<p>有一个无限大的二维平面。</p>

<p>给你一个正整数&nbsp;<code>k</code>&nbsp;，同时给你一个二维数组&nbsp;<code>queries</code>&nbsp;，包含一系列查询：</p>

<ul>
	<li><code>queries[i] = [x, y]</code>&nbsp;：在平面上坐标&nbsp;<code>(x, y)</code>&nbsp;处建一个障碍物，数据保证之前的查询 <strong>不会</strong> 在这个坐标处建立任何障碍物。</li>
</ul>

<p>每次查询后，你需要找到离原点第 <code>k</code>&nbsp;<strong>近</strong>&nbsp;障碍物到原点的 <strong>距离</strong>&nbsp;。</p>

<p>请你返回一个整数数组&nbsp;<code>results</code>&nbsp;，其中&nbsp;<code>results[i]</code>&nbsp;表示建立第 <code>i</code>&nbsp;个障碍物以后，离原地第 <code>k</code>&nbsp;近障碍物距离原点的距离。如果少于 <code>k</code>&nbsp;个障碍物，<code>results[i] == -1</code>&nbsp;。</p>

<p><strong>注意</strong>，一开始&nbsp;<strong>没有</strong>&nbsp;任何障碍物。</p>

<p>坐标在&nbsp;<code>(x, y)</code>&nbsp;处的点距离原点的距离定义为&nbsp;<code>|x| + |y|</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>queries = [[1,2],[3,4],[2,3],[-3,0]], k = 2</span></p>

<p><span class="example-io"><b>输出：</b>[-1,7,5,3]</span></p>

<p><strong>解释：</strong></p>

<p>最初，不存在障碍物。</p>

<ul>
	<li><code>queries[0]</code>&nbsp;之后，少于 2 个障碍物。</li>
	<li><code>queries[1]</code>&nbsp;之后，&nbsp;两个障碍物距离原点的距离分别为 3 和 7 。</li>
	<li><code>queries[2]</code>&nbsp;之后，障碍物距离原点的距离分别为 3 ，5 和 7 。</li>
	<li><code>queries[3]</code>&nbsp;之后，障碍物距离原点的距离分别为 3，3，5 和 7 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>queries = [[5,5],[4,4],[3,3]], k = 1</span></p>

<p><span class="example-io"><b>输出：</b>[10,8,6]</span></p>

<p><b>解释：</b></p>

<ul>
	<li><code>queries[0]</code>&nbsp;之后，只有一个障碍物，距离原点距离为 10 。</li>
	<li><code>queries[1]</code>&nbsp;之后，障碍物距离原点距离分别为 8 和 10 。</li>
	<li><code>queries[2]</code>&nbsp;之后，障碍物距离原点的距离分别为 6， 8 和10 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= queries.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li>所有&nbsp;<code>queries[i]</code>&nbsp;互不相同。</li>
	<li><code>-10<sup>9</sup> &lt;= queries[i][0], queries[i][1] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

思路类似 [347. 前 K 个高频元素](https://leetcode.cn/problems/top-k-frequent-elements/)。

维护前 $k$ 小元素，可以用**最大堆**。

遍历 $\textit{queries}$，计算点 $(x,y)$ 到原点的曼哈顿距离 $d=|x|+|y|$。

把 $d$ 入堆，如果堆大小超过 $k$，就弹出堆顶（最大的元素）。

当堆的大小等于 $k$ 时，堆顶就是第 $k$ 小的距离。

``` 
func resultsArray(queries [][]int, k int) []int {
	m := len(queries)
	ans := make([]int, m)
	if m < k {
		for i := range ans {
			ans[i] = -1
		}
		return ans
	}
	h := hp{make([]int, k)}
	for i, q := range queries[:k] {
		h.IntSlice[i] = abs(q[0]) + abs(q[1])
		ans[i] = -1
	}
	heap.Init(&h)
	ans[k-1] = h.IntSlice[0]
	for i := k; i < m; i++ {
		q := queries[i]
		d := abs(q[0]) + abs(q[1])
		if d < h.IntSlice[0] {
			h.IntSlice[0] = d
			heap.Fix(&h, 0)
		}
		ans[i] = h.IntSlice[0]
	}
	return ans
}

type hp struct{ sort.IntSlice }
func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k + (m-k)\log k)$，其中 $m$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(k)$。返回值不计入。

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