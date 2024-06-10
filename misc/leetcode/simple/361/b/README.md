#### 题目  

<p>给你一个下标从 <strong>0</strong> 开始的字符串 <code>num</code> ，表示一个非负整数。</p>

<p>在一次操作中，您可以选择 <code>num</code> 的任意一位数字并将其删除。请注意，如果你删除 <code>num</code> 中的所有数字，则 <code>num</code> 变为 <code>0</code>。</p>

<p>返回最少需要多少次操作可以使 <code>num</code> 变成特殊数字。</p>

<p>如果整数 <code>x</code> 能被 <code>25</code> 整除，则该整数 <code>x</code> 被认为是特殊数字。</p>

<p> </p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><strong>输入：</strong>num = &#34;2245047&#34;
<strong>输出：</strong>2
<strong>解释：</strong>删除数字 num[5] 和 num[6] ，得到数字 &#34;22450&#34; ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 2 位数字。</pre>

<p><strong>示例 2：</strong></p>

<pre><strong>输入：</strong>num = &#34;2908305&#34;
<strong>输出：</strong>3
<strong>解释：</strong>删除 num[3]、num[4] 和 num[6] ，得到数字 &#34;2900&#34; ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 3 位数字。</pre>

<p><strong>示例 3：</strong></p>

<pre><strong>输入：</strong>num = &#34;10&#34;
<strong>输出：</strong>1
<strong>解释：</strong>删除 num[0] ，得到数字 &#34;0&#34; ，可以被 25 整除。
可以证明要使数字变成特殊数字，最少需要删除 1 位数字。
</pre>

<p> </p>

<p><strong>提示</strong></p>

<ul>
	<li><code>1 &lt;= num.length &lt;= 100</code></li>
	<li><code>num</code> 仅由数字 <code>&#39;0&#39;</code> 到 <code>&#39;9&#39;</code> 组成</li>
	<li><code>num</code> 不含任何前导零</li>
</ul>
 
#### 思路  

一个数能被 $25$ 整除，有如下五种情况：
- 这个数是 $0$。
- 这个数的末尾是 $00$。
- 这个数的末尾是 $25$。
- 这个数的末尾是 $50$。
- 这个数的末尾是 $75$。
  
设 $\textit{num}$ 的长度为 $n$。  
首先，根据题目说的，我们可以把 $\textit{num}$ 中的所有数字都删除，得到 $0$，这需要删除 $n$ 次。  
但如果 $\textit{num}$ 中有 $0$，那么删除 $n-1$ 也可以得到 $0$。  
接下来，看示例 1。以 $50$ 为例说明：  
1. 从右到左遍历 $\textit{num}$，找到第一个 $0$。  
2. 继续向左遍历，找到第一个 $5$，设下标为 $i$。  
3. 删除 $i$ 右侧的所有非 $0$ 非 $5$ 数字，这样就得到了一个以 $50$ 结尾的数字。  
4. 删除次数为 $n-i-2$，例如示例 1 中 $5$ 的下标是 $i=3$，需要删除 $7-3-2=2$ 次。
   
其余 $00,25,75$ 的计算方式同理，取 $n-i-2$ 的最小值作为答案。  
注意：  
1. 如果没有找到要找的字符，则跳过。  
2. 不可能在删除后只得到 $00$，因为题目保证 $\textit{num}$ 不含前导零，如果有多个 $0$，那么 $0$ 的左侧必然有非 $0$ 数字。

```go 
func minimumOperations(num string) int {
	ans := len(num)
	if strings.Contains(num, "0") {
		ans-- // 可以删除 len(num)-1 次得到 "0"
	}
	f := func(tail string) {
		i := strings.LastIndexByte(num, tail[1])
		if i < 0 {
			return
		}
		i = strings.LastIndexByte(num[:i], tail[0])
		if i < 0 {
			return
		}
		ans = min(ans, len(num)-i-2)
	}
	f("00")
	f("25")
	f("50")
	f("75")
	return ans
}

func min(a, b int) int { if b < a { return b }; return a }
```

#### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{num}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。