#### 题目

<p>给你一个字符串 <code>word</code> 和一个整数 <code>numFriends</code>。</p>

<p>Alice 正在为她的 <code>numFriends</code> 位朋友组织一个游戏。游戏分为多个回合，在每一回合中：</p>

<ul>
	<li><code>word</code> 被分割成 <code>numFriends</code> 个&nbsp;<strong>非空&nbsp;</strong>字符串，且该分割方式与之前的任意回合所采用的都 <strong>不完全相同&nbsp;</strong>。</li>
	<li>所有分割出的字符串都会被放入一个盒子中。</li>
</ul>

<p>在所有回合结束后，找出盒子中&nbsp;<strong>字典序最大的&nbsp;</strong>字符串。</p>

<p>字符串 <code>a</code> 的字典序&nbsp;<strong>小于&nbsp;</strong>字符串 <code>b</code> 的前提是：在两个字符串上第一处不同的位置上，<code>a</code> 的字母在字母表中的顺序早于 <code>b</code> 中对应的字母。<br />
如果前 <code>min(a.length, b.length)</code> 个字符都相同，那么较短的字符串字典序更小。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">word = "dbca", numFriends = 2</span></p>

<p><strong>输出:</strong> <span class="example-io">"dbc"</span></p>

<p><strong>解释:</strong>&nbsp;</p>

<p>所有可能的分割方式为：</p>

<ul>
	<li><code>"d"</code> 和 <code>"bca"</code>。</li>
	<li><code>"db"</code> 和 <code>"ca"</code>。</li>
	<li><code>"dbc"</code> 和 <code>"a"</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入:</strong> <span class="example-io">word = "gggg", numFriends = 4</span></p>

<p><strong>输出:</strong> <span class="example-io">"g"</span></p>

<p><strong>解释:</strong>&nbsp;</p>

<p>唯一可能的分割方式为：<code>"g"</code>, <code>"g"</code>, <code>"g"</code>, 和 <code>"g"</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 5&nbsp;* 10<sup>3</sup></code></li>
	<li><code>word</code> 仅由小写英文字母组成。</li>
	<li><code>1 &lt;= numFriends &lt;= word.length</code></li>
</ul>

#### 思路

为方便描述，下文把 $\textit{word}$ 简称为 $s$，把 $\textit{numFriends}$ 简称为 $k$。

## 题意

把 $s$ 分割为 $k$ 个非空子串，返回其中字典序最大的子串。

## 思路

如果固定子串的左端点，那么**子串越长，字典序越大**。
单个子串的长度不能超过多少？
由于其余 $k-1$ 个子串必须是非空的，取长度为 $1$，其余子串的长度之和**至少**为 $k-1$。
所以单个子串的长度**至多**为 $n-(k-1)$。
暴力枚举子串的左端点，计算满足长度上限的最大子串。
注意特判 $k=1$ 的情况，此时无法分割，答案就是 $s$。


```
func answerString(s string, k int) (ans string) {
	if k == 1 {
		return s
	}
	n := len(s)
	for i := range n {
		ans = max(ans, s[i:min(i+n-k+1, n)])
	}
	return
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n-k))$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n-k)$ 或 $\mathcal{O}(1)$。Go 的切片不会发生拷贝，只需要 $\mathcal{O}(1)$ 额外空间。

注：利用**字符串哈希+二分长度**或者**后缀数组**，可以快速比较 $s$ 任意两个子串的大小，做到 $\mathcal{O}(n\log (n-k))$ 或者 $\mathcal{O}(n\log n)$。


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