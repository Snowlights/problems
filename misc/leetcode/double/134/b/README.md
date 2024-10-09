#### 题目

<p>给你一个下标从 <strong>0</strong>&nbsp;开始的整数数组&nbsp;<code>enemyEnergies</code>&nbsp;，它表示一个下标从 <strong>0</strong>&nbsp;开始的敌人能量数组。</p>

<p>同时给你一个整数&nbsp;<code>currentEnergy</code>&nbsp;，它表示你一开始拥有的能量值总量。</p>

<p>你一开始的分数为&nbsp;<code>0</code>&nbsp;，且一开始所有的敌人都未标记。</p>

<p>你可以通过以下操作 <b>之一</b>&nbsp;<strong>任意次</strong>（也可以&nbsp;<strong>0</strong>&nbsp;次）来得分：</p>

<ul>
	<li>选择一个 <strong>未标记</strong>&nbsp;且满足&nbsp;<code>currentEnergy &gt;= enemyEnergies[i]</code>&nbsp;的敌人&nbsp;<code>i</code>&nbsp;。在这个操作中：
	<ul>
		<li>你会获得 <code>1</code>&nbsp;分。</li>
		<li>你的能量值减少&nbsp;<code>enemyEnergies[i]</code>&nbsp;，也就是说&nbsp;<code>currentEnergy = currentEnergy - enemyEnergies[i]</code>&nbsp;。</li>
	</ul>
	</li>
	<li>如果你目前&nbsp;<strong>至少</strong>&nbsp;有 <code>1</code>&nbsp;分，你可以选择一个&nbsp;<strong>未标记</strong>&nbsp;的敌人&nbsp;<code>i</code>&nbsp;。在这个操作中：
	<ul>
		<li>你的能量值增加 <code>enemyEnergies[i]</code>&nbsp;，也就是说&nbsp;<code>currentEnergy = currentEnergy + enemyEnergies[i]</code>&nbsp;。</li>
		<li>敌人&nbsp;<code>i</code> <strong>被标记</strong>&nbsp;。</li>
	</ul>
	</li>
</ul>

<p>请你返回通过以上操作，<strong>最多</strong>&nbsp;可以获得多少分。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><b>输入：</b>enemyEnergies = [3,2,2], currentEnergy = 2</p>

<p><b>输出：</b>3</p>

<p><strong>解释：</strong></p>

<p>通过以下操作可以得到最大得分 3 分：</p>

<ul>
	<li>对敌人 1 使用第一种操作：<code>points</code>&nbsp;增加 1 ，<code>currentEnergy</code>&nbsp;减少 2 。所以&nbsp;<code>points = 1</code>&nbsp;且&nbsp;<code>currentEnergy = 0</code>&nbsp;。</li>
	<li>对敌人 0 使用第二种操作：<code>currentEnergy</code>&nbsp;增加 3 ，敌人 0 被标记。所以&nbsp;<code>points = 1</code>&nbsp;，<code>currentEnergy = 3</code>&nbsp;，被标记的敌人包括&nbsp;<code>[0]</code>&nbsp;。</li>
	<li>对敌人 2 使用第一种操作：<code>points</code>&nbsp;增加 1 ，<code>currentEnergy</code>&nbsp;减少 2 。所以&nbsp;<code>points = 2</code>&nbsp;且&nbsp;<code>currentEnergy = 1</code>&nbsp;，被标记的敌人包括<code>[0]</code>&nbsp;。</li>
	<li>对敌人 2 使用第二种操作：<code>currentEnergy</code>&nbsp;增加 2 ，敌人 2 被标记。所以&nbsp;<code>points = 2</code>&nbsp;，<code>currentEnergy = 3</code>&nbsp;且被标记的敌人包括&nbsp;<code>[0, 2]</code>&nbsp;。</li>
	<li>对敌人 1 使用第一种操作：<code>points</code>&nbsp;增加 1 ，<code>currentEnergy</code>&nbsp;减少 2 。所以&nbsp;<code>points = 3</code>&nbsp;，<code>currentEnergy = 1</code>&nbsp;，被标记的敌人包括&nbsp;<code>[0, 2]</code>&nbsp;。</li>
</ul>

<p><strong>示例 2：</strong></p>

<p><b>输入：</b>enemyEnergies =&nbsp;[2], currentEnergy = 10</p>

<p><b>输出：</b>5</p>

<p><strong>解释：</strong></p>

<p>通过对敌人 0 进行第一种操作 5 次，得到最大得分。</p>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= enemyEnergies.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= enemyEnergies[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= currentEnergy &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

设 $\textit{enemyEnergies}$ 中的最小值为 $\textit{mn}$，元素和为 $s$。

如果 $\textit{currentEnergy} < \textit{mn}$，那么操作一无法执行，无法得到任何分数，所以也无法执行操作二，返回 $0$。
否则，操作顺序如下：
1. 对 $\textit{mn}$ 执行操作一，得到 $1$ 分。
2. 对除了 $\textit{mn}$ 以外的敌人执行操作二，得到 $s - \textit{mn}$ 的能量。
3. 对 $\textit{mn}$ 执行操作一，直到能量不足。

也可以理解为，先得到 $s - \textit{mn}$ 的能量，再不断对 $\textit{mn}$ 执行操作一，所以得分为

$$
\left\lfloor\dfrac{\textit{currentEnergy} + s - \textit{mn}}{\textit{mn}}\right\rfloor
$$

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{enemyEnergies}$ 的长度。
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