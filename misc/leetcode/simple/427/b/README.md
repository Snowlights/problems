#### 题目

<p>给你一个数组 <code>points</code>，其中 <code>points[i] = [x<sub>i</sub>, y<sub>i</sub>]</code> 表示无限平面上一点的坐标。</p>

<p>你的任务是找出满足以下条件的矩形可能的&nbsp;<strong>最大&nbsp;</strong>面积：</p>

<ul>
	<li>矩形的四个顶点必须是数组中的&nbsp;<strong>四个&nbsp;</strong>点。</li>
	<li>矩形的内部或边界上&nbsp;<strong>不能&nbsp;</strong>包含任何其他点。</li>
	<li>矩形的边与坐标轴&nbsp;<strong>平行&nbsp;</strong>。</li>
</ul>

<p>返回可以获得的&nbsp;<strong>最大面积&nbsp;</strong>，如果无法形成这样的矩形，则返回 -1。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">points = [[1,1],[1,3],[3,1],[3,3]]</span></p>

<p><strong>输出：</strong>4</p>

<p><strong>解释：</strong></p>

<p><strong class="example"><img alt="示例 1 图示" src="https://assets.leetcode.com/uploads/2024/11/02/example1.png" style="width: 229px; height: 228px;" /></strong></p>

<p>我们可以用这 4 个点作为顶点构成一个矩形，并且矩形内部或边界上没有其他点。因此，最大面积为 4 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">points = [[1,1],[1,3],[3,1],[3,3],[2,2]]</span></p>

<p><strong>输出：</strong>-1</p>

<p><strong>解释：</strong></p>

<p><strong class="example"><img alt="示例 2 图示" src="https://assets.leetcode.com/uploads/2024/11/02/example2.png" style="width: 229px; height: 228px;" /></strong></p>

<p>唯一一组可能构成矩形的点为 <code>[1,1], [1,3], [3,1]</code> 和 <code>[3,3]</code>，但点 <code>[2,2]</code> 总是位于矩形内部。因此，返回 -1 。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">points = [[1,1],[1,3],[3,1],[3,3],[1,2],[3,2]]</span></p>

<p><strong>输出：</strong>2</p>

<p><strong>解释：</strong></p>

<p><strong class="example"><img alt="示例 3 图示" src="https://assets.leetcode.com/uploads/2024/11/02/example3.png" style="width: 229px; height: 228px;" /></strong></p>

<p>点 <code>[1,3], [1,2], [3,2], [3,3]</code>&nbsp;可以构成面积最大的矩形，面积为 2。此外，点 <code>[1,1], [1,2], [3,1], [3,2]</code> 也可以构成一个符合题目要求的矩形，面积相同。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= points.length &lt;= 10</code></li>
	<li><code>points[i].length == 2</code></li>
	<li><code>0 &lt;= x<sub>i</sub>, y<sub>i</sub> &lt;= 100</code></li>
	<li>给定的所有点都是 <strong>唯一</strong> 的。</li>
</ul>

#### 思路

本题解旨在解决两个更一般的问题：
1. 找到所有矩形，这些矩形的边界（除去顶点）不含其他点。
2. 计算每个矩形区域中的点的个数。本题要求点的个数恰好等于 $4$（四个顶点）。

由于矩形边界不能包含其他点，那么枚举矩形的右上角 $(x_2,y_2)$，它正左边的点就是左上角，正下方的点就是右下角。

左上角的横坐标 $x_1$ 和右下角的纵坐标 $y_1$ 就组成了矩形的左下角 $(x_1,y_1)$。
> 需要判断：矩形右下角的左边的点的横坐标必须是 $x_1$，矩形左上角的下边的点的纵坐标必须是 $y_1$。

把点按照横坐标分组，即可快速知道每个点的正下方的点；按照纵坐标分组，即可快速知道每个点的正左边的点。
现在，问题变成一个（静态的）**二维数点问题**：
- 给你一些询问，对于每个询问，你需要计算横坐标在 $[x_1,x_2]$ 中且纵坐标在 $[y_1,y_2]$ 中的点的个数。如果恰好有 $4$ 个点，用矩形面积 $(x_2-x_1)\cdot(y_2-y_1)$ 更新答案的最大值。

