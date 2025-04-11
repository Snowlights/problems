### 题目

<p>给你一个只包含数字的字符串&nbsp;<code>s</code>&nbsp;。如果 <code>s</code>&nbsp;中两个 <strong>相邻</strong>&nbsp;的数字满足以下条件，我们称它们是 <strong>合法的</strong>&nbsp;：</p>

<ul>
	<li>前面的数字 <strong>不等于</strong> 第二个数字。</li>
	<li>两个数字在 <code>s</code>&nbsp;中出现的次数 <strong>恰好</strong>&nbsp;分别等于这个数字本身。</li>
</ul>

<p>请你从左到右遍历字符串 <code>s</code>&nbsp;，并返回最先找到的 <strong>合法</strong>&nbsp;相邻数字。如果这样的相邻数字不存在，请你返回一个空字符串。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "2523533"</span></p>

<p><span class="example-io"><b>输出：</b>"23"</span></p>

<p><strong>解释：</strong></p>

<p>数字&nbsp;<code>'2'</code>&nbsp;出现 2 次，数字&nbsp;<code>'3'</code>&nbsp;出现 3 次。<code>"23"</code>&nbsp;中每个数字在 <code>s</code>&nbsp;中出现的次数都恰好分别等于数字本身。所以输出&nbsp;<code>"23"</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "221"</span></p>

<p><span class="example-io"><b>输出：</b>"21"</span></p>

<p><strong>解释：</strong></p>

<p>数字&nbsp;<code>'2'</code>&nbsp;出现 2 次，数字&nbsp;<code>'1'</code>&nbsp;出现 1 次。所以输出&nbsp;<code>"21"</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "22"</span></p>

<p><span class="example-io"><b>输出：</b>""</span></p>

<p><strong>解释：</strong></p>

<p>没有合法的相邻数字。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code>&nbsp;只包含&nbsp;<code>'1'</code> 到&nbsp;<code>'9'</code> 的数字。</li>
</ul>

### 思路

统计 $s$ 中每个数字的出现次数。

遍历 $s$，如果相邻数字满足要求，返回这两个数字组成的字符串。

如果没有满足要求的相邻数字，返回空字符串。

```
func findValidPair(s string) (ans string) {
	cnt := [10]int{}
	for _, b := range s {
		cnt[b-'0']++
	}
	for i := 1; i < len(s); i++ {
		x, y := s[i-1]-'0', s[i]-'0'
		if x != y && cnt[x] == int(x) && cnt[y] == int(y) {
			return s[i-1 : i+1]
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=10$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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