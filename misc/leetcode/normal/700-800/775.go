package _00_800

import (
	"math"
	"math/bits"
	"strings"
)

// 782
func getMoves(mask uint, count, n int) int {
	ones := bits.OnesCount(mask)
	if n&1 > 0 {
		// 如果 n 为奇数，则每一行中 1 与 0 的数目相差为 1，且满足相邻行交替
		if abs(n-2*ones) != 1 || abs(n-2*count) != 1 {
			return -1
		}
		if ones == n>>1 {
			// 偶数位变为 1 的最小交换次数
			return n/2 - bits.OnesCount(mask&0xAAAAAAAA)
		} else {
			// 奇数位变为 1 的最小交换次数
			return (n+1)/2 - bits.OnesCount(mask&0x55555555)
		}
	} else {
		// 如果 n 为偶数，则每一行中 1 与 0 的数目相等，且满足相邻行交替
		if ones != n>>1 || count != n>>1 {
			return -1
		}
		// 偶数位变为 1 的最小交换次数
		count0 := n/2 - bits.OnesCount(mask&0xAAAAAAAA)
		// 奇数位变为 1 的最小交换次数
		count1 := n/2 - bits.OnesCount(mask&0x55555555)
		return min(count0, count1)
	}
}

func movesToChessboard(board [][]int) int {
	n := len(board)
	// 棋盘的第一行与第一列
	rowMask, colMask := 0, 0
	for i := 0; i < n; i++ {
		rowMask |= board[0][i] << i
		colMask |= board[i][0] << i
	}
	reverseRowMask := 1<<n - 1 ^ rowMask
	reverseColMask := 1<<n - 1 ^ colMask
	rowCnt, colCnt := 0, 0
	for i := 0; i < n; i++ {
		currRowMask, currColMask := 0, 0
		for j := 0; j < n; j++ {
			currRowMask |= board[i][j] << j
			currColMask |= board[j][i] << j
		}
		if currRowMask != rowMask && currRowMask != reverseRowMask || // 检测每一行的状态是否合法
			currColMask != colMask && currColMask != reverseColMask { // 检测每一列的状态是否合法
			return -1
		}
		if currRowMask == rowMask {
			rowCnt++ // 记录与第一行相同的行数
		}
		if currColMask == colMask {
			colCnt++ // 记录与第一列相同的列数
		}
	}
	rowMoves := getMoves(uint(rowMask), rowCnt, n)
	colMoves := getMoves(uint(colMask), colCnt, n)
	if rowMoves == -1 || colMoves == -1 {
		return -1
	}
	return rowMoves + colMoves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 784
func letterCasePermutation(s string) []string {
	ans := make([]string, 0)
	var bfs func(str string, i int)
	bfs = func(str string, i int) {
		if len(str) == len(s) {
			ans = append(ans, str)
			return
		}
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			bfs(str+strings.ToLower(string(s[i])), i+1)
			bfs(str+strings.ToUpper(string(s[i])), i+1)
		} else {
			bfs(str+string(s[i]), i+1)
		}
	}
	bfs("", 0)
	return ans
}

// 785
const (
	UNCOLORED, RED, GREEN = 0, 1, 2
)

func isBipartite(graph [][]int) bool {
	n := len(graph)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		if color[i] == UNCOLORED {
			q := []int{i}
			color[i] = RED
			for len(q) > 0 {
				tmp := q
				q = nil
				for _, v := range tmp {
					cNei := RED
					if color[v] == RED {
						cNei = GREEN
					}
					for _, vv := range graph[v] {
						if color[vv] == UNCOLORED {
							q = append(q, vv)
							color[vv] = cNei
						} else if color[vv] != cNei {
							return false
						}
					}
				}
			}
		}
	}
	return true
}

// 787
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	if n < 1 {
		return -1
	}
	matrix := make([][]int, n)
	save := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		save[i] = make([]int, K+1)
	}

	for i := 0; i < len(flights); i++ {
		s, e, v := flights[i][0], flights[i][1], flights[i][2]
		matrix[s][e] = v
	}

	res := math.MaxInt32
	queue := make([]int, 0, 2*n)
	queue = append(queue, src)
	step := 0

	for len(queue) > 0 && step <= K {
		preLen := len(queue)
		for j := 0; j < preLen; j++ {
			s := queue[0]
			queue = queue[1:]
			for i := 0; i < n; i++ {
				if matrix[s][i] <= 0 {
					continue
				}
				if step == 0 {
					save[i][step] = 0 + matrix[s][i]
				} else {
					s := save[s][step-1] + matrix[s][i]
					if s < save[i][step] || save[i][step] == 0 {
						save[i][step] = s
					} else {
						continue
					}
				}
				if save[i][step] > res {
					continue
				}

				if i == dst {
					res = save[i][step]
				}
				queue = append(queue, i)
			}
		}
		step++
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

// 792
func numMatchingSubseq(s string, words []string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	nxt := make([][26]int, n)
	for i := n - 1; i >= 0; i-- {
		nxt[i] = pos
		pos[s[i]-'a'] = i
	}

	ans := 0
	check := func(word string) bool {
		i, j := 0, 0
		if word[0] == s[0] {
			j = 1
		}
		for ; j < len(word); j++ {
			next := nxt[i][word[j]-'a']
			if next == n {
				return false
			}
			i = next
		}
		return true
	}
	for _, w := range words {
		if check(w) {
			ans++
		}
	}
	return ans
}

// 797
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans := [][]int{}

	var bfs func([]int)
	bfs = func(res []int) {
		val := res[len(res)-1]
		for _, v := range graph[val] {
			if v == n-1 {
				ans = append(ans, append([]int{}, append(res, v)...))
			}
			bfs(append(res, v))
		}
	}
	bfs([]int{0})
	return ans
}
