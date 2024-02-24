//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// 首先，令 x=d^k * s，其中 k > 0,d | != s。
//那么 k-1 个 d 和一个 d\times s 即为一种分解方案。
//接下来考虑构造另一种。
// 1.若 k=1，则显然没有第二种方案。
//则接下来的讨论中均有 k>1。
// 2.若 s 为合数，则将其分解为两个数，分别乘在两个 d 上，则得到第二种方案。
//则接下来的讨论中 s 均为质数或 1。
//那么 s 拆不动了，只能拆 d。
// 3.若 d 为质数，那么它就没得拆，所以没有第二种方案。
//则现在剩下的情况是 k>1，s 不是合数，d 是合数。
// 4.若 d 中含有不同于 s 的质因数，则将这个质因数乘在一个 d 上，剩下的部分与 s 一起乘在另一个 d 上，则得到第二种方案，连上被拆的一共需要 3 个 d，所以需要 k>2。
//那么注意到，接下来的讨论中有 d=s^q，其中 q>1。
// 5.若 q>2，则 d 可以拆为 s 和 s^{q-1}。将 s^2 和 s^{q-1} 分别乘在一个 d 上，则得到第二种方案，连上被拆的一共需要 3 个 d，所以需要 k>2。
// 6.若 q=2，则 d 可以拆为两个 s。现在不能将 s^2 乘在 d 上了（因为结果不是漂亮数），所以只能将三个 s 分别乘在三个 d 上，连上被拆的一共需要 4 个 d，所以需要 k>3。
//注意到后三类可以整合为对 s^2 与 d 是否相等的判断。

func CF1647D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, x, d int
	const (
		Yes = "YES"
		No  = "NO"
	)

	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &d)
		check := func(n int) int {
			if n < 4 {
				return 1
			}
			if n%2 == 0 || n%3 == 0 {
				return 0
			}
			for i := 5; i*i <= n; i += 6 {
				if n%i == 0 || n%(i+2) == 0 {
					return 0
				}
			}
			return 1
		}
		b2i := func(i bool) int {
			if i {
				return 1
			}
			return 0
		}
		f := func() string {
			k := 0
			for ; x%d == 0; x /= d {
				k++
			}
			if k < 2 {
				return No
			}
			if check(x) == 0 {
				return Yes
			}
			if check(d) == 1 {
				return No
			}
			if k > b2i(x*x == d)+2 {
				return Yes
			}
			return No
		}

		Fprintln(out, f())
	}

}

func main() { CF1647D(os.Stdin, os.Stdout) }
