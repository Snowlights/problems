#### 题目

<p>你有 <code>n</code> 颗处理器，每颗处理器都有 <code>4</code> 个核心。现有 <code>n * 4</code> 个待执行任务，每个核心只执行 <strong>一个</strong> 任务。</p>

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>processorTime</code> ，表示每颗处理器最早空闲时间。另给你一个下标从 <strong>0</strong> 开始的整数数组 <code>tasks</code> ，表示执行每个任务所需的时间。返回所有任务都执行完毕需要的 <strong>最小时间</strong> 。</p>

<p>注意：每个核心独立执行任务。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>processorTime = [8,10], tasks = [2,2,3,1,8,7,4,5]
<strong>输出：</strong>16
<strong>解释：</strong>
最优的方案是将下标为 4, 5, 6, 7 的任务分配给第一颗处理器（最早空闲时间 time = 8），下标为 0, 1, 2, 3 的任务分配给第二颗处理器（最早空闲时间 time = 10）。 
第一颗处理器执行完所有任务需要花费的时间 = max(8 + 8, 8 + 7, 8 + 4, 8 + 5) = 16 。
第二颗处理器执行完所有任务需要花费的时间 = max(10 + 2, 10 + 2, 10 + 3, 10 + 1) = 13 。
因此，可以证明执行完所有任务需要花费的最小时间是 16 。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>processorTime = [10,20], tasks = [2,3,1,2,5,8,4,3]
<strong>输出：</strong>23
<strong>解释：</strong>
最优的方案是将下标为 1, 4, 5, 6 的任务分配给第一颗处理器（最早空闲时间 time = 10），下标为 0, 2, 3, 7 的任务分配给第二颗处理器（最早空闲时间 time = 20）。 
第一颗处理器执行完所有任务需要花费的时间 = max(10 + 3, 10 + 5, 10 + 8, 10 + 4) = 18 。 
第二颗处理器执行完所有任务需要花费的时间 = max(20 + 2, 20 + 1, 20 + 2, 20 + 3) = 23 。 
因此，可以证明执行完所有任务需要花费的最小时间是 23 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n == processorTime.length <= 25000</code></li>
	<li><code>1 <= tasks.length <= 10<sup>5</sup></code></li>
	<li><code>0 <= processorTime[i] <= 10<sup>9</sup></code></li>
	<li><code>1 <= tasks[i] <= 10<sup>9</sup></code></li>
	<li><code>tasks.length == 4 * n</code></li>
</ul>

#### 思路

一颗处理器完成它的 $4$ 个任务，完成的时间取决于这 $4$ 个任务中的 $\textit{tasks}$ 的最大值。
直觉上来说，最早空闲时间越大的处理器，处理 $\textit{tasks}$ 越小的任务，那么完成时间越早。
证明：对于两个最早空闲时间分别为 $p_1$ 和 $p_2$ 的处理器，不妨设 $p_1 \le p_2$。完成的 $4$ 个任务中的最大值分别为 $t_1$ 和 $t_2$，不妨设 $t_1 \le t_2$。
如果 $t_1$ 给 $p_1$，$t_2$ 给 $p_2$，那么最后完成时间为

$$
\max(p_1+t_1, p_2+t_2) = p_2+t_2

$$

如果 $t_1$ 给 $p_2$，$t_2$ 给 $p_1$，那么最后完成时间为

$$
\max(p_1+t_2, p_2+t_1) \le \max(p_2+t_2, p_2+t_2) = p_2+t_2

$$

上式表明，最早空闲时间越大的处理器，处理 $\textit{tasks}$ 越小的任务，那么完成时间不会变的更晚。
我们可以把 $\textit{processorTime}$ 从小到大排序，$\textit{tasks}$ 从大到小排序，那么答案就是

$$
\textit{processorTime}[i] + \textit{tasks}[4i]

$$

的最大值。

```go  
func minProcessingTime(processorTime []int, tasks []int) (ans int) {
	sort.Ints(processorTime)
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	for i, v := range processorTime {
		ans = max(ans, v+tasks[i*4])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{processorTime}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。Python 忽略切片开销。
