#### 题目

<p>给你一个 <span data-keyword="binary-string">二进制字符串</span> <code>s</code>。</p>

<p>你可以对这个字符串执行 <strong>任意次</strong> 下述操作：</p>

<ul>
	<li>选择字符串中的任一下标 <code>i</code>（ <code>i + 1 < s.length</code> ），该下标满足 <code>s[i] == '1'</code> 且 <code>s[i + 1] == '0'</code>。</li>
	<li>将字符 <code>s[i]</code> 向 <strong>右移 </strong>直到它到达字符串的末端或另一个 <code>'1'</code>。例如，对于 <code>s = "010010"</code>，如果我们选择 <code>i = 1</code>，结果字符串将会是 <code>s = "0<strong><u>001</u></strong>10"</code>。</li>
</ul>

<p>返回你能执行的<strong> 最大 </strong>操作次数。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "1001101"</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>可以执行以下操作：</p>

<ul>
	<li>选择下标 <code>i = 0</code>。结果字符串为 <code>s = "<u><strong>001</strong></u>1101"</code>。</li>
	<li>选择下标 <code>i = 4</code>。结果字符串为 <code>s = "0011<u><strong>01</strong></u>1"</code>。</li>
	<li>选择下标 <code>i = 3</code>。结果字符串为 <code>s = "001<strong><u>01</u></strong>11"</code>。</li>
	<li>选择下标 <code>i = 2</code>。结果字符串为 <code>s = "00<strong><u>01</u></strong>111"</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "00111"</span></p>

<p><strong>输出：</strong> <span class="example-io">0</span></p>
</div>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= s.length <= 10<sup>5</sup></code></li>
	<li><code>s[i]</code> 为 <code>'0'</code> 或 <code>'1'</code>。</li>
</ul>

#### 思路

把 $1$ 当作**车**，想象有一条长为 $n$ 的道路上有一些车。
我们要把所有的车都排到最右边，例如 $011010$ 最终要变成 $000111$。
下文用 $s\\rightarrow t$ 表示车的位置变化。
如果我们优先操作右边的车，那么每辆车都只需操作一次：

$$
\begin{aligned}
& 011010      \\
\rightarrow{} & 011001        \\
\rightarrow{} & 010011        \\
\rightarrow{} & 000111        \\
\end{aligned}
$$

一共需要操作 $3$ 次。

> 注意一次操作可以让一辆车移动多次。

如果我们优先操作左边的（能移动的）车，这会制造大量的「**堵车**」，那么每辆车的操作次数就会更多。

$$
\begin{aligned}
& 011010      \\
\rightarrow{} & 010110        \\
\rightarrow{} & 001110        \\
\rightarrow{} & 001101        \\
\rightarrow{} & 001011        \\
\rightarrow{} & 000111        \\
\end{aligned}
$$

一共需要操作 $5$ 次。

## 算法

1. 从左到右遍历 $s$，同时用一个变量 $\\textit{cnt}_1$ 维护遍历到的 $1$ 的个数。
2. 如果 $s[i]$ 是 $1$，把 $\\textit{cnt}_1$ 增加 $1$。
3. 如果 $s[i]$ 是 $0$ 且 $s[i-1]$ 是 $1$，意味着我们找到了一段道路，可以让 $i$ **左边的每辆车都操作一次**，把答案增加 $\\textit{cnt}_1$。
4. 遍历结束，返回答案。

```
func maxOperations(s string) (ans int) {
	cnt1 := 0
	for i, c := range s {
		if c == '1' {
			cnt1++
		} else if i > 0 && s[i-1] == '1' {
			ans += cnt1
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
