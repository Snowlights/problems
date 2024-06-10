#### 题目

<p>给你一个由 <strong>正</strong> 整数组成的数组 <code>nums</code> 。</p>

<p>如果数组中的某个子数组满足下述条件，则称之为 <strong>完全子数组</strong> ：</p>

<ul>
	<li>子数组中 <strong>不同</strong> 元素的数目等于整个数组不同元素的数目。</li>
</ul>

<p>返回数组中 <strong>完全子数组</strong> 的数目。</p>

<p><strong>子数组</strong> 是数组中的一个连续非空序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,1,2,2]
<strong>输出：</strong>4
<strong>解释：</strong>完全子数组有：[1,3,1,2]、[1,3,1,2,2]、[3,1,2] 和 [3,1,2,2] 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [5,5,5,5]
<strong>输出：</strong>10
<strong>解释：</strong>数组仅由整数 5 组成，所以任意子数组都满足完全子数组的条件。子数组的总数为 10 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 1000</code></li>
	<li><code>1 <= nums[i] <= 2000</code></li>
</ul>

#### 思路

经典滑窗

```go  
func countCompleteSubarrays(nums []int) (ans int) {
	set := map[int]struct{}{}
	for _, v := range nums {
		set[v] = struct{}{}
	}
	m := len(set)

	cnt := map[int]int{}
	left := 0
	for _, v := range nums { // 枚举子数组右端点 v=nums[i]
		ans += left // 子数组左端点 < left 的都是合法的
		cnt[v]++
		for len(cnt) == m {
			ans++ // 子数组左端点等于 left 是合法的
			x := nums[left]
			cnt[x]--
			if cnt[x] == 0 {
				delete(cnt, x)
			}
			left++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。
