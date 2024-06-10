#### 题目

<p>给你一个下标从 <strong>0</strong>&nbsp;开始的正整数数组&nbsp;<code>heights</code>&nbsp;，其中&nbsp;<code>heights[i]</code>&nbsp;表示第 <code>i</code>&nbsp;栋建筑的高度。</p>

<p>如果一个人在建筑&nbsp;<code>i</code>&nbsp;，且存在&nbsp;<code>i &lt; j</code>&nbsp;的建筑&nbsp;<code>j</code>&nbsp;满足&nbsp;<code>heights[i] &lt; heights[j]</code>&nbsp;，那么这个人可以移动到建筑&nbsp;<code>j</code>&nbsp;。</p>

<p>给你另外一个数组&nbsp;<code>queries</code>&nbsp;，其中&nbsp;<code>queries[i] = [a<sub>i</sub>, b<sub>i</sub>]</code>&nbsp;。第&nbsp;<code>i</code>&nbsp;个查询中，Alice 在建筑&nbsp;<code>a<sub>i</sub></code> ，Bob 在建筑&nbsp;<code>b<sub>i</sub></code><sub>&nbsp;</sub>。</p>

<p>请你能返回一个数组&nbsp;<code>ans</code>&nbsp;，其中&nbsp;<code>ans[i]</code>&nbsp;是第&nbsp;<code>i</code>&nbsp;个查询中，Alice 和 Bob 可以相遇的&nbsp;<strong>最左边的建筑</strong>&nbsp;。如果对于查询&nbsp;<code>i</code>&nbsp;，Alice<em> </em>和<em> </em>Bob 不能相遇，令&nbsp;<code>ans[i]</code> 为&nbsp;<code>-1</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>heights = [6,4,8,5,2,7], queries = [[0,1],[0,3],[2,4],[3,4],[2,2]]
<b>输出：</b>[2,5,-1,5,2]
<b>解释：</b>第一个查询中，Alice 和 Bob 可以移动到建筑 2 ，因为 heights[0] &lt; heights[2] 且 heights[1] &lt; heights[2] 。
第二个查询中，Alice 和 Bob 可以移动到建筑 5 ，因为 heights[0] &lt; heights[5] 且 heights[3] &lt; heights[5] 。
第三个查询中，Alice 无法与 Bob 相遇，因为 Alice 不能移动到任何其他建筑。
第四个查询中，Alice 和 Bob 可以移动到建筑 5 ，因为 heights[3] &lt; heights[5] 且 heights[4] &lt; heights[5] 。
第五个查询中，Alice 和 Bob 已经在同一栋建筑中。
对于 ans[i] != -1 ，ans[i] 是 Alice 和 Bob 可以相遇的建筑中最左边建筑的下标。
对于 ans[i] == -1 ，不存在 Alice 和 Bob 可以相遇的建筑。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>heights = [5,3,8,2,6,1,4,6], queries = [[0,7],[3,5],[5,2],[3,0],[1,6]]
<b>输出：</b>[7,6,-1,4,6]
<strong>解释：</strong>第一个查询中，Alice 可以直接移动到 Bob 的建筑，因为 heights[0] &lt; heights[7] 。
第二个查询中，Alice 和 Bob 可以移动到建筑 6 ，因为 heights[3] &lt; heights[6] 且 heights[5] &lt; heights[6] 。
第三个查询中，Alice 无法与 Bob 相遇，因为 Bob 不能移动到任何其他建筑。
第四个查询中，Alice 和 Bob 可以移动到建筑 4 ，因为 heights[3] &lt; heights[4] 且 heights[0] &lt; heights[4] 。
第五个查询中，Alice 可以直接移动到 Bob 的建筑，因为 heights[1] &lt; heights[6] 。
对于 ans[i] != -1 ，ans[i] 是 Alice 和 Bob 可以相遇的建筑中最左边建筑的下标。
对于 ans[i] == -1 ，不存在 Alice 和 Bob 可以相遇的建筑。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= heights.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>1 &lt;= heights[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>queries[i] = [a<sub>i</sub>, b<sub>i</sub>]</code></li>
	<li><code>0 &lt;= a<sub>i</sub>, b<sub>i</sub> &lt;= heights.length - 1</code></li>
</ul>

#### 思路


## 方法一：离线做法+最小堆

> 离线：按照自己定义的某种顺序回答询问（而不是按照输入顺序一个个地回答）。

不妨设 $a_i \le b_i$。

如果 $a_i = b_i$ 或者 $\textit{heights}[a_i] <\textit{heights}[b_i]$，那么 Alice 可以直接跳到 Bob 的位置，即 $\textit{ans}[i] = b_i$。

否则，我们可以在 $b_i$ 处**记录**「左边有个 $a_i$，它属于第 $i$ 个询问」。

然后遍历 $\textit{heights}$，同时用一个**最小堆**维护上面说的「记录」：遍历到 $\textit{heights}[i]$ 时，把在下标 $i$ 处的「记录」全部加到最小堆中。

在加到最小堆之前，我们可以回答左边所有高度小于 $\textit{heights}[i]$ 的「记录」，其答案就是 $i$。

#### 总结

这个算法涉及到三个位置，按照**从左到右**的顺序，它们分别是：

1. $a_i$：回答询问时，用它的高度来和当前高度判断。
2. $b_i$：决定了在什么位置把询问加入堆中。
3. 回答询问的位置。

```go  
func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i := range ans {
		ans[i] = -1
	}
	left := make([][]pair, len(heights))
	for qi, q := range queries {
		i, j := q[0], q[1]
		if i > j {
			i, j = j, i // 保证 i <= j
		}
		if i == j || heights[i] < heights[j] {
			ans[qi] = j // i 直接跳到 j
		} else {
			left[j] = append(left[j], pair{heights[i], qi}) // 离线
		}
	}

	h := hp{}
	for i, x := range heights { // 从小到大枚举下标 i
		for h.Len() > 0 && h[0].h < x {
			ans[heap.Pop(&h).(pair).qi] = i // 可以跳到 i（此时 i 是最小的）
		}
		for _, p := range left[i] {
			heap.Push(&h, p) // 后面再回答
		}
	}
	return ans
}

type pair struct{ h, qi int }
type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].h < h[j].h }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log q)$，其中 $n$ 为 $\textit{heights}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n + q)$。


## 方法二：在线做法+线段树二分

构建一棵维护区间**最大值** $\textit{mx}$ 的线段树。

方法一中用堆回答的询问，相当于问区间 $[j+1,n-1]$ 中的大于 $v = \textit{heights}[i]$ 的最小下标。

由于代码中线段树的下标是从 $1$ 开始的，所以区间是 $[j+2,n]$。不过为了避免讨论 $j+2>n$ 的情况，代码中用的 $j+1$。

- 如果当前区间 $\textit{mx}\le v$，则整个区间都不存在大于 $v$ 的数，返回 $0$。
- 如果当前区间只包含一个元素，则找到答案，返回该元素的下标。
- 如果左子树包含下标 $j+1$，则递归左子树。
- 如果左子树返回 $0$，则返回递归右子树的结果。


```go  
type seg []int

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o] = max(t[o<<1], t[o<<1|1])
}

