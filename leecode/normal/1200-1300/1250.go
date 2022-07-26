package _200_1300

// 1254
//0 表示陆地，用 1 表示海水：
func closedIsland(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		//递归出口  越界 或者遇到海水
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 1 {
			return
		}
		//递归出口  已经遍历过(i,j)
		if visited[i][j] {
			return
		}
		//遍历位置(i,j)
		visited[i][j] = true
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	//首先把四周给遍历了
	for j := 0; j < n; j++ {
		dfs(0, j)   //最上面
		dfs(m-1, j) //最下面
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)   //最左边
		dfs(i, n-1) //最右边
	}
	//遍历二维数组
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//满足条件
			if grid[i][j] == 0 && visited[i][j] == false {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}
