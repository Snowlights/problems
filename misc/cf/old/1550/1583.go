package _550

import (
	"bufio"
	"fmt"
	"io"
)

func CF1583A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var n int
		fmt.Fscan(r, &n)
		arr := make([]int, n)
		var total int
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &arr[i])
			total += int(arr[i])
		}
		found := 0
		for j := 2; j*j <= total; j++ {
			if total%j == 0 {
				found = 1
				break
			}
		}
		fmt.Fprintln(w, n+found-1)
		for i, v := range arr {
			if v%2 == 1 && found == 0 {
				found = 1
			} else {
				fmt.Fprint(w, i+1, " ")
			}
		}
		fmt.Fprintln(w)
	}
}

func CF1583B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T int
	for fmt.Fscan(r, &T); T > 0; T-- {
		var n, m int
		fmt.Fscan(r, &n, &m)
		v := make([]int, n+1)
		for i := 0; i < m; i++ {
			var a, b, c int
			fmt.Fscan(r, &a, &b, &c)
			v[b] = 1
		}
		for i := 1; i <= n; i++ {
			if v[i] == 0 {
				for j := 1; j <= n; j++ {
					if i != j {
						fmt.Fprintln(w, i, " ", j)
					}
				}
				break
			}
		}
	}
}
