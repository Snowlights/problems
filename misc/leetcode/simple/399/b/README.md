#### 题目

<p>给你一个字符串 <code>word</code>，请你使用以下算法进行压缩：</p>

<ul>
	<li>从空字符串 <code>comp</code> 开始。当 <code>word</code> <strong>不为空</strong> 时，执行以下操作：
	<ul>
		<li>移除 <code>word</code> 的最长单字符前缀，该前缀由单一字符 <code>c</code> 重复多次组成，且该前缀长度 <strong>最多 </strong>为 9 。</li>
		<li>将前缀的长度和字符 <code>c</code> 追加到 <code>comp</code> 。</li>
	</ul>
	</li>
</ul>

<p>返回字符串 <code>comp</code> 。</p>

<p>&nbsp;</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "abcde"</span></p>

<p><strong>输出：</strong><span class="example-io">"1a1b1c1d1e"</span></p>

<p><strong>解释：</strong></p>

<p>初始时，<code>comp = ""</code> 。进行 5 次操作，每次操作分别选择 <code>"a"</code>、<code>"b"</code>、<code>"c"</code>、<code>"d"</code> 和 <code>"e"</code> 作为前缀。</p>

<p>对每个前缀，将 <code>"1"</code> 和对应的字符追加到 <code>comp</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">word = "aaaaaaaaaaaaaabb"</span></p>

<p><strong>输出：</strong><span class="example-io">"9a5a2b"</span></p>

<p><strong>解释：</strong></p>

<p>初始时，<code>comp = ""</code>。进行 3 次操作，每次操作分别选择 <code>"aaaaaaaaa"</code>、<code>"aaaaa"</code> 和 <code>"bb"</code> 作为前缀。</p>

<ul>
	<li>对于前缀 <code>"aaaaaaaaa"</code>，将 <code>"9"</code> 和 <code>"a"</code> 追加到 <code>comp</code>。</li>
	<li>对于前缀 <code>"aaaaa"</code>，将 <code>"5"</code> 和 <code>"a"</code> 追加到 <code>comp</code>。</li>
	<li>对于前缀 <code>"bb"</code>，将 <code>"2"</code> 和 <code>"b"</code> 追加到 <code>comp</code>。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= word.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>word</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

统计连续相同字母的长度，按题意模拟即可。


``` 
func compressedString(word string) string {
	t := []byte{}
	i0 := -1
	for i := range word {
		c := word[i]
		if i+1 == len(word) || c != word[i+1] {
			k := i - i0
			t = append(t, bytes.Repeat([]byte{'9', c}, k/9)...)
			if k%9 > 0 {
				t = append(t, '0'+byte(k%9), c)
			}
			i0 = i
		}
	}
	return string(t)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{word}$ 的长度。
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
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)