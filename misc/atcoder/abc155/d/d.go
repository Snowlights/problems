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

	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	sort.Ints(a)
	p := sort.SearchInts(a, 0)
	q := sort.SearchInts(a, 1)

	neg := (n - q) * p
	if neg >= k {
		k = neg - k + 1
		Fprintln(out, -sort.Search(1e18, func(val int) bool {
			cnt, j := 0, q
			for _, v := range a[:p] {
				for j < n && a[j]*-v <= val {
					j++
				}
				cnt += j - q
			}
			return cnt >= k
		}))
		return
	}
	k -= neg

	zero := (n-q+p)*(q-p) + (q-p)*(q-p-1)/2
	if zero >= k {
		Fprintln(out, 0)
		return
	}
	k -= zero

	Fprintln(out, sort.Search(1e18, func(val int) bool {
		cnt := 0
		i, j := 0, p-1
		for i < j {
			if a[i]*a[j] > val {
				i++
			} else {
				cnt += j - i
				j--
			}
		}
		i, j = q, n-1
		for i < j {
			if a[i]*a[j] > val {
				j--
			} else {
				cnt += j - i
				i++
			}
		}
		return cnt >= k
	}))

}

func main() { run(os.Stdin, os.Stdout) }
