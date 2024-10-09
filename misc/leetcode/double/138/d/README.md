#### 题目

<p>给你一个整数&nbsp;<code>power</code>&nbsp;和两个整数数组&nbsp;<code>damage</code> 和&nbsp;<code>health</code>&nbsp;，两个数组的长度都为&nbsp;<code>n</code>&nbsp;。</p>

<p>Bob 有&nbsp;<code>n</code>&nbsp;个敌人，如果第&nbsp;<code>i</code>&nbsp;个敌人还活着（也就是健康值&nbsp;<code>health[i] &gt; 0</code>&nbsp;的时候），每秒钟会对 Bob 造成&nbsp;<code>damage[i]</code>&nbsp;<strong>点</strong>&nbsp;伤害。</p>

<p>每一秒中，在敌人对 Bob 造成伤害 <strong>之后</strong>&nbsp;，Bob 会选择 <strong>一个</strong>&nbsp;还活着的敌人进行攻击，该敌人的健康值减少 <code>power</code>&nbsp;。</p>

<p>请你返回 Bob 将 <strong>所有</strong>&nbsp;<code>n</code>&nbsp;个敌人都消灭之前，<strong>最少</strong>&nbsp;会收到多少伤害。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>power = 4, damage = [1,2,3,4], health = [4,5,6,8]</span></p>

<p><span class="example-io"><b>输出：</b>39</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>最开始 2 秒内都攻击敌人 3 ，然后敌人 3 会被消灭，这段时间内对 Bob 的总伤害是&nbsp;<code>10 + 10 = 20</code>&nbsp;点。</li>
	<li>接下来 2 秒内都攻击敌人 2 ，然后敌人 2 会被消灭，这段时间内对 Bob 的总伤害是&nbsp;<code>6 + 6 = 12</code>&nbsp;点。</li>
	<li>接下来 1 秒内都攻击敌人 0 ，然后敌人 0 会被消灭，这段时间内对 Bob 的总伤害是&nbsp;<code>3</code>&nbsp;点。</li>
	<li>接下来 2 秒内都攻击敌人 1 ，然后敌人 1 会被消灭，这段时间内对 Bob 的总伤害是&nbsp;<code>2 + 2 = 4</code>&nbsp;点。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>power = 1, damage = [1,1,1,1], health = [1,2,3,4]</span></p>

<p><span class="example-io"><b>输出：</b>20</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>最开始 1 秒内都攻击敌人 0 ，然后敌人 0 会被消灭，这段时间对 Bob 的总伤害是&nbsp;<code>4</code>&nbsp;点。</li>
	<li>接下来 2 秒内都攻击敌人 1 ，然后敌人 1 会被消灭，这段时间对 Bob 的总伤害是&nbsp;<code>3 + 3 = 6</code>&nbsp;点。</li>
	<li>接下来 3 秒内都攻击敌人 2 ，然后敌人 2 会被消灭，这段时间对 Bob 的总伤害是&nbsp;<code>2 + 2 + 2 = 6</code>&nbsp;点。</li>
	<li>接下来 4 秒内都攻击敌人 3 ，然后敌人 3 会被消灭，这段时间对 Bob 的总伤害是&nbsp;<code>1 + 1 + 1 + 1 = 4</code>&nbsp;点。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>power = 8, damage = [40], health = [59]</span></p>

<p><span class="example-io"><b>输出：</b>320</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= power &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= n == damage.length == health.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= damage[i], health[i] &lt;= 10<sup>4</sup></code></li>
</ul>

#### 思路

首先，一直攻击同一个敌人，相比来回攻击多个敌人（雨露均沾）更好，因为这样我们被敌人攻击的次数更少。

**从特殊到一般**，想一想，如果只有**两个**敌人 A 和 B，我们应该先攻击谁？

消灭 A 需要攻击的次数为

$$
k_A = \left\lceil\dfrac{\textit{health}_A}{\textit{power}}\right\rceil = \left\lfloor\dfrac{\textit{health}_A-1}{\textit{power}}\right\rfloor + 1
$$

> 讨论 $\textit{health}_A$ 被 $\textit{power}$ 整除，和不被 $\textit{power}$ 整除两种情况，可以证明上式的正确性。

同理可得消灭 B 需要的攻击次数，记作 $k_B$。

如果先消灭 A，再消灭 B，那么受到的伤害总和为

$$
k_A\cdot \textit{damage}_A + (k_A+k_B)\cdot \textit{damage}_B
$$

如果先消灭 B，再消灭 A，那么受到的伤害总和为

$$
k_B\cdot \textit{damage}_B + (k_A+k_B)\cdot \textit{damage}_A
$$

如果先消灭 A，再消灭 B 更好，那么有

$$
k_A\cdot \textit{damage}_A + (k_A+k_B)\cdot \textit{damage}_B < k_B\cdot \textit{damage}_B + (k_A+k_B)\cdot \textit{damage}_A
$$

化简得

$$
k_A\cdot \textit{damage}_B < k_B\cdot \textit{damage}_A
$$

> 也就是优先消灭 $k/\textit{damage}$ 更小的敌人。结合你玩过的某些游戏，应当优先消灭那个**又脆皮伤害又高**的敌人。

推广到更多的敌人，可以按照上式对 $\textit{damage}$ 和 $\textit{health}$ 排序，理由如下。

先假定按照输入的顺序消灭敌人。如果发现相邻两个敌人不满足上面的不等式，就交换这两个敌人的位置，这可以让受到的总伤害变小。

不断交换敌人，直到所有相邻敌人都满足上面的不等式。

本质上来说，这个不断交换相邻敌人的过程，和**冒泡排序**是一样的。那么按照不等式对数组排序即可。

排序后，按照顺序消灭敌人。用一个变量 $s$ 维护从一开始到击败当前敌人，所经过的秒数。把 $s\cdot \textit{damage}[i]$ 加入答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ajHYeoEG5/) 第四题，欢迎点赞关注~

## 答疑

**问**：本题可以用 $k/\textit{damage}$ 比大小吗？除法的结果需要用浮点数表示，精度误差是否会影响结果？

**答**：本题 $10^5$ 的值域范围是可以的，但如果范围是 $10^9$ 就可能会有问题，见 [2280. 表示一个折线图的最少线段数](https://leetcode.cn/problems/minimum-lines-to-represent-a-line-chart/)，可以用浮点数写试试。

```
func minDamage(power int, damage, health []int) (ans int64) {
	type pair struct{ k, d int }
	a := make([]pair, len(health))
	for i, h := range health {
		a[i] = pair{(h-1)/power + 1, damage[i]}
	}
	slices.SortFunc(a, func(p, q pair) int { return p.k*q.d - q.k*p.d })

	s := 0
	for _, p := range a {
		s += p.k
		ans += int64(s) * int64(p.d)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{damage}$ 的长度。瓶颈在排序上。
- 时间复杂度：$\mathcal{O}(n)$。

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