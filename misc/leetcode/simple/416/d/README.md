#### 题目

<p>给你两个字符串&nbsp;<code>word1</code> 和&nbsp;<code>word2</code>&nbsp;。</p>

<p>如果一个字符串 <code>x</code>&nbsp;重新排列后，<code>word2</code>&nbsp;是重排字符串的&nbsp;<span data-keyword="string-prefix">前缀</span>&nbsp;，那么我们称字符串&nbsp;<code>x</code>&nbsp;是&nbsp;<strong>合法的</strong>&nbsp;。</p>

<p>请你返回 <code>word1</code>&nbsp;中 <strong>合法</strong>&nbsp;<span data-keyword="substring-nonempty">子字符串</span>&nbsp;的数目。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>word1 = "bcca", word2 = "abc"</span></p>

<p><span class="example-io"><b>输出：</b>1</span></p>

<p><strong>解释：</strong></p>

<p>唯一合法的子字符串是&nbsp;<code>"bcca"</code>&nbsp;，可以重新排列得到&nbsp;<code>"abcc"</code>&nbsp;，<code>"abc"</code>&nbsp;是它的前缀。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>word1 = "abcabc", word2 = "abc"</span></p>

<p><span class="example-io"><b>输出：</b>10</span></p>

<p><strong>解释：</strong></p>

<p>除了长度为 1 和 2 的所有子字符串都是合法的。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>word1 = "abcabc", word2 = "aaabc"</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>
</div>

<p>&nbsp;</p>

<p><strong>解释：</strong></p>

<ul>
	<li><code>1 &lt;= word1.length &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= word2.length &lt;= 10<sup>4</sup></code></li>
	<li><code>word1</code> and <code>word2</code> consist only of lowercase English letters.</li>
</ul>

#### 思路

下文把 $\textit{word}_1$ 和 $\textit{word}_2$ 简称为 $s$ 和 $t$。

由于子串可以重排，只要子串可以涵盖（见 76 题）字符串 $t$，那么子串就可以通过重排，使得 $t$ 是子串的前缀。

所以本题是 [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/) 的求个数版本，做法都是**滑动窗口**，请看 [我的题解](https://leetcode.cn/problems/minimum-window-substring/solutions/2713911/liang-chong-fang-fa-cong-o52mn-dao-omnfu-3ezz/)。

滑动窗口的内层循环结束时，右端点**固定**在 $\textit{right}$，左端点在 $0,1,2,\cdots,\textit{left}-1$ 的所有子串都是合法的，这一共有 $\textit{left}$ 个，把 $\textit{left}$ 加入答案。

```
func validSubstringCount(s, t string) (ans int64) {
	if len(s) < len(t) {
		return 0
	}
	cnt := [26]int{} // t 的字母出现次数与 s 的字母出现次数之差
	for _, b := range t {
		cnt[b-'a']++
	}
	less := 0 // 统计窗口内有多少个字母的出现次数比 t 的少
	for _, c := range cnt {
		if c > 0 {
			less++
		}
	}

	left := 0
	for _, b := range s {
		cnt[b-'a']--
		if cnt[b-'a'] == 0 {
			// 窗口内 b 的出现次数和 t 一样
			less--
		}
		for less == 0 { // 窗口符合要求
			if cnt[s[left]-'a'] == 0 {
				// s[left] 移出窗口后，窗口内 s[left] 的出现次数比 t 的少
				less++
			}
			cnt[s[left]-'a']++
			left++
		}
		ans += int64(left)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+|\Sigma|)$，其中 $n$ 是 $\textit{nums}$ 的长度，$|\Sigma|=26$ 是字符集合的大小。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(|\Sigma|)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
