#### 题目  

<p>一个长度为 <code>n</code> 下标从 <strong>0</strong> 开始的整数数组 <code>arr</code> 的 <strong>不平衡数字</strong> 定义为，在 <code>sarr = sorted(arr)</code> 数组中，满足以下条件的下标数目：</p>

<ul>
	<li><code>0 &lt;= i &lt; n - 1</code> ，和</li>
	<li><code>sarr[i+1] - sarr[i] &gt; 1</code></li>
</ul>

<p>这里，<code>sorted(arr)</code> 表示将数组 <code>arr</code> 排序后得到的数组。</p>

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> ，请你返回它所有 <strong>子数组</strong> 的 <strong>不平衡数字</strong> 之和。</p>

<p>子数组指的是一个数组中连续一段 <strong>非空</strong> 的元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [2,3,1,4]
<b>输出：</b>3
<b>解释：</b>总共有 3 个子数组有非 0 不平衡数字：
- 子数组 [3, 1] ，不平衡数字为 1 。
- 子数组 [3, 1, 4] ，不平衡数字为 1 。
- 子数组 [1, 4] ，不平衡数字为 1 。
其他所有子数组的不平衡数字都是 0 ，所以所有子数组的不平衡数字之和为 3 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>nums = [1,3,3,3,5]
<b>输出：</b>8
<b>解释：</b>总共有 7 个子数组有非 0 不平衡数字：
- 子数组 [1, 3] ，不平衡数字为 1 。
- 子数组 [1, 3, 3] ，不平衡数字为 1 。
- 子数组 [1, 3, 3, 3] ，不平衡数字为 1 。
- 子数组 [1, 3, 3, 3, 5] ，不平衡数字为 2 。
- 子数组 [3, 3, 3, 5] ，不平衡数字为 1 。
- 子数组 [3, 3, 5] ，不平衡数字为 1 。
- 子数组 [3, 5] ，不平衡数字为 1 。
其他所有子数组的不平衡数字都是 0 ，所以所有子数组的不平衡数字之和为 8 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 1000</code></li>
	<li><code>1 &lt;= nums[i] &lt;= nums.length</code></li>
</ul>
 
#### 思路  

由于 $n$ 至多为 $1000$，我们可以从左到右枚举 $i$，然后从 $i+1$ 开始向右枚举 $j$。 一边枚举 $j$，一边维护不平衡度 $\textit{cnt}$：
- 如果 $x=\textit{nums}[j]$ 之前出现过，那么子数组排序后必然会和另一个 $x$ 相邻，$\textit{cnt}$ 不变；
- 如果 $x=\textit{nums}[j]$ 之前没出现过，先把 $\textit{cnt}$ 加一，然后看 $x-1$ 和 $x+1$ 是否出现过，出现一个 $\textit{cnt}$ 减一，
  两个都出现 $\textit{cnt}$ 减二。
  
遍历过程中，累加 $\textit{cnt}$，即为答案。

```go 
func sumImbalanceNumbers(nums []int) (ans int) {
	n := len(nums)
	for i, x := range nums {
		vis := make([]int, n+2)
		vis[x] = 1
		cnt := 0
		for j := i + 1; j < n; j++ {
			if x := nums[j]; vis[x] == 0 {
				cnt += 1 - vis[x-1] - vis[x+1]
				vis[x] = 1
			}
			ans += cnt
		}
	}
	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。