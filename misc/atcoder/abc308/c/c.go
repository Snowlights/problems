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

	var n int
	Fscan(in, &n)
	type pair struct {
		a, b, i int
	}
	pairList := make([]*pair, n)
	for i := range pairList {
		pairList[i] = &pair{}
		Fscan(in, &pairList[i].a, &pairList[i].b)
		pairList[i].i = i + 1
	}

	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].a*(pairList[j].a+pairList[j].b) > pairList[j].a*(pairList[i].a+pairList[i].b) ||
			pairList[i].a*(pairList[j].a+pairList[j].b) == pairList[j].a*(pairList[i].a+pairList[i].b) && pairList[i].i < pairList[j].i
	})

	for _, v := range pairList {
		Fprint(out, v.i, " ")
	}

}

func main() { run(os.Stdin, os.Stdout) }
