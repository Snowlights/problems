#### 题目

<p>车尔尼有一个数组&nbsp;<code>nums</code>&nbsp;，它只包含 <strong>正</strong>&nbsp;整数，所有正整数的数位长度都 <strong>相同</strong>&nbsp;。</p>

<p>两个整数的 <strong>数位不同</strong>&nbsp;指的是两个整数 <b>相同</b>&nbsp;位置上不同数字的数目。</p>

<p>请车尔尼返回 <code>nums</code>&nbsp;中 <strong>所有</strong>&nbsp;整数对里，<strong>数位不同之和。</strong></p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [13,23,12]</span></p>

<p><b>输出：</b>4</p>

<p><strong>解释：</strong><br />
计算过程如下：<br />
-&nbsp;<strong>1</strong>3 和&nbsp;<strong>2</strong>3 的数位不同为&nbsp;1 。<br />
- 1<strong>3</strong> 和 1<strong>2</strong>&nbsp;的数位不同为&nbsp;1 。<br />
-&nbsp;<strong>23</strong> 和&nbsp;<strong>12</strong>&nbsp;的数位不同为&nbsp;2 。<br />
所以所有整数数对的数位不同之和为&nbsp;<code>1 + 1 + 2 = 4</code>&nbsp;。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [10,10,10,10]</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><strong>解释：</strong><br />
数组中所有整数都相同，所以所有整数数对的数位不同之和为 0 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt; 10<sup>9</sup></code></li>
	<li><code>nums</code>&nbsp;中的整数都有相同的数位长度。</li>
</ul>

#### 思路

横看成岭侧成峰，换一个角度，把每一位拆开，先计算个位数中的不同数对个数，再计算十位数中的不同数对个数，然后是百位数中的不同数对个数，依此类推。

此时问题变成：

- 给你一个长为 $n$ 的数组 $a$，只包含数字 $0$ 到 $9$，其中有多少个不同的数对？

做法有多种，一次遍历的做法如下。

遍历 $a$，同时用数组 $\textit{cnt}$ 统计 $0$ 到 $9$ 中每个数字的出现次数。假设现在遍历到 $d=a[k]$，那么前面有 $k$ 个数字，其中有 $\textit{cnt}[d]$ 个数和 $d$ 是一样的，所以有

$$
k - \textit{cnt}[d]
$$

个数和 $d$ 是不一样的，这正是我们要统计的，加入答案。

代码实现时，既可以外层循环枚举个位数、十位数、百位数等，内层循环枚举 $\textit{nums}$；也可以外层循环枚举 $\textit{nums}$，内层循环枚举个位数、十位数、百位数等。下面代码用的后者。


``` go
func sumDigitDifferences(nums []int) int64 {
	n, m := len(nums), len(strconv.Itoa(nums[0]))
	ans := m * n * (n - 1) / 2
	cnt := make([][10]int, m)
	for _, x := range nums {
		for i := 0; x > 0; x /= 10 {
			d := x % 10
			ans -= cnt[i][d]
			cnt[i][d]++
			i++
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\textit{nums}[0]$。
- 空间复杂度：$\mathcal{O}(D\log U)$，其中 $D=10$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)