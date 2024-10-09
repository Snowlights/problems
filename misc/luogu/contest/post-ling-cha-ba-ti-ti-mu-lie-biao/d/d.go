package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func() {
		var n, ans int
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			a[i] += a[i-1]
		}
		for i := 0; i < 25; i++ {
			zero := []int{}
			one := []int{}
			mask, cnt := 1<<i-1, 0
			for _, v := range a {
				m := v & mask
				sort.Ints(one)
				sort.Ints(zero)
				if v&(1<<i) > 0 {
					cnt += len(one) - sort.SearchInts(one, m+1)
					cnt += sort.SearchInts(zero, m+1)
					one = append(one, m)
				} else {
					cnt += sort.SearchInts(one, m+1)
					cnt += len(zero) - sort.SearchInts(zero, m+1)
					zero = append(zero, m)
				}
			}
			if cnt%2 == 1 {
				ans |= 1 << i
			}
		}

		Fprint(out, ans)
	}

	solve()

}

func main() { run(os.Stdin, os.Stdout) }
