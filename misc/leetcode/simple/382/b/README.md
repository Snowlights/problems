#### 题目

<p>给你一个<strong> 正整数 </strong>数组 <code>nums</code> 。</p>

<p>你需要从数组中选出一个满足下述条件的<span data-keyword="subset">子集</span>：</p>

<ul>
	<li>你可以将选中的元素放置在一个下标从 <strong>0</strong> 开始的数组中，并使其遵循以下模式：<code>[x, x<sup>2</sup>, x<sup>4</sup>, ..., x<sup>k/2</sup>, x<sup>k</sup>, x<sup>k/2</sup>, ..., x<sup>4</sup>, x<sup>2</sup>, x]</code>（<strong>注意</strong>，<code>k</code> 可以是任何 <strong>非负</strong> 的 2 的幂）。例如，<code>[2, 4, 16, 4, 2]</code> 和 <code>[3, 9, 3]</code> 都符合这一模式，而 <code>[2, 4, 8, 4, 2]</code> 则不符合。</li>
</ul>

<p>返回满足这些条件的子集中，元素数量的 <strong>最大值 </strong><em>。</em></p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [5,4,1,2,2]
<strong>输出：</strong>3
<strong>解释：</strong>选择子集 {4,2,2} ，将其放在数组 [2,4,2] 中，它遵循该模式，且 2<sup>2</sup> == 4 。因此答案是 3 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,3,2,4]
<strong>输出：</strong>1
<strong>解释：</strong>选择子集 {1}，将其放在数组 [1] 中，它遵循该模式。因此答案是 1 。注意我们也可以选择子集 {2} 、{4} 或 {3} ，可能存在多个子集都能得到相同的答案。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

设数组的最大值为 $m$，当 $xge 2$ 时，我们有

$$
x^{2^p} le m

$$

解得

$$
p le log_2 log_x m

$$

在本题数据范围下，$p$ 至多为 $4$。
所以暴力枚举数组中的数，作为 $x$，然后不断看 $x^2,x^4,cdots$ 在数组中的个数。
直到个数不足 $2$ 个为止，退出循环。
注意模式的正中间的数字只取一个。如果最后 $x$ 有一个，那么个数加一，否则个数减一。
注意特判 $x=1$ 的情况。

```go [sol]
func maximumLength(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	ans := cnt[1] - (cnt[1]%2 ^ 1)
	delete(cnt, 1)
	for x := range cnt {
		res := 0
		for ; cnt[x] > 1; x *= x {
			res += 2
		}
		res += cnt[x]
		ans = max(ans, res-(res%2^1)) // 保证 res 是奇数
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log \log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。\
- 空间复杂度：$\mathcal{O}(n)$。
