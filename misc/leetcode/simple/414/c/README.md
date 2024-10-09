#### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>你的目标是从下标 <code>0</code>&nbsp;出发，到达下标 <code>n - 1</code>&nbsp;处。每次你只能移动到&nbsp;<strong>更大</strong>&nbsp;的下标处。</p>

<p>从下标 <code>i</code>&nbsp;跳到下标 <code>j</code>&nbsp;的得分为&nbsp;<code>(j - i) * nums[i]</code>&nbsp;。</p>

<p>请你返回你到达最后一个下标处能得到的 <strong>最大总得分</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,3,1,5]</span></p>

<p><b>输出：</b>7</p>

<p><b>解释：</b></p>

<p>一开始跳到下标 1 处，然后跳到最后一个下标处。总得分为&nbsp;<code>1 * 1 + 2 * 3 = 7</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [4,3,1,3,2]</span></p>

<p><b>输出：</b>16</p>

<p><strong>解释：</strong></p>

<p>直接跳到最后一个下标处。总得分为&nbsp;<code>4 * 4 = 16</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>

#### 思路

![lc414c-c.png](https://pic.leetcode.cn/1725769175-EJDMxR-lc414c-c.png)

代码实现时，可以视作有 $n-1$ 个底边长为 $1$ 的矩形。

最大化每个矩形的高，即可最大化所有矩形的面积之和。

从左到右遍历，维护遍历到的数的最大值，作为矩形的高。


```
func findMaximumScore(nums []int) (ans int64) {
	mx := 0
	for _, x := range nums[:len(nums)-1] {
		mx = max(mx, x)
		ans += int64(mx)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。Python 忽略切片的空间。


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