这可以**离线**解决，核心想法是：
- 计算横坐标 $\le x_2$ 且纵坐标在 $[y_1,y_2]$ 中的点的个数。
- 计算横坐标 $\le x_1-1$ 且纵坐标在 $[y_1,y_2]$ 中的点的个数。
- 二者相减，即为矩形区域中的点的个数。

具体来说：
1. 把询问按照 $x_1-1$ 和 $x_2$ 的值分组。
2. 对于 $x_1-1$，记录如下信息：
    - 询问的编号。
    - 计算点的个数时，这部分要减去，系数是 $-1$。
    - 纵坐标 $y_1$ 和 $y_2$。
3. 对于 $x_2$，记录如下信息：
    - 询问的编号。
    - 计算点的个数时，这部分要加上，系数是 $1$。
    - 纵坐标 $y_1$ 和 $y_2$。

然后用 [树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/) 按照 $x$ 从小到大的顺序记录、计算点的个数（区间中的 $y$ 的个数）。
由于数据范围很大，需要离散化。

```
type fenwick []int

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func maxRectangleArea(xCoord, yCoord []int) int64 {
	xMap := map[int][]int{} // 同一列的所有点的纵坐标
	yMap := map[int][]int{} // 同一行的所有点的横坐标
	for i, x := range xCoord {
		y := yCoord[i]
		xMap[x] = append(xMap[x], y)
		yMap[y] = append(yMap[y], x)
	}

	// 预处理每个点的正下方的点
	type pair struct{ x, y int }
	below := map[pair]int{}
	for x, ys := range xMap {
		slices.Sort(ys)
		for i := 1; i < len(ys); i++ {
			below[pair{x, ys[i]}] = ys[i-1]
		}
	}

	// 预处理每个点的正左边的点
	left := map[pair]int{}
	for y, xs := range yMap {
		slices.Sort(xs)
		for i := 1; i < len(xs); i++ {
			left[pair{xs[i], y}] = xs[i-1]
		}
	}

	// 离散化用
	xs := slices.Sorted(maps.Keys(xMap))
	ys := slices.Sorted(maps.Keys(yMap))

	// 收集询问：矩形区域（包括边界）的点的个数
	type query struct{ x1, x2, y1, y2, area int }
	queries := []query{}
	// 枚举 (x2,y2) 作为矩形的右上角
	for x2, listY := range xMap {
		for i := 1; i < len(listY); i++ {
			// 计算矩形左下角 (x1,y1)
			y2 := listY[i]
			x1, ok := left[pair{x2, y2}]
			if !ok {
				continue
			}
			y1 := listY[i-1] // (x2,y2) 下面的点（矩形右下角）的纵坐标
			// 矩形右下角的左边的点的横坐标必须是 x1
			if x, ok := left[pair{x2, y1}]; !ok || x != x1 {
				continue
			}
			// 矩形左上角的下边的点的纵坐标必须是 y1
			if y, ok := below[pair{x1, y2}]; !ok || y != y1 {
				continue
			}
			queries = append(queries, query{
				sort.SearchInts(xs, x1), // 离散化
				sort.SearchInts(xs, x2),
				sort.SearchInts(ys, y1),
				sort.SearchInts(ys, y2),
				(x2 - x1) * (y2 - y1),
			})
		}
	}

	// 离线询问
	type data struct{ qid, sign, y1, y2 int }
	qs := make([][]data, len(xs))
	for i, q := range queries {
		if q.x1 > 0 {
			qs[q.x1-1] = append(qs[q.x1-1], data{i, -1, q.y1, q.y2})
		}
		qs[q.x2] = append(qs[q.x2], data{i, 1, q.y1, q.y2})
	}

	// 回答询问
	res := make([]int, len(queries))
	tree := make(fenwick, len(ys)+1)
	for i, x := range xs {
		// 把横坐标为 x 的所有点都加到树状数组中
		for _, y := range xMap[x] {
			tree.add(sort.SearchInts(ys, y) + 1) // 离散化
		}
		for _, q := range qs[i] {
			// 查询 [y1,y2] 中点的个数
			res[q.qid] += q.sign * tree.query(q.y1+1, q.y2+1)
		}
	}

	ans := -1
	for i, cnt := range res {
		if cnt == 4 { // 矩形区域（包括边界）恰好有 4 个点
			ans = max(ans, queries[i].area)
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{xCoord}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)