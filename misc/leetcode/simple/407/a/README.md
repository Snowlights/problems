#### 题目

<p>给你两个正整数 <code>n</code> 和 <code>k</code>。</p>

<p>你可以选择 <code>n</code> 的 <strong>二进制表示</strong> 中任意一个值为 1 的位，并将其改为 0。</p>

<p>返回使得 <code>n</code> 等于 <code>k</code> 所需要的更改次数。如果无法实现，返回 -1。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 13, k = 4</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong><br />
最初，<code>n</code> 和 <code>k</code> 的二进制表示分别为 <code>n = (1101)<sub>2</sub></code> 和 <code>k = (0100)<sub>2</sub></code>，</p>

<p>我们可以改变 <code>n</code> 的第一位和第四位。结果整数为 <code>n = (<u><strong>0</strong></u>10<u><strong>0</strong></u>)<sub>2</sub> = k</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 21, k = 21</span></p>

<p><strong>输出：</strong> <span class="example-io">0</span></p>

<p><strong>解释：</strong><br />
<code>n</code> 和 <code>k</code> 已经相等，因此不需要更改。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 14, k = 13</span></p>

<p><strong>输出：</strong> <span class="example-io">-1</span></p>

<p><strong>解释：</strong><br />
无法使 <code>n</code> 等于 <code>k</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n, k &lt;= 10<sup>6</sup></code></li>
</ul>

#### 思路

**请先阅读**：[从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)
从集合的角度理解，$k$ 必须是 $n$ 的**子集**。如果不是，返回 $-1$。怎么用位运算判断，见上面的文章链接。
如果 $k$ 是 $n$ 的子集，答案为从 $n$ 中去掉 $k$ 后的集合大小，即 $n\oplus k$ 的二进制中的 $1$ 的个数。

> 注：也可以计算 $n-k$ 的二进制中的 $1$ 的个数。

```
func minChanges(n, k int) int {
	if n&k != k {
		return -1
	}
	return bits.OnesCount(uint(n ^ k))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。
- 
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)