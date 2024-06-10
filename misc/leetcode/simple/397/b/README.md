#### 题目

<p>在神秘的地牢中，<code>n</code> 个魔法师站成一排。每个魔法师都拥有一个属性，这个属性可以给你提供能量。有些魔法师可能会给你负能量，即从你身上吸取能量。</p>

<p>你被施加了一种诅咒，当你从魔法师 <code>i</code> 处吸收能量后，你将被立即传送到魔法师 <code>(i + k)</code> 处。这一过程将重复进行，直到你到达一个不存在 <code>(i + k)</code> 的魔法师为止。</p>

<p>换句话说，你将选择一个起点，然后以 <code>k</code> 为间隔跳跃，直到到达魔法师序列的末端，<strong>在过程中吸收所有的能量</strong>。</p>

<p>给定一个数组 <code>energy</code> 和一个整数<code>k</code>，返回你能获得的<strong> 最大 </strong>能量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="
    border-color: var(--border-tertiary);
    border-left-width: 2px;
    color: var(--text-secondary);
    font-size: .875rem;
    margin-bottom: 1rem;
    margin-top: 1rem;
    overflow: visible;
    padding-left: 1rem;
">
<p><strong>输入：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;
"> energy = [5,2,-10,-5,1], k = 3</span></p>

<p><strong>输出：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;
"> 3</span></p>

<p><strong>解释：</strong>可以从魔法师 1 开始，吸收能量 2 + 1 = 3。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="
    border-color: var(--border-tertiary);
    border-left-width: 2px;
    color: var(--text-secondary);
    font-size: .875rem;
    margin-bottom: 1rem;
    margin-top: 1rem;
    overflow: visible;
    padding-left: 1rem;
">
<p><strong>输入：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;
"> energy = [-2,-3,-1], k = 2</span></p>

<p><strong>输出：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;
"> -1</span></p>

<p><strong>解释：</strong>可以从魔法师 2 开始，吸收能量 -1。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= energy.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-1000 &lt;= energy[i] &lt;= 1000</code></li>
	<li><code>1 &lt;= k &lt;= energy.length - 1</code></li>
</ul>

#### 思路

我们要计算的是一个类似后缀和的概念，例如 $k=1$ 时，计算的就是后缀和的最大值，$k=2$ 时，计算的是间隔一个数的后缀和。

枚举终点为 $i$，我们要计算的是下标为

$$
i,i-k,i-2k,\cdots
$$

的元素之和的最小值。

一边遍历，一边累加元素和，取元素和的最大值，即为答案。

$i$ 可以从 $n-k$ 开始枚举到 $n-1$。

``` 
func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	ans := math.MinInt
	for i := n - k; i < n; i++ {
		s := 0
		for j := i; j >= 0; j -= k {
			s += energy[j]
			ans = max(ans, s)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{energy}$ 的长度。每个元素恰好访问一次。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)