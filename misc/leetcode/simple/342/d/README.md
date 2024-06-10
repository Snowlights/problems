#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的 <strong>正</strong> 整数数组 <code>nums</code> 。你可以对数组执行以下操作 <strong>任意</strong> 次：</p>

<ul>
	<li>选择一个满足 <code>0 <= i < n - 1</code> 的下标 <code>i</code> ，将 <code>nums[i]</code> 或者 <code>nums[i+1]</code> 两者之一替换成它们的最大公约数。</li>
</ul>

<p>请你返回使数组 <code>nums</code> 中所有元素都等于 <code>1</code> 的 <strong>最少</strong> 操作次数。如果无法让数组全部变成 <code>1</code> ，请你返回 <code>-1</code> 。</p>

<p>两个正整数的最大公约数指的是能整除这两个数的最大正整数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [2,6,3,4]
<b>输出：</b>4
<b>解释：</b>我们可以执行以下操作：
- 选择下标 i = 2 ，将 nums[2] 替换为 gcd(3,4) = 1 ，得到 nums = [2,6,1,4] 。
- 选择下标 i = 1 ，将 nums[1] 替换为 gcd(6,1) = 1 ，得到 nums = [2,1,1,4] 。
- 选择下标 i = 0 ，将 nums[0] 替换为 gcd(2,1) = 1 ，得到 nums = [1,1,1,4] 。
- 选择下标 i = 2 ，将 nums[3] 替换为 gcd(1,4) = 1 ，得到 nums = [1,1,1,1] 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [2,10,6,14]
<b>输出：</b>-1
<b>解释：</b>无法将所有元素都变成 1 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= nums.length <= 50</code></li>
	<li><code>1 <= nums[i] <= 10<sup>6</sup></code></li>
</ul>

#### 思路

## 方法一：计算最短的 GCD 等于 1 的子数组

#### 提示 1

首先，如果所有数的 GCD（最大公约数）大于 $1$，那么无论如何都无法操作出 $1$，我们返回 $-1$。
如果 $\textit{nums}$ 中有一个 $1$，那么从 $1$ 向左向右不断替换就能把所有数变成 $1$。
例如 $[2,2,1,2,2]\rightarrow[2,\underline{1},1,2,2]\rightarrow[\underline{1},1,1,2,2]\rightarrow[1,1,1,\underline{1},2]\rightarrow[1,1,1,1,\underline{1}]$，一共 $n-1=5-1=4$ 次操作。
如果有多个 $1$，那么每个 $1$ 只需要向左修改，最后一个 $1$ 向右修改剩余的数字。
例如 $[2,1,2,1,2]\rightarrow[\underline{1},1,2,1,2]\rightarrow[1,1,\underline{1},1,2]\rightarrow[1,1,1,1,\underline{1}]$，一共 $n-\textit{cnt}_1=5-2=3$ 次操作。这里 $\textit{cnt}_1$ 表示 $\textit{nums}$ 中 $1$ 的个数。
所以如果 $\textit{nums}$ 中有 $1$，答案为

$$
n-\textit{cnt}_1

$$

如果 $\textit{nums}$ 中没有 $1$ 呢？

#### 提示 2

如果 $\textit{nums}$ 中没有 $1$，想办法花费尽量少的操作得出一个 $1$。由于只能操作相邻的数，所以这个 $1$ 必然是一个连续子数组的 GCD。（如果在不连续的情况下得到了 $1$，那么这个 $1$ 只能属于其中某个连续子数组，其余的操作是多余的。）
那么找到最短的 GCD 为 $1$ 的子数组，设其长度为 $\textit{minSize}$，那么我们需要操作 $\textit{minSize}-1$ 次得到 $1$。例如 $[2,6,3,4]$ 中的 $[3,4]$ 可以操作 $2-1=1$ 次得到 $1$。然后就转化成提示 1 中的情况了，最终答案为

$$
(\textit{minSize}-1) + (n-1) = \textit{minSize}+n-2

$$

```go
func minOperations(nums []int) int {
	n, gcdAll, cnt1 := len(nums), 0, 0
	for _, x := range nums {
		gcdAll = gcd(gcdAll, x)
		if x == 1 {
			cnt1++
		}
	}
	if gcdAll > 1 {
		return -1
	}
	if cnt1 > 0 {
		return n - cnt1
	}

	minSize := n
	for i := range nums {
		g := 0
		for j, x := range nums[i:] {
			g = gcd(g, x)
			if g == 1 {
				// 这里本来是 j+1，把 +1 提出来合并到 return 中
				minSize = min(minSize, j)
				break
			}
		}
	}
	return minSize + n - 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n(n+\log U))$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。外层循环时，单看 $g=\textit{nums}[i]$，它因为求 GCD 减半的次数是 $\mathcal{O}(\log U)$ 次，因此内层循环的时间复杂度为 $\mathcal{O}(n+\log U)$，所以总的时间复杂度为 $\mathcal{O}(n(n+\log U))$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。

## 方法二：利用 GCD 的性质

如果 $n=10^5$ 要怎么做？

原理见 [这篇题解的方法二](https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/solution/by-endlesscheng-zai1/)

```go
func minOperations(nums []int) int {
	n, gcdAll, cnt1 := len(nums), 0, 0
	for _, x := range nums {
		gcdAll = gcd(gcdAll, x)
		if x == 1 {
			cnt1++
		}
	}
	if gcdAll > 1 {
		return -1
	}
	if cnt1 > 0 {
		return n - cnt1
	}

	minSize := n
	type result struct{ gcd, i int }
	a := []result{}
	for i, x := range nums {
		for j, r := range a {
			a[j].gcd = gcd(r.gcd, x)
		}
		a = append(a, result{x, i})

		// 去重
		j := 0
		for _, q := range a[1:] {
			if a[j].gcd != q.gcd {
				j++
				a[j] = q
			} else {
				a[j].i = q.i // 相同 gcd 保存最右边的位置
			}
		}
		a = a[:j+1]

		if a[0].gcd == 1 {
			// 这里本来是 i-a[0].i+1，把 +1 提出来合并到 return 中
			minSize = min(minSize, i-a[0].i)
		}
	}
	return minSize + n - 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。单看每个元素，它因为求 GCD 减半的次数是 $\mathcal{O}(\log U)$ 次，并且每次去重的时间复杂度也为 $\mathcal{O}(\log U)$，因此时间复杂度为 $\mathcal{O}(n\log U)$。
- 空间复杂度：$\mathcal{O}(\log U)$。
