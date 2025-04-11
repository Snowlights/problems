### 题目

<p>给你两个 <strong>正</strong>&nbsp;整数&nbsp;<code>n</code> 和&nbsp;<code>k</code>&nbsp;。</p>

<p>如果一个整数&nbsp;<code>x</code>&nbsp;满足以下条件，那么它被称为 <strong>k</strong><strong>&nbsp;回文</strong>&nbsp;整数&nbsp;。</p>

<ul>
	<li><code>x</code>&nbsp;是一个&nbsp;<span data-keyword="palindrome-integer">回文整数 。</span></li>
	<li><code>x</code>&nbsp;能被 <code>k</code>&nbsp;整除。</li>
</ul>

<p>如果一个整数的数位重新排列后能得到一个 <strong>k 回文整数</strong>&nbsp;，那么我们称这个整数为&nbsp;<strong>好 </strong>整数。比方说，<code>k = 2</code>&nbsp;，那么&nbsp;2020 可以重新排列得到 2002 ，2002 是一个 k 回文串，所以 2020 是一个好整数。而 1010 无法重新排列数位得到一个 k 回文整数。</p>

<p>请你返回 <code>n</code>&nbsp;个数位的整数中，有多少个 <strong>好</strong>&nbsp;整数。</p>

<p><b>注意</b>&nbsp;，任何整数在重新排列数位之前或者之后 <strong>都不能</strong> 有前导 0 。比方说 1010 不能重排列得到&nbsp;101 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 3, k = 5</span></p>

<p><span class="example-io"><b>输出：</b>27</span></p>

<p><b>解释：</b></p>

<p>部分好整数如下：</p>

<ul>
	<li>551 ，因为它可以重排列得到 515 。</li>
	<li>525 ，因为它已经是一个 k 回文整数。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 1, k = 4</span></p>

<p><span class="example-io"><b>输出：</b>2</span></p>

<p><strong>解释：</strong></p>

<p>两个好整数分别是 4 和 8 。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>n = 5, k = 6</span></p>

<p><span class="example-io"><b>输出：</b>2468</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 10</code></li>
	<li><code>1 &lt;= k &lt;= 9</code></li>
</ul>


### 思路

考虑枚举所有长为 $n$ 的回文数。

首先，知道了回文数的左半边，就知道了回文数的右半边，所以可以枚举回文数的左半边。

设 $m = \left\lfloor\dfrac{n-1}{2}\right\rfloor$，设 $\textit{base} = 10^m$。

在 $[\textit{base}, 10\cdot\textit{base})$ 范围内枚举所有长为 $n$ 的回文数的左半边。

如果回文数 $x$ 能被 $k$ 整除，那么接下来需要解决的问题有两个：

1. 计算 $x$ 有多少个不同的排列。
2. 不能重复统计。

为了保证不重复统计，可以像 [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/solutions/2718519/ha-xi-biao-fen-zu-jian-ji-xie-fa-pythonj-1ukv/) 那样，把 $x$ 的十进制字符串 $s$ **排序**，如果之前遇到过同样的字符串 $t$，那么 $s$ 生成的所有排列，$t$ 也能生成。用哈希表记录排序后的字符串，如果 $s$ 排序后在哈希表中，那么就跳过。

下面是组合数学时间。

本质上计算的是「**有重复元素的排列个数**」。

统计 $s$ 中的每个数字的出现次数 $\textit{cnt}$。

先填最高位。由于不能有前导零，最高位可以填的数有 $n-\textit{cnt}_0$ 个。其余 $n-1$ 个数随便排，有 $(n-1)!$ 种方案。

当然，这里面有重复的，例如 $x=34543$，其中两个 $3$ 和两个 $4$ 的排列就是重复的，由于这两个 $3$ **无法区分**，两个 $4$ **无法区分**，方案数要除以 $2!2!$。为什么是除法。

综上，排列个数为

$$
\dfrac{(n-\textit{cnt}_0)\cdot (n-1)!}{\prod\limits_{i=0}^{9}\textit{cnt}_i!}
$$

加入答案。

```
func countGoodIntegers(n, k int) (ans int64) {
	factorial := make([]int, n+1)
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	base := int(math.Pow10((n - 1) / 2))
	vis := map[string]bool{}
	for i := base; i < base*10; i++ { // 枚举回文数左半边
		x := i
		t := i
		if n%2 > 0 {
			t /= 10
		}
		for ; t > 0; t /= 10 {
			x = x*10 + t%10
		}
		if x%k > 0 { // 回文数不能被 k 整除
			continue
		}

		bs := []byte(strconv.Itoa(x))
		slices.Sort(bs)
		s := string(bs)
		if vis[s] { // 不能重复统计
			continue
		}
		vis[s] = true

		cnt := [10]int{}
		for _, c := range bs {
			cnt[c-'0']++
		}
		res := (n - cnt[0]) * factorial[n-1]
		for _, c := range cnt {
			res /= factorial[c]
		}
		ans += int64(res)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(10^m\cdot n\log n)$，其中 $m = \left\lfloor\dfrac{n-1}{2}\right\rfloor$。
- 空间复杂度：$\mathcal{O}(10^m\cdot n)$。


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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)