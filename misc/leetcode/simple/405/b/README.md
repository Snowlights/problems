#### 题目

<p>给你一个正整数 <code>n</code>。</p>

<p>如果一个二进制字符串 <code>x</code> 的所有长度为 2 的<span data-keyword="substring-nonempty">子字符串</span>中包含 <strong>至少</strong> 一个 <code>"1"</code>，则称 <code>x</code> 是一个<strong> 有效</strong> 字符串。</p>

<p>返回所有长度为 <code>n</code> 的<strong> 有效</strong> 字符串，可以以任意顺序排列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 3</span></p>

<p><strong>输出：</strong> <span class="example-io">["010","011","101","110","111"]</span></p>

<p><strong>解释：</strong></p>

<p>长度为 3 的有效字符串有：<code>"010"</code>、<code>"011"</code>、<code>"101"</code>、<code>"110"</code> 和 <code>"111"</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">n = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">["0","1"]</span></p>

<p><strong>解释：</strong></p>

<p>长度为 1 的有效字符串有：<code>"0"</code> 和 <code>"1"</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 18</code></li>
</ul>

#### 思路

枚举 $[0,2^n-1]$ 中的 $i$，如果 $i$ 的长为 $n$ 的二进制中，没有相邻的 $0$，那么将其二进制字符串加入答案。
怎么判断二进制中是否有相邻的 $0$？
我们可以把 $i$ 取反（保留低 $n$ 位），记作 $x$。问题变成：判断 $x$ 中是否有相邻的 $1$。
这可以用 `x & (x >> 1)` 来判断，如果这个值等于 $0$，则说明 $x$ 中没有相邻的 $1$。

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(2^n)$。更细致的分析表明（见视频讲解），答案的长度是**斐波那契数**，所以枚举 $i$ 的复杂度更高，时间复杂度为 $\mathcal{O}(2^n)$。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)