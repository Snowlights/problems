#### 题目

<p>给你一个由 <strong>正整数 </strong>组成的数组 <code>nums</code> 。</p>

<p>返回数组 <code>nums</code> 中所有具有 <strong>最大 </strong>频率的元素的 <strong>总频率 </strong>。</p>

<p>元素的 <strong>频率 </strong>是指该元素在数组中出现的次数。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,2,3,1,4]
<strong>输出：</strong>4
<strong>解释：</strong>元素 1 和 2 的频率为 2 ，是数组中的最大频率。
因此具有最大频率的元素在数组中的数量是 4 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3,4,5]
<strong>输出：</strong>5
<strong>解释：</strong>数组中的所有元素的频率都为 1 ，是最大频率。
因此具有最大频率的元素在数组中的数量是 5 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 100</code></li>
	<li><code>1 <= nums[i] <= 100</code></li>
</ul>

#### 思路

遍历 $\textit{nums}$，同时用哈希表统计每个元素的出现次数，
并维护出现次数的最大值 $\textit{maxCnt}$：

- 如果出现次数 $c > \textit{maxCnt}$， 那么更新 $\textit{maxCnt}=c$，答案 $\textit{ans} = c$。
- 如果出现次数 $c = \textit{maxCnt}$，那么答案增加 $c$。

```go [sol]
func maxFrequencyElements(nums []int) int {
	cnt, ans, mx := make(map[int]int), 0, 0
	for _, v := range nums {
		cnt[v]++
		c := cnt[v]
		if c > mx {
			mx = c
			ans = c
		} else if c == mx {
			ans += c
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
