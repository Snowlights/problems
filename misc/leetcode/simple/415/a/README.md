#### 题目

<p>数字小镇 Digitville 中，存在一个数字列表 <code>nums</code>，其中包含从 <code>0</code> 到 <code>n - 1</code> 的整数。每个数字本应 <strong>只出现一次</strong>，然而，有 <strong>两个 </strong>顽皮的数字额外多出现了一次，使得列表变得比正常情况下更长。</p>

<p>为了恢复 Digitville 的和平，作为小镇中的名侦探，请你找出这两个顽皮的数字。</p>

<p>返回一个长度为 2 的数组，包含这两个数字（顺序任意）。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [0,1,1,0]</span></p>

<p><strong>输出：</strong> <span class="example-io">[0,1]</span></p>

<p><strong>解释：</strong></p>

<p>数字 0 和 1 分别在数组中出现了两次。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [0,3,2,1,3,2]</span></p>

<p><strong>输出：</strong> <span class="example-io">[2,3]</span></p>

<p><strong>解释: </strong></p>

<p>数字 2 和 3 分别在数组中出现了两次。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [7,1,5,4,3,4,6,0,9,5,8,2]</span></p>

<p><strong>输出：</strong> <span class="example-io">[4,5]</span></p>

<p><strong>解释: </strong></p>

<p>数字 4 和 5 分别在数组中出现了两次。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 100</code></li>
	<li><code>nums.length == n + 2</code></li>
	<li><code>0 &lt;= nums[i] &lt; n</code></li>
	<li>输入保证 <code>nums</code> 中<strong> 恰好 </strong>包含两个重复的元素。</li>
</ul>

#### 思路

本题和 [2965. 找出缺失和重复的数字](https://leetcode.cn/problems/find-missing-and-repeated-values/) 本质是一样的，见 [我的题解](https://leetcode.cn/problems/find-missing-and-repeated-values/solutions/2569783/mo-ni-pythonjavacgo-by-endlesscheng-mexz/)，有位运算和数学两种做法。

## 位运算

需要两次遍历。一次遍历见下面的数学做法。

```
func getSneakyNumbers(nums []int) []int {
	n := len(nums) - 2
	a := -n * (n - 1) / 2
	b := -n * (n - 1) * (n*2 - 1) / 6
	for _, x := range nums {
		a += x
		b += x * x
	}
	x := int((float64(a) - math.Sqrt(float64(b*2-a*a))) / 2)
	return []int{x, a - x}
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)