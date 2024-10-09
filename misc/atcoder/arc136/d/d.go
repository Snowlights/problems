package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, ans int
	Fscan(in, &n)
	s, a := [11][11][11][11][11][11]int{}, make([]int, n)
	p := make([]int, 10)
	for i := range a {
		Fscan(in, &a[i])
		for x, j := a[i], 1; j <= 6; j++ {
			p[j] = x%10 + 1
			x /= 10
		}
		s[p[1]][p[2]][p[3]][p[4]][p[5]][p[6]]++
	}

	equal := func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	}

	for i0 := 1; i0 <= 6; i0++ {
		for i1 := 1; i1 <= 10; i1++ {
			for i2 := 1; i2 <= 10; i2++ {
				for i3 := 1; i3 <= 10; i3++ {
					for i4 := 1; i4 <= 10; i4++ {
						for i5 := 1; i5 <= 10; i5++ {
							for i6 := 1; i6 <= 10; i6++ {
								s[i1][i2][i3][i4][i5][i6] +=
									s[i1-equal(i0, 1)][i2-equal(i0, 2)][i3-equal(i0, 3)][i4-equal(i0, 4)][i5-equal(i0, 5)][i6-equal(i0, 6)]
							}
						}
					}
				}
			}
		}
	}

	btoi := func(f bool) int {
		if f {
			return 1
		}
		return 0
	}

	for i := 0; i < n; i++ {
		f := true
		for x, j := a[i], 1; j <= 6; j++ {
			p[j] = x % 10
			f = f && (p[j] < 5)
			x /= 10
		}
		ans += s[10-p[1]][10-p[2]][10-p[3]][10-p[4]][10-p[5]][10-p[6]]
		ans -= btoi(f)
	}

	Fprintln(out, ans/2)
}

func main() { run(os.Stdin, os.Stdout) }
