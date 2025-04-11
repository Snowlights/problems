#### 题目

<p>给你一个整数数组 <code>nums</code> 和一个整数 <code>k</code>。</p>

<p>你可以对数组中的每个元素&nbsp;<strong>最多</strong> 执行 <strong>一次&nbsp;</strong>以下操作：</p>

<ul>
	<li>将一个在范围&nbsp;<code>[-k, k]</code> 内的整数加到该元素上。</li>
</ul>

<p>返回执行这些操作后，<code>nums</code> 中可能拥有的不同元素的&nbsp;<strong>最大&nbsp;</strong>数量。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,2,3,3,4], k = 2</span></p>

<p><strong>输出：</strong> <span class="example-io">6</span></p>

<p><strong>解释：</strong></p>

<p>对前四个元素执行操作，<code>nums</code> 变为 <code>[-1, 0, 1, 2, 3, 4]</code>，可以获得 6 个不同的元素。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [4,4,4,4], k = 1</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong></p>

<p>对 <code>nums[0]</code> 加 -1，以及对 <code>nums[1]</code> 加 1，<code>nums</code> 变为 <code>[3, 5, 4, 4]</code>，可以获得 3 个不同的元素。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>0 &lt;= k &lt;= 10<sup>9</sup></code></li>
</ul>

#### 思路

来看一个现实生活中的例子：

- 军训的某一天，大家在场地上，某些同学聚在一起。现在教官想让这些同学排成一排，那么最靠左的同学，就**尽量往左边移动**，腾出位置。他旁边的同学，可以**紧挨着**最靠左的同学。依此类推。

把 $\textit{nums}$ 视作 $n$ 个人在一维数轴中的位置，从最左边的人的位置 $a$ 开始思考。

这位同学尽量往左移，位置变成 $a-k$。

$\textit{nums}$ 的次小值 $b$ 呢？这位同学也尽量往左移：

- 比如 $a=4,b=6,k=3$，那么 $a$ 变成 $a-k=1$，$b$ 变成 $b-k=3$。
- 比如 $a=4,b=4,k=3$，那么 $a$ 变成 $a'=a-k=1$，$b$ 变成 $b-k=1$ 就和 $a'$ 一样了，可以稍微大一点（紧挨着 $a'$），把 $b$ 变成 $a'+1=2$。

一般地，$b$ 变成

$$
\max(b-k,a'+1)
$$

但这不能超过 $b+k$，所以最终要变成

$$
\min(\max(b-k,a'+1),b+k)
$$

> 相当于让 $a'+1$ 落在 $[b-k,b+k]$ 中，若超出范围则修正。

第三小的数也同理，通过前一个数可以算出当前元素能变成多少。

最后答案为 $\textit{nums}$ 中的不同元素个数。我们可以在修改的同时统计，如果发现当前元素修改后的值，比上一个元素修改后的值大，那么答案加一。

为了方便计算，把 $\textit{nums}$ 从小到大排序，这样从左到右遍历数组，就相当于从最左边的人开始计算了。

```
func maxDistinctElements(nums []int, k int) (ans int) {
	slices.Sort(nums)
	pre := math.MinInt // 记录每个人左边的人的位置
	for _, x := range nums {
		x = min(max(x-k, pre+1), x+k)
		if x > pre {
			ans++
			pre = x
		}
	}
	return
}
```

### 优化

什么情况下，可以直接返回 $n$？

如果所有元素相同，那么我们只能把元素变成 $[x-k,x+k]$ 范围内的数，这一共有 $2k+1$ 个数。所以当 $2k+1 \ge n$ 时，我们可以让所有数都不同，直接返回 $n$。

如果有不同元素，那么当 $2k+1 \ge n$ 时，就更加可以把所有数都变成不同的。

所以只要 $2k+1 \ge n$，就可以直接返回 $n$。


```
func maxDistinctElements(nums []int, k int) (ans int) {
	n := len(nums)
	if k*2+1 >= n {
		return n
	}

	slices.Sort(nums)
	pre := math.MinInt // 记录每个人左边的人的位置
	for _, x := range nums {
		x = min(max(x-k, pre+1), x+k)
		if x > pre {
			ans++
			pre = x
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。


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