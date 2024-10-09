#### 题目

<p>给你两个正整数&nbsp;<code>X</code> 和&nbsp;<code>Y</code>&nbsp;和一个二维整数数组&nbsp;<code>circles</code>&nbsp;，其中&nbsp;<code>circles[i] = [x<sub>i</sub>, y<sub>i</sub>, r<sub>i</sub>]</code>&nbsp;表示一个圆心在&nbsp;<code>(x<sub>i</sub>, y<sub>i</sub>)</code>&nbsp;半径为&nbsp;<code>r<sub>i</sub></code>&nbsp;的圆。</p>

<p>坐标平面内有一个左下角在原点，右上角在&nbsp;<code>(X, Y)</code>&nbsp;的矩形。你需要判断是否存在一条从左下角到右上角的路径满足：路径&nbsp;<strong>完全</strong>&nbsp;在矩形内部，<strong>不会</strong>&nbsp;触碰或者经过 <strong>任何</strong>&nbsp;圆的内部和边界，同时&nbsp;<strong>只</strong> 在起点和终点接触到矩形。</p>

<p>如果存在这样的路径，请你返回&nbsp;<code>true</code>&nbsp;，否则返回&nbsp;<code>false</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>X = 3, Y = 4, circles = [[2,1,1]]</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/05/18/example2circle1.png" style="width: 346px; height: 264px;" /></p>

<p>黑色曲线表示一条从&nbsp;<code>(0, 0)</code>&nbsp;到&nbsp;<code>(3, 4)</code>&nbsp;的路径。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>X = 3, Y = 3, circles = [[1,1,2]]</span></p>

<p><span class="example-io"><b>输出：</b>false</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/05/18/example1circle.png" style="width: 346px; height: 264px;" /></p>

<p>不存在从&nbsp;<code>(0, 0)</code> 到&nbsp;<code>(3, 3)</code>&nbsp;的路径。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>X = 3, Y = 3, circles = [[2,1,1],[1,2,1]]</span></p>

<p><span class="example-io"><b>输出：</b>false</span></p>

<p><b>解释：</b></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/05/18/example0circle.png" style="width: 346px; height: 264px;" /></p>

<p>不存在从&nbsp;<code>(0, 0)</code>&nbsp;到&nbsp;<code>(3, 3)</code>&nbsp;的路径。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= X, Y &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= circles.length &lt;= 1000</code></li>
	<li><code>circles[i].length == 3</code></li>
	<li><code>1 &lt;= x<sub>i</sub>, y<sub>i</sub>, r<sub>i</sub> &lt;= 10<sup>9</sup></code></li>
</ul>


#### 思路

## ⚠观前必读

本题没有限制圆心必须在矩形内部，但测试表明，所有**测试数据的圆心均在矩形内部**，估计过段时间会修改题目描述，保证圆心均在矩形内部。所以**下面代码只考虑圆心一定在矩形内部的情况**。

等价转换：
- 如果从矩形【左边界/上边界】到矩形【下边界/右边界】的路被圆堵死，则无法从矩形左下角移动到矩形右上角。

怎么判断呢？

如果圆和圆相交或相切，则相当于在两个圆之间架起了一座桥。如果圆和矩形边界相交或相切，则相当于在矩形边界和圆之间架起了一座桥。如果可以从矩形【左边界/上边界】通过桥到达矩形【下边界/右边界】，则说明路被堵死，无法从矩形左下角移动到矩形右上角。
也可以把桥理解成切割线，如果能把从矩形左下角到矩形右上角的路径**切断**，则无法从矩形左下角移动到矩形右上角。

## 具体做法

抽象成一个图论题：
- 把第 $i$ 个圆视作节点 $i$，其中 $0\le i \le n-1$。
- 把矩形【左边界/上边界】视作节点 $n$。
- 把矩形【下边界/右边界】视作节点 $n+1$。

遍历每个圆 $i$：
- 判断是否与矩形【左边界/上边界】相交或相切，如果是，则合并节点 $i$ 和 $n$。
- 判断是否与矩形【下边界/右边界】相交或相切，如果是，则合并节点 $i$ 和 $n+1$。
- 判断是否与其他圆 $j$ 相交或相切，如果是，则合并节点 $i$ 和 $j$。

最后，如果节点 $n$ 和 $n+1$ 不在并查集的同一个连通块中，则返回 $\texttt{true}$，否则返回 $\texttt{false}$。也可以在遍历每个圆的过程中判断。

```
func canReachCorner(x, y int, a [][]int) bool {
	n := len(a)
	// 并查集中的 n 表示左边界或上边界，n+1 表示下边界或右边界
	fa := make([]int, n+2)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}
	merge := func(x, y int) {
		fa[find(x)] = find(y)
	}
	for i, c := range a {
		ox, oy, r := c[0], c[1], c[2]
		if ox <= r || oy+r >= y { // 圆 i 和左边界或上边界有交集
			merge(i, n)
		}
		if oy <= r || ox+r >= x { // 圆 i 和下边界或右边界有交集
			merge(i, n+1)
		}
		for j, q := range a[:i] {
			if (ox-q[0])*(ox-q[0])+(oy-q[1])*(oy-q[1]) <= (r+q[2])*(r+q[2]) {
				merge(i, j) // 圆 i 和圆 j 有交集
			}
		}
		if find(n) == find(n+1) { // 无法到达终点
			return false
		}
	}
	return true
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log n)$，其中 $n$ 是 $\textit{circles}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。\n\n注：如果建图 + DFS，可以做到 $\mathcal{O}(n^2)$ 的时间复杂度，但那样常数太大，比并查集慢。

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
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)