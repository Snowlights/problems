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

### 方法一：正序回答询问+有序集合+线段树

![b131d.png](https://pic.leetcode.cn/1716685098-OuZZol-b131d.png)

看示例 2：

- 首先，在 $x=7$ 处放置一个障碍物，这会产生一个长度为 $7$ 的没有障碍物的区域。
- 然后，在 $x=2$ 处放置一个障碍物，去掉原来的 $7$，产生两个长度分别为 $2,5$ 的没有障碍物的区域。

这里的区域长度 $2,5,7$ 应该保存在哪个位置，从而方便我们查询？

查询的范围是 $[0,x]$，这可以分为两个部分。以 $x=8$ 为例：

- 设 $\textit{pre}$ 是在 $x$ 左侧最近的障碍物的位置，这里 $\textit{pre}=7$。
- 查询「完整」的没有障碍物的区域，这有 $[0,2],[2,7]$ 两段。
- 查询「不完整」的没有障碍物的区域，也就是 $[7,8]$ 这一段。

如果区域的**右端点**在 $[0,x]$ 中，这个区域一定是「完整」的。

所以，**把区域的长度保存在区域的右端点处。**

设 $d[x]$ 为右端点为 $x$ 的区域的长度。例如示例 2：

- 一开始所有 $d[x]=0$。
- 首先，在 $x=7$ 处放置一个障碍物，现在 $d[7]=7$。
- 然后，在 $x=2$ 处放置一个障碍物，现在 $d[2]=2,\ d[7]=5$。

问题变成如何维护和查询 $d$ 数组，我们需要支持单点更新，区间查询，这可以用**线段树**解决。

此外，我们还需要知道离 $x$ 左右最近的障碍物的位置，这可以用平衡树维护。

```
type seg []int

// 把 i 处的值改成 val
func (t seg) update(o, l, r, i, val int) {
	if l == r {
		t[o] = val
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(o<<1, l, m, i, val)
	} else {
		t.update(o<<1|1, m+1, r, i, val)
	}
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 查询 [0,R] 中的最大值
func (t seg) query(o, l, r, R int) int {
	if r <= R {
		return t[o]
	}
	m := (l + r) >> 1
	if R <= m {
		return t.query(o<<1, l, m, R)
	}
	return max(t[o<<1], t.query(o<<1|1, m+1, r, R))
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	for _, q := range queries {
		m = max(m, q[1])
	}
	m++

	set := redblacktree.New[int, struct{}]()
	set.Put(0, struct{}{}) // 哨兵
	set.Put(m, struct{}{})
	t := make(seg, 2<<bits.Len(uint(m)))

	for _, q := range queries {
		x := q[1]
        pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			nxt, _ := set.Ceiling(x) // x 右侧最近障碍物的位置
			set.Put(x, struct{}{})
			t.update(1, 0, m, x, x-pre.Key)       // 更新 d[x] = x - pre
			t.update(1, 0, m, nxt.Key, nxt.Key-x) // 更新 d[nxt] = nxt - x
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.query(1, 0, m, pre.Key), x-pre.Key)
			ans = append(ans, maxGap >= q[2])
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(U + q\log U)$，其中 $q$ 是 $\textit{queries}$ 的长度，$U$ 是 $\textit{x}$ 的最大值。注意题目保证 $U\le 3q$。
- 空间复杂度：$\mathcal{O}(U)$。

注：如果要做到值域无关，可以用动态开点线段树。

```
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	set := redblacktree.New[int, struct{}]()
	set.Put(0, struct{}{})
	pos := []int{0}
	for _, q := range queries {
		m = max(m, q[1])
		if q[0] == 1 {
			set.Put(q[1], struct{}{})
			pos = append(pos, q[1])
		}
	}
	m++
	set.Put(m, struct{}{}) // 哨兵

	t := make(fenwick, m)
	slices.Sort(pos)
	for i := 1; i < len(pos); i++ {
		t.update(pos[i], pos[i]-pos[i-1])
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		x := q[1]
		pre, _ := set.Floor(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			set.Remove(x)
			nxt, _ := set.Ceiling(x) // x 右侧最近障碍物的位置
			t.update(nxt.Key, nxt.Key-pre.Key) // 更新 d[nxt] = nxt - pre
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.preMax(pre.Key), x-pre.Key)
			ans = append(ans, maxGap >= q[2])
		}
	}
	slices.Reverse(ans)
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(U + q\log U)$，其中 $q$ 是 $\textit{queries}$ 的长度，$U$ 是 $\textit{x}$ 的最大值。注意题目保证 $U\le 3q$。
- 空间复杂度：$\mathcal{O}(U)$。


