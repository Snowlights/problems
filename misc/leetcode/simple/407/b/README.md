#### 题目

<p>小红和小明在玩一个字符串元音游戏。</p>

<p>给你一个字符串 <code>s</code>，小红和小明将轮流参与游戏，小红<strong> 先 </strong>开始：</p>

<ul>
	<li>在小红的回合，她必须移除 <code>s</code> 中包含 <strong>奇数 </strong>个元音的任意 <strong>非空</strong> <span data-keyword="substring">子字符串</span>。</li>
	<li>在小明的回合，他必须移除 <code>s</code> 中包含 <strong>偶数 </strong>个元音的任意 <strong>非空</strong> <span data-keyword="substring">子字符串</span>。</li>
</ul>

<p>第一个无法在其回合内进行移除操作的玩家输掉游戏。假设小红和小明都采取 <strong>最优策略 </strong>。</p>

<p>如果小红赢得游戏，返回 <code>true</code>，否则返回 <code>false</code>。</p>

<p>英文元音字母包括：<code>a</code>, <code>e</code>, <code>i</code>, <code>o</code>, 和 <code>u</code>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "leetcoder"</span></p>

<p><strong>输出：</strong> <span class="example-io">true</span></p>

<p><strong>解释：</strong><br />
小红可以执行如下移除操作来赢得游戏：</p>

<ul>
	<li>小红先手，她可以移除加下划线的子字符串 <code>s = "<u><strong>leetco</strong></u>der"</code>，其中包含 3 个元音。结果字符串为 <code>s = "der"</code>。</li>
	<li>小明接着，他可以移除加下划线的子字符串 <code>s = "<u><strong>d</strong></u>er"</code>，其中包含 0 个元音。结果字符串为 <code>s = "er"</code>。</li>
	<li>小红再次操作，她可以移除整个字符串 <code>s = "<strong><u>er</u></strong>"</code>，其中包含 1 个元音。</li>
	<li>又轮到小明，由于字符串为空，无法执行移除操作，因此小红赢得游戏。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "bbcd"</span></p>

<p><strong>输出：</strong> <span class="example-io">false</span></p>

<p><strong>解释：</strong><br />
小红在她的第一回合无法执行移除操作，因此小红输掉了游戏。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 10<sup>5</sup></code></li>
	<li><code>s</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

分类讨论：
- 如果 $s$ 不包含任何元音，小红输。
- 如果 $s$ 包含奇数个元音，小红可以直接把整个 $s$ 移除，小红赢。
- 如果 $s$ 包含正偶数个元音，由于**偶数减奇数等于奇数**，小红移除任意包含奇数个元音的子串后，剩余元音个数仍然为奇数。由于**奇数减偶数还是奇数**，所以无论小明怎么操作，仍然会剩下奇数个元音，此时小红可以直接把整个 $s$ 移除，小红赢。

所以只要 $s$ 包含元音，就返回 $\texttt{true}$，否则返回 $\texttt{false}$。

```
func doesAliceWin(s string) bool {
	return strings.ContainsAny(s, "aeiou")
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)