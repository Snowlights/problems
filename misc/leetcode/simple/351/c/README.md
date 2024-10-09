#### 题目  

<p>给你一个二元数组 <code>nums</code> 。</p>

<p>如果数组中的某个子数组 <strong>恰好</strong> 只存在 <strong>一</strong> 个值为 <code>1</code> 的元素，则认为该子数组是一个 <strong>好子数组</strong> 。</p>

<p>请你统计将数组 <code>nums</code> 划分成若干 <strong>好子数组</strong> 的方法数，并以整数形式返回。由于数字可能很大，返回其对 <code>10<sup>9</sup> + 7</code> <strong>取余 </strong>之后的结果。</p>

<p>子数组是数组中的一个连续 <strong>非空</strong> 元素序列。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>nums = [0,1,0,0,1]
<strong>输出：</strong>3
<strong>解释：</strong>存在 3 种可以将 nums 划分成若干好子数组的方式：
- [0,1] [0,0,1]
- [0,1,0] [0,1]
- [0,1,0,0] [1]
</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>nums = [0,1,0]
<strong>输出：</strong>1
<strong>解释：</strong>存在 1 种可以将 nums 划分成若干好子数组的方式：
- [0,1,0]
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 1</code></li>
</ul>
 
#### 思路  

根据题意，需要在**每**两个 $1$ 之间画一条分割线，有 $x$ 个 $0$ 就可以画 $x+1$ 条分割线。
根据乘法原理，答案为所有分割线的方案数的乘积。
特别地，如果数组中没有 $1$，那么答案为 $0$。如果数组只有一个 $1$，那么答案为 $1$。

```go 
const mod int = 1e9 + 7
func numberOfGoodSubarraySplits(a []int) (ans int) {

	ans, pre := 1, -1
	for i, v := range a {
		if v == 1 {
			if pre != -1 {
				ans = ans * (i - pre) % mod
			}
			pre = i
		}
	}
	if pre == -1 {
		return 0
	}
	ans = (ans%mod + mod) % mod
	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。