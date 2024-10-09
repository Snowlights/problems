package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewScanner(_r)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// 支持负数的写法
	r := func() (x int) {
		in.Scan()
		data := in.Bytes()
		if data[0] == '-' {
			for _, b := range data[1:] {
				x = x*10 + int(b&15)
			}
			return -x
		}
		for _, b := range data {
			x = x*10 + int(b&15)
		}
		return
	}

	var T, n, k, ans, i, j int
	var a []int
	for T = r(); T > 0; T-- {
		n, k = r(), r()
		a, ans = make([]int, n), 0
		for i := range a {
			a[i] = r()
			ans += a[i] / k
			a[i] %= k
		}
		sort.Ints(a)
		i, j = 0, n-1
		for i < j {
			if a[i]+a[j] >= k {
				ans++
				j--
			}
			i++
		}
		Fprintln(out, ans)
	}

}

func main() { run(os.Stdin, os.Stdout) }
