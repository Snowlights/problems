#### 题目

<p>给你三个<strong> 正整数 </strong><code>n</code> 、<code>x</code> 和 <code>y</code> 。</p>

<p>在城市中，存在编号从 <code>1</code> 到 <code>n</code> 的房屋，由 <code>n</code> 条街道相连。对所有 <code>1 &lt;= i &lt; n</code> ，都存在一条街道连接编号为 <code>i</code> 的房屋与编号为 <code>i + 1</code> 的房屋。另存在一条街道连接编号为 <code>x</code> 的房屋与编号为 <code>y</code> 的房屋。</p>

<p>对于每个 <code>k</code>（<code>1 &lt;= k &lt;= n</code>），你需要找出所有满足要求的 <strong>房屋对 </strong><code>[house<sub>1</sub>, house<sub>2</sub>]</code> ，即从 <code>house<sub>1</sub></code> 到 <code>house<sub>2</sub></code> 需要经过的<strong> 最少</strong> 街道数为 <code>k</code> 。</p>

<p>返回一个下标从 <strong>1</strong> 开始且长度为 <code>n</code> 的数组 <code>result</code> ，其中 <code>result[k]</code> 表示所有满足要求的房屋对的数量，即从一个房屋到另一个房屋需要经过的<strong> 最少 </strong>街道数为 <code>k</code> 。</p>

<p><strong>注意</strong>，<code>x</code> 与 <code>y</code> 可以 <strong>相等 </strong>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/20/example2.png" style="width: 474px; height: 197px;" />
<pre>
<strong>输入：</strong>n = 3, x = 1, y = 3
<strong>输出：</strong>[6,0,0]
<strong>解释：</strong>让我们检视每个房屋对
- 对于房屋对 (1, 2)，可以直接从房屋 1 到房屋 2。
- 对于房屋对 (2, 1)，可以直接从房屋 2 到房屋 1。
- 对于房屋对 (1, 3)，可以直接从房屋 1 到房屋 3。
- 对于房屋对 (3, 1)，可以直接从房屋 3 到房屋 1。
- 对于房屋对 (2, 3)，可以直接从房屋 2 到房屋 3。
- 对于房屋对 (3, 2)，可以直接从房屋 3 到房屋 2。
</pre>

<p><strong class="example">示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/20/example3.png" style="width: 668px; height: 174px;" />
<pre>
<strong>输入：</strong>n = 5, x = 2, y = 4
<strong>输出：</strong>[10,8,2,0,0]
<strong>解释：</strong>对于每个距离 k ，满足要求的房屋对如下：
- 对于 k == 1，满足要求的房屋对有 (1, 2), (2, 1), (2, 3), (3, 2), (2, 4), (4, 2), (3, 4), (4, 3), (4, 5), 以及 (5, 4)。
- 对于 k == 2，满足要求的房屋对有 (1, 3), (3, 1), (1, 4), (4, 1), (2, 5), (5, 2), (3, 5), 以及 (5, 3)。
- 对于 k == 3，满足要求的房屋对有 (1, 5)，以及 (5, 1) 。
- 对于 k == 4 和 k == 5，不存在满足要求的房屋对。
</pre>

<p><strong>示例 3：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2023/12/20/example5.png" style="width: 544px; height: 130px;" />
<pre>
<strong>输入：</strong>n = 4, x = 1, y = 1
<strong>输出：</strong>[6,4,2,0]
<strong>解释：</strong>对于每个距离 k ，满足要求的房屋对如下：
- 对于 k == 1，满足要求的房屋对有 (1, 2), (2, 1), (2, 3), (3, 2), (3, 4), 以及 (4, 3)。
- 对于 k == 2，满足要求的房屋对有 (1, 3), (3, 1), (2, 4), 以及 (4, 2)。
- 对于 k == 3，满足要求的房屋对有 (1, 4), 以及 (4, 1)。
- 对于 k == 4，不存在满足要求的房屋对。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= x, y &lt;= n</code></li>
</ul>

#### 思路

分类讨论+直接计算

题目说 $(i,j)$ 和 $(j,i)$ 这两个房屋对我们都要统计，我们也可以只统计 $i<j$ 的房屋对， 最后把答案乘 $2$。
分类讨论：

- 如果 $i\le x$，那么：
    - 设 $k = (x+y+1)/2$，对于编号在 $[i+1, k]$ 内的房子，我们可以直接到达，把区间 $[1,k-i]$ 都加一。
    - 对于编号在 $[k+1, y-1]$ 内的房子，我们可以用 $x$ 到 $y$ 的边到达，把区间 $[x-i+2, x-i+y-k]$ 都加一。注：从 $i$ 经过 $x$ 到 $y$ 的边，到达 $k+1$ 走过的路径长度为 $x-i+1+y-(k+1) = x-i+y-k$。
    - 对于编号在 $[y, n]$ 内的房子，我们也可以用 $x$ 到 $y$ 的边到达，把区间 $[x-i+1, x-i+1+n-y]$ 都加一。
- 如果 $x < i < \left\lfloor\dfrac{x+y}{2}\right\rfloor$，那么同上：
    - 设 $k = i + (y-x+1)/2$（见方法一），对于编号在 $[i+1, k]$ 内的房子，我们可以直接到达，把区间 $[1,k-i]$ 都加一。
    - 对于编号在 $[k+1, y-1]$ 内的房子，我们可以用 $x$ 到 $y$ 的边到达，把区间 $[i-x+2, i-x+y-k]$ 都加一。注：从 $i$ 经过 $x$ 到 $y$ 的边，到达 $k+1$ 走过的路径长度为 $i-x+1+y-(k+1) = i-x+y-k$。
    - 对于编号在 $[y, n]$ 内的房子，我们也可以用 $x$ 到 $y$ 的边到达，把区间 $[i-x+1, i-x+1+n-y]$ 都加一。
- 对于更大的 $i$，它到右侧的房屋无需通过 $x$ 到 $y$ 的边，我们只需把区间 $[1,n-i]$ 都加一。
  此外，当 $x+1\ge y$ 时，我们可以直接算出 $\textit{ans}[i]=2\cdot (n-i)$，其中 $i$ 从 $1$ 开始。

```go [sol]
func countOfPairs(n, x, y int) []int64 {
	if x > y {
		x, y = y, x
	}

	ans := make([]int64, n)
	if x+1 >= y {
		for i := 1; i < n; i++ {
			ans[i-1] = int64(n-i) * 2
		}
		return ans
	}

	diff := make([]int, n+1)
	add := func(l, r int) {
		diff[l]++
		diff[r+1]--
	}

	for i := 1; i < n; i++ {
		if i <= x {
			k := (x + y + 1) / 2
			add(1, k-i)
			add(x-i+2, x-i+y-k)
			add(x-i+1, x-i+1+n-y)
		} else if i < (x+y)/2 {
			k := i + (y-x+1)/2
			add(1, k-i)
			add(i-x+2, i-x+y-k)
			add(i-x+1, i-x+1+n-y)
		} else {
			add(1, n-i)
		}
	}

	sumD := int64(0)
	for i, d := range diff[1:] {
		sumD += int64(d)
		ans[i] = sumD * 2
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。
