#### 题目

<p>给你一个整数数组 <code>nums</code>。</p>

<p><strong>因子得分 </strong>定义为数组所有元素的最小公倍数（LCM）与最大公约数（GCD）的<strong> 乘积</strong>。</p>

<p>在 <strong>最多</strong> 移除一个元素的情况下，返回 <code>nums</code> 的<strong> 最大因子得分</strong>。</p>

<p><strong>注意</strong>，单个数字的 LCM 和 GCD 都是其本身，而<strong> </strong><strong>空数组</strong> 的因子得分为 0。</p>

<p><code>lcm(a, b)</code> 表示 <code>a</code> 和 <code>b</code> 的 <strong>最小公倍数</strong>。</p>

<p><code>gcd(a, b)</code> 表示 <code>a</code> 和 <code>b</code> 的<strong> 最大公约数</strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [2,4,8,16]</span></p>

<p><strong>输出：</strong> <span class="example-io">64</span></p>

<p><strong>解释：</strong></p>

<p>移除数字 2 后，剩余元素的 GCD 为 4，LCM 为 16，因此最大因子得分为 <code>4 * 16 = 64</code>。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [1,2,3,4,5]</span></p>

<p><strong>输出：</strong> <span class="example-io">60</span></p>

<p><strong>解释：</strong></p>

<p>无需移除任何元素即可获得最大因子得分 60。</p>
</div>

<p><strong class="example">示例 3：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">nums = [3]</span></p>

<p><strong>输出：</strong> 9</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 100</code></li>
	<li><code>1 &lt;= nums[i] &lt;= 30</code></li>
</ul>

#### 思路

思路同 [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)，本题维护前后缀的 GCD 和 LCM。

细节：

1. 由于 $0$ 可以被任何非零整数整除，所以 $\text{GCD}(0,x)=x$。
2. 由于任何整数都是 $1$ 的倍数，所以 $\text{LCM}(1,x)=x$。


```
func maxScore(nums []int) int64 {
	n := len(nums)
	sufGcd := make([]int, n+1)
	sufLcm := make([]int, n+1)
	sufLcm[n] = 1
	for i, x := range slices.Backward(nums) {
		sufGcd[i] = gcd(sufGcd[i+1], x)
		sufLcm[i] = lcm(sufLcm[i+1], x)
	}

	ans := sufGcd[0] * sufLcm[0]
	preGcd, preLcm := 0, 1
	for i, x := range nums {
		ans = max(ans, gcd(preGcd, sufGcd[i+1])*lcm(preLcm, sufLcm[i+1]))
		preGcd = gcd(preGcd, x)
		preLcm = lcm(preLcm, x)
	}
	return int64(ans)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
```


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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)