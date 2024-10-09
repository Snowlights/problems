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

	var n, ans int
	Fscan(in, &n)

	a, b := make([]int, n), make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := range b {
		Fscan(in, &b[i])
	}

	for k := 0; k < 29; k++ {
		mask := 1<<(k+1) - 1
		sort.Slice(a, func(i, j int) bool { return a[i]&mask < a[j]&mask })
		sort.Slice(b, func(i, j int) bool { return b[i]&mask < b[j]&mask })
		cnt := 0
		i, j, p := n-1, n-1, n-1
		l1, r1, l2 := 1<<k, 1<<(k+1)-1, 3<<k
		for _, v := range a {
			for i >= 0 && v&mask+b[i]&mask >= l1 {
				i--
			}
			for j >= 0 && v&mask+b[j]&mask > r1 {
				j--
			}
			for p >= 0 && v&mask+b[p]&mask >= l2 {
				p--
			}
			cnt ^= i ^ j ^ p
		}
		println(cnt)
		ans |= cnt & 1 << k
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
