#### 题目

<p>给你一个整数数组 <code>nums</code>。该数组包含 <code>n</code> 个元素，其中&nbsp;<strong>恰好&nbsp;</strong>有 <code>n - 2</code> 个元素是&nbsp;<strong>特殊数字&nbsp;</strong>。剩下的&nbsp;<strong>两个&nbsp;</strong>元素中，一个是这些&nbsp;<strong>特殊数字&nbsp;</strong>的 <strong>和</strong> ，另一个是&nbsp;<strong>异常值&nbsp;</strong>。</p>

<p><strong>异常值</strong> 的定义是：既不是原始特殊数字之一，也不是表示这些数字元素和的数字。</p>

<p><strong>注意</strong>，特殊数字、和 以及 异常值 的下标必须&nbsp;<strong>不同&nbsp;</strong>，但可以共享&nbsp;<strong>相同</strong> 的值。</p>

<p>返回 <code>nums</code> 中可能的&nbsp;<strong>最大</strong><strong>异常值</strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [2,3,5,10]</span></p>

<p><strong>输出：</strong> <span class="example-io">10</span></p>

<p><strong>解释：</strong></p>

<p>特殊数字可以是 2 和 3，因此和为 5，异常值为 10。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [-2,-1,-3,-6,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>特殊数字可以是 -2、-1 和 -3，因此和为 -6，异常值为 4。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,1,1,1,1,5,5]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong></p>

<p>特殊数字可以是 1、1、1、1 和 1，因此和为 5，另一个 5 为异常值。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-1000 &lt;= nums[i] &lt;= 1000</code></li>
	<li>输入保证 <code>nums</code> 中至少存在&nbsp;<strong>一个&nbsp;</strong>可能的异常值。</li>
</ul>

#### 思路

## 分析

设异常值为 $x$，元素和为 $y$，那么其余 $n-2$ 个数的和也是 $y$，所以 $x+2y$ 就是整个 $\textit{nums}$ 数组的元素和 $\textit{total}$，即

$$
x+2y = \textit{total}
$$

也就是说，问题相当于从 $\textit{nums}$ 中选出两个（下标不同的）数 $x$ 和 $y$，满足 $x+2y$ 等于一个定值。你需要计算 $x$ 的最大值是多少。

这是我们最熟悉的 [1. 两数之和](https://leetcode.cn/problems/two-sum/)。

## 方法一：枚举异常值

枚举异常值 $x=\textit{nums}[i]$，那么有

$$
2y = \textit{total}-x
$$

如果 $\textit{total}-x$ 是偶数，且 $y=\dfrac{\textit{total}-x}{2}$ 在（除去 $x$ 的）其余 $n-1$ 个数中，那么 $x$ 就是一个异常值。

用哈希表记录每个数的出现次数，从而加快判断。

### 写法一

```
func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, x := range nums {
		cnt[x]--
		if (total-x)%2 == 0 && cnt[(total-x)/2] > 0 {
			ans = max(ans, x)
		}
		cnt[x]++
	}
	return ans
}
```

### 写法二

```
func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, x := range nums {
		if (total-x)%2 == 0 {
			y := (total - x) / 2
			if cnt[y] > 1 || cnt[y] > 0 && y != x {
				ans = max(ans, x)
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：枚举元素和

枚举 $y=\textit{nums}[i]$，那么异常值等于

$$
\textit{total} - 2y
$$

如果 $\textit{total} - 2y$ 在（去掉 $y$ 之后的）哈希表中，那么 $\textit{total} - 2y$ 就是一个异常值。

```
func getLargestOutlier(nums []int) int {
	cnt := map[int]int{}
	total := 0
	for _, x := range nums {
		cnt[x]++
		total += x
	}

	ans := math.MinInt
	for _, y := range nums {
		t := total - y*2
		if cnt[t] > 1 || cnt[t] > 0 && t != y {
			ans = max(ans, t)
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

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
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)