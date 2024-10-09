### 题目

<p>给你一个二进制数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>你可以对数组执行以下操作 <strong>任意</strong>&nbsp;次（也可以 0 次）：</p>

<ul>
	<li>选择数组中 <strong>任意连续</strong>&nbsp;3 个元素，并将它们 <strong>全部反转</strong>&nbsp;。</li>
</ul>

<p><strong>反转</strong>&nbsp;一个元素指的是将它的值从 0 变 1 ，或者从 1 变 0 。</p>

<p>请你返回将 <code>nums</code>&nbsp;中所有元素变为 1 的 <strong>最少</strong>&nbsp;操作次数。如果无法全部变成 1 ，返回 -1 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [0,1,1,1,0,0]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong><br />
我们可以执行以下操作：</p>

<ul>
	<li>选择下标为 0 ，1 和 2 的元素并反转，得到&nbsp;<code>nums = [<u><strong>1</strong></u>,<u><strong>0</strong></u>,<u><strong>0</strong></u>,1,0,0]</code>&nbsp;。</li>
	<li>选择下标为 1 ，2 和 3 的元素并反转，得到&nbsp;<code>nums = [1,<u><strong>1</strong></u>,<u><strong>1</strong></u>,<strong><u>0</u></strong>,0,0]</code>&nbsp;。</li>
	<li>选择下标为 3 ，4 和 5 的元素并反转，得到&nbsp;<code>nums = [1,1,1,<strong><u>1</u></strong>,<u><strong>1</strong></u>,<u><strong>1</strong></u>]</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [0,1,1,1]</span></p>

<p><span class="example-io"><b>输出：</b>-1</span></p>

<p><strong>解释：</strong><br />
无法将所有元素都变为 1 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 1</code></li>
</ul>

### 思路

分类讨论：

- 如果 $\textit{nums}[0]=1$，无需修改，问题变成剩下 $n-1$ 个数如何操作。接下来考虑 $\textit{nums}[1]$。
- 如果 $\textit{nums}[0]=0$，修改，问题变成剩下 $n-1$ 个数如何操作。接下来考虑 $\textit{nums}[1]$。

所以从左到右遍历数组，一边遍历一边修改。

``` 
func minOperations(nums []int) (ans int) {
	n := len(nums)
	for i, x := range nums[:n-2] {
		if x == 0 {
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			ans++
		}
	}
	if nums[n-2] == 0 || nums[n-1] == 0 {
		return -1
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
