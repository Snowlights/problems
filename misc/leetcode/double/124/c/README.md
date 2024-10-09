### 题目

<p>给你一个整数数组 <code>nums</code> ，如果 <code>nums</code> <strong>至少</strong> 包含 <code>2</code> 个元素，你可以执行以下操作中的 <strong>任意</strong> 一个：</p>

<ul>
	<li>选择 <code>nums</code> 中最前面两个元素并且删除它们。</li>
	<li>选择 <code>nums</code> 中最后两个元素并且删除它们。</li>
	<li>选择 <code>nums</code> 中第一个和最后一个元素并且删除它们。</li>
</ul>

<p>一次操作的 <strong>分数</strong> 是被删除元素的和。</p>

<p>在确保<strong> 所有操作分数相同</strong> 的前提下，请你求出 <strong>最多</strong> 能进行多少次操作。</p>

<p>请你返回按照上述要求 <strong>最多</strong> 可以进行的操作次数。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>nums = [3,2,1,2,3,4]
<b>输出：</b>3
<b>解释：</b>我们执行以下操作：
- 删除前两个元素，分数为 3 + 2 = 5 ，nums = [1,2,3,4] 。
- 删除第一个元素和最后一个元素，分数为 1 + 4 = 5 ，nums = [2,3] 。
- 删除第一个元素和最后一个元素，分数为 2 + 3 = 5 ，nums = [] 。
由于 nums 为空，我们无法继续进行任何操作。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>nums = [3,2,6,1,4]
<b>输出：</b>2
<b>解释：</b>我们执行以下操作：
- 删除前两个元素，分数为 3 + 2 = 5 ，nums = [6,1,4] 。
- 删除最后两个元素，分数为 1 + 4 = 5 ，nums = [6] 。
至多进行 2 次操作。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= nums.length <= 2000</code></li>
	<li><code>1 <= nums[i] <= 1000</code></li>
</ul>

### 思路

由于只能从两端删除，所以子问题是「从两侧向内缩小的」，使用区间 DP 解决。

如果确定了第一次删除的元素和，那么后续删除的元素和也就确定了，所以至多有三种不同的元素和，分别用三次区间 DP 解决。  
定义 $\textit{dfs}(i,j)$ 表示当前剩余元素从 $\textit{nums}[i]$ 到 $\textit{nums}[j]$，此时最多可以进行的操作次数。  
枚举三种操作方式，分别从 $\textit{dfs}(i+2,j)+1, \textit{dfs}(i,j-2)+1, \textit{dfs}(i+1,j-1)+1$ 转移过来，取最大值，即为 $\textit{dfs}(i,j)$。
如果三种操作方式都不行，那么 $\textit{dfs}(i,j)=0$。  
递归终点：如果 $i\ge j$，此时至多剩下一个数，无法操作，返回 $0$。  
递归入口：根据三种初始操作，分别为 $\textit{dfs}(2,n-1), \textit{dfs}(0,n-3), \textit{dfs}(1,n-2)$。

```go [sol]
func maxOperations(nums []int) int {
	n := len(nums)
	res1 := helper(nums[2:], nums[0]+nums[1])       // 最前面两个
	res2 := helper(nums[:n-2], nums[n-2]+nums[n-1]) // 最后两个
	res3 := helper(nums[1:n-1], nums[0]+nums[n-1])  // 第一个和最后一个
	return max(res1, max(res2, res3)) + 1           // 加上第一次操作
}

func helper(a []int, target int) int {
	n := len(a)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i >= j {
			return
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		if a[i]+a[i+1] == target { // 最前面两个
			res = max(res, dfs(i+2, j)+1)
		}
		if a[j-1]+a[j] == target { // 最后两个
			res = max(res, dfs(i, j-2)+1)
		}
		if a[i]+a[j] == target { // 第一个和最后一个
			res = max(res, dfs(i+1, j-1)+1)
		}
		*p = res
		return
	}
	return dfs(0, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n^2)$。
