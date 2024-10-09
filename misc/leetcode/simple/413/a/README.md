#### 题目

<p>给你两个字符串 <code>coordinate1</code> 和 <code>coordinate2</code>，代表 <code>8 x 8</code> 国际象棋棋盘上的两个方格的坐标。</p>

<p>以下是棋盘的参考图。</p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/07/17/screenshot-2021-02-20-at-22159-pm.png" style="width: 400px; height: 396px;" /></p>

<p>如果这两个方格颜色相同，返回 <code>true</code>，否则返回 <code>false</code>。</p>

<p>坐标总是表示有效的棋盘方格。坐标的格式总是先字母（表示列），再数字（表示行）。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">coordinate1 = "a1", coordinate2 = "c3"</span></p>

<p><strong>输出：</strong> <span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<p>两个方格均为黑色。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">coordinate1 = "a1", coordinate2 = "h3"</span></p>

<p><strong>输出：</strong> <span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<p>方格 <code>"a1"</code> 是黑色，而 <code>"h3"</code> 是白色。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>coordinate1.length == coordinate2.length == 2</code></li>
	<li><code>'a' &lt;= coordinate1[0], coordinate2[0] &lt;= 'h'</code></li>
	<li><code>'1' &lt;= coordinate1[1], coordinate2[1] &lt;= '8'</code></li>
</ul>


#### 思路

把 $\textit{coordinate}_1$ 和 $\textit{coordinate}_2$ 简记为 $s$ 和 $t$。

根据题目中的图片，如果 $s[0]$ 和 $s[1]$ 的 ASCII 值的奇偶性相同，那么格子是黑格，否则是白格。

进一步地，由于奇数+奇数=偶数，偶数+偶数=偶数，所以如果 $(s[0] + s[1])\bmod 2$ 是偶数，那么格子是黑格；否则奇数+偶数=奇数，格子是白格。

如果

$$
(s[0] + s[1])\bmod 2 = (t[0] + t[1])\bmod 2
$$

那么两个格子颜色相同，否则不同。

也可以取 $(s[0] \oplus s[1])$ 的最低位，其中 $\oplus$ 是异或运算。

```
func checkTwoChessboards(s, t string) bool {
	return (s[0]^s[1])&1 == (t[0]^t[1])&1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。


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