### 方法二：倒序回答询问+有序集合+树状数组

倒序回答询问，在 $x$ 处添加障碍物就变成删除障碍物了。

设 $x$ 左右最近元素分别为 $\textit{pre}$ 和 $\textit{nxt}$。

删除后，需要把 $d[\textit{nxt}]$ 更新为 $\textit{nxt}-\textit{pre}$。注意这会让 $d[\textit{nxt}]$ 增加。

由于询问的是前缀，并且 $d$ 不会变小，所以可以用树状数组维护 $d$。


### 方法三：倒序回答询问+并查集+树状数组

用并查集实现方法二中有序集合的删除、查找前驱和查找后继。

需要维护两个并查集，分别记作 $\textit{left}$ 和 $\textit{right}$。

- 删除 $x$：把 $\textit{left}$ 中的 $x$ 指向 $x-1$，把 $\textit{right}$ 中的 $x$ 指向 $x+1$。如果要批量删除开区间 $(p,q)$ 中的元素 $x$，可以把 $\textit{left}$ 中的 $x$ 指向 $p$，把 $\textit{right}$ 中的 $x$ 指向 $q$。
- 查找 $x$ 的前驱：$x-1$ 在 $\textit{left}$ 中的代表元。
- 查找 $x$ 的后继：$x+1$ 在 $\textit{right}$ 中的代表元。如果 $x$ 已经删除，也可以查询 $x$ 在 $\textit{right}$ 中的代表元。

```
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = max(f[i], val)
	}
}

func (f fenwick) preMax(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res = max(res, f[i])
	}
	return res
}

type uf []int

func (f uf) find(x int) int {
	if f[x] != x {
		f[x] = f.find(f[x])
	}
	return f[x]
}

func getResults(queries [][]int) (ans []bool) {
	m := 0
	pos := []int{0}
	for _, q := range queries {
		m = max(m, q[1])
		if q[0] == 1 {
			pos = append(pos, q[1])
		}
	}
	m++

	left := make(uf, m+1)
	right := make(uf, m+1)
	for i := range left {
		left[i] = i
		right[i] = i
	}
	t := make(fenwick, m)
	slices.Sort(pos)
	for i := 1; i < len(pos); i++ {
		p, q := pos[i-1], pos[i]
		t.update(q, q-p)
		for j := p + 1; j < q; j++ {
			left[j] = p // 删除 j
			right[j] = q
		}
	}
	for j := pos[len(pos)-1] + 1; j < m; j++ {
		left[j] = pos[len(pos)-1] // 删除 j
		right[j] = m
	}

	for i := len(queries) - 1; i >= 0; i-- {
		q := queries[i]
		x := q[1]
		pre := left.find(x - 1) // x 左侧最近障碍物的位置
		if q[0] == 1 {
			left[x] = x - 1 // 删除 x
			right[x] = x + 1
			nxt := right.find(x)   // x 右侧最近障碍物的位置
			t.update(nxt, nxt-pre) // 更新 d[nxt] = nxt - pre
		} else {
			// 最大长度要么是 [0,pre] 中的最大 d，要么是 [pre,x] 这一段的长度
			maxGap := max(t.preMax(pre), x-pre)
			ans = append(ans, maxGap >= q[2])
		}
	}
	slices.Reverse(ans)
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(U + q\log U)$，其中 $q$ 是 $\textit{queries}$ 的长度，$U$ 是 $\textit{x}$ 的最大值。注意题目保证 $U\le 3q$。
- 空间复杂度：$\mathcal{O}(U)$。

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