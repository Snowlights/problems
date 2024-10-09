#### 题目

<p>给你一个长度为 <code>n</code> 的二维整数数组 <code>items</code> 和一个整数 <code>k</code> 。</p>

<p><code>items[i] = [profit<sub>i</sub>, category<sub>i</sub>]</code>，其中 <code>profit<sub>i</sub></code> 和 <code>category<sub>i</sub></code> 分别表示第 <code>i</code> 个项目的利润和类别。</p>

<p>现定义 <code>items</code> 的 <strong>子序列</strong> 的 <strong>优雅度</strong> 可以用 <code>total_profit + distinct_categories<sup>2</sup></code> 计算，其中 <code>total_profit</code> 是子序列中所有项目的利润总和，<code>distinct_categories</code> 是所选子序列所含的所有类别中不同类别的数量。</p>

<p>你的任务是从 <code>items</code> 所有长度为 <code>k</code> 的子序列中，找出 <strong>最大优雅度</strong> 。</p>

<p>用整数形式表示并返回 <code>items</code> 中所有长度恰好为 <code>k</code> 的子序列的最大优雅度。</p>

<p><strong>注意：</strong>数组的子序列是经由原数组删除一些元素（可能不删除）而产生的新数组，且删除不改变其余元素相对顺序。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>items = [[3,2],[5,1],[10,1]], k = 2
<strong>输出：</strong>17
<strong>解释：
</strong>在这个例子中，我们需要选出长度为 2 的子序列。
其中一种方案是 items[0] = [3,2] 和 items[2] = [10,1] 。
子序列的总利润为 3 + 10 = 13 ，子序列包含 2 种不同类别 [2,1] 。
因此，优雅度为 13 + 2<sup>2</sup> = 17 ，可以证明 17 是可以获得的最大优雅度。 
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>items = [[3,1],[3,1],[2,2],[5,3]], k = 3
<strong>输出：</strong>19
<strong>解释：</strong>
在这个例子中，我们需要选出长度为 3 的子序列。 
其中一种方案是 items[0] = [3,1] ，items[2] = [2,2] 和 items[3] = [5,3] 。
子序列的总利润为 3 + 2 + 5 = 10 ，子序列包含 3 种不同类别 [1, 2, 3] 。 
因此，优雅度为 10 + 3<sup>2</sup> = 19 ，可以证明 19 是可以获得的最大优雅度。</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>items = [[1,1],[2,1],[3,1]], k = 3
<strong>输出：</strong>7
<strong>解释：
</strong>在这个例子中，我们需要选出长度为 3 的子序列。
我们需要选中所有项目。
子序列的总利润为 1 + 2 + 3 = 6，子序列包含 1 种不同类别 [1] 。
因此，最大优雅度为 6 + 1<sup>2</sup> = 7 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= items.length == n <= 10<sup>5</sup></code></li>
	<li><code>items[i].length == 2</code></li>
	<li><code>items[i][0] == profit<sub>i</sub></code></li>
	<li><code>items[i][1] == category<sub>i</sub></code></li>
	<li><code>1 <= profit<sub>i</sub> <= 10<sup>9</sup></code></li>
	<li><code>1 <= category<sub>i</sub> <= n </code></li>
	<li><code>1 <= k <= n</code></li>
</ul>

#### 思路

按照利润从大到小排序。先把前 $k$ 个项目选上。
考虑选第 $k+1$ 个项目，为了选它，我们必须从前 $k$ 个项目中**移除**一个项目。由于已经按照利润从大到小排序，选这个项目不会让 $\textit{total\_profit}$ 变大，所以我们重点考虑能否让 $\textit{distinct\_categories}$ 变大。分类讨论：

- 如果第 $k+1$ 个项目和前面已选项目的类别重复了，那么选这个项目后，$\textit{distinct\_categories}$ 保持不变，所以不会让优雅度变大，此时无需选择这个项目。
- 如果第 $k+1$ 个项目和前面任何已选项目的类别都不重复，考虑移除前面已选项目中的哪一个：
  - 如果移除的项目的类别只出现一次，那么选第 $k+1$ 个项目后，$\textit{distinct\_categories}$ 一减一增，保持不变，所以不会让优雅度变大，此时无需选择这个项目。
  - 如果移除的项目的类别重复出现多次，$\textit{distinct\_categories}$ 会增加一，此时有可能会让优雅度变大，可以选择这个项目。

按照这个过程，继续考虑选择后面的项目。计算优雅度，取最大值，即为答案。
代码实现时，我们应当移除已选项目中类别和前面重复且利润最小的项目，这可以用一个**最小堆**来维护。
注：这个算法叫做**反悔贪心**，使用的堆叫做**反悔堆**。

```go  
func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	sort.Slice(items, func(i, j int) bool { return items[i][0] > items[j][0] })
	ans, totalProfit := 0, 0
	vis := map[int]bool{}
	duplicate := hp{} // 重复类别的利润
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit
			if !vis[category] {
				vis[category] = true
			} else { // 重复类别
				heap.Push(&duplicate, profit)
			}
		} else if !vis[category] && duplicate.Len() > 0 {
			vis[category] = true
			totalProfit += profit - heap.Pop(&duplicate).(int) // 选一个重复类别中的最小利润替换
		} // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }
func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{items}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
