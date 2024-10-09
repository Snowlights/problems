package _350

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

func CF1368A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		var a, b, c int64
		fmt.Fscan(r, &a, &b, &c)
		if a < b {
			a, b = b, a
		}
		ans := 0
		for a <= c {
			t := a
			a += b
			b = t
			ans++
		}
		fmt.Fprintln(w, ans)
	}
}

func CF1368B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var n int64
	fmt.Fscan(r, &n)

	s := "codeforces"
	arr := make([]int64, 10)
	for i := 0; i < 10; i++ {
		arr[i] = 1
	}
	idx := 0
	for {
		now := int64(1)
		for i := range arr {
			now *= arr[i]
		}
		if now >= n {
			break
		}
		arr[idx%10]++
		idx++
	}
	for i, v := range arr {
		fmt.Fprint(w, string(bytes.Repeat([]byte{byte(s[i])}, int(v))))
	}
}

func CF1368C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var n int64
	fmt.Fscan(r, &n)

	fmt.Fprintln(w, 3*n+4)
	fmt.Fprintln(w, 0, " ", 0)
	for i := int64(1); i <= n+1; i++ {
		fmt.Fprintln(w, i, " ", i-1)
		fmt.Fprintln(w, i-1, " ", i)
		fmt.Fprintln(w, i, " ", i)
	}
}

func CF1368D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var n int64
	fmt.Fscan(r, &n)
	arr := make([]int64, n)
	b := make([]int64, 20)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(r, &arr[i])
		for j := range b {
			b[j] += arr[i] >> j & 1
		}
	}

	ans := int64(0)
	for ; n > 0; n-- {
		v := 0
		for j := range b {
			if b[j] > 0 {
				v |= 1 << j
				b[j]--
			}
		}
		ans += int64(v * v)
	}
	fmt.Fprintln(w, ans)
}
