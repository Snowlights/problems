#### 题目

<p>给你一个整数数组&nbsp;<code>nums</code>&nbsp;，一个整数&nbsp;<code>k</code>&nbsp;&nbsp;和一个整数&nbsp;<code>multiplier</code>&nbsp;。</p>

<p>你需要对 <code>nums</code>&nbsp;执行 <code>k</code>&nbsp;次操作，每次操作中：</p>

<ul>
	<li>找到 <code>nums</code>&nbsp;中的 <strong>最小</strong>&nbsp;值&nbsp;<code>x</code>&nbsp;，如果存在多个最小值，选择最 <strong>前面</strong>&nbsp;的一个。</li>
	<li>将 <code>x</code>&nbsp;替换为&nbsp;<code>x * multiplier</code>&nbsp;。</li>
</ul>

<p><code>k</code>&nbsp;次操作以后，你需要将 <code>nums</code>&nbsp;中每一个数值对&nbsp;<code>10<sup>9</sup> + 7</code>&nbsp;取余。</p>

<p>请你返回执行完 <code>k</code>&nbsp;次乘运算以及取余运算之后，最终的 <code>nums</code>&nbsp;数组。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [2,1,3,5,6], k = 5, multiplier = 2</span></p>

<p><span class="example-io"><b>输出：</b>[8,4,6,5,6]</span></p>

<p><strong>解释：</strong></p>

<table>
	<tbody>
		<tr>
			<th>操作</th>
			<th>结果</th>
		</tr>
		<tr>
			<td>1 次操作后</td>
			<td>[2, 2, 3, 5, 6]</td>
		</tr>
		<tr>
			<td>2 次操作后</td>
			<td>[4, 2, 3, 5, 6]</td>
		</tr>
		<tr>
			<td>3 次操作后</td>
			<td>[4, 4, 3, 5, 6]</td>
		</tr>
		<tr>
			<td>4 次操作后</td>
			<td>[4, 4, 6, 5, 6]</td>
		</tr>
		<tr>
			<td>5 次操作后</td>
			<td>[8, 4, 6, 5, 6]</td>
		</tr>
		<tr>
			<td>取余操作后</td>
			<td>[8, 4, 6, 5, 6]</td>
		</tr>
	</tbody>
</table>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>nums = [100000,2000], k = 2, multiplier = 1000000</span></p>

<p><span class="example-io"><b>输出：</b>[999999307,999999993]</span></p>

<p><strong>解释：</strong></p>

<table>
	<tbody>
		<tr>
			<th>操作</th>
			<th>结果</th>
		</tr>
		<tr>
			<td>1 次操作后</td>
			<td>[100000, 2000000000]</td>
		</tr>
		<tr>
			<td>2 次操作后</td>
			<td>[100000000000, 2000000000]</td>
		</tr>
		<tr>
			<td>取余操作后</td>
			<td>[999999307, 999999993]</td>
		</tr>
	</tbody>
</table>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= k &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= multiplier &lt;= 10<sup>6</sup></code></li>
</ul>


#### 思路

## 写法一

**核心观察**：对于两个数 $x$ 和 $y$，如果 $x$ 在 $y$ 左边，且 $x\le y$ 以及 $x\cdot \textit{multiplier} > y$，那么操作 $y$ 之后，根据 $x\le y$，我们有 $x\cdot \textit{multiplier} \le y\cdot \textit{multiplier}$，这意味着下一次一定会操作 $x$。继续推导下去，后面的操作顺序是 $y,x,y,x,\cdots$
这意味着当两个数接近时，我们会**交替操作**这两个数，而**不会连续操作同一个数**。
对于更多的数的情况也同理，当这些数接近时，我们会按照从小到大的顺序依次操作这些数。
那么，首先用最小堆手动模拟操作，直到原数组的最大值 $\textit{mx}$ 成为这 $n$ 个数的最小值。根据上面的结论，后面的操作就不需要手动模拟了。
设此时还剩下 $k$ 次操作，那么：
- 对于前 $k\bmod n$ 小的数，还可以再操作 $\left\lfloor\dfrac{k}{n}\right\rfloor+1$ 次。
- 其余元素，还可以再操作 $\left\lfloor\dfrac{k}{n}\right\rfloor$ 次。

