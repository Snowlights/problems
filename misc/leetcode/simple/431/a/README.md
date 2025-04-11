#### 题目

<p>给你一个由&nbsp;<strong>正整数&nbsp;</strong>组成的数组 <code>nums</code>。</p>

<p>如果一个数组 <code>arr</code> 满足 <code>prod(arr) == lcm(arr) * gcd(arr)</code>，则称其为&nbsp;<strong>乘积等价数组&nbsp;</strong>，其中：</p>

<ul>
	<li><code>prod(arr)</code> 表示 <code>arr</code> 中所有元素的乘积。</li>
	<li><code>gcd(arr)</code> 表示 <code>arr</code> 中所有元素的最大公因数 (GCD)。</li>
	<li><code>lcm(arr)</code> 表示 <code>arr</code> 中所有元素的最小公倍数 (LCM)。</li>
</ul>

<p>返回数组 <code>nums</code> 的&nbsp;<strong>最长</strong> <strong>乘积等价子数组&nbsp;</strong>的长度。</p>

<p><strong>子数组&nbsp;</strong>是数组中连续的、非空的元素序列。</p>

<p>术语 <code>gcd(a, b)</code> 表示 <code>a</code> 和 <code>b</code> 的&nbsp;<strong>最大公因数&nbsp;</strong>。</p>

<p>术语 <code>lcm(a, b)</code> 表示 <code>a</code> 和 <code>b</code> 的&nbsp;<strong>最小公倍数&nbsp;</strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,1,2,1,1,1]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>

<p><strong>解释：</strong>&nbsp;</p>

<p>最长的乘积等价子数组是 <code>[1, 2, 1, 1, 1]</code>，其中&nbsp;<code>prod([1, 2, 1, 1, 1]) = 2</code>，&nbsp;<code>gcd([1, 2, 1, 1, 1]) = 1</code>，以及&nbsp;<code>lcm([1, 2, 1, 1, 1]) = 2</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [2,3,4,5,6]</span></p>

<p><strong>输出：</strong> <span class="example-io">3</span></p>

<p><strong>解释：</strong>&nbsp;</p>

<p>最长的乘积等价子数组是 <code>[3, 4, 5]</code>。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3,1,4,5,1]</span></p>

<p><strong>输出：</strong> <span class="example-io">5</span></p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10</code></li>
</ul>

#### 思路

## 方法一：枚举子数组左右端点

暴力枚举所有子数组（左右端点），计算子数组的元素乘积、$\text{LCM}$ 和 $\text{GCD}$。

注意乘积太大会溢出。设 $\textit{allLcm}$ 为所有元素的 $\text{LCM}$，我们可以枚举到子数组乘积超过 $\textit{allLcm}\cdot \max(\textit{nums})$ 为止，这二者分别是 $\text{LCM}$ 的最大值和 $\text{GCD}$ 的最大值。

注意题目保证 $n\ge 2$，又由于两个数一定满足题目要求（见方法二），所以合法子数组是一定存在的，且长度至少是 $2$。

> 注：$\text{LCM}(1,2,\ldots,10)=2520$。

```
func maxLength(nums []int) (ans int) {
	mx := slices.Max(nums)
	allLcm := 1
	for _, x := range nums {
		allLcm = lcm(allLcm, x)
	}

	for i := range nums {
		m, l, g := 1, 1, 0
		for j := i; j < len(nums) && m <= allLcm*mx; j++ {
			x := nums[j]
			m *= x
			l = lcm(l, x)
			g = gcd(g, x)
			if m == l*g {
				ans = max(ans, j-i+1)
			}
		}
	}
	return
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
func lcm(a, b int) int { return a / gcd(a, b) * b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：滑动窗口

考虑乘积、$\text{LCM}$ 和 $\text{GCD}$ 的质因数分解中的一个质数 $p$ 的指数 $E$：

- 在子数组乘积中，$E$ 为各个元素的 $p$ 的指数 $e_i$ 之和，即 $E=e_1+e_2+\cdots+e_k$，其中 $k$ 为子数组的长度。
- 在子数组 $\text{LCM}$ 中，$E$ 为各个元素的 $p$ 的指数 $e_i$ 的最大值，即 $E=\max(e_1,e_2,\ldots,e_k)$。
- 在子数组 $\text{GCD}$ 中，$E$ 为各个元素的 $p$ 的指数 $e_i$ 的最小值，即 $E=\min(e_1,e_2,\ldots,e_k)$。

如果 $k=2$，那么 $e_1+e_2 = \max(e_1,e_2) + \min(e_1,e_2)$ 恒成立。所以答案至少是 $2$。

如果 $k=3$，那么 $e_1+e_2+e_3 = \max(e_1,e_2,e_3) + \min(e_1,e_2,e_3)$ 当且仅当至少两个 $e_i=0$ 时成立。证明：

- 如果三个数都是 $0$，那么等式显然成立。
- 如果有两个数是 $0$，设非零的数是 $e_1$，那么等式左边是 $e_1$，右边是 $e_1+0=e_1$，所以等式成立。
- 如果有一个数是 $0$，设非零的数是 $e_1$ 和 $e_2$，那么等式左边是 $e_1+e_2$，右边是 $e_1+0=e_1$，所以等式不成立（左边大于右边）。
- 如果没有 $0$，那么等式左边是 $e_1+e_2+e_3$，右边是 $e_1+e_3$（假设 $e_2$ 是中间大小的数），所以等式不成立（左边大于右边）。

推广到一般情况，当 $k\ge 3$ 时，$e_1+e_2+\cdots+e_k = \max(e_1,e_2,\ldots,e_k) + \min(e_1,e_2,\ldots,e_k)$ 当且仅当至少 $k-1$ 个 $e_i=0$ 时成立。

这意味着，如果 $k\ge 3$，那么这 $k$ 个数必须**两两互质**。如果两个数有大于 $1$ 的公因子，那么不满足「至少 $k-1$ 个 $e_i=0$」的要求。

两两互质是一个很强的性质，我们可以把子数组的所有数「压缩」到乘积 $m$ 中，计算乘积与即将加入子数组的数 $x$ 的 $\text{GCD}(m,x)$，如果 $\text{GCD}(m,x)>1$，则说明 $x$ 与子数组的某个数不是互质的。

类似 [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) 的 [滑动窗口做法](https://leetcode.cn/problems/longest-substring-without-repeating-characters/solutions/1959540/xia-biao-zong-suan-cuo-qing-kan-zhe-by-e-iaks/)，先去掉子数组中的与 $x$ 不互质的数，再把 $x$ 加到子数组中。

```
func maxLength(nums []int) int {
	ans, mul, left := 2, 1, 0
	for right, x := range nums {
		for gcd(mul, x) > 1 {
			mul /= nums[left]
			left++
		}
		mul *= x
		ans = max(ans, right-left+1)
	}
	return ans
}

func gcd(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
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
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)