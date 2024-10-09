#### 题目

<p>给你一个长度为 <code>3</code> 的整数数组 <code>nums</code>。</p>

<p>现以某种顺序<strong> 连接 </strong>数组 <code>nums</code> 中所有元素的 <strong>二进制表示</strong> ，请你返回可以由这种方法形成的 <strong>最大 </strong>数值。</p>

<p><strong>注意</strong> 任何数字的二进制表示<em> </em><strong>不含</strong><em> </em>前导零。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1:</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">nums = [1,2,3]</span></p>

<p><strong>输出:</strong> 30</p>

<p><strong>解释:</strong></p>

<p>按照顺序 <code>[3, 1, 2]</code> 连接数字的二进制表示，得到结果 <code>"11110"</code>，这是 30 的二进制表示。</p>
</div>

<p><strong class="example">示例 2:</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">nums = [2,8,16]</span></p>

<p><strong>输出:</strong> 1296</p>

<p><strong>解释:</strong></p>

<p>按照顺序 <code>[2, 8, 16]</code> 连接数字的二进制表述，得到结果 <code>"10100010000"</code>，这是 1296 的二进制表示。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>nums.length == 3</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 127</code></li>
</ul>

#### 思路

本题是 [179. 最大数](https://leetcode.cn/problems/largest-number/) 的二进制版本，做法是一样的。

把 $\textit{nums}$ 排序，对于相比较的两个元素 $a$ 和 $b$，如果其二进制表示（字符串）满足 $a+b>b+a$，那么 $a$ 排在 $b$ 的左边，否则 $b$ 排在 $a$ 的左边。

```
func maxGoodNumber(nums []int) (ans int) {
	slices.SortFunc(nums, func(a, b int) int {
		lenA := bits.Len(uint(a))
		lenB := bits.Len(uint(b))
		return (b<<lenA | a) - (a<<lenB | b)
	})

	for _, x := range nums {
		ans = ans<<bits.Len(uint(x)) | x
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

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