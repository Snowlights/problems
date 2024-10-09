#### 题目

<p>给你一个整数 <code>mountainHeight</code> 表示山的高度。</p>

<p>同时给你一个整数数组 <code>workerTimes</code>，表示工人们的工作时间（单位：<strong>秒</strong>）。</p>

<p>工人们需要 <strong>同时 </strong>进行工作以 <strong>降低 </strong>山的高度。对于工人 <code>i</code> :</p>

<ul>
	<li>山的高度降低 <code>x</code>，需要花费 <code>workerTimes[i] + workerTimes[i] * 2 + ... + workerTimes[i] * x</code> 秒。例如：
	<ul>
		<li>山的高度降低 1，需要 <code>workerTimes[i]</code> 秒。</li>
		<li>山的高度降低 2，需要 <code>workerTimes[i] + workerTimes[i] * 2</code> 秒，依此类推。</li>
	</ul>
	</li>
</ul>

<p>返回一个整数，表示工人们使山的高度降低到 0 所需的 <strong>最少</strong> 秒数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">mountainHeight = 4, workerTimes = [2,1,1]</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p>将山的高度降低到 0 的一种方式是：</p>

<ul>
	<li>工人 0 将高度降低 1，花费 <code>workerTimes[0] = 2</code> 秒。</li>
	<li>工人 1 将高度降低 2，花费 <code>workerTimes[1] + workerTimes[1] * 2 = 3</code> 秒。</li>
	<li>工人 2 将高度降低 1，花费 <code>workerTimes[2] = 1</code> 秒。</li>
</ul>

<p>因为工人同时工作，所需的最少时间为 <code>max(2, 3, 1) = 3</code> 秒。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">mountainHeight = 10, workerTimes = [3,2,2,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">12</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>工人 0 将高度降低 2，花费 <code>workerTimes[0] + workerTimes[0] * 2 = 9</code> 秒。</li>
	<li>工人 1 将高度降低 3，花费 <code>workerTimes[1] + workerTimes[1] * 2 + workerTimes[1] * 3 = 12</code> 秒。</li>
	<li>工人 2 将高度降低 3，花费 <code>workerTimes[2] + workerTimes[2] * 2 + workerTimes[2] * 3 = 12</code> 秒。</li>
	<li>工人 3 将高度降低 2，花费 <code>workerTimes[3] + workerTimes[3] * 2 = 12</code> 秒。</li>
</ul>

<p>所需的最少时间为 <code>max(9, 12, 12, 12) = 12</code> 秒。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">mountainHeight = 5, workerTimes = [1]</span></p>

<p><strong>输出：</strong> <span class="example-io">15</span></p>

<p><strong>解释：</strong></p>

<p>这个示例中只有一个工人，所以答案是 <code>workerTimes[0] + workerTimes[0] * 2 + workerTimes[0] * 3 + workerTimes[0] * 4 + workerTimes[0] * 5 = 15</code> 秒。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= mountainHeight &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= workerTimes.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= workerTimes[i] &lt;= 10<sup>6</sup></code></li>
</ul>

#### 思路

## 方法一：最小堆模拟

循环 $\textit{mountainHeight}$ 次，每次选一个「工作后总用时」最短的工人，把山的高度降低 $1$。


```
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	h := make(hp, len(workerTimes))
	for i, t := range workerTimes {
		h[i] = worker{t, t, t}
	}
	heap.Init(&h)

	ans := 0
	for ; mountainHeight > 0; mountainHeight-- {
		ans = max(ans, h[0].nxt)
		h[0].delta += h[0].base
		h[0].nxt += h[0].delta
		heap.Fix(&h, 0)
	}
	return int64(ans)
}

// 工作后总用时，当前工作（山高度降低 1）用时，workerTimes[i]
type worker struct{ nxt, delta, base int }
type hp []worker
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].nxt < h[j].nxt }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\textit{mountainHeight}\log n)$，其中 $n$ 是 $\textit{workerTimes}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：二分答案

由于花的时间越多，能够降低的高度也越多，所以有**单调性**，可以二分答案。

问题变成：

- 每个工人**至多**花费 $m$ 秒，总共降低的高度是多少？能否大于等于 $\textit{mountainHeight}$？

遍历 $\textit{workerTimes}$，设 $t=\textit{workerTimes}[i]$，那么有

$$
t + 2t+ \cdots + xt = t\cdot \dfrac{x(x+1)}{2} \le m
$$

即

$$
\dfrac{x(x+1)}{2} \le \left\lfloor\dfrac{m}{t}\right\rfloor = k
$$

解得

$$
x \le \dfrac{-1 + \sqrt{1 + 8k}}{2}
$$

所以第 $i$ 名工人可以把山的高度降低

$$
\left\lfloor \dfrac{-1 + \sqrt{1 + 8k}}{2} \right\rfloor
$$

累加上式，如果和 $\ge \textit{mountainHeight}$，则说明答案 $\le m$，否则说明答案 $> m$。

最后，讨论二分的上下界。这里用开区间二分，其他二分写法也是可以的。

- 开区间二分下界：$0$，无法把山的高度降低到 $0$。
- 开区间二分上界：设 $\textit{maxT}$ 为 $\textit{workerTimes}$ 的最大值，假设每个工人都是最慢的 $\textit{maxT}$，那么单个工人要把山降低 $h=\left\lceil\dfrac{mountainHeight}{n}\right\rceil$，耗时 $\textit{maxT}\cdot(1+2+\cdots+h)=\textit{maxT}\cdot\dfrac{h(h+1)}{2}$，将其作为开区间的二分上界，一定可以把山的高度降低到 $\le 0$。

```
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxT := slices.Max(workerTimes)
	h := (mountainHeight-1)/len(workerTimes) + 1
	ans := 1 + sort.Search(maxT*h*(h+1)/2, func(m int) bool {
		m++
		leftH := mountainHeight
		for _, t := range workerTimes {
			leftH -= int((math.Sqrt(float64(m/t*8+1)) - 1) / 2)
			if leftH <= 0 {
				return true
			}
		}
		return false
	})
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{workerTimes}$ 的长度，$U\le 5\cdot 10^{10}(10^5+1)$ 是二分上界。二分 $\mathcal{O}(\log U)$ 次，每次 $\mathcal{O}(n)$ 时间。开平方有专门的 CPU 指令，可以视作 $\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

更多相似题目，见下面二分题单中的「**二分答案：求最小**」。


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