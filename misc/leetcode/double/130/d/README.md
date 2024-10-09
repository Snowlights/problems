### 题目

<p>一个整数 <code>x</code>&nbsp;的 <strong>强数组</strong>&nbsp;指的是满足和为 <code>x</code> 的二的幂的最短有序数组。比方说，11 的强数组为&nbsp;<code>[1, 2, 8]</code>&nbsp;。</p>

<p>我们将每一个正整数 <code>i</code>&nbsp;（即1，2，3等等）的 <strong>强数组</strong>&nbsp;连接得到数组&nbsp;<code>big_nums</code>&nbsp;，<code>big_nums</code>&nbsp;开始部分为&nbsp;<code>[<u>1</u>, <u>2</u>, <u>1, 2</u>, <u>4</u>, <u>1, 4</u>, <u>2, 4</u>, <u>1, 2, 4</u>, <u>8</u>, ...]</code>&nbsp;。</p>

<p>给你一个二维整数数组&nbsp;<code>queries</code>&nbsp;，其中&nbsp;<code>queries[i] = [from<sub>i</sub>, to<sub>i</sub>, mod<sub>i</sub>]</code>&nbsp;，你需要计算&nbsp;<code>(big_nums[from<sub>i</sub>] * big_nums[from<sub>i</sub> + 1] * ... * big_nums[to<sub>i</sub>]) % mod<sub>i</sub></code>&nbsp;。</p>

<p>请你返回一个整数数组&nbsp;<code>answer</code>&nbsp;，其中&nbsp;<code>answer[i]</code>&nbsp;是第 <code>i</code>&nbsp;个查询的答案。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>queries = [[1,3,7]]</span></p>

<p><span class="example-io"><b>输出：</b>[4]</span></p>

<p><strong>解释：</strong></p>

<p>只有一个查询。</p>

<p><code>big_nums[1..3] = [2,1,2]</code>&nbsp;。它们的乘积为 4 ，4 对 7 取余数得到 4 。</p>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><span class="example-io"><b>输入：</b>queries = [[2,5,3],[7,7,4]]</span></p>

<p><span class="example-io"><b>输出：</b>[2,2]</span></p>

<p><strong>解释：</strong></p>

<p>有两个查询。</p>

<p>第一个查询：<code>big_nums[2..5] = [1,2,4,1]</code>&nbsp;。它们的乘积为 8 ，8 对 3 取余数得到 2 。</p>

<p>第二个查询：<code>big_nums[7] = 2</code>&nbsp;，2 对 4 取余数得到 2 。</p>
</div>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= queries.length &lt;= 500</code></li>
	<li><code>queries[i].length == 3</code></li>
	<li><code>0 &lt;= queries[i][0] &lt;= queries[i][1] &lt;= 10<sup>15</sup></code></li>
	<li><code>1 &lt;= queries[i][2] &lt;= 10<sup>5</sup></code></li>
</ul>

### 思路

## 转换

由于要计算的是 $\textit{bigNums}$ 数组的元素**乘积**，而这些元素都是 $2$ 的幂，所以先计算出**幂次之和**。

$\textit{bigNums}$ 的幂次数组为

$$
[0] + [1] + [0,1] + [2] + [0,2] + [1,2] + [0,1,2] + [3] + \cdots
$$

其中每个小数组内的是 $1,2,3,4,5,6,7,8,\cdots$ 对应的强数组的幂次。

根据前缀和的思想，计算从 $\textit{from}$ 到 $\textit{to}$ 的幂次之和，等于「前 $\textit{to}+1$ 个幂次之和」减去「前 $\textit{from}$ 个幂次之和」。

## 计算幂次个数

为了计算前 $k$ 个幂次之和，我们首先要计算出，这 $k$ 个幂次是由多少个强数组组成的。

例如前 $n=3$ 个强数组，组成了前 $1+1+2=4$ 个幂次。这也等于 $[1,n]$ 中数字的二进制中的 $1$ 的个数之和。

定义 $\texttt{ones}(n)$ 为 $[0,n-1]$ 中数字的二进制中的 $1$ 的个数之和。（$0$ 中没有 $1$，为方便描述，把左边界设为 $0$。）

如果 $n=2^i$，可以证明：

$$
\texttt{ones}(2^i) = i\cdot 2^{i-1}
$$

例如 $\texttt{ones}(2^2)$ 为 $[0,3]$ 中数字的二进制中的 $1$ 的个数之和，即 $1+1+2=4$。

我们要找一个最大的 $n$，满足 $\texttt{ones}(n)\le k$。

这可以用**试填法**计算。以 $k=10$ 为例：

