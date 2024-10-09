### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>请你将 <code>nums</code>&nbsp;中每一个元素都替换为它的各个数位之 <strong>和</strong>&nbsp;。</p>

<p>请你返回替换所有元素以后 <code>nums</code>&nbsp;中的 <strong>最小</strong>&nbsp;元素。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [10,12,13,14]</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;替换后变为&nbsp;<code>[1, 3, 4, 5]</code>&nbsp;，最小元素为 1 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3,4]</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><b>解释：</b></p>

<p><code>nums</code>&nbsp;替换后变为&nbsp;<code>[1, 2, 3, 4]</code>&nbsp;，最小元素为 1 。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [999,19,199]</span></p>

<p><span class="example-io"><b>输出：</b>10</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;替换后变为&nbsp;<code>[27, 10, 19]</code>&nbsp;，最小元素为 10 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
</ul>


### 思路

设 $x=\textit{nums}[i]$，如何遍历 $x$ 的每个数位？

- 取 $x$ 的个位数，即 $x\bmod 10$，加到 $s$ 中。
- 把 $x$ 更新为 $\left\lfloor\dfrac{x}{10}\right\rfloor$，这样十位数变成个位数，百位数变成十位数，……
- 重复上述过程，直到 $x=0$。

用 $s$ 更新答案的最小值。

```
func minElement(nums []int) int {
	ans := math.MaxInt
	for _, x := range nums {
		s := 0
		for x > 0 {
			s += x % 10
			x /= 10
		}
		ans = min(ans, s)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
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