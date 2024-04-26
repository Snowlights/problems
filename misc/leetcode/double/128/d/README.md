### 题目

<p>给你一个 <strong>正</strong>&nbsp;整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>请你求出&nbsp;<code>nums</code>&nbsp;中有多少个子数组，满足子数组中&nbsp;<strong>第一个</strong>&nbsp;和 <strong>最后一个</strong>&nbsp;元素都是这个子数组中的 <strong>最大</strong>&nbsp;值。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,4,3,3,2]</span></p>

<p><span class="example-io"><b>输出：</b>6</span></p>

<p><strong>解释：</strong></p>

<p>总共有 6 个子数组满足第一个元素和最后一个元素都是子数组中的最大值：</p>

<ul>
	<li>子数组&nbsp;<code>[<u><em><strong>1</strong></em></u>,4,3,3,2]</code>&nbsp;，最大元素为 1 ，第一个和最后一个元素都是 1 。</li>
	<li>子数组&nbsp;<code>[1,<u><em><strong>4</strong></em></u>,3,3,2]</code>&nbsp;，最大元素为 4 ，第一个和最后一个元素都是 4 。</li>
	<li>子数组&nbsp;<code>[1,4,<u><em><strong>3</strong></em></u>,3,2]</code>&nbsp;，最大元素为 3 ，第一个和最后一个元素都是 3 。</li>
	<li>子数组&nbsp;<code>[1,4,3,<u><em><strong>3</strong></em></u>,2]</code>&nbsp;，最大元素为 3 ，第一个和最后一个元素都是 3 。</li>
	<li>子数组&nbsp;<code>[1,4,3,3,<u><em><strong>2</strong></em></u>]</code>&nbsp;，最大元素为 2 ，第一个和最后一个元素都是 2 。</li>
	<li>子数组&nbsp;<code>[1,4,<u><em><strong>3,3</strong></em></u>,2]</code>&nbsp;，最大元素为 3 ，第一个和最后一个元素都是 3 。</li>
</ul>

<p>所以我们返回 6 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [3,3,3]</span></p>

<p><span class="example-io"><b>输出：</b>6</span></p>

<p><strong>解释：</strong></p>

<p>总共有 6 个子数组满足第一个元素和最后一个元素都是子数组中的最大值：</p>

<ul>
	<li>子数组 <code>[<u><em><strong>3</strong></em></u>,3,3]</code>&nbsp;，最大元素为 3&nbsp;，第一个和最后一个元素都是 3&nbsp;。</li>
	<li>子数组 <code>[3,<u><em><strong>3</strong></em></u>,3]</code>&nbsp;，最大元素为 3&nbsp;，第一个和最后一个元素都是 3&nbsp;。</li>
	<li>子数组 <code>[3,3,<u><em><strong>3</strong></em></u>]</code>&nbsp;，最大元素为 3&nbsp;，第一个和最后一个元素都是 3&nbsp;。</li>
	<li>子数组 <code>[<u><em><strong>3,3</strong></em></u>,3]</code>&nbsp;，最大元素为 3&nbsp;，第一个和最后一个元素都是 3&nbsp;。</li>
	<li>子数组 <code>[3,<u><em><strong>3,3</strong></em></u>]</code>&nbsp;，最大元素为 3&nbsp;，第一个和最后一个元素都是 3&nbsp;。</li>
	<li>子数组 <code>[<u><em><strong>3,3,3</strong></em></u>]</code>&nbsp;，最大元素为 3&nbsp;，第一个和最后一个元素都是 3&nbsp;。</li>
</ul>

<p>所以我们返回 6 。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1]</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;中只有一个子数组&nbsp;<code>[<em><strong>1</strong></em>]</code>&nbsp;，最大元素为 1 ，第一个和最后一个元素都是 1 。</p>

<p>所以我们返回 1 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

### 思路

例如 $\textit{nums}=[4,3,1,2,1]$，在从左到右遍历的过程中，由于 $2$ 的出现，左边的 $1$ 永远不可能与右边的 $1$ 组成一个题目要求的子数组。
所以当遍历到 $2$ 时，左边的 $1$ 就是**无用数据**了，可以清除。
清除后我们会得到一个从左到右递减的数据结构。

这个性质和单调栈很像，启发我们用单调栈思考。具体来说：
1. 初始化答案等于 $n$，因为每个元素可以单独组成一个长为 $1$ 的子数组，满足题目要求。
2. 维护一个底大顶小的单调栈，记录元素及其出现次数。
3. 从左到右遍历 $\textit{nums}$。
4. 只要 $x=\textit{nums}[i]$ 大于栈顶，就把栈顶出栈。
5. 如果 $x$ 小于栈顶，把 $x$ 及其出现次数 $1$ 入栈。
6. 如果 $x$ 等于栈顶，设栈顶记录的出现次数为 $\textit{cnt}$，那么 $x$ 可以和左边 $\textit{cnt}$ 个 $x$ 组成 $\textit{cnt}$ 个满足要求的子数组，把答案增加 $\textit{cnt}$，然后把 $\textit{cnt}$ 加一。

注意可能出现某个元素 $v$ 出栈后，又重新入栈的情况，此时 $v$ 的出现次数会重置成 $1$。代码实现时，可以往栈底加入一个 $\infty$ 哨兵，从而简化判断逻辑。

```
func numberOfSubarrays(nums []int) int64 {
	ans := len(nums)
	type pair struct{ x, cnt int }
	st := []pair{{math.MaxInt, 0}}
	for _, x := range nums {
		for x > st[len(st)-1].x {
			st = st[:len(st)-1]
		}
		if x == st[len(st)-1].x {
			ans += st[len(st)-1].cnt
			st[len(st)-1].cnt++
		} else {
			st = append(st, pair{x, 1})
		}
	}
	return int64(ans)
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
