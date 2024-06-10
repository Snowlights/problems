#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的数组 <code>points</code> ，它表示二维平面上一些点的整数坐标，其中 <code>points[i] = [x<sub>i</sub>, y<sub>i</sub>]</code> 。</p>

<p>两点之间的距离定义为它们的<span data-keyword="manhattan-distance">曼哈顿距离</span>。</p>

<p>请你恰好移除一个点，返回移除后任意两点之间的 <strong>最大 </strong>距离可能的 <strong>最小 </strong>值。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>points = [[3,10],[5,15],[10,2],[4,4]]
<strong>输出：</strong>12
<strong>解释：</strong>移除每个点后的最大距离如下所示：
- 移除第 0 个点后，最大距离在点 (5, 15) 和 (10, 2) 之间，为 |5 - 10| + |15 - 2| = 18 。
- 移除第 1 个点后，最大距离在点 (3, 10) 和 (10, 2) 之间，为 |3 - 10| + |10 - 2| = 15 。
- 移除第 2 个点后，最大距离在点 (5, 15) 和 (4, 4) 之间，为 |5 - 4| + |15 - 4| = 12 。
- 移除第 3 个点后，最大距离在点 (5, 15) 和 (10, 2) 之间的，为 |5 - 10| + |15 - 2| = 18 。
在恰好移除一个点后，任意两点之间的最大距离可能的最小值是 12 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>points = [[1,1],[1,1],[1,1]]
<strong>输出：</strong>0
<strong>解释：</strong>移除任一点后，任意两点之间的最大距离都是 0 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= points.length &lt;= 10<sup>5</sup></code></li>
	<li><code>points[i].length == 2</code></li>
	<li><code>1 &lt;= points[i][0], points[i][1] &lt;= 10<sup>8</sup></code></li>
</ul>

#### 思路

![w391-c.png](https://pic.leetcode.cn/1711856129-vydCEa-w391-c.png)

这两种投影长度，其中较大者为曼哈顿距离（较小者是两段折线的投影长度之差，不合法），即如下恒等式

$$
|x_1 - x_2| + |y_1 - y_2| = \max(|x_1'-x_2'|,|y_1'-y_2'|)
$$

其中等式左侧为 $(x_1,y_1)$ 和 $(x_2,y_2)$ 的**曼哈顿距离**，等式右侧 $(x',y') = (x+y,y-x)$，
计算的是 $(x_1',y_1')$ 和 $(x_2',y_2')$ 两点的曼哈顿距离投影到 $x'$ 轴和 $y'$ 轴的线段长度的最大值， 即**切比雪夫距离**。

所以要求任意两点曼哈顿距离的最大值，根据上面的恒等式，我们只需要计算任意两个 $(x',y')$ 切比雪夫距离的最大值，
即横纵坐标差的最大值

$$
\max(\max(x') - \min(x'), \max(y') - \min(y'))
$$

## 方法一：有序集合

本题可以删除一个点，我们可以枚举要删除的点，并用有序集合维护其它 $n-1$ 个点的 $x'$ 和 $y'$ 的最大值和最小值。

```go [sol-Go]
func minimumDistance(points [][]int) int {
	xs := redblacktree.New[int, int]()
	ys := redblacktree.New[int, int]()
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		put(xs, x)
		put(ys, y)
	}
	ans := math.MaxInt
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		remove(xs, x)
		remove(ys, y)
		ans = min(ans, max(xs.Right().Key-xs.Left().Key, ys.Right().Key-ys.Left().Key))
		put(xs, x)
		put(ys, y)
	}
	return ans
}

func put(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	t.Put(v, c+1)
}

func remove(t *redblacktree.Tree[int, int], v int) {
	c, _ := t.Get(v)
	if c == 1 {
		t.Remove(v)
	} else {
		t.Put(v, c-1)
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：维护最大次大、最小次小

优化：如果把最大的 $x'$ 删掉，那么次大 $x'$ 就作为剩下 $n-1$ 个 $x'$ 中的最大值了，对于最小值也同理。
所以我们可以维护 $x'$ 和 $y'$ 的最大次大、最小次小，一共 $8$ 个数。

```go [sol-Go]
func minimumDistance2(points [][]int) int {
	const inf = math.MaxInt
	maxX1, maxX2, maxY1, maxY2 := -inf, -inf, -inf, -inf
	minX1, minX2, minY1, minY2 := inf, inf, inf, inf

	for _, p := range points {
		x := p[0] + p[1]
		y := p[1] - p[0]

		// x 最大次大
		if x > maxX1 {
			maxX2 = maxX1
			maxX1 = x
		} else if x > maxX2 {
			maxX2 = x
		}

		// x 最小次小
		if x < minX1 {
			minX2 = minX1
			minX1 = x
		} else if x < minX2 {
			minX2 = x
		}

		// y 最大次大
		if y > maxY1 {
			maxY2 = maxY1
			maxY1 = y
		} else if y > maxY2 {
			maxY2 = y
		}

		// y 最小次小
		if y < minY1 {
			minY2 = minY1
			minY1 = y
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := inf
	for _, p := range points {
		x := p[0] + p[1]
		y := p[1] - p[0]
		dx := f(x, maxX1, maxX2) - f(x, minX1, minX2)
		dy := f(y, maxY1, maxY2) - f(y, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(v, v1, v2 int) int {
	if v == v1 {
		return v2
	}
	return v1
}
```

进一步地，要删除的点只能是 $x'$ 或 $y'$ 最大最小的点（不然删不删都一样），所以额外维护最大最小值的下标，可以少一次对 $\textit{points}$ 的遍历。\

```go
func minimumDistance(points [][]int) int {
	const inf = math.MaxInt
	maxX1, maxX2, maxY1, maxY2 := -inf, -inf, -inf, -inf
	minX1, minX2, minY1, minY2 := inf, inf, inf, inf
	var maxXi, minXi, maxYi, minYi int

	for i, p := range points {
		x := p[0] + p[1]
		y := p[1] - p[0]

		// x 最大次大
		if x > maxX1 {
			maxX2 = maxX1
			maxX1 = x
			maxXi = i
		} else if x > maxX2 {
			maxX2 = x
		}

		// x 最小次小
		if x < minX1 {
			minX2 = minX1
			minX1 = x
			minXi = i
		} else if x < minX2 {
			minX2 = x
		}

		// y 最大次大
		if y > maxY1 {
			maxY2 = maxY1
			maxY1 = y
			maxYi = i
		} else if y > maxY2 {
			maxY2 = y
		}

		// y 最小次小
		if y < minY1 {
			minY2 = minY1
			minY1 = y
			minYi = i
		} else if y < minY2 {
			minY2 = y
		}
	}

	ans := math.MaxInt
	for _, i := range []int{maxXi, minXi, maxYi, minYi} {
		dx := f(i != maxXi, maxX1, maxX2) - f(i != minXi, minX1, minX2)
		dy := f(i != maxYi, maxY1, maxY2) - f(i != minYi, minY1, minY2)
		ans = min(ans, max(dx, dy))
	}
	return ans
}

func f(b bool, v1, v2 int) int {
	if b {
		return v1
	}
	return v2
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{points}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
