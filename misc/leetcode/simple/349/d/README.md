#### 题目  

<p>给你两个长度为 <code>n</code> 、下标从 <strong>0</strong> 开始的整数数组 <code>nums1</code> 和 <code>nums2</code> ，另给你一个下标从 <strong>1</strong> 开始的二维数组 <code>queries</code> ，其中 <code>queries[i] = [x<sub>i</sub>, y<sub>i</sub>]</code> 。</p>

<p>对于第 <code>i</code> 个查询，在所有满足 <code>nums1[j] &gt;= x<sub>i</sub></code> 且 <code>nums2[j] &gt;= y<sub>i</sub></code> 的下标 <code>j</code> <code>(0 &lt;= j &lt; n)</code> 中，找出 <code>nums1[j] + nums2[j]</code> 的 <strong>最大值</strong> ，如果不存在满足条件的 <code>j</code> 则返回 <strong>-1</strong> 。</p>

<p>返回数组<em> </em><code>answer</code><em> ，</em>其中<em> </em><code>answer[i]</code><em> </em>是第 <code>i</code> 个查询的答案。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums1 = [4,3,1,2], nums2 = [2,4,9,5], queries = [[4,1],[1,3],[2,5]]
<strong>输出：</strong>[6,10,7]
<strong>解释：</strong>
对于第 1 个查询：<code>x<sub>i</sub> = 4</code> 且 <code>y<sub>i</sub> = 1</code> ，可以选择下标 <code>j = 0</code> ，此时 <code>nums1[j] &gt;= 4</code> 且 <code>nums2[j] &gt;= 1</code> 。<code>nums1[j] + nums2[j]</code> 等于 6 ，可以证明 6 是可以获得的最大值。
对于第 2 个查询：<code>x<sub>i</sub> = 1</code> 且 <code>y<sub>i</sub> = 3</code> ，可以选择下标 <code>j = 2</code> ，此时 <code>nums1[j] &gt;= 1</code> 且 <code>nums2[j] &gt;= 3</code> 。<code>nums1[j] + nums2[j]</code> 等于 10 ，可以证明 10 是可以获得的最大值。
对于第 3 个查询：<code>x<sub>i</sub> = 2</code> 且 <code>y<sub>i</sub> = 5</code> ，可以选择下标 <code>j = 3</code> ，此时 <code>nums1[j] &gt;= 2</code> 且 <code>nums2[j] &gt;= 5</code> 。<code>nums1[j] + nums2[j]</code> 等于 7 ，可以证明 7 是可以获得的最大值。
因此，我们返回 <code>[6,10,7]</code> 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums1 = [3,2,5], nums2 = [2,3,4], queries = [[4,4],[3,2],[1,1]]
<strong>输出：</strong>[9,9,9]
<strong>解释：</strong>对于这个示例，我们可以选择下标 <code>j = 2</code> ，该下标可以满足每个查询的限制。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>nums1 = [2,1], nums2 = [2,3], queries = [[3,3]]
<strong>输出：</strong>[-1]
<strong>解释：</strong>示例中的查询 <code>x<sub>i</sub></code> = 3 且 <code>y<sub>i</sub></code> = 3 。对于每个下标 j ，都只满足 nums1[j] &lt; <code>x<sub>i</sub></code> 或者 nums2[j] &lt; <code>y<sub>i</sub></code> 。因此，不存在答案。 
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>nums1.length == nums2.length</code> </li>
	<li><code>n == nums1.length </code></li>
	<li><code>1 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums1[i], nums2[i] &lt;= 10<sup>9</sup> </code></li>
	<li><code>1 &lt;= queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>x<sub>i</sub> == queries[i][1]</code></li>
	<li><code>y<sub>i</sub> == queries[i][2]</code></li>
	<li><code>1 &lt;= x<sub>i</sub>, y<sub>i</sub> &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

为方便处理，可以先把 $\textit{nums}_1$ 和询问中的 $x_i$ 排序。

这样就可以把重心放在 $\textit{nums}_2[j]$ 与 $y_i$ 的大小关系上。

我们可以按照 $x_i$ 从大到小、$\textit{nums}_1[j]$ 从大到小的顺序处理，同时**增量地**维护满足 $\textit{nums}_1[j]\ge x_i$ 的 $\textit{nums}_2[j]$。

如何维护？分类讨论：

- 如果 $\textit{nums}_2[j]$ 比之前遍历过的 $\textit{nums}_2[j']$ 还要小，那么由于 $\textit{nums}_1[j]$ 是从大到小处理的，所以 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 也比之前遍历过的 $\textit{nums}_1[j']+\textit{nums}_2[j']$ 要小。那么在回答 $\le \textit{nums}_2[j]$ 的 $y_i$ 时，最大值不可能是 $\textit{nums}_1[j]+\textit{nums}_2[j]$，所以无需考虑这样的 $\textit{nums}_2[j]$。（这种单调性启发我们用**单调栈**来维护。）
- 如果是相等，那么同理，无需考虑。
- 如果是大于，那么就可以入栈。在入栈前还要去掉一些无效数据：
    - 如果 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 不低于栈顶的 $\textit{nums}_1[j']+\textit{nums}_2[j']$，那么可以弹出栈顶。因为更大的 $\textit{nums}_2[j]$ 更能满足 $\ge y_i$ 的要求，栈顶的 $\textit{nums}_1[j']+\textit{nums}_2[j']$ 在后续的询问中，永远不会是最大值。
    - 代码实现时，可以直接比较 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 与栈顶的值，这是因为如果这一条件成立，由于 $\textit{nums}_1[j]$ 是从大到小处理的，$\textit{nums}_1[j]+\textit{nums}_2[j]$ 能比栈顶的大，说明 $\textit{nums}_2[j]$ 必然不低于栈顶的 $\textit{nums}_2[j']$。

这样我们会得到一个从栈底到栈顶，$\textit{nums}_2[j]$ 递增，$\textit{nums}_1[j]+\textit{nums}_2[j]$ 递减的单调栈。

最后在单调栈中二分 $\ge y_i$ 的最小的 $\textit{nums}_2[j]$，对应的 $\textit{nums}_1[j]+\textit{nums}_2[j]$ 就是最大的。

```go 
func maximumSumQueries(a []int, b []int, qs [][]int) (ans []int) {
	type node struct {
		x, y int
	}
	nodeList := make([]*node, 0, len(a))
	for i, v := range a {
		nodeList = append(nodeList, &node{v, b[i]})
	}
	for i := range qs {
		qs[i] = append(qs[i], i)
	}
	ans = make([]int, len(qs))

	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].x < nodeList[j].x
	})
	sort.Slice(qs, func(i, j int) bool {
		return qs[i][0] > qs[j][0]
	})

	q, i := []*node{}, len(a)-1
	for _, qu := range qs {
		for i >= 0 && nodeList[i].x >= qu[0] {
			for len(q) > 0 && q[len(q)-1].y < nodeList[i].x+nodeList[i].y {
				q = q[:len(q)-1]
			}
			if len(q) == 0 || q[len(q)-1].x < nodeList[i].y {
				q = append(q, &node{nodeList[i].y, nodeList[i].x + nodeList[i].y})
			}
			i--
		}
		j := sort.Search(len(q), func(i int) bool {
			return q[i].x >= qu[1]
		})
		if j < len(q) {
			ans[qu[2]] = q[j].y
		} else {
			ans[qu[2]] = -1
		}
	}

	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n + q\log n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计。