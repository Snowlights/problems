#### 题目  

<p>给你一个由正整数组成的数组 <code>nums</code> 和一个 <strong>正</strong> 整数 <code>k</code> 。</p>

<p>如果 <code>nums</code> 的子集中，任意两个整数的绝对差均不等于 <code>k</code> ，则认为该子数组是一个 <strong>美丽</strong> 子集。</p>

<p>返回数组 <code>nums</code> 中 <strong>非空</strong> 且 <strong>美丽</strong> 的子集数目。</p>

<p><code>nums</code> 的子集定义为：可以经由 <code>nums</code> 删除某些元素（也可能不删除）得到的一个数组。只有在删除元素时选择的索引不同的情况下，两个子集才会被视作是不同的子集。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [2,4,6], k = 2
<strong>输出：</strong>4
<strong>解释：</strong>数组 nums 中的美丽子集有：[2], [4], [6], [2, 6] 。
可以证明数组 [2,4,6] 中只存在 4 个美丽子集。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [1], k = 1
<strong>输出：</strong>1
<strong>解释：</strong>数组 nums 中的美丽数组有：[1] 。
可以证明数组 [1] 中只存在 1 个美丽子集。 
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 20</code></li>
	<li><code>1 &lt;= nums[i], k &lt;= 1000</code></li>
</ul>
 
#### 思路  

在枚举 [78. 子集](https://leetcode.cn/problems/subsets/) 的基础上加个判断。  
在选择 $x=\textit{nums}[i]$ 的时候，如果之前选过 $x-k$ 或 $x+k$，则不能选，否则可以选。  
代码实现时，可以用哈希表或者数组来记录选过的数，从而 $O(1)$ 判断 $x-k$ 和 $x+k$ 是否选过。

```go 
func beautifulSubsets(nums []int, k int) int {
	ans := -1                    // 去掉空集
	cnt := make([]int, 1001+k*2) // 用数组实现比哈希表更快
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ans++
			return
		}
		dfs(i + 1)       // 不选
		x := nums[i] + k // 避免负数下标
		if cnt[x-k] == 0 && cnt[x+k] == 0 {
			cnt[x]++ // 选
			dfs(i + 1)
			cnt[x]-- // 恢复现场
		}
	}
	dfs(0)
	return ans
}
```

#### 复杂度分析  

- 时间复杂度：$O(2^n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。用哈希表实现是 $O(n)$。

