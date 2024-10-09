package c_168

import (
	"bufio"
	"fmt"
	"io"
)

// https://atcoder.jp/contests/abc168/tasks/abc168_e
//
//输入 n(≤2e5) 和 n 个点 (xi, yi)，范围 [-1e18,1e18]。
//你需要从这 n 个点中选出一个非空子集，满足子集中任意两点都有 xi*xj+yi*yj ≠ 0。
//子集的大小可以为 1。
//输出有多少种不同的选法，模 1e9+7。
//
//注意：可能有重复的点。

// https://atcoder.jp/contests/abc168/submissions/36333766
//把点看成向量，公式看成向量不能垂直。
//根据对称性，可以把在 x 轴下方或 y 轴负半轴的向量，按原点对称。
//然后分别统计在坐标原点的、在第一象限或 x 正半轴的（集合 P）、在第二象限或 y 正半轴的（集合 Q)，
//其中 P 和 Q 是有可能垂直的，而 P Q 内部的向量是不会垂直的。
//P 中的每个向量和其在 Q 中垂直的向量是不能同时选的，把这些找出来，当成一组，计算方案数。
//具体见代码。
//根据乘法原理。每组的方案数可以相乘。
//最后统计 Q 中剩余向量的方案数；以及零向量的方案数，由于零向量只能选一个，所以方案数是 cnt0；
//别忘了去掉一个都不选的方案。

const mod int = 1e9 + 7

func AtCoderABC173F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, x, y, cntO, left int
	type pair struct{ x, y int }
	cnt := map[pair]int{}
	cnt2 := map[pair]int{}
	for fmt.Fscan(in, &n); n > 0; n-- {
		fmt.Fscan(in, &x, &y)
		if x == 0 && y == 0 {
			cntO++
			continue
		}
		// 标准化
		if y < 0 {
			x = -x
			y = -y
		} else if y == 0 {
			x = abs(x)
		}
		g := gcd(abs(x), y)
		x /= g
		y /= g
		if x > 0 {
			cnt[pair{x, y}]++ // 在第一象限或 x 正半轴的向量
		} else {
			cnt2[pair{y, -x}]++ // 在第二象限或 y 正半轴的向量 => 有可能与 cnt 中的向量垂直
			left++
		}
	}

	ans := 1
	for p, c := range cnt {
		ans = ans * (pow(2, c) + pow(2, cnt2[p]) - 1) % mod // 单独选 p + 单独选与 p 垂直的 // （注意空集重复统计了一次）
		left -= cnt2[p]
	}
	fmt.Fprint(out, (ans*pow(2, left)+cntO-1+mod)%mod) // 最后单独选 cnt2 中的，-1 表示减去空集
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
