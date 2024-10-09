package _000

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CF1016A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, a, curm int64
	fmt.Fscan(r, &n, &m)
	curm = m
	for ; n > 0; n-- {
		fmt.Fscan(r, &a)
		p := a / m
		a = a % m
		if a > 0 {
			if a >= curm {
				a -= curm
				p++
				curm = m - a
			} else {
				curm -= a
			}
		}
		fmt.Fprint(w, p, " ")
	}

}

func CF1016B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, q, left, right int64
	var s, t string
	fmt.Fscan(r, &n, &m, &q, &s, &t)
	cnt := make([]int64, n+1)
	for i := int64(0); i <= n-m; i++ {
		if s[i:i+m] == t {
			cnt[i+1] = cnt[i] + 1
		} else {
			cnt[i+1] = cnt[i]
		}
	}

	for ; q > 0; q-- {
		fmt.Fscan(r, &left, &right)
		if right-left+1 < m {
			fmt.Fprintln(w, 0)
			continue
		}
		fmt.Fprintln(w, cnt[right-m+1]-cnt[left-1])
	}

}

func CF1016D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, xorA, xorB int64
	fmt.Fscan(r, &n, &m)
	a := make([]int64, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
		xorA ^= a[i]
	}
	b := make([]int64, m)
	for i := range b {
		fmt.Fscan(r, &b[i])
		xorB ^= b[i]
	}

	if xorA != xorB {
		fmt.Fprint(w, "NO")
		return
	}
	fmt.Fprintln(w, "YES")
	fmt.Fprint(w, xorA^a[0]^b[0]) // 左上角
	for _, x := range b[1:] {
		fmt.Fprint(w, " ", x)
	}
	fmt.Fprintln(w)
	zeros := strings.Repeat("0 ", int(m-1))
	for _, x := range a[1:] {
		fmt.Fprintln(w, x, zeros)
	}
}
