#### 题目

<p>给你一个由数字组成的字符串 <code>s</code>&nbsp;。重复执行以下操作，直到字符串恰好包含&nbsp;<strong>两个&nbsp;</strong>数字：</p>

<ul>
	<li>从第一个数字开始，对于 <code>s</code> 中的每一对连续数字，计算这两个数字的和&nbsp;<strong>模&nbsp;</strong>10。</li>
	<li>用计算得到的新数字依次替换 <code>s</code>&nbsp;的每一个字符，并保持原本的顺序。</li>
</ul>

<p>如果 <code>s</code>&nbsp;最后剩下的两个数字 <strong>相同</strong> ，返回 <code>true</code>&nbsp;。否则，返回 <code>false</code>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "3902"</span></p>

<p><strong>输出：</strong> <span class="example-io">true</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>一开始，<code>s = "3902"</code></li>
	<li>第一次操作：
	<ul>
		<li><code>(s[0] + s[1]) % 10 = (3 + 9) % 10 = 2</code></li>
		<li><code>(s[1] + s[2]) % 10 = (9 + 0) % 10 = 9</code></li>
		<li><code>(s[2] + s[3]) % 10 = (0 + 2) % 10 = 2</code></li>
		<li><code>s</code> 变为 <code>"292"</code></li>
	</ul>
	</li>
	<li>第二次操作：
	<ul>
		<li><code>(s[0] + s[1]) % 10 = (2 + 9) % 10 = 1</code></li>
		<li><code>(s[1] + s[2]) % 10 = (9 + 2) % 10 = 1</code></li>
		<li><code>s</code> 变为 <code>"11"</code></li>
	</ul>
	</li>
	<li>由于 <code>"11"</code> 中的数字相同，输出为 <code>true</code>。</li>
</ul>
</div>

<p><strong class="example">示例 2：</strong></p>

<div class="example-block">
<p><strong>输入：</strong> <span class="example-io">s = "34789"</span></p>

<p><strong>输出：</strong> <span class="example-io">false</span></p>

<p><strong>解释：</strong></p>

<ul>
	<li>一开始，<code>s = "34789"</code>。</li>
	<li>第一次操作后，<code>s = "7157"</code>。</li>
	<li>第二次操作后，<code>s = "862"</code>。</li>
	<li>第三次操作后，<code>s = "48"</code>。</li>
	<li>由于 <code>'4' != '8'</code>，输出为 <code>false</code>。</li>
</ul>

<p>&nbsp;</p>
</div>

<p><strong>提示：</strong></p>

<ul>
	<li><code>3 &lt;= s.length &lt;= 100</code></li>
	<li><code>s</code> 仅由数字组成。</li>
</ul>

#### 思路

## 公式推导

先不考虑取模，考察最终的元素和是由哪些元素相加得到的。

把字符串视作一个数字数组。例如数组为 $[a,b,c]$，那么操作一次后变成 $[a+ b, b+ c]$，再操作一次（假设操作到只剩一个数），得到 $a+ 2b+ c$。

又例如数组为 $[a,b,c,d]$，那么操作一次后变成 $[a+b,b+c,c+d]$，再操作一次，变成 $[a+2b+c,b+2c+d]$，再操作一次，得到 $a+3b+3c+d$。

又例如数组为 $[a,b,c,d,e]$，最后可以操作得到 $a+4b+6c+4d+e$。

可以发现，在最终结果中，原数组的第 $i$ 个元素的系数（出现次数）是一个组合数

$$
\dbinom {m-1} {i}
$$

其中 $m$ 是数组长度。

所以本题相当于计算 $s[0]$ 到 $s[n-2]$ 和 $s[1]$ 到 $s[n-1]$ 这两个数组的如下结果：（把数组记作 $a$，长度为 $m=n-1$）

$$
\sum_{i=0}^{m-1} \dbinom {m-1} {i} \cdot a[i]
$$

也可以计算两个数组的结果的差值

$$
\sum_{i=0}^{n-2} \dbinom {n-2} {i} \cdot (s[i] - s[i+1])
$$

判断上式模 $10$ 的结果是否为 $0$。

## 方法一：提取因子 2 和 5 + 欧拉定理

关于组合数取模，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

本题由于模数不是质数，计算逆元无法用费马小定理。怎么办？

把每个数中的因子 $2$ 和 $5$ 单独提取出来，单独统计因子个数。一个数去掉其中所有因子 $2$ 和 $5$ 之后，和 $10$ 互质，这样可以用**欧拉定理**计算整数 $a$ 在模 $10$ 下的逆元，即 $a^{\varphi(10)-1} = a^3$。


### 优化前

