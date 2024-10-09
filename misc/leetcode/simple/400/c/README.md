#### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;。它可能包含任意数量的&nbsp;<code>'*'</code>&nbsp;字符。你的任务是删除所有的&nbsp;<code>'*'</code>&nbsp;字符。</p>

<p>当字符串还存在至少一个&nbsp;<code>'*'</code>&nbsp;字符时，你可以执行以下操作：</p>

<ul>
	<li>删除最左边的&nbsp;<code>'*'</code>&nbsp;字符，同时删除该星号字符左边一个字典序 <strong>最小</strong>&nbsp;的字符。如果有多个字典序最小的字符，你可以删除它们中的任意一个。</li>
</ul>

<p>请你返回删除所有&nbsp;<code>'*'</code>&nbsp;字符以后，剩余字符连接而成的 <span data-keyword="lexicographically-smaller-string">字典序最小</span> 的字符串。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "aaba*"</span></p>

<p><span class="example-io"><b>输出：</b>"aab"</span></p>

<p><strong>解释：</strong></p>

<p>删除 <code>'*'</code>&nbsp;号和它左边的其中一个&nbsp;<code>'a'</code>&nbsp;字符。如果我们选择删除&nbsp;<code>s[3]</code>&nbsp;，<code>s</code>&nbsp;字典序最小。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abc"</span></p>

<p><span class="example-io"><b>输出：</b>"abc"</span></p>

<p><strong>解释：</strong></p>

<p>字符串中没有&nbsp;<code>'*'</code>&nbsp;字符。<!-- notionvc: ff07e34f-b1d6-41fb-9f83-5d0ba3c1ecde --></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code>&nbsp;只含有小写英文字母和&nbsp;<code>'*'</code>&nbsp;字符。</li>
	<li>输入保证操作可以删除所有的&nbsp;<code>'*'</code>&nbsp;字符。</li>
</ul>

#### 思路

**核心思路**：由于要去掉最小的字母，为了让字典序尽量小，相比去掉前面的字母，去掉后面的字母更好。
从左到右遍历 $s$，用 $26$ 个栈模拟。
第 $i$ 个栈维护第 $i$ 个小写字母的下标。
遇到 `*` 时，弹出第一个非空栈的栈顶下标。
最后把所有栈顶下标对应的字母组合起来，即为答案。

``` 
func clearStars(s string) string {
	st := [26][]int{}
	for i, c := range s {
		if c != '*' {
			st[c-'a'] = append(st[c-'a'], i)
			continue
		}
		for j, p := range st {
			if len(p) > 0 {
				st[j] = p[:len(p)-1]
				break
			}
		}
	}

	idx := []int{}
	for _, p := range st {
		idx = append(idx, p...)
	}
	slices.Sort(idx)

	t := make([]byte, len(idx))
	for i, j := range idx {
		t[i] = s[j]
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(n+|\Sigma|)$。\n\n注：用最小堆或者有序集合，可以把 $\mathcal{O}(n|\Sigma|)$ 优化至 $\mathcal{O}(n\log |\Sigma|)$。


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)