#### 题目

<p>给你一个长度为 <strong>偶数</strong> <code>n</code> ，下标从 <strong>0</strong> 开始的字符串 <code>s</code> 。</p>

<p>同时给你一个下标从 <strong>0</strong> 开始的二维整数数组 <code>queries</code> ，其中 <code>queries[i] = [a<sub>i</sub>, b<sub>i</sub>, c<sub>i</sub>, d<sub>i</sub>]</code> 。</p>

<p>对于每个查询 <code>i</code> ，你需要执行以下操作：</p>

<ul>
	<li>将下标在范围 <code>0 <= a<sub>i</sub> <= b<sub>i</sub> < n / 2</code> 内的 <strong>子字符串</strong> <code>s[a<sub>i</sub>:b<sub>i</sub>]</code> 中的字符重新排列。</li>
	<li>将下标在范围 <code>n / 2 <= c<sub>i</sub> <= d<sub>i</sub> < n</code> 内的 <strong>子字符串</strong> <code>s[c<sub>i</sub>:d<sub>i</sub>]</code> 中的字符重新排列。</li>
</ul>

<p>对于每个查询，你的任务是判断执行操作后能否让 <code>s</code> 变成一个 <strong>回文串</strong> 。</p>

<p>每个查询与其他查询都是 <strong>独立的</strong> 。</p>

<p>请你返回一个下标从 <strong>0</strong> 开始的数组<em> </em><code>answer</code> ，如果第 <code>i</code> 个查询执行操作后，可以将 <code>s</code> 变为一个回文串，那么<em> </em><code>answer[i] = true</code>，否则为<em> </em><code>false</code> 。</p>

<ul>
	<li><strong>子字符串</strong> 指的是一个字符串中一段连续的字符序列。</li>
	<li><code>s[x:y]</code> 表示 <code>s</code> 中从下标 <code>x</code> 到 <code>y</code> 且两个端点 <strong>都包含</strong> 的子字符串。</li>
</ul>

<p> </p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<b>输入：</b>s = "abcabc", queries = [[1,1,3,5],[0,2,5,5]]
<b>输出：</b>[true,true]
<b>解释：</b>这个例子中，有 2 个查询：
第一个查询：
- a<sub>0</sub> = 1, b<sub>0</sub> = 1, c<sub>0</sub> = 3, d<sub>0</sub> = 5
- 你可以重新排列 s[1:1] => a<em><strong>b</strong></em>cabc 和 s[3:5] => abc<em><strong>abc</strong></em> 。
- 为了让 s 变为回文串，s[3:5] 可以重新排列得到 => abc<strong><em>cba </em></strong>。
- 现在 s 是一个回文串。所以 answer[0] = true 。
第二个查询：
- a<sub>1</sub> = 0, b<sub>1</sub> = 2, c<sub>1</sub> = 5, d<sub>1</sub> = 5.
- 你可以重新排列 s[0:2] => <em><strong>abc</strong></em>abc 和 s[5:5] => abcab<strong><em>c</em></strong> 。
- 为了让 s 变为回文串，s[0:2] 可以重新排列得到 => <em><strong>cba</strong></em>abc 。
- 现在 s 是一个回文串，所以 answer[1] = true 。
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<b>输入：</b>s = "abbcdecbba", queries = [[0,2,7,9]]
<b>输出：</b>[false]
<b>解释：</b>这个示例中，只有一个查询。
a<sub>0</sub> = 0, b<sub>0</sub> = 2, c<sub>0</sub> = 7, d<sub>0</sub> = 9.
你可以重新排列 s[0:2] => <em><strong>abb</strong></em>cdecbba 和 s[7:9] => abbcdec<em><strong>bba</strong></em> 。
无法通过重新排列这些子字符串使 s 变为一个回文串，因为 s[3:6] 不是一个回文串。
所以 answer[0] = false 。</pre>

<p><strong class="example">示例 3：</strong></p>

<pre>
<b>输入：</b>s = "acbcab", queries = [[1,2,4,5]]
<b>输出：</b>[true]
<strong>解释：</strong>这个示例中，只有一个查询。
a<sub>0</sub> = 1, b<sub>0</sub> = 2, c<sub>0</sub> = 4, d<sub>0</sub> = 5.
你可以重新排列 s[1:2] => a<em><strong>cb</strong></em>cab 和 s[4:5] => acbc<strong><em>ab</em></strong> 。
为了让 s 变为回文串，s[1:2] 可以重新排列得到 => a<em><strong>bc</strong></em>cab<code> </code>。
然后 s[4:5] 重新排列得到 abcc<em><strong>ba</strong></em> 。
现在 s 是一个回文串，所以 answer[0] = true 。</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= n == s.length <= 10<sup>5</sup></code></li>
	<li><code>1 <= queries.length <= 10<sup>5</sup></code></li>
	<li><code>queries[i].length == 4</code></li>
	<li><code>a<sub>i</sub> == queries[i][0], b<sub>i</sub> == queries[i][1]</code></li>
	<li><code>c<sub>i</sub> == queries[i][2], d<sub>i</sub> == queries[i][3]</code></li>
	<li><code>0 <= a<sub>i</sub> <= b<sub>i</sub> < n / 2</code></li>
	<li><code>n / 2 <= c<sub>i</sub> <= d<sub>i</sub> < n </code></li>
	<li><code>n</code> 是一个偶数。</li>
	<li><code>s</code> 只包含小写英文字母。</li>
</ul>

#### 思路


