#### 题目

<p><strong>注意：</strong>在这个问题中，操作次数增加为至多&nbsp;<strong>两次</strong>&nbsp;。</p>

<p>给你一个正整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>如果我们执行以下操作 <strong>至多<u>两次</u></strong>&nbsp;可以让两个整数&nbsp;<code>x</code> 和&nbsp;<code>y</code>&nbsp;相等，那么我们称这个数对是 <strong>近似相等</strong>&nbsp;的：</p>

<ul>
	<li>选择&nbsp;<code>x</code> <strong>或者</strong>&nbsp;<code>y</code> &nbsp;之一，将这个数字中的两个数位交换。</li>
</ul>

<p>请你返回 <code>nums</code>&nbsp;中，下标 <code>i</code>&nbsp;和 <code>j</code>&nbsp;满足&nbsp;<code>i &lt; j</code>&nbsp;且&nbsp;<code>nums[i]</code> 和&nbsp;<code>nums[j]</code> <strong>近似相等</strong>&nbsp;的数对数目。</p>

<p><b>注意</b>&nbsp;，执行操作后得到的整数可以有前导 0 。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1023,2310,2130,213]</span></p>

<p><span class="example-io"><b>输出：</b>4</span></p>

<p><strong>解释：</strong></p>

<p>近似相等数对包括：</p>

<ul>
	<li>1023 和 2310 。交换 1023 中数位 1 和 2 ，然后交换数位 0 和 3 ，得到 2310 。</li>
	<li>1023 和 213 。交换 1023 中数位 1 和 0 ，然后交换数位 1 和 2 ，得到 0213 ，也就是 213 。</li>
	<li>2310 和 213 。交换 2310 中数位 2 和 0 ，然后交换数位 3 和 2 ，得到 0213 ，也就是 213 。</li>
	<li>2310 和 2130 。交换 2310 中数位 3 和 1 ，得到 2130 。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [1,10,100]</span></p>

<p><span class="example-io"><b>输出：</b>3</span></p>

<p><strong>解释：</strong></p>

<p>近似相等数对包括：</p>

<ul>
	<li>1 和 10 。交换 10 中数位 1 和 0&nbsp;，得到 01 ，也就是 1&nbsp;。</li>
	<li>1 和 100 。交换 100 中数位 1 和从左往右的第二个 0 ，得到 001 ，也就是 1 。</li>
	<li>10 和 100 。交换 100 中数位 1 和从左往右的第一个 0 ，得到 010 ，也就是 10 。</li>
</ul>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 5000</code></li>
	<li><code>1 &lt;= nums[i] &lt;&nbsp;10<sup>7</sup></code></li>
</ul>

#### 思路

**核心思想**：枚举 $j$ 并**暴力交换** $\textit{nums}[j]$ 中的数位。设交换后的数字为 $x$，统计左边有多少个 $\textit{nums}[i]$ 等于 $x$。\n\n为避免重复统计，可以先用一个哈希集合记录交换后的数字，再去遍历哈希集合中的元素 $x$，统计左边有多少个 $\textit{nums}[i]$ 等于 $x$。\n\n但是，如果 $100$ 这样的数在左边，我们枚举的 $1$ 这样的数在右边，就没法找到近似相等的数对。\n\n怎么办？把 $\textit{nums}$ **按照数字长度从小到大排序**（也就是元素值从小到大排序），即可解决此问题。\n\n⚠**注意**：代码在第二次交换时，$p$ 是从 $i+1$ 开始枚举，而不是从 $i$ 开始枚举的，因为在 $i$ 和 $j$ 以及 $i$ 和 $k$ 上的交换，一定等价于在 $i$ 和 $k$ 以及 $j$ 和 $k$ 上的交换。例如 $456$ 先交换 $i=0$ 和 $j=1$，得 $546$，然后交换 $i=0$ 和 $k=2$，得 $645$；我们也可以先交换 $i=0$ 和 $k=2$，得 $654$，再交换 $j=1$ 和 $k=2$，也可以得到 $645$。


``` 
func countPairs(nums []int) (ans int) {
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
				for p := i + 1; p < m; p++ {
					for q := p + 1; q < m; q++ {
						s[p], s[q] = s[q], s[p]
						set[atoi(s)] = struct{}{} // 交换两次
						s[p], s[q] = s[q], s[p]
					}
				}
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

- 时间复杂度：$\mathcal{O}(n\log n + n\log^5 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log^4 U)$。

## 优化

把 $x=\textit{nums}[i]$ 的数位拆开，放入一个数组 $a$ 中。例如 $x=123$ 拆开得到 $a=[3,2,1]$，注意低位在数组左边，这样下标与 $10$ 的幂次可以对应上。
交换 $a[i]$ 和 $a[j]$，相当于 $x$ 增加了

$$
\begin{aligned}
& (a[j]-a[i])\cdot 10^i + (a[i]-a[j]) \cdot 10^j      \\
={} & (a[j]-a[i])(10^i-10^j)        \\
\end{aligned}
$$

这样就不需要把字符串再转成数字了，预处理 $10^i$ 后，可以 $\mathcal{O}(1)$ 算出交换后的数字。
此外，我们可以在 $a[i]=a[j]$ 时，直接 `continue`，因为交换不改变 $x$ 的值。

``` 
var pow10 = [...]int{1, 10, 100, 1000, 10000, 100000, 1000000}

func countPairs(nums []int) int {
	slices.Sort(nums)
	ans := 0
	cnt := make(map[int]int)
	a := [len(pow10)]int{}
	for _, x := range nums {
		st := map[int]struct{}{x: {}} // 不交换
		m := 0
		for v := x; v > 0; v /= 10 {
			a[m] = v % 10
			m++
		}
		for i := 0; i < m; i++ {
			for j := i + 1; j < m; j++ {
				if a[i] == a[j] { // 小优化
					continue
				}
				y := x + (a[j]-a[i])*(pow10[i]-pow10[j])
				st[y] = struct{}{} // 交换一次
				a[i], a[j] = a[j], a[i]
				for p := i + 1; p < m; p++ {
					for q := p + 1; q < m; q++ {
						st[y+(a[q]-a[p])*(pow10[p]-pow10[q])] = struct{}{} // 交换两次
					}
				}
				a[i], a[j] = a[j], a[i]
			}
		}
		for x := range st {
			ans += cnt[x]
		}
		cnt[x]++
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n + n\log^4 U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(n + \log^4 U)$。

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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)