#### 题目

<p>给你一个正整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>如果我们执行以下操作 <strong>至多一次</strong>&nbsp;可以让两个整数&nbsp;<code>x</code> 和&nbsp;<code>y</code>&nbsp;相等，那么我们称这个数对是 <strong>近似相等</strong>&nbsp;的：</p>

<ul>
	<li>选择&nbsp;<code>x</code> <strong>或者</strong>&nbsp;<code>y</code> &nbsp;之一，将这个数字中的两个数位交换。</li>
</ul>

<p>请你返回 <code>nums</code>&nbsp;中，下标 <code>i</code>&nbsp;和 <code>j</code>&nbsp;满足&nbsp;<code>i &lt; j</code>&nbsp;且&nbsp;<code>nums[i]</code> 和&nbsp;<code>nums[j]</code> <strong>近似相等</strong>&nbsp;的数对数目。</p>

<p><b>注意</b>&nbsp;，执行操作后一个整数可以有前导 0 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [3,12,30,17,21]</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>近似相等数对包括：</p>

<ul>
	<li>3 和 30 。交换 30 中的数位 3 和 0 ，得到 3 。</li>
	<li>12 和 21 。交换12 中的数位 1 和 2 ，得到 21 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,1,1,1,1]</span></p>

<p><span class="example-io"><b>输出：</b>10</span></p>

<p><strong>解释：</strong></p>

<p>数组中的任意两个元素都是近似相等的。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [123,231]</span></p>

<p><span class="example-io"><b>输出：</b>0</span></p>

<p><strong>解释：</strong></p>

<p>我们无法通过交换 <code>123</code>&nbsp;或者 <code>321</code>&nbsp;中的两个数位得到另一个数。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>

#### 思路

本题和周赛第四题类似，但至多交换一次，
请看 [题解](https://leetcode.cn/problems/count-almost-equal-pairs-ii/solutions/2892072/pai-xu-mei-ju-you-wei-hu-zuo-pythonjavac-vbg0/)。
把第二次交换的代码移除即可得到本题的代码。

``` func countPairs(nums []int) (ans int) {
	slices.Sort(nums)
	cnt := map[int]int{}
	for _, x := range nums {
		set := map[int]struct{}{x: {}} // 不交换
		s := []byte(strconv.Itoa(x))
		m := len(s)
		for i := range s {
			for j := i + 1; j < m; j++ {
				s[i], s[j] = s[j], s[i]
				set[atoi(s)] = struct{}{} // 交换一次
				s[i], s[j] = s[j], s[i]
			}
		}
		for x := range set {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return
}

// 手动转 int 快一些
func atoi(s []byte) (res int) {
	for _, b := range s {
		res = res*10 + int(b&15)
	}
	return
}
```

#### 复杂度分析
 
- 时间复杂度：$\mathcal{O}(n\log n + n\log^3 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log^2 U)$。

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