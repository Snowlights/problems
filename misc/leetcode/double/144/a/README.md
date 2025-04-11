### 题目

<p>Alice 和 Bob 在玩一个游戏，他们俩轮流从一堆石头中移除石头，Alice 先进行操作。</p>

<ul>
	<li>Alice 在第一次操作中移除 <strong>恰好</strong>&nbsp;10 个石头。</li>
	<li>接下来的每次操作中，每位玩家移除的石头数 <strong>恰好</strong>&nbsp;为另一位玩家上一次操作的石头数减 1 。</li>
</ul>

<p>第一位没法进行操作的玩家输掉这个游戏。</p>

<p>给你一个正整数&nbsp;<code>n</code>&nbsp;表示一开始石头的数目，如果 Alice 赢下这个游戏，请你返回&nbsp;<code>true</code>&nbsp;，否则返回 <code>false</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 12</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>Alice 第一次操作中移除 10 个石头，剩下 2 个石头给 Bob 。</li>
	<li>Bob 无法移除 9 个石头，所以 Alice 赢下游戏。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 1</span></p>

<p><span class="example-io"><b>输出：</b>false</span></p>

<p><b>解释：</b></p>

<ul>
	<li>Alice 无法移除 10 个石头，所以 Alice 输掉游戏。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 50</code></li>
</ul>


### 思路

```
func canAliceWin(n int) bool {
	pick := 10
	for n >= pick {
		n -= pick
		pick--
	}
	return (10-pick)%2 > 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(C)$。其中 $C=10$。
- 空间复杂度：$\mathcal{O}(1)$。


## 方法二：数学公式

两人交替操作，设两人一共操作了 $x$ 次。

那么有

$$
10+9+8+\cdots+(11-x) = \dfrac{(21-x)x}{2} \le n
$$

解一元二次不等式，得

$$
x \le \dfrac{21 - \sqrt {21^2-8n}}{2}
$$

由于 $x$ 是整数，所以 $x$ 的最大值为

$$
\left\lfloor\dfrac{21 - \sqrt {21^2-8n}}{2}\right\rfloor = \left\lfloor\dfrac{21 - \lceil\sqrt {21^2-8n}\rceil}{2}\right\rfloor
$$

如果 $x$ 是奇数，Alice 胜，否则 Bob 胜。

```
func canAliceWin(n int) bool {
	x := (21 - int(math.Ceil(math.Sqrt(float64(441-n*8))))) / 2
	return x%2 > 0
}
```

#### 复杂度分析

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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)