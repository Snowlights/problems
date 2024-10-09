### 题目

<p>给你一个正整数 <code>num</code> ，请你将它分割成两个非负整数 <code>num1</code> 和 <code>num2</code> ，满足：</p>

<ul>
	<li><code>num1</code> 和 <code>num2</code> 直接连起来，得到 <code>num</code> 各数位的一个排列。

<ul>
	<li>换句话说，<code>num1</code> 和 <code>num2</code> 中所有数字出现的次数之和等于 <code>num</code> 中所有数字出现的次数。</li>
</ul>
</li>
<li><code>num1</code> 和 <code>num2</code> 可以包含前导 0 。</li>
</ul>

<p>请你返回 <code>num1</code> 和 <code>num2</code> 可以得到的和的 <strong>最小</strong> 值。</p>

<p><strong>注意：</strong></p>

<ul>
	<li><code>num</code> 保证没有前导 0 。</li>
	<li><code>num1</code> 和 <code>num2</code> 中数位顺序可以与 <code>num</code> 中数位顺序不同。</li>
</ul>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>num = 4325
<b>输出：</b>59
<b>解释：</b>我们可以将 4325 分割成 <code>num1 </code>= 24 和 num2<code> = </code>35 ，和为 59 ，59 是最小和。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>num = 687
<b>输出：</b>75
<b>解释：</b>我们可以将 687 分割成 <code>num1</code> = 68 和 <code>num2 </code>= 7 ，和为最优值 75 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>10 <= num <= 10<sup>9</sup></code></li>
</ul>

### 思路

思考：$4325$ 怎么分？  
如果分成 $432$ 和 $5$，那么 $4$ 显然放在 $5$ 这边更优，所以要**均匀分**。  
如果分成 $32$ 和 $45$，那么 $23$ 比 $32$ 更好；进一步地，分成 $24$ 和 $35$ 更好，所以**把小的排在高位更优，大的排在低位更优**。  
设 $s$ 是 $\textit{num}$ 的字符串形式，这等价于把 $s$ 排序后，按照奇偶下标分组。  

```go  
func splitNum(n int) (ans int) {
	s := []byte(strconv.Itoa(n))
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	nums := [2]int{}
	for i, v := range s {
		nums[i%2] = nums[i%2]*10 + int(v-'0')
	}
	return nums[0] + nums[1]
}
```
### 复杂度分析

- 时间复杂度：$O(m\log m)$，其中 $m$ 为 $\textit{num}$ 转成字符串后的长度。
- 空间复杂度：$O(m)$。
