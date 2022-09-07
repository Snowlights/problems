package _00_600

// 556
func matrixReshape(mat [][]int, r int, c int) [][]int {
	row, col := len(mat), len(mat[0])
	if row*col != r*c {
		return mat
	}
	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}
	for i := 0; i < row*col; i++ {
		ans[i/c][i%c] = mat[i/col][i%col]
	}
	return ans
}
