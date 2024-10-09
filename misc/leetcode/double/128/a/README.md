### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;。一个字符串的&nbsp;<strong>分数</strong>&nbsp;定义为相邻字符 <strong>ASCII</strong>&nbsp;码差值绝对值的和。</p>

<p>请你返回 <code>s</code>&nbsp;的 <strong>分数</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "hello"</span></p>

<p><span class="example-io"><b>输出：</b>13</span></p>

<p><strong>解释：</strong></p>

<p><code>s</code>&nbsp;中字符的 <strong>ASCII </strong>码分别为：<code>'h' = 104</code>&nbsp;，<code>'e' = 101</code>&nbsp;，<code>'l' = 108</code>&nbsp;，<code>'o' = 111</code>&nbsp;。所以&nbsp;<code>s</code>&nbsp;的分数为&nbsp;<code>|104 - 101| + |101 - 108| + |108 - 108| + |108 - 111| = 3 + 7 + 0 + 3 = 13</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "zaz"</span></p>

<p><span class="example-io"><b>输出：</b>50</span></p>

<p><strong>解释：</strong></p>

<p><code>s</code>&nbsp;中字符的 <strong>ASCII&nbsp;</strong>码分别为：<code>'z' = 122</code>&nbsp;，<code>'a' = 97</code>&nbsp;。所以&nbsp;<code>s</code>&nbsp;的分数为&nbsp;<code>|122 - 97| + |97 - 122| = 25 + 25 = 50</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
</ul>

### 思路

模拟

``` go  
func scoreOfString(s string) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		ans += abs(int(s[i]) - int(s[i-1]))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
