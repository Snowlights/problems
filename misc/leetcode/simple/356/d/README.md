#### 题目

<p>给你两个正整数 <code>low</code> 和 <code>high</code> ，都用字符串表示，请你统计闭区间 <code>[low, high]</code> 内的 <strong>步进数字</strong> 数目。</p>

<p>如果一个整数相邻数位之间差的绝对值都 <strong>恰好</strong> 是 <code>1</code> ，那么这个数字被称为 <strong>步进数字</strong> 。</p>

<p>请你返回一个整数，表示闭区间 <code>[low, high]</code> 之间步进数字的数目。</p>

<p>由于答案可能很大，请你将它对 <code>10<sup>9</sup> + 7</code> <strong>取余</strong> 后返回。</p>

<p><b>注意：</b>步进数字不能有前导 0 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>low = "1", high = "11"
<b>输出：</b>10
<strong>解释：</strong>区间 [1,11] 内的步进数字为 1 ，2 ，3 ，4 ，5 ，6 ，7 ，8 ，9 和 10 。总共有 10 个步进数字。所以输出为 10 。</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>low = "90", high = "101"
<b>输出：</b>2
<strong>解释：</strong>区间 [90,101] 内的步进数字为 98 和 101 。总共有 2 个步进数字。所以输出为 2 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= int(low) <= int(high) < 10<sup>100</sup></code></li>
	<li><code>1 <= low.length, high.length <= 100</code></li>
	<li><code>low</code> 和 <code>high</code> 只包含数字。</li>
	<li><code>low</code> 和 <code>high</code> 都不含前导 0 。</li>
</ul>

#### 思路

定义 $\text{calc}(n)$ 表示不超过 $n$ 的步进数字数目。那么答案就是

$$
\begin{aligned}&\text{calc}(\textit{high}) - \text{calc}(\textit{low}-1)\
=\ &\text{calc}(\textit{high}) - \text{calc}(\textit{low}) + \text{valid}(\textit{low})
\end{aligned}

$$

由于 $\textit{low}$ 是个很大的数字，不方便减一（Python 用户可以无视），所以用 $\text{valid}(\textit{low})$ 表示： 如果 $\textit{low}$ 是步进数字，那么多减了 $1$，再加 $1$ 补回来。如何计算 $\text{calc}(n)$ 呢？（下文用 $s$ 表示 $n$ 的字符串形式）一种通用套路是，定义 $f(i,\textit{pre}, \textit{isLimit},\textit{isNum})$ 表示构造第 $i$ 位及其之后数位的合法方案数，其余参数的含义为：

- $\textit{pre}$ 表示上一个数位的值。如果 $\textit{isNum}$ 为 `false`，可以忽略 $\textit{pre}$。
- $\textit{isLimit}$ 表示当前是否受到了 $n$ 的约束（注意要构造的数字不能超过 $n$）。若为真，则第 $i$ 位填入的数字至多为 $s[i]$，否则可以是 $9$。如果在受到约束的情况下填了 $s[i]$，那么后续填入的数字仍会受到 $n$ 的约束。例如 $n=123$，那么 $i=0$ 填的是 $1$ 的话，$i=1$ 的这一位至多填 $2$。
- $\textit{isNum}$ 表示 $i$ 前面的数位是否填了数字。若为假，则当前位可以跳过（不填数字），或者要填入的数字至少为 $1$；若为真，则要填入的数字可以从 $0$ 开始。例如 $n=123$，在 $i=0$ 时跳过的话，相当于后面要构造的是一个 $99$ 以内的数字了，如果 $i=1$ 不跳过，那么相当于构造一个 $10$ 到 $99$ 的两位数，如果 $i=1$ 也跳过，相当于构造的是一个 $9$ 以内的数字。

#### 实现细节

递归入口：`f(0, 0, true, false)`，表示：

- 从 $s[0]$ 开始枚举；
- 一开始集合中没有数字；
- 一开始要受到 $n$ 的约束（否则就可以随意填了，这肯定不行）；
- 一开始没有填数字。

递归中：

- 如果 $\textit{isNum}$ 为假，说明前面没有填数字，那么当前也可以不填数字。一旦从这里递归下去，$\textit{isLimit}$ 就可以置为 `false` 了，这是因为 $s[0]$ 必然是大于 $0$ 的，后面就不受到 $n$ 的约束了。或者说，最高位不填数字，后面无论怎么填都比 $n$ 小。
- 如果 $\textit{isNum}$ 为真，那么当前必须填一个数字。枚举填入的数字，根据 $\textit{isNum}$ 和 $\textit{isLimit}$ 来决定填入数字的范围。

递归终点：当 $i$ 等于 $s$ 长度时，如果 $\textit{isNum}$ 为真，则表示得到了一个合法数字（因为不合法的不会继续递归下去），返回 $1$，否则返回 $0$。
关于取模的细节，见文末的讲解。

#### 答疑

