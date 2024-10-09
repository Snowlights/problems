#### 题目

<p><strong>有效单词</strong> 需要满足以下几个条件：</p>

<ul>
	<li><strong>至少 </strong>包含 3 个字符。</li>
	<li>由数字 0-9 和英文大小写字母组成。（不必包含所有这类字符。）</li>
	<li><strong>至少</strong> 包含一个 <strong>元音字母 </strong>。</li>
	<li><strong>至少</strong> 包含一个 <strong>辅音字母 </strong>。</li>
</ul>

<p>给你一个字符串 <code>word</code> 。如果 <code>word</code> 是一个有效单词，则返回 <code>true</code> ，否则返回 <code>false</code> 。</p>

<p><strong>注意：</strong></p>

<ul>
	<li><code>'a'</code>、<code>'e'</code>、<code>'i'</code>、<code>'o'</code>、<code>'u'</code> 及其大写形式都属于<strong> 元音字母 </strong>。</li>
	<li>英文中的 <strong>辅音字母 </strong>是指那些除元音字母之外的字母。</li>
</ul>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "234Adas"</span></p>

<p><strong>输出：</strong><span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<p>这个单词满足所有条件。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "b3"</span></p>

<p><strong>输出：</strong><span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<p>这个单词的长度少于 3 且没有包含元音字母。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "a3$e"</span></p>

<p><strong>输出：</strong><span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<p>这个单词包含了 <code>'$'</code> 字符且没有包含辅音字母。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 20</code></li>
	<li><code>word</code> 由英文大写和小写字母、数字、<code>'@'</code>、<code>'#'</code> 和 <code>'$'</code> 组成。</li>
</ul>


#### 思路

用两个变量 $f_0$ 和 $f_1$ 记录字符串中是否有辅音或元音，必须都为 $\texttt{true}$ 才返回 $\texttt{true}$。

如果字符串长度不足 $3$ 或者包含除了数字或字母以外的字符，返回 $\texttt{false}$。

``` go
func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	var f0, f1 bool
	for _, c := range word {
		if unicode.IsLetter(c) {
			if strings.ContainsRune("aeiou", unicode.ToLower(c)) {
				f1 = true
			} else {
				f0 = true
			}
		} else if !unicode.IsDigit(c) {
			return false
		}
	}
	return f0 && f1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
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


[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)