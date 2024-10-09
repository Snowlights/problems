#### 题目

<p>给你一个字符串 <code>s</code> 和一个整数 <code>k</code>。请你使用以下算法加密字符串：</p>

<ul>
	<li>对于字符串 <code>s</code> 中的每个字符 <code>c</code>，用字符串中 <code>c</code> 后面的第 <code>k</code> 个字符替换 <code>c</code>（以循环方式）。</li>
</ul>

<p>返回加密后的字符串。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "dart", k = 3</span></p>

<p><strong>输出：</strong> <span class="example-io">"tdar"</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>对于 <code>i = 0</code>，<code>'d'</code> 后面的第 3 个字符是 <code>'t'</code>。</li>
	<li>对于 <code>i = 1</code>，<code>'a'</code> 后面的第 3 个字符是 <code>'d'</code>。</li>
	<li>对于 <code>i = 2</code>，<code>'r'</code> 后面的第 3 个字符是 <code>'a'</code>。</li>
	<li>对于 <code>i = 3</code>，<code>'t'</code> 后面的第 3 个字符是 <code>'r'</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "aaa", k = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">"aaa"</span></p>

<p><strong>解释：</strong></p>

<p>由于所有字符都相同，加密后的字符串也将相同。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>4</sup></code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

先把 $k$ 更新为 $k\bmod n$。\n\n根据题意，我们相当于把 $s$ **循环左移** $k$ 位，所以答案为 $s[k:] + s[:k]$。

```
func f(n, m int) int {
	odd := int(math.Sqrt(float64(n)))
	even := int((math.Sqrt(float64(m*4+1)) - 1) / 2)
	if odd > even {
		return even*2 + 1
	}
	return odd * 2
}

func maxHeightOfTriangle(red, blue int) int {
	return max(f(red, blue), f(blue, red))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。如果不计入返回值则为 $\mathcal{O}(1)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)