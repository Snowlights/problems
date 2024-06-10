#### 题目  

<p>给你两个正整数 <code>low</code> 和 <code>high</code> 。</p>

<p>对于一个由 <code>2 * n</code> 位数字组成的整数 <code>x</code> ，如果其前 <code>n</code> 位数字之和与后 <code>n</code> 位数字之和相等，则认为这个数字是一个对称整数。</p>

<p>返回在 <code>[low, high]</code> 范围内的 <strong>对称整数的数目</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><strong>输入：</strong>low = 1, high = 100
<strong>输出：</strong>9
<strong>解释：</strong>在 1 到 100 范围内共有 9 个对称整数：11、22、33、44、55、66、77、88 和 99 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><strong>输入：</strong>low = 1200, high = 1230
<strong>输出：</strong>4
<strong>解释：</strong>在 1200 到 1230 范围内共有 4 个对称整数：1203、1212、1221 和 1230 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= low &lt;= high &lt;= 10<sup>4</sup></code></li>
</ul>
 
#### 思路  

枚举即可

```go 
func countSymmetricIntegers(low int, high int) (ans int) {
	for i := low; i <= high; i++ {
		s := strconv.Itoa(i)
		n := len(s)
		if n&1 == 1 {
			continue
		}
		a, b := 0, 0
		for _, v := range s[:n/2] {
			a += int(v - '0')
		}
		for _, v := range s[n/2:] {
			b += int(v - '0')
		}
		if a == b {
			ans++
		}
	}

	return
}
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}((\textit{high} - \textit{low})\log \textit{high})$。
- 空间复杂度：$\mathcal{O}(\log \textit{high})$。