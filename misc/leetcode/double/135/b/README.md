#### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;。</p>

<p>你需要对 <code>s</code>&nbsp;执行以下操作 <strong>任意</strong>&nbsp;次：</p>

<ul>
	<li>选择一个下标 <code>i</code>&nbsp;，满足 <code>s[i]</code>&nbsp;左边和右边都&nbsp;<strong>至少</strong>&nbsp;有一个字符与它相同。</li>
	<li>删除 <code>s[i]</code>&nbsp;<strong>左边</strong>&nbsp;离它 <strong>最近</strong>&nbsp;且相同的字符。</li>
	<li>删除 <code>s[i]</code>&nbsp;<strong>右边</strong>&nbsp;离它 <strong>最近</strong>&nbsp;且相同的字符。</li>
</ul>

<p>请你返回执行完所有操作后， <code>s</code>&nbsp;的 <strong>最短</strong>&nbsp;长度。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abaacbcbb"</span></p>

<p><span class="example-io"><b>输出：</b>5</span></p>

<p><strong>解释：</strong><br />
我们执行以下操作：</p>

<ul>
	<li>选择下标 2 ，然后删除下标 0 和 3 处的字符，得到&nbsp;<code>s = "bacbcbb"</code>&nbsp;。</li>
	<li>选择下标 3 ，然后删除下标 0 和 5 处的字符，得到&nbsp;<code>s = "acbcb"</code>&nbsp;。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "aa"</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong><br />
无法对字符串进行任何操作，所以返回初始字符串的长度。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
</ul>

#### 思路

操作次数取决于每种字母的出现次数，与字母的位置无关。

假设某个字母出现了 $c$ 次，那么操作后该字母最少能剩下多少？

根据题意，只有当 $c\ge 3$ 时才能操作，每次操作可以把 $c$ 减少 $2$。

- 如果 $c=3,5,7,\cdots$ 是奇数，那么不断减 $2$，最终 $c=1$。
- 如果 $c=4,6,8,\cdots$ 是偶数，那么不断减 $2$，最终 $c=2$。

这两种情况可以合并，最终剩下

$$
(c-1)\bmod 2 + 1
$$

个字母。注意上式同时兼顾 $c=0,1,2$ 的情况。

累加每种字母最终剩下的 $c$，即为答案。

```
func minimumLength(s string) (ans int) {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	for _, c := range cnt {
		ans += (c-1)%2 + 1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 是字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)