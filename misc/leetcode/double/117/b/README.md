### 题目

<p>给你两个正整数 <code>n</code> 和 <code>limit</code> 。</p>

<p>请你将 <code>n</code> 颗糖果分给 <code>3</code> 位小朋友，确保没有任何小朋友得到超过 <code>limit</code> 颗糖果，请你返回满足此条件下的 <strong>总方案数</strong> 。</p>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>n = 5, limit = 2
<b>输出：</b>3
<b>解释：</b>总共有 3 种方法分配 5 颗糖果，且每位小朋友的糖果数不超过 2 ：(1, 2, 2) ，(2, 1, 2) 和 (2, 2, 1) 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>n = 3, limit = 3
<b>输出：</b>10
<b>解释：</b>总共有 10 种方法分配 3 颗糖果，且每位小朋友的糖果数不超过 3 ：(0, 0, 3) ，(0, 1, 2) ，(0, 2, 1) ，(0, 3, 0) ，(1, 0, 2) ，(1, 1, 1) ，(1, 2, 0) ，(2, 0, 1) ，(2, 1, 0) 和 (3, 0, 0) 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= n <= 10<sup>6</sup></code></li>
	<li><code>1 <= limit <= 10<sup>6</sup></code></li>
</ul>

### 思路

枚举第一个人得到的糖果、分情况讨论:剩余数量设为 $x$

- 剩余的糖果 $x < limit$, 方案数为 $x+1$
- 剩余的糖果 $x <= 2*limit$, 方案数为 $2*limit - x + 1$
- 剩余的糖果 $x > 2*limit$, 方案数为 $0$

```go
func distributeCandies(n int, limit int) int64 {
	var ans int64
	for i := 0; i <= limit && i <= n; i++ {
		x := n - i
		if x < limit {
			ans += int64(x+1)
		} else if x <= 2 *limit {
			ans += int64(2 *limit -x + 1)
		}
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。
