#### 题目

<p>有一个&nbsp;<code>m x n</code>&nbsp;大小的矩形蛋糕，需要切成&nbsp;<code>1 x 1</code>&nbsp;的小块。</p>

<p>给你整数&nbsp;<code>m</code>&nbsp;，<code>n</code>&nbsp;和两个数组：</p>

<ul>
	<li><code>horizontalCut</code> 的大小为&nbsp;<code>m - 1</code>&nbsp;，其中&nbsp;<code>horizontalCut[i]</code>&nbsp;表示沿着水平线 <code>i</code>&nbsp;切蛋糕的开销。</li>
	<li><code>verticalCut</code> 的大小为&nbsp;<code>n - 1</code>&nbsp;，其中&nbsp;<code>verticalCut[j]</code>&nbsp;表示沿着垂直线&nbsp;<code>j</code>&nbsp;切蛋糕的开销。</li>
</ul>

<p>一次操作中，你可以选择任意不是&nbsp;<code>1 x 1</code>&nbsp;大小的矩形蛋糕并执行以下操作之一：</p>

<ol>
	<li>沿着水平线&nbsp;<code>i</code>&nbsp;切开蛋糕，开销为&nbsp;<code>horizontalCut[i]</code>&nbsp;。</li>
	<li>沿着垂直线&nbsp;<code>j</code>&nbsp;切开蛋糕，开销为&nbsp;<code>verticalCut[j]</code>&nbsp;。</li>
</ol>

<p>每次操作后，这块蛋糕都被切成两个独立的小蛋糕。</p>

<p>每次操作的开销都为最开始对应切割线的开销，并且不会改变。</p>

<p>请你返回将蛋糕全部切成&nbsp;<code>1 x 1</code>&nbsp;的蛋糕块的&nbsp;<strong>最小</strong>&nbsp;总开销。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>m = 3, n = 2, horizontalCut = [1,3], verticalCut = [5]</span></p>

<p><span class="example-io"><b>输出：</b>13</span></p>

<p><strong>解释：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2024/06/04/ezgifcom-animated-gif-maker-1.gif" style="width: 280px; height: 320px;" /></p>

<ul>
	<li>沿着垂直线 0 切开蛋糕，开销为 5 。</li>
	<li>沿着水平线 0 切开&nbsp;<code>3 x 1</code>&nbsp;的蛋糕块，开销为 1 。</li>
	<li>沿着水平线 0 切开 <code>3 x 1</code>&nbsp;的蛋糕块，开销为 1 。</li>
	<li>沿着水平线 1 切开 <code>2 x 1</code>&nbsp;的蛋糕块，开销为 3 。</li>
	<li>沿着水平线 1 切开 <code>2 x 1</code>&nbsp;的蛋糕块，开销为 3 。</li>
</ul>

<p>总开销为&nbsp;<code>5 + 1 + 1 + 3 + 3 = 13</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>m = 2, n = 2, horizontalCut = [7], verticalCut = [4]</span></p>

<p><span class="example-io"><b>输出：</b>15</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>沿着水平线 0 切开蛋糕，开销为 7 。</li>
	<li>沿着垂直线 0 切开&nbsp;<code>1 x 2</code>&nbsp;的蛋糕块，开销为 4 。</li>
	<li>沿着垂直线 0 切开&nbsp;<code>1 x 2</code>&nbsp;的蛋糕块，开销为 4 。</li>
</ul>

<p>总开销为&nbsp;<code>7 + 4 + 4 = 15</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= m, n &lt;= 20</code></li>
	<li><code>horizontalCut.length == m - 1</code></li>
	<li><code>verticalCut.length == n - 1</code></li>
	<li><code>1 &lt;= horizontalCut[i], verticalCut[i] &lt;= 10<sup>3</sup></code></li>
</ul>

#### 思路

首先，每条水平线和垂直线，最终都要全部切完。

- 水平线（横切）开销 $\textit{horizontalCut}[i]$ 对答案的**贡献**，等于 $\textit{horizontalCut}[i]$ 乘以横切次数（经过多少块蛋糕），即在此之前的竖切次数加一。
- 垂直线（竖切）开销 $\textit{verticalCut}[j]$ 对答案的**贡献**，等于 $\textit{verticalCut}[j]$ 乘以竖切次数（经过多少块蛋糕），即在此之前的横切次数加一。

