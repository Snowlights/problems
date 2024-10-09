#### 题目

<p>给你一个字符串&nbsp;<code>s</code>&nbsp;，它由某个字符串&nbsp;<code>t</code>&nbsp;和若干&nbsp;<code>t</code>&nbsp; 的&nbsp;<strong>同位字符串</strong>&nbsp;连接而成。</p>

<p>请你返回字符串 <code>t</code>&nbsp;的 <strong>最小</strong>&nbsp;可能长度。</p>

<p><strong>同位字符串</strong>&nbsp;指的是重新排列一个单词得到的另外一个字符串，原来字符串中的每个字符在新字符串中都恰好只使用一次。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "abba"</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>一个可能的字符串&nbsp;<code>t</code>&nbsp;为&nbsp;<code>"ba"</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>s = "cdef"</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p>一个可能的字符串&nbsp;<code>t</code>&nbsp;为&nbsp;<code>"cdef"</code>&nbsp;，注意&nbsp;<code>t</code>&nbsp;可能等于&nbsp;<code>s</code>&nbsp;。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code>&nbsp;只包含小写英文字母。</li>
</ul>

#### 思路

设 $s$ 的长度为 $n$，$t$ 的长度为 $k$。

由于 $s$ 是由若干长度为 $k$ 的字符串拼接而成，所以 $k$ 一定是 $n$ 的因子。

由于 $10^5$ 以内的因子个数至多为 $128$（$83160$ 的因子个数），所以我们可以暴力枚举 $n$ 的因子作为 $k$。

然后比较所有首字母下标为 $0,k,2k,3k,\cdots,n-k$ 的长为 $k$ 的子串，所包含的字母及其个数是否一样（同位字符串）。

注意只需枚举小于 $n$ 的因子，如果这些因子都不满足要求，答案一定是 $n$（如示例 2）。

``` go
func minAnagramLength(s string) int {
	n := len(s)
next:
	for k := 1; k <= n/2; k++ {
		if n%k > 0 {
			continue
		}
		cnt0 := [26]int{}
		for _, b := range s[:k] {
			cnt0[b-'a']++
		}
		for i := k * 2; i <= len(s); i += k {
			cnt := [26]int{}
			for _, b := range s[i-k : i] {
				cnt[b-'a']++
			}
			if cnt != cnt0 {
				continue next
			}
		}
		return k
	}
	return n
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(A\cdot n)$，其中 $n$ 为 $s$ 的长度，$A$ 为 $n$ 的因子个数。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。其中 $|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。

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