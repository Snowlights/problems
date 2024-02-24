//go:build main
// +build main

package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
	"runtime/debug"
)

func init() { debug.SetGCPercent(-1) }

type trieNode struct {
	son [2]*trieNode
	cnt int32
}

type trie struct{ root *trieNode }

const trieBitLen = 32

func (t *trie) put(v int) *trieNode {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		o.cnt++
	}
	return o
}

func (t *trie) maxXorKth(v int, k int32) (ans int) {
	o := t.root
	for i := trieBitLen - 1; i >= 0; i-- {
		b := v >> i & 1
		if o.son[b^1] != nil {
			if k <= o.son[b^1].cnt {
				ans |= 1 << i
				b ^= 1
			} else {
				k -= o.son[b^1].cnt
			}
		}
		o = o.son[b]
	}
	return
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, ans int
	Fscan(in, &n, &k)
	a := make([]int, n+1)
	t := &trie{&trieNode{}}
	t.put(0)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		a[i] ^= a[i-1]
		t.put(a[i])
	}

	h := make(hp, n+1)
	for i, v := range a {
		h[i] = tuple{t.maxXorKth(v, 1), i, 1}
	}
	heap.Init(&h)

	for k *= 2; k > 0; k-- {
		p := &h[0]
		ans += p.xor
		p.k++
		p.xor = t.maxXorKth(a[p.i], p.k)
		heap.Fix(&h, 0)
	}
	Fprint(out, ans/2)
}

func main() { run(os.Stdin, os.Stdout) }

type tuple struct {
	xor, i int
	k      int32
}
type hp []tuple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].xor > h[j].xor }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }
