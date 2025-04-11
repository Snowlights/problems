### 题目

<p>给你一个二维整数数组 <code>squares</code>&nbsp;，其中&nbsp;<code>squares[i] = [x<sub>i</sub>, y<sub>i</sub>, l<sub>i</sub>]</code> 表示一个与 x 轴平行的正方形的左下角坐标和正方形的边长。</p>

<p>找到一个<strong>最小的</strong> y 坐标，它对应一条水平线，该线需要满足它以上正方形的总面积 <strong>等于</strong> 该线以下正方形的总面积。</p>

<p>答案如果与实际答案的误差在 <code>10<sup>-5</sup></code> 以内，将视为正确答案。</p>

<p><strong>注意</strong>：正方形&nbsp;<strong>可能会&nbsp;</strong>重叠。重叠区域只&nbsp;<strong>统计一次&nbsp;</strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">squares = [[0,0,1],[2,2,1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">1.00000</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://pic.leetcode.cn/1739609602-zhNmeC-4065example1drawio.png" style="width: 269px; height: 203px;" /></p>

<p>任何在 <code>y = 1</code> 和 <code>y = 2</code> 之间的水平线都会有 1 平方单位的面积在其上方，1 平方单位的面积在其下方。最小的 y 坐标是 1。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">squares = [[0,0,2],[1,1,1]]</span></p>

<p><strong>输出：</strong> <span class="example-io">1.00000</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://pic.leetcode.cn/1739609605-ezeVgk-4065example2drawio.png" style="width: 269px; height: 203px;" /></p>

<p>由于蓝色正方形和红色正方形有重叠区域且重叠区域只统计一次。所以直线&nbsp;<code>y = 1</code>&nbsp;将正方形分割成两部分且面积相等。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= squares.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>squares[i] = [x<sub>i</sub>, y<sub>i</sub>, l<sub>i</sub>]</code></li>
	<li><code>squares[i].length == 3</code></li>
	<li><code>0 &lt;= x<sub>i</sub>, y<sub>i</sub> &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= l<sub>i</sub> &lt;= 10<sup>9</sup></code></li>
</ul>


### 思路

