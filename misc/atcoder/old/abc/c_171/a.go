package c_171

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc171/tasks/abc171_f
//
//输入 k(≤1e6) 和一个长度不超过 1e6 的字符串 s，由小写字母组成。
//你需要在 s 中插入恰好 k 个小写字母。
//输出你能得到的字符串的个数，模 1e9+7。

// https://atcoder.jp/contests/abc171/submissions/36296507
//设 s 的长度为 n。
//提示 1：如何避免重复统计？做一个规定，插入在 s[i] 左侧的字符，
// 不能和 s[i] 相同，这不会影响答案的正确性。
//提示 2：枚举最后一个字符的右侧插入了多少个字符，设为 i，
// 这些字符没有限制，有 26^i 种方案。
//提示 3：剩下 (n-1) + (k-i) 个字符，我们需要考虑其中 n-1 个字符的位置，
// 这就是 C(n-1+k-i, n-1)。
//提示 4：其余插入字符的方案数就是 25^(k-i)。
//因此答案为 ∑26^i * C(n-1+k-i, n-1) * 25^(k-i), i=[0,k]
//不知道组合数怎么算的，需要学一下逆元。

func AtCoderABC171F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const mod int64 = 1e9 + 7
	const mx int = 2e6
	// 阶乘
	pow := func(x, n int64) int64 {
		//x %= mod
		res := int64(1)
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}
	F := [mx + 1]int64{1}
	for i := 1; i <= mx; i++ {
		F[i] = F[i-1] * int64(i) % mod
	}
	invF := [...]int64{mx: pow(F[mx], mod-2)}
	for i := mx; i > 0; i-- {
		invF[i-1] = invF[i] * int64(i) % mod
	}
	C := func(n, k int) int64 {
		if k < 0 || k > n {
			return 0
		}
		return F[n] * invF[k] % mod * invF[n-k] % mod
	}

	var k, ans int64
	var s string
	fmt.Fscan(r, &k, &s)
	n := int64(len(s))
	p25, p26 := pow(25, int64(k)), int64(1)
	invF25 := pow(25, mod-2)
	for i := int64(0); i <= k; i++ {
		ans = (ans + p26*C(int(n-1+k-i), int(n-1))%mod*p25%mod) % mod
		p26 = p26 * 26 % mod
		p25 = p25 * invF25 % mod
	}
	fmt.Fprintln(w, ans)
}
