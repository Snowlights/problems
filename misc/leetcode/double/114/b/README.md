### 题目  

<p>给你一个下标从 <strong>0</strong>&nbsp;开始的正整数数组&nbsp;<code>nums</code>&nbsp;。</p>

<p>你可以对数组执行以下两种操作 <strong>任意次</strong>&nbsp;：</p>

<ul>
	<li>从数组中选择 <strong>两个</strong>&nbsp;值 <strong>相等</strong>&nbsp;的元素，并将它们从数组中 <strong>删除</strong>&nbsp;。</li>
	<li>从数组中选择 <strong>三个</strong>&nbsp;值 <strong>相等</strong>&nbsp;的元素，并将它们从数组中 <strong>删除</strong>&nbsp;。</li>
</ul>

<p>请你返回使数组为空的 <strong>最少</strong>&nbsp;操作次数，如果无法达成，请返回 <code>-1</code>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,3,3,2,2,4,2,3,4]
<b>输出：</b>4
<b>解释：</b>我们可以执行以下操作使数组为空：
- 对下标为 0 和 3 的元素执行第一种操作，得到 nums = [3,3,2,4,2,3,4] 。
- 对下标为 2 和 4 的元素执行第一种操作，得到 nums = [3,3,4,3,4] 。
- 对下标为 0 ，1 和 3 的元素执行第二种操作，得到 nums = [4,4] 。
- 对下标为 0 和 1 的元素执行第一种操作，得到 nums = [] 。
至少需要 4 步操作使数组为空。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [2,1,2,2,3,3]
<b>输出：</b>-1
<b>解释：</b>无法使数组为空。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>

 
### 思路 

先统计每个元素的出现次数。  
考虑出现了 $c$ 次的数：  
- 如果 $c=1$，无法操作，返回 $-1$。
- 如果 $c$ 是 $3$ 的倍数，那么可以用 $\\dfrac{c}{3}$ 次操作删除。
- 如果 $c$ 模 $3$ 为 $1$，那么 $c=(c-4)+4$，其中 $c-4$ 是 $3$ 的倍数，剩余的 $4$ 可以用两次操作完成。
- 如果 $c$ 模 $3$ 为 $2$，那么 $c=(c-2)+2$，其中 $c-2$ 是 $3$ 的倍数，剩余的 $2$ 可以用一次操作完成。
- 总的来说，都需要 $\\left\\lceil\\dfrac{c}{3}\\right\\rceil=\\left\\lfloor\\dfrac{c+2}{3}\\right\\rfloor$ 次操作完成。

```go 
func minOperations(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	for _, c := range cnt {
		if c == 1 {
			return -1
		}
		ans += (c + 2) / 3
	}
	return
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$ 。