**问**：记忆化四个状态有点麻烦，能不能只记忆化 $i$ 和 $\textit{pre}$ 这两个状态？**答**：可以的！比如 $n=234$，第一位填 $2$，第二位填 $3$，后面无论怎么递归，都不会再次递归到第一位填 $2$，第二位填 $3$ 的情况，所以不需要记录。又比如，第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。根据这个例子，我们可以只记录不受到约束时的状态 $(i,\textit{mask},\text{false},\text{true})$。比如 $n=456$，第一位（最高位）填的 $3$，那么继续递归，后面就可以随便填，所以状态 $(1,3,\text{false},\text{true})$ 就表示 $i=0$ 填 $3$，从 $i=1$ 往后随便填的方案数。由于后面两个参数恒为 $\text{false}$ 和 $\text{true}$，所以可以不用记忆化，只记忆化 $i$ 和 $\textit{pre}$。

> 注：Python 有 `@cache`，可以无视上面说的。

**问**：能不能只记忆化 $i$？
**答**：这是不行的。想一想，我们为什么要用记忆化？如果递归到同一个状态时，计算出的结果是一样的，那么第二次递归到同一个状态，就可以直接返回第一次计算的结果了。通过保存第一次计算的结果，来优化时间复杂度。
由于前面选的数字会影响后面选的数字，两次递归到相同的 $i$，如果前面选的数字不一样，计算出的结果就可能是不一样的。如果只记忆化 $i$，就可能会算出错误的结果。
也可以这样理解：记忆化搜索要求递归函数无副作用（除了修改 `memo` 数组），从而保证递归到同一个状态时，计算出的结果是一样的。

```go
const mod int = 1e9 + 7

func countSteppingNumbers(low string, high string) (ans int) {

	var f func(s string) int
	f = func(s string) int {
		dp := make([][10]int, len(s))
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = -1 // -1 表示没有计算过
			}
		}
		var dfs func(i, pre int, limit, num bool) int
		dfs = func(i, pre int, limit, num bool) int {
			if i == len(s) {
				if num {
					return 1
				}
				return 0
			}
			res := 0
			if !limit && num {
				dv := &dp[i][pre]
				if *dv >= 0 {
					return *dv
				}
				defer func() {
					if !limit && num {
						*dv = res
					}
				}()
			}
			if !num {
				res += dfs(i+1, pre, false, false)
			}

			down, up := 0, 9
			if !num {
				down = 1
			}
			if limit {
				up = int(s[i] - '0')
			}

			for ; down <= up; down++ {
				if !num || abs(down-pre) == 1 {
					res += dfs(i+1, down, limit && down == up, true) % mod
				}
			}
			return res % mod
		}
		return dfs(0, 0, true, false)
	}

	valid := func(s string) int {
		for i := 1; i < len(s); i++ {
			if abs(int(s[i-1])-int(s[i])) != 1 {
				return 0
			}
		}
		return 1
	}

	ans = f(high) - f(low) + valid(low)
	ans = (ans%mod + mod) % mod
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nD^2)$，其中 $n$ 为 $\textit{high}$ 的长度，$D=10$。由于每个状态只会计算一次，因此动态规划的时间复杂度 $=$ 状态个数 $\times$ 单个状态的计算时间。本题状态个数为 $\mathcal{O}(nD)$，单个状态的计算时间为 $\mathcal{O}(D)$，因此时间复杂度为 $\mathcal{O}(nD^2)$。
- 空间复杂度：$\mathcal{O}(nD)$

#### 附：模运算

如果让你计算 $1234\cdot 6789$ 的**个位数**，你会如何计算？
由于只有个位数会影响到乘积的个位数，那么 $4\cdot 9=36$ 的个位数 $6$ 就是答案。
对于 $1234+6789$ 的个位数，同理，$4+9=13$ 的个位数 $3$ 就是答案。
你能把这个结论抽象成数学等式吗？
一般地，涉及到取模的题目，通常会用到如下等式（上面计算的是 $m=10$）：

$$
(a+b)\bmod m = ((a\bmod m) + (b\bmod m)) \bmod m

$$

$$
(a\cdot b) \bmod m=((a\bmod m)\cdot  (b\bmod m)) \bmod m

$$

证明：根据**带余除法**，任意整数 $a$ 都可以表示为 $a=km+r$，这里 $r$ 相当于 $a\bmod m$。那么设 $a=k_1m+r_1,\ b=k_2m+r_2$。
第一个等式：

$$
\begin{aligned}
&\ (a+b) \bmod m\\
=&\ ((k_1+k_2) m+r_1+r_2)\bmod m\\
=&\ (r_1+r_2)\bmod m\\
=&\ ((a\bmod m) + (b\bmod m)) \bmod m \end{aligned}

$$

第二个等式：

$$
\begin{aligned}
&\ (a\cdot b) \bmod m\\
=&\ (k_1k_2m^2+(k_1r_2+k_2r_1)m+r_1r_2)\bmod m\\
=&\ (r_1r_2)\bmod m\\
=&\ ((a\bmod m)\cdot  (b\bmod m)) \bmod m
\end{aligned}

$$

**根据这两个恒等式，可以随意地对代码中的加法和乘法的结果取模**。
