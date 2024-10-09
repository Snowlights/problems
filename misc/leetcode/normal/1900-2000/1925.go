package _900_2000

import (
	"sort"
	"strings"
)

// 1926
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	out, vis := make([][]bool, m), make([][]bool, m)
	for i := range out {
		out[i] = make([]bool, n)
		vis[i] = make([]bool, n)
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for i := 0; i < m; i++ {
		if maze[i][0] == '.' {
			out[i][0] = true
		}
		if maze[i][n-1] == '.' {
			out[i][n-1] = true
		}
	}
	for i := 0; i < n; i++ {
		if maze[0][i] == '.' {
			out[0][i] = true
		}
		if maze[m-1][i] == '.' {
			out[m-1][i] = true
		}
	}

	q, l := [][]int{entrance}, 0
	vis[entrance[0]][entrance[1]] = true
	out[entrance[0]][entrance[1]] = false
	for len(q) > 0 {
		t := q
		q = nil
		for _, v := range t {
			x, y := v[0], v[1]
			if out[x][y] {
				return l
			}
			for _, d := range dir {
				xx, yy := x+d[0], y+d[1]
				if 0 <= xx && xx < m && 0 <= yy && yy < n && maze[xx][yy] == '.' && !vis[xx][yy] {
					q = append(q, []int{xx, yy})
					vis[xx][yy] = true
				}
			}
		}
		l++
	}

	return -1
}

// 1948
type folder struct {
	son map[string]*folder
	val string // 文件夹名称
	del bool   // 删除标记
}

func deleteDuplicateFolder(paths [][]string) (ans [][]string) {
	root := &folder{}
	for _, path := range paths {
		// 将 path 加入字典树
		f := root
		for _, s := range path {
			if f.son == nil {
				f.son = map[string]*folder{}
			}
			if f.son[s] == nil {
				f.son[s] = &folder{}
			}
			f = f.son[s]
			f.val = s
		}
	}

	folders := map[string][]*folder{} // 存储括号表达式及其对应的文件夹节点列表
	var dfs func(*folder) string
	dfs = func(f *folder) string {
		if f.son == nil {
			return "(" + f.val + ")"
		}
		expr := make([]string, 0, len(f.son))
		for _, son := range f.son {
			expr = append(expr, dfs(son))
		}
		sort.Strings(expr)
		subTreeExpr := strings.Join(expr, "") // 按字典序拼接所有子树
		folders[subTreeExpr] = append(folders[subTreeExpr], f)
		return "(" + f.val + subTreeExpr + ")"
	}
	dfs(root)

	for _, fs := range folders {
		if len(fs) > 1 { // 将括号表达式对应的节点个数大于 1 的节点全部删除
			for _, f := range fs {
				f.del = true
			}
		}
	}

	// 再次 DFS 这颗字典树，仅访问未被删除的节点，并将路径记录到答案中
	path := []string{}
	var dfs2 func(*folder)
	dfs2 = func(f *folder) {
		if f.del {
			return
		}
		path = append(path, f.val)
		ans = append(ans, append([]string(nil), path...))
		for _, son := range f.son {
			dfs2(son)
		}
		path = path[:len(path)-1]
	}
	for _, son := range root.son {
		dfs2(son)
	}
	return
}
