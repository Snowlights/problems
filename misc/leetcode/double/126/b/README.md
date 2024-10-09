### 题目

<p>给你一个长度为 <code>n</code>&nbsp;下标从 <strong>0</strong>&nbsp;开始的正整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>同时给你一个长度为 <code>m</code>&nbsp;的二维操作数组&nbsp;<code>queries</code>&nbsp;，其中&nbsp;<code>queries[i] = [index<sub>i</sub>, k<sub>i</sub>]</code>&nbsp;。</p>

<p>一开始，数组中的所有元素都 <strong>未标记</strong>&nbsp;。</p>

<p>你需要依次对数组执行 <code>m</code>&nbsp;次操作，第 <code>i</code>&nbsp;次操作中，你需要执行：</p>

<ul>
	<li>如果下标&nbsp;<code>index<sub>i</sub></code>&nbsp;对应的元素还没标记，那么标记这个元素。</li>
	<li>然后标记&nbsp;<code>k<sub>i</sub></code>&nbsp;个数组中还没有标记的&nbsp;<strong>最小</strong>&nbsp;元素。如果有元素的值相等，那么优先标记它们中下标较小的。如果少于&nbsp;<code>k<sub>i</sub></code>&nbsp;个未标记元素存在，那么将它们全部标记。</li>
</ul>

<p>请你返回一个长度为 <code>m</code>&nbsp;的数组 <code>answer</code>&nbsp;，其中<em>&nbsp;</em><code>answer[i]</code>是第&nbsp;<code>i</code>&nbsp;次操作后数组中还没标记元素的&nbsp;<strong>和</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [1,2,2,1,2,3,1], queries = [[1,2],[3,3],[4,2]]</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">[8,3,0]</span></p>

<p><strong>解释：</strong></p>

<p>我们依次对数组做以下操作：</p>

<ul>
	<li>标记下标为&nbsp;<code>1</code>&nbsp;的元素，同时标记&nbsp;<code>2</code>&nbsp;个未标记的最小元素。标记完后数组为&nbsp;<code>nums = [<em><strong>1</strong></em>,<em><strong>2</strong></em>,2,<em><strong>1</strong></em>,2,3,1]</code>&nbsp;。未标记元素的和为&nbsp;<code>2 + 2 + 3 + 1 = 8</code>&nbsp;。</li>
	<li>标记下标为&nbsp;<code>3</code>&nbsp;的元素，由于它已经被标记过了，所以我们忽略这次标记，同时标记最靠前的&nbsp;<code>3</code>&nbsp;个未标记的最小元素。标记完后数组为&nbsp;<code>nums = [<em><strong>1</strong></em>,<em><strong>2</strong></em>,<em><strong>2</strong></em>,<em><strong>1</strong></em>,<em><strong>2</strong></em>,3,<em><strong>1</strong></em>]</code>&nbsp;。未标记元素的和为&nbsp;<code>3</code>&nbsp;。</li>
	<li>标记下标为 <code>4</code>&nbsp;的元素，由于它已经被标记过了，所以我们忽略这次标记，同时标记最靠前的 <code>2</code>&nbsp;个未标记的最小元素。标记完后数组为&nbsp;<code>nums = [<em><strong>1</strong></em>,<em><strong>2</strong></em>,<em><strong>2</strong></em>,<em><strong>1</strong></em>,<em><strong>2</strong></em>,<em><strong>3</strong></em>,<em><strong>1</strong></em>]</code>&nbsp;。未标记元素的和为&nbsp;<code>0</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">nums = [1,4,2,3], queries = [[0,1]]</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">[7]</span></p>

<p><strong>解释：</strong>我们执行一次操作，将下标为&nbsp;<code>0</code>&nbsp;处的元素标记，并且标记最靠前的&nbsp;<code>1</code>&nbsp;个未标记的最小元素。标记完后数组为&nbsp;<code>nums = [<em><strong>1</strong></em>,4,<em><strong>2</strong></em>,3]</code>&nbsp;。未标记元素的和为&nbsp;<code>4 + 3 = 7</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums.length</code></li>
	<li><code>m == queries.length</code></li>
	<li><code>1 &lt;= m &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>0 &lt;= index<sub>i</sub>, k<sub>i</sub> &lt;= n - 1</code></li>
</ul>

### 思路

由于要按照元素 从小到大标记，但又不能直接对数组排序（因为有对特定 $\textit{index}$ 的标记），我们可以创建一个 $\textit{ids}$ 数组，其中 $\textit{ids}[i]=i$，然后对该数组按照 $\textit{nums}[\textit{ids}[i]]$ 从小到大排序。注意要使用**稳定排序**，因为相同元素值的下标需要按照下标从小到大排。也可以对于相同元素值按照下标从小到大排序。

设 $\textit{nums}$ 的元素和为 $s$。对于每个询问，我们先将 $s$ 减少 $\textit{nums}[\textit{index}]$，然后将 $\textit{nums}[\textit{index}]$ 置为 $0$，就相当于标记了这个数（因为题目保证数组元素都是正数）。然后依照 $\textit{ids}$ 找 $k$ 个最小的没有被标记的数，将其标记，标记的同时维护 $s$。

```go [sol-Go]
func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	s, n := 0, len(nums)
	ids := make([]int, n)
	for i, x := range nums {
		s += x
		ids[i] = i
	}
	slices.SortStableFunc(ids, func(i, j int) int { return nums[i] - nums[j] })

	ans := make([]int64, len(queries))
	j := 0
	for qi, p := range queries {
		i, k := p[0], p[1]
		s -= nums[i]
		nums[i] = 0 // 标记
		for ; j < n && k > 0; j++ {
			i := ids[j]
			if nums[i] > 0 { // 没有被标记
				s -= nums[i]
				nums[i] = 0
				k--
			}
		}
		ans[qi] = int64(s)
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。忽略返回值的空间。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

