#### 题目  

<p>给你一个长度为 <code>n</code> 的整数数组 <code>nums</code> ，下标从 <strong>0</strong> 开始。</p>

<p>如果在下标 <code>i</code> 处 <strong>分割</strong> 数组，其中 <code>0 &lt;= i &lt;= n - 2</code> ，使前 <code>i + 1</code> 个元素的乘积和剩余元素的乘积互质，则认为该分割 <strong>有效</strong> 。</p>

<ul>
	<li>例如，如果 <code>nums = [2, 3, 3]</code> ，那么在下标 <code>i = 0</code> 处的分割有效，因为 <code>2</code> 和 <code>9</code> 互质，而在下标 <code>i = 1</code> 处的分割无效，因为 <code>6</code> 和 <code>3</code> 不互质。在下标 <code>i = 2</code> 处的分割也无效，因为 <code>i == n - 1</code> 。</li>
</ul>

<p>返回可以有效分割数组的最小下标 <code>i</code> ，如果不存在有效分割，则返回 <code>-1</code> 。</p>

<p>当且仅当 <code>gcd(val1, val2) == 1</code> 成立时，<code>val1</code> 和 <code>val2</code> 这两个值才是互质的，其中 <code>gcd(val1, val2)</code> 表示 <code>val1</code> 和 <code>val2</code> 的最大公约数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2022/12/14/second.PNG" style="width: 450px; height: 211px;"/></p>

<pre><strong>输入：</strong>nums = [4,7,8,15,3,5]
<strong>输出：</strong>2
<strong>解释：</strong>上表展示了每个下标 i 处的前 i + 1 个元素的乘积、剩余元素的乘积和它们的最大公约数的值。
唯一一个有效分割位于下标 2 。</pre>

<p><strong>示例 2：</strong></p>

<p><img alt="" src="https://assets.leetcode.com/uploads/2022/12/14/capture.PNG" style="width: 450px; height: 215px;"/></p>

<pre><strong>输入：</strong>nums = [4,7,15,8,3,5]
<strong>输出：</strong>-1
<strong>解释：</strong>上表展示了每个下标 i 处的前 i + 1 个元素的乘积、剩余元素的乘积和它们的最大公约数的值。
不存在有效分割。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == nums.length</code></li>
	<li><code>1 &lt;= n &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
</ul>
 
#### 思路  

对于一个质因子 $p$，设它在数组中的最左和最右的位置为 $\textit{left}$ 和 $\textit{right}$。  
那么答案是不能在区间 $[\textit{left},\textit{right})$ 中的。注意区间右端点可能为答案。  
因此这题本质上和 [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/) 是类似的，找从 $0$ 出发，最远遇到的区间右端点，即为答案。

```go 
func findValidSplit(nums []int) int {
	left := map[int]int{}           // left[p] 表示质数 p 首次出现的下标
	right := make([]int, len(nums)) // right[i] 表示左端点为 i 的区间的右端点的最大值
	f := func(p, i int) {
		if l, ok := left[p]; ok {
			right[l] = i // 记录左端点 l 对应的右端点的最大值
		} else {
			left[p] = i // 第一次遇到质数 p
		}
	}

	for i, x := range nums {
		for d := 2; d*d <= x; d++ { // 分解质因数
			if x%d == 0 {
				f(d, i)
				for x /= d; x%d == 0; x /= d {
				}
			}
		}
		if x > 1 {
			f(x, i)
		}
	}

	maxR := 0
	for l, r := range right {
		if l > maxR { // 最远可以遇到 maxR
			return maxR // 也可以写 l-1
		}
		maxR = max(maxR, r)
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

#### 复杂度分析  

- 时间复杂度：$O(n\sqrt U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$O\left(n + \dfrac{U}{\log U}\right)$。$U$ 范围内的质数个数有 $O\left(\dfrac{U}{\log U}\right)$ 个。