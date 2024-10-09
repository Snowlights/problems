#### 题目

<p>给你一个数组&nbsp;<code>maximumHeight</code>&nbsp;，其中&nbsp;<code>maximumHeight[i]</code>&nbsp;表示第 <code>i</code>&nbsp;座塔可以达到的 <strong>最大</strong>&nbsp;高度。</p>

<p>你的任务是给每一座塔分别设置一个高度，使得：</p>

<ol>
	<li>第 <code>i</code>&nbsp;座塔的高度是一个正整数，且不超过&nbsp;<code>maximumHeight[i]</code>&nbsp;。</li>
	<li>所有塔的高度互不相同。</li>
</ol>

<p>请你返回设置完所有塔的高度后，可以达到的 <strong>最大</strong>&nbsp;总高度。如果没有合法的设置，返回 <code>-1</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><b>输入：</b>maximumHeight<span class="example-io"> = [2,3,4,3]</span></p>

<p><span class="example-io"><b>输出：</b>10</span></p>

<p><strong>解释：</strong></p>

<p>我们可以将塔的高度设置为：<code>[1, 2, 4, 3]</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><b>输入：</b>maximumHeight<span class="example-io"> = [15,10]</span></p>

<p><span class="example-io"><b>输出：</b>25</span></p>

<p><strong>解释：</strong></p>

<p>我们可以将塔的高度设置为：<code>[15, 10]</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><b>输入：</b>maximumHeight<span class="example-io"> = [2,2,1]</span></p>

<p><span class="example-io"><b>输出：</b>-1</span></p>

<p><b>解释：</b></p>

<p>无法设置塔的高度为正整数且高度互不相同。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= maximumHeight.length&nbsp;&lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= maximumHeight[i] &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

从最大的元素开始思考：

- 数组的最大值 $m$ 不变是最好的。反证：如果把 $m$ 变小是最优的，那么把 $m$ 恢复成其原来的值，仍然满足题目要求，且我们得到了更优的答案，矛盾。
- 数组的次大值呢？如果它等于 $m$，那么它必须变成 $m-1$，否则不变。
- 依此类推。

为了方便计算，先把数组从大到小排序，那么 $\textit{maximumHeight}[i]$ 的实际值为

$$
\min(\textit{maximumHeight}[i], \textit{maximumHeight}[i-1] - 1)
$$

如果元素值 $\le 0$，不符合题目要求，返回 $-1$。

最终答案为 $\textit{maximumHeight}$ 的元素之和。

```
func maximumTotalSum(maximumHeight []int) int64 {
	slices.SortFunc(maximumHeight, func(a, b int) int { return b - a })
	ans := maximumHeight[0]
	for i := 1; i < len(maximumHeight); i++ {
		maximumHeight[i] = min(maximumHeight[i], maximumHeight[i-1]-1)
		if maximumHeight[i] <= 0 {
			return -1
		}
		ans += maximumHeight[i]
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{maximumHeight}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)