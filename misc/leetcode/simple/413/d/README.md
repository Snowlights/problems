#### 题目

<p>给你一个由 <code>n</code> 个整数组成的数组 <code>nums</code>，以及一个大小为 <code>q</code> 的二维整数数组 <code>queries</code>，其中 <code>queries[i] = [l<sub>i</sub>, r<sub>i</sub>]</code>。</p>

<p>对于每一个查询，你需要找出 <code>nums[l<sub>i</sub>..r<sub>i</sub>]</code> 中任意 <span data-keyword="subarray">子数组</span> 的 <strong>最大异或值</strong>。</p>

<p><strong>数组的异或值 </strong>需要对数组 <code>a</code> 反复执行以下操作，直到只剩一个元素，剩下的那个元素就是 <strong>异或值</strong>：</p>

<ul>
	<li><span class="text-only" data-eleid="9" style="white-space: pre;">对于除最后一个下标以外的所有下标</span> <code>i</code>，同时将 <code>a[i]</code> 替换为 <code>a[i] XOR a[i + 1]</code> 。</li>
	<li>移除数组的最后一个元素。</li>
</ul>

<p>返回一个大小为 <code>q</code> 的数组 <code>answer</code>，其中 <code>answer[i]</code> 表示查询 <code>i</code> 的答案。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [2,8,4,32,16,1], queries = [[0,2],[1,4],[0,5]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[12,60,60]</span></p>

<p><strong>解释：</strong></p>

<p>在第一个查询中，<code>nums[0..2]</code> 的子数组分别是 <code>[2]</code>, <code>[8]</code>, <code>[4]</code>, <code>[2, 8]</code>, <code>[8, 4]</code>, 和 <code>[2, 8, 4]</code>，它们的异或值分别为 2, 8, 4, 10, 12, 和 6。查询的答案是 12，所有异或值中的最大值。</p>

<p>在第二个查询中，<code>nums[1..4]</code> 的子数组中最大的异或值是子数组 <code>nums[1..4]</code> 的异或值，为 60。</p>

<p>在第三个查询中，<code>nums[0..5]</code> 的子数组中最大的异或值是子数组 <code>nums[1..4]</code> 的异或值，为 60。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [0,7,3,2,8,5,1], queries = [[0,3],[1,5],[2,4],[2,6],[5,6]]</span></p>

<p><strong>输出：</strong> <span class="example-io">[7,14,11,14,5]</span></p>

<p><strong>解释：</strong></p>

<table height="70" width="472">
	<thead>
		<tr>
			<th>下标</th>
			<th>nums[l<sub>i</sub>..r<sub>i</sub>]</th>
			<th>最大异或值子数组</th>
			<th>子数组最大异或值</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>0</td>
			<td>[0, 7, 3, 2]</td>
			<td>[7]</td>
			<td>7</td>
		</tr>
		<tr>
			<td>1</td>
			<td>[7, 3, 2, 8, 5]</td>
			<td>[7, 3, 2, 8]</td>
			<td>14</td>
		</tr>
		<tr>
			<td>2</td>
			<td>[3, 2, 8]</td>
			<td>[3, 2, 8]</td>
			<td>11</td>
		</tr>
		<tr>
			<td>3</td>
			<td>[3, 2, 8, 5, 1]</td>
			<td>[2, 8, 5, 1]</td>
			<td>14</td>
		</tr>
		<tr>
			<td>4</td>
			<td>[5, 1]</td>
			<td>[5]</td>
			<td>5</td>
		</tr>
	</tbody>
</table>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == nums.length &lt;= 2000</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 2<sup>31</sup> - 1</code></li>
	<li><code>1 &lt;= q == queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>queries[i] = [l<sub>i</sub>, r<sub>i</sub>]</code></li>
	<li><code>0 &lt;= l<sub>i</sub> &lt;= r<sub>i</sub> &lt;= n - 1</code></li>
</ul>

#### 思路

## 数组的异或值

下文用 $\oplus$ 表示异或。

考虑数组的异或值（最后剩下的元素）是由哪些元素异或得到的。

例如数组为 $[a,b,c]$，那么操作一次后变成 $[a\oplus b,\ b\oplus c]$，再操作一次，得到 $a\oplus b\oplus b\oplus c$。其中 $b$ 异或了 $2$ 次。

为方便描述，下面把 $a\oplus b$ 简记为 $ab$，把 $a\oplus b\oplus b\oplus c$ 简记为 $ab^2c$。

又例如数组为 $[a,b,c,d]$，那么操作一次后变成 $[ab,bc,cd]$，再操作一次，变成 $[ab^2c,bc^2d]$，再操作一次，得到 $ab^3c^3d$。

可以发现，$ab^3c^3d$ 相当于数组 $[a,b,c]$ 的异或值，再异或 $[b,c,d]$ 的异或值。

> 当然，你可以用组合数计算出幂次，那就是另一个思路了，具体见 [本题视频讲解](https://www.bilibili.com/video/BV142Hae7E5y/)。

## 第一个区间 DP

定义 $f[i][j]$ 表示下标从 $i$ 到 $j$ 的子数组的「数组的异或值」，根据上面的讨论，有

$$
f[i][j] = f[i][j-1]\oplus f[i+1][j]
$$

初始值：$f[i][i]=\textit{nums}[i]$。

## 第二个区间 DP

为了回答询问，我们需要计算下标从 $i$ 到 $j$ 的子数组中的所有子数组的 $f$ 值的最大值，将其记作 $\textit{mx}[i][j]$。

分类讨论：

- 取 $f[i][j]$ 作为最大值。
- 最大值对应的子数组，右端点不是 $j$，那么问题变成下标从 $i$ 到 $j-1$ 的子数组中的所有子数组的 $f$ 值的最大值，即 $\textit{mx}[i][j-1]$。
- 最大值对应的子数组，左端点不是 $i$，那么问题变成下标从 $i+1$ 到 $j$ 的子数组中的所有子数组的 $f$ 值的最大值，即 $\textit{mx}[i+1][j]$。

三者取最大值，得

$$
\textit{mx}[i][j] = \max(f[i][j], \textit{mx}[i][j-1], \textit{mx}[i+1][j])
$$

初始值：$\textit{mx}[i][i]=\textit{nums}[i]$。

回答询问时直接查询 $\textit{mx}$ 数组。

```
func maximumSubarrayXor(nums []int, queries [][]int) []int {
	n := len(nums)
	f := make([][]int, n)
	mx := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		mx[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		f[i][i] = nums[i]
		mx[i][i] = nums[i]
		for j := i + 1; j < n; j++ {
			f[i][j] = f[i][j-1] ^ f[i+1][j]
			mx[i][j] = max(f[i][j], mx[i+1][j], mx[i][j-1])
		}
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = mx[q[0]][q[1]]
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度
- 空间复杂度：$\mathcal{O}(n^2)$。返回值不计入。


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
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)