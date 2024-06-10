#### 题目

<p>给你一个下标从 <strong>0</strong> 开始、长度为 <code>n</code> 的整数数组 <code>nums</code> ，和一个整数 <code>k</code> 。</p>

<p>你可以执行下述 <strong>递增</strong> 运算 <strong>任意</strong> 次（可以是 <strong>0</strong> 次）：</p>

<ul>
	<li>从范围 <code>[0, n - 1]</code> 中选则一个下标 <code>i</code> ，并将 <code>nums[i]</code> 的值加 <code>1</code> 。</li>
</ul>

<p>如果数组中任何长度 <strong>大于或等于 3</strong> 的子数组，其 <strong>最大</strong> 元素都大于或等于 <code>k</code> ，则认为数组是一个 <strong>美丽数组</strong> 。</p>

<p>以整数形式返回使数组变为 <strong>美丽数组</strong> 需要执行的 <strong>最小</strong> 递增运算数。</p>

<p>子数组是数组中的一个连续 <strong>非空</strong> 元素序列。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2,3,0,0,2], k = 4
<strong>输出：</strong>3
<strong>解释：</strong>可以执行下述递增运算，使 nums 变为美丽数组：
选择下标 i = 1 ，并且将 nums[1] 的值加 1 -> [2,4,0,0,2] 。
选择下标 i = 4 ，并且将 nums[4] 的值加 1 -> [2,4,0,0,3] 。
选择下标 i = 4 ，并且将 nums[4] 的值加 1 -> [2,4,0,0,4] 。
长度大于或等于 3 的子数组为 [2,4,0], [4,0,0], [0,0,4], [2,4,0,0], [4,0,0,4], [2,4,0,0,4] 。
在所有子数组中，最大元素都等于 k = 4 ，所以 nums 现在是美丽数组。
可以证明无法用少于 3 次递增运算使 nums 变为美丽数组。
因此，答案为 3 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [0,1,3,3], k = 5
<strong>输出：</strong>2
<strong>解释：</strong>可以执行下述递增运算，使 nums 变为美丽数组：
选择下标 i = 2 ，并且将 nums[2] 的值加 1 -> [0,1,4,3] 。
选择下标 i = 2 ，并且将 nums[2] 的值加 1 -> [0,1,5,3] 。
长度大于或等于 3 的子数组为 [0,1,5]、[1,5,3]、[0,1,5,3] 。
在所有子数组中，最大元素都等于 k = 5 ，所以 nums 现在是美丽数组。
可以证明无法用少于 2 次递增运算使 nums 变为美丽数组。 
因此，答案为 2 。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,1,2], k = 1
<strong>输出：</strong>0
<strong>解释：</strong>在这个示例中，只有一个长度大于或等于 3 的子数组 [1,1,2] 。
其最大元素 2 已经大于 k = 1 ，所以无需执行任何增量运算。
因此，答案为 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 <= n == nums.length <= 10<sup>5</sup></code></li>
	<li><code>0 <= nums[i] <= 10<sup>9</sup></code></li>
	<li><code>0 <= k <= 10<sup>9</sup></code></li>
</ul>

#### 思路

## 写法一：记忆化搜索

把大于 $k$ 的元素视作 $k$。
由于大于 $3$ 的子数组必然包含等于 $3$ 的子数组，问题转换成：每个长为 $3$ 的子数组都需要包含至少一个 $k$。考虑最后一个元素「选或不选」，即是否增大：

- 增大到 $k$：那么对于左边那个数来说，它右边就有一个 $k$ 了。
- 不增大：那么对于左边那个数来说，它右边有一个没有增大的数。

进一步地，如果倒数第二个数也不增大，那么对于倒数第三个数，它右边就有两个没有增大的数，那么它一定要增大（不用管右边那两个数是否为 $k$，因为下面的 $\textit{dfs}$ 会考虑到所有的情况，取最小值）。因此，用 $i$ 表示当前位置，$j$ 表示右边有几个没有增大的数。定义 $\textit{dfs}(i,j)$ 表示在该状态下对于前 $i$ 个数的最小递增运算数。

- 增大到 $k$，即 $\textit{dfs}(i-1,0) + \max(k-\textit{nums}[i], 0)$。
- 如果 $j<2$，则可以不增大，即 $\textit{dfs}(i-1,j+1)$。

这两种情况取最小值。
递归边界：当 $i<0$ 时返回 $0$。递归入口：$\textit{dfs}(n-1,0)$，即答案。

```go  
func minIncrementOperations(nums []int, k int) int64 {
	n := len(nums)
	memo := make([][3]int, n)
	for i := range memo {
		memo[i] = [3]int{-1, -1, -1}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		res := dfs(i-1, 0) + max(k-nums[i], 0) // nums[i] 增大
		if j < 2 {
			res = min(res, dfs(i-1, j+1)) // nums[i] 不增大
		}
		*p = res
		return res
	}
	return int64(dfs(n-1, 0))
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func minIncrementOperations(nums []int, k int) int64 {
	var f0, f1, f2 int
	for _, x := range nums {
		inc := f0 + max(k-x, 0)
		f0 = min(inc, f1)
		f1 = min(inc, f2)
		f2 = inc
	}
	return int64(f0)
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。
