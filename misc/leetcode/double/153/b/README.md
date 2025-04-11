### 题目

<p>给你一个长度为 <code>n</code> 的二进制字符串 <code>s</code>，其中：</p>

<ul>
	<li><code>'1'</code> 表示一个 <strong>活跃</strong> 区段。</li>
	<li><code>'0'</code> 表示一个 <strong>非活跃</strong> 区段。</li>
</ul>

<p>你可以执行 <strong>最多一次操作</strong>&nbsp;来最大化 <code>s</code> 中的活跃区段数量。在一次操作中，你可以：</p>

<ul>
	<li>将一个被 <code>'0'</code> 包围的连续 <code>'1'</code> 区块转换为全 <code>'0'</code>。</li>
	<li>然后，将一个被 <code>'1'</code> 包围的连续 <code>'0'</code> 区块转换为全 <code>'1'</code>。</li>
</ul>

<p>返回在执行最优操作后，<code>s</code> 中的 <strong>最大</strong> 活跃区段数。</p>

<p><strong>注意：</strong>处理时需要在 <code>s</code> 的两侧加上 <code>'1'</code> ，即 <code>t = '1' + s + '1'</code>。这些加上的 <code>'1'</code> 不会影响最终的计数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "01"</span></p>

<p><strong>输出：</strong> <span class="example-io">1</span></p>

<p><strong>解释：</strong></p>

<p>因为没有被 <code>'0'</code> 包围的 <code>'1'</code> 区块，因此无法进行有效操作。最大活跃区段数为 1。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "0100"</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>字符串 <code>"0100"</code> → 两端加上 <code>'1'</code>&nbsp;后得到&nbsp;<code>"101001"</code>&nbsp;。</li>
	<li>选择 <code>"0100"</code>，<code>"10<u><strong>1</strong></u>001"</code> → <code>"1<u><strong>0000</strong></u>1"</code> → <code>"1<u><strong>1111</strong></u>1"</code>&nbsp;。</li>
	<li>最终的字符串去掉两端的 <code>'1'</code>&nbsp;后为 <code>"1111"</code>&nbsp;。最大活跃区段数为 4。</li>
</ul>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "1000100"</span></p>

<p><strong>输出：</strong> <span class="example-io">7</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>字符串 <code>"1000100"</code> → 两端加上 <code>'1'</code>&nbsp;后得到 <code>"110001001"</code>&nbsp;。</li>
	<li>选择 <code>"000100"</code>，<code>"11000<u><strong>1</strong></u>001"</code> → <code>"11<u><strong>000000</strong></u>1"</code> → <code>"11<u><strong>111111</strong></u>1"</code>。</li>
	<li>最终的字符串去掉两端的 <code>'1'</code>&nbsp;后为 <code>"1111111"</code>。最大活跃区段数为 7。</li>
</ul>
</div>

<p><strong class="example">示例 4：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "01010"</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>字符串 <code>"01010"</code> → 两端加上 <code>'1'</code>&nbsp;后得到 <code>"1010101"</code>。</li>
	<li>选择 <code>"010"</code>，<code>"10<u><strong>1</strong></u>0101"</code> → <code>"1<u><strong>000</strong></u>101"</code> → <code>"1<u><strong>111</strong></u>101"</code>。</li>
	<li>最终的字符串去掉两端的 <code>'1'</code>&nbsp;后为 <code>"11110"</code>。最大活跃区段数为 4。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n == s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s[i]</code> 仅包含 <code>'0'</code> 或 <code>'1'</code></li>
</ul>

### 思路

⚠**注意**：题目求的是 $\texttt{1}$ 的个数，并没有要求这些 $\texttt{1}$ 是连续的。

根据题意，答案为 $s$ 中 $\texttt{1}$ 的个数，加上一个 $\texttt{010}$ 子串中的 $\texttt{0}$ 的个数。这里 $\texttt{010}$ 子串是指一段连续的 $\texttt{0}$，紧跟着一段连续的 $\texttt{1}$，再紧跟着一段连续的 $\texttt{0}$。

我们需要找一个 $\texttt{0}$ 最多的 $\texttt{010}$ 子串。

遍历 $s$ 的过程中，记录连续相同段的长度 $\textit{cnt}$，以及上一段连续 $\texttt{0}$ 的个数 $\textit{pre}_0$。如果当前这段是 $\texttt{0}$，那么用 $\textit{pre}_0+\textit{cnt}$ 更新 $\textit{mx}$。

最终答案为 $s$ 中 $\texttt{1}$ 的个数加上 $\textit{mx}$。

```
func maxActiveSectionsAfterTrade(s string) (ans int) {
	mx := 0
	pre0 := math.MinInt
	cnt := 0
	for i := range len(s) {
		cnt++
		if i == len(s)-1 || s[i] != s[i+1] {
			if s[i] == '1' {
				ans += cnt
			} else {
				mx = max(mx, pre0+cnt)
				pre0 = cnt
			}
			cnt = 0
		}
	}
	return ans + mx
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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