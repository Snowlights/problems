#### 题目

<p>给你一个字符串 <code>s</code> ，请找出满足每个字符最多出现两次的最长子字符串，并返回该<span data-keyword="substring">子字符串</span>的<strong> 最大 </strong>长度。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "bcbbbcba"</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>以下子字符串长度为 4，并且每个字符最多出现两次：<code>"bcbb<u>bcba</u>"</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "aaaa"</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>以下子字符串长度为 2，并且每个字符最多出现两次：<code>"<u>aa</u>aa"</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul><!-- 字符串 s 的长度在 2 到 100 之间 -->
	<li><code>2 &lt;= s.length &lt;= 100</code></li>
	<!-- 字符串 s 仅包含小写英文字母 -->
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

本题相当于把 [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) 中的「出现次数不超过 $1$」改成「出现次数不超过 $2$」，

```go [sol-Go]
func maximumLengthSubstring(s string) (ans int) {
	cnt := [26]int{}
	left := 0
	for i, b := range s {
		b -= 'a'
		cnt[b]++
		for cnt[b] > 2 {
			cnt[s[left]-'a']--
			left++
		}
		ans = max(ans, i-left+1)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 为 $s$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

