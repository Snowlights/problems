#### 题目

<p>给你一个下标从<strong> 0</strong> 开始的字符串 <code>s</code> ，该字符串由用户输入。按键变更的定义是：使用与上次使用的按键不同的键。例如 <code>s = "ab"</code> 表示按键变更一次，而 <code>s = "bBBb"</code> 不存在按键变更。</p>

<p>返回用户输入过程中按键变更的次数。</p>

<p><strong>注意：</strong><code>shift</code> 或 <code>caps lock</code> 等修饰键不计入按键变更，也就是说，如果用户先输入字母 <code>'a'</code> 然后输入字母 <code>'A'</code> ，不算作按键变更。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "aAbBcC"
<strong>输出：</strong>2
<strong>解释：</strong> 
从 s[0] = 'a' 到 s[1] = 'A'，不存在按键变更，因为不计入 caps lock 或 shift 。
从 s[1] = 'A' 到 s[2] = 'b'，按键变更。
从 s[2] = 'b' 到 s[3] = 'B'，不存在按键变更，因为不计入 caps lock 或 shift 。
从 s[3] = 'B' 到 s[4] = 'c'，按键变更。
从 s[4] = 'c' 到 s[5] = 'C'，不存在按键变更，因为不计入 caps lock 或 shift 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "AaAaAaaA"
<strong>输出：</strong>0
<strong>解释：</strong> 不存在按键变更，因为这个过程中只按下字母 'a' 和 'A' ，不需要进行按键变更。<!-- notionvc: 8849fe75-f31e-41dc-a2e0-b7d33d8427d2 -->
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= s.length <= 100</code></li>
	<li><code>s</code> 仅由英文大写字母和小写字母组成。</li>
</ul>

#### 思路

由于同一字母的大小写，ASCII 码的低 $5$ 位是相同的，我们只需要统计

$$
s[i-1] \& 31 \ne s[i]\& 31

$$

的个数。其中 $\& 31$ 表示取二进制的低 $5$ 位。

```go [sol]
func countKeyChanges(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		if s[i-1]&31 != s[i]&31 {
			ans++
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。
