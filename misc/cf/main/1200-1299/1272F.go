//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"strings"
)

type Node struct {
	i, j, k, nl int
	flag        bool
}

var (
	dp  = [205][205][405]int{}
	pre = [205][205][405]*Node{}
)

// Memory limit exceeded
func CF1272F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)

	lens, lent := len(s), len(t)
	s, t = " "+s, " "+t
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = math.MaxInt32
			}
		}
	}

	dp[0][0][0] = 0
	for i := 0; i <= lens; i++ {
		for j := 0; j <= lent; j++ {
			for k := 0; k <= lens+lent; k++ {
				if dp[i][j][k] == math.MaxInt32 {
					continue
				}

				// 加个右括号
				if k > 0 {
					sf, tf := 0, 0
					if i < lens && s[i+1] == ')' {
						sf = 1
					}
					if j < lent && t[j+1] == ')' {
						tf = 1
					}
					if dp[i][j][k]+1 < dp[i+sf][j+tf][k-1] {
						node := &Node{
							i:    i,
							j:    j,
							k:    k,
							nl:   dp[i][j][k],
							flag: false,
						}
						pre[i+sf][j+tf][k-1] = node
						dp[i+sf][j+tf][k-1] = dp[i][j][k] + 1
					}
				}
				// 加个左括号

				sf, tf := 0, 0
				if i < lens && s[i+1] == '(' {
					sf = 1
				}
				if j < lent && t[j+1] == '(' {
					tf = 1
				}
				if dp[i][j][k]+1 < dp[i+sf][j+tf][k+1] {
					node := &Node{
						i:    i,
						j:    j,
						k:    k,
						nl:   dp[i][j][k],
						flag: true,
					}
					pre[i+sf][j+tf][k+1] = node
					dp[i+sf][j+tf][k+1] = dp[i][j][k] + 1
				}
			}

		}
	}

	mink := 0
	for k := 1; k <= lens+lent; k++ {
		if dp[lens][lent][k]+k < dp[lens][lent][mink]+mink {
			mink = k
		}
	}

	var buf []string
	var nodePrint func(node *Node)
	nodePrint = func(node *Node) {
		if node.flag {
			buf = append(buf, "(")
		} else {
			buf = append(buf, ")")
		}
		if node.nl > 0 {
			nodePrint(pre[node.i][node.j][node.k])
		}
	}
	for i := 0; i < mink; i++ {
		buf = append(buf, ")")
	}
	nodePrint(pre[lens][lent][mink])
	for i := 0; i < len(buf)/2; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}
	Fprintln(out, strings.Join(buf, ""))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CF1272FV2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s, t string
	Fscan(in, &s, &t)
	lens, lent := len(s), len(t)
	s += "$"
	t += "$"
	C, c := "()", []int{1, -1}
	// dp[i,j,k]表示 s 匹配到 i , t 匹配到 j , 栈中有 k 个左括号，的最小新序列长度
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	flag := func(f bool) int {
		if f {
			return 1
		}
		return 0
	}
	var S func(i, j, k int) int
	S = func(i, j, k int) int {
		if k > 200 || k < 0 {
			return 1e9
		}

		if k == 0 && i == lens && j == lent {
			return 0
		}
		dv := &dp[i][j][k]
		if *dv != -1 {
			return *dv
		}
		res := int(1e9)
		*dv = 1e9
		for x := 0; x < 2; x++ {
			res = min(res, 1+S(i+flag(C[x] == s[i]), j+flag(C[x] == t[j]), k+c[x]))
		}
		*dv = res
		return res
	}
	var P func(i, j, k int)
	P = func(i, j, k int) {
		if k == 0 && i == lens && j == lent {
			return
		}
		r := S(i, j, k)
		for x := 0; x < 2; x++ {
			if r == 1+S(i+flag(C[x] == s[i]), j+flag(C[x] == t[j]), k+c[x]) {
				Fprintf(out, string(C[x]))
				P(i+flag(C[x] == s[i]), j+flag(C[x] == t[j]), k+c[x])
				break
			}
		}
	}
	P(0, 0, 0)
}

func main() { CF1272FV2(os.Stdin, os.Stdout) }
