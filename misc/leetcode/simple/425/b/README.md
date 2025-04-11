#### 题目

<p>给你两个字符串 <code>s</code> 和 <code>t</code>（它们互为字母异位词），以及一个整数 <code>k</code>。</p>

<p>你的任务是判断是否可以将字符串 <code>s</code> 分割成 <code>k</code> 个等长的子字符串，然后重新排列这些子字符串，并以任意顺序连接它们，使得最终得到的新字符串与给定的字符串 <code>t</code> 相匹配。</p>

<p>如果可以做到，返回 <code>true</code>；否则，返回 <code>false</code>。</p>

<p><strong>字母异位词&nbsp;</strong>是指由另一个单词或短语的所有字母重新排列形成的单词或短语，使用所有原始字母恰好一次。</p>

<p><strong>子字符串&nbsp;</strong>是字符串中的一个连续&nbsp;<b>非空&nbsp;</b>字符序列。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "abcd", t = "cdab", k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>将 <code>s</code> 分割成 2 个长度为 2 的子字符串：<code>["ab", "cd"]</code>。</li>
	<li>重新排列这些子字符串为 <code>["cd", "ab"]</code>，然后连接它们得到 <code>"cdab"</code>，与 <code>t</code> 相匹配。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "aabbcc", t = "bbaacc", k = 3</span></p>

<p><strong>输出：</strong> <span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>将 <code>s</code> 分割成 3 个长度为 2 的子字符串：<code>["aa", "bb", "cc"]</code>。</li>
	<li>重新排列这些子字符串为 <code>["bb", "aa", "cc"]</code>，然后连接它们得到 <code>"bbaacc"</code>，与 <code>t</code> 相匹配。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "aabbcc", t = "bbaacc", k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>将 <code>s</code> 分割成 2 个长度为 3 的子字符串：<code>["aab", "bcc"]</code>。</li>
	<li>这些子字符串无法重新排列形成 <code>t = "bbaacc"</code>，所以输出 <code>false</code>。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length == t.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>1 &lt;= k &lt;= s.length</code></li>
	<li><code>s.length</code> 能被 <code>k</code> 整除。</li>
	<li><code>s</code> 和 <code>t</code> 仅由小写英文字母组成。</li>
	<li>输入保证 <code>s</code> 和 <code>t</code> 互为字母异位词。</li>
</ul>

#### 思路

最终只需要关心 $s$ 是否等于 $t$，那么交换 $s$ 中的两个子串，等同于交换 $t$ 中的两个子串。

所以可以视作 $s$ 和 $t$ 中的子串可以随意重排。

那么只需判断 $s$ 子串和 $t$ 子串排序后的结果是否一样。

也可以用哈希表统计每个子串的出现次数。

```
func isPossibleToRearrange(s, t string, k int) bool {
	a := make([]string, 0, k) // 预分配空间
	b := make([]string, 0, k)
	n := len(s)
	k = n / k
	for i := k; i <= n; i += k {
		a = append(a, s[i-k:i])
		b = append(b, t[i-k:i])
	}
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n\log k)$，其中 $n$ 是 $s$ 的长度。
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
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)