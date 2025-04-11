#### 题目

<p>给你一个字符串 <code>target</code>。</p>

<p>Alice 将会使用一种特殊的键盘在她的电脑上输入 <code>target</code>，这个键盘<strong> 只有两个 </strong>按键：</p>

<ul>
	<li>按键 1：在屏幕上的字符串后追加字符 <code>'a'</code>。</li>
	<li>按键 2：将屏幕上字符串的 <strong>最后一个 </strong>字符更改为英文字母表中的 <strong>下一个</strong> 字符。例如，<code>'c'</code> 变为 <code>'d'</code>，<code>'z'</code> 变为 <code>'a'</code>。</li>
</ul>

<p><strong>注意</strong>，最初屏幕上是一个<em>空</em>字符串 <code>""</code>，所以她<strong> 只能</strong> 按按键 1。</p>

<p>请你考虑按键次数 <strong>最少</strong> 的情况，按字符串出现顺序，返回 Alice 输入 <code>target</code> 时屏幕上出现的所有字符串列表。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">target = "abc"</span></p>

<p><strong>输出：</strong> <span class="example-io">["a","aa","ab","aba","abb","abc"]</span></p>

<p><strong>解释：</strong></p>

<p>Alice 按键的顺序如下：</p>

<ul>
	<li>按下按键 1，屏幕上的字符串变为 <code>"a"</code>。</li>
	<li>按下按键 1，屏幕上的字符串变为 <code>"aa"</code>。</li>
	<li>按下按键 2，屏幕上的字符串变为 <code>"ab"</code>。</li>
	<li>按下按键 1，屏幕上的字符串变为 <code>"aba"</code>。</li>
	<li>按下按键 2，屏幕上的字符串变为 <code>"abb"</code>。</li>
	<li>按下按键 2，屏幕上的字符串变为 <code>"abc"</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">target = "he"</span></p>

<p><strong>输出：</strong> <span class="example-io">["a","b","c","d","e","f","g","h","ha","hb","hc","hd","he"]</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= target.length &lt;= 400</code></li>
	<li><code>target</code> 仅由小写英文字母组成。</li>
</ul>

#### 思路

注意我们只能添加字母，不能删除字母，也不能修改最后一个字母左边的字母，所以只有把第一个字母改成和 $\textit{target}[0]$ 一样后， 才能继续向后添加/修改字母。
比如 $\textit{target}=\texttt{ccc}$，那么操作过程中产生的字符串一定是 

$$
\texttt{a},\texttt{b},\texttt{c},\texttt{ca},\texttt{cb},\texttt{cc},\texttt{cca},\texttt{ccb},\texttt{ccc}
$$

所以操作方式是**唯一**的，模拟即可。


```
func stringSequence(target string) (ans []string) {
	s := make([]byte, len(target))
	for i, c := range target {
		for j := byte('a'); j <= byte(c); j++ {
			s[i] = j
			ans = append(ans, string(s[:i+1]))
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(|\Sigma|n^2)$，其中 $n$ 是 $\textit{target}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。
- 空间复杂度：$\mathcal{O}(n)$。返回值不计入。

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