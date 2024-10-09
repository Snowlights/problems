#### 题目

<p>给你一个字符串 <code>date</code>，它的格式为 <code>yyyy-mm-dd</code>，表示一个公历日期。</p>

<p><code>date</code> 可以重写为二进制表示，只需要将年、月、日分别转换为对应的二进制表示（不带前导零）并遵循 <code>year-month-day</code> 的格式。</p>

<p>返回 <code>date</code> 的 <strong>二进制</strong> 表示。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">date = "2080-02-29"</span></p>

<p><strong>输出：</strong> <span class="example-io">"100000100000-10-11101"</span></p>

<p><strong>解释：</strong></p>

<p><span class="example-io">100000100000, 10 和 11101 分别是 2080, 02 和 29 的二进制表示。</span></p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">date = "1900-01-01"</span></p>

<p><strong>输出：</strong> <span class="example-io">"11101101100-1-1"</span></p>

<p><strong>解释：</strong></p>

<p><span class="example-io">11101101100, 1 和 1 分别是 1900, 1 和 1 的二进制表示。</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>date.length == 10</code></li>
	<li><code>date[4] == date[7] == '-'</code>，其余的 <code>date[i]</code> 都是数字。</li>
	<li>输入保证 <code>date</code> 代表一个有效的公历日期，日期范围从 1900 年 1 月 1 日到 2100 年 12 月 31 日（包括这两天）。</li>
</ul>

#### 思路

1. 把字符串 $\textit{date}$ 按照 `-` 分割。
2. 对于每个字符串，首先转成数字。
3. 然后转成二进制字符串。
4. 最后，用 `-` 拼接每个二进制字符串。
5. 
```
func convertDateToBinary(date string) string {
	a := strings.Split(date, "-")
	for i := range a {
		x, _ := strconv.Atoi(a[i])
		a[i] = strconv.FormatUint(uint64(x), 2)
	}
	return strings.Join(a, "-")
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{date}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。


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