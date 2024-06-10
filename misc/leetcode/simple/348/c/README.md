#### 题目  

<p>给你一个整数 <code>n</code> 和一个下标从 <strong>0</strong> 开始的 <strong>二维数组</strong> <code>queries</code> ，其中 <code>queries[i] = [type<sub>i</sub>, index<sub>i</sub>, val<sub>i</sub>]</code> 。</p>

<p>一开始，给你一个下标从 <strong>0</strong> 开始的 <code>n x n</code> 矩阵，所有元素均为 <code>0</code> 。每一个查询，你需要执行以下操作之一：</p>

<ul>
	<li>如果 <code>type<sub>i</sub> == 0</code> ，将第 <code>index<sub>i</sub></code> 行的元素全部修改为 <code>val<sub>i</sub></code> ，覆盖任何之前的值。</li>
	<li>如果 <code>type<sub>i</sub> == 1</code> ，将第 <code>index<sub>i</sub></code> 列的元素全部修改为 <code>val<sub>i</sub></code> ，覆盖任何之前的值。</li>
</ul>

<p>请你执行完所有查询以后，返回矩阵中所有整数的和。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/05/11/exm1.png" style="width: 681px; height: 161px;"/></p>

<pre><b>输入：</b>n = 3, queries = [[0,0,1],[1,2,2],[0,2,3],[1,0,4]]
<b>输出：</b>23
<b>解释：</b>上图展示了每个查询以后矩阵的值。所有操作执行完以后，矩阵元素之和为 23 。
</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2023/05/11/exm2.png" style="width: 681px; height: 331px;"/></p>

<pre><b>输入：</b>n = 3, queries = [[0,0,4],[0,1,2],[1,0,1],[0,2,3],[1,2,1]]
<b>输出：</b>17
<b>解释：</b>上图展示了每一个查询操作之后的矩阵。所有操作执行完以后，矩阵元素之和为 17 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>queries[i].length == 3</code></li>
	<li><code>0 &lt;= type<sub>i</sub> &lt;= 1</code></li>
	<li><code>0 &lt;= index<sub>i</sub> &lt; n</code></li>
	<li><code>0 &lt;= val<sub>i</sub> &lt;= 10<sup>5</sup></code></li>
</ul>
 
#### 思路  

- 如果对同一行反复操作，那么只有最后一次对这行的操作会计入答案。列同理。
- 倒序操作 queries
- 以行为例。如果 $queries[i]$ 操作的是行，那么需要知道  
  - 这一行之前有没有操作过（这里「之前」指大于 $i$ 的操作）。对此可以用哈希表 $\textit{visRow}$ 记录被操作过的行号，哈希表 $\textit{visCol}$ 记录被操作过的列号。
  - 这一行有多少列之前操作过，这就是 $\textit{visCol}$ 的长度 $m$。那么剩余可以填入的格子为 $n-m$，答案增加了 $(n-m)\cdot \textit{val}_i$。
    
这样可以做到 $\mathcal{O}(q)$ 的时间复杂度（与 $n$ 无关）
    

```go 
func matrixSumQueries(n int, qs [][]int) int64 {
	ans := 0
	vis := [2]map[int]bool{{}, {}}
	for i := len(qs) - 1; i >= 0; i-- {
		q := qs[i]
		t, idx, v := q[0], q[1], q[2]
		if !vis[t][idx] {
			ans += v * (n - len(vis[t^1]))
			vis[t][idx] = true
		}
	}

	return int64(ans)
}
```

```python
class Solution:
    def matrixSumQueries(self, n: int, queries: List[List[int]]) -> int:
        vis = [set, set]
        ans = 0
        for type, index, val in reversed(queries):
            if index not in vis[type]:
                ans += val * (n - len(vis[type ^ 1]))
                vis[type].add(index)
        return ans
```

#### 复杂度分析  
- 时间复杂度：$\mathcal{O}(q)$，其中 $q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(\min\{q,n\})$。哈希表中至多有 $\mathcal{O}(n)$ 个数