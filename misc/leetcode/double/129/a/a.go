package main

func canMakeSquare(grid [][]byte) bool {
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 1; j++ {
			bCnt := 0
			if grid[i][j] == 'B' {
				bCnt++
			}
			if grid[i][j+1] == 'B' {
				bCnt++
			}
			if grid[i+1][j] == 'B' {
				bCnt++
			}
			if grid[i+1][j+1] == 'B' {
				bCnt++
			}
			if bCnt <= 1 || bCnt >= 3 {
				return true
			}
		}
	}
	return false
}
