#### 题目

<p>给你一个下标从 <strong>0</strong> 开始的整数数组 <code>nums</code> 和一个 <strong>正</strong> 整数 <code>k</code> 。</p>

<p>你可以对数组执行以下操作 <strong>任意次</strong> ：</p>

<ul>
	<li>选择两个互不相同的下标 <code>i</code> 和 <code>j</code> ，<strong>同时</strong> 将 <code>nums[i]</code> 更新为 <code>(nums[i] AND nums[j])</code> 且将 <code>nums[j]</code> 更新为 <code>(nums[i] OR nums[j])</code> ，<code>OR</code> 表示按位 <strong>或</strong> 运算，<code>AND</code> 表示按位 <strong>与</strong> 运算。</li>
</ul>

<p>你需要从最终的数组里选择 <code>k</code> 个元素，并计算它们的 <strong>平方</strong> 之和。</p>

<p>请你返回你可以得到的 <strong>最大</strong> 平方和。</p>

<p>由于答案可能会很大，将答案对 <code>10<sup>9</sup> + 7</code> <strong>取余</strong> 后返回。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><b>输入：</b>nums = [2,6,5,8], k = 2
<b>输出：</b>261
<b>解释：</b>我们可以对数组执行以下操作：
- 选择 i = 0 和 j = 3 ，同时将 nums[0] 变为 (2 AND 8) = 0 且 nums[3] 变为 (2 OR 8) = 10 ，结果数组为 nums = [0,6,5,10] 。
- 选择 i = 2 和 j = 3 ，同时将 nums[2] 变为 (5 AND 10) = 0 且 nums[3] 变为 (5 OR 10) = 15 ，结果数组为 nums = [0,6,0,15] 。
从最终数组里选择元素 15 和 6 ，平方和为 15<sup>2</sup> + 6<sup>2</sup> = 261 。
261 是可以得到的最大结果。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>nums = [4,5,4,7], k = 3
<b>输出：</b>90
<b>解释：</b>不需要执行任何操作。
选择元素 7 ，5 和 4 ，平方和为 7<sup>2</sup> + 5<sup>2</sup> + 4<sup>2</sup> = 90 。
90 是可以得到的最大结果。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= k <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= nums[i] <= 10<sup>9</sup></code></li>
</ul>

#### 思路

## 提示 1

对于同一个比特位，由于 AND 和 OR 不会改变都为 $0$ 和都为 $1$ 的情况，所以操作等价于：
把一个数的 $0$ 和另一个数的同一个比特位上的 $1$ **交换**。

## 提示 2

设交换前两个数是 $x$ 和 $y$，且 $x>y$。把小的数上的 $1$ 给大的数，假设交换后 $x$ 增加了 $d$，那么 $y$ 也减少了 $d$。
交换前：$x^2+y^2$。
交换后：$(x+d)^2+(y-d)^2 = x^2+y^2+2d(x-y)+2d^2 > x^2+y^2$。
这说明应该通过交换，让一个数越大越好。
相当于把 $1$ 都**聚集**在一个数中，比分散到不同的数更好。

## 提示 3

由于可以操作任意次，那么一定可以「组装」出尽量大的数，做法如下：

1. 对于每个比特位，统计 $\textit{nums}$ 在这个比特位上有多少个 $1$，记到一个长（至多）为 $30$ 的 $\textit{cnt}$ 数组中（因为 $10^9 < 2^{30}$）。
2. 循环 $k$ 次。
3. 每次循环，「组装」一个数（记作 $x$）：遍历 $\textit{cnt}$，只要 $\textit{cnt}[i]>0$ 就将其减一，同时将 $2^i$ 加到 $x$ 中。这样相当于把 $1$ 尽量聚集在一个数中。
4. 把 $x^2$ 加到答案中。

```go  
const mod int = 1e9 + 7

func maxSum(a []int, k int) (ans int) {
	cnt := [32]int{}
	for _, v := range a {
		for i := range cnt {
			cnt[i] += v >> i & 1
		}
	}
	for ; k > 0; k-- {
		v := 0
		for i := range cnt {
			if cnt[i] > 0 {
				v |= 1 << i
				cnt[i]--
			}
		}
		ans = (ans + v * v) % mod
	}

	ans = (ans%mod + mod) % mod
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}(\log U)$。
