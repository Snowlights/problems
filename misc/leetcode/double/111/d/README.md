### 题目

<p>给你正整数 <code>low</code> ，<code>high</code> 和 <code>k</code> 。</p>

<p>如果一个数满足以下两个条件，那么它是 <strong>美丽的</strong> ：</p>

<ul>
	<li>偶数数位的数目与奇数数位的数目相同。</li>
	<li>这个整数可以被 <code>k</code> 整除。</li>
</ul>

<p>请你返回范围 <code>[low, high]</code> 中美丽整数的数目。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>low = 10, high = 20, k = 3
<b>输出：</b>2
<b>解释：</b>给定范围中有 2 个美丽数字：[12,18]
- 12 是美丽整数，因为它有 1 个奇数数位和 1 个偶数数位，而且可以被 k = 3 整除。
- 18 是美丽整数，因为它有 1 个奇数数位和 1 个偶数数位，而且可以被 k = 3 整除。
以下是一些不是美丽整数的例子：
- 16 不是美丽整数，因为它不能被 k = 3 整除。
- 15 不是美丽整数，因为它的奇数数位和偶数数位的数目不相等。
给定范围内总共有 2 个美丽整数。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>low = 1, high = 10, k = 1
<b>输出：</b>1
<b>解释：</b>给定范围中有 1 个美丽数字：[10]
- 10 是美丽整数，因为它有 1 个奇数数位和 1 个偶数数位，而且可以被 k = 1 整除。
给定范围内总共有 1 个美丽整数。
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>low = 5, high = 5, k = 2
<b>输出：</b>0
<b>解释：</b>给定范围中有 0 个美丽数字。
- 5 不是美丽整数，因为它的奇数数位和偶数数位的数目不相等。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 < low <= high <= 10<sup>9</sup></code></li>
	<li><code>0 < k <= 20</code></li>
</ul>

### 思路

数位dp

```go  
func numberOfBeautifulIntegers(low int, high int, k int) int {

	cal := func(s string) int{
		n := len(s)
		var dp = [10][10][10][21]int{}
		for i := range dp {
			for j := range dp[i] {
				for k := range dp[i][j] {
					for l := range dp[i][j][k] {
						dp[i][j][k][l] = -1
					}
				}
			}
		}
		var dfs func(i, pre, odd, even, mod int, limit, num bool) int
		dfs = func(i, pre, odd, even, mod int, limit, num bool) int {
			if i == n {
				if odd == even && mod == 0 {
					return 1
				}
				return 0
			}

			res := 0
			if num && !limit {
				dv := &dp[i][odd][even][mod]
				if *dv >= 0 {
					return *dv
				}
				defer func() {
					*dv = res
				}()
			}

			if !num {
				res += dfs(i+1, 0, 0, 0, 0, false, false)
			}

			low, up := 0, 9
			if !num {
				low = 1
			}
			if limit {
				up = int(s[i] - '0')
			}

			for ; low <= up; low ++ {
				res += dfs(i+1, pre*10 + low, odd + low % 2, even + (low+1)%2, (pre*10 + low)%k, limit && low == up, true)
			}
			return res
		}
		return dfs(0, 0, 0, 0, 0, true, false)
	}
	return cal(strconv.Itoa(high)) - cal(strconv.Itoa(low-1))
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2kD)$，其中 $n=\mathcal{O}(\log\textit{high})$，也就是 $\textit{high}$ 的十进制表示的长度；$D=10$。由于每个状态只会计算一次，动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(n^2k)$，单个状态的计算时间为 $\mathcal{O}(D)$，所以动态规划的时间复杂度为 $\mathcal{O}(n^2kD)$
- 空间复杂度：$\mathcal{O}(nm)$。
