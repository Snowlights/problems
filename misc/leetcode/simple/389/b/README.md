#### 题目

<p>给你一个字符串 <code>s</code> 和一个字符 <code>c </code>。返回在字符串 <code>s</code> 中并且以 <code>c</code> 字符开头和结尾的<span data-keyword="substring-nonempty">非空子字符串</span>的总数。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">s = "abada", c = "a"</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">6</span></p>

<p><strong>解释：</strong>以 <code>"a"</code> 开头和结尾的子字符串有： <code>"<strong><u>a</u></strong>bada"</code>、<code>"<u><strong>aba</strong></u>da"</code>、<code>"<u><strong>abada</strong></u>"</code>、<code>"ab<u><strong>a</strong></u>da"</code>、<code>"ab<u><strong>ada</strong></u>"</code>、<code>"abad<u><strong>a</strong></u>"</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block" style="border-color: var(--border-tertiary); border-left-width: 2px; color: var(--text-secondary); font-size: .875rem; margin-bottom: 1rem; margin-top: 1rem; overflow: visible; padding-left: 1rem;">
<p><strong>输入：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">s = "zzz", c = "z"</span></p>

<p><strong>输出：</strong><span class="example-io" style="font-family: Menlo,sans-serif; font-size: 0.85rem;">6</span></p>

<p><strong>解释：</strong>字符串 <code>s</code> 中总共有 <code>6</code> 个子字符串，并且它们都以 <code>"z"</code> 开头和结尾。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> 和 <code>c</code> 均由小写英文字母组成。</li>
</ul>

#### 思路

设 $s$ 中有 $k$ 个 $c$。

每个 $c$ 都可以和它自己或左边的 $c$ 形成（对应着）一个符合要求的子串，所以一共有

$$
1 + 2 + \cdots + k = \dfrac{k(k+1)}{2}
$$

个符合要求的子串。

```go [sol-Go]
func countSubstrings(s string, c byte) int64 {
	k := int64(strings.Count(s, string(c)))
	return k * (k + 1) / 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

