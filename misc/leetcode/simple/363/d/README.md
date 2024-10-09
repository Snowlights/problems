#### 题目

<p>给你一个下标从 <strong>1</strong> 开始、由 <code>n</code> 个整数组成的数组。</p>

<p>如果一组数字中每对元素的乘积都是一个完全平方数，则称这组数字是一个 <strong>完全集</strong> 。</p>

<p>下标集 <code>{1, 2, ..., n}</code> 的子集可以表示为 <code>{i<sub>1</sub>, i<sub>2</sub>, ..., i<sub>k</sub>}</code>，我们定义对应该子集的 <strong>元素和</strong> 为 <code>nums[i<sub>1</sub>] + nums[i<sub>2</sub>] + ... + nums[i<sub>k</sub>]</code> 。</p>

<p>返回下标集 <code>{1, 2, ..., n}</code> 的 <strong>完全子集</strong> 所能取到的 <strong>最大元素和</strong> 。</p>

<p>完全平方数是指可以表示为一个整数和其自身相乘的数。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [8,7,3,5,7,2,4,9]
<strong>输出：</strong>16
<strong>解释：</strong>除了由单个下标组成的子集之外，还有两个下标集的完全子集：{1,4} 和 {2,8} 。
与下标 1 和 4 对应的元素和等于 nums[1] + nums[4] = 8 + 5 = 13 。
与下标 2 和 8 对应的元素和等于 nums[2] + nums[8] = 7 + 9 = 16 。
因此，下标集的完全子集可以取到的最大元素和为 16 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [5,10,3,10,1,13,7,9,4]
<strong>输出：</strong>19
<strong>解释：</strong>除了由单个下标组成的子集之外，还有四个下标集的完全子集：{1,4}、{1,9}、{2,8}、{4,9} 和 {1,4,9} 。 
与下标 1 和 4 对应的元素和等于 nums[1] + nums[4] = 5 + 10 = 15 。 
与下标 1 和 9 对应的元素和等于 nums[1] + nums[9] = 5 + 4 = 9 。 
与下标 2 和 8 对应的元素和等于 nums[2] + nums[8] = 10 + 9 = 19 。
与下标 4 和 9 对应的元素和等于 nums[4] + nums[9] = 10 + 4 = 14 。 
与下标 1、4 和 9 对应的元素和等于 nums[1] + nums[4] + nums[9] = 5 + 10 + 4 = 19 。 
因此，下标集的完全子集可以取到的最大元素和为 19 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n == nums.length <= 10<sup>4</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

## 方法一：

按照下标的 core 值分组
定义 $\text{core}(n)$ 为 $n$ 除去完全平方因子后的剩余结果。
例如 $\text{core}(8)=8/4=2,\ \text{core}(12)=12/4=3, \text{core}(25)=25/25=1, \text{core}(5)=5/1=5$。
计算方式同质因数分解，把 $n$ 的所有出现次数为奇数的质因子相乘，即为 $\text{core}(n)$。
根据题意，如果同一组中有两个数，它们的下标的 $\text{core}$ 值不同，那么这两个数相乘，就不是一个完全平方数。
所以，同一组内的数字下标的 $\text{core}$ 值必须都一样。
那么按照下标的 $\text{core}$ 值分组，累加同一组的元素和，最大元素和即为答案。

```go  
func core(n int) int {
	res := 1
	for i := 2; i*i <= n; i++ {
		e := 0
		for n%i == 0 {
			e ^= 1
			n /= i
		}
		if e == 1 {
			res *= i
		}
	}
	if n > 1 {
		res *= n
	}
	return res
}

func maximumSum(nums []int) (ans int64) {
	sum := make([]int64, len(nums)+1)
	for i, x := range nums {
		c := core(i + 1)
		sum[c] += int64(x)
		ans = max(ans, sum[c])
	}
	return
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\sqrt{n})$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：枚举 core

我们还可以从 $1$ 开始枚举 $i$ 和所有 $\text{core}$ 值等于 $i$ 的下标，也就是 $i\cdots j^2$。
例如 $i=3$ 时，我们可以枚举

$$
3\cdot 1, 3\cdot 4, 3\cdot 9, 3\cdot 16,3\cdot 25,cdots

$$

这些下标对应的元素值都在同一组中。
可能你会觉得这不对，比如 $i=4$ 时，我们枚举的是

$$
4\cdot 1, 4\cdot 4, 4\cdot 9, 4\cdot 16,4\cdot 25,cdots

$$

而这些数的 $\text{core}$ 值都等于 $1$，而不是 $4$。
没有关系！在 $i=1$ 时，上面这些数都枚举到了（相当于是 $i=1$ 时枚举的数的子集），所以上面这些数的元素和不会超过 $i=1$ 时计算出来的元素和。**我们不会漏掉最大的元素和。**

```go  
func maximumSum(nums []int) (ans int64) {
	n := len(nums)
	for i := 1; i <= n; i++ {
		sum := int64(0)
		for j := 1; i*j*j <= n; j++ {
			sum += int64(nums[i*j*j-1]) // -1 是因为数组下标从 0 开始
		}
		ans = max(ans, sum)
	}
	return
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。循环次数约为 $\sqrt{n}\cdot(1/\sqrt{1} + 1/\sqrt{2} + ...)$，由 $f(x)=1/\sqrt{x}$ 的积分可知，$1/\sqrt{1} + 1/\sqrt{2} + ... = \mathcal{O}(\sqrt{n})$，所以总的循环次数为 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。
