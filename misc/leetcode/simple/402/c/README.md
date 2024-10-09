#### 题目

<p>一个魔法师有许多不同的咒语。</p>

<p>给你一个数组&nbsp;<code>power</code>&nbsp;，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。</p>

<p>已知魔法师使用伤害值为&nbsp;<code>power[i]</code>&nbsp;的咒语时，他们就&nbsp;<strong>不能</strong>&nbsp;使用伤害为&nbsp;<code>power[i] - 2</code>&nbsp;，<code>power[i] - 1</code>&nbsp;，<code>power[i] + 1</code>&nbsp;或者&nbsp;<code>power[i] + 2</code>&nbsp;的咒语。</p>

<p>每个咒语最多只能被使用 <strong>一次</strong>&nbsp;。</p>

<p>请你返回这个魔法师可以达到的伤害值之和的 <strong>最大值</strong>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>power = [1,1,3,4]</span></p>

<p><span class="example-io"><b>输出：</b>6</span></p>

<p><strong>解释：</strong></p>

<p>可以使用咒语 0，1，3，伤害值分别为 1，1，4，总伤害值为 6 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>power = [7,1,6,6]</span></p>

<p><span class="example-io"><b>输出：</b>13</span></p>

<p><strong>解释：</strong></p>

<p>可以使用咒语 1，2，3，伤害值分别为 1，6，6，总伤害值为 13 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= power.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= power[i] &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

本题和 [740. 删除并获得点数](https://leetcode.cn/problems/delete-and-earn/) 几乎一样，都是在**值域**上的 [198. 打家劫舍](https://leetcode.cn/problems/house-robber/)。

## 一、记忆化搜索

**前置知识**：[动态规划入门：从记忆化搜索到递推](https://www.bilibili.com/video/BV1Xj411K7oF/)，其中包含把记忆化搜索 1:1 翻译成递推的技巧。

统计每个元素的出现次数，记到哈希表 $\textit{cnt}$ 中。将哈希表的 key 整理到数组 $a$ 中，把 $a$ 按照从小到大的顺序排序。
定义 $\textit{dfs}(i)$ 表示从 $a[0]$ 到 $a[i]$ 中选择，可以得到的伤害值之和的最大值。
考虑 $a[i]$ 选或不选：
- 不选：问题变成从 $a[0]$ 到 $a[i-1]$ 中选择，可以得到的伤害值之和的最大值，即 $\textit{dfs}(i) = \textit{dfs}(i-1)$。
- 选：那么伤害值等于 $a[i]-2$ 和 $a[i]-1$ 的数不能选，问题变成从 $a[0]$ 到 $a[j-1]$ 中选择，可以得到的伤害值之和的最大值，其中 $j$ 是最小的满足 $a[j] \ge a[i]-2$ 的数。那么 $\textit{dfs}(i) = \textit{dfs}(j-1) + a[i]\cdot \textit{cnt}[a[i]]$。

两种情况取最大值，得

$$
\textit{dfs}(i) = \max(\textit{dfs}(i-1), \textit{dfs}(j-1) + a[i]\cdot \textit{cnt}[a[i]])
$$

递归边界：$\textit{dfs}(-1) = 0$。没有数可以选，伤害值之和为 $0$。
递归入口：$\textit{dfs}(n-1)$，即答案。注意这里 $n$ 是 $a$ 的长度，即 $\textit{power}$ 中的不同元素个数。
代码实现时，$j$ 的计算可以用二分查找，也可以暴力用循环查找。

``` 
func maximumTotalDamage(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}
	slices.Sort(a)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		x := a[i]
		j := i
		for j > 0 && a[j-1] >= x-2 {
			j--
		}
		*p = max(dfs(i-1), dfs(j-1)+x*cnt[x])
		return *p
	}
	return int64(dfs(n - 1))
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{power}$ 的长度。瓶颈在排序上，记忆化搜索的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。


## 二、1:1 翻译成递推

我们可以去掉递归中的「递」，只保留「归」的部分，即自底向上计算。
具体来说，$f[i]$ 的定义和 $\textit{dfs}(i)$ 的定义是一样的，都表示从 $a[0]$ 到 $a[i]$ 中选择，可以得到的伤害值之和的最大值。
相应的递推式（状态转移方程）也和 $\textit{dfs}$ 一样：

$$
f[i] = \max(f[i-1], f[j-1] + a[i]\cdot \textit{cnt}[a[i]])
$$

但是，这种定义方式**没有状态能表示递归边界**，即 $i=-1$ 的情况。
解决办法：在 $f$ 数组的最左边插入一个 $0$，那么其余状态全部向右偏移一位，把 $f[i]$ 改为 $f[i+1]$。
修改后 $f[i+1]$ 表示从 $a[0]$ 到 $a[i]$ 中选择，可以得到的伤害值之和的最大值。此时 $f[0]$ 就对应递归边界了。
修改后的递推式为

$$
f[i+1] = \max(f[i], f[j] + a[i]\cdot \textit{cnt}[a[i]])
$$

> 问：为什么 $a$ 的下标不用变？
> 答：既然是在 $f$ 的最左边插入一个状态，那么就只需要修改和 $f$ 有关的下标，其余任何逻辑都无需修改。或者说，如果把 $a[i]$ 也改成 $a[i+1]$，那么 $a[0]$ 就被我们给忽略掉了。

初始值 $f[0]=0$，翻译自递归边界 $\textit{dfs}(-1)=0$。
答案为 $f[n]$，翻译自递归入口 $\textit{dfs}(n-1)$。
代码实现时，$j$ 的计算可以用二分查找，也可以暴力用循环查找，也可以用双指针，其中双指针最快且适用性更广（把题目的 $2$ 改成 $k$ 也可以过）。

``` 
func maximumTotalDamage(power []int) int64 {
	cnt := map[int]int{}
	for _, x := range power {
		cnt[x]++
	}

	n := len(cnt)
	a := make([]int, 0, n)
	for x := range cnt {
		a = append(a, x)
	}
	slices.Sort(a)

	f := make([]int, n+1)
	j := 0
	for i, x := range a {
		for a[j] < x-2 {
			j++
		}
		f[i+1] = max(f[i], f[j]+x*cnt[x])
	}
	return int64(f[n])
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{power}$ 的长度。瓶颈在排序上。双指针的时间复杂度为 $\mathcal{O}(n)$，因为 `j++` 至多执行 $\mathcal{O}(n)$ 次。
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)