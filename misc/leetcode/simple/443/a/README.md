#### 题目

<p data-end="438" data-start="104">给你一个长度为 <code>n</code> 的整数数组 <code data-end="119" data-start="113">cost</code> 。当前你位于位置 <code data-end="166" data-start="163">n</code>（队伍的末尾），队伍中共有 <code data-end="187" data-start="180">n + 1</code> 人，编号从 0 到 <code>n</code> 。</p>

<p data-end="438" data-start="104">你希望在队伍中向前移动，但队伍中每个人都会收取一定的费用才能与你 <strong>交换</strong>位置。与编号 <code data-end="375" data-start="372">i</code> 的人交换位置的费用为 <code data-end="397" data-start="388">cost[i]</code> 。</p>

<p data-end="487" data-start="440">你可以按照以下规则与他人交换位置：</p>

<ul data-end="632" data-start="488">
	<li data-end="572" data-start="488">如果对方在你前面，你 <strong>必须</strong> 支付 <code data-end="546" data-start="537">cost[i]</code> 费用与他们交换位置。</li>
	<li data-end="632" data-start="573">如果对方在你后面，他们可以免费与你交换位置。</li>
</ul>

<p data-end="755" data-start="634">返回一个大小为 <code>n</code> 的数组 <code>answer</code>，其中 <code>answer[i]</code> 表示到达队伍中每个位置 <code>i</code> 所需的 <strong data-end="680" data-start="664">最小</strong> 总费用。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">cost = [5,3,4,1,3,2]</span></p>

<p><strong>输出:</strong> <span class="example-io">[5,3,3,1,1,1]</span></p>

<p><strong>解释:</strong></p>

<p>我们可以通过以下方式到达每个位置：</p>

<ul>
	<li><code>i = 0</code>。可以花费 5 费用与编号 0 的人交换位置。</li>
	<li><span class="example-io"><code>i = 1</code>。可以花费 3 费用与编号 1 的人交换位置。</span></li>
	<li><span class="example-io"><code>i = 2</code>。可以花费 3 费用与编号 1 的人交换位置，然后免费与编号 2 的人交换位置。</span></li>
	<li><span class="example-io"><code>i = 3</code>。可以花费 1 费用与编号 3 的人交换位置。</span></li>
	<li><span class="example-io"><code>i = 4</code>。可以花费 1 费用与编号 3 的人交换位置，然后免费与编号 4 的人交换位置。</span></li>
	<li><span class="example-io"><code>i = 5</code>。可以花费 1 费用与编号 3 的人交换位置，然后免费与编号 5 的人交换位置。</span></li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">cost = [1,2,4,6,7]</span></p>

<p><strong>输出:</strong> <span class="example-io">[1,1,1,1,1]</span></p>

<p><strong>解释:</strong></p>

<p>可以花费 1 费用与编号 0 的人交换位置，然后可以免费到达队伍中的任何位置 <code>i</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示</strong></p>

<ul>
	<li><code>1 &lt;= n == cost.length &lt;= 100</code></li>
	<li><code>1 &lt;= cost[i] &lt;= 100</code></li>
</ul>

#### 思路

要想到达 $i$，我们可以先找一个 $[0,i]$ 中 $\textit{cost}$ 最小的人，和他交换，然后再和右边的人免费交换，直到到达 $i$。

所以 $\textit{answer}[i]$ 就是 $[0,i]$ 中 $\textit{cost}$ 的最小值（前缀最小值）。

这可以从左到右遍历 $\textit{cost}$ 算出来。

代码实现时，直接原地修改 $\textit{cost}$，这样可以做到 $\mathcal{O}(1)$ 空间。

```
func minCosts(cost []int) []int {
	for i := 1; i < len(cost); i++ {
		cost[i] = min(cost[i], cost[i-1])
	}
	return cost
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{cost}$ 的长度。
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