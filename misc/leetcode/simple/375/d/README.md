#### 题目

<p>给你一个下标从 <strong>0</strong> 开始、由 <strong>正整数</strong> 组成的数组 <code>nums</code>。</p>

<p>将数组分割成一个或多个<strong> 连续</strong> 子数组，如果不存在包含了相同数字的两个子数组，则认为是一种 <strong>好分割方案</strong> 。</p>

<p>返回 <code>nums</code> 的 <strong>好分割方案</strong> 的 <strong>数目</strong>。</p>

<p>由于答案可能很大，请返回答案对 <code>10<sup>9</sup> + 7</code> <strong>取余</strong> 的结果。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3,4]
<strong>输出：</strong>8
<strong>解释：</strong>有 8 种 <strong>好分割方案 </strong>：([1], [2], [3], [4]), ([1], [2], [3,4]), ([1], [2,3], [4]), ([1], [2,3,4]), ([1,2], [3], [4]), ([1,2], [3,4]), ([1,2,3], [4]) 和 ([1,2,3,4]) 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,1,1,1]
<strong>输出：</strong>1
<strong>解释：</strong>唯一的 <strong>好分割方案</strong> 是：([1,1,1,1]) 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,1,3]
<strong>输出：</strong>2
<strong>解释：</strong>有 2 种<strong> 好分割方案 </strong>：([1,2,1], [3]) 和 ([1,2,1,3]) 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

考虑如下数组

$$
[3,1,2,1,2,4,4]

$$

题目要求相同数字必须在同一个子数组中，所以两个 $1$ 必须在同一个子数组，
两个 $2$ 也必须在同一个子数组。所以 $[1,2,1,2]$ 这一段必须是完整的，不能分割。

把该数组分到无法再分，得到

$$
[3] + [1,2,1,2] + [4,4]

$$

考虑每个 $+$ 号**选或不选**，一共有 $2^2=4$ 种好分割方案，即

$$
\begin{aligned}
&[3] + [1,2,1,2] + [4,4]\\
&[3] + [1,2,1,2,4,4]\\
&[3,1,2,1,2] + [4,4]\\
&[3,1,2,1,2,4,4]
\end{aligned}

$$



只记录右端点
1. 遍历数组，用哈希表 $r$ 记录 $\textit{nums}[i]$ 最右边的下标。
2. 再次遍历数组，那么第一个区间就是 $[0, r[\textit{nums}[0]]]$。
3. 如果第二个区间和第一个区间有交集，那么合并区间，维护合并后的区间的右端点 $\textit{maxR}$。
4. 如果第二个区间和第一个区间没有交集，把合并后的区间个数 $m$ 加一。怎么判断没有交集？只要第一个区间的 $\textit{maxR}=i$ 就表示没有交集。
5. 返回 $2^{m-1}$。记得取模。

```go  
func numberOfGoodPartitions(nums []int) int {
	h, ans, r := make(map[int]int), 0, -1
	for i, v := range nums {
		h[v] = i
	}
	const mod int = 1e9+7
	pow := func(x, n int) int {
		//x %= mod
		res := 1
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	for i, v := range nums {
		r = max(r, h[v])
		if r == i {
			ans++
		}
	}
	return pow(2, ans -1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $m$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
