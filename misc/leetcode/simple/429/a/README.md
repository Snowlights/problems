#### 题目

<p>给你一个整数数组 <code>nums</code>，你需要确保数组中的元素&nbsp;<strong>互不相同&nbsp;</strong>。为此，你可以执行以下操作任意次：</p>

<ul>
	<li>从数组的开头移除 3 个元素。如果数组中元素少于 3 个，则移除所有剩余元素。</li>
</ul>

<p><strong>注意：</strong>空数组也视作为数组元素互不相同。返回使数组元素互不相同所需的&nbsp;<strong>最少操作次数&nbsp;</strong>。<!-- notionvc: 210ee4f2-90af-4cdf-8dbc-96d1fa8f67c7 --></p>

<p>&nbsp;</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3,4,2,3,3,5,7]</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>第一次操作：移除前 3 个元素，数组变为 <code>[4, 2, 3, 3, 5, 7]</code>。</li>
	<li>第二次操作：再次移除前 3 个元素，数组变为 <code>[3, 5, 7]</code>，此时数组中的元素互不相同。</li>
</ul>

<p>因此，答案是 2。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [4,5,6,4,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>第一次操作：移除前 3 个元素，数组变为 <code>[4, 4]</code>。</li>
	<li>第二次操作：移除所有剩余元素，数组变为空。</li>
</ul>

<p>因此，答案是 2。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [6,7,8,9]</span></p>

<p><strong>输出：</strong> <span class="example-io">0</span></p>

<p><strong>解释：</strong></p>

<p>数组中的元素已经互不相同，因此不需要进行任何操作，答案是 0。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 100</code></li>
</ul>

#### 思路

倒着遍历 $\textit{nums}$。如果 $\textit{nums}[i]$ 在右边遇到过，那么说明下标在 $[0,i]$ 中的元素都要移除，这需要操作

$$
\left\lceil\dfrac{i+1}{3}\right\rceil = \left\lfloor\dfrac{i}{3}\right\rfloor + 1
$$

次。

如果 $\textit{nums}$ 本来就没有重复元素，返回 $0$。

```
func minimumOperations(nums []int) int {
	seen := map[int]bool{}
	for i, x := range slices.Backward(nums) {
		if seen[x] {
			return i/3 + 1
		}
		seen[x] = true
	}
	return 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。


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