### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;。</p>

<p>你的任务是重复以下操作删除 <strong>所有</strong>&nbsp;数字字符：</p>

<ul>
	<li>删除 <strong>第一个数字字符</strong>&nbsp;以及它左边 <strong>最近</strong>&nbsp;的 <strong>非数字</strong>&nbsp;字符。</li>
</ul>

<p>请你返回删除所有数字字符以后剩下的字符串。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abc"</span></p>

<p><span class="example-io"><b>输出：</b>"abc"</span></p>

<p><strong>解释：</strong></p>

<p>字符串中没有数字。<!-- notionvc: ff07e34f-b1d6-41fb-9f83-5d0ba3c1ecde --></p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "cb34"</span></p>

<p><span class="example-io"><b>输出：</b>""</span></p>

<p><b>解释：</b></p>

<p>一开始，我们对&nbsp;<code>s[2]</code>&nbsp;执行操作，<code>s</code>&nbsp;变为&nbsp;<code>"c4"</code>&nbsp;。</p>

<p>然后对&nbsp;<code>s[1]</code>&nbsp;执行操作，<code>s</code>&nbsp;变为&nbsp;<code>""</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母和数字字符。</li>
	<li>输入保证所有数字都可以按以上操作被删除。</li>
</ul>

### 思路

用栈维护，遇到数字则弹出栈顶，否则把字符入栈。最后从栈底到栈顶就是答案。

``` go  
func clearDigits(s string) string {
	st := []rune{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}

```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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