用**快速幂**计算操作这么多次后的结果，原理见[【图解】一张图秒懂快速幂](https://leetcode.cn/problems/powx-n/solution/tu-jie-yi-zhang-tu-miao-dong-kuai-su-mi-ykp3i/)。

``` 
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	heap.Init(&h)

	// 模拟，直到堆顶是 mx
	for ; k > 0 && h[0].x < mx; k-- {
		h[0].x *= multiplier
		heap.Fix(&h, 0)
	}

	// 剩余的操作可以直接用公式计算
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		e := k / n
		if i < k%n {
			e++
		}
		nums[p.i] = p.x % mod * pow(multiplier, e) % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n\log_{m} U)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$m=\textit{multiplier}$。瓶颈在模拟那，每个数至多操作 $\mathcal{O}(\log_{m} U)$ 次。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法二：优化

设把每个 $\textit{nums}[i]$ 都操作到至少为 $\textit{nums}$ 的最大值，总共需要操作 $\textit{total}$ 次。
分类讨论：
- 如果 $k < \textit{total}$，那么直接用最小堆暴力模拟 $k$ 次。
- 如果 $k\ge \textit{total}$，我们可以先把每个 $\textit{nums}[i]$ 暴力操作到 $\textit{nums}$ 的最大值，然后剩余的操作直接用公式计算。

此外，可以先把 `pow` 算出来，而不是在循环中计算。


``` 
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	clone := slices.Clone(h)

	// 每个数直接暴力操作到 >= mx
	left := k
outer:
	for i := range h {
		for h[i].x < mx {
			h[i].x *= multiplier
			left--
			if left < 0 {
				break outer
			}
		}
	}

	if left < 0 {
		// 暴力模拟
		h = clone
		heap.Init(&h)
		for ; k > 0; k-- {
			h[0].x *= multiplier
			heap.Fix(&h, 0)
		}
		for _, p := range h {
			nums[p.i] = p.x % mod
		}
		return nums
	}

	// 剩余的操作可以直接用公式计算
	k = left
	pow1 := pow(multiplier, k/n)
	pow2 := pow1 * multiplier % mod
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		pw := pow1
		if i < k%n {
			pw = pow2
		}
		nums[p.i] = p.x % mod * pw % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
```

#### 复杂度分析

设 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$，$m=\textit{multiplier}$。

- 时间复杂度：
  - $k$ 较小时为 $\mathcal{O}(n+k\log n)$。Java 是 $\mathcal{O}(n\log n+k\log n)$。
  - $k$ 较大时为 $\mathcal{O}(n\log_{m} U+ n\log n)$ 或 $\mathcal{O}(n\log_{m} U)$。如果像 C++ 那样使用**快速选择**，时间复杂度为 $\mathcal{O}(n\log_{m} U)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 写法三：进一步优化\

方法二瓶颈在暴力把每个 $\textit{nums}[i]$ 乘到至少为 $\textit{mx}$ 上。
通过预处理，把 $\textit{nums}[i]$ 乘到至少为 $\textit{mx}$，这一步可以做到 $\mathcal{O}(1)$。
根据 [@hqztrue](/u/hqztrue) 的 [题解](https://leetcode.cn/problems/final-array-state-after-k-multiplication-operations-ii/solutions/2892649/xian-xing-suan-fa-by-hqztrue-30fk/)，代码实现如下。
> 注：由于常数比较大，该优化并不明显。仅用来说明当 $k$ 比较大时，存在 $\mathcal{O}(n)$ 的做法

``` 
const mod = 1_000_000_007

func getFinalState(nums []int, k int, multiplier int) []int {
	if multiplier == 1 { // 数组不变
		return nums
	}

	n := len(nums)
	mx := 0
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	clone := slices.Clone(h)

	// 打表，计算出最小的 e 满足 multiplier^e >= 2^i
	mxLen := bits.Len(uint(mx))
	type ep struct{ e, powM int }
	ePowM := make([]ep, 0, mxLen)
	for pow2, powM, e := 1, 1, 0; pow2 <= mx; pow2 <<= 1 {
		if powM < pow2 { // 由于 multiplier >= 2，这里只需写 if 而不是 for
			powM *= multiplier
			e++
		}
		ePowM = append(ePowM, ep{e, powM})
	}

	// 把每个数都操作到 >= mx
	left := k
	for i := range h {
		x := h[i].x
		p := ePowM[mxLen-bits.Len(uint(x))]
		e, powM := p.e, p.powM
		if powM/multiplier*x >= mx { // 多操作了一次
			powM /= multiplier
			e--
		} else if x*powM < mx { // 少操作了一次
			powM *= multiplier
			e++
		}
		left -= e
		if left < 0 {
			break
		}
		h[i].x *= powM
	}

	if left < 0 {
		// 暴力模拟
		h = clone
		heap.Init(&h)
		for ; k > 0; k-- {
			h[0].x *= multiplier
			heap.Fix(&h, 0)
		}
		sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
		for _, p := range h {
			nums[p.i] = p.x % mod
		}
		return nums
	}

	// 剩余的操作可以直接用公式计算
	k = left
	pow1 := pow(multiplier, k/n)
	pow2 := pow1 * multiplier % mod
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		pw := pow1
		if i < k%n {
			pw = pow2
		}
		nums[p.i] = p.x % mod * pw % mod
	}
	return nums
}

type pair struct{ x, i int }
func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair
func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
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

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