把输入的字符串记作大写 $S$。为方便计算，把 $S$ 均分为左右两个字符串， 其中右半部分反转。左半和右半分别记作 $s$ 和 $t$。
我们需要预处理三个前缀和：
- $s$ 的每个前缀的每个字符的出现次数，记作 $\textit{sumS}$。
- $t$ 的每个前缀的每个字符的出现次数，记作 $\textit{sumT}$。
- 前缀中 $s[i] \ne t[i]$ 的次数，记作 $\textit{sumNe}$（Not Equal）。

由于 $S$ 的右半部分反转了，所以询问中的 $c$ 和 $d$ 要替换成 $m-1-d$ 和 $m-1-c$，其中 $m$ 是 $S$ 的长度。
为方便描述，把 $a,b,c,d$ 改称为 $l_1, r_1, l_2, r_2$。不失一般性，假设 $l_1 \le l_2$（如果不成立就交换），分类讨论：
- 首先，$[0,l_1-1]$ 和 $[\max(r_1,r_2)+1,n-1]$ 这两个区间，区间内的每个下标 $i$ 必须满足 $s[i]=t[i]$， 这可以用 $\textit{sumNe}$ 判断。
- 如果**区间包含**，即 $r_2 \le r_1$。这是最简单的情况。由于在 $[l_1,r_1]$ 中 $s$ 可以随意排列，所以甚至不需要排列 $t$ 中的字符，我们只需要判断 $s$ 和 $t$ 在 $[l_1,r_1]$ 中每个字符的出现次数是否相等即可。
- 如果**区间不相交**，即 $r_1 < l_2$。先判断 $[r_1+1,l_2-1]$ 内的每个下标 $i$ 是否满足 $s[i]=t[i]$，不满足直接返回 `false`。然后算出 $s$ 的 $[l_1,r_1]$ 中每个字符的出现次数，这必须与 $t$ 中 $[l_1,r_1]$ 中每个字符的出现次数相等。同样地，$t$ 的 $[l_2,r_2]$ 中每个字符的出现次数，必须与 $s$ 中 $[l_2,r_2]$ 中每个字符的出现次数相等。
- 如果**区间相交但不包含**，即 $l_2 \le r_1$。先算出 $s$ 的 $[l_1,r_1]$ 中每个字符的出现次数，减去 $t$ 中 $[l_1,l_2-1]$ 中每个字符的出现次数。然后算出 $t$ 中 $[l_2,r2]$ 中每个字符的出现次数，减去 $s$ 中 $[r_1+1,r_2]$ 中每个字符的出现次数。最后判断出现次数不能为负数，且剩余的 $s$ 和 $t$ 中的字符出现次数必须相等。

```go  [sol]
func canMakePalindromeQueries(s string, queries [][]int) []bool {
	// 分成左右两半，右半反转
	n := len(s) / 2
	t := []byte(s[n:])
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
	s = s[:n]

	// 预处理三种前缀和
	sumS := make([][26]int, n+1)
	for i, b := range s {
		sumS[i+1] = sumS[i]
		sumS[i+1][b-'a']++
	}

	sumT := make([][26]int, n+1)
	for i, b := range t {
		sumT[i+1] = sumT[i]
		sumT[i+1][b-'a']++
	}

	sumNe := make([]int, n+1)
	for i := range s {
		sumNe[i+1] = sumNe[i]
		if s[i] != t[i] {
			sumNe[i+1]++
		}
	}

	// 计算子串中各个字符的出现次数，闭区间 [l,r]
	count := func(sum [][26]int, l, r int) []int {
		res := sum[r+1]
		for i, s := range sum[l][:] {
			res[i] -= s
		}
		return res[:]
	}

	subtract := func(s1, s2 []int) []int {
		for i, s := range s2 {
			s1[i] -= s
			if s1[i] < 0 {
				return nil
			}
		}
		return s1
	}
	equal := func(s1, s2 []int) bool {
		for i, s := range s1 {
			if s != s2[i] {
				return false
			}
		}
		return true
	}

	check := func(l1, r1, l2, r2 int, sumS, sumT [][26]int) bool {
		if sumNe[l1] > 0 || // [0,l1-1] 有 s[i] != t[i]
			sumNe[n]-sumNe[max(r1, r2)+1] > 0 { // [max(r1,r2)+1,n-1] 有 s[i] != t[i]
			return false
		}
		if r2 <= r1 { // 区间包含
			return equal(count(sumS, l1, r1), count(sumT, l1, r1))
		}
		if r1 < l2 { // 区间不相交
			return sumNe[l2]-sumNe[r1+1] == 0 && // [r1+1,l2-1] 都满足 s[i] == t[i]
				equal(count(sumS, l1, r1), count(sumT, l1, r1)) &&
				equal(count(sumS, l2, r2), count(sumT, l2, r2))
		}
		// 区间相交但不包含
		s1 := subtract(count(sumS, l1, r1), count(sumT, l1, l2-1))
		s2 := subtract(count(sumT, l2, r2), count(sumS, r1+1, r2))
		return s1 != nil && s2 != nil && equal(s1, s2)
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		l1, r1, l2, r2 := q[0], q[1], n*2-1-q[3], n*2-1-q[2]
		if l1 <= l2 {
			ans[i] = check(l1, r1, l2, r2, sumS, sumT)
		} else {
			ans[i] = check(l2, r2, l1, r1, sumT, sumS)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+q)|\Sigma|)$，其中 $n$ 为 $s$ 的长度，$q$ 为 $\textit{queries}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。回答每个询问的时间是 $\mathcal{O}(|\Sigma|)$。
- 空间复杂度：$\mathcal{O}(n|\Sigma|)$。返回值的空间不计入。