#### 题目

<p>给你一个 <strong>正整数 </strong>数组 <code>nums</code>。</p>

<p>小红和小明正在玩游戏。在游戏中，小红可以从 <code>nums</code> 中选择所有个位数 <strong>或</strong> 所有两位数，剩余的数字归小明所有。如果小红所选数字之和 <strong>严格大于 </strong>小明的数字之和，则小红获胜。</p>

<p>如果小红能赢得这场游戏，返回 <code>true</code>；否则，返回 <code>false</code>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [1,2,3,4,10]</span></p>

<p><strong>输出：</strong><span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<p>小红不管选个位数还是两位数都无法赢得比赛。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [1,2,3,4,5,14]</span></p>

<p><strong>输出：</strong><span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<p>小红选择个位数可以赢得比赛，所选数字之和为 15。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong><span class="example-io">nums = [5,5,5,25]</span></p>

<p><strong>输出：</strong><span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<p>小红选择两位数可以赢得比赛，所选数字之和为 25。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 99</code></li>
</ul>

#### 思路

设 $s_1$ 为 $\textit{nums}$ 中的所有个位数之和，$s_2$ 为 $\textit{nums}$ 中的所有两位数之和。注意题目保证只有个位数和两位数。
小红若要获胜，必须满足 $s_1 > s_2$ 或者 $s_2 > s_1$，即

$$
s_1 \ne s_2
$$

代码实现时，可以令 $s = s_1 - s_2$，即累加 $\textit{nums}$ 的所有元素，把其中的两位数变成相反数累加。这样最后只需判断 $s\ne 0$ 即可。

> 注：也可以计算 $n-k$ 的二进制中的 $1$ 的个数。

```
func canAliceWin(nums []int) bool {
	s := 0
	for _, x := range nums {
		if x < 10 {
			s += x
		} else {
			s -= x
		}
	}
	return s != 0
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)