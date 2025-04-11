### 题目

<p>Alice 正在她的电脑上输入一个字符串。但是她打字技术比较笨拙，她&nbsp;<strong>可能</strong>&nbsp;在一个按键上按太久，导致一个字符被输入&nbsp;<strong>多次</strong>&nbsp;。</p>

<p>尽管 Alice 尽可能集中注意力，她仍然可能会犯错 <strong>至多</strong>&nbsp;一次。</p>

<p>给你一个字符串&nbsp;<code>word</code> ，它表示 <strong>最终</strong>&nbsp;显示在 Alice 显示屏上的结果。</p>

<p>请你返回 Alice 一开始可能想要输入字符串的总方案数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>word = "abbcccc"</span></p>

<p><span class="example-io"><b>输出：</b>5</span></p>

<p><strong>解释：</strong></p>

<p>可能的字符串包括：<code>"abbcccc"</code>&nbsp;，<code>"abbccc"</code>&nbsp;，<code>"abbcc"</code>&nbsp;，<code>"abbc"</code>&nbsp;和&nbsp;<code>"abcccc"</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>word = "abcd"</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong></p>

<p>唯一可能的字符串是&nbsp;<code>"abcd"</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>word = "aaaa"</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 100</code></li>
	<li><code>word</code>&nbsp;只包含小写英文字母。</li>
</ul>

### 思路

根据题意：

- 如果所有相邻字母都互不相同，那么 Alice 不可能犯错，所以方案数只有 $1$ 种。
- 如果有 $1$ 对相邻字母相同，那么 Alice 可以在这里犯错一次，例如 $\texttt{abb}$，一开始想要输入的可能是 $\texttt{abb}$，也可能是 $\texttt{ab}$，其中 $\texttt{b}$ 多打了一次，所以方案数有 $2$ 种。
- 如果有 $2$ 对相邻字母相同，那么一开始想要输入的字符串会再多一种。例如 $\texttt{abbb}$，一开始想要输入的可能是 $\texttt{abbb}$，也可能是 $\texttt{abb}$，也可能是 $\texttt{ab}$，所以方案数有 $3$ 种。
- 依此类推，方案数等于相邻相同字母的个数加一。

```
func possibleStringCount(word string) int {
	ans := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			ans++
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
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