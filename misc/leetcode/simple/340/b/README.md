#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 。现有一个长度等于 <code>nums.length</code> 的数组 <code>arr</code> 。对于满足 <code>nums[j] == nums[i]</code> 且 <code>j != i</code> 的所有 <code>j</code> ，<code>arr[i]</code> 等于所有 <code>|i - j|</code> 之和。如果不存在这样的 <code>j</code> ，则令 <code>arr[i]</code> 等于 <code>0</code> 。</p>

<p>返回数组<em> </em><code>arr</code><em> 。</em></p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [1,3,1,1,2]
<strong>输出：</strong>[5,0,3,4,0]
<strong>解释：</strong>
i = 0 ，nums[0] == nums[2] 且 nums[0] == nums[3] 。因此，arr[0] = |0 - 2| + |0 - 3| = 5 。 
i = 1 ，arr[1] = 0 因为不存在值等于 3 的其他下标。
i = 2 ，nums[2] == nums[0] 且 nums[2] == nums[3] 。因此，arr[2] = |2 - 0| + |2 - 3| = 3 。
i = 3 ，nums[3] == nums[0] 且 nums[3] == nums[2] 。因此，arr[3] = |3 - 0| + |3 - 2| = 4 。 
i = 4 ，arr[4] = 0 因为不存在值等于 2 的其他下标。
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [0,5,3]
<strong>输出：</strong>[0,0,0]
<strong>解释：</strong>因为 nums 中的元素互不相同，对于所有 i ，都有 arr[i] = 0 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
</ul>
 
#### 思路  

## 方法一：相同元素分组+前缀和

按照相同元素分组后再计算。
这里的思路和 [2602. 使数组元素全部相等的最少操作次数（题解）](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/solution/yi-tu-miao-dong-pai-xu-qian-zhui-he-er-f-nf55/)
是一样的。由于目标位置就是数组中的下标，无需二分。

```go 
func distance(nums []int) []int64 {
	groups := map[int][]int{}
	for i, x := range nums {
		groups[x] = append(groups[x], i) // 相同元素分到同一组，记录下标
	}
	ans := make([]int64, len(nums))
	for _, a := range groups {
		n := len(a)
		s := make([]int, n+1)
		for i, x := range a {
			s[i+1] = s[i] + x // 前缀和
		}
		for i, target := range a {
			left := target*i - s[i] // 蓝色面积
			right := s[n] - s[i] - target*(n-i) // 绿色面积
			ans[target] = int64(left + right)
		}
	}
	return ans
}
```

#### 复杂度分析  

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。

## 方法二：相同元素分组+考虑增量

分组后，对于其中一个组 $a$，我们先暴力计算出 $a[0]$ 到其它元素的距离之和，设为 $s$。  
然后计算 $a[1]$，我们不再暴力计算，而是思考：$s$ 增加了多少？（增量可以是负数）  
设 $n$ 为 $a$ 的长度，从 $a[0]$ 到 $a[1]$，有 $1$ 个数的距离变大了 $a[1]-a[0]$，有 $n-1$ 个数的距离变小了 $a[1]-a[0]$。  
所以对于 $a[1]$，我们可以 $O(1)$ 地知道它到其它元素的距离之和为

$$
s + (2-n) \cdot (a[1]-a[0])
$$

一般地，设 $a[i-1]$ 到其它元素的距离之和为 $s$，那么 $a[i]$ 到其它元素的距离之和为

$$
s + (2i-n) \cdot (a[i]-a[i-1])
$$

```go  
func distance(nums []int) []int64 {
	groups := map[int][]int{}
	for i, x := range nums {
		groups[x] = append(groups[x], i) // 相同元素分到同一组，记录下标
	}
	ans := make([]int64, len(nums))
	for _, a := range groups {
		n := len(a)
		s := int64(0)
		for _, x := range a {
			s += int64(x - a[0]) // a[0] 到其它下标的距离之和
		}
		ans[a[0]] = s
		for i := 1; i < n; i++ {
			// 从计算 a[i-1] 到计算 a[i]，考虑 s 增加了多少
			s += int64(i*2-n) * int64(a[i]-a[i-1])
			ans[a[i]] = s
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。