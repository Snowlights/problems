### 题目

<p>给你两个 <strong>正</strong>&nbsp;整数&nbsp;<code>x</code>&nbsp;和&nbsp;<code>y</code>&nbsp;，分别表示价值为 75 和 10 的硬币的数目。</p>

<p>Alice 和 Bob 正在玩一个游戏。每一轮中，Alice&nbsp;先进行操作，Bob 后操作。每次操作中，玩家需要拿出价值 <b>总和</b>&nbsp;为 115 的硬币。如果一名玩家无法执行此操作，那么这名玩家 <strong>输掉</strong>&nbsp;游戏。</p>

<p>两名玩家都采取 <strong>最优</strong>&nbsp;策略，请你返回游戏的赢家。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>x = 2, y = 7</span></p>

<p><span class="example-io"><b>输出：</b>"Alice"</span></p>

<p><strong>解释：</strong></p>

<p>游戏一次操作后结束：</p>

<ul>
	<li>Alice 拿走 1 枚价值为 75 的硬币和 4 枚价值为 10 的硬币。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>x = 4, y = 11</span></p>

<p><span class="example-io"><b>输出：</b>"Bob"</span></p>

<p><strong>解释：</strong></p>

<p>游戏 2 次操作后结束：</p>

<ul>
	<li>Alice 拿走&nbsp;1 枚价值为 75 的硬币和 4 枚价值为 10 的硬币。</li>
	<li>Bob 拿走&nbsp;1 枚价值为 75 的硬币和 4 枚价值为 10 的硬币。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= x, y &lt;= 100</code></li>
</ul>


### 思路

因为 $10$ 的倍数不可能等于 $115$，所以面额为 $10$ 的硬币不能单独选，至少要选 $1$ 个面额为 $75$ 的硬币。

又由于 $75\cdot 2=150>115$，所以面额为 $75$ 的硬币要**恰好**选 $1$ 个。

由于 $75+10\cdot 4 = 115$，所以面额为 $10$ 的硬币要**恰好**选 $4$ 个。

> 本质上来说，我们在求解二元一次不定方程 $75a+10b=115$，它有唯一正整数解 $a=1,b=4$。

如果一开始 Alice 就没法选，或者偶数轮后 Alice 没法选，那么 Bob 胜出，否则 Alice 胜出。

## 优化

设 $k = \min(x, \lfloor y/4 \rfloor)$，这是能玩的回合数，判断 $k$ 的奇偶性即可。

```
func losingPlayer(x, y int) string {
	return [2]string{"Bob", "Alice"}[min(x, y/4)%2]
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
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