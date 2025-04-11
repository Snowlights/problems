#### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p><strong>分区</strong>&nbsp;是指将数组按照下标&nbsp;<code>i</code>&nbsp;（<code>0 &lt;= i &lt; n - 1</code>）划分成两个 <strong>非空</strong> 子数组，其中：</p>

<ul>
	<li>左子数组包含区间&nbsp;<code>[0, i]</code>&nbsp;内的所有下标。</li>
	<li>右子数组包含区间&nbsp;<code>[i + 1, n - 1]</code>&nbsp;内的所有下标。</li>
</ul>

<p>对左子数组和右子数组先求元素 <strong>和</strong> 再做 <strong>差</strong> ，统计并返回差值为 <strong>偶数</strong> 的 <strong>分区</strong> 方案数。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><b>输入：</b><span class="example-io">nums = [10,10,3,7,6]</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><b>解释：</b></p>

<p>共有 4 个满足题意的分区方案：</p>

<ul>
	<li><code>[10]</code>、<code>[10, 3, 7, 6]</code>&nbsp;元素和的差值为&nbsp;<code>10 - 26 = -16</code>&nbsp;，是偶数。</li>
	<li><code>[10, 10]</code>、<code>[3, 7, 6]</code> 元素和的差值为&nbsp;<code>20 - 16 = 4</code>，是偶数。</li>
	<li><code>[10, 10, 3]</code>、<code>[7, 6]</code> 元素和的差值为&nbsp;<code>23 - 13 = 10</code>，是偶数。</li>
	<li><code>[10, 10, 3, 7]</code>、<code>[6]</code> 元素和的差值为&nbsp;<code>30 - 6 = 24</code>，是偶数。</li>
</ul>
</div>

<p><b>示例 2：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,2]</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><b>解释：</b></p>

<p>不存在元素和的差值为偶数的分区方案。</p>
</div>

<p><b>示例 3：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,4,6,8]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><b>解释：</b></p>

<p>所有分区方案都满足元素和的差值为偶数。</p>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>2 &lt;= n == nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 100</code></li>
</ul>

#### 思路

设 $\textit{nums}$ 的元素和为 $S$，左子数组元素和为 $L$，那么右子数组的元素和为 $S-L$。

题目要求 $L - (S-L) = 2L - S$ 是偶数。由于 $2L$ 一定是偶数，所以问题变成判断 $S$ 是否为偶数。

注意这和 $i$ 无关。换句话说，如果 $S$ 是奇数，那么答案是 $0$；如果 $S$ 是偶数，那么所有分区方案都符合要求，答案是 $n-1$。


```
func countPartitions(nums []int) int {
	s := 0
	for _, x := range nums {
		s += x
	}
	if s%2 == 0 {
		return len(nums) - 1
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)