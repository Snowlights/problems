#### 题目

<p>给你一个字符串 <code>s</code> 和一个整数 <code>k</code>，在 <code>s</code> 的所有子字符串中，请你统计并返回 <strong>至少有一个 </strong>字符 <strong>至少出现</strong> <code>k</code> 次的子字符串总数。</p>

<p><strong>子字符串 </strong>是字符串中的一个连续、<b> 非空</b> 的字符序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "abacb", k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>符合条件的子字符串如下：</p>

<ul>
	<li><code>"aba"</code>（字符 <code>'a'</code> 出现 2 次）。</li>
	<li><code>"abac"</code>（字符 <code>'a'</code> 出现 2 次）。</li>
	<li><code>"abacb"</code>（字符 <code>'a'</code> 出现 2 次）。</li>
	<li><code>"bacb"</code>（字符 <code>'b'</code> 出现 2 次）。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "abcde", k = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">15</span></p>

<p><strong>解释：</strong></p>

<p>所有子字符串都有效，因为每个字符至少出现一次。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 3000</code></li>
	<li><code>1 &lt;= k &lt;= s.length</code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

**前置知识**：[滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

从小到大枚举子串右端点 $\textit{right}$，如果子串符合要求，则右移左端点 $\textit{left}$。
滑动窗口的内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\cdots,\textit{left}-1$ 的所有子串都是合法的，
这一共有 $\textit{left}$ 个，加入答案。

## 答疑

**问**：为什么只需判断 `cnt[c - 'a'] >= k`？

**答**：因为只有 `cnt[c - 'a']++` 导致了 `cnt[c - 'a']` 达到 $k$，所以其余字母的出现次数必然小于 $k$，无需判断。当然，这里的 `>=` 写成 `==` 也是可以的。

```
func numberOfSubstrings(s string, k int) (ans int) {
	cnt := [26]int{}
	left := 0
	for _, c := range s {
		cnt[c-'a']++
		for cnt[c-'a'] >= k {
			cnt[s[left]-'a']--
			left++
		}
		ans += left
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以滑动窗口的时间复杂度为 $\mathcal{O}(n)$。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)