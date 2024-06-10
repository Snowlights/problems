#### 题目  

<p>给你两个数字字符串 <code>num1</code> 和 <code>num2</code> ，以及两个整数 <code>max_sum</code> 和 <code>min_sum</code> 。如果一个整数 <code>x</code> 满足以下条件，我们称它是一个好整数：</p>

<ul>
	<li><code>num1 &lt;= x &lt;= num2</code></li>
	<li><code>min_sum &lt;= digit_sum(x) &lt;= max_sum</code>.</li>
</ul>

<p>请你返回好整数的数目。答案可能很大，请返回答案对 <code>10<sup>9</sup> + 7</code> 取余后的结果。</p>

<p>注意，<code>digit_sum(x)</code> 表示 <code>x</code> 各位数字之和。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>num1 = &#34;1&#34;, num2 = &#34;12&#34;, min_num = 1, max_num = 8
<b>输出：</b>11
<b>解释：</b>总共有 11 个整数的数位和在 1 到 8 之间，分别是 1,2,3,4,5,6,7,8,10,11 和 12 。所以我们返回 11 。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>num1 = &#34;1&#34;, num2 = &#34;5&#34;, min_num = 1, max_num = 5
<b>输出：</b>5
<b>解释：</b>数位和在 1 到 5 之间的 5 个整数分别为 1,2,3,4 和 5 。所以我们返回 5 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= num1 &lt;= num2 &lt;= 10<sup>22</sup></code></li>
	<li><code>1 &lt;= min_sum &lt;= max_sum &lt;= 400</code></li>
</ul>
 
#### 思路  

数位dp

```go 
const mod int = 1e9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func count(x string, y string, min_sum int, max_sum int) (ans int) {

	var f func(s string) int
	f = func(s string) int {
		n := len(s)
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, min(9*n, max_sum)+1)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}

		var dfs func(i, sum int, limit bool) int
		dfs = func(i, sum int, limit bool) int {
			if sum > max_sum {
				return 0
			}
			if i == n {
				if sum >= min_sum {
					return 1
				}
				return 0
			}
			res := 0
			if !limit {
				dv := &dp[i][sum]
				if *dv >= 0 {
					return *dv
				}
				defer func() {
					*dv = res
				}()
			}

			down, up := 0, 9
			if limit {
				up = int(s[i] - '0')
			}
			for ; down <= up; down++ {
				res += dfs(i+1, sum+down, limit && down == up) % mod
			}
			return res
		}
		return dfs(0, 0, true)
	}

	ans = f(y) - f(x) + mod
	sum := 0
	for _, v := range x {
		sum += int(v - '0')
	}
	if min_sum <= sum && sum <= max_sum {
		ans++
	}

	ans = (ans%mod + mod) % mod
	return
}

```

```python
from functools import cache
class Solution:
    def count(self, num1, num2: str, min_sum, max_sum: int) -> int:
        mod = int(1e9+7)

        def f(s:str) -> int:
            n = len(s)
            @cache
            def dfs(i, sum : int, limit: bool) -> int:
                if sum > max_sum:
                    return 0
                if i == n:
                    if sum >= min_sum:
                        return 1
                    else:
                        return 0
                res, up = 0, 9
                if limit:
                    up = int(s[i])
                for down in range(up+1):
                    res = (res + dfs(i+1, sum + down, limit and down == up) ) % mod
                return res
            return dfs(0, 0, True) % mod
        return (f(num2) - f(num1) + mod + (min_sum <= sum(map(int, num1)) <= max_sum)) % mod
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(10nm)$，其中 $n$ 为 $\textit{nums}_2$ 的长度，$m=\min\{9n, \textit{maxSum}\}$。动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题中状态个数等于 $\mathcal{O}(nm)$，单个状态的计算时间为 $\mathcal{O}(10)$，因此时间复杂度为 $\mathcal{O}(10nm)$。
- 空间复杂度：$\mathcal{O}(nm)$。