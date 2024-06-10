#### 题目

<p>给你一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>maxHeights</code> 。</p>
<p>你的任务是在坐标轴上建 <code>n</code> 座塔。第 <code>i</code> 座塔的下标为 <code>i</code> ，高度为 <code>heights[i]</code> 。</p>

<p>如果以下条件满足，我们称这些塔是 <strong>美丽</strong> 的：</p>

<ol>
	<li><code>1 <= heights[i] <= maxHeights[i]</code></li>
	<li><code>heights</code> 是一个 <strong>山状</strong> 数组。</li>
</ol>

<p>如果存在下标 <code>i</code> 满足以下条件，那么我们称数组 <code>heights</code> 是一个 <strong>山状</strong> 数组：</p>

<ul>
	<li>对于所有 <code>0 < j <= i</code> ，都有 <code>heights[j - 1] <= heights[j]</code></li>
	<li>对于所有 <code>i <= k < n - 1</code> ，都有 <code>heights[k + 1] <= heights[k]</code></li>
</ul>

<p>请你返回满足 <b>美丽塔</b> 要求的方案中，<strong>高度和的最大值</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>maxHeights = [5,3,4,1,1]
<b>输出：</b>13
<b>解释：</b>和最大的美丽塔方案为 heights = [5,3,3,1,1] ，这是一个美丽塔方案，因为：
- 1 <= heights[i] <= maxHeights[i]  
- heights 是个山状数组，峰值在 i = 0 处。
13 是所有美丽塔方案中的最大高度和。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>maxHeights = [6,5,3,9,2,7]
<b>输出：</b>22
<strong>解释：</strong> 和最大的美丽塔方案为 heights = [3,3,3,9,2,2] ，这是一个美丽塔方案，因为：
- 1 <= heights[i] <= maxHeights[i]
- heights 是个山状数组，峰值在 i = 3 处。
22 是所有美丽塔方案中的最大高度和。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>maxHeights = [3,2,5,5,2,3]
<b>输出：</b>18
<strong>解释：</strong>和最大的美丽塔方案为 heights = [2,2,5,5,2,2] ，这是一个美丽塔方案，因为：
- 1 <= heights[i] <= maxHeights[i]
- heights 是个山状数组，最大值在 i = 2 处。
注意，在这个方案中，i = 3 也是一个峰值。
18 是所有美丽塔方案中的最大高度和。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n == maxHeights <= 10<sup>5</sup></code></li>
	<li><code>1 <= maxHeights[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

下面把 $\textit{maxHeights}$ 简记为 $a$。

- 计算从 $a[0]$ 到 $a[i]$ 形成山状数组的左侧递增段，元素和最大是多少，记到数组 $\textit{pre}[i]$ 中。
- 计算从 $a[i]$ 到 $a[n-1]$ 形成山状数组的右侧递减段，元素和最大是多少，记到数组 $\textit{suf}[i]$ 中。

那么答案就是 $\textit{pre}[i]+\textit{suf}[i+1]$ 的最大值。如何计算 $\textit{pre}$ 和 $\textit{suf}$ 呢？用**单调栈**，元素值从栈底到栈顶严格递增。以 $\textit{suf}$ 为例，我们从右往左遍历 $a$，设当前得到的元素和为 $\textit{sum}$。

- 如果 $a[i]$ 大于栈顶的元素值，那么直接把 $a[i]$ 加到 $\textit{sum}$ 中，同时把 $i$ 入栈（栈中只需要保存下标）。
- 否则，只要 $a[i]$ 小于等于栈顶元素值，就不断循环，把之前加到 $\textit{sum}$ 的**撤销**掉。循环结束后，从 $a[i]$ 到 $a[j-1]$（假设现在栈顶下标是 $j$）都必须是 $a[i]$，把 $a[i]\cdot (j-i)$ 加到 $\textit{sum}$ 中。

```go
func maximumSumOfHeights(a []int) int64 {
	ans := 0
	n := len(a)
	suf := make([]int, n+1)
	st := []int{n} // 哨兵
	sum := 0
	for i := n - 1; i >= 0; i-- {
		x := a[i]
		for len(st) > 1 && x <= a[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			sum -= a[j] * (st[len(st)-1] - j) // 撤销之前加到 sum 中的
		}
		sum += x * (st[len(st)-1] - i) // 从 i 到 st[len(st)-1]-1 都是 x
		suf[i] = sum
		st = append(st, i)
	}
	ans = sum

	st = []int{-1} // 哨兵
	pre := 0
	for i, x := range a {
		for len(st) > 1 && x <= a[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			pre -= a[j] * (j - st[len(st)-1]) // 撤销之前加到 pre 中的
		}
		pre += x * (i - st[len(st)-1]) // 从 st[len(st)-1]+1 到 i 都是 x
		ans = max(ans, pre+suf[i+1])
		st = append(st, i)
	}
	return int64(ans)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{maxHeights}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
