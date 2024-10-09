### 题目  

<p>给你一个三位数整数 <code>n</code> 。</p>

<p>如果经过以下修改得到的数字 <strong>恰好</strong> 包含数字 <code>1</code> 到 <code>9</code> 各一次且不包含任何 <code>0</code> ，那么我们称数字 <code>n</code> 是 <strong>迷人的</strong> ：</p>

<ul>
	<li>将 <code>n</code> 与数字 <code>2 * n</code> 和 <code>3 * n</code> <strong>连接</strong> 。</li>
</ul>

<p>如果 <code>n</code> 是迷人的，返回 <code>true</code>，否则返回 <code>false</code> 。</p>

<p><strong>连接</strong> 两个数字表示把它们首尾相接连在一起。比方说 <code>121</code> 和 <code>371</code> 连接得到 <code>121371</code> 。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre><b>输入：</b>n = 192
<b>输出：</b>true
<b>解释：</b>我们将数字 n = 192 ，2 * n = 384 和 3 * n = 576 连接，得到 192384576 。这个数字包含 1 到 9 恰好各一次。
</pre>

<p><strong>示例 2：</strong></p>

<pre><b>输入：</b>n = 100
<b>输出：</b>false
<b>解释：</b>我们将数字 n = 100 ，2 * n = 200 和 3 * n = 300 连接，得到 100200300 。这个数字不符合上述条件。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>100 &lt;= n &lt;= 999</code></li>
</ul>
 
### 思路  

根据题意，$3n$ 必须是一个三位数，即 $3n\le 999$，所以 $n\le 333$。    
用一位数记录 $n$ 、 $2*n$ 、 $3*n$ 出现的数字
和 $1<<10 - 2$ (111111110) 判断是否相等

```go 
func isFascinating(n int) (ans bool) {

	if n > 333 {
		return false
	}

	mask := 0
	for _, s := range []string{strconv.Itoa(n), strconv.Itoa(n * 2), strconv.Itoa(n * 3)} {
		for _, v := range s {
			mask |= 1 << (v - '0')
		}
	}

	return mask == 1<<10-2
}
```

### 复杂度分析  

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$ 。 仅用到部分变量