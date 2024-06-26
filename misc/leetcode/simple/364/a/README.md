#### 题目

<p>给你一个 <strong>二进制</strong> 字符串 <code>s</code> ，其中至少包含一个 <code>'1'</code> 。</p>

<p>你必须按某种方式 <strong>重新排列</strong> 字符串中的位，使得到的二进制数字是可以由该组合生成的 <strong>最大二进制奇数</strong> 。</p>

<p>以字符串形式，表示并返回可以由给定组合生成的最大二进制奇数。</p>

<p><strong>注意 </strong>返回的结果字符串 <strong>可以</strong> 含前导零。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "010"
<strong>输出：</strong>"001"
<strong>解释：</strong>因为字符串 s 中仅有一个 '1' ，其必须出现在最后一位上。所以答案是 "001" 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "0101"
<strong>输出：</strong>"1001"
<strong>解释：</strong>其中一个 '1' 必须出现在最后一位上。而由剩下的数字可以生产的最大数字是 "100" 。所以答案是 "1001" 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code> 仅由 <code>'0'</code> 和 <code>'1'</code> 组成</li>
	<li><code>s</code> 中至少包含一个 <code>'1'</code></li>
</ul>

#### 思路

把一个 $1$ 放末尾，其余全部放在开头。

```go  
func maximumOddBinaryNumber(s string) string {
	cnt1 := strings.Count(s, "1")
	return strings.Repeat("1", cnt1-1) + strings.Repeat("0", len(s)-cnt1) + "1"
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。其中 $n$ 为 $\textit{s}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。