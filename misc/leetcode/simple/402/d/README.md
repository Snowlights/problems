#### 题目

<p>数组 <code>arr</code>&nbsp;中&nbsp;<strong>大于</strong>&nbsp;前面和后面相邻元素的元素被称为 <strong>峰值</strong>&nbsp;元素。</p>

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;和一个二维整数数组&nbsp;<code>queries</code>&nbsp;。</p>

<p>你需要处理以下两种类型的操作：</p>

<ul>
	<li><code>queries[i] = [1, l<sub>i</sub>, r<sub>i</sub>]</code>&nbsp;，求出子数组&nbsp;<code>nums[l<sub>i</sub>..r<sub>i</sub>]</code>&nbsp;中&nbsp;<strong>峰值</strong>&nbsp;元素的数目。<!-- notionvc: 73b20b7c-e1ab-4dac-86d0-13761094a9ae --></li>
	<li><code>queries[i] = [2, index<sub>i</sub>, val<sub>i</sub>]</code>&nbsp;，将&nbsp;<code>nums[index<sub>i</sub>]</code>&nbsp;变为&nbsp;<code><font face="monospace">val<sub>i</sub></font></code><font face="monospace">&nbsp;。</font></li>
</ul>

<p>请你返回一个数组&nbsp;<code>answer</code>&nbsp;，它依次包含每一个第一种操作的答案。<!-- notionvc: a9ccef22-4061-4b5a-b4cc-a2b2a0e12f30 --></p>

<p><strong>注意：</strong></p>

<ul>
	<li>子数组中 <strong>第一个</strong>&nbsp;和 <strong>最后一个</strong>&nbsp;元素都 <strong>不是</strong>&nbsp;峰值元素。</li>
</ul>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [3,1,4,2,5], queries = [[2,3,4],[1,0,4]]</span></p>

<p><span class="example-io"><b>输出：</b>[0]</span></p>

<p><strong>解释：</strong></p>

<p>第一个操作：我们将&nbsp;<code>nums[3]</code>&nbsp;变为&nbsp;4 ，<code>nums</code>&nbsp;变为&nbsp;<code>[3,1,4,4,5]</code>&nbsp;。</p>

<p>第二个操作：<code>[3,1,4,4,5]</code>&nbsp;中峰值元素的数目为 0 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [4,1,4,2,1,5], queries = [[2,2,4],[1,0,2],[1,0,4]]</span></p>

<p><span class="example-io"><b>输出：</b>[0,1]</span></p>

<p><strong>解释：</strong></p>

<p>第一个操作：<code>nums[2]</code>&nbsp;变为 4 ，它已经是 4 了，所以保持不变。</p>

<p>第二个操作：<code>[4,1,4]</code>&nbsp;中峰值元素的数目为 0 。</p>

<p>第三个操作：第二个 4 是&nbsp;<code>[4,1,4,2,1]</code>&nbsp;中的峰值元素。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= queries.length &lt;= 10<sup>5</sup></code></li>
	<li><code>queries[i][0] == 1</code>&nbsp;或者&nbsp;<code>queries[i][0] == 2</code></li>
	<li>对于所有的 <code>i</code>&nbsp;，都有：
	<ul>
		<li><code>queries[i][0] == 1</code>&nbsp;：<code>0 &lt;= queries[i][1] &lt;= queries[i][2] &lt;= nums.length - 1</code></li>
		<li><code>queries[i][0] == 2</code> ：<code>0 &lt;= queries[i][1] &lt;= nums.length - 1</code>, <code>1 &lt;= queries[i][2] &lt;= 10<sup>5</sup></code></li>
	</ul>
	</li>
</ul>

#### 思路

为方便描述，把 $\textit{nums}$ 记作 $a$。
如果 $1\le i \le n-2$ 且 $a[i-1] < a[i] > a[i+1]$，那么把 $a[i]$ 视作 $1$，否则视作 $0$。
如此转换后，操作 1 相当于计算从 $l+1$ 到 $r-1$ 的子数组的元素和。请注意，题目说的是「子数组」的第一个和最后一个不是峰值元素，注意是子数组，不是整个数组。
由于操作 2 要动态修改数组，我们可以用**树状数组**维护，具体请看 [带你发明树状数组](https://leetcode.cn/problems/range-sum-query-mutable/solution/dai-ni-fa-ming-shu-zhuang-shu-zu-fu-shu-lyfll/)。
具体来说：
1. 先把区间 $[\max(i-1,1),\min(i+1,n-2)]$ 中的峰值元素从树状数组中去掉。
2. 然后修改 $a[i]=\textit{val}$。
3. 最后再把区间 $[\max(i-1,1),\min(i+1,n-2)]$ 中的峰值元素加入树状数组。

``` 
type fenwick []int

func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return res
}

func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func countOfPeaks(nums []int, queries [][]int) (ans []int) {
	n := len(nums)
	f := make(fenwick, n-1)
	update := func(i, val int) {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			f.update(i, val)
		}
	}
	for i := 1; i < n-1; i++ {
		update(i, 1)
	}

	for _, q := range queries {
		if q[0] == 1 {
			ans = append(ans, f.query(q[1]+1, q[2]-1))
			continue
		}
		i := q[1]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			update(j, -1)
		}
		nums[i] = q[2]
		for j := max(i-1, 1); j <= min(i+1, n-2); j++ {
			update(j, 1)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度，$q$ 是 $\textit{queries}$ 的长度。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)