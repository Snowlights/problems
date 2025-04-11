#### 题目

<p>给你一个字符串 <code>initialCurrency</code>，表示初始货币类型，并且你一开始拥有 <code>1.0</code> 单位的 <code>initialCurrency</code>。</p>

<p>另给你四个数组，分别表示货币对（字符串）和汇率（实数）：</p>

<ul>
	<li><code>pairs1[i] = [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code> 表示在&nbsp;<strong>第 1 天</strong>，可以按照汇率 <code>rates1[i]</code> 将 <code>startCurrency<sub>i</sub></code> 转换为 <code>targetCurrency<sub>i</sub></code>。</li>
	<li><code>pairs2[i] = [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code> 表示在&nbsp;<strong>第 2 天</strong>，可以按照汇率 <code>rates2[i]</code> 将 <code>startCurrency<sub>i</sub></code> 转换为 <code>targetCurrency<sub>i</sub></code>。</li>
	<li>此外，每种 <code>targetCurrency</code> 都可以以汇率 <code>1 / rate</code> 转换回对应的 <code>startCurrency</code>。</li>
</ul>

<p>你可以在&nbsp;<strong>第 1 天&nbsp;</strong>使用 <code>rates1</code> 进行任意次数的兑换（包括 0 次），然后在&nbsp;<strong>第 2 天&nbsp;</strong>使用 <code>rates2</code> 再进行任意次数的兑换（包括 0 次）。</p>

<p>返回在两天兑换后，最大可能拥有的 <code>initialCurrency</code> 的数量。</p>

<p><strong>注意：</strong>汇率是有效的，并且第 1 天和第 2 天的汇率之间相互独立，不会产生矛盾。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">initialCurrency = "EUR", pairs1 = [["EUR","USD"],["USD","JPY"]], rates1 = [2.0,3.0], pairs2 = [["JPY","USD"],["USD","CHF"],["CHF","EUR"]], rates2 = [4.0,5.0,6.0]</span></p>

<p><strong>输出：</strong> <span class="example-io">720.00000</span></p>

<p><strong>解释：</strong></p>

<p>根据题目要求，需要最大化最终的 <strong>EUR</strong> 数量，从 1.0 <strong>EUR</strong> 开始：</p>

<ul>
	<li><strong>第 1 天：</strong>
	<ul>
		<li>将 <strong>EUR</strong> 换成 <strong>USD</strong>，得到 2.0&nbsp;<strong>USD</strong>。</li>
		<li>将 <strong>USD</strong> 换成 <strong>JPY</strong>，得到 6.0 <strong>JPY</strong>。</li>
	</ul>
	</li>
	<li><strong>第 2 天：</strong>
	<ul>
		<li>将 <strong>JPY</strong> 换成 <strong>USD</strong>，得到 24.0 <strong>USD</strong>。</li>
		<li>将 <strong>USD</strong> 换成 <strong>CHF</strong>，得到 120.0 <strong>CHF</strong>。</li>
		<li>最后将 <strong>CHF</strong> 换回 <strong>EUR</strong>，得到 720.0 <strong>EUR</strong>。</li>
	</ul>
	</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">initialCurrency = "NGN", pairs1 = [["NGN","EUR"]], rates1 = [9.0], pairs2 = [["NGN","EUR"]], rates2 = [6.0]</span></p>

<p><strong>输出：</strong> <span class="example-io">1.50000</span></p>

<p><strong>解释：</strong></p>

<p>在第 1 天将 <strong>NGN</strong> 换成 <strong>EUR</strong>，并在第 2 天用反向汇率将 <strong>EUR</strong> 换回 <strong>NGN</strong>，可以最大化最终的 <strong>NGN</strong> 数量。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">initialCurrency = "USD", pairs1 = [["USD","EUR"]], rates1 = [1.0], pairs2 = [["EUR","JPY"]], rates2 = [10.0]</span></p>

<p><strong>输出：</strong> <span class="example-io">1.00000</span></p>

<p><strong>解释：</strong></p>

<p>在这个例子中，不需要在任何一天进行任何兑换。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= initialCurrency.length &lt;= 3</code></li>
	<li><code>initialCurrency</code> 仅由大写英文字母组成。</li>
	<li><code>1 &lt;= n == pairs1.length &lt;= 10</code></li>
	<li><code>1 &lt;= m == pairs2.length &lt;= 10</code></li>
	<li><code>pairs1[i] == [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code></li>
	<li><code>pairs2[i] == [startCurrency<sub>i</sub>, targetCurrency<sub>i</sub>]</code></li>
	<li><code>1 &lt;= startCurrency<sub>i</sub>.length, targetCurrency<sub>i</sub>.length &lt;= 3</code></li>
	<li><code>startCurrency<sub>i</sub></code> 和 <code>targetCurrency<sub>i</sub></code> 仅由大写英文字母组成。</li>
	<li><code>rates1.length == n</code></li>
	<li><code>rates2.length == m</code></li>
	<li><code>1.0 &lt;= rates1[i], rates2[i] &lt;= 10.0</code></li>
	<li>输入保证两个转换图在各自的天数中没有矛盾或循环。</li>
</ul>

#### 思路

1. 根据 $\textit{pairs}_1$ 和 $\textit{rates}_1$ 建图。
2. 从 $\textit{initialCurrency}$ 开始，自顶向下 DFS 这张图，递归的同时维护金额。记录把 $\textit{initialCurrency}$ 兑换成其他货币的金额 $\textit{day1Amount}$。
3. 根据 $\textit{pairs}_2$ 和 $\textit{rates}_2$ 建图。
4. 同样地，从 $\textit{initialCurrency}$ 开始，自顶向下 DFS 这张图，递归的同时维护金额。记录把 $\textit{initialCurrency}$ 兑换成其他货币的金额 $\textit{day2Amount}$。金额的倒数，就是从其他货币兑换成 $\textit{initialCurrency}$ 的金额

枚举中转货币 $x$，答案为 $\dfrac{\textit{day1Amount}[x]}{\textit{day2Amount}[x]}$ 的最大值。


```
type pair struct {
	to   string
	rate float64
}

func calcAmount(pairs [][]string, rates []float64, initialCurrency string) map[string]float64 {
	g := map[string][]pair{}
	for i, p := range pairs {
		x, y, r := p[0], p[1], rates[i]
		g[x] = append(g[x], pair{y, r})
		g[y] = append(g[y], pair{x, 1 / r})
	}

	amount := map[string]float64{}
	var dfs func(string, float64)
	dfs = func(x string, curAmount float64) {
		amount[x] = curAmount
		for _, e := range g[x] {
			// 每个节点只需递归一次（重复递归算出来的结果是一样的，因为题目保证汇率没有矛盾）
			if amount[e.to] == 0 {
				dfs(e.to, curAmount*e.rate)
			}
		}
	}
	dfs(initialCurrency, 1)
	return amount
}

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) (ans float64) {
	day1Amount := calcAmount(pairs1, rates1, initialCurrency)
	day2Amount := calcAmount(pairs2, rates2, initialCurrency)
	for x, a2 := range day2Amount {
		ans = max(ans, day1Amount[x]/a2)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)L)$，其中 $n$ 是 $\textit{pairs}_1$ 的长度，$m$ 是 $\textit{pairs}_2$ 的长度，$L$ 是单个字符串的长度（不超过 $3$）。
- 空间复杂度：$\mathcal{O}((n+m)L)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)