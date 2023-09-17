### 题目

<p>给你一个字符串 <code>s</code> 和一个整数 <code>k</code> 。</p>

<p><strong>k 子序列</strong>指的是 <code>s</code> 的一个长度为 <code>k</code> 的 <strong>子序列</strong> ，且所有字符都是 <strong>唯一</strong> 的，也就是说每个字符在子序列里只出现过一次。</p>

<p>定义 <code>f(c)</code> 为字符 <code>c</code> 在 <code>s</code> 中出现的次数。</p>

<p>k 子序列的 <strong>美丽值</strong> 定义为这个子序列中每一个字符 <code>c</code> 的 <code>f(c)</code> 之 <strong>和</strong> 。</p>

<p>比方说，<code>s = "abbbdd"</code> 和 <code>k = 2</code> ，我们有：</p>

<ul>
	<li><code>f('a') = 1</code>, <code>f('b') = 3</code>, <code>f('d') = 2</code></li>
	<li><code>s</code> 的部分 k 子序列为：
	<ul>
		<li><code>"<em><strong>ab</strong></em>bbdd"</code> -> <code>"ab"</code> ，美丽值为 <code>f('a') + f('b') = 4</code></li>
		<li><code>"<em><strong>a</strong></em>bbb<em><strong>d</strong></em>d"</code> -> <code>"ad"</code> ，美丽值为 <code>f('a') + f('d') = 3</code></li>
		<li><code>"a<em><strong>b</strong></em>bb<em><strong>d</strong></em>d"</code> -> <code>"bd"</code> ，美丽值为 <code>f('b') + f('d') = 5</code></li>
	</ul>
	</li>
</ul>

<p>请你返回一个整数，表示所有 <strong>k 子序列 </strong>里面 <strong>美丽值 </strong>是 <strong>最大值</strong> 的子序列数目。由于答案可能很大，将结果对 <code>10<sup>9</sup> + 7</code> 取余后返回。</p>

<p>一个字符串的子序列指的是从原字符串里面删除一些字符（也可能一个字符也不删除），不改变剩下字符顺序连接得到的新字符串。</p>

<p><strong>注意：</strong></p>

<ul>
	<li><code>f(c)</code> 指的是字符 <code>c</code> 在字符串 <code>s</code> 的出现次数，不是在 k 子序列里的出现次数。</li>
	<li>两个 k 子序列如果有任何一个字符在原字符串中的下标不同，则它们是两个不同的子序列。所以两个不同的 k 子序列可能产生相同的字符串。</li>
</ul>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre><b>输入：</b>s = "bcca", k = 2
<b>输出：</b>4
<b>解释：</b><span style="white-space: normal">s 中我们有 f('a') = 1 ，f('b') = 1 和 f('c') = 2 。</span>
s 的 k 子序列为：
<em><strong>bc</strong></em>ca ，美丽值为 f('b') + f('c') = 3
<em><strong>b</strong></em>c<em><strong>c</strong></em>a ，美丽值为 f('b') + f('c') = 3
<em><strong>b</strong></em>cc<em><strong>a</strong></em> ，美丽值为 f('b') + f('a') = 2
b<em><strong>c</strong></em>c<em><strong>a</strong></em><strong> </strong>，美丽值为 f('c') + f('a') = 3
bc<em><strong>ca</strong></em> ，美丽值为 f('c') + f('a') = 3
总共有 4 个 k 子序列美丽值为最大值 3 。
所以答案为 4 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre><b>输入：</b>s = "abbcd", k = 4
<b>输出：</b>2
<b>解释：</b><span style="white-space: normal">s 中我们有 f('a') = 1 ，f('b') = 2 ，f('c') = 1 和</span> f('d') = 1 。
s 的 k 子序列为：
<em><strong>ab</strong></em>b<em><strong>cd</strong></em> ，美丽值为 f('a') + f('b') + f('c') + f('d') = 5
<span style="white-space: normal;"><b><i>a</i></b></span>b<em><strong>bcd</strong></em> ，美丽值为 f('a') + f('b') + f('c') + f('d') = 5 
总共有 2 个 k 子序列美丽值为最大值 5 。
所以答案为 2 。
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= s.length <= 2 * 10<sup>5</sup></code></li>
	<li><code>1 <= k <= s.length</code></li>
	<li><code>s</code> 只包含小写英文字母。</li>
</ul>

### 思路

如果 $k >= 26$ 或者 $s$ 中去重的字符个数 $< k$ 则结果必然位0，否则必然存在:

剩下的结果中，前 $m$ 个取 最大的几个，
如果最后取得值有多个($count$个)则从 $count$ 个数中取 $n-m$ 个数

```go  
const (
	mod int = 1e9 + 7
	mx      = 30
)

func countKSubsequencesWithMaxBeauty(s string, k int) (ans int) {

	F := [mx + 1]int{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int(i) % mod
	}
	// 阶乘
	pow := func(x, n int) int {
		//x %= mod
		res := int(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	// n阶乘的逆元
	invF := [...]int{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int(i) % mod
	}
	// C(n, k) 组合：n中取k，有多少种取法
	C := func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	cnt, cc := make([]int, 26), make(map[int]bool)
	for _, v := range s {
		cnt[int(v-'a')]++
		cc[int(v-'a')] = true
	}
	if k > 26 || len(cc) < k {
		return 0
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	val, count := cnt[k-1], 0
	ans = 1
	for _, v := range cnt {
		if v == val {
			count++
			continue
		}
		if v < val {
			break
		}
		k--
		ans = ans * v % mod
	}

	ans = ans * C(count, k) % mod * pow(val, k) % mod

	ans = (ans%mod + mod) % mod
	return
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$ 。
