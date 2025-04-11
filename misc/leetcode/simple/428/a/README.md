#### 题目

<p>给你一个二维数组 <code>events</code>，表示孩子在键盘上按下一系列按钮触发的按钮事件。</p>

<p>每个 <code>events[i] = [index<sub>i</sub>, time<sub>i</sub>]</code> 表示在时间 <code>time<sub>i</sub></code> 时，按下了下标为 <code>index<sub>i</sub></code> 的按钮。</p>

<ul>
	<li>数组按照 <code>time</code> 的递增顺序<strong>排序</strong>。</li>
	<li>按下一个按钮所需的时间是连续两次按钮按下的时间差。按下第一个按钮所需的时间就是其时间戳。</li>
</ul>

<p>返回按下时间&nbsp;<strong>最长&nbsp;</strong>的按钮的 <code>index</code>。如果有多个按钮的按下时间相同，则返回 <code>index</code> 最小的按钮。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">events = [[1,2],[2,5],[3,9],[1,15]]</span></p>

<p><strong>输出：</strong> <span class="example-io">1</span></p>

<p><strong>解释：</strong></p>
<ul>
	<li>下标为 1 的按钮在时间 2 被按下。</li>
	<li>下标为 2 的按钮在时间 5 被按下，因此按下时间为 <code>5 - 2 = 3</code>。</li>
	<li>下标为 3 的按钮在时间 9 被按下，因此按下时间为 <code>9 - 5 = 4</code>。</li>
	<li>下标为 1 的按钮再次在时间 15 被按下，因此按下时间为 <code>15 - 9 = 6</code>。</li>
</ul>

<p>最终，下标为 1 的按钮按下时间最长，为 6。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">events = [[10,5],[1,7]]</span></p>

<p><strong>输出：</strong> <span class="example-io">10</span></p>

<p><strong>解释：</strong></p>
<ul>
	<li>下标为 10 的按钮在时间 5 被按下。</li>
	<li>下标为 1 的按钮在时间 7 被按下，因此按下时间为 <code>7 - 5 = 2</code>。</li>
</ul>

<p>最终，下标为 10 的按钮按下时间最长，为 5。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= events.length &lt;= 1000</code></li>
	<li><code>events[i] == [index<sub>i</sub>, time<sub>i</sub>]</code></li>
	<li><code>1 &lt;= index<sub>i</sub>, time<sub>i</sub> &lt;= 10<sup>5</sup></code></li>
	<li>输入保证数组 <code>events</code> 按照 <code>time<sub>i</sub></code> 的递增顺序排序。</li>
</ul>

#### 思路

遍历的过程中，维护答案 $\textit{idx}$，以及最大时间差 $\textit{maxDiff}$。

```
func buttonWithLongestTime(events [][]int) int {
	idx, maxDiff := events[0][0], events[0][1]
	for i := 1; i < len(events); i++ {
		p, q := events[i-1], events[i]
		d := q[1] - p[1]
		if d > maxDiff || d == maxDiff && q[0] < idx {
			idx, maxDiff = q[0], d
		}
	}
	return idx
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{events}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。


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