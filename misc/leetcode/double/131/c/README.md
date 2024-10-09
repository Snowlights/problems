### 题目

<p>给你一个整数&nbsp;<code>limit</code>&nbsp;和一个大小为 <code>n x 2</code>&nbsp;的二维数组&nbsp;<code>queries</code>&nbsp;。</p>

<p>总共有&nbsp;<code>limit + 1</code>&nbsp;个球，每个球的编号为&nbsp;<code>[0, limit]</code>&nbsp;中一个&nbsp;<strong>互不相同</strong>&nbsp;的数字。一开始，所有球都没有颜色。<code>queries</code>&nbsp;中每次操作的格式为&nbsp;<code>[x, y]</code>&nbsp;，你需要将球&nbsp;<code>x</code>&nbsp;染上颜色&nbsp;<code>y</code>&nbsp;。每次操作之后，你需要求出所有球中&nbsp;<strong>不同</strong>&nbsp;颜色的数目。</p>

<p>请你返回一个长度为 <code>n</code>&nbsp;的数组&nbsp;<code>result</code>&nbsp;，其中&nbsp;<code>result[i]</code>&nbsp;是第 <code>i</code>&nbsp;次操作以后不同颜色的数目。</p>

<p><strong>注意</strong>&nbsp;，没有染色的球不算作一种颜色。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>limit = 4, queries = [[1,4],[2,5],[1,3],[3,4]]</span></p>

<p><span class="example-io"><b>输出：</b>[1,2,2,3]</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/04/17/ezgifcom-crop.gif" style="width: 455px; height: 145px;" /></p>

<ul>
	<li>操作 0&nbsp;后，球 1 颜色为 4 。</li>
	<li>操作 1 后，球 1 颜色为&nbsp;4 ，球 2 颜色为 5 。</li>
	<li>操作 2 后，球 1 颜色为 3 ，球 2 颜色为 5 。</li>
	<li>操作 3 后，球 1 颜色为 3 ，球 2 颜色为 5 ，球 3 颜色为 4 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>limit = 4, queries = [[0,1],[1,2],[2,2],[3,4],[4,5]]</span></p>

<p><span class="example-io"><b>输出：</b>[1,2,2,3,4]</span></p>

<p><strong>解释：</strong></p>

<p><strong><img alt="" src="https://assets.leetcode.com/uploads/2024/04/17/ezgifcom-crop2.gif" style="width: 457px; height: 144px;" /></strong></p>

<ul>
	<li>操作 0&nbsp;后，球 0&nbsp;颜色为 1&nbsp;。</li>
	<li>操作 1&nbsp;后，球 0&nbsp;颜色为 1 ，球 1 颜色为 2 。</li>
	<li>操作 2&nbsp;后，球 0&nbsp;颜色为 1 ，球 1 和 2&nbsp;颜色为 2 。</li>
	<li>操作 3 后，球 0&nbsp;颜色为 1 ，球 1 和 2&nbsp;颜色为 2 ，球 3 颜色为 4 。</li>
	<li>操作 4&nbsp;后，球 0&nbsp;颜色为 1 ，球 1 和 2&nbsp;颜色为 2 ，球 3 颜色为 4 ，球 4 颜色为 5 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= limit &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= n == queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 2</code></li>
	<li><code>0 &lt;= queries[i][0] &lt;= limit</code></li>
	<li><code>1 &lt;= queries[i][1] &lt;= 10<sup>9</sup></code></li>
</ul>


### 思路

用哈希表 $\textit{color}$ 维护第 $x$ 个球的颜色，哈希表 $\textit{cnt}$ 维护每种颜色的出现次数。

当把球 $x$ 的颜色改成 $y$ 时：

1. 如果 $x$ 在 $\textit{color}$ 中，设 $c=\textit{color}[x]$，先把 $\textit{cnt}[c]$ 减少一，如果 $\textit{cnt}[c]$ 变成 $0$，则从 $\textit{cnt}$ 中删除 $c$。
2. 更新 $\textit{color}[x] = y$。
3. 把 $\textit{cnt}[y]$ 增加一。
4. 把 $\textit{cnt}$ 的大小加入答案。

### 复杂度分析

- 时间复杂度：$\mathcal{O}(q)$，其中 $q$ 是 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(q)$。

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