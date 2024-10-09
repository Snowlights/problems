#### 题目

<p>给你一个长度为 <code>n</code> 的字符串 <code>word</code> 和一个整数 <code>k</code> ，其中 <code>k</code> 是 <code>n</code> 的因数。</p>

<p>在一次操作中，你可以选择任意两个下标 <code>i</code> 和 <code>j</code>，其中 <code>0 &lt;= i, j &lt; n</code> ，且这两个下标都可以被 <code>k</code> 整除，然后用从 <code>j</code> 开始的长度为 <code>k</code> 的子串替换从 <code>i</code> 开始的长度为 <code>k</code> 的子串。也就是说，将子串 <code>word[i..i + k - 1]</code> 替换为子串 <code>word[j..j + k - 1]</code> 。</p>

<p>返回使 <code>word</code> 成为 <strong>K 周期字符串</strong> 所需的<strong> 最少</strong> 操作次数。</p>

<p>如果存在某个长度为 <code>k</code> 的字符串 <code>s</code>，使得 <code>word</code> 可以表示为任意次数连接 <code>s</code> ，则称字符串 <code>word</code> 是 <strong>K 周期字符串</strong> 。例如，如果 <code>word == "ababab"</code>，那么 <code>word</code> 就是 <code>s = "ab"</code> 时的 2 周期字符串 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="
    border-color: var(--border-tertiary);
    border-left-width: 2px;
    color: var(--text-secondary);
    font-size: .875rem;
    margin-bottom: 1rem;
    margin-top: 1rem;
    overflow: visible;
    padding-left: 1rem;
">
<p><strong>输入：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;
">word = "leetcodeleet", k = 4</span></p>

<p><strong>输出：</strong><span class="example-io" style="
font-family: Menlo,sans-serif;
font-size: 0.85rem;
">1</span></p>

<p><strong>解释：</strong>可以选择 i = 4 和 j = 0 获得一个 4 周期字符串。这次操作后，word 变为 "leetleetleet" 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="
    border-color: var(--border-tertiary);
    border-left-width: 2px;
    color: var(--text-secondary);
    font-size: .875rem;
    margin-bottom: 1rem;
    margin-top: 1rem;
    overflow: visible;
    padding-left: 1rem;
">
<p><strong>输入：</strong><span class="example-io" style="
    font-family: Menlo,sans-serif;
    font-size: 0.85rem;
">word = "leetcoleet", k = 2</span></p>

<p><strong>输出：</strong>3</p>

<p><strong>解释：</strong>可以执行以下操作获得一个 2 周期字符串。</p>

<table border="1" bordercolor="#ccc" cellpadding="5" cellspacing="0" height="146" style="border-collapse:collapse; text-align: center; vertical-align: middle;">
	<tbody>
		<tr>
			<th>i</th>
			<th>j</th>
			<th>word</th>
		</tr>
		<tr>
			<td style="padding: 5px 15px;">0</td>
			<td style="padding: 5px 15px;">2</td>
			<td style="padding: 5px 15px;">etetcoleet</td>
		</tr>
		<tr>
			<td style="padding: 5px 15px;">4</td>
			<td style="padding: 5px 15px;">0</td>
			<td style="padding: 5px 15px;">etetetleet</td>
		</tr>
		<tr>
			<td style="padding: 5px 15px;">6</td>
			<td style="padding: 5px 15px;">0</td>
			<td style="padding: 5px 15px;">etetetetet</td>
		</tr>
	</tbody>
</table>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == word.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= k &lt;= word.length</code></li>
	<li><code>k</code> 能整除 <code>word.length</code> 。</li>
	<li><code>word</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

根据题意，我们只能选择首字母下标为 $0,k,2k,3k,\cdots,n-k$ 的长为 $k$ 的子串来操作（替换）。

并且，$k$ 周期字符串意味着，所有首字母下标为 $0,k,2k,3k,\cdots,n-k$ 的长为 $k$ 的子串均相等。

为使操作次数尽量少，我们可以计算最多保留多少个子串**不变**。也就是统计 $\textit{word}$ 中的这些首字母下标为 $0,k,2k,3k,\cdots,n-k$ 的长为 $k$ 的子串中，出现次数最多的子串的出现次数 $\textit{mx}$。**用出现次数最多的子串，替换其余子串。**

所以用子串的个数 $\dfrac{n}{k}$ 减去 $\textit{mx}$，就是最少操作次数。

``` 
func minimumOperationsToMakeKPeriodic(word string, k int) int {
	n := len(word)
	cnt := map[string]int{}
	for i := k; i <= n; i += k {
		cnt[word[i-k:i]]++
	}
	mx := 0
	for _, c := range cnt {
		mx = max(mx, c)
	}
	return n/k - mx
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{word}$ 的长度。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)