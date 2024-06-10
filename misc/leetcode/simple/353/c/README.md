#### 题目  

<p>给你两个下标从 <strong>0</strong> 开始的整数数组 <code>nums1</code> 和 <code>nums2</code> ，长度均为 <code>n</code> 。</p>

<p>让我们定义另一个下标从 <strong>0</strong> 开始、长度为 <code>n</code> 的整数数组，<code>nums3</code> 。对于范围 <code>[0, n - 1]</code> 的每个下标 <code>i</code> ，你可以将 <code>nums1[i]</code> 或 <code>nums2[i]</code> 的值赋给 <code>nums3[i]</code> 。</p>

<p>你的任务是使用最优策略为 <code>nums3</code> 赋值，以最大化 <code>nums3</code> 中 <strong>最长非递减子数组</strong> 的长度。</p>

<p>以整数形式表示并返回 <code>nums3</code> 中 <strong>最长非递减</strong> 子数组的长度。</p>

<p><strong>注意：子数组</strong> 是数组中的一个连续非空元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums1 = [2,3,1], nums2 = [1,2,1]
<strong>输出：</strong>2
<strong>解释：</strong>构造 nums3 的方法之一是： 
nums3 = [nums1[0], nums2[1], nums2[2]] =&gt; [2,2,1]
从下标 0 开始到下标 1 结束，形成了一个长度为 2 的非递减子数组 [2,2] 。 
可以证明 2 是可达到的最大长度。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums1 = [1,3,2,1], nums2 = [2,2,3,4]
<strong>输出：</strong>4
<strong>解释：</strong>构造 nums3 的方法之一是： 
nums3 = [nums1[0], nums2[1], nums2[2], nums2[3]] =&gt; [1,2,3,4]
整个数组形成了一个长度为 4 的非递减子数组，并且是可达到的最大长度。
</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>nums1 = [1,1], nums2 = [2,2]
<strong>输出：</strong>2
<strong>解释：</strong>构造 nums3 的方法之一是： 
nums3 = [nums1[0], nums1[1]] =&gt; [1,1] 
整个数组形成了一个长度为 2 的非递减子数组，并且是可达到的最大长度。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums1.length == nums2.length == n &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums1[i], nums2[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

为方便后面翻译成递推，这里从右往左递归。
定义 $\textit{dfs}(i,j)$ 表示以 $\textit{nums}_j[i]$ 结尾的最长非递减子数组的长度。
用「枚举选哪个」来思考：
- 如果 $\textit{nums}_1[i-1]\le \textit{nums}_j[i]$，那么下一步选 $\textit{nums}_1[i-1]$，有 $\textit{dfs}(i,j) = \textit{dfs}(i-1,0)+1$。
- 如果 $\textit{nums}_2[i-1]\le \textit{nums}_j[i]$，那么下一步选 $\textit{nums}_2[i-1]$，有 $\textit{dfs}(i,j) = \textit{dfs}(i-1,1)+1$。
- 如果都不成立，那么 $\textit{dfs}(i,j)=1$。
  
这几种情况取最大值，即为 $\textit{dfs}(i,j)$。
递归边界：$\textit{dfs}(0)=1$。
递归入口：$\textit{dfs}(i,j)$。遍历所有 $i,j$ 取 $\textit{dfs}(i,j)$ 的最大值，即为答案。

```go 
func maxNonDecreasingLength(nums1, nums2 []int) (ans int) {
	n := len(nums1)
	nums := [2][]int{nums1, nums2}
	memo := make([][2]int, n)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示没有计算过
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := 1
		if nums1[i-1] <= nums[j][i] {
			res = dfs(i-1, 0) + 1
		}
		if nums2[i-1] <= nums[j][i] {
			res = max(res, dfs(i-1, 1)+1)
		}
		*p = res // 记忆化
		return res
	}
	for j := 0; j < 2; j++ {
		for i := 0; i < n; i++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return
}

func max(a, b int) int { if b > a { return b }; return a }
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n)$，单个状态的计算时间为 $\mathcal{O}(1)$，所以动态规划的时间复杂度为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。