func (t seg) query(o, l, r, L, v int) int {
	if v >= t[o] {
		return 0
	}
	if l == r {
		return l
	}
	m := (l + r) >> 1
	if L <= m {
		pos := t.query(o<<1, l, m, L, v)
		if pos > 0 {
			return pos
		}
	}
	return t.query(o<<1|1, m+1, r, L, v)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func leftmostBuildingQueriesSeg(heights []int, queries [][]int) []int {
	n := len(heights)
	t := make(seg, n*4)
	t.build(heights, 1, 1, n)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		i, j := q[0], q[1]
		if i > j {
			i, j = j, i
		}
		if i == j || heights[i] < heights[j] {
			ans[qi] = j
		} else {
			pos := t.query(1, 1, n, j+1, heights[i])
			ans[qi] = pos - 1 // 不存在时刚好得到 -1
		}
	}
	return ans

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 为 $\textit{heights}$ 的长度，$q$ 为 $\textit{queries}$ 的长度。对于左子树的递归，时间是 $\mathcal{O}(\log n)$ 的（同单点更新）；对于右子树的递归，由于区间满足 $\textit{max}\le v$ 则不递归，否则只会向下递归，所以这部分的时间也是 $\mathcal{O}(\log n)$ 的，所以线段树二分的时间复杂度为 $\mathcal{O}(\log n)$。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。
