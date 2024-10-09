package _50

import (
	"bufio"
	"fmt"
	"io"
)

func CF954A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var s string
	var n int
	fmt.Fscan(r, &n, &s)
	ans := 0
	i := 1
	for i < n {
		if s[i] != s[i-1] {
			i++
			ans++
		}
		i++
	}

	fmt.Fprintln(w, len(s)-ans)
}

func CF954B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var s string
	var n int
	fmt.Fscan(r, &n, &s)
	ans := 0
	for i := 1; i <= len(s)/2; i++ {
		if s[:i] == s[i:2*i] {
			ans = i
		}
	}
	if ans > 0 {
		fmt.Fprintln(w, len(s)-ans+1)
	} else {
		fmt.Fprintln(w, len(s))
	}
}

func CF954C(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	fmt.Fscan(r, &a[0])
	for i := 1; i < n; i++ {
		fmt.Fscan(r, &a[i])
		d := a[i] - a[i-1]
		if d < 0 {
			d = -d
		}
		if d == 0 {
			fmt.Fprint(w, "NO")
			return
		}
		if d > 1 {
			if m == 0 {
				m = d
			} else if d != m {
				fmt.Fprint(w, "NO")
				return
			}
		}
	}
	if m == 0 {
		fmt.Fprintln(w, "YES")
		fmt.Fprint(w, int(1e9), 1)
		return
	}
	for i := 1; i < n; i++ {
		if a[i] == a[i-1]+1 && a[i-1]%m == 0 || a[i] == a[i-1]-1 && a[i]%m == 0 {
			fmt.Fprint(w, "NO")
			return
		}
	}
	fmt.Fprintln(w, "YES")
	fmt.Fprint(w, int(1e9), m)
}
