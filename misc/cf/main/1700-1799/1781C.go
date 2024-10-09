//go:build main
// +build main

package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"os"
	"sort"
)

func CF1781C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
	id := make([]int, 26)
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		for i := range id {
			id[i] = i
		}
		sort.Slice(id, func(i, j int) bool { return cnt[id[i]] > cnt[id[j]] })
		maxSave := 0
		idx := 0
		for i := 1; i <= 26; i++ {
			if n%i > 0 {
				continue
			}
			save := 0
			for _, j := range id[:i] {
				save += min(cnt[j], n/i)
			}
			if save > maxSave {
				maxSave = save
				idx = i
			}
		}
		ex := []byte{}
		for _, i := range id[:idx] {
			if cnt[i] > n/idx {
				cnt[i] = n / idx
			} else {
				ex = append(ex, bytes.Repeat([]byte{'a' + byte(i)}, n/idx-cnt[i])...)
			}
		}
		for _, i := range id[idx:] {
			cnt[i] = 0
		}
		ei := 0
		for i, b := range s {
			b -= 'a'
			if cnt[b] > 0 {
				cnt[b]--
			} else {
				s[i] = ex[ei]
				ei++
			}
		}
		Fprintln(out, n-maxSave)
		Fprintf(out, "%s\n", s)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() { CF1781C(os.Stdin, os.Stdout) }
