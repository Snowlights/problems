#### 题目  

<p>你有一台电脑，它可以 <strong>同时</strong> 运行无数个任务。给你一个二维整数数组 <code>tasks</code> ，其中 <code>tasks[i] = [start<sub>i</sub>, end<sub>i</sub>, duration<sub>i</sub>]</code> 表示第 <code>i</code> 个任务需要在 <strong>闭区间</strong> 时间段 <code>[start<sub>i</sub>, end<sub>i</sub>]</code> 内运行 <code>duration<sub>i</sub></code> 个整数时间点（但不需要连续）。</p>

<p>当电脑需要运行任务时，你可以打开电脑，如果空闲时，你可以将电脑关闭。</p>

<p>请你返回完成所有任务的情况下，电脑最少需要运行多少秒。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>tasks = [[2,3,1],[4,5,1],[1,5,2]]
<b>输出：</b>2
<b>解释：</b>
- 第一个任务在闭区间 [2, 2] 运行。
- 第二个任务在闭区间 [5, 5] 运行。
- 第三个任务在闭区间 [2, 2] 和 [5, 5] 运行。
电脑总共运行 2 个整数时间点。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>tasks = [[1,3,2],[2,5,3],[5,6,2]]
<b>输出：</b>4
<b>解释：</b>
- 第一个任务在闭区间 [2, 3] 运行
- 第二个任务在闭区间 [2, 3] 和 [5, 5] 运行。
- 第三个任务在闭区间 [5, 6] 运行。
电脑总共运行 4 个整数时间点。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= tasks.length &lt;= 2000</code></li>
	<li><code>tasks[i].length == 3</code></li>
	<li><code>1 &lt;= start<sub>i</sub>, end<sub>i</sub> &lt;= 2000</code></li>
	<li><code>1 &lt;= duration<sub>i</sub> &lt;= end<sub>i</sub> - start<sub>i</sub> + 1 </code></li>
</ul>
 
#### 思路  

# 方法一：贪心+暴力

#### 提示 1

按照右端点排序。

#### 提示 2

对于 $\textit{tasks}[i]$ 来说，它右侧的任务要么和它没有交集，要么包含它的区间**后缀**。

#### 提示 3

遍历排序后的任务，先统计区间内的已有的电脑运行时间点，如果个数小于 $\textit{duration}$，则需要新增时间点。
根据提示 2，尽量把新增的时间点安排在区间 $[\textit{start},\textit{end}]$ 的后缀上，这样下一个区间就能统计到更多已有的时间点。

```go 
func findMinimumTime(tasks [][]int) (ans int) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	run := make([]bool, tasks[len(tasks)-1][1]+1)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		for _, b := range run[start : end+1] { // 去掉运行中的时间点
			if b {
				d--
			}
		}
		for i := end; d > 0; i-- { // 剩余的 d 填充区间后缀
			if !run[i] {
				run[i] = true
				d--
				ans++
			}
		}
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$O(nU)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{end}_i)$。
- 空间复杂度：$O(U)$。

# 方法二：贪心+线段树优化

在方法一的暴力更新上优化。  
由于有区间更新操作，需要用 lazy 线段树，对于本题，在更新的时候需要优先递归右子树，从而保证是从右往左更新。其余细节见代码注释。
> 注：由于线段树常数比较大，在数据范围只有几百几千的情况下，不一定比方法一的暴力快。

```go  
type seg []struct {
	l, r, cnt int
	todo      bool
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(i int) {
	o := &t[i]
	o.cnt = o.r - o.l + 1
	o.todo = true
}

func (t seg) spread(o int) {
	if t[o].todo {
		t[o].todo = false
		t.do(o << 1)
		t.do(o<<1 | 1)
	}
}

// 查询区间 [l,r]   o=1
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].cnt
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

// 新增区间 [l,r] 后缀的 suffix 个时间点   o=1
// 相当于把线段树二分和线段树更新合并成了一个函数，时间复杂度为 O(log u)
func (t seg) update(o, l, r int, suffix *int) {
	size := t[o].r - t[o].l + 1
	if t[o].cnt == size { // 全部为运行中
		return
	}
	if l <= t[o].l && t[o].r <= r && size-t[o].cnt <= *suffix { // 整个区间全部改为运行中
		*suffix -= size - t[o].cnt
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r > m { // 先更新右子树
		t.update(o<<1|1, l, r, suffix)
	}
	if *suffix > 0 { // 再更新左子树（如果还有需要新增的时间点）
		t.update(o<<1, l, r, suffix)
	}
	t[o].cnt = t[o<<1].cnt + t[o<<1|1].cnt
}

func findMinimumTime(tasks [][]int) (ans int) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	u := tasks[len(tasks)-1][1]
	st := make(seg, u*4)
	st.build(1, 1, u)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		d -= st.query(1, start, end) // 去掉运行中的时间点
		if d > 0 {
			st.update(1, start, end, &d) // 新增时间点
		}
	}
	return st[1].cnt
}
```


#### 复杂度分析

- 时间复杂度：$O(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{end}_i)$。
- 空间复杂度：$O(U)$。
**注**：如果改用动态开点线段树，可以做到 $O(n\log n)$ 时间和 $O(n\log n)$ 空间。

# 方法三：贪心+栈优化+二分查找

由于每次都是从右到左新增时间点，相当于把若干右侧的区间合并成一个大区间，因此可以用栈来优化。  
栈中保存闭区间的左右端点，以及从栈底到栈顶的区间长度之和（类似前缀和）。  
合并前先在栈中二分查找左端点所在区间，由于我们保存了长度之和，所以可以算出 $[\textit{start},\textit{end}]$ 范围内的运行中的时间点。  
如果还需要新增时间点，那么就从右到左合并，具体细节见代码。

```go  
func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	type tuple struct{ l, r, s int }
	st := []tuple{{-2, -2, 0}} // 闭区间左右端点，栈底到栈顶的区间长度的和
	for _, p := range tasks {
		start, end, d := p[0], p[1], p[2]
		i := sort.Search(len(st), func(i int) bool { return st[i].l >= start }) - 1
		d -= st[len(st)-1].s - st[i].s // 去掉运行中的时间点
		if start <= st[i].r { // start 在区间 st[i] 内
			d -= st[i].r - start + 1 // 去掉运行中的时间点
		}
		if d <= 0 {
			continue
		}
		for end-st[len(st)-1].r <= d { // 剩余的 d 填充区间后缀
			top := st[len(st)-1]
			st = st[:len(st)-1]
			d += top.r - top.l + 1 // 合并区间
		}
		st = append(st, tuple{end - d + 1, end, st[len(st)-1].s + d})
	}
	return st[len(st)-1].s
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$