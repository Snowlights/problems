#### 题目

<p>给定一个长度为 <code>n</code> 的整数数组 <code>nums</code> 和一个二维数组 <code>queries</code>，其中 <code>queries[i] = [l<sub>i</sub>, r<sub>i</sub>]</code>。</p>

<p>对于每个查询&nbsp;<code>queries[i]</code>：</p>

<ul>
	<li>在&nbsp;<code>nums</code>&nbsp;的下标范围&nbsp;<code>[l<sub>i</sub>, r<sub>i</sub>]</code>&nbsp;内选择一个下标子集。</li>
	<li>将选中的每个下标对应的元素值减 1。</li>
</ul>

<p><strong>零数组&nbsp;</strong>是指所有元素都等于 0 的数组。</p>

<p>如果在按顺序处理所有查询后，可以将 <code>nums</code> 转换为&nbsp;<strong>零数组&nbsp;</strong>，则返回 <code>true</code>，否则返回 <code>false</code>。</p>

<p>数组的&nbsp;<strong>子集&nbsp;</strong>是对数组元素的选择（可能为空）。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,0,1], queries = [[0,2]]</span></p>

<p><strong>输出：</strong> <span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><strong>对于 i = 0：</strong>
	<ul>
		<li>选择下标子集 <code>[0, 2]</code> 并将这些下标处的值减 1。</li>
		<li>数组将变为 <code>[0, 0, 0]</code>，这是一个零数组。</li>
	</ul>
	</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [4,3,2,1], queries = [[1,3],[0,2]]</span></p>

<p><strong>输出：</strong> <span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li><strong>对于 i = 0：</strong>&nbsp;
	<ul>
		<li>选择下标子集 <code>[1, 2, 3]</code> 并将这些下标处的值减 1。</li>
		<li>数组将变为 <code>[4, 2, 1, 0]</code>。</li>
	</ul>
	</li>
	<li><strong>对于 i = 1：</strong>
	<ul>
		<li>选择下标子集 <code>[0, 1, 2]</code> 并将这些下标处的值减 1。</li>
		<li>数组将变为 <code>[3, 1, 0, 0]</code>，这不是一个零数组。</li>
	</ul>
	</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>0 &lt;= l<sub>i</sub> &lt;= r<sub>i</sub> &lt; nums.length</code></li>
</ul>

#### 思路

题意可以转换成：

- 把 $[l_i,r_i]$ 中的元素都减一，最终数组中的所有元素是否都 $\le 0$？

如果所有元素都 $\le 0$，那么我们可以撤销一部分元素的减一，使其调整为 $0$，从而满足原始题意的要求。
这可以用**差分数组**计算，

```
func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)
	for _, q := range queries {
		// 区间 [l,r] 中的数都加一
		diff[q[0]]++
		diff[q[1]+1]--
	}

	sumD := 0
	for i, x := range nums {
		sumD += diff[i]
		// 此时 sumD 表示 x=nums[i] 要减掉多少
		if x > sumD { // x 无法变成 0
			return false
		}
	}
	return true
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度
- 空间复杂度：$\mathcal{O}(n)$。




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
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)