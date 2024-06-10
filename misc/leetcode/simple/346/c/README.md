#### 题目

<p>给你一个正整数&nbsp;<code>n</code>&nbsp;，请你返回&nbsp;<code>n</code>&nbsp;的&nbsp;<strong>惩罚数</strong>&nbsp;。</p>

<p><code>n</code>&nbsp;的 <strong>惩罚数</strong>&nbsp;定义为所有满足以下条件 <code>i</code>&nbsp;的数的平方和：</p>

<ul>
	<li><code>1 &lt;= i &lt;= n</code></li>
	<li><code>i * i</code> 的十进制表示的字符串可以分割成若干连续子字符串，且这些子字符串对应的整数值之和等于 <code>i</code> 。</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<b>输入：</b>n = 10
<b>输出：</b>182
<b>解释：</b>总共有 3 个整数 i 满足要求：
- 1 ，因为 1 * 1 = 1
- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
因此，10 的惩罚数为 1 + 81 + 100 = 182
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<b>输入：</b>n = 37
<b>输出：</b>1478
<b>解释：</b>总共有 4 个整数 i 满足要求：
- 1 ，因为 1 * 1 = 1
- 9 ，因为 9 * 9 = 81 ，且 81 可以分割成 8 + 1 。
- 10 ，因为 10 * 10 = 100 ，且 100 可以分割成 10 + 0 。
- 36 ，因为 36 * 36 = 1296 ，且 1296 可以分割成 1 + 29 + 6 。
因此，37 的惩罚数为 1 + 81 + 100 + 1296 = 1478
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 1000</code></li>
</ul>

#### 思路

判断 $[1,1000]$ 的每个数字 $i$ 是否为惩罚数，并预处理 $[1,i]$ 内的惩罚数的元素和 $\textit{preSum}$。  
对于每个数字 $i$，把它转成字符串 $s$ 后，写一个回溯，枚举第一个子串、第二个子串、……，
累加所有子串对应的整数值之和 $\textit{sum}$。如果存在 $\textit{sum}=i$，则说明 $i$ 是惩罚数。

```go 
func punishmentNumber(n int) (ans int) {

	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n {
				return sum == i
			}
			x := 0
			for j := p; j < n; j++ {
				x = x*10 + int(s[j]-'0')
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		if dfs(0, 0) {
			ans += i * i
		}
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：预处理 $\mathcal{O}(U^{1 + 2\log_{10} 2})\approx\mathcal{O}(U^{1.602})$，其中 $U=1000$。对于数字 $i^2$，它转成字符串后的长度为 $m=\lfloor1+2\log_{10} i\rfloor$，所以回溯需要 $\mathcal{O}(2^m)=\mathcal{O}(i^{2\log_{10} 2})$ 的时间，对其积分可知，整个预处理需要 $\mathcal{O}(U^{1 + 2\log_{10} 2})$ 的时间。
- 空间复杂度：预处理 $\mathcal{O}(U)$。