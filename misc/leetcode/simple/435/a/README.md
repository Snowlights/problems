#### 题目

<p>给你一个由小写英文字母组成的字符串&nbsp;<code>s</code> 。请你找出字符串中两个字符的出现频次之间的 <strong>最大</strong> 差值，这两个字符需要满足：</p>

<ul>
	<li>一个字符在字符串中出现 <strong>偶数次</strong> 。</li>
	<li>另一个字符在字符串中出现 <strong>奇数次</strong>&nbsp;。</li>
</ul>

<p>返回 <strong>最大</strong> 差值，计算方法是出现 <strong>奇数次</strong> 字符的次数 <strong>减去</strong> 出现 <strong>偶数次</strong> 字符的次数。</p>

<p>&nbsp;</p>

<p><b>示例 1：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "aaaaabbc"</span></p>

<p><b>输出：</b>3</p>

<p><b>解释：</b></p>

<ul>
	<li>字符&nbsp;<code>'a'</code>&nbsp;出现 <strong>奇数次</strong>&nbsp;，次数为&nbsp;<code><font face="monospace">5</font></code><font face="monospace"> ；字符</font>&nbsp;<code>'b'</code>&nbsp;出现 <strong>偶数次</strong> ，次数为&nbsp;<code><font face="monospace">2</font></code>&nbsp;。</li>
	<li>最大差值为&nbsp;<code>5 - 2 = 3</code>&nbsp;。</li>
</ul>
</div>

<p><b>示例 2：</b></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abcabcab"</span></p>

<p><b>输出：</b>1</p>

<p><b>解释：</b></p>

<ul>
	<li>字符&nbsp;<code>'a'</code>&nbsp;出现 <strong>奇数次</strong>&nbsp;，次数为&nbsp;<code><font face="monospace">3</font></code><font face="monospace"> ；字符</font>&nbsp;<code>'c'</code>&nbsp;出现 <strong>偶数次</strong>&nbsp;，次数为&nbsp;<font face="monospace">2 。</font></li>
	<li>最大差值为&nbsp;<code>3 - 2 = 1</code> 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><b>提示：</b></p>

<ul>
	<li><code>3 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code>&nbsp;仅由小写英文字母组成。</li>
	<li><code>s</code>&nbsp;至少由一个出现奇数次的字符和一个出现偶数次的字符组成。</li>
</ul>


#### 思路

统计每种字母的出现次数 $\textit{cnt}$。

要让奇数减偶数尽量大，那么奇数要尽量大，偶数要尽量小。

遍历所有出现次数 $c = \textit{cnt}[i]$，计算：

- 最大的奇数 $c$，记作 $\textit{max}_1$。
- 最小的正偶数 $c$，记作 $\textit{min}_0$。

答案为

$$
\textit{max}_1 - \textit{min}_0
$$


```
func maxDifference(s string) int {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}

	max1, min0 := 0, math.MaxInt
	for _, c := range cnt {
		if c%2 > 0 {
			max1 = max(max1, c)
		} else if c > 0 {
			min0 = min(min0, c)
		}
	}
	return max1 - min0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $s$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。


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