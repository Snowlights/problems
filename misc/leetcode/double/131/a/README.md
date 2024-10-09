### 题目

<p>给你一个数组&nbsp;<code>nums</code>&nbsp;，数组中的数字 <strong>要么</strong> 出现一次，<strong>要么</strong>&nbsp;出现两次。</p>

<p>请你返回数组中所有出现两次数字的按位<em>&nbsp;</em><code>XOR</code>&nbsp;值，如果没有数字出现过两次，返回 0 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,1,3]</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;中唯一出现过两次的数字是 1 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,3]</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><strong>解释：</strong></p>

<p><code>nums</code>&nbsp;中没有数字出现两次。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,2,2,1]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<p>数字 1 和&nbsp;2 出现过两次。<code>1 XOR 2 == 3</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 50</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 50</code></li>
	<li><code>nums</code>&nbsp;中每个数字要么出现过一次，要么出现过两次。</li>
</ul>

### 思路

遍历 $\textit{nums}$，同时用一个 $\textit{vis}$ 集合记录遇到的数字。

- 设 $x=\textit{nums}[i]$。
- 如果 $x$ 不在 $\textit{vis}$ 中，说明是第一次遇到，加入 $\textit{vis}$。
- 如果 $x$ 在 $\textit{vis}$ 中，说明是第二次遇到（注意每个数至多出现两次），加入答案（异或）。



``` go  
func duplicateNumbersXOR(nums []int) (ans int) {
	vis := 0
	for _, x := range nums {
		if vis>>x&1 > 0 {
			ans ^= x
		} else {
			vis |= 1 << x
		}
	}
	return
}
```

### 复杂度分析

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