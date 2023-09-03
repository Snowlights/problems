### 题目  

<p>定义一个数组 <code>arr</code> 的 <strong>转换数组</strong> <code>conver</code> 为：</p>

<ul>
	<li><code>conver[i] = arr[i] + max(arr[0..i])</code>，其中 <code>max(arr[0..i])</code> 是满足 <code>0 &lt;= j &lt;= i</code> 的所有 <code>arr[j]</code> 中的最大值。</li>
</ul>

<p>定义一个数组 <code>arr</code> 的 <strong>分数</strong> 为 <code>arr</code> 转换数组中所有元素的和。</p>

<p>给你一个下标从 <strong>0</strong> 开始长度为 <code>n</code> 的整数数组 <code>nums</code> ，请你返回一个长度为 <code>n</code> 的数组<em> </em><code>ans</code> ，其中 <code>ans[i]</code>是前缀 <code>nums[0..i]</code> 的分数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>nums = [2,3,7,5,10]
<b>输出：</b>[4,10,24,36,56]
<b>解释：</b>
对于前缀 [2] ，转换数组为 [4] ，所以分数为 4 。
对于前缀 [2, 3] ，转换数组为 [4, 6] ，所以分数为 10 。
对于前缀 [2, 3, 7] ，转换数组为 [4, 6, 14] ，所以分数为 24 。
对于前缀 [2, 3, 7, 5] ，转换数组为 [4, 6, 14, 12] ，所以分数为 36 。
对于前缀 [2, 3, 7, 5, 10] ，转换数组为 [4, 6, 14, 12, 20] ，所以分数为 56 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [1,1,2,4,8,16]
<b>输出：</b>[2,4,8,16,32,64]
<b>解释：</b>
对于前缀 [1] ，转换数组为 [2] ，所以分数为 2 。
对于前缀 [1, 1]，转换数组为 [2, 2] ，所以分数为 4 。
对于前缀 [1, 1, 2]，转换数组为 [2, 2, 4] ，所以分数为 8 。
对于前缀 [1, 1, 2, 4]，转换数组为 [2, 2, 4, 8] ，所以分数为 16 。
对于前缀 [1, 1, 2, 4, 8]，转换数组为 [2, 2, 4, 8, 16] ，所以分数为 32 。
对于前缀 [1, 1, 2, 4, 8, 16]，转换数组为 [2, 2, 4, 8, 16, 32] ，所以分数为 64 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
### 思路  

一边遍历，一边计算前缀最大值 $\textit{mx}$，以及前缀的得分之和

```go 
func findPrefixScore(nums []int) []int64 {
	ans := make([]int64, len(nums))
	mx, s := 0, 0
	for i, x := range nums {
		mx = max(mx, x) // 前缀最大值
		s += x + mx     // 累加前缀的得分
		ans[i] = int64(s)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

### 复杂度分析  

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(1)$。返回值不计入，仅用到若干额外变量。