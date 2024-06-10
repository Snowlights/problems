#### 题目

<p>给你两个字符串数组&nbsp;<code>wordsContainer</code> 和&nbsp;<code>wordsQuery</code>&nbsp;。</p>

<p>对于每个&nbsp;<code>wordsQuery[i]</code>&nbsp;，你需要从&nbsp;<code>wordsContainer</code>&nbsp;中找到一个与&nbsp;<code>wordsQuery[i]</code>&nbsp;有&nbsp;<strong>最长公共后缀</strong>&nbsp;的字符串。如果 <code>wordsContainer</code>&nbsp;中有两个或者更多字符串有最长公共后缀，那么答案为长度<strong>&nbsp;最短</strong>&nbsp;的。如果有超过两个字符串有&nbsp;<strong>相同</strong>&nbsp;最短长度，那么答案为它们在&nbsp;<code>wordsContainer</code>&nbsp;中出现&nbsp;<strong>更早</strong>&nbsp;的一个。</p>

<p>请你返回一个整数数组<em>&nbsp;</em><code>ans</code>&nbsp;，其中<em>&nbsp;</em><code>ans[i]</code>是<em>&nbsp;</em><code>wordsContainer</code>中与&nbsp;<code>wordsQuery[i]</code>&nbsp;有&nbsp;<strong>最长公共后缀</strong>&nbsp;字符串的下标。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>wordsContainer = ["abcd","bcd","xbcd"], wordsQuery = ["cd","bcd","xyz"]</span></p>

<p><span class="example-io"><b>输出：</b>[1,1,1]</span></p>

<p><strong>解释：</strong></p>

<p>我们分别来看每一个&nbsp;<code>wordsQuery[i]</code>&nbsp;：</p>

<ul>
	<li>对于&nbsp;<code>wordsQuery[0] = "cd"</code>&nbsp;，<code>wordsContainer</code>&nbsp;中有最长公共后缀&nbsp;<code>"cd"</code>&nbsp;的字符串下标分别为&nbsp;0 ，1 和&nbsp;2 。这些字符串中，答案是下标为 1 的字符串，因为它的长度为 3 ，是最短的字符串。</li>
	<li>对于&nbsp;<code>wordsQuery[1] = "bcd"</code>&nbsp;，<code>wordsContainer</code>&nbsp;中有最长公共后缀&nbsp;<code>"bcd"</code>&nbsp;的字符串下标分别为 0 ，1 和 2 。这些字符串中，答案是下标为 1 的字符串，因为它的长度为 3 ，是最短的字符串。</li>
	<li>对于&nbsp;<code>wordsQuery[2] = "xyz"</code>&nbsp;，<code>wordsContainer</code>&nbsp;中没有字符串跟它有公共后缀，所以最长公共后缀为&nbsp;<code>""</code>&nbsp;，下标为&nbsp;0 ，1 和 2 的字符串都得到这一公共后缀。这些字符串中，&nbsp;答案是下标为 1 的字符串，因为它的长度为 3 ，是最短的字符串。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>wordsContainer = ["abcdefgh","poiuygh","ghghgh"], wordsQuery = ["gh","acbfgh","acbfegh"]</span></p>

<p><span class="example-io"><b>输出：</b>[2,0,2]</span></p>

<p><strong>解释：</strong></p>

<p>我们分别来看每一个&nbsp;<code>wordsQuery[i]</code>&nbsp;：</p>

<ul>
	<li>对于&nbsp;<code>wordsQuery[0] = "gh"</code>&nbsp;，<code>wordsContainer</code>&nbsp;中有最长公共后缀&nbsp;<code>"gh"</code>&nbsp;的字符串下标分别为 0 ，1 和 2 。这些字符串中，答案是下标为 2 的字符串，因为它的长度为 6 ，是最短的字符串。</li>
	<li>对于&nbsp;<code>wordsQuery[1] = "acbfgh"</code>&nbsp;，只有下标为 0 的字符串有最长公共后缀&nbsp;<code>"fgh"</code>&nbsp;。所以尽管下标为 2 的字符串是最短的字符串，但答案是 0 。</li>
	<li>对于&nbsp;<code>wordsQuery[2] = "acbfegh"</code>&nbsp;，<code>wordsContainer</code>&nbsp;中有最长公共后缀&nbsp;<code>"gh"</code>&nbsp;的字符串下标分别为 0 ，1 和 2 。这些字符串中，答案是下标为 2 的字符串，因为它的长度为 6 ，是最短的字符串。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= wordsContainer.length, wordsQuery.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= wordsContainer[i].length &lt;= 5 * 10<sup>3</sup></code></li>
	<li><code>1 &lt;= wordsQuery[i].length &lt;= 5 * 10<sup>3</sup></code></li>
	<li><code>wordsContainer[i]</code>&nbsp;只包含小写英文字母。</li>
	<li><code>wordsQuery[i]</code>&nbsp;只包含小写英文字母。</li>
	<li><code>wordsContainer[i].length</code>&nbsp;的和至多为&nbsp;<code>5 * 10<sup>5</sup></code>&nbsp;。</li>
	<li><code>wordsQuery[i].length</code>&nbsp;的和至多为&nbsp;<code>5 * 10<sup>5</sup></code>&nbsp;。</li>
</ul>

#### 思路

**前置知识**：[208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)。

从左到右遍历 $\textit{wordsContainer}$，设 $s=\textit{wordsContainer}[i]$。

倒着遍历 $s$，插入字典树。插入时，**对于每个经过的节点**，更新节点对应的最小字符串长度及其下标。

对于查询 $s=\textit{wordsQuery}[i]$，仍然倒着遍历 $s$。在字典树上**找到最后一个匹配的节点**，那么该节点保存的下标就是答案。

```go [sol-Go]
func stringIndices(wordsContainer, wordsQuery []string) []int {
	type node struct {
		son     [26]*node
		minL, i int
	}
	root := &node{minL: math.MaxInt}

	for idx, s := range wordsContainer {
		l := len(s)
		cur := root
		if l < cur.minL {
			cur.minL, cur.i = l, idx
		}
		for i := len(s) - 1; i >= 0; i-- {
			b := s[i] - 'a'
			if cur.son[b] == nil {
				cur.son[b] = &node{minL: math.MaxInt}
			}
			cur = cur.son[b]
			if l < cur.minL {
				cur.minL, cur.i = l, idx
			}
		}
	}

	ans := make([]int, len(wordsQuery))
	for idx, s := range wordsQuery {
		cur := root
		for i := len(s) - 1; i >= 0 && cur.son[s[i]-'a'] != nil; i-- {
			cur = cur.son[s[i]-'a']
		}
		ans[idx] = cur.i
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(L_1|\Sigma| + L_2)$，其中 $L_1$ 为 $\textit{wordsContainer}$ 中的所有字符串的长度**之和**，$L_2$ 为 $\textit{wordsQuery}$ 中的所有字符串的长度**之和**，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(L_1|\Sigma|)$。返回值不计入。

#### 相似题目

- [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/)
- [2416. 字符串的前缀分数和](https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/) 1725
- [336. 回文对](https://leetcode.cn/problems/palindrome-pairs/)
- [745. 前缀和后缀搜索](https://leetcode.cn/problems/prefix-and-suffix-search/)
- [3045. 统计前后缀下标对 II](https://leetcode.cn/problems/count-prefix-and-suffix-pairs-ii/) 2328
- [527. 单词缩写](https://leetcode.cn/problems/word-abbreviation/)（会员题）
- [1804. 实现 Trie （前缀树） II](https://leetcode.cn/problems/implement-trie-ii-prefix-tree/)（会员题）

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
