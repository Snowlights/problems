### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;，一个整数数组&nbsp;<code>queries</code>&nbsp;和一个整数&nbsp;<code>x</code>&nbsp;。</p>

<p>对于每个查询&nbsp;<code>queries[i]</code>&nbsp;，你需要找到&nbsp;<code>nums</code>&nbsp;中第&nbsp;<code>queries[i]</code>&nbsp;个&nbsp;<code>x</code>&nbsp;的位置，并返回它的下标。如果数组中&nbsp;<code>x</code>&nbsp;的出现次数少于&nbsp;<code>queries[i]</code>&nbsp;，该查询的答案为 -1 。</p>

<p>请你返回一个整数数组&nbsp;<code>answer</code>&nbsp;，包含所有查询的答案。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,3,1,7], queries = [1,3,2,4], x = 1</span></p>

<p><span class="example-io"><b>输出：</b>[0,-1,2,-1]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>第 1 个查询，第一个 1 出现在下标 0 处。</li>
	<li>第 2 个查询，<code>nums</code>&nbsp;中只有两个 1 ，所以答案为 -1 。</li>
	<li>第 3 个查询，第二个 1 出现在下标 2 处。</li>
	<li>第 4 个查询，<code>nums</code>&nbsp;中只有两个 1 ，所以答案为 -1 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3], queries = [10], x = 5</span></p>

<p><span class="example-io"><b>输出：</b>[-1]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>第 1 个查询，<code>nums</code>&nbsp;中没有 5 ，所以答案为 -1 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length, queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= queries[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i], x &lt;= 10<sup>4</sup></code></li>
</ul>

### 思路

用 $\textit{pos}$ 数组记录所有等于 $x$ 的 $\textit{nums}[i]$ 的下标 $i$。

对于每个询问 $q=\textit{queries}[i]$，如果 $q$ 大于 $\textit{pos}$ 的长度，则答案为 $-1$，否则答案为 $\textit{pos}[q-1]$。

``` 
func occurrencesOfElement(nums, queries []int, x int) []int {
	pos := []int{}
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}
	for i, q := range queries {
		if q > len(pos) {
			queries[i] = -1
		} else {
			queries[i] = pos[q-1]
		}
	}
	return queries
}

```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。
- 
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
