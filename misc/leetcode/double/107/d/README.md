### 题目

<p>给你一个整数 <code>n</code> ，表示服务器的总数目，再给你一个下标从 <strong>0</strong> 开始的 <strong>二维</strong> 整数数组 <code>logs</code> ，其中 <code>logs[i] = [server_id, time]</code> 表示 id 为 <code>server_id</code> 的服务器在 <code>time</code> 时收到了一个请求。</p>

<p>同时给你一个整数 <code>x</code> 和一个下标从 <strong>0</strong> 开始的整数数组 <code>queries</code>  。</p>

<p>请你返回一个长度等于 <code>queries.length</code> 的数组 <code>arr</code> ，其中 <code>arr[i]</code> 表示在时间区间 <code>[queries[i] - x, queries[i]]</code> 内没有收到请求的服务器数目。</p>

<p>注意时间区间是个闭区间。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>n = 3, logs = [[1,3],[2,6],[1,5]], x = 5, queries = [10,11]
<b>输出：</b>[1,2]
<b>解释：</b>
对于 queries[0]：id 为 1 和 2 的服务器在区间 [5, 10] 内收到了请求，所以只有服务器 3 没有收到请求。
对于 queries[1]：id 为 2 的服务器在区间 [6,11] 内收到了请求，所以 id 为 1 和 3 的服务器在这个时间段内没有收到请求。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>n = 3, logs = [[2,4],[2,1],[1,2],[3,1]], x = 2, queries = [3,4]
<b>输出：</b>[0,1]
<b>解释：</b>
对于 queries[0]：区间 [1, 3] 内所有服务器都收到了请求。
对于 queries[1]：只有 id 为 3 的服务器在区间 [2,4] 内没有收到请求。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n <= 10<sup>5</sup></code></li>
	<li><code>1 <= logs.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= queries.length <= 10<sup>5</sup></code></li>
	<li><code>logs[i].length == 2</code></li>
	<li><code>1 <= logs[i][0] <= n</code></li>
	<li><code>1 <= logs[i][1] <= 10<sup>6</sup></code></li>
	<li><code>1 <= x <= 10<sup>5</sup></code></li>
	<li><code>x < queries[i] <= 10<sup>6</sup></code></li>
</ul>

### 思路

将日志和查询按照时间排序,之后遍历查询的时间，滑动窗口统计在范围内的服务器个数，离线回答。

```go  
func countServers(n int, logs [][]int, x int, qs []int) (ans []int) {

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	type q struct {
		i, t int
	}
	qList := make([]*q, 0, len(qs))
	for i, v := range qs {
		qList = append(qList, &q{i, v})
	}
	sort.Slice(qList, func(i, j int) bool {
		return qList[i].t < qList[j].t
	})

	ans = make([]int, len(qs))
	h := make(map[int]int)
	l, r := 0, 0

	for _, q := range qList {
		for r < len(logs) && logs[r][1] <= q.t {
			h[logs[r][0]]++
			r++
		}
		for l < r && logs[l][1] < q.t-x {
			h[logs[l][0]]--
			if h[logs[l][0]] == 0 {
				delete(h, logs[l][0])
			}
			l++
		}
		ans[q.i] = n - len(h)
	}

	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(m+n+n\log (n)+m\log (m))$, $n$ 为 $\textit{logs}$ 的长度, $m$ 为 $\textit{qs}$ 的长度, 瓶颈在排序
- 空间复杂度：$\mathcal{O}(n)$ 。
