#### 题目

<p>给你一个长度为 <code>n</code>&nbsp;的字符串 <code>s</code>&nbsp;和一个整数&nbsp;<code>k</code>&nbsp;，<code>n</code>&nbsp;是 <code>k</code>&nbsp;的 <strong>倍数</strong>&nbsp;。你的任务是将字符串 <code>s</code>&nbsp;哈希为一个长度为 <code>n / k</code>&nbsp;的新字符串&nbsp;<code>result</code>&nbsp;。</p>

<p>首先，将&nbsp;<code>s</code>&nbsp;分割成&nbsp;<code>n / k</code>&nbsp;个&nbsp;<strong><span data-keyword="substring-nonempty">子字符串</span></strong>&nbsp;，每个子字符串的长度都为&nbsp;<code>k</code>&nbsp;。然后，将&nbsp;<code>result</code>&nbsp;初始化为一个 <strong>空</strong>&nbsp;字符串。</p>

<p>我们依次从前往后处理每一个 <strong>子字符串</strong>&nbsp;：</p>

<ul>
	<li>一个字符的 <strong>哈希值</strong>&nbsp;是它在 <strong>字母表</strong>&nbsp;中的下标（也就是&nbsp;<code>'a' →<!-- notionvc: d3f8e4c2-23cd-41ad-a14b-101dfe4c5aba --> 0</code>&nbsp;，<code>'b' →<!-- notionvc: d3f8e4c2-23cd-41ad-a14b-101dfe4c5aba --> 1</code>&nbsp;，... ，<code>'z' →<!-- notionvc: d3f8e4c2-23cd-41ad-a14b-101dfe4c5aba --> 25</code>）。</li>
	<li>将子字符串中字幕的 <strong>哈希值</strong>&nbsp;求和。</li>
	<li>将和对 26 取余，将结果记为&nbsp;<code>hashedChar</code>&nbsp;。</li>
	<li>找到小写字母表中 <code>hashedChar</code>&nbsp;对应的字符。</li>
	<li>将该字符添加到&nbsp;<code>result</code>&nbsp;的末尾。</li>
</ul>

<p>返回&nbsp;<code>result</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abcd", k = 2</span></p>

<p><span class="example-io"><b>输出：</b>"bf"</span></p>

<p><b>解释：</b></p>

<p>第一个字符串为&nbsp;<code>"ab"</code>&nbsp;，<code>0 + 1 = 1</code>&nbsp;，<code>1 % 26 = 1</code>&nbsp;，<code>result[0] = 'b'</code>&nbsp;。</p>

<p>第二个字符串为： <code>"cd"</code>&nbsp;，<code>2 + 3 = 5</code>&nbsp;，<code>5 % 26 = 5</code>&nbsp;，<code>result[1] = 'f'</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "mxz", k = 3</span></p>

<p><span class="example-io"><b>输出：</b>"i"</span></p>

<p><b>解释：</b></p>

<p>唯一的子字符串为&nbsp;<code>"mxz"</code>&nbsp;，<code>12 + 23 + 25 = 60</code>&nbsp;，<code>60 % 26 = 8</code>&nbsp;，<code>result[0] = 'i'</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= 100</code></li>
	<li><code>k &lt;= s.length &lt;= 1000</code></li>
	<li><code>s.length</code>&nbsp;能被 <code>k</code>&nbsp;整除。</li>
	<li><code>s</code> 只含有小写英文字母。</li>
</ul>

#### 思路

直接写

```
func stringHash(s string, k int) string {
	n := len(s)
	ans := make([]byte, n/k)
	for i := 0; i < n; i += k {
		sum := 0
		for _, b := range s[i : i+k] {
			sum += int(b - 'a')
		}
		ans[i/k] = 'a' + byte(sum%26)
	}
	return string(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 时间复杂度：$\mathcal{O}(1)$。返回值不计入。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)