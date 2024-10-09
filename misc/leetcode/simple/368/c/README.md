#### 题目

<p>给你一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。</p>

<p>我们想将下标进行分组，使得 <code>[0, n - 1]</code> 内所有下标 <code>i</code> 都 <strong>恰好</strong> 被分到其中一组。</p>

<p>如果以下条件成立，我们说这个分组方案是合法的：</p>

<ul>
	<li>对于每个组 <code>g</code> ，同一组内所有下标在 <code>nums</code> 中对应的数值都相等。</li>
	<li>对于任意两个组 <code>g<sub>1</sub></code> 和 <code>g<sub>2</sub></code> ，两个组中 <strong>下标数量</strong> 的 <strong>差值不超过 </strong><code>1</code> 。</li>
</ul>

<p>请你返回一个整数，表示得到一个合法分组方案的 <strong>最少</strong> 组数。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,2,3,2,3]
<b>输出：</b>2
<b>解释：</b>一个得到 2 个分组的方案如下，中括号内的数字都是下标：
组 1 -> [0,2,4]
组 2 -> [1,3]
所有下标都只属于一个组。
组 1 中，nums[0] == nums[2] == nums[4] ，所有下标对应的数值都相等。
组 2 中，nums[1] == nums[3] ，所有下标对应的数值都相等。
组 1 中下标数目为 3 ，组 2 中下标数目为 2 。
两者之差不超过 1 。
无法得到一个小于 2 组的答案，因为如果只有 1 组，组内所有下标对应的数值都要相等。
所以答案为 2 。</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [10,10,10,3,1,1]
<b>输出：</b>4
<b>解释：</b>一个得到 2 个分组的方案如下，中括号内的数字都是下标：
组 1 -> [0]
组 2 -> [1,2]
组 3 -> [3]
组 4 -> [4,5]
分组方案满足题目要求的两个条件。
无法得到一个小于 4 组的答案。
所以答案为 4 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

统计每个数字的出现次数，记在哈希表 $\textit{cnt}$ 中。
假设可以分成大小为 $k$ 和 $k+1$ 的组，现在需要算出每个 $\textit{cnt}[x]$ 最少可以分成多少组。
举例说明，假设 $\textit{\textit{cnt}}[x]=32$，$k=10$，那么 $32=10+10+10+2$，多出的 $2$ 可以分成两个 $1$，加到两个 $10$ 中，从而得到 $11,11,10$ 这三组。
但如果 $\textit{\textit{cnt}}[x]=34$，那么 $34=10+10+10+4$，多出的 $4$ 无法加到另外三个 $10$ 中。
设 $q=\left\lfloor\dfrac{\textit{cnt}[x]}{k}\right\rfloor$，$r = \textit{cnt}[x] \bmod k$。
如果 $q < r$ 则无法分成 $k$ 和 $k+1$ 组，否则一定可以分组。
由于现在只有 $k$ 和 $k+1$，在可以分组的前提下，按照 $k+1$ 来划分，相比按照 $k$ 分出的组数更少，所以最少可以分出

$$
\left\lceil\dfrac{\textit{cnt}[x]}{k+1}\right\rceil

$$

组。累加组数即为分组个数。
例如 $\textit{cnt}[x] = 10$，如果按照 $k=2$ 划分，可以分出 $5$ 组（$10=2+2+2+2+2$），但是按照 $k+1=3$ 来划分，可以分出 $4$ 组（$10=3+3+2+2$）。
你可能会问：为什么不在 $k=3$ 的时候去计算组数呢？那是因为 $k=3$ 算的是分出大小为 $3$ 和 $4$ 的组，这里是分出大小为 $2$ 和 $3$ 的组。

从 $\min(\textit{cnt}[x])$ 开始倒着枚举 $k$，只要可以分，就立刻返回答案。

```go  
func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, c := range cnt {
		k = min(k, c)
	}
	for ; ; k-- {
		ans := 0
		for _, c := range cnt {
			if c/k < c%k {
				ans = 0
				break
			}
			ans += (c + k) / (k + 1)
		}
		if ans > 0 {
			return ans
		}
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\sqrt{n})$。哈希表中至多有 $\mathcal{O}(\sqrt{n})$ 个 key。