```
const mod = 10

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

const mx = 100_000

var f, invF, p2, p5 [mx + 1]int

func init() {
	f[0] = 1
	for i := 1; i <= mx; i++ {
		x := i
		// 2 的幂次
		e2 := bits.TrailingZeros(uint(x))
		x >>= e2
		// 5 的幂次
		e5 := 0
		for x%5 == 0 {
			e5++
			x /= 5
		}
		f[i] = f[i-1] * x % mod
		p2[i] = p2[i-1] + e2
		p5[i] = p5[i-1] + e5
	}

	invF[mx] = pow(f[mx], 3) // 欧拉定理
	for i := mx; i > 0; i-- {
		x := i
		x >>= bits.TrailingZeros(uint(x))
		for x%5 == 0 {
			x /= 5
		}
		invF[i-1] = invF[i] * x % mod
	}
}

func comb(n, k int) int {
	// 由于每项都 < 10，所以无需中途取模
	return f[n] * invF[k] * invF[n-k] *
		pow(2, p2[n]-p2[k]-p2[n-k]) *
		pow(5, p5[n]-p5[k]-p5[n-k])
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%mod == 0
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n\log U)$，其中 $n$ 是 $s$ 的长度，$\log U$ 是计算快速幂的时间。
- 空间复杂度：$\mathcal{O}(1)$。

### 优化

预处理 $2$ 的幂模 $10$ 和 $5$ 的幂模 $10$。

由于 $2^i\bmod 10\ (i>0)$ 按照 $2,4,8,6$ 的周期循环，只需预处理一个长为 $4$ 的数组。

当 $i>0$ 时，$5^i\bmod 10 = 5$ 恒成立，所以无需预处理。

```
const mod = 10

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

const mx = 100_000

var f, invF, p2, p5 [mx + 1]int

func init() {
	f[0] = 1
	for i := 1; i <= mx; i++ {
		x := i
		// 2 的幂次
		e2 := bits.TrailingZeros(uint(x))
		x >>= e2
		// 5 的幂次
		e5 := 0
		for x%5 == 0 {
			e5++
			x /= 5
		}
		f[i] = f[i-1] * x % mod
		p2[i] = p2[i-1] + e2
		p5[i] = p5[i-1] + e5
	}

	invF[mx] = pow(f[mx], 3) // 欧拉定理
	for i := mx; i > 0; i-- {
		x := i
		x >>= bits.TrailingZeros(uint(x))
		for x%5 == 0 {
			x /= 5
		}
		invF[i-1] = invF[i] * x % mod
	}
}

var pow2 = [4]int{2, 4, 8, 6}

func comb(n, k int) int {
	res := f[n] * invF[k] * invF[n-k]
	e2 := p2[n] - p2[k] - p2[n-k]
	if e2 > 0 {
		res *= pow2[(e2-1)%4]
	}
	if p5[n]-p5[k]-p5[n-k] > 0 {
		res *= 5
	}
	return res
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%mod == 0
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：扩展 Lucas

证明见 [Lucas 定理](https://oi-wiki.org/math/number-theory/lucas/)。

Lucas 定理可以帮助我们快速计算 $r_2 = \dbinom {n} {k}\bmod 2$ 和 $r_5 = \dbinom {n} {k}\bmod 5$。

根据 [中国剩余定理](https://oi-wiki.org/math/number-theory/crt/) 的计算方法，知道了 $r_2$ 和 $r_5$，那么有

$$
\dbinom {n} {k}\bmod 10 = (5r_2+6r_5)\bmod 10
$$

```
const mx = 5

var c [mx][mx]int

func init() {
    // 预处理组合数
	for i := range mx {
		c[i][0], c[i][i] = 1, 1
		for j := 1; j < i; j++ {
			c[i][j] = c[i-1][j-1] + c[i-1][j]
		}
	}
}

// 计算 C(n, k) % p，要求 p 是质数
func lucas(n, k, p int) int {
	if k == 0 {
		return 1
	}
	return c[n%p][k%p] * lucas(n/p, k/p, p) % p
}

func comb10(n, k int) int {
	// 结果至多为 5 + 4 * 6 = 29，无需中途取模
	return lucas(n, k, 2)*5 + lucas(n, k, 5)*6
}

func hasSameDigits(s string) bool {
	diff := 0
	for i := range len(s) - 1 {
		diff += comb10(len(s)-2, i) * (int(s[i]) - int(s[i+1]))
	}
	return diff%10 == 0
}
```

#### 复杂度分析

预处理的时间和空间忽略不计。

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $s$ 的长度。计算 $\texttt{lucas}$ 的时间为 $\mathcal{O}(\log_{p} n)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单


- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
- [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
- [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
- [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
- [贪心（基本贪心策略/反悔/区间/字典序/数学/思维/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
- [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
- [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)