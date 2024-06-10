#### 题目

<p>给你一个二进制字符串 <code>s</code> 和一个正整数 <code>k</code> 。</p>

<p>如果 <code>s</code> 的某个子字符串中 <code>1</code> 的个数恰好等于 <code>k</code> ，则称这个子字符串是一个 <strong>美丽子字符串</strong> 。</p>

<p>令 <code>len</code> 等于 <strong>最短</strong> 美丽子字符串的长度。</p>

<p>返回长度等于 <code>len</code> 且字典序 <strong>最小</strong> 的美丽子字符串。如果 <code>s</code> 中不含美丽子字符串，则返回一个 <strong>空</strong> 字符串。</p>

<p>对于相同长度的两个字符串 <code>a</code> 和 <code>b</code> ，如果在 <code>a</code> 和 <code>b</code> 出现不同的第一个位置上，<code>a</code> 中该位置上的字符严格大于 <code>b</code> 中的对应字符，则认为字符串 <code>a</code> 字典序 <strong>大于</strong> 字符串 <code>b</code> 。</p>

<ul>
	<li>例如，<code>"abcd"</code> 的字典序大于 <code>"abcc"</code> ，因为两个字符串出现不同的第一个位置对应第四个字符，而 <code>d</code> 大于 <code>c</code> 。</li>
</ul>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "100011001", k = 3
<strong>输出：</strong>"11001"
<strong>解释：</strong>示例中共有 7 个美丽子字符串：
1. 子字符串 "<em><strong>100011</strong></em>001" 。
2. 子字符串 "<strong><em>1000110</em></strong>01" 。
3. 子字符串 "<strong><em>100011001</em></strong>" 。
4. 子字符串 "1<strong><em>00011001</em></strong>" 。
5. 子字符串 "10<strong><em>0011001</em></strong>" 。
6. 子字符串 "100<em><strong>011001</strong></em>" 。
7. 子字符串 "1000<strong><em>11001</em></strong>" 。
最短美丽子字符串的长度是 5 。
长度为 5 且字典序最小的美丽子字符串是子字符串 "11001" 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "1011", k = 2
<strong>输出：</strong>"11"
<strong>解释：</strong>示例中共有 3 个美丽子字符串：
1. 子字符串 "<em><strong>101</strong></em>1" 。
2. 子字符串 "1<em><strong>011</strong></em>" 。
3. 子字符串 "10<em><strong>11</strong></em>" 。
最短美丽子字符串的长度是 2 。
长度为 2 且字典序最小的美丽子字符串是子字符串 "11" 。 
</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "000", k = 1
<strong>输出：</strong>""
<strong>解释：</strong>示例中不存在美丽子字符串。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 100</code></li>
	<li><code>1 &lt;= k &lt;= s.length</code></li>
</ul>


#### 思路

## 方法一：枚举

首先，如果 $s$ 中 $1$ 的个数不足 $k$，直接返回空串。

否则一定有解。

从 $k$ 开始枚举答案的长度 $\textit{size}$，然后在 $s$ 中枚举所有长为 $\textit{size}$ 的子串，同时维护字典序最小的子串。如果存在一个子串，其中 $1$ 的个数等于 $k$，则返回字典序最小的子串。

```go [sol-Go]
func shortestBeautifulSubstring(s string, k int) string {
	// 1 的个数不足 k
	if strings.Count(s, "1") < k {
		return ""
	}
	// 否则一定有解
	for size := k; ; size++ {
		ans := ""
		for i := size; i <= len(s); i++ {
			t := s[i-size : i]
			if (ans == "" || t < ans) && strings.Count(t, "1") == k {
				ans = t
			}
		}
		if ans != "" {
			return ans
		}
	}
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^3)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。字符串切片需要 $\mathcal{O}(n)$ 的空间，Go 除外。

## 方法二：滑动窗口

由于答案中恰好有 $k$ 个 $1$，我们也可以用滑动窗口找最短的答案。

如果窗口内的 $1$ 的个数超过 $k$，或者窗口端点是 $0$，就可以缩小窗口。

> 注：利用字符串哈希（或者后缀数组等），可以把比较字典序的时间降至 $\mathcal{O}(n\log n)$，这样可以做到 $\mathcal{O}(n\log n)$ 的时间复杂度。

```go [sol-Go]
func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans, cnt, left := s, 0, 0
	for right, b := range s {
		cnt += int(b & 1)
		for cnt > k || s[left] == '0' {
			cnt -= int(s[left] & 1)
			left++
		}
		if cnt == k {
			t := s[left : right+1]
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 为 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。字符串切片需要 $\mathcal{O}(n)$ 的空间，Go 除外。
