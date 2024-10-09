#### 题目

<p>你有一个初始为空的浮点数数组 <code>averages</code>。另给你一个包含 <code>n</code> 个整数的数组 <code>nums</code>，其中 <code>n</code> 为偶数。</p>

<p>你需要重复以下步骤 <code>n / 2</code> 次：</p>

<ul>
	<li>从 <code>nums</code> 中移除<strong> 最小 </strong>的元素 <code>minElement</code> 和<strong> 最大 </strong>的元素 <code>maxElement</code>。</li>
	<li>将 <code>(minElement + maxElement) / 2</code> 加入到 <code>averages</code> 中。</li>
</ul>

<p>返回 <code>averages</code> 中的 <strong>最小 </strong>元素。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [7,8,3,4,15,13,4,1]</span></p>

<p><strong>输出：</strong> <span class="example-io">5.5</span></p>

<p><strong>解释：</strong></p>

<table>
	<tbody>
		<tr>
			<th>步骤</th>
			<th>nums</th>
			<th>averages</th>
		</tr>
		<tr>
			<td>0</td>
			<td>[7,8,3,4,15,13,4,1]</td>
			<td>[]</td>
		</tr>
		<tr>
			<td>1</td>
			<td>[7,8,3,4,13,4]</td>
			<td>[8]</td>
		</tr>
		<tr>
			<td>2</td>
			<td>[7,8,4,4]</td>
			<td>[8,8]</td>
		</tr>
		<tr>
			<td>3</td>
			<td>[7,4]</td>
			<td>[8,8,6]</td>
		</tr>
		<tr>
			<td>4</td>
			<td>[]</td>
			<td>[8,8,6,5.5]</td>
		</tr>
	</tbody>
</table>
返回 averages 中最小的元素，即 5.5。</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,9,8,3,10,5]</span></p>

<p><strong>输出：</strong> <span class="example-io">5.5</span></p>

<p><strong>解释：</strong></p>

<table>
	<tbody>
		<tr>
			<th>步骤</th>
			<th>nums</th>
			<th>averages</th>
		</tr>
		<tr>
			<td>0</td>
			<td>[1,9,8,3,10,5]</td>
			<td>[]</td>
		</tr>
		<tr>
			<td>1</td>
			<td>[9,8,3,5]</td>
			<td>[5.5]</td>
		</tr>
		<tr>
			<td>2</td>
			<td>[8,5]</td>
			<td>[5.5,6]</td>
		</tr>
		<tr>
			<td>3</td>
			<td>[]</td>
			<td>[5.5,6,6.5]</td>
		</tr>
	</tbody>
</table>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3,7,8,9]</span></p>

<p><strong>输出：</strong> <span class="example-io">5.0</span></p>

<p><strong>解释：</strong></p>

<table>
	<tbody>
		<tr>
			<th>步骤</th>
			<th>nums</th>
			<th>averages</th>
		</tr>
		<tr>
			<td>0</td>
			<td>[1,2,3,7,8,9]</td>
			<td>[]</td>
		</tr>
		<tr>
			<td>1</td>
			<td>[2,3,7,8]</td>
			<td>[5]</td>
		</tr>
		<tr>
			<td>2</td>
			<td>[3,7]</td>
			<td>[5,5]</td>
		</tr>
		<tr>
			<td>3</td>
			<td>[]</td>
			<td>[5,5,5]</td>
		</tr>
	</tbody>
</table>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n == nums.length &lt;= 50</code></li>
	<li><code>n</code> 为偶数。</li>
	<li><code>1 &lt;= nums[i] &lt;= 50</code></li>
</ul>

#### 思路

把 $\textit{nums}$ 从小到大排序。

排序后，计算所有 $\textit{nums}[i]+\textit{nums}[n-1-i]$ 的最小值，返回其除以 $2$ 的结果。

```
func minimumAverage(nums []int) float64 {
	slices.Sort(nums)
	ans := math.MaxInt
	for i, n := 0, len(nums); i < n/2; i++ {
		ans = min(ans, nums[i]+nums[n-1-i])
	}
	return float64(ans) / 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)