**前置题目**：[850. 矩形面积 II](https://leetcode.cn/problems/rectangle-area-ii/)，[我的题解](https://leetcode.cn/problems/rectangle-area-ii/solutions/3078272/lazy-xian-duan-shu-sao-miao-xian-pythonj-4tkr/)。

首先用 850 题的扫描线方法，求出所有正方形的面积并 $\textit{totArea}$。

然后再次扫描，设扫描线下方的面积和为 $\textit{area}$，那么扫描线上方的面积和为 $\textit{totArea} - \textit{area}$。

题目要求

$$
\textit{area} = \textit{totArea} - \textit{area}
$$

即

$$
\textit{area}\cdot 2 = \textit{totArea}
$$

设当前扫描线的纵坐标为 $y$，下一个需要经过的正方形上/下边界的纵坐标为 $y'$，被至少一个正方形覆盖的底边长之和为 $\textit{sumLen}$，那么新的面积和为

$$
\textit{area} + \textit{sumLen} \cdot (y'-y)
$$

如果发现

$$
(\textit{area} + \textit{sumLen} \cdot (y'-y))\cdot 2 \ge \textit{totArea}
$$

取等号，解得

$$
y' = y + \dfrac{\textit{totalArea}/2 - \textit{area}}{\textit{sumL}} = y + \dfrac{\textit{totalArea} - \textit{area}\cdot 2}{\textit{sumL}\cdot 2}
$$

即为答案。

**编程技巧**：把第一次扫描过程中的关键数据 $\textit{area}$ 和 $\textit{sumLen}$ 记录到一个数组中，然后遍历数组（或者二分），这样可以避免跑两遍线段树（空间换时间）。

```
// 线段树每个节点维护一段横坐标区间 [lx, rx]
type seg []struct {
    l, r        int
    minCoverLen int // 区间内被矩形覆盖次数最少的底边长之和
    minCover    int // 区间内被矩形覆盖的最小次数
    todo        int // 子树内的所有节点的 minCover 需要增加的量，注意这可以是负数
}

// 根据左右儿子的信息，更新当前节点的信息
func (t seg) maintain(o int) {
    lo, ro := &t[o<<1], &t[o<<1|1]
    mn := min(lo.minCover, ro.minCover)
    t[o].minCover = mn
    t[o].minCoverLen = 0
    if lo.minCover == mn { // 只统计等于 minCover 的底边长之和
        t[o].minCoverLen = lo.minCoverLen
    }
    if ro.minCover == mn {
        t[o].minCoverLen += ro.minCoverLen
    }
}

// 仅更新节点信息，不下传懒标记 todo
func (t seg) do(o, v int) {
    t[o].minCover += v
    t[o].todo += v
}

// 下传懒标记 todo
func (t seg) spread(o int) {
    v := t[o].todo
    if v != 0 {
        t.do(o<<1, v)
        t.do(o<<1|1, v)
        t[o].todo = 0
    }
}

// 建树
func (t seg) build(xs []int, o, l, r int) {
    t[o].l, t[o].r = l, r
    t[o].todo = 0
    if l == r {
        t[o].minCover = 0
        t[o].minCoverLen = xs[l+1] - xs[l]
        return
    }
    m := (l + r) >> 1
    t.build(xs, o<<1, l, m)
    t.build(xs, o<<1|1, m+1, r)
    t.maintain(o)
}

// 区间更新
func (t seg) update(o, l, r, v int) {
    if l <= t[o].l && t[o].r <= r {
        t.do(o, v)
        return
    }
    t.spread(o)
    m := (t[o].l + t[o].r) >> 1
    if l <= m {
        t.update(o<<1, l, r, v)
    }
    if m < r {
        t.update(o<<1|1, l, r, v)
    }
    t.maintain(o)
}

// 代码逻辑同 850 题，增加一个 records 数组记录关键数据
func separateSquares(squares [][]int) float64 {
    m := len(squares) * 2
    xs := make([]int, 0, m)
    type event struct{ y, lx, rx, delta int }
    events := make([]event, 0, m)
    for _, sq := range squares {
        lx, y, l := sq[0], sq[1], sq[2]
        rx := lx + l
        xs = append(xs, lx, rx)
        events = append(events, event{y, lx, rx, 1}, event{y + l, lx, rx, -1})
    }

    // 排序去重，方便离散化
    slices.Sort(xs)
    xs = slices.Compact(xs)

    // 初始化线段树
    n := len(xs) - 1 // len(xs) 个横坐标有 len(xs)-1 个差值
    t := make(seg, 2<<bits.Len(uint(n-1)))
    t.build(xs, 1, 0, n-1)

    // 模拟扫描线从下往上移动
    slices.SortFunc(events, func(a, b event) int { return a.y - b.y })
    type pair struct{ area, sumLen int }
    records := make([]pair, m-1)
    totArea := 0
    for i, e := range events[:m-1] {
        l := sort.SearchInts(xs, e.lx)
        r := sort.SearchInts(xs, e.rx) - 1 // 注意 r 对应着 xs[r] 与 xs[r+1]=e.rx 的差值
        t.update(1, l, r, e.delta)         // 更新被 [e.lx, e.rx] 覆盖的次数
        sumLen := xs[len(xs)-1] - xs[0]    // 总的底边长度
        if t[1].minCover == 0 {            // 需要去掉没被矩形覆盖的长度
            sumLen -= t[1].minCoverLen
        }
        records[i] = pair{totArea, sumLen} // 记录关键数据
        totArea += sumLen * (events[i+1].y - e.y) // 新增面积 = 被至少一个矩形覆盖的底边长之和 * 矩形高度
    }

    // 二分找最后一个 < totArea / 2 的面积
    i := sort.Search(m-1, func(i int) bool { return records[i].area*2 >= totArea }) - 1
    return float64(events[i].y) + float64(totArea-records[i].area*2)/float64(records[i].sumLen*2)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{squares}$ 的长度。
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