⚠注意：本题虽然不要求「一切到底」，但计算贡献的式子是不变的。无论是否一切到底，最终得到的式子，合并同类项之后，仍然是「横切开销」乘以「横切次数」，以及「竖切开销」乘以「竖切次数」，**计算结果与一切到底是一样的**。所以下面的分析，**只需考虑一切到底的情况**。

以示例 1 为例，其操作序列为

$$
竖切\ 0,\ 横切\ 0,\ 横切\ 1
$$

最小总开销为

$$
\textit{verticalCut}[0]\cdot 1 + \textit{horizontalCut}[0] \cdot 2 + \textit{horizontalCut}[1] \cdot 2
$$

对于一个操作序列，交换其中两个相邻的横切，不影响答案；交换其中两个相邻的竖切，也不影响答案。所以重点考察交换两个相邻的横切和竖切。

> 为方便讨论，交换相邻的，也可以交换不相邻的，最终得出的结论是一样的。

设横切的开销为 $h$，如果先横切，设需要横切 $\textit{cntH}$ 次。

设竖切的开销为 $v$，如果先竖切，设需要竖切 $\textit{cntV}$ 次。

- 先横切，再竖切，那么竖切的次数（这一刀经过的蛋糕块数）要多 $1$，开销为
  $$
  h\cdot \textit{cntH} + v\cdot (\textit{cntV}+1)
  $$
- 先竖切，再横切，那么横切的次数（这一刀经过的蛋糕块数）要多 $1$，开销为
  $$
  v\cdot \textit{cntV} + h\cdot (\textit{cntH}+1)
  $$

如果先横再竖开销更小，则有

$$
h\cdot \textit{cntH} + v\cdot (\textit{cntV}+1) < v\cdot \textit{cntV} + h\cdot (\textit{cntH}+1)
$$

化简得

$$
h > v
$$

这意味着，**谁的开销更大，就先切谁**，并且这个先后顺序与 $\textit{cntH}$ 和 $\textit{cntV}$ 无关。换句话说，按照该规则去切蛋糕，得到的操作序列，如果把开销大的操作移动后面，必然会得到更大的总开销。

## 写法一

1. 把 $\textit{horizontalCut}$ 和 $\textit{verticalCut}$ 从大到小排序。
2. 初始化 $\textit{cntH} = 1, \textit{cntV} = 1, i = 0, j = 0$。
3. 双指针遍历 $\textit{horizontalCut}$ 和 $\textit{verticalCut}$。
4. 如果 $\textit{horizontalCut}[i] > \textit{verticalCut}[j]$，那么优先横切，把 $\textit{horizontalCut}[i]\cdot \textit{cntH}$ 加入答案，$i$ 增加 $1$，然后需要竖切的次数增加，把 $\textit{cntV}$ 增加 $1$；否则优先竖切，把 $\textit{verticalCut}[j]\cdot \textit{cntV}$ 加入答案，$j$ 增加 $1$，然后需要横切的次数增加，把 $\textit{cntH}$ 增加 $1$。
5. 循环直到两个数组都遍历完。
6. 返回答案。

注意 $i=m-1$ 和 $j=n-1$ 的情况。

```
func minimumCost(m, n int, horizontalCut, verticalCut []int) int64 {
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	cntH, cntV := 1, 1
	i, j := 0, 0
	for i < m-1 || j < n-1 {
		if j == n-1 || i < m-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * cntH // 横切
			i++
			cntV++ // 需要竖切的蛋糕块增加
		} else {
			ans += verticalCut[j] * cntV // 竖切
			j++
			cntH++ // 需要横切的蛋糕块增加
		}
	}
	return int64(ans)
}
```

## 写法二（优化）

$\textit{cntH}$ 和 $\textit{cntV}$ 这两个变量可以省略，因为从上面的过程可以发现，$\textit{cntH}=j+1,\ \textit{cntV}=i+1$。

```
func minimumCost(m, n int, horizontalCut, verticalCut []int) int64 {
	slices.SortFunc(horizontalCut, func(a, b int) int { return b - a })
	slices.SortFunc(verticalCut, func(a, b int) int { return b - a })
	ans := 0
	i, j := 0, 0
	for i < m-1 || j < n-1 {
		if j == n-1 || i < m-1 && horizontalCut[i] > verticalCut[j] {
			ans += horizontalCut[i] * (j + 1) // 横切
			i++
		} else {
			ans += verticalCut[j] * (i + 1) // 竖切
			j++
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m\log m + n\log n)$。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。排序的栈开销不计入。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)