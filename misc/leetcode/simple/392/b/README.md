#### 题目

<p>给你一个字符串 <code>s</code> 和一个整数 <code>k</code> 。</p>

<p>定义函数 <code>distance(s<sub>1</sub>, s<sub>2</sub>)</code> ，用于衡量两个长度为 <code>n</code> 的字符串 <code>s<sub>1</sub></code> 和 <code>s<sub>2</sub></code> 之间的距离，即：</p>

<ul>
	<li>字符 <code>'a'</code> 到 <code>'z'</code> 按 <strong>循环 </strong>顺序排列，对于区间 <code>[0, n - 1]</code> 中的 <code>i</code> ，计算所有「 <code>s<sub>1</sub>[i]</code> 和 <code>s<sub>2</sub>[i]</code> 之间<strong> 最小距离</strong>」的 <strong>和 </strong>。</li>
</ul>

<p>例如，<code>distance("ab", "cd") == 4</code> ，且 <code>distance("a", "z") == 1</code> 。</p>

<p>你可以对字符串 <code>s</code> 执行<strong> 任意次 </strong>操作。在每次操作中，可以将 <code>s</code> 中的一个字母 <strong>改变 </strong>为<strong> 任意 </strong>其他小写英文字母。</p>

<p>返回一个字符串，表示在执行一些操作后你可以得到的 <strong>字典序最小</strong> 的字符串 <code>t</code> ，且满足 <code>distance(s, t) &lt;= k</code> 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "zbbz", k = 3
<strong>输出：</strong>"aaaz"
<strong>解释：</strong>在这个例子中，可以执行以下操作：
将 s[0] 改为 'a' ，s 变为 "abbz" 。
将 s[1] 改为 'a' ，s 变为 "aabz" 。
将 s[2] 改为 'a' ，s 变为 "aaaz" 。
"zbbz" 和 "aaaz" 之间的距离等于 k = 3 。
可以证明 "aaaz" 是在任意次操作后能够得到的字典序最小的字符串。
因此，答案是 "aaaz" 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "xaxcd", k = 4
<strong>输出：</strong>"aawcd"
<strong>解释：</strong>在这个例子中，可以执行以下操作：
将 s[0] 改为 'a' ，s 变为 "aaxcd" 。
将 s[2] 改为 'w' ，s 变为 "aawcd" 。
"xaxcd" 和 "aawcd" 之间的距离等于 k = 4 。
可以证明 "aawcd" 是在任意次操作后能够得到的字典序最小的字符串。
因此，答案是 "aawcd" 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "lol", k = 0
<strong>输出：</strong>"lol"
<strong>解释：</strong>在这个例子中，k = 0，更改任何字符都会使得距离大于 0 。
因此，答案是 "lol" 。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>0 &lt;= k &lt;= 2000</code></li>
	<li><code>s</code> 只包含小写英文字母。</li>
</ul>

#### 思路

- 每次操作可以把 $s[i]$ 加一或减一，求在操作次数不超过 $k$ 的前提下，$s$ 的最小字典序。

算法：

1. 从左到右遍历 $s$。
2. 如果把 $s[i]$ 变成 $\texttt{a}$ 的操作次数 $\textit{dis} \le k$，那么就把 $s[i]$ 变成 $\texttt{a}$，同时 $k$ 减少 $\textit{dis}$。
3. 否则，把 $s[i]$ 减少 $k$，退出循环。

``` go 
func getSmallestString(s string, k int) string {
	t := []byte(s)
	for i, c := range t {
		dis := int(min(c-'a', 'z'-c+1))
		if dis > k {
			t[i] -= byte(k)
			break
		}
		t[i] = 'a'
		k -= dis
	}
	return string(t)
}
```


#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)