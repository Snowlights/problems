#### 题目

<p>Alice 和 Bob 正在玩一个游戏。最初，Alice 有一个字符串 <code>word = "a"</code>。</p>

<p>给定一个<strong>正整数</strong> <code>k</code>。</p>

<p>现在 Bob 会要求 Alice 执行以下操作<strong> 无限次 </strong>:</p>

<ul>
	<li>将 <code>word</code> 中的每个字符<strong> 更改 </strong>为英文字母表中的<strong> 下一个 </strong>字符来生成一个新字符串，并将其<strong> 追加 </strong>到原始的 <code>word</code>。</li>
</ul>

<p>例如，对 <code>"c"</code> 进行操作生成 <code>"cd"</code>，对 <code>"zb"</code> 进行操作生成 <code>"zbac"</code>。</p>

<p>在执行足够多的操作后， <code>word</code> 中 <strong>至少 </strong>存在 <code>k</code> 个字符，此时返回 <code>word</code> 中第 <code>k</code> 个字符的值。</p>

<p><strong>注意</strong>，在操作中字符 <code>'z'</code> 可以变成 <code>'a'</code>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1:</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">k = 5</span></p>

<p><strong>输出：</strong><span class="example-io">"b"</span></p>

<p><strong>解释：</strong></p>

<p>最初，<code>word = "a"</code>。需要进行三次操作:</p>

<ul>
	<li>生成的字符串是 <code>"b"</code>，<code>word</code> 变为 <code>"ab"</code>。</li>
	<li>生成的字符串是 <code>"bc"</code>，<code>word</code> 变为 <code>"abbc"</code>。</li>
	<li>生成的字符串是 <code>"bccd"</code>，<code>word</code> 变为 <code>"abbcbccd"</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2:</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">k = 10</span></p>

<p><strong>输出：</strong><span class="example-io">"c"</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= k &lt;= 500</code></li>
</ul>

#### 思路

本题相当于周赛第四题所有 $\textit{operations}[i]=1$ 的版本，做法是一样的，请先看 [我的题解](https://leetcode.cn/problems/find-the-k-th-character-in-string-game-ii/solutions/2934284/liang-chong-zuo-fa-di-gui-die-dai-python-5f6z/)。

```
func kthCharacter(k int) byte {
	k--
	ans := byte('a')
	for i := bits.Len(uint(k)) - 1; i >= 0; i-- {
		if k >= 1<<i { // k 在右半边
			ans++
			k -= 1 << i
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log k)$。
- 空间复杂度：$\mathcal{O}(1)$。


上面的代码相当于，每次遇到 $k-1$ 二进制中的 $1$，就把答案加一。

所以答案为 $\texttt{a}$ 加上 $k-1$ 二进制中的 $1$ 的个数。

注意在本题的数据范围下，无需和 $26$ 取模。

```
func kthCharacter(k int) byte {
	return byte('a' + bits.OnesCount(uint(k-1)))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)