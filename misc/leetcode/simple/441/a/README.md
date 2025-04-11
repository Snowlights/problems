#### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>你可以从数组 <code>nums</code> 中删除任意数量的元素，但不能将其变为 <strong>空</strong> 数组。执行删除操作后，选出&nbsp;<code>nums</code>&nbsp;中满足下述条件的一个子数组：</p>

<ol>
	<li>子数组中的所有元素 <strong>互不相同</strong> 。</li>
	<li><strong>最大化</strong> 子数组的元素和。</li>
</ol>

<p>返回子数组的 <strong>最大元素和</strong> 。</p>
<strong>子数组</strong> 是数组的一个连续、<strong>非空</strong> 的元素序列。

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3,4,5]</span></p>

<p><span class="example-io"><b>输出：</b>15</span></p>

<p><b>解释：</b></p>

<p>不删除任何元素，选中整个数组得到最大元素和。</p>
</div>

<p><b>示例 2：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b></span><span class="example-io">nums = [1,1,0,1,1]</span></p>

<p><span class="example-io"><b>输出：</b></span>1</p>

<p><b>解释：</b></p>

<p>删除元素&nbsp;<code>nums[0] == 1</code>、<code>nums[1] == 1</code>、<code>nums[2] == 0</code>&nbsp;和&nbsp;<code>nums[3] == 1</code>&nbsp;。选中整个数组&nbsp;<code>[1]</code>&nbsp;得到最大元素和。</p>
</div>

<p><b>示例 3：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b></span><span class="example-io">nums = [1,2,-1,-2,1,0,-1]</span></p>

<p><span class="example-io"><b>输出：</b></span>3</p>

<p><b>解释：</b></p>

<p>删除元素&nbsp;<code>nums[2] == -1</code>&nbsp;和&nbsp;<code>nums[3] == -2</code>&nbsp;，从&nbsp;<code>[1, 2, 1, 0, -1]</code>&nbsp;中选中子数组&nbsp;<code>[2, 1]</code>&nbsp;以获得最大元素和。</p>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>-100 &lt;= nums[i] &lt;= 100</code></li>
</ul>

#### 思路

负数肯定不能选，先都去掉。
「子数组中的所有元素互不相同」意味着每个非负数只能选一个，所以答案就是 $\textit{nums}$ 中的非负数去重后的元素和。
⚠**注意**：如果数组中都是负数，由于题目规定不能都去掉，只能选一个最大的负数（绝对值最小的负数）作为答案。

```
func maxSum(nums []int) (ans int) {
	set := map[int]struct{}{}
	mx := math.MinInt
	for _, x := range nums { // 一次遍历
		if x < 0 {
			mx = max(mx, x)
		} else if _, ok := set[x]; !ok {
			set[x] = struct{}{}
			ans += x
		}
	}
	if len(set) == 0 {
		return mx
	}
	return
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