### 题目

<p>电子表格是一个网格，它有 26 列（从 <code>'A'</code> 到 <code>'Z'</code>）和指定数量的 <code>rows</code>。每个单元格可以存储一个 0 到 10<sup>5</sup>&nbsp;之间的整数值。</p>

<p>请你实现一个&nbsp;<code>Spreadsheet</code> 类：</p>

<ul>
	<li><code>Spreadsheet(int rows)</code> 初始化一个具有 26 列（从 <code>'A'</code> 到 <code>'Z'</code>）和指定行数的电子表格。所有单元格最初的值都为 0 。</li>
	<li><code>void setCell(String cell, int value)</code> 设置指定单元格的值。单元格引用以 <code>"AX"</code> 的格式提供（例如，<code>"A1"</code>，<code>"B10"</code>），其中字母表示列（从 <code>'A'</code> 到 <code>'Z'</code>），数字表示从<strong>&nbsp;</strong><strong>1</strong>&nbsp;开始的行号。</li>
	<li><code>void resetCell(String cell)</code> 重置指定单元格的值为 0 。</li>
	<li><code>int getValue(String formula)</code> 计算一个公式的值，格式为 <code>"=X+Y"</code>，其中 <code>X</code> 和 <code>Y</code>&nbsp;<strong>要么</strong> 是单元格引用，要么非负整数，返回计算的和。</li>
</ul>

<p><strong>注意：</strong> 如果 <code>getValue</code> 引用一个未通过 <code>setCell</code> 明确设置的单元格，则该单元格的值默认为 0 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><br />
<span class="example-io">["Spreadsheet", "getValue", "setCell", "getValue", "setCell", "getValue", "resetCell", "getValue"]<br />
[[3], ["=5+7"], ["A1", 10], ["=A1+6"], ["B2", 15], ["=A1+B2"], ["A1"], ["=A1+B2"]]</span></p>

<p><strong>输出：</strong><br />
<span class="example-io">[null, 12, null, 16, null, 25, null, 15] </span></p>

<p><strong>解释</strong></p>
Spreadsheet spreadsheet = new Spreadsheet(3); // 初始化一个具有 3 行和 26 列的电子表格<br data-end="321" data-start="318" />
spreadsheet.getValue("=5+7"); // 返回 12 (5+7)<br data-end="373" data-start="370" />
spreadsheet.setCell("A1", 10); // 设置 A1 为 10<br data-end="423" data-start="420" />
spreadsheet.getValue("=A1+6"); // 返回 16 (10+6)<br data-end="477" data-start="474" />
spreadsheet.setCell("B2", 15); // 设置 B2 为 15<br data-end="527" data-start="524" />
spreadsheet.getValue("=A1+B2"); // 返回 25 (10+15)<br data-end="583" data-start="580" />
spreadsheet.resetCell("A1"); // 重置 A1 为 0<br data-end="634" data-start="631" />
spreadsheet.getValue("=A1+B2"); // 返回 15 (0+15)</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= rows &lt;= 10<sup>3</sup></code></li>
	<li><code>0 &lt;= value &lt;= 10<sup>5</sup></code></li>
	<li>公式保证采用 <code>"=X+Y"</code> 格式，其中 <code>X</code> 和 <code>Y</code> 要么是有效的单元格引用，要么是小于等于 <code>10<sup>5</sup></code> 的 <strong>非负</strong> 整数。</li>
	<li>每个单元格引用由一个大写字母 <code>'A'</code> 到 <code>'Z'</code> 和一个介于 <code>1</code> 和 <code>rows</code> 之间的行号组成。</li>
	<li><strong>总共</strong> 最多会对 <code>setCell</code>、<code>resetCell</code> 和 <code>getValue</code> 调用 <code>10<sup>4</sup></code> 次。</li>
</ul>

### 思路

创建一个 key 为字符串，value 为整数的哈希表。

- $\texttt{setCell}$：把 $\textit{cell}$ 和 $\textit{value}$ 插入哈希表。注意不需要解析 $\textit{cell}$。
- $\texttt{resetCell}$：把 $\textit{cell}$ 从哈希表中删除。
- $\texttt{getValue}$：去掉第一个字符，然后用 $\texttt{+}$ 号分割字符串，查找哈希表，把两部分的和加入答案。

```
type Spreadsheet map[string]int

func Constructor(int) Spreadsheet {
	return Spreadsheet{}
}

func (s Spreadsheet) SetCell(cell string, value int) {
	s[cell] = value
}

func (s Spreadsheet) ResetCell(cell string) {
	delete(s, cell)
}

func (s Spreadsheet) GetValue(formula string) (ans int) {
	for _, cell := range strings.Split(formula[1:], "+") {
		if unicode.IsUpper(rune(cell[0])) {
			ans += s[cell]
		} else {
			x, _ := strconv.Atoi(cell)
			ans += x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：初始化为 $\mathcal{O}(1)$，其余为 $\mathcal{O}(L)$，其中 $L$ 是 $\textit{cell}$（或者 $\textit{formula}$）的长度。
- 空间复杂度：$\mathcal{O}(qL)$。其中 $q$ 为 $\texttt{setCell}$ 的调用次数。

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