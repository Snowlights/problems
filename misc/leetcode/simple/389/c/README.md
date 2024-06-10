#### 题目

<p>给你一个字符串 <code>word</code> 和一个整数 <code>k</code>。</p>

<p>如果&nbsp;<code>|freq(word[i]) - freq(word[j])| &lt;= k</code> 对于字符串中所有下标 <code>i</code> 和 <code>j</code>&nbsp; 都成立，则认为 <code>word</code> 是 <strong>k 特殊字符串</strong>。</p>

<p>此处，<code>freq(x)</code> 表示字符 <code>x</code> 在 <code>word</code> 中的<span data-keyword="frequency-letter">出现频率</span>，而 <code>|y|</code> 表示 <code>y</code> 的绝对值。</p>

<p>返回使 <code>word</code> 成为 <strong>k 特殊字符串</strong> 需要删除的字符的最小数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">word = "aabcaba", k = 0</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">3</span></p>

<p><strong>解释：</strong>可以删除 <code>2</code> 个 <code>"a"</code> 和 <code>1</code> 个 <code>"c"</code> 使 <code>word</code> 成为 <code>0</code> 特殊字符串。<code>word</code> 变为 <code>"baba"</code>，此时 <code>freq('a') == freq('b') == 2</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">word = "dabdcbdcdcd", k = 2</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">2</span></p>

<p><strong>解释：</strong>可以删除 <code>1</code> 个 <code>"a"</code> 和 <code>1</code> 个 <code>"d"</code> 使 <code>word</code> 成为 <code>2</code> 特殊字符串。<code>word</code> 变为 <code>"bdcbdcdcd"</code>，此时 <code>freq('b') == 2</code>，<code>freq('c') == 3</code>，<code>freq('d') == 4</code>。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">word = "aaabaaa", k = 2</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">1</span></p>

<p><strong>解释：</strong>可以删除<strong> </strong>1 个 <code>"b"</code> 使 <code>word</code> 成为 <code>2</code>特殊字符串。因此，<code>word</code> 变为 <code>"aaaaaa"</code>，此时每个字母的频率都是 <code>6</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= k &lt;= 10<sup>5</sup></code></li>
	<li><code>word</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

**定理**：必然有一种字母是不需要删除的。

**反证法**：如果每种字母都至少删除一个，那么可以都增加一，不影响字母数量之差。

统计 $\textit{word}$ 中每个字母的出现次数，记到一个数组 $\textit{cnt}$ 中。

枚举 $i$ 作为出现次数最小的字母，为了保留尽量多的字母，字母 $i$ 肯定不需要删除。此外，出现次数最多的字母，其出现次数不能超过 $\textit{cnt}[i]+k$。

分类讨论：

- 出现次数小于 $\textit{cnt}[i]$ 的字母，全部删除。
- 出现次数大于等于 $\textit{cnt}[i]$ 的字母 $j$，保留 $\min(\textit{cnt}[j], \textit{cnt}[i] + k)$ 个。累加保留的字母个数，更新最多保留的字母个数 $\textit{maxSave}$ 的最大值。

最后用 $\textit{word}$ 的长度，减去 $\textit{maxSave}$，即为答案。

代码实现时，为方便计算，把 $\textit{cnt}$ 从小到大排序。

```go [sol-Go]
func minimumDeletions(word string, k int) int {
	cnt := make([]int, 26)
	for _, b := range word {
		cnt[b-'a']++
	}
	sort.Ints(cnt)

	maxSave := 0
	for i, base := range cnt {
		sum := 0
		for _, c := range cnt[i:] {
			sum += min(c, base+k)
		}
		maxSave = max(maxSave, sum)
	}
	return len(word) - maxSave
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + |\Sigma|^2)$，其中 $n$ 为 $\textit{word}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
