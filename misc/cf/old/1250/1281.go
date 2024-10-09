package _250

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func CF1281A(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const (
		po       = "po"
		FILIPINO = "FILIPINO"
		desu     = "desu"
		masu     = "masu"
		JAPANESE = "JAPANESE"
		mnida    = "mnida"
		KOREAN   = "KOREAN"
	)

	var T int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		var s string
		fmt.Fscan(r, &s)
		if strings.HasSuffix(s, po) {
			fmt.Fprintln(w, FILIPINO)
		} else if strings.HasSuffix(s, desu) || strings.HasSuffix(s, masu) {
			fmt.Fprintln(w, JAPANESE)
		} else if strings.HasSuffix(s, mnida) {
			fmt.Fprintln(w, KOREAN)
		}
	}

}

func CF1281B(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	var s, t []byte
	fmt.Fscan(r, &n)
	for ; n > 0; n-- {
		fmt.Fscan(r, &s, &t)
		last := make(map[byte]int)
		for i, b := range s {
			last[b] = i
		}
		for i, v := range s {
			if v == 'A' {
				continue
			}
			exit, idx := false, i
			for j := 'A'; j <= 'Z'; j++ {
				if last[byte(j)] > i && byte(j) < byte(v) {
					exit = true
					idx = last[byte(j)]
					break
				}
			}
			if exit {
				s[i], s[idx] = s[idx], s[i]
				break
			}
		}
		// fmt.Println(string(s), string(t))
		if string(s) < string(t) {
			fmt.Fprintln(w, string(s))
		} else {
			fmt.Fprintln(w, "---")
		}
	}

}

func CF1281C(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	mod := int64(1e9 + 7)
	var T int64
	for fmt.Fscan(r, &T); T > 0; T-- {
		var x int64
		var s []byte
		fmt.Fscan(r, &x, &s)
		sLen := int64(len(s))
		for i := int64(0); i < x; i++ {
			for j := int64(0); j < int64(s[i]-'1'); j++ {
				for k := i + 1; int64(len(s)) < x && k < int64(sLen); k++ {
					s = append(s, s[k])
				}
			}
			sLen = (sLen + (sLen-(i+1)+mod)*int64((s[i]-'1'))) % mod
		}
		fmt.Fprintln(w, sLen)
	}

}
