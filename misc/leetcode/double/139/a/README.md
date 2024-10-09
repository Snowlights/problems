### 题目

<p>有&nbsp;<code>n</code>&nbsp;座山排成一列，每座山都有一个高度。给你一个整数数组&nbsp;<code>height</code>&nbsp;，其中&nbsp;<code>height[i]</code>&nbsp;表示第 <code>i</code>&nbsp;座山的高度，再给你一个整数&nbsp;<code>threshold</code>&nbsp;。</p>

<p>对于下标不为 <code>0</code>&nbsp;的一座山，如果它左侧相邻的山的高度 <strong>严格</strong><strong>大于</strong>&nbsp;<code>threshold</code>&nbsp;，那么我们称它是 <strong>稳定</strong>&nbsp;的。我们定义下标为 <code>0</code>&nbsp;的山 <strong>不是</strong>&nbsp;稳定的。</p>

<p>请你返回一个数组，包含所有 <strong>稳定</strong>&nbsp;山的下标，你可以以 <strong>任意</strong>&nbsp;顺序返回下标数组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>height = [1,2,3,4,5], threshold = 2</span></p>

<p><span class="example-io"><b>输出：</b>[3,4]</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>下标为 3 的山是稳定的，因为&nbsp;<code>height[2] == 3</code>&nbsp;大于&nbsp;<code>threshold == 2</code>&nbsp;。</li>
	<li>下标为 4 的山是稳定的，因为&nbsp;<code>height[3] == 4</code> 大于 <code>threshold == 2</code>.</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>height = [10,1,10,1,10], threshold = 3</span></p>

<p><span class="example-io"><b>输出：</b>[1,3]</span></p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>height = [10,1,10,1,10], threshold = 10</span></p>

<p><span class="example-io"><b>输出：</b>[]</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n == height.length &lt;= 100</code></li>
	<li><code>1 &lt;= height[i] &lt;= 100</code></li>
	<li><code>1 &lt;= threshold &lt;= 100</code></li>
</ul>


### 思路

遍历下标 $0$ 到 $n-2$，如果发现 $\textit{height}[i]>\textit{threshold}$，就说明 $i+1$ 是稳定的，加入答案。

```
func stableMountains(height []int, threshold int) (ans []int) {
	for i, h := range height[:len(height)-1] {
		if h > threshold {
			ans = append(ans, i+1)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{height}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略切片空间。

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