- 从 $i=3$ 开始考虑，因为 $i=4$ 的时候 $[0,2^i-1]$ 中的元素个数已经大于 $k$ 了，即使每个数都只算一个 $1$ 也超过了。
- 假设 $n$ 二进制从低到高第 $i=3$ 位填 $1$，即 $n=8$，那么 $\texttt{ones}(n) = 3\cdot 2^{3-1} = 12 > k$，不能填 $1$，只能填 $0$。
- 假设 $n$ 二进制从低到高第 $i=2$ 位填 $1$，即 $n=4$，那么 $\texttt{ones}(n) = 2\cdot 2^{2-1} = 4 \le k$，可以填 $1$，现在 $n=4$。
- 在 $n=4$ 的基础上，假设 $n$ 二进制从低到高第 $i=1$ 位填 $1$，即 $n=6$，我们来计算相比 $n=4$，$1$ 的个数**增加**了多少，即 $[4,6-1]$ 中的 $1$ 的个数。由于第 $2$ 位都是 $1$，所以增加量是之前填的 $1$ 的个数（$1$ 个）乘以 $[4,6-1]$ 中的元素个数（$2$ 个），加上低两位，即 $[0, 2-1]$ 中的 $1$ 的个数（$1$ 个）。现在 $1$ 的个数为 $4 + 1\cdot 2 + 1 = 7\le k$，可以填 $1$，现在 $n=6$。
- 在 $n=6$ 的基础上，假设 $n$ 二进制从低到高第 $i=0$ 位填 $1$，即 $n=7$，我们来计算相比 $n=6$，$1$ 的个数**增加**了多少，即 $[6,7-1]$ 中的 $1$ 的个数。同上，增加量是之前填的 $1$ 的个数（$2$ 个）乘以 $[6,7-1]$ 中的元素个数（$1$ 个），加上低一位，即 $[0, 1-1]$ 中的 $1$ 的个数（$0$ 个）。现在 $1$ 的个数为 $7 + 2\cdot 1 = 9\le k$，可以填 $1$，所以 $n=7$。
- 注意现在算出的 $1$ 的个数是 $9$，相比 $k=10$ 还差一个，这可以用 $n$ 的最低位补充。
- 总结一下，前 $k=10$ 个幂次分别来自 $1,2,3,4,5,6$ 的强数组中的幂次，以及 $7$ 中最低位的幂次。

## 计算幂次之和

知道了 $n$，现在来计算幂次之和。

定义 $\texttt{sumE}(n)$ 为 $[0,n-1]$ 中数字的强数组的幂次之和。（规定 $0$ 的强数组的幂次之和为 $0$。）

如果 $n=2^i$，可以证明：

$$
\texttt{sumE}(2^i) = \dfrac{i(i-1)}{2}\cdot 2^{i-1}
$$


对于一般的 $n$，计算方式同「计算幂次个数」，如果 $n$ 二进制从低到高第 $i$ 位是 $1$，那么幂次之和的**增加量**，分为如下两部分：

- 之前填的 $1$ 的幂次之和 $\textit{sumI}$ 乘以因为填 $1$ 新增加的元素个数 $2^i$。
- $\texttt{sumE}(2^i)$。

可以在「计算幂次个数」的同时计算 $\texttt{sumE}$。

得到了幂次之和，用**快速幂**计算 $2$ 的幂模 $\textit{mod}$，即为答案。关于快速幂，见 [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)。


```
func sumE(k int) (res int) {
	var n, cnt1, sumI int
	for i := bits.Len(uint(k+1)) - 1; i > 0; i-- {
		c := cnt1<<i + i<<(i-1) // 新增的幂次个数
		if c <= k {
			k -= c
			res += sumI<<i + i*(i-1)/2<<(i-1)
			sumI += i   // 之前填的 1 的幂次之和
			cnt1++      // 之前填的 1 的个数
			n |= 1 << i // 填 1
		}
	}
	// 最低位单独计算
	if cnt1 <= k {
		k -= cnt1
		res += sumI
		n++ // 填 1
	}
	// 剩余的 k 个幂次，由 n 的低 k 个 1 补充
	for ; k > 0; k-- {
		res += bits.TrailingZeros(uint(n))
		n &= n - 1
	}
	return
}

func findProductsOfElements(queries [][]int64) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		er := sumE(int(q[1]) + 1)
		el := sumE(int(q[0]))
		ans[i] = pow(2, er-el, int(q[2]))
	}
	return ans
}

func pow(x, n, mod int) int {
	res := 1 % mod
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

- 时间复杂度：$\mathcal{O}(q\log r)$，其中 $q$ 为 $\textit{queries}$ 的长度，$r=\max(\textit{to}_i)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)