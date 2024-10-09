#### 题目

<p>给你一个整数数组 <code>nums</code>。</p>

<p><code>nums</code> 的子序列 <code>sub</code> 的长度为 <code>x</code> ，如果其满足以下条件，则称其为 <strong>有效子序列</strong>：</p>

<ul>
	<li><code>(sub[0] + sub[1]) % 2 == (sub[1] + sub[2]) % 2 == ... == (sub[x - 2] + sub[x - 1]) % 2</code></li>
</ul>

<p>返回 <code>nums</code> 的 <strong>最长的有效子序列</strong> 的长度。</p>

<p>一个&nbsp;<strong>子序列</strong>&nbsp;指的是从原数组中删除一些元素（也可以不删除任何元素），剩余元素保持原来顺序组成的新数组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3,4]</span></p>

<p><strong>输出：</strong> <span class="example-io">4</span></p>

<p><strong>解释：</strong></p>

<p>最长的有效子序列是 <code>[1, 2, 3, 4]</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,1,1,2,1,2]</span></p>

<p><strong>输出：</strong> 6</p>

<p><strong>解释：</strong></p>

<p>最长的有效子序列是 <code>[1, 2, 1, 2, 1, 2]</code>。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,3]</span></p>

<p><strong>输出：</strong> <span class="example-io">2</span></p>

<p><strong>解释：</strong></p>

<p>最长的有效子序列是 <code>[1, 3]</code>。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 2 * 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>7</sup></code></li>
</ul>

#### 思路

## 方法一：考察子序列的最后两项

从左到右遍历 $\textit{nums}$，遍历的同时，维护一个二维数组 $f[y][x]$，表示最后两项模 $k$ 分别为 $y$ 和 $x$ 的子序列的长度。

对于 $x=\textit{nums}[i]\bmod k$，我们可以在「最后两项模 $k$ 分别为 $x$ 和 $y$ 的子序列」的末尾添加 $\textit{nums}[i]$，那么「最后两项模 $k$ 分别为 $y$ 和 $x$ 的子序列」的长度会增加 $1$，即

$$
f[y][x] = f[x][y] + 1
$$

最后答案为 $f[i][j]$ 的最大值。

### 答疑

**问**：如何理解这个递推？它和记忆化搜索的区别是什么？

**答**：对比二者的**计算顺序**。如果用记忆化搜索来做，需要单独计算「最左（或者最右）两项模 $k$ 分别为 $x$ 和 $y$ 的子序列」的长度，这是「单线程」，必须**查找下一个元素的位置**。而递推的计算顺序是，（假设我们先遍历到了元素 $2$，然后遍历到了元素 $3$）一会计算一下「最后两项模 $k$ 分别为 $y$ 和 $2$ 的子序列」，一会又计算一下「最后两项模 $k$ 分别为 $y$ 和 $3$ 的子序列」，这是「多线程」，**没有查找元素位置的过程，遇到谁就处理谁**。


```
func maximumLength(nums []int, k int) (ans int) {
	f := make([][]int, k)
	for i := range f {
		f[i] = make([]int, k)
	}
	for _, x := range nums {
		x %= k
		for y, fxy := range f[x] {
			f[y][x] = fxy + 1
			ans = max(ans, f[y][x])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k^2 + nk)$，其中 $n$ 是 $\textit{nums}$ 的长度。注意创建大小为 $k^2$ 的二维数组需要 $\mathcal{O}(k^2)$ 的时间。
- 空间复杂度：$\mathcal{O}(k^2)$。

## 方法二：枚举余数，考察子序列的最后一项

枚举子序列相邻两项模 $k$ 的结果为 $m=0,1,2,\cdots, k-1$。

如果知道了子序列的最后一项（假设是 $x$），那么子序列的倒数第二项也就确定了，即

$$
(m - x\bmod k + k) \bmod k
$$

加 $k$ 再模 $k$ 是为了在 $m < x\bmod k$ 时，保证计算结果非负。

类似方法一，从左到右遍历 $\textit{nums}$ 的同时，维护一个数组 $f[x]$，表示最后一项模 $k$ 为 $x$ 的子序列的长度。

对于 $x=\textit{nums}[i]\bmod k$，我们可以在「最后一项模 $k$ 为 $(m - x\bmod k + k) \bmod k$ 的子序列」的末尾添加 $\textit{nums}[i]$，那么「最后一项模 $k$ 为 $x$ 的子序列」的长度会增加 $1$，即

$$
f[x] = f[(m - x\bmod k + k) \bmod k] + 1
$$

> Python 更简单，由于允许负数下标，可以直接用 $f[m-x\bmod k]$ 作为转移来源。

遍历结束后（或者遍历中），用 $f[i]$ 更新答案的最大值。

```
func maximumLength(nums []int, k int) (ans int) {
	f := make([]int, k)
	for m := 0; m < k; m++ {
		clear(f)
		for _, x := range nums {
			x %= k
			f[x] = f[(m-x+k)%k] + 1
			ans = max(ans, f[x])
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(k(k+n))$，其中 $n$ 是 $\textit{nums}$ 的长度。注意创建大小为 $k$ 的数组需要 $\mathcal{O}(k)$ 的时间。
- 空间复杂度：$\mathcal{O}(k)$。


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