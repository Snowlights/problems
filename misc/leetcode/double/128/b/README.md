### 题目

<p>给你一个二维整数数组&nbsp;<code>point</code>&nbsp;，其中&nbsp;<code>points[i] = [x<sub>i</sub>, y<sub>i</sub>]</code>&nbsp;表示二维平面内的一个点。同时给你一个整数&nbsp;<code>w</code>&nbsp;。你需要用矩形&nbsp;<strong>覆盖所有</strong>&nbsp;点。</p>

<p>每个矩形的左下角在某个点&nbsp;<code>(x<sub>1</sub>, 0)</code>&nbsp;处，且右上角在某个点&nbsp;<code>(x<sub>2</sub>, y<sub>2</sub>)</code>&nbsp;处，其中&nbsp;<code>x<sub>1</sub> &lt;= x<sub>2</sub></code>&nbsp;且&nbsp;<code>y<sub>2</sub> &gt;= 0</code>&nbsp;，同时对于每个矩形都&nbsp;<strong>必须</strong>&nbsp;满足&nbsp;<code>x<sub>2</sub> - x<sub>1</sub> &lt;= w</code>&nbsp;。</p>

<p>如果一个点在矩形内或者在边上，我们说这个点被矩形覆盖了。</p>

<p>请你在确保每个点都 <strong>至少</strong>&nbsp;被一个矩形覆盖的前提下，<strong>最少</strong>&nbsp;需要多少个矩形。</p>

<p><strong>注意：</strong>一个点可以被多个矩形覆盖。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/03/04/screenshot-from-2024-03-04-20-33-05.png" style="width: 205px; height: 300px;" /></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p style=""><span class="example-io" style="font-size: 8.75px;"><b>输入：</b></span><span class="example-io" style="font-size: 0.85rem; font-family: Menlo, sans-serif;">points = [[2,1],[1,0],[1,4],[1,8],[3,5],[4,6]], w = 1</span></p>

<p style=""><span class="example-io" style="font-size: 8.75px;"><b>输出：</b></span><span class="example-io" style="font-size: 0.85rem; font-family: Menlo, sans-serif;">2</span></p>

<p style="font-size: 0.875rem;"><strong>解释：</strong></p>

<p style="font-size: 0.875rem;">上图展示了一种可行的矩形放置方案：</p>

<ul style="font-size: 0.875rem;">
	<li>一个矩形的左下角在&nbsp;<code>(1, 0)</code>&nbsp;，右上角在&nbsp;<code>(2, 8)</code>&nbsp;。</li>
	<li>一个矩形的左下角在&nbsp;<code>(3, 0)</code>&nbsp;，右上角在&nbsp;<code>(4, 8)</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/03/04/screenshot-from-2024-03-04-18-59-12.png" style="width: 260px; height: 250px;" /></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p style=""><span class="example-io" style="font-size: 8.75px;"><b>输入：</b></span><span class="example-io" style="font-size: 0.85rem; font-family: Menlo, sans-serif;">points = [[0,0],[1,1],[2,2],[3,3],[4,4],[5,5],[6,6]], w = 2</span></p>

<p style=""><span class="example-io" style="font-size: 8.75px;"><b>输出：</b></span><span class="example-io" style="font-size: 0.85rem; font-family: Menlo, sans-serif;">3</span></p>

<p style="font-size: 0.875rem;"><b>解释：</b></p>

<p style="font-size: 0.875rem;">上图展示了一种可行的矩形放置方案：</p>

<ul style="font-size: 0.875rem;">
	<li>一个矩形的左下角在&nbsp;<code>(0, 0)</code>&nbsp;，右上角在&nbsp;<code>(2, 2)</code>&nbsp;。</li>
	<li>一个矩形的左下角在&nbsp;<code>(3, 0)</code>&nbsp;，右上角在&nbsp;<code>(5, 5)</code>&nbsp;。</li>
	<li>一个矩形的左下角在&nbsp;<code>(6, 0)</code>&nbsp;，右上角在&nbsp;<code>(6, 6)</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/03/04/screenshot-from-2024-03-04-20-24-03.png" style="height: 150px; width: 127px;" /></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p style=""><span class="example-io" style="font-size: 8.75px;"><b>输入：</b></span><span class="example-io" style="font-size: 0.85rem; font-family: Menlo, sans-serif;">points = [[2,3],[1,2]], w = 0</span></p>

<p style=""><span class="example-io" style="font-size: 8.75px;"><b>输出：</b></span><span class="example-io" style="font-size: 0.85rem; font-family: Menlo, sans-serif;">2</span></p>

<p style="font-size: 0.875rem;"><strong>解释：</strong></p>

<p style="font-size: 0.875rem;">上图展示了一种可行的矩形放置方案：</p>

<ul style="font-size: 0.875rem;">
	<li>一个矩形的左下角在&nbsp;<code>(1, 0)</code>&nbsp;，右上角在&nbsp;<code>(1, 2)</code>&nbsp;。</li>
	<li>一个矩形的左下角在&nbsp;<code>(2, 0)</code>&nbsp;，右上角在&nbsp;<code>(2, 3)</code>&nbsp;。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= points.length &lt;= 10<sup>5</sup></code></li>
	<li><code>points[i].length == 2</code></li>
	<li><code>0 &lt;= x<sub>i</sub> == points[i][0] &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= y<sub>i</sub> == points[i][1] &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= w &lt;= 10<sup>9</sup></code></li>
	<li>所有点坐标&nbsp;<code>(x<sub>i</sub>, y<sub>i</sub>)</code>&nbsp;互不相同。</li>
</ul>

### 思路

由于矩形的高没有限制，所以我们只需考虑 $\textit{points}$ 的横坐标。
矩形越大越好，所以 $x_2$ 应该恰好等于 $x_1+w$。
算法如下：

1. 把横坐标按照从小到大的顺序排序。
2. 为方便计算，假设第一个矩形左边还有一个矩形，初始化 $x_2 = -1$，因为所有横坐标都是非负数。
3. 遍历横坐标 $x=\textit{points}[i][0]$，如果 $x > x_2$，我们需要一个新的 $x_1=x$ 的矩形，答案加一，然后把 $x_2$ 更新为 $x+w$。
4. 最后，返回答案。

``` 
func minRectanglesToCoverPoints(points [][]int, w int) int {
	x, ans := make([]int, 0, len(points)), 1
	for _, v := range points {
		x = append(x, v[0])
	}
	sort.Ints(x)
	base := x[0]
	for _, v := range x {
		if base+w < v {
			ans++
			base = v
		}
	}

	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{points}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

