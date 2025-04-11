### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;和一个模式字符串&nbsp;<code>p</code>&nbsp;，其中&nbsp;<code>p</code> <strong>恰好</strong>&nbsp;包含 <strong>一个</strong>&nbsp;<code>'*'</code>&nbsp;符号。</p>

<p><code>p</code>&nbsp;中的 <code>'*'</code>&nbsp;符号可以被替换为零个或多个字符组成的任意字符序列。</p>

<p>如果 <code>p</code>&nbsp;可以变成 <code>s</code>&nbsp;的子字符串，那么返回&nbsp;<code>true</code>&nbsp;，否则返回 <code>false</code>&nbsp;。</p>

<p><strong>子字符串</strong>&nbsp;指的是字符串中一段连续 <strong>非空</strong>&nbsp;的字符序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "leetcode", p = "ee*e"</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><b>解释：</b></p>

<p>将&nbsp;<code>'*'</code>&nbsp;替换为&nbsp;<code>"tcod"</code>&nbsp;，子字符串&nbsp;<code>"eetcode"</code>&nbsp;匹配模式串。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "car", p = "c*v"</span></p>

<p><span class="example-io"><b>输出：</b>false</span></p>

<p><strong>解释：</strong></p>

<p>不存在匹配模式串的子字符串。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "luck", p = "u*"</span></p>

<p><span class="example-io"><b>输出：</b>true</span></p>

<p><b>解释：</b></p>

<p>子字符串&nbsp;<code>"u"</code>&nbsp;，<code>"uc"</code>&nbsp;和&nbsp;<code>"uck"</code>&nbsp;都匹配模式串。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 50</code></li>
	<li><code>1 &lt;= p.length &lt;= 50 </code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
	<li><code>p</code>&nbsp;只包含小写英文字母和一个&nbsp;<code>'*'</code> 符号。</li>
</ul>

### 思路

算法：

1. 找到 $p$ 中 $\texttt{*}$ 的位置，记作 $\textit{star}$。
2. 把 $p$ 分成两部分，$\texttt{*}$ 左边的记作 $p[:\textit{star}]$，右边的记作 $p[\textit{star}+1:]$。
3. 在 $s$ 中找 $p[:\textit{star}]$ 首次出现的位置 $i$。如果没找到，返回 $\texttt{false}$。
4. 继续匹配 $\texttt{*}$ 右边的内容，也就是判断 $p[\textit{star}+1:]$ 是否在 $s[i+\textit{star}:]$ 中。


```
func hasMatch(s, p string) bool {
	star := strings.IndexByte(p, '*')
	i := strings.Index(s, p[:star])
	return i >= 0 && strings.Contains(s[i+star:], p[star+1:])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nm)$，其中 $n$ 是 $s$ 的长度，$m$ 是 $p$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$ 或 $\mathcal{O}(1)$。Go 的切片没有拷贝，是 $\mathcal{O}(1)$ 的。


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