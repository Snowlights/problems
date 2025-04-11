### 题目

<p>一个任务管理器系统可以让用户管理他们的任务，每个任务有一个优先级。这个系统需要高效地处理添加、修改、执行和删除任务的操作。</p>

<p>请你设计一个&nbsp;<code>TaskManager</code>&nbsp;类：</p>

<ul>
	<li>
	<p><code>TaskManager(vector&lt;vector&lt;int&gt;&gt;&amp; tasks)</code>&nbsp;初始化任务管理器，初始化的数组格式为&nbsp;<code>[userId, taskId, priority]</code>&nbsp;，表示给 <code>userId</code>&nbsp;添加一个优先级为 <code>priority</code>&nbsp;的任务 <code>taskId</code>&nbsp;。</p>
	</li>
	<li>
	<p><code>void add(int userId, int taskId, int priority)</code>&nbsp;表示给用户 <code>userId</code>&nbsp;添加一个优先级为 <code>priority</code>&nbsp;的任务 <code>taskId</code>&nbsp;，输入 <strong>保证&nbsp;</strong><code>taskId</code>&nbsp;不在系统中。</p>
	</li>
	<li>
	<p><code>void edit(int taskId, int newPriority)</code>&nbsp;更新已经存在的任务&nbsp;<code>taskId</code>&nbsp;的优先级为&nbsp;<code>newPriority</code>&nbsp;。输入 <strong>保证</strong>&nbsp;<code>taskId</code>&nbsp;存在于系统中。</p>
	</li>
	<li>
	<p><code>void rmv(int taskId)</code>&nbsp;从系统中删除任务&nbsp;<code>taskId</code>&nbsp;。输入 <strong>保证</strong>&nbsp;<code>taskId</code>&nbsp;存在于系统中。</p>
	</li>
	<li>
	<p><code>int execTop()</code>&nbsp;执行所有用户的任务中优先级 <strong>最高</strong>&nbsp;的任务，如果有多个任务优先级相同且都为 <strong>最高</strong>&nbsp;，执行&nbsp;<code>taskId</code>&nbsp;最大的一个任务。执行完任务后，<code>taskId</code><strong>&nbsp;</strong>从系统中 <strong>删除</strong>&nbsp;。同时请你返回这个任务所属的用户&nbsp;<code>userId</code>&nbsp;。如果不存在任何任务，返回&nbsp;-1 。</p>
	</li>
</ul>

<p><strong>注意</strong> ，一个用户可能被安排多个任务。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><br />
<span class="example-io">["TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"]<br />
[[[[1, 101, 10], [2, 102, 20], [3, 103, 15]]], [4, 104, 5], [102, 8], [], [101], [5, 105, 15], []]</span></p>

<p><strong>输出：</strong><br />
<span class="example-io">[null, null, null, 3, null, null, 5] </span></p>

<p><strong>解释：</strong></p>
TaskManager taskManager = new TaskManager([[1, 101, 10], [2, 102, 20], [3, 103, 15]]); // 分别给用户 1 ，2 和 3 初始化一个任务。<br />
taskManager.add(4, 104, 5); // 给用户 4 添加优先级为 5 的任务 104 。<br />
taskManager.edit(102, 8); // 更新任务 102 的优先级为 8 。<br />
taskManager.execTop(); // 返回 3 。执行用户 3 的任务 103 。<br />
taskManager.rmv(101); // 将系统中的任务 101 删除。<br />
taskManager.add(5, 105, 15); // 给用户 5 添加优先级为 15 的任务 105 。<br />
taskManager.execTop(); // 返回 5 。执行用户 5 的任务 105 。</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= tasks.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= userId &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= taskId &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= priority &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= newPriority &lt;= 10<sup>9</sup></code></li>
	<li><code>add</code>&nbsp;，<code>edit</code>&nbsp;，<code>rmv</code>&nbsp;和&nbsp;<code>execTop</code>&nbsp;的总操作次数 <strong>加起来</strong>&nbsp;不超过&nbsp;<code>2 * 10<sup>5</sup></code> 次。</li>
</ul>

### 思路

用**最大堆**维护优先级、任务编号、用户编号。

同时额外用一个哈希表记录每个任务编号对应的**最新的**优先级和用户编号。

- 对于 $\texttt{edit}$，直接把一个新的优先级、任务编号、用户编号三元组入堆。同时更新哈希表。
- 对于 $\texttt{rmv}$，直接把元素从哈希表删掉（或者优先级改成 $-1$）。
- 对于 $\texttt{execTop}$，不断弹出「货不对板」的堆顶，也就是任务编号和哈希表中记录的数据不一致的堆顶。直到找到一致的数据或者堆为空。

> 注 1：如果你学过 Dijkstra 算法，其中的堆就是懒删除堆。
> 注 2：题目保证输入的 $\textit{tasks}$ 数组中没有重复的 $\textit{taskId}$，但没有明说。

```
type pair struct{ priority, userId int }

type TaskManager struct {
	h  *hp          // (priority, taskId, userId)
	mp map[int]pair // taskId -> (priority, userId)
}

func Constructor(tasks [][]int) TaskManager {
	h := hp{}
	mp := map[int]pair{}
	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		mp[taskId] = pair{priority, userId}
		h = append(h, tuple{priority, taskId, userId})
	}
	heap.Init(&h)
	return TaskManager{&h, mp}
}

func (tm *TaskManager) Add(userId, taskId, priority int) {
	tm.mp[taskId] = pair{priority, userId}
	heap.Push(tm.h, tuple{priority, taskId, userId})
}

func (tm *TaskManager) Edit(taskId, newPriority int) {
	// 懒修改
	tm.Add(tm.mp[taskId].userId, taskId, newPriority)
}

func (tm *TaskManager) Rmv(taskId int) {
	// 懒删除
	tm.mp[taskId] = pair{-1, -1}
}

func (tm *TaskManager) ExecTop() int {
	for tm.h.Len() > 0 {
		top := heap.Pop(tm.h).(tuple)
		priority, taskId, userId := top.priority, top.taskId, top.userId
		// 如果货不对板，堆顶和 mp 中记录的不一样，说明这个数据已被修改/删除，不做处理
		if tm.mp[taskId] == (pair{priority, userId}) {
			tm.Rmv(taskId)
			return userId
		}
	}
	return -1
}

type tuple struct{ priority, taskId, userId int }
type hp []tuple
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return cmp.Or(h[i].priority-h[j].priority, h[i].taskId-h[j].taskId) > 0 }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

#### 复杂度分析

- 时间复杂度：
    - 初始化：$\mathcal{O}(n)$ 或者 $\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{tasks}$ 的长度。Python 和 Go 使用了堆化，复杂度是 $\mathcal{O}(n)$ 的。
    - $\texttt{add}$ 和 $\texttt{edit}$：$\mathcal{O}(\log (n+q))$，其中 $q$ 是 $\texttt{add}$ 和 $\texttt{edit}$ 的调用次数之和。
    - $\texttt{rmv}$：$\mathcal{O}(1)$。
    - $\texttt{execTop}$：均摊 $\mathcal{O}(\log (n+q))$。每个元素至多入堆出堆各一次。
- 空间复杂度：$\mathcal{O}(n+q)$。


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