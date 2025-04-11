#### 题目

<p>给你一个字符串 <code>s</code> 和一个整数 <code>k</code>。</p>

<p>判断是否存在一个长度&nbsp;<strong>恰好&nbsp;</strong>为 <code>k</code> 的子字符串，该子字符串需要满足以下条件：</p>

<ol>
	<li>该子字符串&nbsp;<strong>只包含一个唯一字符</strong>（例如，<code>"aaa"</code> 或 <code>"bbb"</code>）。</li>
	<li>如果该子字符串的&nbsp;<strong>前面&nbsp;</strong>有字符，则该字符必须与子字符串中的字符不同。</li>
	<li>如果该子字符串的&nbsp;<strong>后面&nbsp;</strong>有字符，则该字符也必须与子字符串中的字符不同。</li>
</ol>

<p>如果存在这样的子串，返回 <code>true</code>；否则，返回 <code>false</code>。</p>

<p><strong>子字符串&nbsp;</strong>是字符串中的连续、非空字符序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "aaabaaa", k = 3</span></p>

<p><strong>输出：</strong> <span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<p>子字符串 <code>s[4..6] == "aaa"</code> 满足条件：</p>
<ul>
	<li>长度为 3。</li>
	<li>所有字符相同。</li>
	<li>子串 <code>"aaa"</code> 前的字符是 <code>'b'</code>，与 <code>'a'</code> 不同。</li>
	<li>子串 <code>"aaa"</code> 后没有字符。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "abc", k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<p>不存在长度为 2 、仅由一个唯一字符组成且满足所有条件的子字符串。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

**题意**：找一个连续相同子串 $s[i]$ 到 $s[j]$，要求长度恰好等于 $k$，且 $s[i-1]\ne s[i]$ 且 $s[j]\ne s[j+1]$（如果有）。
遍历 $s$ 的同时，维护连续相同子串长度 $\textit{cnt}$，遇到子串末尾就看看 $\textit{cnt}=k$ 是否成立，成立就立刻返回 $\texttt{true}$。


```
func hasSpecialSubstring(s string, k int) bool {
	cnt := 0
	for i := range s {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt == k {
				return true
			}
			cnt = 0
		}
	}
	return false
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。


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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)