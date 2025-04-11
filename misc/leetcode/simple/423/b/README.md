#### 题目

<p>给你一个由 <code>n</code> 个整数组成的数组 <code>nums</code> ，请你找出 <code>k</code> 的 <strong>最大值</strong>，使得存在 <strong>两个</strong> <strong>相邻</strong> 且长度为 <code>k</code> 的 <strong>严格递增</strong> 子数组。具体来说，需要检查是否存在从下标 <code>a</code> 和 <code>b</code> (<code>a &lt; b</code>) 开始的 <strong>两个</strong> 子数组，并满足下述全部条件：</p>

<ul>
	<li>这两个子数组 <code>nums[a..a + k - 1]</code> 和 <code>nums[b..b + k - 1]</code> 都是 <strong>严格递增</strong> 的。</li>
	<li>这两个子数组必须是 <strong>相邻的</strong>，即 <code>b = a + k</code>。</li>
</ul>

<p>返回 <code>k</code> 的 <strong>最大可能 </strong>值。</p>

<p><strong>子数组</strong> 是数组中的一个连续<b> 非空</b> 的元素序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [2,5,7,8,9,2,3,4,3,1]</span></p>

<p><strong>输出：</strong><span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>从下标 2 开始的子数组是 <code>[7, 8, 9]</code>，它是严格递增的。</li>
	<li>从下标 5 开始的子数组是 <code>[2, 3, 4]</code>，它也是严格递增的。</li>
	<li>这两个子数组是相邻的，因此 3 是满足题目条件的 <strong>最大</strong> <code>k</code> 值。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [1,2,3,4,4,4,4,5,6,7]</span></p>

<p><strong>输出：</strong><span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>从下标 0 开始的子数组是 <code>[1, 2]</code>，它是严格递增的。</li>
	<li>从下标 2 开始的子数组是 <code>[3, 4]</code>，它也是严格递增的。</li>
	<li>这两个子数组是相邻的，因此 2 是满足题目条件的 <strong>最大</strong> <code>k</code> 值。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

遍历 $\textit{nums}$，寻找严格递增段（子数组）。

设当前严格递增段的长度为 $\textit{cnt}$，上一个严格递增段的长度为 $\textit{preCnt}$。

答案有两种情况：

- 两个子数组属于同一个严格递增段，那么 $k$ 最大是 $\left\lfloor\dfrac{\textit{cnt}}{2}\right\rfloor$。
- 两个子数组分别属于一对相邻严格递增段，那么 $k$ 最大是 $\min(\textit{preCnt}, \textit{cnt})$。

```
func maxIncreasingSubarrays(nums []int) (ans int) {
	preCnt, cnt := 0, 0
	for i, x := range nums {
		cnt++
		if i == len(nums)-1 || x >= nums[i+1] { // i 是严格递增段的末尾
			ans = max(ans, cnt/2, min(preCnt, cnt))
			preCnt = cnt
			cnt = 0
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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