package _50

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func CF175A(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var s []byte
	fmt.Fscan(r, &s)
	if len(s) < 3 {
		fmt.Fprintln(w, -1)
		return
	}

	const (
		oneMillion = 1000000
	)

	ans := int64(-1)
	for i := 1; i <= len(s)-2 && i <= 7; i++ {
		for j := i + 1; j <= len(s)-1 && j <= i+7; j++ {
			if len(s)-j > 7 {
				continue
			}
			if s[0] == '0' && i > 1 {
				continue
			}
			if s[i] == '0' && j > i+1 {
				continue
			}
			if s[j] == '0' && len(s) > j+1 {
				continue
			}
			a, _ := strconv.Atoi(string(s[:i]))
			b, _ := strconv.Atoi(string(s[i:j]))
			c, _ := strconv.Atoi(string(s[j:]))
			if a <= oneMillion && b <= oneMillion && c <= oneMillion {
				ans = max(ans, int64(a+b+c))
			}
		}
	}
	fmt.Fprintln(w, ans)
}

func CF175B(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	type node struct {
		name  string
		score int64
	}
	var n int64
	fmt.Fscan(r, &n)
	nodeList, nodeMap := make([]*node, 0, n), make(map[string]*node, n)
	for ; n > 0; n-- {
		var name string
		var score int64
		fmt.Fscan(r, &name, &score)
		no, ok := nodeMap[name]
		if !ok {
			no = &node{
				name: name,
			}
			nodeMap[name] = no
			nodeList = append(nodeList, no)
		}
		no.score = max(no.score, score)
	}

	sort.Slice(nodeList, func(i, j int) bool { return nodeList[i].score > nodeList[j].score })
	ll := len(nodeList)
	fmt.Fprintln(w, ll)
	for _, node := range nodeList {
		s := 0
		for _, m := range nodeList {
			if node.score >= m.score {
				s++
			}
		}
		s *= 100
		fmt.Fprint(w, node.name, " ")
		if s >= ll*99 {
			fmt.Fprintln(w, "pro")
		} else if s >= ll*90 {
			fmt.Fprintln(w, "hardcore")
		} else if s >= ll*80 {
			fmt.Fprintln(w, "average")
		} else if s >= ll*50 {
			fmt.Fprintln(w, "random")
		} else {
			fmt.Fprintln(w, "noob")
		}
	}
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func CF175C(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	type master struct {
		num, score int64
	}
	var N int64
	fmt.Fscan(r, &N)
	masterList := make([]*master, 0, N)
	for i := int64(0); i < N; i++ {
		var n, s int64
		fmt.Fscan(r, &n, &s)
		masterList = append(masterList, &master{
			num:   n,
			score: s,
		})
	}
	var T int64
	fmt.Fscan(r, &T)
	arr := make([]int64, T)
	for i := int64(0); i < T; i++ {
		fmt.Fscan(r, &arr[i])
	}
	newArr := make([]int64, T+1)
	newArr[0] = arr[0]
	for i := int64(1); i < T; i++ {
		newArr[i] = arr[i] - arr[i-1]
	}
	newArr[T] = 1e12

	sort.Slice(masterList, func(i, j int) bool {
		return masterList[i].score < masterList[j].score
	})

	ans, idx := int64(0), int64(0)
	for _, m := range masterList {
		for m.num > 0 && idx < T+1 {
			v := min(m.num, newArr[idx])
			ans += v * m.score * (idx + 1)
			newArr[idx] -= v
			if newArr[idx] == 0 {
				idx++
			}
			m.num -= v
		}
		if idx == T+1 {
			break
		}
	}
	fmt.Fprintln(w, ans)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
